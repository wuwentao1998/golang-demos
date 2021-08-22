package tool

import (
	"bytes"
	"sync"
)

var bufferPool *sync.Pool

func init() {
	bufferPool = &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
}

func GetBytesBuf() *bytes.Buffer {
	if buf, ok := bufferPool.Get().(*bytes.Buffer); ok {
		return buf
	}

	return new(bytes.Buffer)
}

func PutBytesBuf(buf *bytes.Buffer) {
	buf.Reset()
	bufferPool.Put(buf)
}
