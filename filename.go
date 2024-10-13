package enumgen

import "strings"

func generateFilename(enumTypeName string) (string, error) {
	// remove all non alphanumeric characters
	filename := strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			return r
		}

		return -1
	}, enumTypeName)

	// if the length is 0, return an error
	if filename == `` {
		return "", &FilenameError{EnumTypeName: enumTypeName}
	}

	filename = strings.ToLower(filename)

	// convert to lower case
	return filename, nil
}
