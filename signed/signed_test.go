package signed

import (
	"bytes"
	"math/rand"
	"testing"
)

func TestXOR(t *testing.T) {
	sign := Sha256([]byte("ololo"))

	payload := make([]byte, 500)
	rand.Read(payload)
	signed := sign.SignXOR(payload, 10)
	unsigned, err := sign.UnsignXOR(signed, 10)

	t.Logf("%s", signed)

	if err != nil {
		t.Error(err)
		return
	}

	if !bytes.Equal(payload, unsigned) {
		t.Errorf("Data is not equal (%s != %s)", payload, unsigned)
	}
}
