package initials

import "testing"

// struct used to test the initials generator
type test struct {
	name    string
	options []Option
	want    string
}

// Test empty name
func TestGetInitials_EmptyName(t *testing.T) {

	tt := test{}

	if got := GetInitials(tt.name, tt.options...); got != tt.want {
		t.Errorf("GetInitials() = %v, want %v", got, tt.want)
	}
}

// Test with a simple name
func TestGetInitials_Simple(t *testing.T) {

	tt := test{
		name: "John Doe",
		want: "JD",
	}

	if got := GetInitials(tt.name, tt.options...); got != tt.want {
		t.Errorf("GetInitials() = %v, want %v", got, tt.want)
	}
}

// Test with a simple name and a separator
func TestGetInitials_SimpleWithSeparator(t *testing.T) {

	tts := []test{
		{
			name:    "John Doe",
			options: []Option{WithSeparator(".")},
			want:    "J.D",
		},
		{
			name:    "John Doe",
			options: []Option{WithSeparator("-")},
			want:    "J-D",
		},
	}

	for _, tt := range tts {
		if got := GetInitials(tt.name, tt.options...); got != tt.want {
			t.Errorf("GetInitials() = %v, want %v", got, tt.want)
		}
	}
}

// Test with different case options
func TestGetInitials_CaseOptions(t *testing.T) {

	tts := []test{
		{
			name:    "John Doe",
			options: []Option{WithSensitive()},
			want:    "JD",
		},
		{
			name:    "John doe",
			options: []Option{WithSensitive()},
			want:    "Jd",
		},
		{
			name:    "john doe",
			options: []Option{WithSensitive()},
			want:    "jd",
		},
		{
			name:    "John Doe",
			options: []Option{WithLowercase()},
			want:    "jd",
		},
		{
			name:    "John doe",
			options: []Option{WithLowercase()},
			want:    "jd",
		},
		{
			name:    "john doe",
			options: []Option{WithLowercase()},
			want:    "jd",
		},
		{
			name:    "John Doe",
			options: []Option{WithUppercase()},
			want:    "JD",
		},
		{
			name:    "John doe",
			options: []Option{WithUppercase()},
			want:    "JD",
		},
		{
			name:    "john doe",
			options: []Option{WithUppercase()},
			want:    "JD",
		},
	}

	for _, tt := range tts {
		if got := GetInitials(tt.name, tt.options...); got != tt.want {
			t.Errorf("GetInitials() = %v, want %v", got, tt.want)
		}
	}
}

// Test with different length options
func TestGetInitials_LengthOptions(t *testing.T) {

	tts := []test{
		{
			name:    "John Doe",
			options: []Option{WithLength(2)},
			want:    "JD",
		},
		{
			name:    "John Doe",
			options: []Option{WithLength(2), WithCamelCase()},
			want:    "JD",
		},
		{
			name:    "John Doe",
			options: []Option{WithLength(3)},
			want:    "JDO",
		},
		{
			name:    "John Doe",
			options: []Option{WithLength(3), WithCamelCase()},
			want:    "JDo",
		},
		{
			name:    "John Doe",
			options: []Option{WithLength(4), WithCamelCase()},
			want:    "JDoe",
		},
		{
			name:    "John",
			options: []Option{WithLength(2)},
			want:    "JO",
		},
		{
			name:    "John",
			options: []Option{WithLength(3)},
			want:    "JOH",
		},
		{
			name:    "John",
			options: []Option{WithLength(2), WithCamelCase()},
			want:    "Jo",
		},
		{
			name:    "John",
			options: []Option{WithLength(3), WithCamelCase()},
			want:    "Joh",
		},
	}

	for _, tt := range tts {
		if got := GetInitials(tt.name, tt.options...); got != tt.want {
			t.Errorf("GetInitials() = %v, want %v", got, tt.want)
		}
	}
}

// Test with length option higher than the number of characters
func TestGetInitials_LengthOptionHigherThanNumberOfCharacters(t *testing.T) {

	tts := []test{
		{
			name:    "John Doe",
			options: []Option{WithLength(10)},
			want:    "JDOE",
		},
		{
			name:    "John Doe",
			options: []Option{WithLength(10), WithCamelCase()},
			want:    "JDoe",
		},
	}

	for _, tt := range tts {
		if got := GetInitials(tt.name, tt.options...); got != tt.want {
			t.Errorf("GetInitials() = %v, want %v", got, tt.want)
		}
	}
}

// Test with dash in the firstname like in Mary-Kate Olsen
func TestGetInitials_Dash(t *testing.T) {

	tts := []test{
		{
			name:    "Mary-Kate Olsen",
			options: []Option{WithLength(2)},
			want:    "MK",
		},
		{
			name:    "Mary-Kate Olsen",
			options: []Option{WithLength(3)},
			want:    "MKO",
		},
		{
			name:    "Mary-Kate Olsen",
			options: []Option{WithLength(4)},
			want:    "MKOL",
		},
		{
			name:    "Mary-Kate Olsen",
			options: []Option{WithLength(4), WithCamelCase()},
			want:    "MKOl",
		},
		{
			name:    "M@ry Olsen",
			options: []Option{WithLength(2)},
			want:    "MO",
		},
		{
			name:    "M@ry Olsen",
			options: []Option{WithLength(3)},
			want:    "MOL",
		},
		{
			name:    "M@ry Olsen",
			options: []Option{WithLength(4)},
			want:    "MOLS",
		},
		{
			name:    "M@ry Olsen",
			options: []Option{WithLength(4), WithCamelCase()},
			want:    "MOls",
		},
	}

	for _, tt := range tts {
		if got := GetInitials(tt.name, tt.options...); got != tt.want {
			t.Errorf("GetInitials() = %v, want %v", got, tt.want)
		}
	}
}

// Test with word length option
func TestGetInitials_WordLengthOption(t *testing.T) {

	tts := []test{
		{
			name:    "John Doe",
			options: []Option{WithWordLength()},
			want:    "JD",
		},
		{
			name:    "Mary-Kate Olsen",
			options: []Option{WithWordLength()},
			want:    "MKO",
		},
		{
			name:    "Mary-Kate Olsen",
			options: []Option{WithWordLength(), WithCamelCase()},
			want:    "MKO",
		},
	}

	for _, tt := range tts {
		if got := GetInitials(tt.name, tt.options...); got != tt.want {
			t.Errorf("GetInitials() = %v, want %v", got, tt.want)
		}
	}
}

// Test with emoji
func TestGetInitials_Emoji(t *testing.T) {

	tts := []test{
		{
			name:    "😅",
			options: []Option{},
			want:    "😅",
		},
		{
			name:    "😅",
			options: []Option{WithLength(2)},
			want:    "😅",
		},
		{
			name:    "😅",
			options: []Option{WithLength(2), WithCamelCase()},
			want:    "😅",
		},
	}

	for _, tt := range tts {
		if got := GetInitials(tt.name, tt.options...); got != tt.want {
			t.Errorf("GetInitials() = %v, want %v", got, tt.want)
		}
	}
}
