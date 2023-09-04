package internal

func ReturnBanks() map[int]ClabeBank {
	banksMap := map[int]ClabeBank{
		1:   {Tag: "BANXICO", Name: "Banco de México"},
		2:   {Tag: "BANameX", Name: "Banco Nacional de México"},
		6:   {Tag: "BANCOMEXT", Name: "Banco Nacional de Comercio Exterior"},
		9:   {Tag: "BANOBRAS", Name: "Banco Nacional de Obras y Servicios Públicos"},
		12:  {Tag: "BBVA BANCOMER", Name: "BBVA Bancomer"},
		14:  {Tag: "SANTANDER", Name: "Banco Santander"},
		19:  {Tag: "BANJERCITO", Name: "Banco Nacional del Ejército, Fuerza Aérea y Armada"},
		21:  {Tag: "HSBC", Name: "HSBC México"},
		22:  {Tag: "GE MONEY", Name: "GE Money Bank"},
		30:  {Tag: "BAJIO", Name: "Banco del Bajío"},
		32:  {Tag: "IXE", Name: "IXE Banco"},
		36:  {Tag: "INBURSA", Name: "Banco Inbursa"},
		37:  {Tag: "INTERACCIONES", Name: "Banco Interacciones"},
		42:  {Tag: "MIFEL", Name: "Banca Mifel"},
		44:  {Tag: "SCOTIABANK", Name: "Scotiabank Inverlat"},
		58:  {Tag: "BANREGIO", Name: "Banco Regional de Monterrey"},
		59:  {Tag: "INVEX", Name: "Banco Invex"},
		60:  {Tag: "BANSI", Name: "Bansi"},
		62:  {Tag: "AFIRME", Name: "Banca Afirme"},
		72:  {Tag: "BANORTE", Name: "Banco Mercantil del Norte"},
		102: {Tag: "ACCENDO BANCO", Name: "ABN AMRO Bank México"},
		103: {Tag: "AMERICAN EXPRESS", Name: "American Express Bank (México)"},
		106: {Tag: "BAMSA", Name: "Bank of America México"},
		108: {Tag: "TOKYO", Name: "Bank of Tokyo-Mitsubishi UFJ (México)"},
		110: {Tag: "JP MORGAN", Name: "Banco J.P. Morgan"},
		112: {Tag: "BMONEX", Name: "Banco Monex"},
		113: {Tag: "VE POR MAS", Name: "Banco Ve por Mas"},
		116: {Tag: "ING", Name: "ING Bank (México)"},
		124: {Tag: "DEUTSCHE", Name: "Deutsche Bank México"},
		126: {Tag: "CREDIT SUISSE", Name: "Banco Credit Suisse (México)"},
		127: {Tag: "AZTECA", Name: "Banco Azteca"},
		128: {Tag: "AUTOFIN", Name: "Banco Autofin México"},
		129: {Tag: "BARCLAYS", Name: "Barclays Bank México"},
		130: {Tag: "COMPARTAMOS", Name: "Banco Compartamos"},
		131: {Tag: "FAMSA", Name: "Banco Ahorro Famsa"},
		132: {Tag: "BMULTIVA", Name: "Banco Multiva"},
		133: {Tag: "ACTINVER", Name: "Banco Actinver"},
		134: {Tag: "WAL-MART", Name: "Banco Wal-Mart de México Adelante"},
		135: {Tag: "NAFIN", Name: "Nacional Financiera"},
		136: {Tag: "INTERBANCO", Name: "Inter Banco"},
		137: {Tag: "BANCOPPEL", Name: "BanCoppel"},
		138: {Tag: "ABC CAPITAL", Name: "ABC Capital"},
		139: {Tag: "UBS BANK", Name: "UBS Banco"},
		140: {Tag: "CONSUBANCO", Name: "Banco Fácil"},
		141: {Tag: "VOLKSWAGEN", Name: "Volkswagen Bank, Institución de Banca Múltiple"},
		143: {Tag: "CIBANCO", Name: "Consultoría Internacional Banco"},
		145: {Tag: "BBASE", Name: "Banco BASE"},
		147: {Tag: "BANKAOOL", Name: "Bankaool, Institución de Banca Múltiple"},
		148: {Tag: "PAGATODO", Name: "Banco PagaTodo Institución de Banca Múltiple"},
		150: {Tag: "INMOBILIARIO", Name: "Banco Inmobiliario Mexicano, Institución de Banca Múltiple"},
		151: {Tag: "DONDE", Name: "Fundación Dondé Banco"},
		152: {Tag: "BANCREA", Name: "Banco Bancrea, Institución de Banca Múltiple"},
		154: {Tag: "BANCO FINTERRA", Name: "Banco Finterra"},
		155: {Tag: "ICBC", Name: "Industrial and Commercial Bank of China"},
		156: {Tag: "SABADELL", Name: "Banco de Sabadell"},
		157: {Tag: "SHINHAN", Name: "Shinhan Bank"},
		158: {Tag: "MIZUHO BANK", Name: "Mizuho Bank"},
		159: {Tag: "BANK OF CHINA", Name: "Bank of China México"},
		160: {Tag: "BANCO S3", Name: "Banco S3 México"},
		166: {Tag: "BANSEFI", Name: "Banco del Ahorro Nacional y Servicios Financieros"},
		168: {Tag: "HIPOTECARIA FED", Name: "Sociedad Hipotecaria Federal"},
		600: {Tag: "MONEXCB", Name: "Monex Casa de Bolsa"},
		601: {Tag: "GBM", Name: "GBM Grupo Bursátil Mexicano"},
		602: {Tag: "MASARI", Name: "Masari Casa de Bolsa"},
		604: {Tag: "C.B. INBURSA", Name: "Inversora Bursátil"},
		605: {Tag: "VALUE", Name: "Valué, Casa de Bolsa"},
		606: {Tag: "ESTRUCTURADORES", Name: "Base Internacional Casa de Bolsa"},
		607: {Tag: "TIBER", Name: "Casa de Cambio Tiber"},
		608: {Tag: "VECTOR", Name: "Vector Casa de Bolsa"},
		610: {Tag: "B&B", Name: "B y B Casa de Cambio"},
		611: {Tag: "INTERCAM", Name: "Intercam Casa de Cambio"},
		613: {Tag: "MULTIVA", Name: "Multivalores Casa de Bolsa, Multiva Gpo. Fin."},
		614: {Tag: "ACCIVAL", Name: "Acciones y Valores BaNamex, Casa de Bolsa"},
		615: {Tag: "MERRILL LYNCH", Name: "Merrill Lynch México, Casa de Bolsa"},
		616: {Tag: "FINameX", Name: "Casa de Bolsa FiNamex"},
		617: {Tag: "VALMEX", Name: "Valores Mexicanos Casa de Bolsa"},
		618: {Tag: "ÚNICA", Name: "Única Casa de Cambio"},
		619: {Tag: "ASEGURADORA MAPFRE", Name: "MAPFRE Tepeyac"},
		620: {Tag: "PROFUTURO", Name: "Profuturo G.N.P., Afore"},
		621: {Tag: "CB ACTINBER", Name: "Actinver Casa de Bolsa"},
		622: {Tag: "ACTINVE SI", Name: "Actinver"},
		623: {Tag: "SKANDIA", Name: "Skandia Vida"},
		624: {Tag: "CONSULTORÍA", Name: "Consultoría Internacional Casa de Cambio"},
		626: {Tag: "CBDEUTSCHE", Name: "Deutsche Securities"},
		627: {Tag: "ZURICH", Name: "Zurich Compañía de Seguros"},
		628: {Tag: "ZURICHVI", Name: "Zurich Vida, Compañía de Seguros"},
		629: {Tag: "HIPOTECARIA SU CASITA", Name: "Hipotecaria su Casita"},
		630: {Tag: "CB INTERCAM", Name: "Intercam Casa de Bolsa"},
		631: {Tag: "CI BOLSA", Name: "CI Casa de Bolsa"},
		632: {Tag: "BULLTICK C.B.", Name: "Bulltick Casa de Bolsa"},
		633: {Tag: "STERLING", Name: "Sterling Casa de Cambio"},
		634: {Tag: "FINCOMUN", Name: "Fincomún, Servicios Financieros Comunitarios"},
		636: {Tag: "HDI SEGUROS", Name: "HDI Seguros"},
		637: {Tag: "ORDER", Name: "OrderExpress Casa de Cambio"},
		638: {Tag: "AKALA", Name: "Akala, Sociedad Financiera Popular"},
		640: {Tag: "CB JPMORGAN", Name: "J.P. Morgan Casa de Bolsa"},
		642: {Tag: "REFORMA", Name: "Operadora de Recursos Reforma"},
		646: {Tag: "STP", Name: "Sistema de Transferencias y Pagos STP, SOFOM E.N.R."},
		647: {Tag: "TELECOMM", Name: "Telecomunicaciones de México"},
		648: {Tag: "EVERCORE", Name: "Evercore Casa de Bolsa"},
		649: {Tag: "SKANDIA", Name: "Skandia Operadora"},
		651: {Tag: "SEGMTY", Name: "Seguros Monterrey New York Life"},
		652: {Tag: "CREDICAPITAL", Name: "Solución Asea, Sociedad Financiera Popular"},
		653: {Tag: "KUSPIT", Name: "Kuspit Casa de Bolsa"},
		655: {Tag: "SOFIEXPRESS", Name: "JP SofiExpress"},
		656: {Tag: "UNAGRA", Name: "UNAGRA, S.F.P."},
		659: {Tag: "ASP INTEGRA OPC", Name: "Opciones Empresariales Del Noreste"},
		661: {Tag: "ALTERNATIVOS", Name: "Servicios Financieros Alternativos (Klar)"},
		670: {Tag: "LIBERTAD", Name: "Libertad Servicios Financieros"},
		674: {Tag: "AXA", Name: "AXA Seguros"},
		677: {Tag: "CAJA POP MEXICA", Name: "Caja Popular Mexicana"},
		679: {Tag: "FND", Name: "Financiera Nacional De Desarrollo Agropecuario"},
		680: {Tag: "CRISTOBAL COLON", Name: "Caja Popular Cristóbal Colón"},
		683: {Tag: "CAJA TELEFONIST", Name: "Caja de Ahorro de los Telefonistas"},
		684: {Tag: "TRANSFER", Name: "Operadora de Pagos Moviles de Mexico"},
		685: {Tag: "FONDO (FIRA)", Name: "Fondo de Garantía y Fomento para la Agricultura, Ganadería y Avicultura (FONDO)"},
		686: {Tag: "INVERCAP", Name: "Afore InverCap"},
		689: {Tag: "FOMPED", Name: "Fondo Mexicano del Petroleo para la Estabilizacion y el Desarrollo"},
		706: {Tag: "ARCUS", Name: "Arcus Financial Intelligence"},
		710: {Tag: "NVIO", Name: "NVIO Pagos México"},
		812: {Tag: "BBVA BANCOMER2", Name: "BBVA Bancomer"},
		814: {Tag: "SANTANDER2", Name: "Banco Santander"},
		821: {Tag: "HSBC2", Name: "HSBC México"},
		846: {Tag: "STP", Name: "Sistema de Transferencias y Pagos STP"},
		901: {Tag: "CLS", Name: "CLS Bank International"},
		902: {Tag: "INDEVAL", Name: "SD. INDEVAL"},
		903: {Tag: "CODI VALIDA", Name: "CoDi Valida"},
		999: {Tag: "N/A", Name: "N/A"},
	}
	return banksMap
}
