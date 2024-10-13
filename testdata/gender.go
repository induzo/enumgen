// Code generated by github.com/induzo/enumgen DO NOT EDIT.
package book

import (
	"fmt"

	"github.com/goccy/go-json"
)

type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
	GenderOther  Gender = "other"
)

func (gdr Gender) String() string {
	return string(gdr)
}

type GenderParseError struct {
	Gender string `json:"data"`
}

func (e GenderParseError) Error() string {
	return "invalid Gender value: " + e.Gender
}

func ParseGenderFromString(gdr string) (Gender, error) {
	switch gdr {
	case GenderMale.String():
		return GenderMale, nil
	case GenderFemale.String():
		return GenderFemale, nil
	case GenderOther.String():
		return GenderOther, nil
	default:
		return "", GenderParseError{Gender: gdr}
	}
}

func MustParseGenderFromString(gdr string) Gender {
	g, err := ParseGenderFromString(gdr)
	if err != nil {
		panic(err)
	}

	return g
}

func (gdr *Gender) MarshalJSON() ([]byte, error) {
	return []byte(`"` + gdr.String() + `"`), nil
}

type UnmarshalJSONGenderError struct {
	Data string `json:"data"`
	Err  error  `json:"error"`
}

func (e UnmarshalJSONGenderError) Error() string {
	return fmt.Sprintf("error parsing Gender `%s`: %v", e.Data, e.Err)
}

func (gdr *Gender) UnmarshalJSON(data []byte) error {
	var tmp string

	if errU := json.Unmarshal(data, &tmp); errU != nil {
		return UnmarshalJSONGenderError{
			Data: string(data),
			Err:  errU,
		}
	}

	var err error

	*gdr, err = ParseGenderFromString(tmp)
	if err != nil {
		return UnmarshalJSONGenderError{
			Data: string(data),
			Err:  err,
		}
	}

	return nil
}
