package internal

import "github.com/brianvoe/gofakeit/v6"

func ReturnRandomAccountNumber() int {
	return gofakeit.Number(00000000001, 99999999999)
}
