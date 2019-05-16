package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

type Bibliography struct {
	db map[string]Source
}

type Source struct {
	Indent       string
	Lang         string
	Type         string
	Title        string
	Desc         string
	Authors      []string
	URL          string
	RefferalDate time.Time
}

func (source Source) AuthorsStr() string {
	return strings.Join(source.Authors, " ")
}

func (source Source) MainAuthor() string {
	if len(source.Authors) > 0 {
		return source.Authors[0]
	}
	panic(`[Source.MainAuthor] source descriptions contains no authors!`)
}

func (source Source) Format(formatter SourceFormatter) (string, error) {
	return formatter(source)
}

type SourceFormatter func(source Source) (string, error)

func GOST(source Source) (string, error) {
	var buf = &bytes.Buffer{}
	var sourceTypes, validLang = sourceTypesLangMappings[source.Lang]
	if !validLang {
		sourceTypes = ruSources
	}
	const dateFormat = ""
	switch source.Type {
	case sourceTypes.Digital, sourceTypes.Digital, sourceTypes.Article:
		fmt.Fprintf(buf, "%s %s [%s] // %s. – режим доступа: %s (дата обращения: %s)",
			source.Title,
			source.AuthorsStr(),
			source.Type,
			source.Desc,
			source.URL,
			source.RefferalDate.Format(dateFormat))
	default:
		fmt.Fprintf(buf, "%s %s [%s] // %s",
			source.AuthorsStr(),
			source.Title,
			source.Type,
			source.Desc)
	}
	return buf.String(), nil
}

type CommonSourceTypes struct {
	Digital string
	Blog    string
	Article string
}

var sourceTypesLangMappings = map[string]CommonSourceTypes{
	"ru": ruSources,
}

var ruSources = CommonSourceTypes{
	Digital: "электронный ресурс",
	Blog:    "блог",
	Article: "статья",
}
