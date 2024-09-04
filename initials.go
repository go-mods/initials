package initials

import (
	"math"
	"regexp"
	"strings"
	"unicode"
)

// generatorOptions is a struct to hold the options for the initials generator
type generatorOptions struct {
	// Name is the name to get the initials from
	Name string

	// Separator is the separator to use between the initials
	Separator string

	// Sensitive is a boolean to indicate if the initials should be case-sensitive
	Sensitive bool

	// Lowercase is a boolean to indicate if the initials should be lowercase
	Lowercase bool

	// Uppercase is a boolean to indicate if the initials should be uppercase
	Uppercase bool

	// CamelCase is a boolean to indicate if the initials should be camel case
	CamelCase bool

	// SpecialChars is a boolean to indicate if the initials should contain special characters
	SpecialChars bool

	// Length is the maximum length of the initials
	Length int

	// WordLength set the length to the number of words
	WordLength bool

	// remainingLength is the remaining length of the initials
	remainingLength int
}

// Option is a function to set the generatorOptions for the initials generator
type Option func(*generatorOptions)

// GetInitials returns the initials of a name
func GetInitials(name string, options ...Option) string {

	// Set the default generator options
	opts := generatorOptions{
		Name:            name,
		Separator:       "",
		Sensitive:       false,
		Lowercase:       false,
		Uppercase:       true,
		CamelCase:       false,
		SpecialChars:    false,
		Length:          2,
		WordLength:      false,
		remainingLength: 2,
	}

	// Apply the provided options to the generator options
	for _, option := range options {
		option(&opts)
	}

	// Replace special characters if specified
	if !opts.SpecialChars {
		opts.Name = replaceSpecialCharacters(opts.Name)
	}

	// Replace with friendly characters
	opts.Name = replaceFriendlyCharacters(opts.Name)

	// Split the name into words by spaces and hyphens
	words := strings.FieldsFunc(opts.Name, func(r rune) bool {
		return r == ' ' || r == '-'
	})

	// If there are no words, return an empty string
	if len(words) == 0 {
		return ""
	}

	// If WordLength is set, adjust the length to the number of words
	if opts.WordLength {
		opts.Length = len(words)
		opts.remainingLength = len(words)
	}

	// Loop through the words and get the initials
	initials := ""
	for i, word := range words {
		// If the initials should be limited to a certain length, return the initials
		if opts.remainingLength == 0 {
			return initials
		}

		// Determine the length of the initial to get from the word
		var length = 1

		// The length is equal to 1 for the first word
		if i == 0 && len(words) > 1 {
			length = 1
		} else
		// The length is equal to 1 if the remaining length is less than the number of words left
		if opts.remainingLength < len(words)-i {
			length = 1
		} else
		// The length is equal to the remaining length divided by the number of words left
		if opts.remainingLength > len(words)-i {
			length = int(math.Floor(float64(opts.remainingLength) / float64(len(words)-i)))
		}

		// The length is equal to the remaining length for the last word
		if i == len(words)-1 {
			length = opts.remainingLength
		}

		// Get the first character of the word
		initial := getInitial(word, length, &opts)

		// Reduce the remaining length
		opts.remainingLength -= len(initial)

		// Add the initial
		initials += initial

		// Add the separator
		if len(initials) > 0 && opts.remainingLength > 0 {
			initials += opts.Separator
		}
	}

	// Return the initials
	return initials
}

// WithSeparator sets the separator to use between the initials
func WithSeparator(separator string) Option {
	return func(o *generatorOptions) {
		o.Separator = separator
	}
}

// WithSensitive sets the initials to be case-sensitive
func WithSensitive() Option {
	return func(o *generatorOptions) {
		o.Sensitive = true
		o.Lowercase = false
		o.Uppercase = false
		o.CamelCase = false
	}
}

// WithLowercase sets the initials to be lowercase
func WithLowercase() Option {
	return func(o *generatorOptions) {
		o.Lowercase = true
		o.Sensitive = false
		o.Uppercase = false
		o.CamelCase = false
	}
}

// WithUppercase sets the initials to be uppercase
func WithUppercase() Option {
	return func(o *generatorOptions) {
		o.Uppercase = true
		o.Sensitive = false
		o.Lowercase = false
		o.CamelCase = false
	}
}

// WithCamelCase sets the initials to be camel case
func WithCamelCase() Option {
	return func(o *generatorOptions) {
		o.CamelCase = true
		o.Sensitive = false
		o.Lowercase = false
		o.Uppercase = false
	}
}

// WithLength sets the maximum length of the initials
func WithLength(length int) Option {
	return func(o *generatorOptions) {
		o.Length = length
		o.remainingLength = length
	}
}

// WithWordLength sets the length to the number of words
func WithWordLength() Option {
	return func(o *generatorOptions) {
		o.WordLength = true
	}
}

// replaceSpecialCharacters removes special characters from the name using regex
func replaceSpecialCharacters(name string) string {
	// replace special characters using regex
	re := regexp.MustCompile("/[!@#$%^&*(),.?\":{}|<>_]/")
	return re.ReplaceAllString(name, "")
}

// replaceFriendlyCharacters replaces friendly characters with their English equivalents
func replaceFriendlyCharacters(name string) string {

	// replace friendly characters with their english equivalent
	// iterate replaceFriendlyCharacters map
	for key, value := range friendlyCharactersMap {
		for _, char := range value {
			name = strings.Replace(name, char, key, -1)
		}
	}

	re := regexp.MustCompile(`/[^\x20-\x7E]/u`)
	return re.ReplaceAllString(name, "")
}

// getInitial returns the first character of a word
func getInitial(word string, length int, opts *generatorOptions) string {

	// If the word is empty, return an empty string
	if len(word) == 0 {
		return ""
	}

	// Get the number of characters defined by the length parameter
	if len(word) < length {
		length = len(word)
	}

	// decompose the word into runes
	runes := []rune(word)

	// adjust the length if the word is shorter than the length
	if len(runes) < length {
		length = len(runes)
	}

	// get the initial
	initial := string(runes[0:length])

	// If the initials should be case-sensitive, return the first character
	if opts.Sensitive {
		return initial
	}

	// If the initials should be lowercase, return the lowercase first character
	if opts.Lowercase {
		return strings.ToLower(initial)
	}

	// If the initials should be uppercase, return the uppercase first character
	if opts.Uppercase {
		return strings.ToUpper(initial)
	}

	// If the initials should be camel case, return the camel case first character
	if opts.CamelCase {
		ir := []rune(initial)
		ir[0] = unicode.ToUpper(ir[0])
		return string(ir)
	}

	// Return the initial
	return initial
}
