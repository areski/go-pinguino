Go-pinguino
===========

This is a Go daemon/service that set and perform a list of actions based on the Http Get or Ping result.

[![circleci](https://circleci.com/gh/areski/go-pinguino.png)](https://circleci.com/gh/areski/go-pinguino)


Disclaimer
----------

This application aim to be run as Go Service (daemon). The project goals are very small, we just want to perform certain actions on an internal network based on microservice states.
We are not trying to replace nagios or any monitoring platform, Pinguino was build to be use on personal computer.


Usage
-----

You may find Pinguino useful if you want to activate/deactivate some local services or take action according to the output of webservices or state of your local network.


Install / Run
-------------

To run this application:

    $ git clone https://github.com/areski/go-pinguino.git
    $ cd go-pinguino
    $ go build .
    $ ./pinguino

Config file need to be installed at the following location /etc/pinguino.yaml


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
- [ ] godoc / https://gowalker.org
- [ ] Review install/deploy documentation
- [ ] Implement checkPing method
