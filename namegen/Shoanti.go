/*
########################################################################################
#  __                                                                                  #
# /__ _                                                                                #
# \_|(_)                                                                               #
#  _______  _______  _______             _______     ______      _______               #
# (  ____ \(       )(  ___  ) Game      (  ____ \   / ___  \    (  __   )              #
# | (    \/| () () || (   ) | Master's  | (    \/   \/   \  \   | (  )  |              #
# | |      | || || || (___) | Assistant | (____        ___) /   | | /   |              #
# | | ____ | |(_)| ||  ___  | (Go Port) (_____ \      (___ (    | (/ /) |              #
# | | \_  )| |   | || (   ) |                 ) )         ) \   |   / | |              #
# | (___) || )   ( || )   ( | Mapper    /\____) ) _ /\___/  / _ |  (__) |              #
# (_______)|/     \||/     \| Client    \______/ (_)\______/ (_)(_______)              #
#                                                                                      #
########################################################################################
*/
//
// Name Generator Cultural Module
// Automatically generated from JavaScript source file
//

package namegen

//
// Shoanti describes the naming conventions for the Shoanti
// culture. Its methods give further details, but generally speaking
// the main operation to perform on these types is to just call the
// Generate and GenerateWithSurnames methods to create new names which
// conform to their cultural patterns.
//
type Shoanti struct {
	BaseCulture
}

//
// defaultMinMax returns the minimum and maximum size of Shoanti names based on gender.
//
func (c Shoanti) defaultMinMax(gender rune) (int, int) {
	switch gender {
	case 'F':
		return 4, 10
	case 'M':
		return 4, 10
	default:
		return 1, 1
	}
}

//
// Genders returns the set of genders defined for the Shoanti culture.
//
func (c Shoanti) Genders() []rune {
	return []rune{'F', 'M'}
}

//
// Name returns the name of the culture, i.e., "Shoanti".
//
func (c Shoanti) Name() string {
	return "Shoanti"
}

//
// HasGender returns true if the specified gender code is defined
// in the Shoanti culture.
//
func (c Shoanti) HasGender(gender rune) bool {
	switch gender {
	case 'F', 'M':
		return true
	default:
		return false
	}
}

//
// db returns the name data for the given gender in the Shoanti culture.
//
func (c Shoanti) db(gender rune) map[string][]nameFragment {
	switch gender {
	case 'F':
		return map[string][]nameFragment{
			"__K": {
				{'i', 0.25},
				{'a', 1.0},
			},
			"__N": {
				{'i', 0.25},
				{'o', 0.5},
				{'a', 1.0},
			},
			"Zon": {
				{'t', 1.0},
			},
			"ena": {
				{'s', 1.0},
			},
			"__Y": {
				{'o', 0.3333},
				{'a', 1.0},
			},
			"ina": {
				{0, 1.0},
			},
			"hna": {
				{'h', 1.0},
			},
			"ayw": {
				{'e', 1.0},
			},
			"_Si": {
				{'t', 1.0},
			},
			"yda": {
				{0, 1.0},
			},
			"ona": {
				{0, 0.75},
				{'h', 1.0},
			},
			"_Ot": {
				{'e', 1.0},
			},
			"_No": {
				{'y', 1.0},
			},
			"hid": {
				{'e', 1.0},
			},
			"_Ay": {
				{'i', 1.0},
			},
			"__C": {
				{'h', 1.0},
			},
			"lal": {
				{'i', 1.0},
			},
			"oly": {
				{0, 1.0},
			},
			"and": {
				{'a', 1.0},
			},
			"kch": {
				{'a', 1.0},
			},
			"ide": {
				{'e', 1.0},
			},
			"lah": {
				{'i', 1.0},
			},
			"npa": {
				{'y', 1.0},
			},
			"ska": {
				{'w', 1.0},
			},
			"Eyo": {
				{'t', 1.0},
			},
			"mam": {
				{'a', 1.0},
			},
			"__I": {
				{'n', 1.0},
			},
			"nol": {
				{'a', 1.0},
			},
			"Ool": {
				{'j', 1.0},
			},
			"zba": {
				{0, 1.0},
			},
			"_Sa": {
				{'s', 0.25},
				{'l', 0.5},
				{'h', 1.0},
			},
			"_Aw": {
				{'i', 0.5},
				{'e', 1.0},
			},
			"lok": {
				{'e', 1.0},
			},
			"jee": {
				{0, 1.0},
			},
			"Hia": {
				{'w', 1.0},
			},
			"sht": {
				{'a', 1.0},
			},
			"san": {
				{'w', 0.5},
				{'i', 1.0},
			},
			"_Do": {
				{'b', 0.25},
				{'l', 0.5},
				{'y', 0.75},
				{'w', 1.0},
			},
			"yol": {
				{0, 1.0},
			},
			"__P": {
				{'t', 1.0},
			},
			"Wak": {
				{'a', 1.0},
			},
			"Ayi": {
				{'t', 1.0},
			},
			"awi": {
				{0, 1.0},
			},
			"__B": {
				{'o', 1.0},
			},
			"Chu": {
				{'m', 1.0},
			},
			"uta": {
				{'h', 1.0},
			},
			"ich": {
				{'a', 1.0},
			},
			"Man": {
				{'a', 1.0},
			},
			"yto": {
				{'o', 1.0},
			},
			"rac": {
				{'u', 1.0},
			},
			"nwe": {
				{'e', 1.0},
			},
			"_At": {
				{'s', 1.0},
			},
			"mas": {
				{'a', 1.0},
			},
			"ela": {
				{0, 1.0},
			},
			"kyo": {
				{0, 1.0},
			},
			"_Mo": {
				{'s', 1.0},
			},
			"ola": {
				{0, 1.0},
			},
			"wan": {
				{'h', 1.0},
			},
			"sul": {
				{'a', 1.0},
			},
			"ays": {
				{'a', 1.0},
			},
			"_Sn": {
				{'a', 1.0},
			},
			"__T": {
				{'s', 0.3333},
				{'o', 0.5},
				{'a', 1.0},
			},
			"was": {
				{'s', 1.0},
			},
			"sga": {
				{0, 1.0},
			},
			"pay": {
				{'t', 1.0},
			},
			"neg": {
				{'a', 1.0},
			},
			"ntk": {
				{'a', 1.0},
			},
			"Sha": {
				{'d', 1.0},
			},
			"hpi": {
				{0, 1.0},
			},
			"zhi": {
				{0, 1.0},
			},
			"Wih": {
				{'a', 1.0},
			},
			"ahy": {
				{0, 1.0},
			},
			"uma": {
				{'n', 1.0},
			},
			"olj": {
				{'e', 1.0},
			},
			"_Oo": {
				{'l', 1.0},
			},
			"_De": {
				{'z', 1.0},
			},
			"_Na": {
				{'s', 0.5},
				{'h', 1.0},
			},
			"_Ad": {
				{'o', 0.5},
				{'s', 1.0},
			},
			"Ats": {
				{'i', 1.0},
			},
			"Eha": {
				{'w', 1.0},
			},
			"Gol": {
				{'a', 1.0},
			},
			"ash": {
				{'n', 0.5},
				{'t', 1.0},
			},
			"Woy": {
				{'a', 1.0},
			},
			"dza": {
				{0, 1.0},
			},
			"nas": {
				{'a', 1.0},
			},
			"ool": {
				{'y', 1.0},
			},
			"__G": {
				{'o', 0.6667},
				{'a', 1.0},
			},
			"yot": {
				{'a', 1.0},
			},
			"oli": {
				{0, 1.0},
			},
			"_As": {
				{'d', 1.0},
			},
			"iwi": {
				{0, 1.0},
			},
			"mah": {
				{0, 1.0},
			},
			"azh": {
				{'i', 1.0},
			},
			"Sna": {
				{'n', 1.0},
			},
			"Aha": {
				{'w', 1.0},
			},
			"caw": {
				{'i', 1.0},
			},
			"Une": {
				{'g', 1.0},
			},
			"sch": {
				{'a', 1.0},
			},
			"ywe": {
				{'e', 1.0},
			},
			"_Go": {
				{'g', 0.5},
				{'l', 1.0},
			},
			"had": {
				{'i', 1.0},
			},
			"sse": {
				{'e', 1.0},
			},
			"ahi": {
				{'m', 0.5},
				{0, 1.0},
			},
			"__E": {
				{'h', 0.5},
				{'y', 1.0},
			},
			"cun": {
				{'y', 1.0},
			},
			"imi": {
				{'m', 1.0},
			},
			"Ads": {
				{'i', 1.0},
			},
			"osi": {
				{0, 1.0},
			},
			"Niy": {
				{'o', 1.0},
			},
			"Mak": {
				{'a', 1.0},
			},
			"_Bo": {
				{'n', 1.0},
			},
			"ino": {
				{'n', 1.0},
			},
			"yay": {
				{'a', 1.0},
			},
			"Sal": {
				{'a', 1.0},
			},
			"uny": {
				{0, 1.0},
			},
			"Cha": {
				{'p', 1.0},
			},
			"_Us": {
				{'d', 1.0},
			},
			"ega": {
				{0, 1.0},
			},
			"ahk": {
				{'y', 1.0},
			},
			"Oji": {
				{'n', 1.0},
			},
			"_Ma": {
				{'n', 0.1429},
				{'c', 0.4286},
				{'g', 0.5714},
				{'p', 0.7143},
				{'k', 1.0},
			},
			"_Yo": {
				{'n', 1.0},
			},
			"_Ts": {
				{'u', 0.5},
				{'o', 1.0},
			},
			"oga": {
				{0, 1.0},
			},
			"akc": {
				{'h', 1.0},
			},
			"Alt": {
				{'s', 1.0},
			},
			"cha": {
				{0, 0.5},
				{'w', 0.75},
				{'h', 1.0},
			},
			"Chl": {
				{'u', 1.0},
			},
			"_Wo": {
				{'y', 1.0},
			},
			"mim": {
				{'e', 1.0},
			},
			"him": {
				{'a', 1.0},
			},
			"sil": {
				{'a', 1.0},
			},
			"___": {
				{'N', 0.0435},
				{'W', 0.1739},
				{'T', 0.2391},
				{'S', 0.3478},
				{'I', 0.3587},
				{'O', 0.3913},
				{'D', 0.4565},
				{'C', 0.5},
				{'Z', 0.5217},
				{'P', 0.5326},
				{'E', 0.5543},
				{'K', 0.5978},
				{'H', 0.6304},
				{'M', 0.7283},
				{'B', 0.7391},
				{'Y', 0.7717},
				{'G', 0.8043},
				{'A', 0.9783},
				{'U', 1.0},
			},
			"Dow": {
				{'a', 1.0},
			},
			"eek": {
				{'o', 1.0},
			},
			"nho": {
				{'w', 1.0},
			},
			"hky": {
				{'o', 1.0},
			},
			"lts": {
				{'o', 1.0},
			},
			"Bon": {
				{'i', 1.0},
			},
			"wee": {
				{0, 1.0},
			},
			"ngp": {
				{'e', 1.0},
			},
			"chi": {
				{'w', 1.0},
			},
			"_Ni": {
				{'y', 1.0},
			},
			"aba": {
				{0, 1.0},
			},
			"ana": {
				{0, 0.5},
				{'b', 1.0},
			},
			"ani": {
				{0, 0.75},
				{'t', 1.0},
			},
			"Ado": {
				{'e', 1.0},
			},
			"man": {
				{'i', 0.6667},
				{'a', 1.0},
			},
			"tas": {
				{'h', 1.0},
			},
			"int": {
				{'k', 1.0},
			},
			"_An": {
				{'a', 0.3333},
				{'g', 0.6667},
				{'p', 1.0},
			},
			"_Ki": {
				{'m', 1.0},
			},
			"_Wa": {
				{'k', 0.3333},
				{'s', 0.6667},
				{'c', 1.0},
			},
			"hlu": {
				{'m', 1.0},
			},
			"_To": {
				{'o', 1.0},
			},
			"Gog": {
				{'a', 1.0},
			},
			"_Un": {
				{'e', 1.0},
			},
			"tsi": {
				{'l', 0.5},
				{0, 1.0},
			},
			"Tsu": {
				{'l', 1.0},
			},
			"_Oj": {
				{'i', 1.0},
			},
			"aka": {
				{0, 0.25},
				{'n', 0.5},
				{'w', 0.75},
				{'y', 1.0},
			},
			"its": {
				{'i', 1.0},
			},
			"ama": {
				{'m', 0.5},
				{0, 1.0},
			},
			"_We": {
				{'a', 0.25},
				{'e', 0.5},
				{'n', 1.0},
			},
			"kay": {
				{'d', 1.0},
			},
			"Dol": {
				{'i', 1.0},
			},
			"asg": {
				{'a', 1.0},
			},
			"Tso": {
				{'m', 1.0},
			},
			"Nas": {
				{'c', 1.0},
			},
			"_In": {
				{'o', 1.0},
			},
			"how": {
				{'e', 1.0},
			},
			"Kam": {
				{'a', 1.0},
			},
			"ooa": {
				{'n', 1.0},
			},
			"asa": {
				{'n', 0.3333},
				{0, 1.0},
			},
			"_Ta": {
				{'k', 0.3333},
				{'l', 0.6667},
				{'y', 1.0},
			},
			"hyo": {
				{'k', 1.0},
			},
			"haw": {
				{'i', 0.3333},
				{'e', 1.0},
			},
			"etu": {
				{0, 1.0},
			},
			"kah": {
				{0, 1.0},
			},
			"yan": {
				{'i', 1.0},
			},
			"ayd": {
				{'a', 1.0},
			},
			"nan": {
				{'a', 1.0},
			},
			"Sah": {
				{'p', 0.5},
				{'k', 1.0},
			},
			"oni": {
				{'t', 1.0},
			},
			"nge": {
				{'e', 1.0},
			},
			"Tay": {
				{'a', 1.0},
			},
			"Tal": {
				{'u', 1.0},
			},
			"yok": {
				{'a', 1.0},
			},
			"Was": {
				{'h', 1.0},
			},
			"_Ha": {
				{'n', 0.5},
				{'l', 1.0},
			},
			"Wic": {
				{'h', 1.0},
			},
			"_Wi": {
				{'n', 0.25},
				{'c', 0.5},
				{'h', 0.75},
				{'t', 1.0},
			},
			"eko": {
				{0, 1.0},
			},
			"tah": {
				{0, 1.0},
			},
			"poo": {
				{'l', 1.0},
			},
			"nta": {
				{0, 0.5},
				{'y', 1.0},
			},
			"nit": {
				{'a', 1.0},
			},
			"hpo": {
				{'o', 1.0},
			},
			"Aga": {
				{'s', 1.0},
			},
			"ira": {
				{'c', 1.0},
			},
			"ula": {
				{0, 1.0},
			},
			"_Zi": {
				{'r', 1.0},
			},
			"eka": {
				{'h', 1.0},
			},
			"oba": {
				{0, 1.0},
			},
			"ini": {
				{'t', 1.0},
			},
			"eay": {
				{'a', 1.0},
			},
			"hiw": {
				{'i', 1.0},
			},
			"Yon": {
				{'a', 1.0},
			},
			"oma": {
				{'h', 1.0},
			},
			"hak": {
				{'a', 1.0},
			},
			"Tak": {
				{'c', 1.0},
			},
			"gpe": {
				{'t', 1.0},
			},
			"inj": {
				{'i', 1.0},
			},
			"ada": {
				{'h', 1.0},
			},
			"alu": {
				{'t', 1.0},
			},
			"Asd": {
				{'z', 1.0},
			},
			"Ino": {
				{'l', 1.0},
			},
			"awe": {
				{'e', 1.0},
			},
			"anw": {
				{'e', 1.0},
			},
			"tuh": {
				{0, 1.0},
			},
			"tte": {
				{0, 1.0},
			},
			"nji": {
				{'n', 1.0},
			},
			"tso": {
				{'b', 1.0},
			},
			"_Sh": {
				{'a', 0.25},
				{'i', 1.0},
			},
			"ezb": {
				{'a', 1.0},
			},
			"iaw": {
				{'a', 1.0},
			},
			"ach": {
				{'i', 0.5},
				{'a', 1.0},
			},
			"_Ya": {
				{'z', 0.5},
				{'n', 1.0},
			},
			"shn": {
				{'a', 1.0},
			},
			"mad": {
				{'a', 1.0},
			},
			"kaw": {
				{'e', 1.0},
			},
			"Dez": {
				{'b', 1.0},
			},
			"Wen": {
				{'o', 1.0},
			},
			"see": {
				{0, 1.0},
			},
			"tek": {
				{'a', 1.0},
			},
			"dah": {
				{'y', 1.0},
			},
			"__S": {
				{'a', 0.4},
				{'h', 0.8},
				{'i', 0.9},
				{'n', 1.0},
			},
			"ass": {
				{'e', 1.0},
			},
			"Ote": {
				{'k', 1.0},
			},
			"sdz": {
				{'a', 1.0},
			},
			"aca": {
				{'w', 1.0},
			},
			"Zir": {
				{'a', 1.0},
			},
			"ont": {
				{'a', 1.0},
			},
			"ett": {
				{'e', 1.0},
			},
			"Dob": {
				{'a', 1.0},
			},
			"lum": {
				{'a', 1.0},
			},
			"anh": {
				{'o', 1.0},
			},
			"Awi": {
				{'n', 1.0},
			},
			"mel": {
				{'a', 1.0},
			},
			"too": {
				{0, 1.0},
			},
			"adi": {
				{0, 1.0},
			},
			"Sas": {
				{'a', 1.0},
			},
			"_Pt": {
				{'a', 1.0},
			},
			"_Eh": {
				{'a', 1.0},
			},
			"ezh": {
				{'i', 1.0},
			},
			"eez": {
				{'h', 1.0},
			},
			"Gal": {
				{'i', 1.0},
			},
			"Wee": {
				{'k', 1.0},
			},
			"_Hi": {
				{'a', 1.0},
			},
			"__D": {
				{'e', 0.1667},
				{'i', 0.3333},
				{'o', 1.0},
			},
			"Noy": {
				{'a', 1.0},
			},
			"awa": {
				{'s', 1.0},
			},
			"ayt": {
				{'o', 1.0},
			},
			"gas": {
				{'k', 0.5},
				{'g', 1.0},
			},
			"Mag": {
				{'a', 1.0},
			},
			"ima": {
				{'n', 0.3333},
				{0, 0.6667},
				{'s', 1.0},
			},
			"__A": {
				{'n', 0.1875},
				{'m', 0.3125},
				{'g', 0.375},
				{'l', 0.4375},
				{'y', 0.5},
				{'w', 0.625},
				{'d', 0.75},
				{'s', 0.8125},
				{'t', 0.875},
				{'h', 1.0},
			},
			"Too": {
				{'a', 1.0},
			},
			"paw": {
				{'e', 1.0},
			},
			"hum": {
				{'a', 1.0},
			},
			"jin": {
				{'t', 0.5},
				{'j', 1.0},
			},
			"lil": {
				{'a', 1.0},
			},
			"oya": {
				{0, 1.0},
			},
			"iyo": {
				{'l', 1.0},
			},
			"_Ag": {
				{'a', 1.0},
			},
			"dsi": {
				{'l', 1.0},
			},
			"hah": {
				{'p', 1.0},
			},
			"alo": {
				{'k', 1.0},
			},
			"_Ka": {
				{'m', 0.3333},
				{'i', 0.6667},
				{'n', 1.0},
			},
			"Han": {
				{'t', 1.0},
			},
			"ime": {
				{'l', 1.0},
			},
			"Pta": {
				{'y', 1.0},
			},
			"pet": {
				{'u', 1.0},
			},
			"ali": {
				{0, 0.5},
				{'l', 1.0},
			},
			"__Z": {
				{'o', 0.5},
				{'i', 1.0},
			},
			"Wit": {
				{'a', 1.0},
			},
			"gee": {
				{0, 1.0},
			},
			"ita": {
				{'s', 0.2},
				{0, 1.0},
			},
			"acu": {
				{'n', 1.0},
			},
			"tka": {
				{0, 1.0},
			},
			"doe": {
				{'t', 1.0},
			},
			"Anp": {
				{'a', 1.0},
			},
			"Usd": {
				{'i', 1.0},
			},
			"__M": {
				{'a', 0.7778},
				{'i', 0.8889},
				{'o', 1.0},
			},
			"ant": {
				{'a', 0.5},
				{'u', 1.0},
			},
			"aga": {
				{'s', 1.0},
			},
			"Min": {
				{'a', 1.0},
			},
			"_Al": {
				{'t', 1.0},
			},
			"Kim": {
				{'i', 1.0},
			},
			"iya": {
				{0, 1.0},
			},
			"ila": {
				{'h', 0.3333},
				{0, 1.0},
			},
			"Mac": {
				{'h', 0.5},
				{'a', 1.0},
			},
			"_Ey": {
				{'o', 1.0},
			},
			"wen": {
				{'a', 1.0},
			},
			"_Am": {
				{'a', 1.0},
			},
			"Yan": {
				{'a', 1.0},
			},
			"ang": {
				{'e', 1.0},
			},
			"ala": {
				{'l', 1.0},
			},
			"Shi": {
				{'m', 0.6667},
				{'d', 1.0},
			},
			"oke": {
				{0, 1.0},
			},
			"Yaz": {
				{'h', 1.0},
			},
			"piy": {
				{'a', 1.0},
			},
			"nda": {
				{0, 1.0},
			},
			"ota": {
				{0, 1.0},
			},
			"api": {
				{'y', 1.0},
			},
			"eno": {
				{'n', 1.0},
			},
			"__H": {
				{'i', 0.3333},
				{'a', 1.0},
			},
			"Wac": {
				{'h', 1.0},
			},
			"iha": {
				{'k', 1.0},
			},
			"Wea": {
				{'y', 1.0},
			},
			"asc": {
				{'h', 1.0},
			},
			"ysa": {
				{'n', 1.0},
			},
			"non": {
				{'a', 1.0},
			},
			"Kai": {
				{0, 1.0},
			},
			"sdi": {
				{0, 1.0},
			},
			"__W": {
				{'e', 0.3333},
				{'a', 0.5833},
				{'o', 0.6667},
				{'i', 1.0},
			},
			"som": {
				{'a', 1.0},
			},
			"oet": {
				{'t', 1.0},
			},
			"Kan": {
				{'g', 1.0},
			},
			"kan": {
				{'d', 1.0},
			},
			"hap": {
				{'a', 1.0},
			},
			"oka": {
				{0, 1.0},
			},
			"ask": {
				{'a', 1.0},
			},
			"lut": {
				{'a', 1.0},
			},
			"nah": {
				{0, 1.0},
			},
			"owe": {
				{'e', 1.0},
			},
			"dee": {
				{'z', 1.0},
			},
			"Ana": {
				{'b', 1.0},
			},
			"Dib": {
				{'e', 1.0},
			},
			"lje": {
				{'e', 1.0},
			},
			"nab": {
				{'a', 1.0},
			},
			"owa": {
				{'n', 1.0},
			},
			"ibe": {
				{0, 1.0},
			},
			"Win": {
				{'o', 1.0},
			},
			"aya": {
				{'y', 0.3333},
				{0, 0.6667},
				{'n', 1.0},
			},
			"yit": {
				{'a', 1.0},
			},
			"__O": {
				{'t', 0.3333},
				{'j', 0.6667},
				{'o', 1.0},
			},
			"Mos": {
				{'i', 1.0},
			},
			"_Ch": {
				{'u', 0.25},
				{'a', 0.75},
				{'l', 1.0},
			},
			"_Ah": {
				{'y', 0.5},
				{'a', 1.0},
			},
			"Ahy": {
				{'o', 1.0},
			},
			"Doy": {
				{'a', 1.0},
			},
			"Map": {
				{'i', 1.0},
			},
			"_Zo": {
				{'n', 1.0},
			},
			"Ama": {
				{'d', 0.5},
				{0, 1.0},
			},
			"tay": {
				{'w', 0.5},
				{'s', 1.0},
			},
			"_Mi": {
				{'n', 1.0},
			},
			"Awe": {
				{'n', 1.0},
			},
			"win": {
				{'i', 1.0},
			},
			"_Di": {
				{'b', 1.0},
			},
			"oan": {
				{'t', 1.0},
			},
			"sob": {
				{'a', 1.0},
			},
			"Sit": {
				{'s', 1.0},
			},
			"Hal": {
				{'o', 1.0},
			},
			"hta": {
				{0, 1.0},
			},
			"Ang": {
				{'p', 1.0},
			},
			"Nah": {
				{'i', 1.0},
			},
			"ahp": {
				{'i', 0.5},
				{'o', 1.0},
			},
			"ntu": {
				{'h', 1.0},
			},
			"_Ga": {
				{'l', 1.0},
			},
			"__U": {
				{'n', 0.5},
				{'s', 1.0},
			},
			"apa": {
				{'w', 0.5},
				{0, 1.0},
			},
		}
	case 'M':
		return map[string][]nameFragment{
			"Una": {
				{'d', 1.0},
			},
			"nac": {
				{'o', 1.0},
			},
			"nag": {
				{'e', 1.0},
			},
			"yto": {
				{'n', 1.0},
			},
			"nto": {
				{'n', 1.0},
			},
			"ela": {
				{0, 1.0},
			},
			"stu": {
				{0, 1.0},
			},
			"var": {
				{'i', 1.0},
			},
			"uli": {
				{0, 1.0},
			},
			"was": {
				{'h', 1.0},
			},
			"Woh": {
				{'a', 1.0},
			},
			"hay": {
				{'t', 1.0},
			},
			"ays": {
				{'h', 1.0},
			},
			"hto": {
				{0, 0.5},
				{'n', 1.0},
			},
			"ahy": {
				{0, 1.0},
			},
			"ulk": {
				{'a', 1.0},
			},
			"Sha": {
				{'p', 1.0},
			},
			"uah": {
				{0, 1.0},
			},
			"Koh": {
				{'a', 1.0},
			},
			"sap": {
				{'a', 1.0},
			},
			"_De": {
				{'g', 1.0},
			},
			"_Na": {
				{'s', 0.2},
				{'n', 0.4},
				{'a', 0.6},
				{'p', 0.8},
				{'k', 1.0},
			},
			"Sic": {
				{'h', 1.0},
			},
			"Dig": {
				{'h', 1.0},
			},
			"aad": {
				{0, 1.0},
			},
			"Ahu": {
				{'l', 1.0},
			},
			"ink": {
				{'s', 1.0},
			},
			"Ada": {
				{'h', 1.0},
			},
			"_Ad": {
				{'o', 0.3333},
				{'i', 0.6667},
				{'a', 1.0},
			},
			"sah": {
				{'a', 1.0},
			},
			"nii": {
				{0, 1.0},
			},
			"gii": {
				{0, 1.0},
			},
			"__N": {
				{'a', 0.7143},
				{'i', 1.0},
			},
			"Tas": {
				{'h', 1.0},
			},
			"Sin": {
				{'t', 1.0},
			},
			"_Si": {
				{'k', 0.3333},
				{'n', 0.6667},
				{'c', 1.0},
			},
			"Gaw": {
				{'o', 1.0},
			},
			"Wes": {
				{'a', 1.0},
			},
			"Mik": {
				{'a', 1.0},
			},
			"ezi": {
				{0, 1.0},
			},
			"nka": {
				{'g', 0.25},
				{0, 1.0},
			},
			"_Ot": {
				{'a', 1.0},
			},
			"Tse": {
				{0, 0.5},
				{'l', 1.0},
			},
			"eto": {
				{'n', 1.0},
			},
			"Wah": {
				{'k', 0.5},
				{'c', 1.0},
			},
			"iil": {
				{0, 1.0},
			},
			"mbl": {
				{'e', 1.0},
			},
			"wes": {
				{'c', 1.0},
			},
			"equ": {
				{'o', 0.5},
				{'a', 1.0},
			},
			"aza": {
				{0, 1.0},
			},
			"_Sa": {
				{'n', 0.5},
				{'t', 1.0},
			},
			"esh": {
				{'a', 0.5},
				{'i', 1.0},
			},
			"ska": {
				{0, 0.7143},
				{'h', 0.8571},
				{'m', 1.0},
			},
			"lah": {
				{0, 1.0},
			},
			"_Co": {
				{'o', 1.0},
			},
			"una": {
				{0, 1.0},
			},
			"yol": {
				{0, 1.0},
			},
			"_Do": {
				{'h', 1.0},
			},
			"sht": {
				{'a', 1.0},
			},
			"zee": {
				{0, 1.0},
			},
			"ish": {
				{0, 1.0},
			},
			"___": {
				{'P', 0.016},
				{'E', 0.04},
				{'Y', 0.08},
				{'B', 0.112},
				{'C', 0.176},
				{'W', 0.296},
				{'O', 0.352},
				{'S', 0.448},
				{'D', 0.504},
				{'U', 0.52},
				{'T', 0.648},
				{'L', 0.664},
				{'N', 0.72},
				{'G', 0.76},
				{'M', 0.816},
				{'A', 0.928},
				{'K', 0.968},
				{'H', 1.0},
			},
			"Yas": {
				{0, 1.0},
			},
			"ite": {
				{'k', 1.0},
			},
			"chi": {
				{'n', 0.6667},
				{'i', 1.0},
			},
			"lni": {
				{'s', 1.0},
			},
			"isk": {
				{'a', 1.0},
			},
			"tas": {
				{0, 1.0},
			},
			"int": {
				{'o', 0.5},
				{'e', 1.0},
			},
			"aht": {
				{'o', 1.0},
			},
			"ana": {
				{'g', 0.25},
				{0, 0.75},
				{'h', 1.0},
			},
			"_Ni": {
				{'i', 0.5},
				{'y', 1.0},
			},
			"akt": {
				{'e', 0.5},
				{'a', 1.0},
			},
			"_Wa": {
				{'m', 0.1818},
				{'n', 0.5455},
				{'h', 0.9091},
				{'y', 1.0},
			},
			"uit": {
				{'o', 1.0},
			},
			"aka": {
				{'n', 0.5},
				{'i', 1.0},
			},
			"its": {
				{'a', 1.0},
			},
			"tsi": {
				{'d', 1.0},
			},
			"_Ko": {
				{'h', 1.0},
			},
			"iki": {
				{'y', 1.0},
			},
			"hii": {
				{0, 1.0},
			},
			"kai": {
				{0, 1.0},
			},
			"_Ta": {
				{'t', 0.125},
				{'k', 0.25},
				{'y', 0.375},
				{'p', 0.5},
				{'h', 0.875},
				{'s', 1.0},
			},
			"ooa": {
				{'n', 1.0},
			},
			"hka": {
				{'n', 0.6667},
				{'h', 1.0},
			},
			"Tee": {
				{'t', 1.0},
			},
			"Nas": {
				{'t', 1.0},
			},
			"_Sk": {
				{'a', 1.0},
			},
			"ast": {
				{'i', 0.5},
				{'a', 1.0},
			},
			"kah": {
				{0, 1.0},
			},
			"_Ap": {
				{'i', 1.0},
			},
			"Wha": {
				{'k', 1.0},
			},
			"wic": {
				{'a', 1.0},
			},
			"oai": {
				{0, 1.0},
			},
			"pco": {
				{0, 1.0},
			},
			"dut": {
				{'i', 1.0},
			},
			"oee": {
				{'t', 1.0},
			},
			"_Go": {
				{'m', 1.0},
			},
			"uoy": {
				{'a', 1.0},
			},
			"kot": {
				{'a', 1.0},
			},
			"ari": {
				{'a', 1.0},
			},
			"Niy": {
				{'o', 1.0},
			},
			"eas": {
				{'e', 1.0},
			},
			"imi": {
				{'k', 1.0},
			},
			"aku": {
				{'w', 0.3333},
				{'l', 1.0},
			},
			"iic": {
				{'h', 1.0},
			},
			"ull": {
				{'a', 1.0},
			},
			"hig": {
				{'a', 1.0},
			},
			"oga": {
				{0, 1.0},
			},
			"Tap": {
				{'c', 1.0},
			},
			"_Oh": {
				{'i', 0.6667},
				{'a', 1.0},
			},
			"_Ts": {
				{'i', 0.25},
				{'e', 0.75},
				{'o', 1.0},
			},
			"onu": {
				{0, 1.0},
			},
			"hat": {
				{'a', 0.6667},
				{'e', 1.0},
			},
			"eni": {
				{0, 1.0},
			},
			"_Bi": {
				{'s', 0.3333},
				{'d', 0.6667},
				{'l', 1.0},
			},
			"ust": {
				{'u', 1.0},
			},
			"idz": {
				{'i', 1.0},
			},
			"sin": {
				{0, 1.0},
			},
			"ete": {
				{0, 1.0},
			},
			"cha": {
				{'a', 1.0},
			},
			"ano": {
				{'w', 0.5},
				{'s', 1.0},
			},
			"oto": {
				{'g', 1.0},
			},
			"eti": {
				{'m', 1.0},
			},
			"Kla": {
				{'h', 1.0},
			},
			"han": {
				{'z', 0.3333},
				{'k', 0.6667},
				{'a', 1.0},
			},
			"shk": {
				{'i', 1.0},
			},
			"_Yi": {
				{'s', 1.0},
			},
			"isi": {
				{'n', 1.0},
			},
			"Oga": {
				{'l', 1.0},
			},
			"_Pa": {
				{'y', 1.0},
			},
			"hch": {
				{'i', 1.0},
			},
			"aag": {
				{'i', 1.0},
			},
			"adi": {
				{0, 1.0},
			},
			"soa": {
				{'i', 1.0},
			},
			"kte": {
				{0, 1.0},
			},
			"Tat": {
				{'o', 1.0},
			},
			"osk": {
				{'a', 1.0},
			},
			"aan": {
				{'a', 1.0},
			},
			"Ahi": {
				{'g', 1.0},
			},
			"tos": {
				{'k', 1.0},
			},
			"kii": {
				{0, 1.0},
			},
			"nze": {
				{'e', 1.0},
			},
			"_En": {
				{'a', 1.0},
			},
			"pia": {
				{'t', 1.0},
			},
			"uwa": {
				{0, 1.0},
			},
			"__D": {
				{'a', 0.1429},
				{'u', 0.2857},
				{'o', 0.5714},
				{'e', 0.7143},
				{'i', 1.0},
			},
			"ton": {
				{'k', 0.4286},
				{0, 0.8571},
				{'g', 1.0},
			},
			"Loo": {
				{'t', 1.0},
			},
			"Adi": {
				{'t', 1.0},
			},
			"Nan": {
				{'t', 1.0},
			},
			"aty": {
				{0, 1.0},
			},
			"idi": {
				{0, 1.0},
			},
			"Ato": {
				{'h', 1.0},
			},
			"hni": {
				{0, 1.0},
			},
			"tan": {
				{'g', 0.1429},
				{0, 0.7143},
				{'t', 0.8571},
				{'w', 1.0},
			},
			"haa": {
				{'d', 1.0},
			},
			"now": {
				{'i', 1.0},
			},
			"oya": {
				{'h', 1.0},
			},
			"anu": {
				{'n', 1.0},
			},
			"iwa": {
				{'l', 1.0},
			},
			"coo": {
				{'w', 1.0},
			},
			"omd": {
				{'a', 1.0},
			},
			"oni": {
				{'i', 1.0},
			},
			"ngy": {
				{'a', 1.0},
			},
			"nik": {
				{'i', 1.0},
			},
			"Wam": {
				{'b', 1.0},
			},
			"San": {
				{'i', 1.0},
			},
			"sti": {
				{0, 0.5},
				{'i', 1.0},
			},
			"min": {
				{'z', 1.0},
			},
			"_Ha": {
				{'n', 0.5},
				{'s', 1.0},
			},
			"nga": {
				{0, 1.0},
			},
			"iat": {
				{'a', 0.5},
				{'i', 1.0},
			},
			"mme": {
				{'d', 1.0},
			},
			"ahe": {
				{'t', 1.0},
			},
			"nit": {
				{'a', 1.0},
			},
			"Mam": {
				{'m', 1.0},
			},
			"eay": {
				{'a', 1.0},
			},
			"tag": {
				{'u', 1.0},
			},
			"iga": {
				{0, 1.0},
			},
			"kod": {
				{'a', 1.0},
			},
			"hak": {
				{'a', 1.0},
			},
			"Ata": {
				{'g', 1.0},
			},
			"oma": {
				{0, 1.0},
			},
			"tuh": {
				{0, 1.0},
			},
			"anw": {
				{'a', 1.0},
			},
			"adu": {
				{'t', 1.0},
			},
			"_Ya": {
				{'n', 0.5},
				{'h', 0.75},
				{'s', 1.0},
			},
			"Gui": {
				{'t', 1.0},
			},
			"ngv": {
				{'a', 1.0},
			},
			"kec": {
				{'h', 1.0},
			},
			"Ska": {
				{'h', 1.0},
			},
			"hin": {
				{'k', 0.3333},
				{0, 0.6667},
				{'t', 1.0},
			},
			"nah": {
				{'t', 1.0},
			},
			"owa": {
				{'h', 0.6667},
				{'s', 1.0},
			},
			"ata": {
				{'n', 1.0},
			},
			"tsa": {
				{'n', 0.5},
				{'d', 1.0},
			},
			"_Ch": {
				{'e', 0.2},
				{'a', 1.0},
			},
			"ike": {
				{0, 1.0},
			},
			"Ona": {
				{'c', 1.0},
			},
			"hiy": {
				{'e', 1.0},
			},
			"Lal": {
				{'l', 1.0},
			},
			"anz": {
				{'e', 1.0},
			},
			"gul": {
				{'k', 1.0},
			},
			"ank": {
				{'o', 1.0},
			},
			"tay": {
				{0, 1.0},
			},
			"gal": {
				{'e', 1.0},
			},
			"_Mi": {
				{'k', 1.0},
			},
			"_Da": {
				{'k', 1.0},
			},
			"oan": {
				{'t', 1.0},
			},
			"_Di": {
				{'g', 0.5},
				{'w', 1.0},
			},
			"yes": {
				{'a', 1.0},
			},
			"hom": {
				{'a', 1.0},
			},
			"Che": {
				{'a', 1.0},
			},
			"ppa": {
				{0, 1.0},
			},
			"__U": {
				{'s', 0.5},
				{'n', 1.0},
			},
			"ahp": {
				{'e', 1.0},
			},
			"Att": {
				{'a', 1.0},
			},
			"sid": {
				{'i', 1.0},
			},
			"ali": {
				{0, 1.0},
			},
			"How": {
				{'a', 1.0},
			},
			"__M": {
				{'a', 0.8571},
				{'i', 1.0},
			},
			"Bes": {
				{'h', 1.0},
			},
			"ita": {
				{0, 1.0},
			},
			"ato": {
				{'s', 0.3333},
				{0, 0.6667},
				{'n', 1.0},
			},
			"ila": {
				{'g', 0.5},
				{'h', 1.0},
			},
			"iya": {
				{0, 1.0},
			},
			"zii": {
				{'l', 1.0},
			},
			"ant": {
				{'a', 0.6667},
				{'u', 1.0},
			},
			"Pay": {
				{'t', 1.0},
			},
			"_Og": {
				{'a', 1.0},
			},
			"ang": {
				{'e', 0.5},
				{'y', 1.0},
			},
			"_Ey": {
				{'a', 1.0},
			},
			"ahc": {
				{'h', 1.0},
			},
			"dat": {
				{'y', 1.0},
			},
			"nis": {
				{'i', 0.5},
				{'h', 1.0},
			},
			"lak": {
				{'u', 1.0},
			},
			"aho": {
				{'m', 1.0},
			},
			"oke": {
				{0, 1.0},
			},
			"yta": {
				{'h', 1.0},
			},
			"iin": {
				{0, 1.0},
			},
			"__H": {
				{'a', 0.5},
				{'o', 1.0},
			},
			"kta": {
				{'y', 1.0},
			},
			"ech": {
				{'e', 1.0},
			},
			"zim": {
				{0, 1.0},
			},
			"__W": {
				{'h', 0.0667},
				{'a', 0.8},
				{'o', 0.8667},
				{'e', 1.0},
			},
			"onk": {
				{'a', 1.0},
			},
			"_Pe": {
				{'z', 1.0},
			},
			"ask": {
				{'a', 1.0},
			},
			"awo": {
				{'n', 1.0},
			},
			"kan": {
				{0, 1.0},
			},
			"Sat": {
				{'a', 1.0},
			},
			"oot": {
				{'a', 1.0},
			},
			"_At": {
				{'o', 0.2},
				{'s', 0.6},
				{'a', 0.8},
				{'t', 1.0},
			},
			"oho": {
				{'s', 1.0},
			},
			"ich": {
				{'e', 0.5},
				{'a', 1.0},
			},
			"_Ca": {
				{'n', 1.0},
			},
			"esc": {
				{'o', 1.0},
			},
			"has": {
				{'k', 1.0},
			},
			"nks": {
				{'a', 1.0},
			},
			"wak": {
				{'u', 1.0},
			},
			"Oda": {
				{'k', 1.0},
			},
			"__T": {
				{'a', 0.5},
				{'o', 0.6875},
				{'e', 0.75},
				{'s', 1.0},
			},
			"Tsi": {
				{'y', 1.0},
			},
			"all": {
				{'o', 1.0},
			},
			"lka": {
				{'l', 1.0},
			},
			"kiy": {
				{'a', 0.5},
				{0, 1.0},
			},
			"pay": {
				{'s', 0.5},
				{0, 1.0},
			},
			"oko": {
				{'t', 1.0},
			},
			"the": {
				{'e', 1.0},
			},
			"eet": {
				{'e', 0.5},
				{'o', 1.0},
			},
			"hee": {
				{'n', 1.0},
			},
			"hei": {
				{'i', 1.0},
			},
			"kag": {
				{'y', 1.0},
			},
			"ria": {
				{'t', 1.0},
			},
			"agi": {
				{'i', 1.0},
			},
			"het": {
				{'o', 0.5},
				{'a', 1.0},
			},
			"ash": {
				{'t', 0.5},
				{'u', 1.0},
			},
			"osa": {
				{'n', 0.5},
				{0, 1.0},
			},
			"Yis": {
				{'k', 1.0},
			},
			"Ats": {
				{'a', 0.5},
				{'i', 1.0},
			},
			"nte": {
				{0, 1.0},
			},
			"_As": {
				{'h', 1.0},
			},
			"ees": {
				{'k', 0.6667},
				{'h', 1.0},
			},
			"__G": {
				{'a', 0.6},
				{'u', 0.8},
				{'o', 1.0},
			},
			"Ena": {
				{'p', 1.0},
			},
			"quo": {
				{'y', 1.0},
			},
			"kul": {
				{'l', 1.0},
			},
			"Eya": {
				{'n', 1.0},
			},
			"__K": {
				{'o', 0.2},
				{'l', 0.4},
				{'i', 0.6},
				{'a', 1.0},
			},
			"hos": {
				{'a', 1.0},
			},
			"Esk": {
				{'a', 1.0},
			},
			"__Y": {
				{'a', 0.8},
				{'i', 1.0},
			},
			"lee": {
				{'s', 0.6667},
				{0, 1.0},
			},
			"Mah": {
				{'p', 0.5},
				{'k', 1.0},
			},
			"Way": {
				{'a', 1.0},
			},
			"ona": {
				{0, 1.0},
			},
			"ilc": {
				{'h', 1.0},
			},
			"tta": {
				{'k', 1.0},
			},
			"nap": {
				{'a', 1.0},
			},
			"_Du": {
				{'s', 1.0},
			},
			"__C": {
				{'h', 0.625},
				{'a', 0.75},
				{'o', 0.875},
				{'e', 1.0},
			},
			"app": {
				{'a', 1.0},
			},
			"Bis": {
				{'a', 1.0},
			},
			"tog": {
				{'a', 1.0},
			},
			"__P": {
				{'e', 0.5},
				{'a', 1.0},
			},
			"san": {
				{0, 1.0},
			},
			"__B": {
				{'e', 0.25},
				{'i', 1.0},
			},
			"Dus": {
				{'t', 1.0},
			},
			"mda": {
				{0, 1.0},
			},
			"_Wh": {
				{'a', 1.0},
			},
			"ghi": {
				{'n', 1.0},
			},
			"hul": {
				{'i', 1.0},
			},
			"wal": {
				{'i', 1.0},
			},
			"ako": {
				{'t', 0.6667},
				{'d', 1.0},
			},
			"dzi": {
				{'i', 1.0},
			},
			"lth": {
				{'e', 1.0},
			},
			"mik": {
				{'a', 1.0},
			},
			"wah": {
				{0, 0.5},
				{'k', 1.0},
			},
			"kuw": {
				{'a', 1.0},
			},
			"_Ki": {
				{'l', 1.0},
			},
			"tak": {
				{'t', 0.5},
				{'u', 1.0},
			},
			"Nii": {
				{'c', 1.0},
			},
			"_An": {
				{'g', 1.0},
			},
			"ani": {
				{'s', 0.1667},
				{0, 0.5},
				{'t', 0.6667},
				{'k', 1.0},
			},
			"Ado": {
				{'e', 1.0},
			},
			"_Un": {
				{'a', 1.0},
			},
			"_To": {
				{'o', 0.3333},
				{'k', 1.0},
			},
			"kal": {
				{'u', 1.0},
			},
			"yah": {
				{0, 1.0},
			},
			"esk": {
				{'a', 1.0},
			},
			"Yah": {
				{'t', 1.0},
			},
			"tio": {
				{'n', 1.0},
			},
			"Diw": {
				{'a', 1.0},
			},
			"aal": {
				{'n', 1.0},
			},
			"seq": {
				{'u', 1.0},
			},
			"lla": {
				{0, 0.5},
				{'k', 1.0},
			},
			"_We": {
				{'a', 0.5},
				{'s', 1.0},
			},
			"_Lo": {
				{'o', 1.0},
			},
			"Tah": {
				{'a', 0.3333},
				{'e', 0.6667},
				{'o', 1.0},
			},
			"_Ce": {
				{'t', 1.0},
			},
			"Tso": {
				{'a', 1.0},
			},
			"sad": {
				{'i', 1.0},
			},
			"Wan": {
				{'a', 0.5},
				{'i', 1.0},
			},
			"yan": {
				{'o', 0.5},
				{'i', 1.0},
			},
			"Ota": {
				{'k', 1.0},
			},
			"amb": {
				{'l', 1.0},
			},
			"_Te": {
				{'e', 1.0},
			},
			"owi": {
				{'c', 1.0},
			},
			"eda": {
				{'t', 1.0},
			},
			"ale": {
				{'e', 1.0},
			},
			"pee": {
				{0, 1.0},
			},
			"sco": {
				{'o', 1.0},
			},
			"__E": {
				{'s', 0.3333},
				{'y', 0.6667},
				{'n', 1.0},
			},
			"Deg": {
				{'o', 1.0},
			},
			"ans": {
				{'a', 0.5},
				{'k', 1.0},
			},
			"sel": {
				{'a', 1.0},
			},
			"_Es": {
				{'k', 1.0},
			},
			"lch": {
				{'i', 1.0},
			},
			"yay": {
				{'a', 1.0},
			},
			"won": {
				{'i', 1.0},
			},
			"Gad": {
				{0, 1.0},
			},
			"hki": {
				{'i', 1.0},
			},
			"_Us": {
				{'t', 1.0},
			},
			"_Ak": {
				{'e', 1.0},
			},
			"_Be": {
				{'s', 1.0},
			},
			"Cha": {
				{'s', 0.25},
				{'n', 0.5},
				{'t', 0.75},
				{'y', 1.0},
			},
			"ohi": {
				{0, 1.0},
			},
			"eii": {
				{0, 1.0},
			},
			"_La": {
				{'l', 1.0},
			},
			"aha": {
				{'t', 0.5},
				{'l', 1.0},
			},
			"_Od": {
				{'a', 1.0},
			},
			"Mat": {
				{'o', 1.0},
			},
			"_Ma": {
				{'h', 0.3333},
				{'z', 0.5},
				{'t', 0.8333},
				{'m', 1.0},
			},
			"oow": {
				{'e', 0.5},
				{'a', 1.0},
			},
			"ahk": {
				{'a', 0.75},
				{'o', 1.0},
			},
			"_Wo": {
				{'h', 1.0},
			},
			"_Se": {
				{'q', 0.3333},
				{'t', 1.0},
			},
			"nos": {
				{'a', 1.0},
			},
			"__S": {
				{'a', 0.1667},
				{'h', 0.4167},
				{'k', 0.5},
				{'i', 0.75},
				{'e', 1.0},
			},
			"Ohi": {
				{'y', 0.5},
				{'t', 1.0},
			},
			"tek": {
				{'a', 1.0},
			},
			"dah": {
				{'y', 1.0},
			},
			"med": {
				{'a', 1.0},
			},
			"ong": {
				{'a', 1.0},
			},
			"ble": {
				{'e', 1.0},
			},
			"Bid": {
				{'z', 1.0},
			},
			"Gom": {
				{'d', 1.0},
			},
			"Oha": {
				{'n', 1.0},
			},
			"gva": {
				{'r', 1.0},
			},
			"toh": {
				{'i', 1.0},
			},
			"Nap": {
				{'a', 1.0},
			},
			"isa": {
				{'h', 1.0},
			},
			"agy": {
				{'a', 1.0},
			},
			"ase": {
				{'q', 1.0},
			},
			"_Ho": {
				{'w', 0.5},
				{'t', 1.0},
			},
			"ysh": {
				{'n', 1.0},
			},
			"esa": {
				{0, 1.0},
			},
			"Ust": {
				{'i', 1.0},
			},
			"Doh": {
				{'a', 0.5},
				{'o', 1.0},
			},
			"qua": {
				{'h', 1.0},
			},
			"gaa": {
				{'n', 1.0},
			},
			"ayt": {
				{'a', 0.5},
				{'o', 1.0},
			},
			"iye": {
				{0, 0.5},
				{'s', 1.0},
			},
			"iyo": {
				{'l', 1.0},
			},
			"__A": {
				{'p', 0.0714},
				{'h', 0.2143},
				{'n', 0.2857},
				{'k', 0.3571},
				{'t', 0.7143},
				{'s', 0.7857},
				{'d', 1.0},
			},
			"Too": {
				{'a', 1.0},
			},
			"_Ka": {
				{'n', 1.0},
			},
			"Han": {
				{'s', 1.0},
			},
			"aco": {
				{'n', 1.0},
			},
			"sha": {
				{0, 1.0},
			},
			"Tay": {
				{'a', 1.0},
			},
			"nge": {
				{'e', 1.0},
			},
			"shi": {
				{'l', 1.0},
			},
			"shu": {
				{'n', 1.0},
			},
			"nta": {
				{0, 0.5},
				{'n', 1.0},
			},
			"tah": {
				{0, 1.0},
			},
			"dak": {
				{'o', 1.0},
			},
			"nad": {
				{'u', 1.0},
			},
			"kam": {
				{'i', 1.0},
			},
			"Ash": {
				{'k', 1.0},
			},
			"eka": {
				{'h', 1.0},
			},
			"Maz": {
				{'a', 1.0},
			},
			"igh": {
				{'i', 1.0},
			},
			"gya": {
				{0, 1.0},
			},
			"Hot": {
				{'a', 1.0},
			},
			"_Gu": {
				{'i', 1.0},
			},
			"__L": {
				{'o', 0.5},
				{'a', 1.0},
			},
			"Bil": {
				{'a', 1.0},
			},
			"alu": {
				{0, 1.0},
			},
			"Coo": {
				{'w', 1.0},
			},
			"Tak": {
				{'o', 1.0},
			},
			"hit": {
				{'e', 1.0},
			},
			"lan": {
				{'i', 1.0},
			},
			"Ake": {
				{'c', 1.0},
			},
			"ica": {
				{'k', 1.0},
			},
			"shn": {
				{'i', 1.0},
			},
			"dit": {
				{'s', 1.0},
			},
			"_Sh": {
				{'i', 0.6667},
				{'a', 1.0},
			},
			"hil": {
				{'a', 0.5},
				{'t', 1.0},
			},
			"owe": {
				{'s', 0.5},
				{0, 1.0},
			},
			"ito": {
				{'n', 1.0},
			},
			"Set": {
				{'i', 0.5},
				{'a', 1.0},
			},
			"ami": {
				{'n', 1.0},
			},
			"ion": {
				{'u', 1.0},
			},
			"koo": {
				{'w', 1.0},
			},
			"_Ah": {
				{'u', 0.5},
				{'i', 1.0},
			},
			"ate": {
				{0, 1.0},
			},
			"__O": {
				{'g', 0.1429},
				{'d', 0.2857},
				{'h', 0.7143},
				{'n', 0.8571},
				{'t', 1.0},
			},
			"aya": {
				{'n', 0.25},
				{0, 0.75},
				{'y', 1.0},
			},
			"oda": {
				{0, 1.0},
			},
			"got": {
				{'o', 1.0},
			},
			"lag": {
				{'a', 1.0},
			},
			"unk": {
				{'a', 1.0},
			},
			"iyi": {
				{0, 1.0},
			},
			"nzi": {
				{'m', 1.0},
			},
			"hal": {
				{'a', 0.5},
				{'i', 1.0},
			},
			"Sik": {
				{'e', 1.0},
			},
			"age": {
				{'e', 1.0},
			},
			"hun": {
				{'k', 1.0},
			},
			"oha": {
				{'n', 0.3333},
				{'t', 0.6667},
				{'l', 1.0},
			},
			"tii": {
				{'n', 1.0},
			},
			"apc": {
				{'o', 1.0},
			},
			"nwa": {
				{'k', 1.0},
			},
			"Kil": {
				{'c', 1.0},
			},
			"hta": {
				{'y', 1.0},
			},
			"eta": {
				{'n', 0.6667},
				{0, 1.0},
			},
			"Can": {
				{'o', 1.0},
			},
			"apa": {
				{0, 0.3333},
				{'y', 1.0},
			},
			"agu": {
				{'l', 1.0},
			},
			"ntu": {
				{'h', 1.0},
			},
			"_Ga": {
				{'w', 0.3333},
				{'a', 0.6667},
				{'d', 1.0},
			},
			"_On": {
				{'a', 1.0},
			},
			"_Kl": {
				{'a', 1.0},
			},
			"Ang": {
				{'v', 1.0},
			},
			"nun": {
				{'a', 1.0},
			},
			"ati": {
				{'o', 1.0},
			},
			"cak": {
				{'t', 1.0},
			},
			"ksa": {
				{'p', 1.0},
			},
			"aln": {
				{'i', 1.0},
			},
			"che": {
				{'t', 0.5},
				{'i', 1.0},
			},
			"siy": {
				{'i', 1.0},
			},
			"con": {
				{'a', 1.0},
			},
			"doe": {
				{'e', 1.0},
			},
			"hpe": {
				{'e', 1.0},
			},
			"Has": {
				{'t', 1.0},
			},
			"inz": {
				{'i', 1.0},
			},
			"gee": {
				{0, 0.5},
				{'s', 1.0},
			},
			"Dak": {
				{'o', 1.0},
			},
			"nsa": {
				{0, 1.0},
			},
			"ika": {
				{0, 1.0},
			},
			"Pez": {
				{'i', 1.0},
			},
			"amm": {
				{'e', 1.0},
			},
			"aga": {
				{'a', 1.0},
			},
			"Yan": {
				{'s', 0.5},
				{'i', 1.0},
			},
			"ego": {
				{'t', 1.0},
			},
			"hko": {
				{'o', 1.0},
			},
			"Shi": {
				{'y', 0.5},
				{'l', 1.0},
			},
			"tim": {
				{'i', 1.0},
			},
			"ala": {
				{'n', 1.0},
			},
			"een": {
				{'i', 1.0},
			},
			"Naa": {
				{'l', 1.0},
			},
			"Gaa": {
				{'g', 1.0},
			},
			"Api": {
				{'a', 1.0},
			},
			"uti": {
				{0, 1.0},
			},
			"nsk": {
				{'a', 1.0},
			},
			"ota": {
				{'h', 0.4},
				{0, 1.0},
			},
			"Cet": {
				{'a', 1.0},
			},
			"llo": {
				{0, 1.0},
			},
			"sta": {
				{'s', 1.0},
			},
			"Nak": {
				{'a', 1.0},
			},
			"hea": {
				{'s', 1.0},
			},
			"Wea": {
				{'y', 1.0},
			},
			"Seq": {
				{'u', 1.0},
			},
			"ilt": {
				{'h', 1.0},
			},
			"nko": {
				{'o', 1.0},
			},
			"hap": {
				{'p', 1.0},
			},
			"Kan": {
				{'u', 0.5},
				{'g', 1.0},
			},
			"Tok": {
				{'o', 0.5},
				{'e', 1.0},
			},
		}
	default:
		return nil
	}
}
//
// End of generated data
//
