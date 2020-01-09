package csgologreceiver

import (
	"net"
	"testing"
)

func TestReceiver(t *testing.T) {

	t.Run("Init and close receiver", func(t *testing.T) {

		// when
		r, _ := New("0.0.0.0", 12345)

		// then
		r.Close()
	})

	t.Run("Init, send message and close receiver", func(t *testing.T) {

		// given
		r, _ := New("0.0.0.0", 12345)
		conn, _ := net.Dial("udp", "0.0.0.0:12345")
		conn.Write([]byte(`xxxxRL foo`))

		// when
		response, _ := r.Read()

		// then
		assert(t, "L foo", response.Message)

		// after
		r.Close()
	})
}

func TestParseHeader(t *testing.T) {

	t.Run("Parse no match", func(t *testing.T) {

		// when
		_, _, _, err := ParseHeader([]byte(`xxxxFOO`))

		// then
		assert(t, ErrorNoMatch, err)
	})

	t.Run("Parse header no secret nor token", func(t *testing.T) {

		// when
		_, _, m, _ := ParseHeader([]byte(`xxxxRL foo`))

		// then
		assert(t, "L foo", m)
	})

	t.Run("Parse header with secret but no token", func(t *testing.T) {

		// when
		s, _, m, _ := ParseHeader([]byte(`xxxxS2fooL foo`))

		// then
		assert(t, "L foo", m)
		assert(t, "2foo", s)
	})

	t.Run("Parse header no secret but with token", func(t *testing.T) {

		// when
		_, tok, m, _ := ParseHeader([]byte(`xxxxRTFOO L foo`))

		// then
		assert(t, "L foo", m)
		assert(t, "FOO", tok)
	})

	t.Run("Parse header with secret and token", func(t *testing.T) {

		// when
		s, tok, m, _ := ParseHeader([]byte(`xxxxS2fooTFOO L foo`))

		// then
		assert(t, "L foo", m)
		assert(t, "2foo", s)
		assert(t, "FOO", tok)
	})
}

// helper

func assert(t *testing.T, want interface{}, have interface{}) {

	// mark as test helper function
	t.Helper()

	if want != have {
		t.Error("Assertion failed for", t.Name(), "\n\twanted:\t", want, "\n\thave:\t", have)
	}
}
