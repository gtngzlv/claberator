package claberator

import (
	"github.com/gtngzlv/claberator/internal"
)

func ReturnRandomClabe() internal.Clabe {
	accountNumber := internal.ReturnRandomAccountNumber()
	bankInfo := internal.ReturnRandomBank()
	cityInfo, cityCode := internal.ReturnRandomCity()

	var clabe internal.Clabe
	generated := clabe.Calculate(bankInfo.Code, cityCode, accountNumber)
	clabe.Number = generated
	clabe.BankInfo = bankInfo
	clabe.CityInfo = cityInfo
	return clabe
}
