package enumgen

import "testing"

func TestFilenameError_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		EnumTypeName string
		want         string
	}{
		{
			name:         "empty enum type name",
			EnumTypeName: "",
			want:         "enum name `` leads to an empty filename",
		},
		{
			name:         "valid enum type name",
			EnumTypeName: "Rating",
			want:         "enum name `Rating` leads to an empty filename",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := &FilenameError{
				EnumTypeName: tt.EnumTypeName,
			}

			if got := e.Error(); got != tt.want {
				t.Errorf("FilenameError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPathError_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		Path string
		want string
	}{
		{
			name: "empty path",
			Path: "",
			want: "path `` is not acceptable",
		},
		{
			name: "non empty path",
			Path: "/xxx/xx",
			want: "path `/xxx/xx` is not acceptable",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := &PathError{
				Path: tt.Path,
			}

			if got := e.Error(); got != tt.want {
				t.Errorf("PathError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
