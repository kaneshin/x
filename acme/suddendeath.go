package acme

import (
	"fmt"
	"strings"

	"github.com/kaneshin/x/utf8string"
)

type SuddenDeath string

func (s SuddenDeath) Bytes() []byte {
	return []byte(s.String())
}

func (s SuddenDeath) String() string {
	contents := strings.Split(strings.TrimRight(string(s), "\n"), "\n")
	output := make([]string, 0, len(contents)+2)

	docs := make([]*utf8string.String, 0, len(contents))
	for i := 0; i < len(contents); i++ {
		docs = append(docs, utf8string.NewString(contents[i]))
	}
	max := utf8string.Docs(docs).StringByMaxWidth()

	size := max.Width/2 + 3
	t := make([]string, size)
	t[0], t[len(t)-1] = "＿", "＿"
	output = append(output, strings.Join(t, "人"))
	for _, v := range docs {
		format := fmt.Sprintf("＞　%%s%%%ds　＜", max.Width-v.Width)
		output = append(output, fmt.Sprintf(format, v.Raw, ""))
	}
	t[0], t[len(t)-1] = "￣", "￣"
	output = append(output, strings.Join(t, "Ｙ"))
	return strings.Join(output, "\n")
}
