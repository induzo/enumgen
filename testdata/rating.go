package book

import (
	"fmt"

	"github.com/goccy/go-json"
)

type Rating string

const (
	RatingNotGood Rating = "not_good"
	RatingOk      Rating = "ok"
	RatingNice    Rating = "nice"
	RatingGreat   Rating = "great"
)

func (rat Rating) String() string {
	return string(rat)
}

type RatingParseError struct {
	Rating string `json:"data"`
}

func (e RatingParseError) Error() string {
	return "invalid Rating value: " + e.Rating
}

func ParseRatingFromString(rat string) (Rating, error) {
	switch rat {
	case RatingNotGood.String():
		return RatingNotGood, nil
	case RatingOk.String():
		return RatingOk, nil
	case RatingNice.String():
		return RatingNice, nil
	case RatingGreat.String():
		return RatingGreat, nil
	default:
		return "", RatingParseError{Rating: rat}
	}
}

func MustParseRatingFromString(rat string) Rating {
	g, err := ParseRatingFromString(rat)
	if err != nil {
		panic(err)
	}

	return g
}

func (rat *Rating) MarshalJSON() ([]byte, error) {
	return []byte(`"` + rat.String() + `"`), nil
}

type UnmarshalRatingError struct {
	Data string `json:"data"`
	Err  error  `json:"error"`
}

func (e UnmarshalRatingError) Error() string {
	return fmt.Sprintf("error parsing Rating `%s`: %v", e.Data, e.Err)
}

func (rat *Rating) UnmarshalJSON(data []byte) error {
	var tmp string

	if errU := json.Unmarshal(data, &tmp); errU != nil {
		return UnmarshalRatingError{
			Data: string(data),
			Err:  errU,
		}
	}

	var err error

	*rat, err = ParseRatingFromString(tmp)
	if err != nil {
		return UnmarshalRatingError{
			Data: string(data),
			Err:  err,
		}
	}

	return nil
}
