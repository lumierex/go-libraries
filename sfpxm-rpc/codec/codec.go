package codec

import "io"

// Header 协议 请求 响应头
type Header struct {
	ServiceMethod string
	Seq           uint64 // request id
	Error         string
}

// Code interface for different codec implements
type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

const (
	GobType  = "application/gob"
	JsonType = "application/json"
)

// codec插件管理
var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
