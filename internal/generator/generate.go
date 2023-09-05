package generator

import "github.com/gtngzlv/claberator/internal"

// залить на личный гитхаб, сделать тег и импортнуть
func ReturnRandomClabe() internal.Clabe {
	accountNumber := internal.ReturnRandomAccountNumber()
	bankInfo, bankCode := internal.ReturnRandomBank()
	cityInfo, cityCode := internal.ReturnRandomCity()

	var clabe internal.Clabe
	generated := clabe.Calculate(bankCode, cityCode, accountNumber)
	clabe.Number = generated
	clabe.BankInfo = bankInfo
	clabe.CityInfo = cityInfo
	return clabe
}
