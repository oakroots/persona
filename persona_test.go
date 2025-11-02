package persona

import (
	"sync"
	"testing"
)

func TestNew_Defaults(t *testing.T) {
	g := New()
	if g.gender != Fantasy {
		t.Fatalf("default gender = %q, want %q", g.gender, Fantasy)
	}
	if g.deterministic {
		t.Fatalf("default deterministic = true, want false")
	}
	if got := g.Seed(); got != 1 {
		t.Fatalf("default seed = %d, want 1", got)
	}
}

func TestOptions_Gender(t *testing.T) {
	tests := []struct {
		name   string
		option Option
		want   Gender
	}{
		{"male", WithGender(Male), Male},
		{"female", WithGender(Female), Female},
		{"fantasy", WithGender(Fantasy), Fantasy},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := New(tc.option)
			if g.gender != tc.want {
				t.Fatalf("gender = %q, want %q", g.gender, tc.want)
			}
		})
	}
}

func TestOptions_SeedAndDeterministic(t *testing.T) {
	g := New(WithSeed(123), WithDeterministic())
	if got := g.Seed(); got != 123 {
		t.Fatalf("seed = %d, want 123", got)
	}
	if !g.deterministic {
		t.Fatalf("deterministic = false, want true")
	}
}

func TestOptions_ChainAndOverrideOrder(t *testing.T) {
	// The last option should take precedence (for gender), seed is set once.
	g := New(
		WithGender(Male),
		WithGender(Female),
		WithSeed(7),
	)
	if g.gender != Female {
		t.Fatalf("gender = %q, want %q (last option wins)", g.gender, Female)
	}
	if got := g.Seed(); got != 7 {
		t.Fatalf("seed = %d, want 7", got)
	}
}

func TestParseGender_OK(t *testing.T) {
	tests := []struct {
		in   string
		want Gender
	}{
		{"m", Male},
		{"f", Female},
		{"u", Fantasy},
		{"M", Male},
		{"F", Female},
		{"U", Fantasy},
	}

	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			got, err := ParseGender(tc.in)
			if err != nil {
				t.Fatalf("ParseGender(%q) error: %v", tc.in, err)
			}
			if got != tc.want {
				t.Fatalf("ParseGender(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestParseGender_Err(t *testing.T) {
	invalid := []string{"", "x", "male", "female", "unknown", "1"}
	for _, in := range invalid {
		t.Run(in, func(t *testing.T) {
			if g, err := ParseGender(in); err == nil {
				t.Fatalf("ParseGender(%q) = %q, want error", in, g)
			}
		})
	}
}

func TestGender_String(t *testing.T) {
	tests := []struct {
		g    Gender
		want string
	}{
		{Male, "m"},
		{Female, "f"},
		{Fantasy, "u"},
		{Gender("whatever"), "?"},
	}
	for _, tc := range tests {
		t.Run(string(tc.g), func(t *testing.T) {
			g := tc.g
			if got := (&g).String(); got != tc.want {
				t.Fatalf("Gender(%q).String() = %q, want %q", g, got, tc.want)
			}
		})
	}
}

func TestSeed_ReadConcurrency(t *testing.T) {
	// Verifies that concurrent reads of Seed() do not cause race conditions.
	g := New(WithSeed(4242))
	var wg sync.WaitGroup
	const readers = 32

	for i := 0; i < readers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if got := g.Seed(); got != 4242 {
				t.Errorf("Seed() = %d, want 4242", got)
			}
		}()
	}
	wg.Wait()
}
