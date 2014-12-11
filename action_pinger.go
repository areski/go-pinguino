//
// This is where the magic happens...
//

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

package main

import (
	// "flag"
	"fmt"
	"github.com/kr/pretty"
	"github.com/takama/daemon"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Config file for actionpinger service
// var default_conf = "/etc/action_pinger.yaml"
var default_conf = "./action_pinger.yaml"

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
	Action_cmd_on  string
	Action_cmd_off string
}

var config = Config{}

const (

	// name of the service, match with executable file name
	name        = "action_pinger"
	description = "Action Pinger Service"
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
	cmd_launcher := make(chan string, 100)
	go procChecker(service.config, cmd_launcher)

	// loop work cycle with accept connections or interrupt
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

// Accept a client connection and collect it in a channel
func procChecker(config Config, cmd_launcher chan<- string) {
	for {
		c := time.Tick(10 * time.Second)
		for now := range c {
			fmt.Printf("%v\n", now)
			cmd_launcher <- "a_command_to_launch"
		}
	}
}

func runCommand(command string) {
	// we will run run a command here
	fmt.Println("hey hey..." + command)
}

func main() {
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

	// ------------------------

	srv, err := daemon.New(name, description)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	service := &Service{daemon: srv, config: config}
	status, err := service.Manage()
	if err != nil {
		fmt.Println(status, "\nError: ", err)
		os.Exit(1)
	}
	fmt.Println(status)
}
