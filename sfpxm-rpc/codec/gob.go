package codec

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
)

type GobCodec struct {
	conn io.ReadWriteCloser
	buf  *bufio.Writer // avoid block
	dec  *gob.Decoder
	enc  *gob.Encoder
}

// _ is Codec Type 
var _ Codec = (*GobCodec)(nil)

func NewGobCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &GobCodec{
		conn: conn,
		buf:  buf,
		dec:  gob.NewDecoder(conn),
		enc:  gob.NewEncoder(buf),
	}
}

func (c GobCodec) Close() error {
	return c.conn.Close()
}

func (c GobCodec) ReadHeader(h *Header) error {
	return c.dec.Decode(h)
}

func (c GobCodec) ReadBody(body interface{}) error {
	return c.dec.Decode(body)
}

func (c GobCodec) Write(h *Header, body interface{}) (err error) {
	// defer func to flush buf, handler error
	defer func() {
		_ = c.buf.Flush()
		// when error close conn
		if err != nil {
			_ = c.Close()
		}
	}()

	if err := c.enc.Encode(h); err != nil {
		log.Println("rpc codec: gob encode header error: ", err)
		return err
	}

	if err := c.enc.Encode(body); err != nil {
		log.Println("rpc codec: gob encode body error: ", err)
		return err
	}

	return nil
}