package main

import (
	"fmt"

	"github.com/gtngzlv/claberator/internal"
)

func main() {
	accountNumber := internal.ReturnRandomAccountNumber()
	_, bankCode := internal.ReturnRandomBank()
	_, cityCode := internal.ReturnRandomCity()

	var clabe internal.Clabe
	generated := clabe.Calculate(bankCode, cityCode, accountNumber)
	fmt.Println(generated)
}
