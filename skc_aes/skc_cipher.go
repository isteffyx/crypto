// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skc_aes

import (
	"crypto/cipher"

	"crypto/aes"

	"github.com/isteffyx/crypto/v4/skc_session"
)

// NewCipher creates and returns a new cipher.Block.
// The key argument should be the AES key,
// either 16, 24, or 32 bytes to select
// AES-128, AES-192, or AES-256.
func NewCipher(key []byte) (cipher.Block, error) {
	return aes.NewCipher(key)
}

//getting session id
func GetSessionId() []byte {
	return skc_session.SessionDetails.SessionId
}
