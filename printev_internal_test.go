package printev

import (
	"testing"
)

func Test_getLang(t *testing.T) {

	tests := map[string]struct {
		ext string
		want Lang
	}{
		"node": {ext:".js", want: Node},
		"ruby": {ext:".rb", want: Ruby},
		"yaml": {ext:".rb", want: Ruby},
		"go": {ext:".go", want: Go},
		"gotemplate": {ext:".go.template", want: Go},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := getLang(tt.ext); got != tt.want {
				t.Errorf("getLang() = %v, want %v", got, tt.want)
			}
		})
	}
}