package enumgen

import (
	"os"
	"testing"
)

func Test_writeToFile(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		filename string
		content  []byte
		wantErr  bool
	}{
		{
			name:     "empty filename",
			filename: "",
			content:  []byte(""),
			wantErr:  true,
		},
		{
			name:     "empty content",
			filename: "test.txt",
			content:  []byte(""),
			wantErr:  false,
		},
		{
			name:     "non-empty content",
			filename: "test.txt",
			content:  []byte("hello world"),
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if err := writeToFile(tt.filename, tt.content); (err != nil) != tt.wantErr {
				t.Errorf("writeToFile() error = %v, wantErr %v", err, tt.wantErr)
			}

			// cleanup
			if _, err := os.Stat(tt.filename); err == nil {
				if err := os.Remove(tt.filename); err != nil {
					t.Errorf("error removing file: %v", err)
				}
			}
		})
	}
}
