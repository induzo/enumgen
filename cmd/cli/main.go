package main

import (
	"flag"
	"log/slog"
	"os"
	"strings"

	"github.com/induzo/enumgen"
)

func main() {
	genPath := flag.String("path", "./testdata", "Path to generate files")
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

	// Get the current working directory (the one `go generate` runs in).
	cwd, errD := os.Getwd()
	if errD != nil {
		slog.Error("error getting current working directory", slog.Any("err", errD))
		os.Exit(1)
	}

	path := cwd + `/` + *genPath

	slog.Info("generating files",
		slog.String("path", path),
		slog.String("package", *packageName),
		slog.String("enum", *enumTypeName),
		slog.String("short", *enumTypeShortName),
		slog.String("values", strings.Join(enumValues, ", ")),
	)

	files, errG := enumgen.GenerateFiles(path, data)
	if errG != nil {
		slog.Error("error generating files", slog.Any("err", errG))

		return
	}

	slog.Info("files generated successfully", slog.String("files", strings.Join(files, ", ")))
}
