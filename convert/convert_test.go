package convert

import (
	"strings"
	"testing"
)

func buildKoord(input string) Mgrs {
	ka := strings.Split(input, ";")

	return Mgrs{Zone: ka[0], Belt: ka[1], Kmkv: ka[2], East: ka[3], North: ka[4], Town: ka[5]}

}

func TestZoneNumeric(t *testing.T) {
	koord := buildKoord("32;U;NG;08600;77000;Somewhere")

	_, err := UtmAbs(koord)
	if err != nil {
		t.Errorf("Zone = %s cannot be converted to int", koord.Zone)
	}

	koord = buildKoord("3x2;U;NG;08600;77000;Somewhere")

	_, err = UtmAbs(koord)
	if err == nil {
		t.Errorf("test fails Zone = %s should not be converted to int", koord.Zone)
	}
}

func TestZoneTable(t *testing.T) {

	var tests = []struct {
		input string
		want  string
	}{
		{"32;U;NG;08600;77000;Somewhere", "pass"},
		{"3x2;U;NG;08600;77000;Somewhere", "fail"},
		{"321;U;NG;08600;77000;Somewhere", "fail"},
		{"-11;U;NG;08600;77000;Somewhere", "fail"},
	}

	for _, test := range tests {
		koord := buildKoord(test.input)
		_, err := UtmAbs(koord)
		if test.want == "pass" && err != nil {
			t.Errorf("Zone = %s cannot be converted to int", koord.Zone)
		}
		if test.want == "fail" && err == nil {
			t.Errorf("test fails Zone = %s er ikke en valid zone", koord.Zone)
		}

	}

}
