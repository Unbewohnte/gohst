/*
            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
                    Version 2, December 2004

 Copyright (C) 2004 Sam Hocevar <sam@hocevar.net>

 Everyone is permitted to copy and distribute verbatim or modified
 copies of this license document, and changing it is allowed as long
 as the name is changed.

            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

  0. You just DO WHAT THE FUCK YOU WANT TO.
*/

// Kasyanov N.A. (Unbewohnte), 2023

package encryption

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// Encodes given string via Base64
func EncodeStringBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Decodes given string via Base64
func DecodeStringBase64(encodedStr string) string {
	decodedBytes, _ := base64.StdEncoding.DecodeString(encodedStr)
	return string(decodedBytes)
}

// Returns HEX string of SHA256'd data
func SHA256Hex(data []byte) string {
	hash := sha256.New()
	hash.Write(data)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
