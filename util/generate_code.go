package util

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateCode() string {
	rand.Seed(time.Now().UnixNano())

	codigo := rand.Intn(100000000)

	codigoFormatado := fmt.Sprintf("%08d", codigo)

	return codigoFormatado
}
