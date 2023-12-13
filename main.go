// Copyright (C) 2023 Mya Pitzeruse
// SPDX-License-Identifier: MIT

package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/spdx/tools-golang/spdx/v2/v2_3"

	"go.pitz.tech/spdx-to-md/internal/templates"
)

//go:generate addlicense -f ./legal/header.txt -skip tmpl .

func run(input, output, report string) (err error) {
	fileIn := os.Stdin
	fileOut := os.Stdout

	if input != "-" {
		fileIn, err = os.Open(input)
		if err != nil {
			return err
		}

		defer fileIn.Close()
	}

	if output != "-" {
		fileOut, err = os.Create(output)
		if err != nil {
			return err
		}

		defer fileOut.Close()
	}

	doc := &v2_3.Document{}

	err = json.NewDecoder(fileIn).Decode(doc)
	if err != nil {
		return err
	}

	sort.Slice(doc.Packages, func(i, j int) bool {
		return strings.Compare(doc.Packages[i].PackageName, doc.Packages[j].PackageName) < 0
	})

	sort.Slice(doc.Files, func(i, j int) bool {
		return strings.Compare(doc.Files[i].FileName, doc.Files[j].FileName) < 0
	})

	contents, ok := templates.Lookup(report)
	if !ok {
		contents, err = os.ReadFile(report)
		if err != nil {
			return err
		}
	}

	t, err := template.New("spdx-to-md").
		Funcs(sprig.FuncMap()).
		Parse(string(contents))

	if err != nil {
		return err
	}

	return t.Execute(fileOut, doc)
}

func main() {
	cli := flag.NewFlagSet("spdx-to-md", flag.ExitOnError)

	input := cli.String("input", "-", "Specify the input spdx.json file. Defaults to '-' for stdin.")
	output := cli.String("output", "-", "Specify the output markdown file. Defaults to '-' for stdout.")
	report := cli.String("report", "spdx", "Specify which report to render (spdx, third-party-licenses, or a file path).")

	err := cli.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	err = run(*input, *output, *report)
	if err != nil {
		log.Fatal(err)
	}
}
