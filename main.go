package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// TODO: directory support

// JUnitReport represents either a single test suite or a collection of test suites
type JUnitReport struct {
	XMLName   xml.Name
	XML       string       `xml:",innerxml"`
	Name      string       `xml:"name,attr"`
	Time      float64      `xml:"time,attr"`
	Tests     uint64       `xml:"tests,attr"`
	Failures  uint64       `xml:"failures,attr"`
	XMLBuffer bytes.Buffer `xml:"-"`
}

var outputFileName string

func init() {
	flag.StringVar(&outputFileName, "o", "", "Merged report filename")
}

func main() {
	flag.Parse()
	files := flag.Args()
	printReport := outputFileName == ""

	if len(files) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	var mergedReport JUnitReport
	startedReading := false
	fileCount := 0

	for _, fileName := range files {
		var report JUnitReport
		in, err := ioutil.ReadFile(filepath.Clean(fileName))
		if err != nil {
			log.Fatalf("Unable to read file '%s': %s", fileName, err)
		}

		if err = xml.Unmarshal(in, &report); err != nil {
			log.Fatalf("Unable to unmarshal report '%s': %s", fileName, err)
		}

		if report.XMLName.Local == "testsuite" {
			log.Fatalf("Unable to read report '%s': reports with a root <testsuite> are not supported", fileName)
		}

		if startedReading && report.Name != mergedReport.Name {
			log.Fatalf("Unable to read report '%s': all reports must have the same <testsuites> name", fileName)
		}

		startedReading = true
		fileCount++
		mergedReport.XMLName = xml.Name{Local: "testsuites"}
		mergedReport.Name = report.Name
		mergedReport.Time += report.Time
		mergedReport.Tests += report.Tests
		mergedReport.Failures += report.Failures
		mergedReport.XMLBuffer.WriteString(report.XML)
	}

	mergedReport.XML = mergedReport.XMLBuffer.String()
	mergedOutput, _ := xml.MarshalIndent(&mergedReport, "", "  ")

	if printReport {
		log.Println(string(mergedOutput))
		return
	}

	if err := ioutil.WriteFile(outputFileName, mergedOutput, 0o600); err != nil {
		log.Fatalf("Unable to save reports to the output file '%s': %s", outputFileName, err)
	}

	log.Printf("Merged %d reports to file '%s'", fileCount, outputFileName)
}
