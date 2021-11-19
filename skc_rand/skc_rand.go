// Package rand implements a cryptographically secure
// random number generator.
package skc_rand

import (
	"crypto/rand"

	"github.com/isteffyx/crypto/v4/skc_session"
)

// Reader is a global, shared instance of a cryptographically
// secure random number generator.
//
// On Linux, FreeBSD, Dragonfly and Solaris, Reader uses getrandom(2) if
// available, /dev/urandom otherwise.
// On OpenBSD and macOS, Reader uses getentropy(2).
// On other Unix-like systems, Reader reads from /dev/urandom.
// On Windows systems, Reader uses the RtlGenRandom API.
// On Wasm, Reader uses the Web Crypto API.
var Reader = rand.Reader

//getting session id
func getSessionId() []byte {
	return skc_session.SessionDetails.SessionId
}
