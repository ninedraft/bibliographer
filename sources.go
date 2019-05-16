package main

type Sources []Source

func (sources Sources) Len() int {
	return len(sources)
}

func (sources Sources) Format(formatter SourceFormatter) ([]string, error) {
	var strs = make([]string, 0, sources.Len())
	for _, source := range sources {
		var formatted, err = source.Format(formatter)
		if err != nil {
			return nil, err
		}
		strs = append(strs, formatted)
	}
	return strs, nil
}
