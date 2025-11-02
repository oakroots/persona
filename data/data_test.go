package data

import (
	"reflect"
	"sync"
	"testing"
)

func TestSplitLines(t *testing.T) {
	tests := []struct {
		input NamesRaw
		want  Names
	}{
		{
			input: NamesRaw(""),
			want:  Names{},
		},
		{input: NamesRaw("Alice"), want: Names{"Alice"}},
		{input: NamesRaw("Alice\nBob"), want: Names{"Alice", "Bob"}},
		{input: NamesRaw("Alice\nBob\n"), want: Names{"Alice", "Bob"}},
		{input: NamesRaw("Alice\n\nBob\n"), want: Names{"Alice", "Bob"}},
		{input: NamesRaw(" Alice \n  Bob  \n\n"), want: Names{"Alice", "Bob"}},
		{input: NamesRaw("First\r\nSecond\r\nThird"), want: Names{"First", "Second", "Third"}},
	}

	for _, tt := range tests {
		got := splitLines(tt.input)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("splitLines(%#v) = %#v, want %#v", tt.input, got, tt.want)
		}
	}
}

func TestEmbeddedNameLists(t *testing.T) {
	tests := []struct {
		name string
		fn   func() Names
	}{
		{name: "FantasyFirstNames", fn: FantasyFirstNames},
		{name: "FantasyLastNames", fn: FantasyLastNames},
		{name: "MaleFirstNames", fn: MaleFirstNames},
		{name: "MaleLastNames", fn: MaleLastNames},
		{name: "FemaleFirstNames", fn: FemaleFirstNames},
		{name: "FemaleLastNames", fn: FemaleLastNames},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := make(chan Names, 5)

			var wg sync.WaitGroup
			start := make(chan struct{})

			for i := 0; i < 5; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					<-start
					result <- tc.fn()
				}()
			}
			close(start)
			wg.Wait()
			close(result)

			var firstRes Names
			for res := range result {
				if firstRes == nil {
					firstRes = res
				} else if !reflect.DeepEqual(res, firstRes) {
					t.Errorf("%s: incosistent results in concurrent calls", tc.name)
				}
			}

			if len(firstRes) == 0 {
				t.Errorf("%s returned an empty list (embed data missing or file empty", tc.name)
				return
			}

			got := tc.fn()
			if !reflect.DeepEqual(got, firstRes) {
				t.Errorf("%s: subsequent call returned different data", tc.name)
			}

		})
	}
}
