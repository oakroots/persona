package persona

import (
	"testing"

	"github.com/oakroots/persona/data"
	"github.com/stretchr/testify/assert"
)

func TestGenerator_GetFirstName(t *testing.T) {
	type fields struct {
		seed          uint32
		deterministic bool
		gender        Gender
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "get deterministic female name",
			fields: fields{
				seed:          1,
				deterministic: true,
				gender:        Female,
			},
			want: "Iga",
		},
		{
			name: "get deterministic male name",
			fields: fields{
				seed:          1,
				deterministic: true,
				gender:        Male,
			},
			want: "Kazimierz",
		},
		{
			name: "get deterministic default name",
			fields: fields{
				seed:          1,
				deterministic: true,
			},
			want: "Kirin",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Generator{
				seed:          tt.fields.seed,
				deterministic: tt.fields.deterministic,
				gender:        tt.fields.gender,
			}
			if got := g.GetFirstName(); got != tt.want {
				t.Errorf("GetFirstName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerator_GetFullName(t *testing.T) {
	type fields struct {
		seed          uint32
		deterministic bool
		gender        Gender
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "get deterministic female name",
			fields: fields{
				seed:          1,
				deterministic: true,
				gender:        Female,
			},
			want: "Iga Milewska",
		},
		{
			name: "get deterministic male name",
			fields: fields{
				seed:          1,
				deterministic: true,
				gender:        Male,
			},
			want: "Kazimierz Milewski",
		},
		{
			name: "get deterministic default name",
			fields: fields{
				seed:          1,
				deterministic: true,
			},
			want: "Kirin Whisperglen",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Generator{
				seed:          tt.fields.seed,
				deterministic: tt.fields.deterministic,
				gender:        tt.fields.gender,
			}
			if got := g.GetFullName(); got != tt.want {
				t.Errorf("GetFullName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerator_GetLastName(t *testing.T) {
	type fields struct {
		seed          uint32
		deterministic bool
		gender        Gender
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "get deterministic female name",
			fields: fields{
				seed:          1,
				deterministic: true,
				gender:        Female,
			},
			want: "Ostrowska",
		},
		{
			name: "get deterministic male name",
			fields: fields{
				seed:          1,
				deterministic: true,
				gender:        Male,
			},
			want: "Ostrowski",
		},
		{
			name: "get deterministic default name",
			fields: fields{
				seed:          1,
				deterministic: true,
			},
			want: "Moonbloom",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Generator{
				seed:          tt.fields.seed,
				deterministic: tt.fields.deterministic,
				gender:        tt.fields.gender,
			}
			if got := g.GetLastName(); got != tt.want {
				t.Errorf("GetLastName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerator_pickOne(t *testing.T) {
	type check func(got string, t *testing.T)
	checks := func(cs ...check) []check { return cs }

	var isEmpty = func() check {
		return func(got string, t *testing.T) {
			t.Helper()
			assert.Empty(t, got)
		}
	}

	var oneOf = func(expected []string) check {
		return func(got string, t *testing.T) {
			t.Helper()
			assert.Contains(t, expected, got)
		}
	}

	tests := []struct {
		name   string
		fetch  func() data.Names
		checks []check
	}{
		{
			name:   "empty list",
			fetch:  func() data.Names { return []string{} },
			checks: checks(isEmpty()),
		},
		{
			name:  "random",
			fetch: func() data.Names { return []string{"Anna", "Lina", "Rosa"} },
			checks: checks(
				oneOf([]string{"Anna", "Lina", "Rosa"}),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Generator{}
			got := g.pickOne(tt.fetch)
			for _, chk := range tt.checks {
				chk(got, t)
			}
		})
	}
}
