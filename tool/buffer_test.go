package tool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyBuffer(t *testing.T) {
	buf := GetBytesBuf()

	assert.Equal(t, buf.Len(), 0)
	assert.Equal(t, buf.Cap(), 0)

	const size = 1000
	buf.Grow(size)
	assert.Equal(t, buf.Cap(), size)
}

func TestRecoveredBuffer(t *testing.T) {
	t.Run("need grow", func(t *testing.T) {
		const cap1 = 500
		buf := GetBytesBuf()
		buf.Grow(cap1)
		PutBytesBuf(buf)

		buf = GetBytesBuf()
		assert.Equal(t, 0, buf.Len())
		assert.Equal(t, cap1, buf.Cap())

		const cap2 = 1000
		buf.Grow(cap2)
		assert.Equal(t, 0, buf.Len())
		assert.LessOrEqual(t, cap2, buf.Cap())
	})

	t.Run("needn't grow", func(t *testing.T) {
		const cap1 = 2000
		buf := GetBytesBuf()
		buf.Grow(cap1)
		PutBytesBuf(buf)

		buf = GetBytesBuf()
		assert.Equal(t, 0, buf.Len())
		assert.Equal(t, cap1, buf.Cap())

		const cap2 = 1000
		buf.Grow(cap2)
		assert.Equal(t, 0, buf.Len())
		assert.Equal(t, cap1, buf.Cap())
	})
}

func TestUsedBuffer(t *testing.T) {
	buf := GetBytesBuf()
	buf.WriteString("12345")

	assert.Equal(t, buf.Len(), 5)

	const size = 100
	buf.Grow(size)

	assert.LessOrEqual(t, size, buf.Cap()-buf.Len())
}
