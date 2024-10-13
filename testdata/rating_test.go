package book

import (
	"errors"
	"reflect"
	"testing"
)

func TestRating_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		payload []byte
		wantErr bool
	}{
		{
			name:    "correct",
			payload: []byte(`"not_good"`),
			wantErr: false,
		},
		{
			name:    "correct",
			payload: []byte(`"ok"`),
			wantErr: false,
		},
		{
			name:    "correct",
			payload: []byte(`"nice"`),
			wantErr: false,
		},
		{
			name:    "correct",
			payload: []byte(`"great"`),
			wantErr: false,
		},
		{
			name:    "error parsing",
			payload: []byte(`"xxx"`),
			wantErr: true,
		},
		{
			name:    "error unmarshalling",
			payload: []byte(`{}`),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var g Rating
			if err := g.UnmarshalJSON(tt.payload); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMustParseRatingFromString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		chRes   string
		want    Rating
		wantErr bool
	}{
		{
			name:  "correct not_good",
			chRes: "not_good",
			want:  RatingNotGood,
		},
		{
			name:  "correct ok",
			chRes: "ok",
			want:  RatingOk,
		},
		{
			name:  "correct nice",
			chRes: "nice",
			want:  RatingNice,
		},
		{
			name:  "correct great",
			chRes: "great",
			want:  RatingGreat,
		},
		{
			name:    "error",
			chRes:   "xxx",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if tt.wantErr {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()

				MustParseRatingFromString(tt.chRes)

				return
			}

			got := MustParseRatingFromString(tt.chRes)
			if got != tt.want {
				t.Errorf("MustParseFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRating_MarshalJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		chRes   Rating
		want    []byte
		wantErr bool
	}{
		{
			name:  "not_good",
			chRes: RatingNotGood,
			want:  []byte(`"not_good"`),
		},
		{
			name:  "ok",
			chRes: RatingOk,
			want:  []byte(`"ok"`),
		},
		{
			name:  "nice",
			chRes: RatingNice,
			want:  []byte(`"nice"`),
		},
		{
			name:  "great",
			chRes: RatingGreat,
			want:  []byte(`"great"`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.chRes.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnmarshalRatingError_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		Data string
		Err  error
		want string
	}{
		{
			name: "correct",
			Data: "data",
			Err:  nil,
			want: "error parsing Rating `data`: <nil>",
		},
		{
			name: "correct",
			Data: "data",
			Err:  errors.New("error"),
			want: "error parsing Rating `data`: error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := UnmarshalRatingError{
				Data: tt.Data,
				Err:  tt.Err,
			}

			if got := e.Error(); got != tt.want {
				t.Errorf("UnmarshalRatingError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRatingParseError_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		Rating string
		want   string
	}{
		{
			name:   "correct",
			Rating: "data",
			want:   "invalid Rating value: data",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := RatingParseError{
				Rating: tt.Rating,
			}

			if got := e.Error(); got != tt.want {
				t.Errorf("RatingParseError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
