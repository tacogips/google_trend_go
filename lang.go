package google_trend_go

type Lang int

func (l Lang) String() string {
	return langCodeMap[l]
}

const (
	LangAB = iota
	LangAA
	LangAF
	LangAK
	LangSQ
	LangAM
	LangAR
	LangAN
	LangHY
	LangAS
	LangAV
	LangAE
	LangAY
	LangAZ
	LangBM
	LangBA
	LangEU
	LangBE
	LangBN
	LangBH
	LangBI
	LangBS
	LangBR
	LangBG
	LangMY
	LangCA
	LangCH
	LangCE
	LangNY
	LangZH
	LangZH_HANS
	LangZH_HANT
	LangCV
	LangKW
	LangCO
	LangCR
	LangHR
	LangCS
	LangDA
	LangDV
	LangNL
	LangDZ
	LangEN
	LangEO
	LangET
	LangEE
	LangFO
	LangFJ
	LangFI
	LangFR
	LangFF
	LangGL
	LangGD
	LangGV
	LangKA
	LangDE
	LangEL
	LangKL
	LangGN
	LangGU
	LangHT
	LangHA
	LangHE
	LangHZ
	LangHI
	LangHO
	LangHU
	LangIS
	LangIO
	LangIG
	LangIN
	LangIA
	LangIE
	LangIU
	LangIK
	LangGA
	LangIT
	LangJA
	LangJV
	LangKN
	LangKR
	LangKS
	LangKK
	LangKM
	LangKI
	LangRW
	LangRN
	LangKY
	LangKV
	LangKG
	LangKO
	LangKU
	LangKJ
	LangLO
	LangLA
	LangLV
	LangLI
	LangLN
	LangLT
	LangLU
	LangLG
	LangLB
	LangMK
	LangMG
	LangMS
	LangML
	LangMT
	LangMI
	LangMR
	LangMH
	LangMO
	LangMN
	LangNA
	LangNV
	LangNG
	LangND
	LangNE
	LangNO
	LangNB
	LangNN
	LangII
	LangOC
	LangOJ
	LangCU
	LangOR
	LangOM
	LangOS
	LangPI
	LangPS
	LangFA
	LangPL
	LangPT
	LangPA
	LangQU
	LangRM
	LangRO
	LangRU
	LangSE
	LangSM
	LangSG
	LangSA
	LangSR
	LangSH
	LangST
	LangTN
	LangSN
	LangSD
	LangSI
	LangSS
	LangSK
	LangSL
	LangSO
	LangNR
	LangES
	LangSU
	LangSW
	LangSV
	LangTL
	LangTY
	LangTG
	LangTA
	LangTT
	LangTE
	LangTH
	LangBO
	LangTI
	LangTO
	LangTS
	LangTR
	LangTK
	LangTW
	LangUG
	LangUK
	LangUR
	LangUZ
	LangVE
	LangVI
	LangVO
	LangWA
	LangCY
	LangWO
	LangFY
	LangXH
	LangYI
	LangYO
	LangZA
	LangZU
)

type langCodeMapType map[Lang]string

var langCodeMap langCodeMapType = map[Lang]string{
	LangAB:      "ab",
	LangAA:      "aa",
	LangAF:      "af",
	LangAK:      "ak",
	LangSQ:      "sq",
	LangAM:      "am",
	LangAR:      "ar",
	LangAN:      "an",
	LangHY:      "hy",
	LangAS:      "as",
	LangAV:      "av",
	LangAE:      "ae",
	LangAY:      "ay",
	LangAZ:      "az",
	LangBM:      "bm",
	LangBA:      "ba",
	LangEU:      "eu",
	LangBE:      "be",
	LangBN:      "bn",
	LangBH:      "bh",
	LangBI:      "bi",
	LangBS:      "bs",
	LangBR:      "br",
	LangBG:      "bg",
	LangMY:      "my",
	LangCA:      "ca",
	LangCH:      "ch",
	LangCE:      "ce",
	LangNY:      "ny",
	LangZH:      "zh",
	LangZH_HANS: "zh-Hans",
	LangZH_HANT: "zh-Hant",
	LangCV:      "cv",
	LangKW:      "kw",
	LangCO:      "co",
	LangCR:      "cr",
	LangHR:      "hr",
	LangCS:      "cs",
	LangDA:      "da",
	LangDV:      "dv",
	LangNL:      "nl",
	LangDZ:      "dz",
	LangEN:      "en",
	LangEO:      "eo",
	LangET:      "et",
	LangEE:      "ee",
	LangFO:      "fo",
	LangFJ:      "fj",
	LangFI:      "fi",
	LangFR:      "fr",
	LangFF:      "ff",
	LangGL:      "gl",
	LangGD:      "gd",
	LangGV:      "gv",
	LangKA:      "ka",
	LangDE:      "de",
	LangEL:      "el",
	LangKL:      "kl",
	LangGN:      "gn",
	LangGU:      "gu",
	LangHT:      "ht",
	LangHA:      "ha",
	LangHE:      "he",
	LangHZ:      "hz",
	LangHI:      "hi",
	LangHO:      "ho",
	LangHU:      "hu",
	LangIS:      "is",
	LangIO:      "io",
	LangIG:      "ig",
	LangIN:      "in",
	LangIA:      "ia",
	LangIE:      "ie",
	LangIU:      "iu",
	LangIK:      "ik",
	LangGA:      "ga",
	LangIT:      "it",
	LangJA:      "ja",
	LangJV:      "jv",
	LangKN:      "kn",
	LangKR:      "kr",
	LangKS:      "ks",
	LangKK:      "kk",
	LangKM:      "km",
	LangKI:      "ki",
	LangRW:      "rw",
	LangRN:      "rn",
	LangKY:      "ky",
	LangKV:      "kv",
	LangKG:      "kg",
	LangKO:      "ko",
	LangKU:      "ku",
	LangKJ:      "kj",
	LangLO:      "lo",
	LangLA:      "la",
	LangLV:      "lv",
	LangLI:      "li",
	LangLN:      "ln",
	LangLT:      "lt",
	LangLU:      "lu",
	LangLG:      "lg",
	LangLB:      "lb",
	LangMK:      "mk",
	LangMG:      "mg",
	LangMS:      "ms",
	LangML:      "ml",
	LangMT:      "mt",
	LangMI:      "mi",
	LangMR:      "mr",
	LangMH:      "mh",
	LangMO:      "mo",
	LangMN:      "mn",
	LangNA:      "na",
	LangNV:      "nv",
	LangNG:      "ng",
	LangND:      "nd",
	LangNE:      "ne",
	LangNO:      "no",
	LangNB:      "nb",
	LangNN:      "nn",
	LangII:      "ii",
	LangOC:      "oc",
	LangOJ:      "oj",
	LangCU:      "cu",
	LangOR:      "or",
	LangOM:      "om",
	LangOS:      "os",
	LangPI:      "pi",
	LangPS:      "ps",
	LangFA:      "fa",
	LangPL:      "pl",
	LangPT:      "pt",
	LangPA:      "pa",
	LangQU:      "qu",
	LangRM:      "rm",
	LangRO:      "ro",
	LangRU:      "ru",
	LangSE:      "se",
	LangSM:      "sm",
	LangSG:      "sg",
	LangSA:      "sa",
	LangSR:      "sr",
	LangSH:      "sh",
	LangST:      "st",
	LangTN:      "tn",
	LangSN:      "sn",
	LangSD:      "sd",
	LangSI:      "si",
	LangSS:      "ss",
	LangSK:      "sk",
	LangSL:      "sl",
	LangSO:      "so",
	LangNR:      "nr",
	LangES:      "es",
	LangSU:      "su",
	LangSW:      "sw",
	LangSV:      "sv",
	LangTL:      "tl",
	LangTY:      "ty",
	LangTG:      "tg",
	LangTA:      "ta",
	LangTT:      "tt",
	LangTE:      "te",
	LangTH:      "th",
	LangBO:      "bo",
	LangTI:      "ti",
	LangTO:      "to",
	LangTS:      "ts",
	LangTR:      "tr",
	LangTK:      "tk",
	LangTW:      "tw",
	LangUG:      "ug",
	LangUK:      "uk",
	LangUR:      "ur",
	LangUZ:      "uz",
	LangVE:      "ve",
	LangVI:      "vi",
	LangVO:      "vo",
	LangWA:      "wa",
	LangCY:      "cy",
	LangWO:      "wo",
	LangFY:      "fy",
	LangXH:      "xh",
	LangYI:      "yi",
	LangYO:      "yo",
	LangZA:      "za",
	LangZU:      "zu",
}
