package pool

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type buffered struct {
	payload    string
	resetCalls int
}

func (b *buffered) Reset() {
	b.payload = ""
	b.resetCalls++
}

func TestNew_ReturnsPool(t *testing.T) {
	p := New(func() *buffered { return &buffered{} })
	require.NotNil(t, p)
}

func TestGet_EmptyPoolCreatesFromFactory(t *testing.T) {
	p := New(func() *buffered { return &buffered{payload: "fresh"} })

	obj := p.Get()

	require.NotNil(t, obj)
	assert.Equal(t, "fresh", obj.payload)
}

func TestGetAfterPut_ReusesObject(t *testing.T) {
	p := New(func() *buffered { return &buffered{} })

	first := p.Get()
	first.payload = "dirty"
	p.Put(first)

	second := p.Get()

	assert.Same(t, first, second, "Put object should be reused from the pool")
	assert.Empty(t, second.payload, "reused object must be reset")
}

func TestPut_ResetsObject(t *testing.T) {
	p := New(func() *buffered { return &buffered{} })

	obj := p.Get()
	obj.payload = "dirty"
	p.Put(obj)

	assert.Equal(t, 1, obj.resetCalls, "Put must call Reset exactly once")
	assert.Empty(t, obj.payload)
}
