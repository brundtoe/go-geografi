package utmabs

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
		field string
		want  string
	}{
		{"32;U;NG;08600;77000;Somewhere", "zone", "pass"},
		{"3x2;U;NG;08600;77000;Somewhere", "zone", "fail"},
		{"321;U;NG;08600;77000;Somewhere", "zone", "fail"},
		{"-11;U;NG;08600;77000;Somewhere", "zone", "fail"},
		{"32;N;NG;08600;77000;Somewhere", "belt", "pass"},
		{"32;X;NG;08600;77000;Somewhere", "belt", "pass"},
		{"32;U;G;08600;77000;Somewhere", "kmkv", "fail"},
		{"32;U;G;08600;77000;Somewhere", "kmkv", "fail"},
		{"32;U;NV;08600;77000;Somewhere", "kmkv", "pass"},
		{"32;U;NW;08600;77000;Somewhere", "kmkv", "fail"},
	}

	for _, test := range tests {
		koord := buildKoord(test.input)
		_, err := UtmAbs(koord)
		if test.want == "pass" && err != nil {
			t.Errorf("Field %s : %s", test.field, err)
		}
		if test.want == "fail" && err == nil {
			t.Errorf("Field %s : %s", test.field, err)
		}

	}
}
