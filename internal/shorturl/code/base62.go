package code

import "math/rand"

type Base62 struct {
	Length int
}

func (g Base62) Generate() string {
	const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, g.Length)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(b)
}
