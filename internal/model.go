package internal

// Clabe представляет набор данных и функций для работы с CLABE-номерами.
type Clabe struct {
	BanksMap  map[int]ClabeBank
	CitiesMap map[int][]ClabeCityInfo
	Cities    []ClabeCityInfo
}

type ClabeBank struct {
	Tag  string
	Name string
}

type ClabeBanksMap map[int]ClabeBank

type ClabeCityInfo struct {
	Code  int
	Name  string
	State string
}

type ClabeCitiesMap map[int][]ClabeCityInfo

type ClabeCheck struct {
	OK       bool     // Проверка на валидность
	FormatOK bool     // Валидная длина и контрольная сумма
	Error    string   // Код ошибки, например: 'invalid-city'
	Message  string   // Информация для отображения статуса
	Clabe    string   // Полный 18-значный номер
	Tag      string   // Краткое название банка, например: 'BANAMEX'
	Bank     string   // Полное название банка, например: 'Banco Nacional'
	City     string   // Название отделения или плазы
	Multiple bool     // Несколько городов имеют один и тот же код
	Total    int      // Количество городов
	Account  string   // 11-значный номер банковского счета с нулями
	Code     struct { // 3-значные коды
		Bank string // Код банка
		City string // Код города
	}
	Checksum int // Контрольная цифра (от 0 до 9)
}
