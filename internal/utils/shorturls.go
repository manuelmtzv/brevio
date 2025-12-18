package utils

import (
	"fmt"
	"math/rand"
)

func GenerateUniqueCode() string {
	code := rand.Intn(1000000)
	return fmt.Sprintf("%06d", code)
}
