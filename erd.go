package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"syscall"
	"text/template"

	flags "github.com/jessevdk/go-flags"
	"golang.org/x/crypto/ssh/terminal"
)

// Options for the command line tool
type Options struct {
	OutFormat  string `short:"f" long:"fmt" description:"output format (dot only)"`
	InputFile  string `short:"i" long:"input" description:"input will be read from the given file."`
	OutputFile string `short:"o" long:"output" description:"output will be written to the given file."`
}

var opts Options

func main() {
	contents := ""
	logStderr := log.New(os.Stderr, "", 0)

	optsParser := flags.NewParser(&opts, flags.Default)
	optsParser.Name = filepath.Base(os.Args[0])
	optsParser.Usage = "[OPTIONS] PATTERN [PATH]"

	args, err := optsParser.Parse()
	if err != nil {
		logStderr.Println(err)
		os.Exit(1)
	}

	if terminal.IsTerminal(int(syscall.Stdin)) {
		if len(args) == 0 && opts.InputFile == "" {
			optsParser.WriteHelp(os.Stdout)
			os.Exit(1)
		}
		buffer, err := ioutil.ReadFile(opts.InputFile)
		if err != nil {
			logStderr.Println(err)
			os.Exit(1)
		}
		contents = string(buffer)
	} else {
		body, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			logStderr.Println(err)
			os.Exit(1)
		}
		contents = string(body)
	}

	parser := &Parser{Buffer: contents}
	parser.Init()
	err = parser.Parse()
	if err != nil {
		logStderr.Println(err)
		os.Exit(1)
	}

	parser.Execute()

	if parser.Erd.IsError {
		logStderr.Println("Parser error")
		os.Exit(1)
	}

	dot, _ := Asset("templates/dot.tmpl")
	tables, _ := Asset("templates/dot_tables.tmpl")
	relations, _ := Asset("templates/dot_relations.tmpl")
	clusters, _ := Asset("templates/dot_subgraphs.tmpl")
	templates := template.Must(
		template.New("").Parse(
			string(dot) +
				string(tables) +
				string(relations) +
				string(clusters)))

	fd := os.Stdout
	if opts.OutputFile != "" {
		fd, err = os.Create(opts.OutputFile)
		if err != nil {
			logStderr.Println(err)
			os.Exit(1)
		}
	}

	templates.ExecuteTemplate(fd, "dot", parser.Erd)
}
