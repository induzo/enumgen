package enumgen

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_generateFromTmpl(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		templateFile      string
		data              TemplateData
		fileToCompareWith string
		want              []byte
		wantErr           bool
	}{
		{
			name:    "empty template file",
			wantErr: true,
		},
		{
			name:         "empty data",
			templateFile: "enum.go.tmpl",
			wantErr:      true,
		},
		{
			name:         "valid data",
			templateFile: "enum.go.tmpl",
			data: TemplateData{
				PackageName:       "book",
				EnumTypeName:      "Rating",
				EnumTypeShortName: "rat",
				EnumValues:        []string{"not_good", "ok", "nice", "great"},
			},
			fileToCompareWith: "testdata/rating.go",
			wantErr:           false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := generateFromTmpl(tt.templateFile, tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateFromTmpl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Read and open fileToCompareWith
			if tt.fileToCompareWith != "" {
				gotFromFile, err := os.ReadFile(tt.fileToCompareWith)
				if err != nil {
					t.Errorf("error reading file: %v", err)
				}

				if !cmp.Equal(got, gotFromFile) {
					t.Errorf("generateFromTmpl() = %s", cmp.Diff(got, gotFromFile))
				}
			}
		})
	}
}
