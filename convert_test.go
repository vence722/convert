package convert

import "testing"

func TestBytesSlice2StrNonCopy(t *testing.T) {
	origin := "Vence is smart"
	originBytes := []byte(origin)
	converted := BytesSlice2StrNonCopy(originBytes)
	if origin != converted {
		t.Fail()
	}
}