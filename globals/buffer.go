package globals

import (
	"sync"

	"github.com/go-playground/log"
)

// ByteBuffer contains all buffer related logic
type ByteBuffer interface {
	Get() []byte
	Put(bytes []byte)
}

// byteBuffer contains all buffer related logic
type byteBuffer struct {
	pool *sync.Pool
}

var _ ByteBuffer = new(byteBuffer)

// newByteBuffer returns a new ByteBuffer instance
func newByteBuffer() ByteBuffer {

	log.Info("Initializing ByteBuffer")

	return &byteBuffer{
		pool: &sync.Pool{New: func() interface{} {
			return make([]byte, 0, 64)
		}},
	}
}

// Get returns buffer from the pool or a new instance of one if none exists
func (b *byteBuffer) Get() []byte {
	return b.pool.Get().([]byte)
}

// Put returns a buffer to the pool + resets it for next use
func (b *byteBuffer) Put(bytes []byte) {
	b.pool.Put(bytes[0:0])
}
