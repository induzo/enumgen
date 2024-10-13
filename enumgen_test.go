package enumgen

import (
	"os"
	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestGenerateFiles(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		path    string
		data    *TemplateData
		wantErr bool
	}{
		{
			name:    "empty path",
			path:    "",
			wantErr: true,
		},
		{
			name:    "/ path",
			path:    "/",
			wantErr: true,
		},
		{
			name: ". path",
			path: ".",
			data: &TemplateData{
				PackageName:       "book",
				EnumTypeName:      "Rating",
				EnumTypeShortName: "rat",
				EnumValues:        []string{"not_good", "ok", "nice", "great"},
			},
			wantErr: false,
		},
		{
			name: "./ path",
			path: "./",
			data: &TemplateData{
				PackageName:       "book",
				EnumTypeName:      "Rating",
				EnumTypeShortName: "rat",
				EnumValues:        []string{"not_good", "ok", "nice", "great"},
			},
			wantErr: false,
		},
		{
			name: "wrong package name",
			path: "./",
			data: &TemplateData{
				PackageName:       "",
				EnumTypeName:      "Rating",
				EnumTypeShortName: "rat",
				EnumValues:        []string{"not_good", "ok", "nice", "great"},
			},
			wantErr: true,
		},
		{
			name: "wrong enum type name",
			path: "./",
			data: &TemplateData{
				PackageName:       "book",
				EnumTypeName:      "",
				EnumTypeShortName: "rat",
				EnumValues:        []string{"not_good", "ok", "nice", "great"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			genFiles, err := GenerateFiles(tt.path, tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateFiles() error = %v, wantErr %v", err, tt.wantErr)
			}

			// cleanup
			for _, genFile := range genFiles {
				if _, err := os.Stat(genFile); err == nil {
					if err := os.Remove(genFile); err != nil {
						t.Errorf("error removing path: %v", err)
					}
				}
			}
		})
	}
}
