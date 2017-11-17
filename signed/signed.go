package signed

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"hash"
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
	return append(sp.Sign(p), p...)
}

func (sk *SignedPayload) Unsign(p []byte) (r []byte, err error) {
	if len(p) < sk.size {
		err = ErrInvalidSignature
		return
	}

	r = p[sk.size:]
	sum := sk.sum(r)
	if !hmac.Equal(sum, p[:sk.size]) {
		err = ErrInvalidSignature
	}

	return

}
