package enumgen

type FilenameError struct {
	EnumTypeName string
}

func (e *FilenameError) Error() string {
	return "enum name `" + e.EnumTypeName + "` leads to an empty filename"
}

type PathError struct {
	Path string
}

func (e *PathError) Error() string {
	return "path `" + e.Path + "` is not acceptable"
}
