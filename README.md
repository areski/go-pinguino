go-actionpinger
===============

This is a Go daemon/service that set and perform a list of actions based on the Http Get or Ping result.


Disclaimer
----------

This application aim to be run as Go Service (daemon). The project goals are very small, we just want to perform certain actions on an internal network based on microservice states.
We are not trying to replace nagios or any monitoring platform, Actionpinger was build to be use on personal computer.


Usage
-----

You may find Actionpinger useful if you want to activate/deactivate some local services or take action according to the output of webservices or state of your local network.


Install / Run
-------------

To run this application:

    $ git clone https://github.com/areski/go-actionpinger.git
    $ cd go-actionpinger
    $ go build .
    $ ./actionpinger

Config file need to be installed at the following location /etc/action_pinger.yaml


Configuration file
------------------

Config file `/etc/action_pinger.yaml`:

    # checker: check to trigger an action (HTTPGet | Ping)
    checker_type: "HTTPGet"

    # checker_source: URL or IP that will be checked
    checker_source: "http://192.168.1.1/"

    # checker_regex: Regular expresion to verify on source
    checker_regex: "RouterOS|WebFig"
    # <title>RouterOS router configuration page</title>

    # checker_freq: Frequence of check in seconds (300 -> 5min)
    checker_freq: 5

    # action to perform when checker_regex is true
    # leave action_cmd_* empty if no action
    action_cmd_on: "echo `date` >> /tmp/actionpinger.txt"

    # action to perform when checker_regex is false
    # leave action_cmd_* empty if no action
    action_cmd_off: "echo oupsss >> /tmp/actionpinger.txt"


License
-------

Go-actionpinger is licensed under MIT, see `LICENSE` file.


TODO
----

- [x] Add logging
- [x] Implement in Goroutine
- [x] Daemonize https://github.com/takama/daemon / https://github.com/sevlyar/go-daemon
- [ ] Implement checkPing method

