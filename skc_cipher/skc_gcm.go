package skc_cipher

import (
	"crypto/cipher"

	"github.com/isteffyx/crypto/v4/skc_session"
)

// AEAD is a cipher mode providing authenticated encryption with associated
// data. For a description of the methodology, see
//	https://en.wikipedia.org/wiki/Authenticated_encryption
type AEAD = cipher.AEAD

// In general, the GHASH operation performed by this implementation of GCM is not constant-time.
// An exception is when the underlying Block was created by aes.NewCipher
// on systems with hardware support for AES. See the crypto/aes package documentation for details.
func NewGCM(cphr cipher.Block) (AEAD, error) {
	return cipher.NewGCM(cphr)
}

//getting session id
func getSessionId() []byte {
	return skc_session.SessionDetails.SessionId
}
