package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
)

func main() {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln("unable to read input", err)
	}

	out, err := yaml.JSONToYAML(in)
	if err != nil {
		log.Fatalln("unable to convert yaml to json", err)
	}

	fmt.Print(string(out))
}
