package internal

import (
	"crypto/rand"
	"log"
	"math/big"
)

func ReturnRandomAccountNumber() int {
	maxNumberValue := 9999
	randNumber, err := rand.Int(rand.Reader, big.NewInt(int64(maxNumberValue)))
	if err != nil {
		log.Println(err)
	}
	return int(randNumber.Int64())
}
