package internal

import (
	"crypto/rand"
	"log"
	"math/big"
)

var banks = ClabeBanks{
	{Code: 1, Tag: "BANXICO", Name: "Banco de México"},
	{Code: 2, Tag: "BANameX", Name: "Banco Nacional de México"},
	{Code: 6, Tag: "BANCOMEXT", Name: "Banco Nacional de Comercio Exterior"},
	{Code: 9, Tag: "BANOBRAS", Name: "Banco Nacional de Obras y Servicios Públicos"},
	{Code: 12, Tag: "BBVA BANCOMER", Name: "BBVA Bancomer"},
	{Code: 14, Tag: "SANTANDER", Name: "Banco Santander"},
	{Code: 19, Tag: "BANJERCITO", Name: "Banco Nacional del Ejército, Fuerza Aérea y Armada"},
	{Code: 21, Tag: "HSBC", Name: "HSBC México"},
	{Code: 30, Tag: "BAJIO", Name: "Banco del Bajío"},
	{Code: 36, Tag: "INBURSA", Name: "Banco Inbursa"},
	{Code: 42, Tag: "MIFEL", Name: "Banca Mifel"},
	{Code: 44, Tag: "SCOTIABANK", Name: "Scotiabank Inverlat"},
	{Code: 58, Tag: "BANREGIO", Name: "Banco Regional de Monterrey"},
	{Code: 59, Tag: "INVEX", Name: "Banco Invex"},
	{Code: 60, Tag: "BANSI", Name: "Bansi"},
	{Code: 62, Tag: "AFIRME", Name: "Banca Afirme"},
	{Code: 72, Tag: "BANORTE", Name: "Banco Mercantil del Norte"},
	{Code: 102, Tag: "ACCENDO BANCO", Name: "ABN AMRO Bank México"},
	{Code: 103, Tag: "AMERICAN EXPRESS", Name: "American Express Bank (México)"},
	{Code: 106, Tag: "BAMSA", Name: "Bank of America México"},
	{Code: 108, Tag: "TOKYO", Name: "Bank of Tokyo-Mitsubishi UFJ (México)"},
	{Code: 110, Tag: "JP MORGAN", Name: "Banco J.P. Morgan"},
	{Code: 112, Tag: "BMONEX", Name: "Banco Monex"},
	{Code: 113, Tag: "VE POR MAS", Name: "Banco Ve por Mas"},
	{Code: 124, Tag: "DEUTSCHE", Name: "Deutsche Bank México"},
	{Code: 126, Tag: "CREDIT SUISSE", Name: "Banco Credit Suisse (México)"},
	{Code: 127, Tag: "AZTECA", Name: "Banco Azteca"},
	{Code: 128, Tag: "AUTOFIN", Name: "Banco Autofin México"},
	{Code: 129, Tag: "BARCLAYS", Name: "Barclays Bank México"},
	{Code: 130, Tag: "COMPARTAMOS", Name: "Banco Compartamos"},
	{Code: 132, Tag: "BMULTIVA", Name: "Banco Multiva"},
	{Code: 133, Tag: "ACTINVER", Name: "Banco Actinver"},
	{Code: 135, Tag: "NAFIN", Name: "Nacional Financiera"},
	{Code: 136, Tag: "INTERBANCO", Name: "Inter Banco"},
	{Code: 137, Tag: "BANCOPPEL", Name: "BanCoppel"},
	{Code: 138, Tag: "ABC CAPITAL", Name: "ABC Capital"},
	{Code: 140, Tag: "CONSUBANCO", Name: "Banco Fácil"},
	{Code: 141, Tag: "VOLKSWAGEN", Name: "Volkswagen Bank, Institución de Banca Múltiple"},
	{Code: 143, Tag: "CIBANCO", Name: "Consultoría Internacional Banco"},
	{Code: 145, Tag: "BBASE", Name: "Banco BASE"},
	{Code: 147, Tag: "BANKAOOL", Name: "Bankaool, Institución de Banca Múltiple"},
	{Code: 148, Tag: "PAGATODO", Name: "Banco PagaTodo Institución de Banca Múltiple"},
	{Code: 150, Tag: "INMOBILIARIO", Name: "Banco Inmobiliario Mexicano, Institución de Banca Múltiple"},
	{Code: 151, Tag: "DONDE", Name: "Fundación Dondé Banco"},
	{Code: 152, Tag: "BANCREA", Name: "Banco Bancrea, Institución de Banca Múltiple"},
	{Code: 154, Tag: "BANCO FINTERRA", Name: "Banco Finterra"},
	{Code: 155, Tag: "ICBC", Name: "Industrial and Commercial Bank of China"},
	{Code: 156, Tag: "SABADELL", Name: "Banco de Sabadell"},
	{Code: 157, Tag: "SHINHAN", Name: "Shinhan Bank"},
	{Code: 158, Tag: "MIZUHO BANK", Name: "Mizuho Bank"},
	{Code: 159, Tag: "BANK OF CHINA", Name: "Bank of China México"},
	{Code: 160, Tag: "BANCO S3", Name: "Banco S3 México"},
	{Code: 166, Tag: "BANSEFI", Name: "Banco del Ahorro Nacional y Servicios Financieros"},
	{Code: 168, Tag: "HIPOTECARIA FED", Name: "Sociedad Hipotecaria Federal"},
	{Code: 600, Tag: "MONEXCB", Name: "Monex Casa de Bolsa"},
	{Code: 601, Tag: "GBM", Name: "GBM Grupo Bursátil Mexicano"},
	{Code: 602, Tag: "MASARI", Name: "Masari Casa de Bolsa"},
	{Code: 605, Tag: "VALUE", Name: "Valué, Casa de Bolsa"},
	{Code: 606, Tag: "ESTRUCTURADORES", Name: "Base Internacional Casa de Bolsa"},
	{Code: 608, Tag: "VECTOR", Name: "Vector Casa de Bolsa"},
	{Code: 613, Tag: "MULTIVA", Name: "Multivalores Casa de Bolsa, Multiva Gpo. Fin."},
	{Code: 616, Tag: "FINameX", Name: "Casa de Bolsa FiNamex"},
	{Code: 617, Tag: "VALMEX", Name: "Valores Mexicanos Casa de Bolsa"},
	{Code: 620, Tag: "PROFUTURO", Name: "Profuturo G.N.P., Afore"},
	{Code: 630, Tag: "CB INTERCAM", Name: "Intercam Casa de Bolsa"},
	{Code: 631, Tag: "CI BOLSA", Name: "CI Casa de Bolsa"},
	{Code: 634, Tag: "FINCOMUN", Name: "Fincomún, Servicios Financieros Comunitarios"},
	{Code: 636, Tag: "HDI SEGUROS", Name: "HDI Seguros"},
	{Code: 638, Tag: "AKALA", Name: "Akala, Sociedad Financiera Popular"},
	{Code: 642, Tag: "REFORMA", Name: "Operadora de Recursos Reforma"},
	{Code: 646, Tag: "STP", Name: "Sistema de Transferencias y Pagos STP, SOFOM E.N.R."},
	{Code: 648, Tag: "EVERCORE", Name: "Evercore Casa de Bolsa"},
	{Code: 652, Tag: "CREDICAPITAL", Name: "Solución Asea, Sociedad Financiera Popular"},
	{Code: 653, Tag: "KUSPIT", Name: "Kuspit Casa de Bolsa"},
	{Code: 656, Tag: "UNAGRA", Name: "UNAGRA, S.F.P."},
	{Code: 659, Tag: "ASP INTEGRA OPC", Name: "Opciones Empresariales Del Noreste"},
	{Code: 661, Tag: "ALTERNATIVOS", Name: "Servicios Financieros Alternativos (Klar)"},
	{Code: 670, Tag: "LIBERTAD", Name: "Libertad Servicios Financieros"},
	{Code: 677, Tag: "CAJA POP MEXICA", Name: "Caja Popular Mexicana"},
	{Code: 680, Tag: "CRISTOBAL COLON", Name: "Caja Popular Cristóbal Colón"},
	{Code: 683, Tag: "CAJA TELEFONIST", Name: "Caja de Ahorro de los Telefonistas"},
	{Code: 684, Tag: "TRANSFER", Name: "Operadora de Pagos Moviles de Mexico"},
	{Code: 685, Tag: "FONDO (FIRA)", Name: "Fondo de Garantía y Fomento para la Agricultura, Ganadería y Avicultura (FONDO)"},
	{Code: 686, Tag: "INVERCAP", Name: "Afore InverCap"},
	{Code: 689, Tag: "FOMPED", Name: "Fondo Mexicano del Petroleo para la Estabilizacion y el Desarrollo"},
	{Code: 706, Tag: "ARCUS", Name: "Arcus Financial Intelligence"},
	{Code: 710, Tag: "NVIO", Name: "NVIO Pagos México"},
	{Code: 812, Tag: "BBVA BANCOMER2", Name: "BBVA Bancomer"},
	{Code: 814, Tag: "SANTANDER2", Name: "Banco Santander"},
	{Code: 821, Tag: "HSBC2", Name: "HSBC México"},
	{Code: 901, Tag: "CLS", Name: "CLS Bank International"},
	{Code: 902, Tag: "INDEVAL", Name: "SD. INDEVAL"},
	{Code: 903, Tag: "CODI VALIDA", Name: "CoDi Valida"},
	{Code: 999, Tag: "N/A", Name: "N/A"},
}

var notParsedBanks = ClabeBanks{
	{Code: 22, Tag: "GE MONEY", Name: "GE Money Bank"},
	{Code: 32, Tag: "IXE", Name: "IXE Banco"},
	{Code: 37, Tag: "INTERACCIONES", Name: "Banco Interacciones"},
	{Code: 116, Tag: "ING", Name: "ING Bank (México)"},
	{Code: 131, Tag: "FAMSA", Name: "Banco Ahorro Famsa"},
	{Code: 134, Tag: "WAL-MART", Name: "Banco Wal-Mart de México Adelante"},
	{Code: 139, Tag: "UBS BANK", Name: "UBS Banco"},
	{Code: 604, Tag: "C.B. INBURSA", Name: "Inversora Bursátil"},
	{Code: 607, Tag: "TIBER", Name: "Casa de Cambio Tiber"},
	{Code: 610, Tag: "B&B", Name: "B y B Casa de Cambio"},
	{Code: 611, Tag: "INTERCAM", Name: "Intercam Casa de Cambio"},
	{Code: 614, Tag: "ACCIVAL", Name: "Acciones y Valores BaNamex, Casa de Bolsa"},
	{Code: 615, Tag: "MERRILL LYNCH", Name: "Merrill Lynch México, Casa de Bolsa"},
	{Code: 618, Tag: "ÚNICA", Name: "Única Casa de Cambio"},
	{Code: 619, Tag: "ASEGURADORA MAPFRE", Name: "MAPFRE Tepeyac"},
	{Code: 621, Tag: "CB ACTINBER", Name: "Actinver Casa de Bolsa"},
	{Code: 622, Tag: "ACTINVE SI", Name: "Actinver"},
	{Code: 623, Tag: "SKANDIA", Name: "Skandia Vida"},
	{Code: 624, Tag: "CONSULTORÍA", Name: "Consultoría Internacional Casa de Cambio"},
	{Code: 626, Tag: "CBDEUTSCHE", Name: "Deutsche Securities"},
	{Code: 627, Tag: "ZURICH", Name: "Zurich Compañía de Seguros"},
	{Code: 628, Tag: "ZURICHVI", Name: "Zurich Vida, Compañía de Seguros"},
	{Code: 629, Tag: "HIPOTECARIA SU CASITA", Name: "Hipotecaria su Casita"},
	{Code: 632, Tag: "BULLTICK C.B.", Name: "Bulltick Casa de Bolsa"},
	{Code: 633, Tag: "STERLING", Name: "Sterling Casa de Cambio"},
	{Code: 637, Tag: "ORDER", Name: "OrderExpress Casa de Cambio"},
	{Code: 640, Tag: "CB JPMORGAN", Name: "J.P. Morgan Casa de Bolsa"},
	{Code: 647, Tag: "TELECOMM", Name: "Telecomunicaciones de México"},
	{Code: 649, Tag: "SKANDIA", Name: "Skandia Operadora"},
	{Code: 651, Tag: "SEGMTY", Name: "Seguros Monterrey New York Life"},
	{Code: 655, Tag: "SOFIEXPRESS", Name: "JP SofiExpress"},
	{Code: 674, Tag: "AXA", Name: "AXA Seguros"},
	{Code: 679, Tag: "FND", Name: "Financiera Nacional De Desarrollo Agropecuario"},
	{Code: 846, Tag: "STP", Name: "Sistema de Transferencias y Pagos STP"}}

func ReturnRandomBank() ClabeBank {
	mapLen := len(banks)
	randNumber, err := rand.Int(rand.Reader, big.NewInt(int64(mapLen-1)))
	if err != nil {
		log.Println(err)
	}
	intRandomNumber := int(randNumber.Int64())
	randomBank := banks[intRandomNumber]
	return randomBank
}

func ReturnNotParsedBank() ClabeBank {
	mapLen := len(notParsedBanks)
	randNumber, err := rand.Int(rand.Reader, big.NewInt(int64(mapLen-1)))
	if err != nil {
		log.Println(err)
	}
	intRandomNumber := int(randNumber.Int64())
	randomBank := notParsedBanks[intRandomNumber]
	return randomBank
}
