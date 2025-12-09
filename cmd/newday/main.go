package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

const dayTemplate = `package day{{.DayPadded}}

func Run() (string, string) {
    data, _ := os.ReadFile("internal/day{{.DayPadded}}/input.txt")
    return Part1(string(data)), Part2(string(data))
}
`

const partTemplate = `package day{{.DayPadded}}

func Part{{.Part}}(input string) string {
    // TODO: implement part {{.Part}}
    return ""
}
`

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: newday <day-number>")
		os.Exit(1)
	}

	day := os.Args[1]
	if len(day) == 1 {
		day = "0" + day
	}

	dayDir := filepath.Join("internal", "day"+day)

	// Create directory
	if err := os.MkdirAll(dayDir, 0755); err != nil {
		panic(err)
	}

	// Create input.txt
	inputPath := filepath.Join(dayDir, "input.txt")
	os.WriteFile(inputPath, []byte(""), 0644)

	// Generate dayXX.go
	writeTemplate(filepath.Join(dayDir, "day"+day+".go"), dayTemplate, map[string]string{
		"DayPadded": day,
	})

	// Generate part1.go
	writeTemplate(filepath.Join(dayDir, "part1.go"), partTemplate, map[string]string{
		"DayPadded": day,
		"Part":      "1",
	})

	// Generate part2.go
	writeTemplate(filepath.Join(dayDir, "part2.go"), partTemplate, map[string]string{
		"DayPadded": day,
		"Part":      "2",
	})

	fmt.Printf("✨ Created template for Day %s!\n", day)
}

func writeTemplate(path string, tpl string, data any) {
	t := template.Must(template.New("file").Parse(tpl))
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	t.Execute(f, data)
}
