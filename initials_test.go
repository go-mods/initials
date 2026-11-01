package initials

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// testCase describes a call to GetInitials and its expected result
type testCase struct {
	title   string
	input   string
	options []Option
	want    string
}

// runTestCases runs each test case in its own subtest
func runTestCases(t *testing.T, tests []testCase) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			assert.Equal(t, tt.want, GetInitials(tt.input, tt.options...))
		})
	}
}

func TestGetInitials_Simple(t *testing.T) {
	runTestCases(t, []testCase{
		{title: "empty name", input: "", want: ""},
		{title: "first name and last name", input: "John Doe", want: "JD"},
	})
}

func TestGetInitials_Separator(t *testing.T) {
	runTestCases(t, []testCase{
		{title: "dot", input: "John Doe", options: []Option{WithSeparator(".")}, want: "J.D"},
		{title: "hyphen", input: "John Doe", options: []Option{WithSeparator("-")}, want: "J-D"},
		{title: "length above the number of characters", input: "John Doe", options: []Option{WithLength(10), WithSeparator(".")}, want: "J.DOE"},
		{title: "length reached before the last word", input: "John David Dugga", options: []Option{WithSeparator(".")}, want: "J.D"},
		{title: "single word", input: "John", options: []Option{WithLength(10), WithSeparator(".")}, want: "JOHN"},
	})
}

func TestGetInitials_CaseOptions(t *testing.T) {
	runTestCases(t, []testCase{
		{title: "sensitive title case", input: "John Doe", options: []Option{WithSensitive()}, want: "JD"},
		{title: "sensitive mixed case", input: "John doe", options: []Option{WithSensitive()}, want: "Jd"},
		{title: "sensitive lowercase", input: "john doe", options: []Option{WithSensitive()}, want: "jd"},
		{title: "lowercase from title case", input: "John Doe", options: []Option{WithLowercase()}, want: "jd"},
		{title: "lowercase from mixed case", input: "John doe", options: []Option{WithLowercase()}, want: "jd"},
		{title: "lowercase from lowercase", input: "john doe", options: []Option{WithLowercase()}, want: "jd"},
		{title: "uppercase from title case", input: "John Doe", options: []Option{WithUppercase()}, want: "JD"},
		{title: "uppercase from mixed case", input: "John doe", options: []Option{WithUppercase()}, want: "JD"},
		{title: "uppercase from lowercase", input: "john doe", options: []Option{WithUppercase()}, want: "JD"},
	})
}

func TestGetInitials_LengthOptions(t *testing.T) {
	runTestCases(t, []testCase{
		{title: "two words length 2", input: "John Doe", options: []Option{WithLength(2)}, want: "JD"},
		{title: "two words length 2 camel case", input: "John Doe", options: []Option{WithLength(2), WithCamelCase()}, want: "JD"},
		{title: "two words length 3", input: "John Doe", options: []Option{WithLength(3)}, want: "JDO"},
		{title: "two words length 3 camel case", input: "John Doe", options: []Option{WithLength(3), WithCamelCase()}, want: "JDo"},
		{title: "two words length 4 camel case", input: "John Doe", options: []Option{WithLength(4), WithCamelCase()}, want: "JDoe"},
		{title: "one word length 2", input: "John", options: []Option{WithLength(2)}, want: "JO"},
		{title: "one word length 3", input: "John", options: []Option{WithLength(3)}, want: "JOH"},
		{title: "one word length 2 camel case", input: "John", options: []Option{WithLength(2), WithCamelCase()}, want: "Jo"},
		{title: "one word length 3 camel case", input: "John", options: []Option{WithLength(3), WithCamelCase()}, want: "Joh"},
		{title: "length above the number of characters", input: "John Doe", options: []Option{WithLength(10)}, want: "JDOE"},
		{title: "length above the number of characters camel case", input: "John Doe", options: []Option{WithLength(10), WithCamelCase()}, want: "JDoe"},
		{title: "zero length", input: "John Doe", options: []Option{WithLength(0)}, want: ""},
		{title: "negative length", input: "John Doe", options: []Option{WithLength(-1)}, want: ""},
		{title: "negative length one word", input: "John", options: []Option{WithLength(-1)}, want: ""},
	})
}

func TestGetInitials_Dash(t *testing.T) {
	runTestCases(t, []testCase{
		{title: "length 2", input: "Mary-Kate Olsen", options: []Option{WithLength(2)}, want: "MK"},
		{title: "length 3", input: "Mary-Kate Olsen", options: []Option{WithLength(3)}, want: "MKO"},
		{title: "length 4", input: "Mary-Kate Olsen", options: []Option{WithLength(4)}, want: "MKOL"},
		{title: "length 4 camel case", input: "Mary-Kate Olsen", options: []Option{WithLength(4), WithCamelCase()}, want: "MKOl"},
	})
}

func TestGetInitials_SpecialCharacters(t *testing.T) {
	runTestCases(t, []testCase{
		{title: "inside a word length 2", input: "M@ry Olsen", options: []Option{WithLength(2)}, want: "MO"},
		{title: "inside a word length 3", input: "M@ry Olsen", options: []Option{WithLength(3)}, want: "MOL"},
		{title: "inside a word length 4", input: "M@ry Olsen", options: []Option{WithLength(4)}, want: "MOLS"},
		{title: "inside a word length 4 camel case", input: "M@ry Olsen", options: []Option{WithLength(4), WithCamelCase()}, want: "MOls"},
		{title: "second letter of a single word", input: "J@hn", want: "JH"},
		{title: "at the start of words", input: "(John) _Doe", want: "JD"},
		{title: "word made of special characters", input: "John # Doe", want: "JD"},
		{title: "only special characters", input: "#!?", want: ""},
	})
}

func TestGetInitials_WordLengthOption(t *testing.T) {
	runTestCases(t, []testCase{
		{title: "two words", input: "John Doe", options: []Option{WithWordLength()}, want: "JD"},
		{title: "three words", input: "Mary-Kate Olsen", options: []Option{WithWordLength()}, want: "MKO"},
		{title: "three words camel case", input: "Mary-Kate Olsen", options: []Option{WithWordLength(), WithCamelCase()}, want: "MKO"},
	})
}

func TestGetInitials_NonASCII(t *testing.T) {
	runTestCases(t, []testCase{
		{title: "emoji", input: "😅", want: "😅"},
		{title: "emoji length 2", input: "😅", options: []Option{WithLength(2)}, want: "😅"},
		{title: "emoji camel case", input: "😅", options: []Option{WithLength(2), WithCamelCase()}, want: "😅"},
		{title: "emoji before a word", input: "😅 Doe", want: "😅D"},
		{title: "word before an emoji", input: "Doe 😅", want: "D😅"},
		{title: "ideographs", input: "张 伟", want: "张伟"},
		{title: "ideographs in a single word", input: "张伟明", want: "张伟"},
		{title: "ideographs with a separator", input: "张 伟", options: []Option{WithSeparator(".")}, want: "张.伟"},
	})
}

func TestGetInitials_FriendlyCharacters(t *testing.T) {
	runTestCases(t, []testCase{
		{title: "accents", input: "Élodie Müller", want: "EM"},
		{title: "letter becoming two letters", input: "Шарль", want: "SH"},
		{title: "cyrillic er", input: "Роман Петров", want: "RP"},
		{title: "cyrillic en", input: "Нина Doe", want: "ND"},
		{title: "cyrillic ghe with upturn", input: "Ґалина Doe", want: "GD"},
		{title: "non-breaking space", input: "John Doe", want: "JD"},
		{title: "character made of several runes", input: "ို", options: []Option{WithLength(1)}, want: "O"},
		{title: "cyrillic capital i", input: "Иван Петров", want: "IP"},
		{title: "cyrillic capital ghe", input: "Георгий Doe", want: "GD"},
		{title: "cyrillic sign removed", input: "Ьа", want: "A"},
		{title: "only cyrillic signs", input: "ьъ", want: ""},
		{title: "extended cyrillic", input: "Җамал Ҽ", want: "ZC"},
		{title: "greek capital gamma", input: "Γιώργος Παπαδόπουλος", want: "GP"},
		{title: "greek theta", input: "Θεόδωρος", want: "TH"},
		{title: "greek eta", input: "Ηλίας", want: "EL"},
		{title: "greek chi", input: "Χρήστος", want: "CH"},
		{title: "polytonic greek", input: "Ῥόδος", want: "RO"},
		{title: "arabic ain", input: "عمر", want: "AM"},
		{title: "arabic qaf", input: "قاسم", want: "QA"},
		{title: "arabic waw", input: "وليد", want: "WL"},
		{title: "georgian", input: "გიორგი", want: "GI"},
		{title: "georgian phar", input: "ფირუზა", want: "PI"},
		{title: "burmese", input: "ဖြိုး", want: "PH"},
		{title: "devanagari", input: "राम", want: "RA"},
		{title: "devanagari conjunct", input: "कृष्ण", want: "KR"},
	})
}

// Each letter of these ranges must have an equivalent, identical for both cases.
func TestFriendlyCharactersMap_Alphabets(t *testing.T) {
	equivalents := make(map[rune]string)
	for equivalent, characters := range friendlyCharactersMap {
		for _, character := range characters {
			if utf8.RuneCountInString(character) == 1 {
				equivalents[[]rune(character)[0]] = equivalent
			}
		}
	}

	alphabets := []struct {
		title    string
		from, to rune
	}{
		{title: "cyrillic", from: 0x0400, to: 0x04FF},
		{title: "greek", from: 0x0370, to: 0x03E1},
		{title: "greek symbols", from: 0x03F0, to: 0x03FF},
		{title: "polytonic greek", from: 0x1F00, to: 0x1FFF},
		{title: "arabic", from: 0x0621, to: 0x064A},
		{title: "burmese", from: 0x1000, to: 0x1021},
		{title: "georgian", from: 0x10D0, to: 0x10FA},
		{title: "georgian capitals", from: 0x1C90, to: 0x1CBA},
		{title: "devanagari", from: 0x0905, to: 0x0939},
	}

	for _, alphabet := range alphabets {
		t.Run(alphabet.title, func(t *testing.T) {
			for r := alphabet.from; r <= alphabet.to; r++ {
				if !unicode.IsLetter(r) || unicode.Is(unicode.Lm, r) {
					continue
				}

				equivalent, found := equivalents[r]
				if !assert.True(t, found, "no equivalent for %c (%U)", r, r) {
					continue
				}

				if upper, found := equivalents[unicode.ToUpper(r)]; found {
					assert.True(t, strings.EqualFold(equivalent, upper), "%c (%U) gives %q, %c gives %q", r, r, equivalent, unicode.ToUpper(r), upper)
				}
			}
		})
	}
}

// A character listed under several equivalents would make the result depend on the order of
// the table: each one must have a single equivalent.
func TestFriendlyCharactersMap_NoAmbiguousCharacter(t *testing.T) {
	equivalents := make(map[string]string)
	for equivalent, characters := range friendlyCharactersMap {
		for _, character := range characters {
			if other, found := equivalents[character]; found {
				assert.Equal(t, other, equivalent, "character %q (%U)", character, []rune(character))
			}
			equivalents[character] = equivalent
		}
	}
	require.NotEmpty(t, equivalents)
}

func FuzzGetInitials(f *testing.F) {
	f.Add("John Doe", 2)
	f.Add("Mary-Kate Olsen", 4)
	f.Add("😅 Doe", 2)
	f.Add("张 伟", 2)
	f.Add("Шарль ို", 3)
	f.Add("(John) _Doe", -1)
	f.Add("\xff\xfe", 1)

	f.Fuzz(func(t *testing.T, name string, length int) {
		got := GetInitials(name, WithLength(length))

		assert.LessOrEqual(t, utf8.RuneCountInString(got), max(length, 0))
		assert.Equal(t, got, GetInitials(name, WithLength(length)))
	})
}

func ExampleGetInitials() {
	fmt.Println(GetInitials("John Doe"))
	fmt.Println(GetInitials("John Doe", WithLength(3), WithCamelCase()))
	fmt.Println(GetInitials("John David Dugga", WithWordLength()))
	// Output:
	// JD
	// JDo
	// JDD
}
