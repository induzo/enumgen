package main

import (
	"flag"
	"log/slog"
	"strings"

	"github.com/induzo/enumgen"
)

func main() {
	genPath := flag.String("path", "./cmd/testdata", "Path to generate files")
	packageName := flag.String("package", "book", "Package name")
	enumTypeName := flag.String("enum", "Rating", "Enum type name")
	enumTypeShortName := flag.String("short", "rat", "Enum type short name")
	enumValuesStr := flag.String("values", "not_good,ok,nice,great", "Comma-separated enum values")

	flag.Parse()

	enumValues := strings.Split(*enumValuesStr, ",")

	data := enumgen.TemplateData{
		PackageName:       *packageName,
		EnumTypeShortName: *enumTypeShortName,
		EnumTypeName:      *enumTypeName,
		EnumValues:        enumValues,
	}

	files, errG := enumgen.GenerateFiles(*genPath, data)
	if errG != nil {
		slog.Error("error generating files", slog.Any("err", errG))

		return
	}

	slog.Info("files generated successfully", slog.String("files", strings.Join(files, ", ")))
}
