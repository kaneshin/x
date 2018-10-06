package x

import "testing"

func Test_SuddenDeath(t *testing.T) {
	tests := []struct {
		in, out string
	}{
		{
			in: `突然の死`,
			out: `＿人人人人人人＿
＞　突然の死　＜
￣ＹＹＹＹＹＹ￣`,
		}, {
			in: `突然



ののののののののの死`,
			out: `＿人人人人人人人人人人人人＿
＞　突然                　＜
＞　                    　＜
＞　                    　＜
＞　                    　＜
＞　ののののののののの死　＜
￣ＹＹＹＹＹＹＹＹＹＹＹＹ￣`,
		}, {
			in: `Sudden Death`,
			out: `＿人人人人人人人人＿
＞　Sudden Death　＜
￣ＹＹＹＹＹＹＹＹ￣`,
		}, {
			in: `突然の Death`,
			out: `＿人人人人人人人人＿
＞　突然の Death　＜
￣ＹＹＹＹＹＹＹＹ￣`,
		}, {
			in: ``,
			out: `＿人人＿
＞　　＜
￣ＹＹ￣`,
		},
	}

	for _, tt := range tests {
		tt := tt
		result := SuddenDeath(tt.in).String()
		if tt.out != result {
			t.Errorf("want %s but got %s", tt.out, result)
		}
	}
}
