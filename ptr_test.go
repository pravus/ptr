package ptr_test

import (
	"testing"

	"github.com/pravus/ptr"
	"github.com/stretchr/testify/require"
)

func testTo[T any](t *testing.T, v T) {
	p := ptr.To(v)
	require.Equal(t, v, *p)
	require.IsType(t, v, *p)
}

func TestTo(t *testing.T) {
	testTo(t, int(10))
	testTo(t, int8(21))
	testTo(t, int16(32))
	testTo(t, int32(43))
	testTo(t, int64(54))

	testTo(t, float32(1.32))
	testTo(t, float64(2.34))

	testTo(t, complex64(1+2i))
	testTo(t, complex128(2+3i))

	testTo(t, byte('g'))
	testTo(t, rune('o'))

	testTo(t, `go`)

	testTo(t, true)
	testTo(t, false)

	{
		v := struct {
			i int
			s string
			b bool
		}{42, `go`, true}
		p := ptr.To(v)
		require.Equal(t, v, *p)
		require.IsType(t, v, *p)
		require.Equal(t, v.i, p.i)
		require.IsType(t, v.i, p.i)
		require.Equal(t, v.s, p.s)
		require.IsType(t, v.s, p.s)
		require.Equal(t, v.b, p.b)
		require.IsType(t, v.b, p.b)

		p.i = 24
		p.s = `og`
		p.b = false
		require.NotEqual(t, v, *p)
		require.NotEqual(t, v.i, p.i)
		require.NotEqual(t, v.s, p.s)
		require.NotEqual(t, v.b, p.b)
	}
}
