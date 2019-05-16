package main

type Bibliography struct {
	Title string
	db    map[string]Source
}

func NewBibliography(title string, sources ...Source) *Bibliography {
	var db = make(map[string]Source, len(sources))
	for _, source := range sources {
		db[source.Indent] = source
	}
	return &Bibliography{
		db: db,
	}
}

func (bibliograpgy *Bibliography) OrderedSources(order []string) Sources {
	var sources = make([]Source, 0, len(order))
	for _, indent := range order {
		sources = append(sources, bibliograpgy.db[indent])
	}
	return sources
}

func (bibliography *Bibliography) Indents() []string {
	var indents = make([]string, 0, len(bibliography.db))
	for indent := range bibliography.db {
		indents = append(indents, indent)
	}
	return indents
}
