package enumgen

import "testing"

func Test_constMaker(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "empty enum type name",
			input: "",
			want:  "",
		},
		{
			name:  "valid enum type name",
			input: "goodbook",
			want:  "Goodbook",
		},
		{
			name:  "valid enum type name with spaces",
			input: "good book",
			want:  "GoodBook",
		},
		{
			name:  "valid enum type name with underscore",
			input: "good_book",
			want:  "GoodBook",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := constMaker(tt.input); got != tt.want {
				t.Errorf("constMaker() = %v, want %v", got, tt.want)
			}
		})
	}
}
