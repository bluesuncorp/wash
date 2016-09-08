package globals

import (
	"sync"

	"github.com/go-playground/log"
)

// ByteBuffer contains all buffer related logic
type ByteBuffer struct {
	pool *sync.Pool
}

// NewByteBuffer returns a new ByteBuffer instance
func NewByteBuffer() *ByteBuffer {

	log.Info("Initializing ByteBuffer")

	return &ByteBuffer{
		pool: &sync.Pool{New: func() interface{} {
			return make([]byte, 0, 64)
		}},
	}
}

// Get returns buffer from the pool or a new instance of one if none exists
func (b *ByteBuffer) Get() []byte {
	return b.pool.Get().([]byte)
}

// Put returns a buffer to the pool + resets it for next use
func (b *ByteBuffer) Put(bytes []byte) {
	b.pool.Put(bytes[0:0])
}
