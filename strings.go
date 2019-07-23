package zimbra

import (
	"crypto/rand"
	"math/big"
)

// GenString makes random strings of any length for passwords and such.
func GenString(size int) string {
	// Somewhat password-friendly.
	valid := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!#-.")
	pw := make([]byte, size)
	for i := 0; i < size; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(valid))))
		if err != nil {
			return ""
		}
		c := valid[n.Int64()]
		pw[i] = c
	}
	return string(pw)
}
