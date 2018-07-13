package signed

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"hash"
	"math/rand"
)

var (
	ErrInvalidSignature = fmt.Errorf("Invalid signature")
)

type SignedPayload struct {
	h    func() hash.Hash
	key  []byte
	size int
}

func Sha256(key []byte) *SignedPayload {
	return New(sha256.New, key)
}

func New(h func() hash.Hash, key []byte) *SignedPayload {
	return &SignedPayload{
		h:    h,
		key:  key,
		size: h().Size(),
	}
}

func (sp *SignedPayload) SignatureSize() int {
	return sp.size
}

func (sp *SignedPayload) sum(p []byte) []byte {
	mac := hmac.New(sp.h, sp.key)
	mac.Write(p)
	return mac.Sum(nil)
}

func (sp *SignedPayload) Sign(p []byte) []byte {
	return append(sp.sum(p), p...)
}

func (sp *SignedPayload) SignXOR(p []byte, overhead int) []byte {
	if overhead <= 0 {
		return sp.Sign(p)
	}

	pLen := len(p)
	buf := make([]byte, overhead+sp.size+pLen)

	if _, err := rand.Read(buf[0:overhead]); err != nil {
		panic(err)
	}

	sum := sp.sum(p)
	offset := 0

	for i := 0; i < sp.size; i++ {
		buf[overhead+offset] = sum[i] ^ buf[offset%overhead]
		offset++
	}

	for i := 0; i < pLen; i++ {
		buf[overhead+offset] = p[i] ^ buf[offset%overhead]
		offset++
	}

	return buf
}

func (sp *SignedPayload) Unsign(p []byte) (r []byte, err error) {
	if len(p) < sp.size {
		err = ErrInvalidSignature
		return
	}

	r = p[sp.size:]
	sum := sp.sum(r)
	if !hmac.Equal(sum, p[:sp.size]) {
		err = ErrInvalidSignature
	}

	return
}

func (sp *SignedPayload) UnsignXOR(p []byte, overhead int) (r []byte, err error) {
	if overhead <= 0 {
		return sp.Unsign(p)
	}

	pLen := len(p)

	if pLen < sp.size+overhead {
		return nil, ErrInvalidSignature
	}

	data := make([]byte, pLen-overhead)
	offset := 0
	for i := overhead; i < pLen; i++ {
		data[offset] = p[i] ^ p[offset%overhead]
		offset++
	}

	r = data[sp.size:]
	sum := sp.sum(r)
	if !hmac.Equal(sum, data[:sp.size]) {
		err = ErrInvalidSignature
	}

	return
}
