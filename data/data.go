package data

import (
	"bytes"
	_ "embed"
	"strings"
	"sync"
)

type (
	NamesRaw []byte
	Names    []string
)

var (
	//go:embed fantasy_unisex_first_names.txt
	fantasyFirstNamesRaw NamesRaw

	//go:embed fantasy_unisex_surnames.txt
	fantasyLastNamesRaw NamesRaw

	//go:embed male_first_names.txt
	maleFirstNamesRaw NamesRaw

	//go:embed male_last_names.txt
	maleLastNamesRaw NamesRaw

	//go:embed female_first_names.txt
	femaleFirstNamesRaw NamesRaw

	//go:embed female_last_names.txt
	femaleLastNamesRaw NamesRaw

	FantasyFirstNames = sync.OnceValue(func() Names { return splitLines(fantasyFirstNamesRaw) })
	FantasyLastNames  = sync.OnceValue(func() Names { return splitLines(fantasyLastNamesRaw) })
	FemaleFirstNames  = sync.OnceValue(func() Names { return splitLines(femaleFirstNamesRaw) })
	FemaleLastNames   = sync.OnceValue(func() Names { return splitLines(femaleLastNamesRaw) })
	MaleFirstNames    = sync.OnceValue(func() Names { return splitLines(maleFirstNamesRaw) })
	MaleLastNames     = sync.OnceValue(func() Names { return splitLines(maleLastNamesRaw) })
)

func splitLines(blob NamesRaw) Names {
	raw := bytes.Split(blob, []byte("\n"))
	out := make([]string, 0, len(raw))
	for _, r := range raw {
		s := strings.TrimSpace(string(r))
		if s != "" {
			out = append(out, s)
		}
	}

	return out
}
