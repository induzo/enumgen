package enumgen

import "testing"

func Test_generateFilename(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		enumTypeName string
		want         string
		wantErr      bool
	}{
		{
			name:         "empty enum type name",
			enumTypeName: "",
			wantErr:      true,
		},
		{
			name:         "valid enum type name",
			enumTypeName: "Rating",
			want:         "rating",
			wantErr:      false,
		},
		{
			name:         "valid enum type name with spaces",
			enumTypeName: "Rating 1",
			want:         "rating1",
			wantErr:      false,
		},
		{
			name:         "valid enum type name with special characters",
			enumTypeName: "Rating!@#$%^&*()",
			want:         "rating",
			wantErr:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := generateFilename(tt.enumTypeName)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateFilename() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if got != tt.want {
				t.Errorf("generateFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}
