package util

import (
	"strings"
	"unicode"
)

var kanaCase = unicode.SpecialCase{
	// ァ-ヴ
	unicode.CaseRange{
		Lo: 0x3040,
		Hi: 0x3094,
		Delta: [unicode.MaxCase]rune{
			0x30a1 - 0x3041,
			0,
			0,
		},
	},
	// ぁ-ゔ
	unicode.CaseRange{
		Lo: 0x30a0,
		Hi: 0x30f4,
		Delta: [unicode.MaxCase]rune{
			0,
			0x3041 - 0x30a1,
			0,
		},
	},
}

// HiraganaToKatakana はひらがなをカタカナに変換する
func HiraganaToKatakana(s string) string {
	return strings.ToUpperSpecial(kanaCase, s)
}

var alphanumConv = unicode.SpecialCase{
	// numbers
	unicode.CaseRange{
		Lo: 0xff10, // '０'
		Hi: 0xff19, // '９'
		Delta: [unicode.MaxCase]rune{
			0,
			0x0030 - 0xff10, // '0' - '０'
			0,
		},
	},
	// uppercase letters
	unicode.CaseRange{
		Lo: 0xff21, // 'Ａ'
		Hi: 0xff3a, // 'Ｚ'
		Delta: [unicode.MaxCase]rune{
			0,
			0x0041 - 0xff21, // 'A' - 'Ａ'
			0,
		},
	},
	// lowercase letters
	unicode.CaseRange{
		Lo: 0xff41, // 'ａ'
		Hi: 0xff5a, // 'ｚ'
		Delta: [unicode.MaxCase]rune{
			0,
			0x0061 - 0xff41, // 'a' - 'ａ'
			0,
		},
	},
}

// 英数字を半角小文字に変換
func AlnumToHalfLowwer(s string) string {
	return strings.ToLowerSpecial(alphanumConv, s)
}
