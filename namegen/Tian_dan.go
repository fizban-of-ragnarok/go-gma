/*
########################################################################################
#  __                                                                                  #
# /__ _                                                                                #
# \_|(_)                                                                               #
#  _______  _______  _______             _______      __                               #
# (  ____ \(       )(  ___  ) Game      (  ____ \    /  \                              #
# | (    \/| () () || (   ) | Master's  | (    \/    \/) )                             #
# | |      | || || || (___) | Assistant | (____        | |                             #
# | | ____ | |(_)| ||  ___  | (Go Port) (_____ \       | |                             #
# | | \_  )| |   | || (   ) |                 ) )      | |                             #
# | (___) || )   ( || )   ( | Mapper    /\____) ) _  __) (_                            #
# (_______)|/     \||/     \| Client    \______/ (_) \____/                            #
#                                                                                      #
########################################################################################
*/
//
// Name Generator Cultural Module
// Automatically generated from JavaScript source file
//

package namegen

//
// TianDan describes the naming conventions for the Tian-dan
// culture. Its methods give further details, but generally speaking
// the main operation to perform on these types is to just call the
// Generate and GenerateWithSurnames methods to create new names which
// conform to their cultural patterns.
//
type TianDan struct {
	BaseCulture
}

//
// nameWords gives the number of parts to generate in Tian-dan given names.
//
func (c TianDan) nameWords(gender rune) int {
	switch gender {
	case 'F':
		return 2
	case 'M':
		return 2
	case 'S':
		return 1
	default:
		return 1
	}
}

//
// prefix gives the prefix/selector string for each Tian-dan gender, or an empty
// string if one is not defined.
//
func (c TianDan) prefix(gender rune) string {
	switch gender {
	case 'F':
		return "__"
	case 'M':
		return "__"
	case 'S':
		return "__"
	default:
		return ""
	}
}

//
// defaultMinMax returns the minimum and maximum size of Tian-dan names based on gender.
//
func (c TianDan) defaultMinMax(gender rune) (int, int) {
	switch gender {
	case 'F':
		return 2, 7
	case 'M':
		return 3, 7
	case 'S':
		return 2, 8
	default:
		return 1, 1
	}
}

//
// Genders returns the set of genders defined for the Tian-dan culture.
//
func (c TianDan) Genders() []rune {
	return []rune{'F', 'M', 'S'}
}

//
// HasSurnames returns true if the Tian-dan culture defines surnames.
//
func (c TianDan) HasSurnames() bool {
	return true
}

//
// Name returns the name of the culture, i.e., "Tian-dan".
//
func (c TianDan) Name() string {
	return "Tian-dan"
}

//
// maxCount returns the maximum number of times a rune can be added to
// a Tian-dan name based on gender.
//
func (c TianDan) maxCount(gender, char rune) int {
	switch gender {
	case 'F':
		switch char {
		case '-':
			return 1
		default:
			return 0
		}
	case 'M':
		switch char {
		case '-':
			return 1
		default:
			return 0
		}
	case 'S':
		switch char {
		default:
			return 0
		}
	default:
		return 0
	}
}

//
// HasGender returns true if the specified gender code is defined
// in the Tian-dan culture.
//
func (c TianDan) HasGender(gender rune) bool {
	switch gender {
	case 'F', 'M', 'S':
		return true
	default:
		return false
	}
}

//
// db returns the name data for the given gender in the Tian-dan culture.
//
func (c TianDan) db(gender rune) map[string][]nameFragment {
	switch gender {
	case 'F':
		return map[string][]nameFragment{
			"ie": {
				{'n', 0.25},
				{'u', 0.75},
				{'t', 0.875},
				{'m', 1.0},
			},
			"ng": {
				{0, 1.0},
			},
			"Ha": {
				{'n', 0.6667},
				{0, 1.0},
			},
			"in": {
				{'h', 1.0},
			},
			"Ta": {
				{'m', 0.3333},
				{'n', 0.6667},
				{'o', 1.0},
			},
			"Ly": {
				{0, 1.0},
			},
			"Xu": {
				{'a', 1.0},
			},
			"Au": {
				{0, 1.0},
			},
			"Tu": {
				{'y', 1.0},
			},
			"_U": {
				{'t', 1.0},
			},
			"Ho": {
				{'a', 0.6667},
				{'n', 1.0},
			},
			"gu": {
				{'y', 1.0},
			},
			"ye": {
				{'t', 0.6667},
				{'n', 1.0},
			},
			"ru": {
				{'c', 1.0},
			},
			"_S": {
				{'o', 1.0},
			},
			"Hi": {
				{'e', 1.0},
			},
			"ro": {
				{'n', 1.0},
			},
			"_K": {
				{'h', 0.25},
				{'i', 1.0},
			},
			"_L": {
				{'y', 0.1},
				{'u', 0.3},
				{'o', 0.4},
				{'i', 0.6},
				{'e', 0.8},
				{'a', 1.0},
			},
			"Ai": {
				{0, 1.0},
			},
			"Kh": {
				{'a', 1.0},
			},
			"Ca": {
				{'m', 0.5},
				{'n', 1.0},
			},
			"ic": {
				{'h', 1.0},
			},
			"Ki": {
				{'m', 0.3333},
				{'e', 1.0},
			},
			"ai": {
				{0, 1.0},
			},
			"ri": {
				{'n', 1.0},
			},
			"Hu": {
				{'y', 0.3333},
				{'o', 0.6667},
				{'e', 1.0},
			},
			"La": {
				{'n', 1.0},
			},
			"ao": {
				{0, 1.0},
			},
			"Ch": {
				{'a', 1.0},
			},
			"Di": {
				{'e', 1.0},
			},
			"So": {
				{0, 1.0},
			},
			"uc": {
				{0, 1.0},
			},
			"un": {
				{'g', 1.0},
			},
			"ha": {
				{'n', 0.6},
				{'o', 0.8},
				{'u', 1.0},
			},
			"_T": {
				{'a', 0.1765},
				{'h', 0.6471},
				{'r', 0.9412},
				{'u', 1.0},
			},
			"hu": {
				{0, 0.4},
				{'y', 0.6},
				{'o', 0.8},
				{'a', 1.0},
			},
			"ue": {
				{0, 1.0},
			},
			"_C": {
				{'h', 0.3333},
				{'a', 1.0},
			},
			"_A": {
				{'i', 0.25},
				{'m', 0.5},
				{'n', 0.75},
				{'u', 1.0},
			},
			"_D": {
				{'i', 1.0},
			},
			"Ph": {
				{'u', 0.5},
				{'a', 1.0},
			},
			"Gu": {
				{'n', 1.0},
			},
			"eu": {
				{0, 1.0},
			},
			"Le": {
				{'e', 0.5},
				{0, 1.0},
			},
			"Vi": {
				{'n', 1.0},
			},
			"on": {
				{0, 0.25},
				{'g', 1.0},
			},
			"im": {
				{0, 1.0},
			},
			"ee": {
				{0, 1.0},
			},
			"__": {
				{'D', 0.0286},
				{'X', 0.0429},
				{'N', 0.1143},
				{'T', 0.3571},
				{'G', 0.3714},
				{'B', 0.4143},
				{'Q', 0.4286},
				{'K', 0.4857},
				{'S', 0.5},
				{'V', 0.5143},
				{'M', 0.5714},
				{'C', 0.6143},
				{'U', 0.6286},
				{'A', 0.6857},
				{'L', 0.8286},
				{'H', 0.9714},
				{'P', 1.0},
			},
			"an": {
				{'h', 0.2308},
				{'g', 0.4615},
				{0, 1.0},
			},
			"en": {
				{0, 1.0},
			},
			"Lu": {
				{'i', 0.5},
				{'a', 1.0},
			},
			"uo": {
				{'n', 1.0},
			},
			"ho": {
				{0, 0.5},
				{'a', 1.0},
			},
			"_V": {
				{'i', 1.0},
			},
			"oc": {
				{0, 1.0},
			},
			"_N": {
				{'u', 0.4},
				{'g', 0.8},
				{'h', 1.0},
			},
			"uy": {
				{'e', 0.6},
				{0, 1.0},
			},
			"Li": {
				{'e', 1.0},
			},
			"Am": {
				{0, 1.0},
			},
			"Bi": {
				{'c', 1.0},
			},
			"nh": {
				{0, 1.0},
			},
			"ua": {
				{'m', 0.25},
				{0, 0.5},
				{'n', 1.0},
			},
			"Nu": {
				{'c', 0.5},
				{0, 1.0},
			},
			"_B": {
				{'i', 0.3333},
				{'u', 0.6667},
				{'e', 1.0},
			},
			"Lo": {
				{'a', 1.0},
			},
			"ch": {
				{0, 1.0},
			},
			"go": {
				{'c', 1.0},
			},
			"em": {
				{0, 1.0},
			},
			"_M": {
				{'y', 0.25},
				{'i', 0.75},
				{'a', 1.0},
			},
			"Ma": {
				{'i', 1.0},
			},
			"Be": {
				{0, 1.0},
			},
			"Bu": {
				{'a', 1.0},
			},
			"My": {
				{0, 1.0},
			},
			"he": {
				{0, 1.0},
			},
			"_X": {
				{'u', 1.0},
			},
			"Tr": {
				{'a', 0.4},
				{'u', 0.6},
				{'o', 0.8},
				{'i', 1.0},
			},
			"ui": {
				{0, 1.0},
			},
			"_Q": {
				{'u', 1.0},
			},
			"et": {
				{0, 1.0},
			},
			"Qu": {
				{'y', 1.0},
			},
			"ra": {
				{'n', 0.5},
				{'c', 1.0},
			},
			"Nh": {
				{'u', 1.0},
			},
			"oa": {
				{'n', 0.25},
				{0, 0.75},
				{'i', 1.0},
			},
			"au": {
				{0, 1.0},
			},
			"_H": {
				{'i', 0.1},
				{'u', 0.4},
				{'o', 0.7},
				{'a', 1.0},
			},
			"Th": {
				{'e', 0.125},
				{'a', 0.375},
				{'u', 0.75},
				{'o', 1.0},
			},
			"ac": {
				{0, 1.0},
			},
			"Ng": {
				{'u', 0.5},
				{'o', 1.0},
			},
			"_G": {
				{'u', 1.0},
			},
			"Ut": {
				{0, 1.0},
			},
			"_P": {
				{'h', 1.0},
			},
			"am": {
				{0, 1.0},
			},
			"Mi": {
				{'e', 0.5},
				{'n', 1.0},
			},
			"An": {
				{0, 1.0},
			},
		}
	case 'M':
		return map[string][]nameFragment{
			"_B": {
				{'a', 0.4286},
				{'i', 0.7143},
				{'u', 1.0},
			},
			"Bi": {
				{'c', 0.5},
				{'n', 1.0},
			},
			"nh": {
				{0, 1.0},
			},
			"ua": {
				{'n', 1.0},
			},
			"ah": {
				{'n', 1.0},
			},
			"go": {
				{'c', 1.0},
			},
			"Te": {
				{'o', 1.0},
			},
			"Ye": {
				{'n', 1.0},
			},
			"Lo": {
				{'a', 0.25},
				{'n', 0.5},
				{'c', 0.75},
				{'i', 1.0},
			},
			"ch": {
				{0, 0.8},
				{'u', 1.0},
			},
			"Li": {
				{'c', 0.3333},
				{'n', 0.6667},
				{'e', 1.0},
			},
			"uy": {
				{'e', 0.25},
				{0, 1.0},
			},
			"_Y": {
				{'u', 0.5},
				{'e', 1.0},
			},
			"Tr": {
				{'o', 0.5},
				{'a', 1.0},
			},
			"_X": {
				{'u', 1.0},
			},
			"eo": {
				{0, 1.0},
			},
			"ui": {
				{0, 1.0},
			},
			"uu": {
				{0, 1.0},
			},
			"oi": {
				{0, 1.0},
			},
			"_M": {
				{'a', 0.5},
				{'i', 0.8333},
				{'u', 1.0},
			},
			"Bu": {
				{'u', 0.5},
				{'i', 1.0},
			},
			"Ma": {
				{'n', 0.6667},
				{'i', 1.0},
			},
			"Qu": {
				{'o', 0.3333},
				{'a', 1.0},
			},
			"Va": {
				{'n', 1.0},
			},
			"hn": {
				{0, 1.0},
			},
			"et": {
				{0, 1.0},
			},
			"Vo": {
				{0, 1.0},
			},
			"oa": {
				{'n', 0.8},
				{0, 1.0},
			},
			"Nh": {
				{'a', 1.0},
			},
			"ra": {
				{'i', 1.0},
			},
			"at": {
				{0, 1.0},
			},
			"Mu": {
				{'o', 1.0},
			},
			"_Q": {
				{'u', 1.0},
			},
			"Ga": {
				{'p', 1.0},
			},
			"_G": {
				{'a', 0.3333},
				{'i', 1.0},
			},
			"da": {
				{'o', 1.0},
			},
			"Mi": {
				{'n', 1.0},
			},
			"An": {
				{0, 0.5},
				{'h', 1.0},
			},
			"am": {
				{0, 1.0},
			},
			"_P": {
				{'i', 0.2},
				{'h', 1.0},
			},
			"To": {
				{'n', 0.5},
				{'a', 1.0},
			},
			"_H": {
				{'y', 0.0625},
				{'u', 0.25},
				{'i', 0.5},
				{'o', 0.6875},
				{'a', 1.0},
			},
			"Ti": {
				{'c', 0.2},
				{'n', 0.6},
				{'e', 1.0},
			},
			"au": {
				{0, 1.0},
			},
			"Ni": {
				{'e', 1.0},
			},
			"ac": {
				{'h', 0.6667},
				{0, 1.0},
			},
			"Ng": {
				{'o', 0.5},
				{0, 1.0},
			},
			"Th": {
				{'u', 0.2727},
				{'i', 0.5455},
				{'o', 0.6364},
				{'a', 1.0},
			},
			"ad": {
				{'a', 0.5},
				{0, 1.0},
			},
			"ye": {
				{'n', 1.0},
			},
			"ay": {
				{0, 1.0},
			},
			"ap": {
				{0, 1.0},
			},
			"_S": {
				{'a', 0.3333},
				{'o', 1.0},
			},
			"in": {
				{'h', 0.6667},
				{0, 1.0},
			},
			"Ha": {
				{'u', 0.2},
				{0, 0.4},
				{'o', 0.6},
				{'i', 0.8},
				{'d', 1.0},
			},
			"Ta": {
				{'n', 0.6},
				{'m', 0.8},
				{'i', 1.0},
			},
			"ie": {
				{'t', 0.1429},
				{'p', 0.2143},
				{'u', 0.4286},
				{0, 0.5},
				{'n', 1.0},
			},
			"ng": {
				{0, 1.0},
			},
			"Tu": {
				{'n', 0.2},
				{'a', 0.4},
				{'y', 0.8},
				{0, 1.0},
			},
			"Ho": {
				{'a', 0.6667},
				{'n', 1.0},
			},
			"Xu": {
				{'a', 1.0},
			},
			"Hy": {
				{0, 1.0},
			},
			"Ki": {
				{'m', 1.0},
			},
			"Ky": {
				{0, 1.0},
			},
			"Ca": {
				{'o', 0.1667},
				{0, 0.3333},
				{'n', 0.6667},
				{'d', 0.8333},
				{'m', 1.0},
			},
			"ic": {
				{'h', 1.0},
			},
			"ai": {
				{0, 1.0},
			},
			"ia": {
				{0, 0.3333},
				{'n', 0.6667},
				{'p', 1.0},
			},
			"Hu": {
				{0, 0.3333},
				{'o', 0.6667},
				{'n', 1.0},
			},
			"Na": {
				{'m', 1.0},
			},
			"Do": {
				{'h', 0.3333},
				{'n', 0.6667},
				{'a', 1.0},
			},
			"hi": {
				{0, 0.1429},
				{'m', 0.2857},
				{'n', 0.4286},
				{'a', 0.5714},
				{'e', 1.0},
			},
			"Hi": {
				{0, 0.25},
				{'e', 1.0},
			},
			"ro": {
				{'n', 1.0},
			},
			"_L": {
				{'i', 0.3333},
				{'o', 0.7778},
				{'u', 1.0},
			},
			"_K": {
				{'i', 0.2},
				{'y', 0.4},
				{'h', 1.0},
			},
			"Kh": {
				{'a', 1.0},
			},
			"Cu": {
				{'o', 1.0},
			},
			"Ty": {
				{0, 1.0},
			},
			"_D": {
				{'a', 0.3684},
				{'i', 0.5789},
				{'o', 0.7368},
				{'u', 1.0},
			},
			"Gi": {
				{'a', 1.0},
			},
			"Ph": {
				{'a', 0.5},
				{'u', 1.0},
			},
			"_A": {
				{'n', 1.0},
			},
			"_C": {
				{'o', 0.0625},
				{'u', 0.125},
				{'a', 0.5},
				{'h', 1.0},
			},
			"eu": {
				{0, 1.0},
			},
			"Ch": {
				{'a', 0.375},
				{'o', 0.5},
				{'i', 1.0},
			},
			"Di": {
				{'e', 0.75},
				{'n', 1.0},
			},
			"ao": {
				{0, 1.0},
			},
			"Sa": {
				{'h', 1.0},
			},
			"_T": {
				{'r', 0.0625},
				{'a', 0.2188},
				{'y', 0.25},
				{'o', 0.3125},
				{'e', 0.3438},
				{'h', 0.6875},
				{'u', 0.8438},
				{'i', 1.0},
			},
			"hu": {
				{'n', 0.1667},
				{'c', 0.3333},
				{'y', 0.5},
				{0, 1.0},
			},
			"Pi": {
				{'n', 1.0},
			},
			"So": {
				{'n', 0.5},
				{0, 1.0},
			},
			"uc": {
				{0, 1.0},
			},
			"ep": {
				{0, 1.0},
			},
			"uk": {
				{0, 1.0},
			},
			"un": {
				{'g', 1.0},
			},
			"ha": {
				{'t', 0.1538},
				{'u', 0.2308},
				{'i', 0.3077},
				{'m', 0.5385},
				{'n', 0.8462},
				{'c', 1.0},
			},
			"Yu": {
				{0, 1.0},
			},
			"Lu": {
				{'u', 0.5},
				{'o', 1.0},
			},
			"ho": {
				{'n', 0.5},
				{0, 1.0},
			},
			"uo": {
				{'n', 0.6},
				{'c', 0.8},
				{'i', 1.0},
			},
			"Du": {
				{'n', 0.2},
				{'c', 0.4},
				{'k', 0.6},
				{0, 0.8},
				{'y', 1.0},
			},
			"_N": {
				{'g', 0.4},
				{'a', 0.6},
				{'h', 0.8},
				{'i', 1.0},
			},
			"_V": {
				{'a', 0.2},
				{'u', 0.6},
				{'o', 0.8},
				{'i', 1.0},
			},
			"oh": {
				{'m', 1.0},
			},
			"oc": {
				{0, 1.0},
			},
			"on": {
				{'g', 0.9091},
				{0, 1.0},
			},
			"hm": {
				{0, 1.0},
			},
			"im": {
				{0, 1.0},
			},
			"Vi": {
				{'e', 1.0},
			},
			"Da": {
				{'i', 0.1429},
				{'o', 0.2857},
				{0, 0.4286},
				{'m', 0.5714},
				{'n', 1.0},
			},
			"Ba": {
				{'t', 0.3333},
				{'c', 0.6667},
				{'y', 1.0},
			},
			"en": {
				{0, 1.0},
			},
			"an": {
				{'h', 0.2083},
				{0, 0.75},
				{'g', 1.0},
			},
			"__": {
				{'S', 0.0216},
				{'C', 0.1367},
				{'L', 0.2014},
				{'G', 0.223},
				{'K', 0.259},
				{'N', 0.295},
				{'Q', 0.3165},
				{'P', 0.3525},
				{'H', 0.4676},
				{'V', 0.5036},
				{'A', 0.518},
				{'B', 0.5683},
				{'Y', 0.5827},
				{'M', 0.6259},
				{'D', 0.7626},
				{'T', 0.9928},
				{'X', 1.0},
			},
			"Co": {
				{'n', 1.0},
			},
			"Vu": {
				{'i', 0.5},
				{0, 1.0},
			},
		}
	case 'S':
		return map[string][]nameFragment{
			"Cu": {
				{0, 1.0},
			},
			"Kh": {
				{'u', 0.5},
				{'o', 1.0},
			},
			"_L": {
				{'e', 0.1111},
				{'u', 0.4444},
				{'a', 0.6667},
				{'o', 0.8889},
				{'y', 1.0},
			},
			"_K": {
				{'i', 0.4},
				{'h', 0.8},
				{'w', 1.0},
			},
			"La": {
				{'n', 0.5},
				{'m', 1.0},
			},
			"ri": {
				{'n', 0.5},
				{'e', 1.0},
			},
			"ia": {
				{0, 0.5},
				{'p', 1.0},
			},
			"Hu": {
				{'n', 0.5},
				{'y', 1.0},
			},
			"ai": {
				{0, 1.0},
			},
			"Do": {
				{'a', 0.5},
				{0, 1.0},
			},
			"hi": {
				{'e', 0.5},
				{'a', 1.0},
			},
			"Ki": {
				{'e', 0.5},
				{'m', 1.0},
			},
			"Ca": {
				{'h', 0.5},
				{'o', 1.0},
			},
			"Ho": {
				{'a', 0.6667},
				{0, 1.0},
			},
			"_U": {
				{'t', 1.0},
			},
			"Tu": {
				{'a', 0.5},
				{'n', 1.0},
			},
			"Xu": {
				{'a', 1.0},
			},
			"Ly": {
				{0, 1.0},
			},
			"Ta": {
				{'n', 1.0},
			},
			"in": {
				{'h', 1.0},
			},
			"Ha": {
				{0, 1.0},
			},
			"yn": {
				{'h', 1.0},
			},
			"ng": {
				{0, 1.0},
			},
			"ie": {
				{'u', 0.5},
				{'n', 1.0},
			},
			"_S": {
				{'a', 0.6667},
				{'o', 1.0},
			},
			"ru": {
				{0, 0.3333},
				{'o', 0.6667},
				{'n', 1.0},
			},
			"ap": {
				{0, 1.0},
			},
			"ye": {
				{'n', 1.0},
			},
			"gu": {
				{'y', 1.0},
			},
			"en": {
				{'g', 0.25},
				{0, 1.0},
			},
			"an": {
				{0, 0.5333},
				{'h', 0.6667},
				{'g', 1.0},
			},
			"Ba": {
				{0, 1.0},
			},
			"Vu": {
				{0, 1.0},
			},
			"__": {
				{'T', 0.2024},
				{'D', 0.3095},
				{'O', 0.3333},
				{'B', 0.369},
				{'V', 0.3929},
				{'W', 0.4048},
				{'P', 0.4881},
				{'L', 0.5952},
				{'U', 0.6071},
				{'C', 0.6786},
				{'M', 0.7262},
				{'N', 0.7976},
				{'K', 0.8571},
				{'H', 0.9405},
				{'G', 0.9524},
				{'S', 0.9881},
				{'X', 1.0},
			},
			"im": {
				{0, 1.0},
			},
			"gh": {
				{'i', 1.0},
			},
			"on": {
				{'g', 0.8571},
				{0, 1.0},
			},
			"Da": {
				{'i', 0.3333},
				{'n', 1.0},
			},
			"_N": {
				{'g', 0.8333},
				{'i', 1.0},
			},
			"Du": {
				{'e', 0.5},
				{'o', 1.0},
			},
			"oc": {
				{0, 1.0},
			},
			"_V": {
				{'u', 0.5},
				{'o', 1.0},
			},
			"ho": {
				{'n', 1.0},
			},
			"uo": {
				{'n', 1.0},
			},
			"Lu": {
				{'o', 0.3333},
				{'u', 0.6667},
				{0, 1.0},
			},
			"ok": {
				{0, 1.0},
			},
			"_T": {
				{'u', 0.1176},
				{'a', 0.1765},
				{'r', 0.5882},
				{'h', 0.9412},
				{'o', 1.0},
			},
			"hu": {
				{'a', 0.125},
				{0, 0.375},
				{'c', 0.75},
				{'n', 0.875},
				{'y', 1.0},
			},
			"Sa": {
				{'m', 0.5},
				{0, 1.0},
			},
			"un": {
				{'g', 1.0},
			},
			"ha": {
				{'i', 0.1667},
				{'n', 0.5},
				{'m', 0.6667},
				{'c', 0.8333},
				{'u', 1.0},
			},
			"So": {
				{'n', 1.0},
			},
			"uc": {
				{0, 1.0},
			},
			"Di": {
				{'n', 1.0},
			},
			"Ch": {
				{'i', 0.3333},
				{'e', 0.6667},
				{'a', 1.0},
			},
			"ao": {
				{0, 1.0},
			},
			"Le": {
				{0, 1.0},
			},
			"eu": {
				{0, 1.0},
			},
			"On": {
				{'g', 1.0},
			},
			"Kw": {
				{'a', 1.0},
			},
			"Ph": {
				{'o', 0.1429},
				{'a', 0.5714},
				{'u', 1.0},
			},
			"Gi": {
				{'a', 1.0},
			},
			"_D": {
				{'o', 0.2222},
				{'i', 0.3333},
				{'e', 0.4444},
				{'a', 0.7778},
				{'u', 1.0},
			},
			"_C": {
				{'u', 0.1667},
				{'h', 0.6667},
				{'a', 1.0},
			},
			"ue": {
				{0, 1.0},
			},
			"Bu": {
				{'i', 1.0},
			},
			"Ma": {
				{'c', 0.3333},
				{0, 0.6667},
				{'i', 1.0},
			},
			"_M": {
				{'a', 0.75},
				{'i', 1.0},
			},
			"uu": {
				{0, 1.0},
			},
			"ui": {
				{0, 1.0},
			},
			"_X": {
				{'u', 1.0},
			},
			"il": {
				{'h', 1.0},
			},
			"Tr": {
				{'i', 0.2857},
				{'u', 0.7143},
				{'a', 1.0},
			},
			"he": {
				{0, 1.0},
			},
			"_W": {
				{'a', 1.0},
			},
			"uy": {
				{'n', 0.3333},
				{0, 0.6667},
				{'e', 1.0},
			},
			"go": {
				{0, 0.5},
				{'c', 1.0},
			},
			"ch": {
				{0, 1.0},
			},
			"Lo": {
				{'k', 0.5},
				{'n', 1.0},
			},
			"_O": {
				{'n', 0.5},
				{0, 1.0},
			},
			"_B": {
				{'u', 0.3333},
				{'a', 0.6667},
				{'i', 1.0},
			},
			"ua": {
				{'n', 1.0},
			},
			"Wa": {
				{'n', 1.0},
			},
			"ah": {
				{'n', 1.0},
			},
			"nh": {
				{0, 1.0},
			},
			"Bi": {
				{'l', 1.0},
			},
			"Ng": {
				{'o', 0.4},
				{'u', 0.6},
				{0, 0.8},
				{'h', 1.0},
			},
			"ac": {
				{'h', 0.5},
				{0, 1.0},
			},
			"De": {
				{0, 1.0},
			},
			"Th": {
				{'a', 0.3333},
				{'u', 1.0},
			},
			"_H": {
				{'e', 0.1429},
				{'a', 0.2857},
				{'u', 0.5714},
				{'o', 1.0},
			},
			"To": {
				{'a', 1.0},
			},
			"Ni": {
				{'e', 1.0},
			},
			"au": {
				{0, 1.0},
			},
			"Mi": {
				{'n', 1.0},
			},
			"am": {
				{0, 1.0},
			},
			"_P": {
				{'h', 1.0},
			},
			"Ut": {
				{0, 1.0},
			},
			"He": {
				{'n', 1.0},
			},
			"_G": {
				{'i', 1.0},
			},
			"lh": {
				{0, 1.0},
			},
			"wa": {
				{'n', 1.0},
			},
			"oa": {
				{'n', 0.75},
				{'i', 1.0},
			},
			"Vo": {
				{0, 1.0},
			},
			"ra": {
				{0, 0.5},
				{'n', 1.0},
			},
			"hn": {
				{0, 1.0},
			},
		}
	default:
		return nil
	}
}
//
// End of generated data
//
