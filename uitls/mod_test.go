package uitls

import "testing"

func TestPkgMode(t *testing.T) {
	result := PkgMode()
	if result != "github.com/go-kenka/esql\n" {
		t.Fail()
	}
}
