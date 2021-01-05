package sfpxm_rpc

import (
	"errors"
	"io"
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
	sending  sync.Mutex // protect following
	header   codec.Header
	mu       sync.Mutex // protect following
	seq      uint64
	pending  map[uint64]*Call
	closing  bool // user close directly
	shutdown bool // server tell
}

var errShoutDown = errors.New("connection is shout down")

// Close implement io close
func (c *Client) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	// user primitive shoutdown
	if c.closing {
		return errShoutDown
	}
	c.closing = true
	return c.cc.Close()

}

var _ io.Closer = (*Client)(nil)
