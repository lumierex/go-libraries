package sfpxm_rpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"sfpxm-rpc/codec"
	"sync"
)

// func(t* T) MethodName(argType T1, replayType* T2) error
type Call struct {
	Seq           uint64      // request sequence represent one request
	ServiceMethod string      // format "<service>.<method>"
	Args          interface{} // arguments to the func
	Reply         interface{} // replay from the funct
	Error         error       // if error occurs, it will be set
	Done          chan *Call
}

func (call *Call) done() {
	call.Done <- call
}

type Client struct {
	cc       codec.Codec
	opt      *Option
	sending  sync.Mutex       // protect following
	header   codec.Header     //
	mu       sync.Mutex       // protect following
	seq      uint64           // request number sequence
	pending  map[uint64]*Call // pending store unfinished call
	closing  bool             // user close directly
	shutdown bool             // server tell
}

var ErrShoutDown = errors.New("connection is shout down")

var _ io.Closer = (*Client)(nil)

// Close implement io close
func (c *Client) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	// unproactive closing
	// 非主动关闭
	if c.closing {
		return ErrShoutDown
	}
	c.closing = true
	return c.cc.Close()
}

// IsAvailable check if client does work
func (c *Client) IsAvailable() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return !c.shutdown && !c.closing
}

// registerCall 往pending 上进行注册
// success return seq number
func (c *Client) registerCall(call *Call) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// still not register and be shoutdown
	if c.closing || c.shutdown {
		return 0, ErrShoutDown
	}
	call.Seq = c.seq
	c.pending[call.Seq] = call
	c.seq++
	return call.Seq, nil
}

// removeCall remove call from client pending
// success return call which was be remove
func (c *Client) removeCall(seq uint64) *Call {
	// map 非线程安全的？
	c.mu.Lock()
	defer c.mu.Unlock()
	call := c.pending[seq]
	delete(c.pending, seq)
	return call
}

func (c *Client) terminateCalls(err error) {
	c.sending.Lock()
	defer c.sending.Unlock()
	c.mu.Lock()
	defer c.mu.Unlock()

	c.shutdown = true
	for _, call := range c.pending {
		// https://geektutu.com/post/geerpc-day2.html
		// take error message to tell all call instance which is pending
		// call err where is call err
		call.Error = err
		call.done()
	}
}

// receive client receive message
// for client receive response has three condition
// 1. call dont exist maybe be canceled or dont send complete and server still handle
// 2. server handler error h.Error(codec) not empty
// 3. call exist server start normal read reply from body
func (c *Client) receive() {

	// 1. remove call from client then get current call instance
	// 2. handle call status

	var err error
	for err == nil {
		var h codec.Header
		if err = c.cc.ReadHeader(&h); err != nil {
			break
		}

		call := c.removeCall(h.Seq)

		switch {
		case call == nil:
			err = c.cc.ReadBody(nil)
		case h.Error != "":
			// server codec error
			call.Error = fmt.Errorf(h.Error)
			err = c.cc.ReadBody(nil)
			call.done()
		default:

			err = c.cc.ReadBody(call.Reply)
			if err != nil {
				call.Error = errors.New("reading body " + err.Error())
			}
			call.done()
		}
	}

	// server error or client happen error
	// notify all pending calls error happend
	// and cancel the pending call
	c.terminateCalls(err)
}

func NewClient(conn net.Conn, opt *Option) (*Client, error) {
	// exchange protocol when client was first created

	f := codec.NewCodecFuncMap[opt.CodeType]
	if f == nil {
		// 没有对应的
		err := fmt.Errorf("invalid code type %s", opt.CodeType)
		log.Println("rpc client codec error ", err)
		return nil, err
	}

	if err := json.NewEncoder(conn).Encode(opt); err != nil {
		log.Println("rpc client: options error", err)
		_ = conn.Close()
		return nil, err
	}

	//
	return nil, nil

}

//
func newClientCodec(cc codec.Codec, opt *Option) *Client {
	client := &Client{
		seq:     1, // seq starts with 1
		cc:      cc,
		opt:     opt,
		pending: make(map[uint64]*Call),
	}

	// receive call message
	go client.receive()
	return client
}
