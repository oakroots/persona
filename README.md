# persona

**persona** is a Go library and command-line tool created out of a simple need —  
for another project, I wanted a way to generate both fantasy-style names and authentic Polish ones.  
Most existing solutions focused only on English names or lacked flexibility,  
so I decided to build my own.

While researching, I came across the excellent project [lzap/deagon](https://github.com/lzap/deagon),  
which shared a similar idea and inspired some of the internal mechanisms used here.

`persona` can be used as a standalone CLI tool for quick name generation,  
or as a Go library that you can integrate directly into your own applications.

---

## ✨ Features

- Deterministic name generation using a 25-bit Linear Feedback Shift Register (LFSR)
- Optional seed for reproducible sequences
- Multiple formatting styles (plain, capitalized, spaced, etc.)
- Built-in datasets for male, female, and fantasy names
- CLI tool (`persona`) for quick name generation from the terminal
- No external dependencies for the core library

---

## 📦 Installation

### CLI

```bash
go install github.com/oakroots/persona/cmd/persona@latest
```

Then run:

```bash
persona --help
```

### Library

```bash
go get github.com/oakroots/persona
```

---

## ⚙️ Usage (CLI)

```bash
# Generate a random name
persona

# Generate a male name
persona --gender m

# Generate a female name
persona --gender f

# Generate a fantasy name
persona --gender u

# Generate a deterministic sequence based on a seed
persona --seed 12345 --count 5

# Use a specific output format
persona --format capitalized-space
```

---

## 🧩 Usage (Library)

Example in Go:

```go
package main

import (
	"fmt"

	"github.com/oakroots/persona"
)

func main() {
	// Simple random name
	fmt.Println(persona.RandomName(persona.DefaultFormatter))

	// Deterministic name from a specific seed
	next, name := persona.PseudoRandomName(12345, false, persona.CapitalizedSpaceFormatter{})
	fmt.Printf("Seed: %d, Name: %s\n", next, name)

	// Using gender filtering
	opts := []persona.Option{
		persona.WithGender(persona.Female),
		persona.WithFormatter(persona.CapitalizedSpaceFormatter{}),
	}
	gen := persona.NewGenerator(opts...)
	fmt.Println(gen.Next())
}
```

---

## 🧠 Design

`persona` uses a 25-bit Linear Feedback Shift Register (LFSR) to generate pseudorandom indices into internal name tables.  
Each generated number corresponds to a specific combination of first and last names.  
When a seed is provided, name sequences can be reproduced exactly.

---

## ⚖️ License

This project is distributed under the MIT License.  
See the [`LICENSE`](./LICENSE) file for details.

---

## 🙌 Acknowledgements

This library and its internal logic were inspired by  
[lzap/deagon](https://github.com/lzap/deagon), created by Lukáš Zapletal.

---

## 🧰 Contributing

Pull requests and feature ideas are welcome.  
If you find a bug or have an idea for improvement, open an issue or submit a PR.

---

## 🏷 Example output

```bash
$ persona --count 5
Alandra Miren
Jareth Korrin
Thalina Veyra
Kaelen Durn
Riven Talar
```

---

Created with ❤️ by [oakroots](https://github.com/oakroots)
