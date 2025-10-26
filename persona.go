package persona

import (
	"fmt"
	"strings"
	"sync"
)

type (
	Gender string

	Generator struct {
		mu            sync.Mutex
		seed          uint32
		deterministic bool
		gender        Gender
	}

	Option func(*Generator)
)

const (
	Male    Gender = "m"
	Female  Gender = "f"
	Fantasy Gender = "u"
)

func (g *Gender) String() string {
	switch *g {
	case Male:
		return "m"
	case Female:
		return "f"
	case Fantasy:
		return "u"
	default:
		return "?"
	}
}

func ParseGender(s string) (Gender, error) {
	switch strings.ToLower(s) {
	case "m":
		return Male, nil
	case "f":
		return Female, nil
	case "u":
		return Fantasy, nil
	}

	return "", fmt.Errorf("invalid gender %q", s)
}

func WithGender(gender Gender) Option {
	return func(g *Generator) { g.gender = gender }
}

func WithSeed(seed uint32) Option {
	return func(g *Generator) { g.seed = seed }
}

func WithDeterministic() Option {
	return func(g *Generator) { g.deterministic = true }
}

func New(opts ...Option) *Generator {
	g := &Generator{
		seed:          1,
		deterministic: false,
		gender:        Fantasy,
	}

	for _, opt := range opts {
		opt(g)
	}

	return g
}

func (g *Generator) Seed() uint32 {
	g.mu.Lock()
	defer g.mu.Unlock()

	return g.seed
}

//func (g *Generator) Deterministic() bool {
//	g.mu.Lock()
//	defer g.mu.Unlock()
//
//	return g.deterministic
//}
//
//func (g *Generator) Gender() Gender {
//	g.mu.Lock()
//	defer g.mu.Unlock()
//
//	return g.gender
//}
