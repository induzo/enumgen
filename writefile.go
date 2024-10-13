package enumgen

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
)

func writeToFile(filename string, content []byte) error {
	// Write the content to a file
	out, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}

	defer func() {
		if errC := out.Close(); errC != nil {
			slog.Error("err closing file", slog.Any("err", errC))
		}
	}()

	writer := bufio.NewWriter(out)

	if _, errW := writer.Write(content); errW != nil {
		return fmt.Errorf("error writing to file: %w", errW)
	}

	if errF := writer.Flush(); errF != nil {
		return fmt.Errorf("error flushing writer: %w", errF)
	}

	return nil
}
