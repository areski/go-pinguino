//
// This is where the magic happens...
//

package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	// "html/template"
	"github.com/kr/pretty"
	"io/ioutil"
	"log"
)

var (
	configfile = flag.String("configfile", "config.yaml", "path and filename of the config file")
)

// Hold the structure for the wiki configuration
type Config struct {
	// First letter of variables need to be capital letter
	Checker_type   string
	Checker_source string
	Checker_regex  string
	Checker_freq   int
	Action_cmd_on  string
	Action_cmd_off string
}

var config = Config{}

func main() {
	// Parse CLI
	flag.Parse()

	// Load configfile and configure template
	if len(*configfile) > 0 {
		source, err := ioutil.ReadFile(*configfile)
		// fmt.Println(string(source))
		if err != nil {
			panic(err)
		}
		// decode the yaml source
		err = yaml.Unmarshal(source, &config)
		if err != nil {
			panic(err)
		}

	} else {
		panic("Config file defined properly.")
	}

	if len(config.Checker_type) == 0 || len(config.Checker_source) == 0 || len(config.Checker_regex) == 0 {
		panic("Settings not properly configured!")
	}

	log.Printf("Starting action pinger...")

	fmt.Println("Let's get the party started...")
	fmt.Printf("%# v", pretty.Formatter(config))

}

// # checker: check to trigger an action (scrape | ping)
// checker_type: "scrape"

// # checker_source: URL or IP that will be checked
// checker_source: "http://192.168.1.1/"

// # checker_regex: Regular expresion to verify on source
// checker_regex: "RouterOS|WebFig"

// # action to perform when checker_regex is true
// # leave action_cmd_* empty if no action
// action_cmd_on: "echo `date` >> /tmp/actionpinger.txt"

// # action to perform when checker_regex is false
// # leave action_cmd_* empty if no action
// action_cmd_off: "echo oupsss >> /tmp/actionpinger.txt"
