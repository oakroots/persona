# persona

[![Go Reference](https://pkg.go.dev/badge/github.com/oakroots/persona.svg)](https://pkg.go.dev/github.com/oakroots/persona)
[![Go Report Card](https://goreportcard.com/badge/github.com/oakroots/persona)](https://goreportcard.com/report/github.com/oakroots/persona)
[![SonarCloud](https://sonarcloud.io/api/project_badges/measure?project=oakroots_persona&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=oakroots_persona)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=oakroots_persona&metric=coverage)](https://sonarcloud.io/summary/new_code?id=oakroots_persona)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Version](https://img.shields.io/github/v/tag/oakroots/persona?label=version)](https://github.com/oakroots/persona/releases)


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
persona --seed 12345 --num 5
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
	p := persona.New()

	fmt.Println(p.GetFullName())

	// Deterministic name from a specific seed
	p = persona.New(
		persona.WithSeed(12345),
		persona.WithDeterministic(),
	)

	for i := 0; i < 5; i++ {
		fmt.Printf("%s\n", p.GetFullName())
	}

	// Using gender filtering
	p = persona.New(
		persona.WithGender(persona.Female),
	)

	fmt.Println(p.GetFullName())
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
$ persona --seed 12345 --num 5
Rell Dawnchaser
Fenra Starglen
Ren Starborne
Quen Duskbloom
Arin Windhollow
7045132
```

---

Created with ❤️ by [oakroots](https://github.com/oakroots)
