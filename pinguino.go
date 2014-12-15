//
// This is where the magic happens...
//

// Configuration file:
// -------------------

// # checker: check to trigger an action (HTTPGet | Ping)
// # NOTE: Ping is not implemented yet
// checker_type: "HTTPGet"

// # checker_source: URL or IP that will be checked
// checker_source: "http://192.168.1.1/"

// # checker_regex: Regular expresion to verify on source
// checker_regex: "RouterOS|WebFig"
// # <title>RouterOS router configuration page</title>

// # checker_freq: Frequence of check in seconds (300 -> 5min)
// checker_freq: 5

// # action to perform when checker_regex is true (leave action_cmd_* empty if no action)
// # Use a tuple to define the command ie [ touch, /tmp/touchedfile.txt, ] or [./runme.sh, ]
// action_cmd_on: ["touch", "/tmp/touchedfile_on.txt", ]

// # action to perform when checker_regex is false ( leave action_cmd_* empty if no action)
// # Use a tuple to define the command ie [ touch, /tmp/touchedfile.txt, ] or [./runme.sh, ]
// action_cmd_off: ["touch", "/tmp/touchedfile_off.txt", ]

package main

import (
	// "flag"
	// "fmt"
	"errors"
	"github.com/codeskyblue/go-sh"
	"github.com/kr/pretty"
	"github.com/takama/daemon"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"
)

const check_HTTPGet string = "HTTPGet"
const check_Ping string = "Ping"

// Config file for actionpinger service
// var default_conf = "/etc/action_pinger.yaml"
var default_conf = "./pinguino.yaml"

// we create a point to string so we can return to use flag.
var configfile = &default_conf

// var (
// 	configfile = flag.String("configfile", "config.yaml", "path and filename of the config file")
// )

// Hold the structure for the wiki configuration
type Config struct {
	// First letter of variables need to be capital letter
	Checker_type   string
	Checker_source string
	Checker_regex  string
	Checker_freq   int
	Action_cmd_on  []string
	Action_cmd_off []string
}

var config = Config{}

const (

	// name of the service, match with executable file name
	name        = "pinguino"
	description = "Pinguino Service"
)

// Service has embedded daemon
type Service struct {
	daemon daemon.Daemon
	config Config
}

// Manage by daemon commands or run the daemon
func (service *Service) Manage() (string, error) {

	usage := "Usage: myservice install | remove | start | stop | status"

	// if received any kind of command, do it
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return service.daemon.Install()
		case "remove":
			return service.daemon.Remove()
		case "start":
			return service.daemon.Start()
		case "stop":
			return service.daemon.Stop()
		case "status":
			return service.daemon.Status()
		default:
			return usage, nil
		}
	}

	// Do something, call your goroutines, etc

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	// set up channel on which to receive communication and launch commands
	cmd_launcher := make(chan []string, 100)
	go performChecker(service.config, cmd_launcher)

	// loop work cycle which listen for command or interrupt
	// by system signal
	for {
		select {
		case command := <-cmd_launcher:
			go runCommand(command)
		case killSignal := <-interrupt:
			log.Println("Got signal:", killSignal)
			if killSignal == os.Interrupt {
				return "Daemon was interruped by system signal", nil
			}
			return "Daemon was killed", nil
		}
	}
	// never happen, but need to complete code
	return usage, nil
}

// function to check if a ping to an ip is successful
func checkPing(ipaddress string, checker_regex string) (bool, error) {
	// TODO: This method is not implemented yet
	log.Printf("checkPing - ipaddress:%s checker_regex:%s\n", ipaddress, checker_regex)
	return true, nil
}

// function to check the content of a URL against a regular expression
func checkHTTPGet(url string, checker_regex string) (bool, error) {
	log.Printf("checkHTTPGet - url:%s checker_regex:%s\n", url, checker_regex)
	response, err := http.Get(url)
	if err != nil {
		return false, err
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		if len(contents) > 0 {
			// fmt.Printf("checks again regular expersion:%s\n", checker_regex)
			// log.Println(string(contents))
			match, regerr := regexp.MatchString(checker_regex, string(contents))
			if regerr != nil {
				return false, errors.New("invalid regular expression")
			}
			return match, nil
		} else {
			return false, errors.New("not content return")
		}
	}
}

// function to launch action based on the Result
func launchCmdAction(rescheck bool, config Config, cmd_launcher chan<- []string) {
	log.Printf("Launch Action based on the result: %v", rescheck)
	// if rescheck is true or false, push command_on or command_off respectivily
	if rescheck && len(config.Action_cmd_on[0]) > 0 {
		cmd_launcher <- config.Action_cmd_on
	} else if !rescheck && len(config.Action_cmd_off[0]) > 0 {
		cmd_launcher <- config.Action_cmd_off
	} else {
		log.Printf("we dont have Action_cmd_on or Action_cmd_off to handle this case\n")
	}
}

// We will run checker here and send command to channel cmd_launcher depending of checker results
func performChecker(config Config, cmd_launcher chan<- []string) {
	// TODO: Unneed loop for?
	for {
		c := time.Tick(time.Duration(config.Checker_freq) * time.Second)
		for now := range c {
			switch config.Checker_type {
			case check_Ping:
				// TODO: This method is not implemented yet
				rescheck, cerr := checkPing(config.Checker_source, config.Checker_regex)
				if cerr != nil {
					log.Println(cerr)
					continue
				}
				launchCmdAction(rescheck, config, cmd_launcher)
			case check_HTTPGet:
				rescheck, cerr := checkHTTPGet(config.Checker_source, config.Checker_regex)
				if cerr != nil {
					log.Println(cerr)
					continue
				}
				launchCmdAction(rescheck, config, cmd_launcher)
			default:
				log.Printf("Checker type is incorrect: %v - %s\n", now, string(config.Checker_type))
				continue
			}
		}
	}
}

// Command runner
func runCommand(command []string) bool {
	if len(command) == 2 && len(command[0]) > 0 && len(command[1]) > 0 {
		log.Println("Run the command: ", command[0], command[1])
		sh.Command(command[0], command[1]).Run()
	} else if len(command) == 1 && len(command[0]) > 0 {
		log.Println("Run the command: ", command[0])
		sh.Command(command[0]).Run()
	} else if len(command) == 0 {
		return false
	}
	return true
}

// Load configuration file
func loadconfig() bool {
	if len(*configfile) > 0 {
		source, err := ioutil.ReadFile(*configfile)
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
	return true
}

func main() {

	// Load config
	loadconfig()

	if len(config.Checker_type) == 0 || len(config.Checker_source) == 0 || len(config.Checker_regex) == 0 {
		panic("Settings not properly configured!")
	}

	log.Println("Let's get the party started...")
	log.Printf("Loaded Config:\n%# v\n\n", pretty.Formatter(config))

	srv, err := daemon.New(name, description)
	if err != nil {
		log.Printf("Error: ", err)
		os.Exit(1)
	}
	service := &Service{daemon: srv, config: config}
	status, err := service.Manage()
	if err != nil {
		log.Printf(status, "\nError: ", err)
		os.Exit(1)
	}
	log.Println(status)
}
