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

}
