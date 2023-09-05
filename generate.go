package main

import "github.com/gtngzlv/claberator/data"

// залить на личный гитхаб, сделать тег и импортнуть
func ReturnRandomClabe() data.Clabe {
	accountNumber := data.ReturnRandomAccountNumber()
	bankInfo, bankCode := data.ReturnRandomBank()
	cityInfo, cityCode := data.ReturnRandomCity()

	var clabe data.Clabe
	generated := clabe.Calculate(bankCode, cityCode, accountNumber)
	clabe.Number = generated
	clabe.BankInfo = bankInfo
	clabe.CityInfo = cityInfo
	return clabe
}
