package json

import "testing"

type neverzero int

func (nz neverzero) IsZero() bool {
	return false
}

type alwayszero int

func (az alwayszero) IsZero() bool {
	return true
}

type OptionalsIsZero struct {
	Sr string `json:"sr"`
	So string `json:"so,omitempty"`
	Sw string `json:"-"`

	Ir int `json:"omitempty"` // actually named omitempty, not an option
	Io int `json:"io,omitempty"`

	Slr []string `json:"slr,random"`
	Slo []string `json:"slo,omitempty"`

	Mr map[string]any `json:"mr"`
	Mo map[string]any `json:",omitempty"`

	Fr float64 `json:"fr"`
	Fo float64 `json:"fo,omitempty"`

	Br bool `json:"br"`
	Bo bool `json:"bo,omitempty"`

	Ur uint `json:"ur"`
	Uo uint `json:"uo,omitempty"`

	Str struct{} `json:"str"`
	Sto struct{} `json:"sto,omitempty"`

	Nzr neverzero `json:"nzr"`
	Nzo neverzero `json:"nzo,omitempty"`

	Azr alwayszero `json:"azr"`
	Azo alwayszero `json:"azo,omitempty"`

	PtrAzo *alwayszero `json:"ptrAzo,omitempty"`
}

var optionalsIsZeroExpected = `{
 "sr": "",
 "omitempty": 0,
 "slr": null,
 "mr": {},
 "fr": 0,
 "br": false,
 "ur": 0,
 "str": {},
 "sto": {},
 "nzr": 0,
 "nzo": 0,
 "azr": 0
}`

func TestOmitEmptyAndIsZero(t *testing.T) {
	var o OptionalsIsZero
	o.Sw = "something"
	o.Mr = map[string]any{}
	o.Mo = map[string]any{}

	got, err := MarshalIndent(&o, "", " ")
	if err != nil {
		t.Fatal(err)
	}
	if got := string(got); got != optionalsIsZeroExpected {
		t.Errorf(" got: %s\nwant: %s\n", got, optionalsIsZeroExpected)
	}
}
