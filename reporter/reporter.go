package reporter

import (
	"fmt"
	"os"
	"text/template"
)

type ReportData struct {
	Image   string
	Machine string
	Command string
	Results string
	Before  string
}

type Reporter struct {
	OutputFile string
}

// very simple markdown template for reporting
var tmpl = `## Benchmark results

{{range .RepData}}
#### {{.Image}}
**Machine**: _{{.Machine}}_

**Commands before benchmark**: _{{.Before}}_

**Benchmark command**: _{{.Command}}_
~~~
{{.Results}}
~~~
{{end}}

<sub><sup>Generated by [ben](https://github.com/drish/ben)</sup></sub>
`

// Creates a new reporter
func NewReporter(outputFile string) *Reporter {

	if outputFile == "" {
		outputFile = "benchmarks.md"
	}

	return &Reporter{
		OutputFile: outputFile,
	}
}

// Writes the benchmark sumary to fs as a markdown file
func (r *Reporter) Run(d []ReportData) error {

	f, _ := os.Create(r.OutputFile)
	defer f.Close()

	t := template.New("")
	t, _ = t.Parse(tmpl)

	t.Execute(f, struct {
		RepData []ReportData
	}{
		RepData: d,
	})

	fmt.Printf("\r  \033[36mwrote results to \033[m %s\n", r.OutputFile)
	return nil
}