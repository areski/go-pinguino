//
// This is where the magic happens...
//

package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	// "html/template"
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
	Action_cmd_on  string
	Action_cmd_off string
}

var config Config

// Default Template and Data directory
var CHECKER_TYPE = "scrape"
var CHECKER_SOURCE = "http://192.168.1.1/"
var CHECKER_REGEX = "RouterOS|WebFig"
var ACTION_CMD_ON = "echo `date` >> /tmp/actionpinger.txt"
var ACTION_CMD_OFF = "echo oupsss >> /tmp/actionpinger.txt"

func main() {
	// Parse CLI
	flag.Parse()

	config = Config{}

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
		// Change global Tempalte & Data vars
		CHECKER_TYPE = config.Checker_type
		CHECKER_SOURCE = config.Checker_source
		CHECKER_REGEX = config.Checker_regex
		ACTION_CMD_ON = config.Action_cmd_on
		ACTION_CMD_OFF = config.Action_cmd_off

	} else {
		panic("Config file defined properly.")
	}

	log.Printf("Starting action pinger...")
	fmt.Println("Welcome to my house!")
}
