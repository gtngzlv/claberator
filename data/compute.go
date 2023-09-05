package data

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ComputeChecksum вычисляет контрольную цифру для CLABE-номера.
func (c *Clabe) ComputeChecksum(clabeNum17 string) int {
	x := func(i int) int { return []int{3, 7, 1}[i%3] }
	add := func(sum int, digit rune, i int) int {
		return sum + (int(digit-'0')*x(i))%10
	}
	compute := func() int {
		sum := 0
		for i, digit := range clabeNum17[:17] {
			sum = add(sum, digit, i)
		}
		return (10 - (sum % 10)) % 10
	}
	if match, _ := regexp.MatchString("^[0-9]{17,18}$", clabeNum17); match {
		return compute()
	}
	return -1
}

// Validate проверяет CLABE-номер и возвращает информацию о нем.
func (c *Clabe) Validate(clabeNum string) ClabeCheck {
	errorMap := map[string]string{
		"length":     "Должно быть ровно 18 цифр",
		"characters": "Должны быть только цифры (без букв)",
		"checksum":   "Неверная контрольная цифра, должна быть: ",
		"bank":       "Неверный код банка: ",
		"city":       "Неверный код города: ",
	}
	if len(clabeNum) != 18 {
		return ClabeCheck{
			Error:   "length",
			Message: errorMap["length"],
		}
	}
	if match, _ := regexp.MatchString("^[0-9]+$", clabeNum); !match {
		return ClabeCheck{
			Error:   "characters",
			Message: errorMap["characters"],
		}
	}
	bankCode := clabeNum[:3]
	cityCode := clabeNum[3:6]
	account := clabeNum[6:17]
	checksum := int(clabeNum[17] - '0')
	bankCodeInt, _ := strconv.Atoi(bankCode)
	bank, bankExists := c.BanksMap[bankCodeInt]
	cityCodeInt, _ := strconv.Atoi(cityCode)
	cities, citiesExist := c.CitiesMap[cityCodeInt]
	realChecksum := c.ComputeChecksum(clabeNum)

	var validationInfo struct {
		Invalid string
		Data    interface{}
	}
	switch {
	case !bankExists:
		validationInfo.Invalid = "bank"
		validationInfo.Data = bankCode
	case !citiesExist:
		validationInfo.Invalid = "city"
		validationInfo.Data = cityCode
	case checksum != realChecksum:
		validationInfo.Invalid = "checksum"
		validationInfo.Data = realChecksum
	}

	if validationInfo.Invalid != "" {
		return ClabeCheck{
			Error:   "invalid-" + validationInfo.Invalid,
			Message: errorMap[validationInfo.Invalid] + fmt.Sprint(validationInfo.Data),
		}
	}

	var cityNames []string
	for _, city := range cities {
		if city.State != "" {
			cityNames = append(cityNames, fmt.Sprintf("%s %s", city.Name, city.State))
		} else {
			cityNames = append(cityNames, city.Name)
		}
	}

	return ClabeCheck{
		OK:       true,
		FormatOK: true,
		Message:  "Действительный",
		Clabe:    clabeNum,
		Tag:      bank.Tag,
		Bank:     bank.Name,
		City:     strings.Join(cityNames, ", "),
		Multiple: len(cities) > 1,
		Total:    len(cities),
		Account:  account,
		Code: struct {
			Bank string
			City string
		}{bankCode, cityCode},
		Checksum: realChecksum,
	}
}

// Calculate вычисляет CLABE-номер.
func (c *Clabe) Calculate(bankCode, cityCode, accountNumber int) string {
	pad := func(text string, length int) string {
		for len(text) < length {
			text = "0" + text
		}
		return text
	}
	fit := func(num, length int) string {
		return pad(fmt.Sprint(num), length)
	}
	clabeNum := fit(bankCode, 3) + fit(cityCode, 3) + fit(accountNumber, 11)
	checksum := c.ComputeChecksum(clabeNum)
	clabeNum += fmt.Sprintf("%d", checksum)
	return clabeNum
}
