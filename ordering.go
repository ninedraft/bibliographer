package main

import (
	"strings"
)

type SourceOrdering struct {
	indexes  map[string]int
	topIndex int
}

func NewSourceOrdering(sources []string) *SourceOrdering {
	var indexes = make(map[string]int, len(sources))
	for _, source := range sources {
		indexes[source] = 0
	}
	return &SourceOrdering{
		indexes:  indexes,
		topIndex: 0,
	}
}

func (ordering *SourceOrdering) IndexSource(source string) int {
	return ordering.indexes[source]
}

func (ordering *SourceOrdering) TopIndex() int {
	return ordering.topIndex
}

func (ordering *SourceOrdering) ScanText(text string) map[string]int {
	var locatedSources = make(map[string]int)
	for source, i := range ordering.indexes {
		var index = strings.Index(text, "["+source+"]")
		if index < 0 {
			continue
		}
		if i == 0 {
			ordering.topIndex++
			ordering.indexes[source] = ordering.topIndex
			locatedSources[source] = ordering.topIndex
		}
	}
	return locatedSources
}

func (ordering *SourceOrdering) Index() map[string]int {
	var cp = make(map[string]int, len(ordering.indexes))
	for source, index := range ordering.indexes {
		if index > 0 {
			cp[source] = index
		}
	}
	return cp
}

func (ordering *SourceOrdering) OrderedSources() []string {
	var orderedSources = make([]string, len(ordering.indexes))
	for source, index := range ordering.indexes {
		if index > 0 {
			orderedSources[index-1] = source
		}
	}
	return orderedSources
}
