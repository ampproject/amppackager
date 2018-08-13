/*
Transform invokes the transformer library on a given input doc to
generate and output transformed AMP HTML. This does not validate the
document.

See flag.Usage in main() for usage instructions.
*/
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	t "github.com/ampproject/amppackager/pkg/transform"
)

var transformersFlag = flag.String("transformers", "", "Comma-separated list of transformers to execute.")
var documentURLFlag = flag.String("document URL", "", "The URL of the document being processed")

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	// Custom usage message.
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nUsage %s [OPTION] [FILE]\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, `
Examples:

# Execute with filename
$GOPATH/bin/transform -transformers=URLTransformer,AMPBoilerplateTransformer \
   /path/to/input.html

# Execute with pipe
cat /path/to/input.html | $GOPATH/bin/transform -transformers=URLTransformer,AMPBoilerplateTransformer
`)
	}

	flag.Parse()
	var transformers []string
	if *transformersFlag != "" {
		transformers = strings.Split(*transformersFlag, ",")
		for i := range transformers {
			transformers[i] = strings.TrimSpace(transformers[i])
		}
	}

	var data []byte
	var err error
	switch flag.NArg() {
	case 0:
		data, err = ioutil.ReadAll(os.Stdin)
	case 1:
		data, err = ioutil.ReadFile(flag.Arg(0))
	default:
		log.Fatal("Input must be from stdin or file.")
	}
	checkErr(err)
	o, err := t.Process(string(data), *documentURLFlag, transformers)
	checkErr(err)
	fmt.Print(o)
}
