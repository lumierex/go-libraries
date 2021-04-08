package sfpxm_web_core

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestWithPort(t *testing.T) {
	var serverOpts ServerOptions
	fServerOpts := WithPort("3000")
	fServerOpts(&serverOpts)
	assert.Equal(t, "3000", serverOpts.port)
	fServerOpts = WithPort("")
	fServerOpts(&serverOpts)
	assert.Equal(t, "", serverOpts.port)
}
