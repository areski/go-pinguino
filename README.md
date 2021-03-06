Go-pinguino
===========

This is a Go daemon/service that performs a list of actions based on the result of Http Get or Ping.

[![circleci](https://circleci.com/gh/areski/go-pinguino.png)](https://circleci.com/gh/areski/go-pinguino)

[![Go Walker](http://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/areski/go-pinguino)


Disclaimer
----------

This application aim to be run as Go Service (daemon). The project goals are small, we just want to perform certain actions on an internal network based on microservice states.
We are not trying to replace nagios or any monitoring platform, Pinguino was build to be use on personal computer,
with minimal effort for deployment and dependencies.


Usage
-----

You may find Pinguino useful if you want to activate/deactivate some services or run custom actions on your computer/server based on the output of webservices and surroundings.


Install / Run
-------------

The config file need to be installed at the following location /etc/pinguino.yaml

To install and run the go-pinguino daemon, follow those steps:

    $ git clone https://github.com/areski/go-pinguino.git
    $ cd go-pinguino
    $ export GOPATH=`pwd`
    $ make build
    $ ./bin/daemon-pinguino


Configuration file
------------------

Config file `/etc/pinguino.yaml`:

    # checker: check to trigger an action (HTTPGet | Ping)
    # NOTE: Ping is not implemented yet
    checker_type: "HTTPGet"

    # checker_source: URL or IP that will be checked
    checker_source: "http://192.168.1.1/"

    # checker_regex: Regular expresion to verify on source
    checker_regex: "RouterOS|WebFig"
    # <title>RouterOS router configuration page</title>

    # checker_freq: Frequence of check in seconds (300 -> 5min)
    checker_freq: 5

    # action to perform when checker_regex is true (leave action_cmd_* empty if no action)
    # Use a tuple to define the command ie [ touch, /tmp/touchedfile.txt, ] or [./runme.sh, ]
    action_cmd_on: ["touch", "/tmp/touchedfile_on.txt", ]

    # action to perform when checker_regex is false ( leave action_cmd_* empty if no action)
    # Use a tuple to define the command ie [ touch, /tmp/touchedfile.txt, ] or [./runme.sh, ]
    action_cmd_off: ["touch", "/tmp/touchedfile_off.txt", ]


License
-------

Go-pinguino is licensed under MIT, see `LICENSE` file.


TODO
----

- [x] Add logging
- [x] Implement in Goroutine
- [x] Daemonize https://github.com/takama/daemon / https://github.com/sevlyar/go-daemon
- [x] Add test / travis-ci / Badge
- [x] godoc / https://gowalker.org
- [x] Review install/deploy documentation
- [ ] Handle several checkers: Update config file to support checker_type[0]: "...", checker_source[0]: "...", etc...
- [ ] Implement checkPing method
