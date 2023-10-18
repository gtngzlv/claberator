package claberator

import (
	"github.com/gtngzlv/claberator/internal"
)

func ReturnRandomClabe() internal.Clabe {
	accountNumber := internal.ReturnRandomAccountNumber()
	bankInfo := internal.ReturnRandomBank()
	cityInfo := internal.ReturnRandomCity()

	var clabe internal.Clabe
	generated := clabe.Calculate(bankInfo.Code, cityInfo.Code, accountNumber)
	clabe.Number = generated
	clabe.BankInfo = bankInfo
	clabe.CityInfo = cityInfo
	return clabe
}
