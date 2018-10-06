package utf8string

import "testing"

func Test_String(t *testing.T) {
	tests := []struct {
		in  string
		out String
	}{
		{
			in: `Hello world!`,
			out: String{
				NumRunes: 12,
				NumASCII: 12,
				Width:    12,
			},
		}, {
			in: `Hello 世界!`,
			out: String{
				NumRunes: 9,
				NumASCII: 7,
				Width:    11,
			},
		}, {
			in: `こんにちは世界`,
			out: String{
				NumRunes: 7,
				NumASCII: 0,
				Width:    14,
			},
		}, {
			in: `こんにちは
世界`,
			out: String{
				NumRunes: 8,
				NumASCII: 1,
				Width:    15,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		result := NewString(tt.in)
		switch {
		case tt.out.NumRunes != result.NumRunes:
			t.Errorf("want %d but got %d", tt.out.NumRunes, result.NumRunes)
		case tt.out.NumASCII != result.NumASCII:
			t.Errorf("want %d but got %d", tt.out.NumASCII, result.NumASCII)
		case tt.out.Width != result.Width:
			t.Errorf("want %d but got %d", tt.out.Width, result.Width)
		}
	}
}

func Test_Docs(t *testing.T) {
	short := NewString(`foo`)
	long := NewString(`foobarbazqux`)
	tests := []struct {
		in  Docs
		out *String
	}{
		{
			in:  Docs([]*String{short, long}),
			out: long,
		},
	}

	for _, tt := range tests {
		tt := tt
		result := tt.in.StringByMaxWidth()
		if tt.out != result {
			t.Errorf("want %v but got %v", *tt.out, *result)
		}
	}
}
