package sfpxm_rpc

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"reflect"
	"sfpxm-rpc/codec"
	"sync"
)

const MagicNumber = 0x3bef5c

type Option struct {
	MagicNumber int
	CodeType    codec.Type
}

var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	CodeType:    codec.GobType,
}

//  | Option{MagicNumber: xxx, CodecType: xxx} | Header{ServiceMethod ...} | Body interface{} |
//  | <------      固定 JSON 编码      ------>  | <-------   编码方式由 CodeType 决定   ------->|

//  | Option | Header1 | Body1 | Header2 | Body2 | ...

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

type request struct {
	h            *codec.Header
	argv, replyv reflect.Value
}

var DefaultServer = NewServer()

func Accept(lis net.Listener)  {
	DefaultServer.Accept(lis)
}

func (s *Server) Accept(lis net.Listener) {
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println("rpc server: accept error", err)
			return
		}
		go s.ServeConn(conn)
	}
}

func (s *Server) ServeConn(c io.ReadWriteCloser) {

	defer func() {
		// 释放请求
		_ = c.Close()
	}()
	// 1. 解析Option
	var opt Option

	if err := json.NewDecoder(c).Decode(&opt); err != nil {
		log.Println("rpc server: options decode error", err)
		return
	}

	// 2. 判断MagicNumber 是否是自定义的协议
	if opt.MagicNumber != MagicNumber {
		log.Println("rpc server: invalid magic number", opt.MagicNumber)
		return
	}
	//f, ok := codec.NewCodecFuncMap[opt.CodeType];
	// 3. 解析CodeType => 根据CodeType 进行协议的解析
	f := codec.NewCodecFuncMap[opt.CodeType]
	if f == nil {
		log.Println("rpc server: invalid code type", opt.CodeType)
		return
	}

	s.ServeCodec(f(c))
}

var invalidRequest = struct{}{}

func (s *Server) ServeCodec(cc codec.Codec) {
	// handle request is concurrent
	// 回复报文逐个发送
	sending := new(sync.Mutex)
	// make sure all request are handled
	wg := new(sync.WaitGroup)

	for {
		// 1. 读取请求
		req, err := s.readRequest(cc)
		if err != nil {
			// 可能recover 所以直接关掉请求
			if req == nil {
				break
			}
			//req.
			req.h.Error = err.Error()
			s.sendResponse(cc, req.h, invalidRequest, sending)
			continue
		}

		wg.Add(1)
		go s.handleRequest(cc, req, sending, wg)
	}
	wg.Wait()
	_ = cc.Close()

}

func (s *Server) readRequestHeader(cc codec.Codec) (*codec.Header, error) {
	var h codec.Header
	if err := cc.ReadHeader(&h); err != nil {
		// 非 io错误
		if err != io.EOF || err != io.ErrUnexpectedEOF {
			log.Println("rpc server: read header error ", err)
		}
		return nil, err
	}
	return &h, nil
}

func (s *Server) readRequest(cc codec.Codec) (*request, error) {

	// 1. 读请求头
	h, err := s.readRequestHeader(cc)
	if err != nil {
		return nil, err
	}

	req := &request{
		h: h,
	}
	req.argv = reflect.New(reflect.TypeOf(""))
	if err := cc.ReadBody(req.argv.Interface()); err != nil {
		log.Println("rpc server: read argv err: ", err)
	}
	return req, nil
}

func (s *Server) sendResponse(cc codec.Codec, h *codec.Header, body interface{}, sending *sync.Mutex) {
	sending.Lock()
	defer sending.Unlock()
	if err := cc.Write(h, body); err != nil {
		log.Println("rpc server: write response err", err)
	}

}

func (s *Server) handleRequest(cc codec.Codec, req *request, sending *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println(req.h, req.argv.Elem())

	//
	req.replyv = reflect.ValueOf(fmt.Sprintf("rpc resp %d", req.h.Seq))
	s.sendResponse(cc, req.h, req.replyv.Interface(), sending)
}
