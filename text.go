package main

import (
	"bytes"
	"fmt"
	"strings"
)

func AddBibliographyToText(text []string, bibliography *Bibliography, formatter SourceFormatter) ([]string, error) {
	var ordering = NewSourceOrdering(bibliography.Indents())
	text = append([]string{}, text...)
	for chunkID, chunk := range text {
		var locatedSources = ordering.ScanText(chunk)
		for source, index := range locatedSources {
			var textRef = fmt.Sprintf("[%s]", source)
			var numRef = fmt.Sprintf("[%d]", index)
			text[chunkID] = strings.ReplaceAll(chunk, textRef, numRef)
		}
	}
	var sources = bibliography.OrderedSources(ordering.OrderedSources())
	var formattedSources, errFormat = sources.Format(formatter)
	if errFormat != nil {
		return nil, errFormat
	}
	var formattedBibliography = bytes.NewBufferString(bibliography.Title + "\n")
	for i, formatted := range formattedSources {
		fmt.Fprintf(formattedBibliography, "%d. %s\n", i+1, formatted)
	}
	return append(text, formattedBibliography.String()), nil
}
