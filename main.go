package main

import (
	"encoding/xml"
	"flag"
	"github.com/suzuken/jenkins-job-generator/jenkins"
	"io/ioutil"
	"os"
)

var (
	file    = flag.String("file", "path/to/file", "file that include shell")
	project = flag.String("project", "project name", "name of project")
)

func main() {
	flag.Parse()
	b, err := ioutil.ReadFile(*file)
	if err != nil {
		panic(err)
	}
	j := jenkins.NewProject(*project, string(b))
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(j); err != nil {
		panic(err)
	}
}
