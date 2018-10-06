package utf8string

import "unicode/utf8"

type String struct {
	Raw      string
	NumRunes int
	NumASCII int
	Width    int
}

func NewString(contents string) *String {
	s := String{
		Raw: contents,
	}
	for i := 0; i < len(s.Raw); i++ {
		if contents[i] < utf8.RuneSelf {
			s.NumASCII++
		}
	}
	if s.NumASCII == len(s.Raw) {
		s.NumRunes = len(s.Raw)
	} else {
		s.NumRunes = utf8.RuneCountInString(s.Raw)
	}
	s.Width = 2*s.NumRunes - s.NumASCII
	return &s
}

type Docs []*String

func (d Docs) StringByMaxWidth() *String {
	if len(d) == 0 {
		return nil
	}

	maxNum := 0
	maxWidth := d[maxNum].Width
	for i := 1; i < len(d); i++ {
		if maxWidth < d[i].Width {
			maxNum = i
			maxWidth = d[i].Width
		}
	}
	return d[maxNum]
}
