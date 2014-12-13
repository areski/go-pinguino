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


License
-------

Go-actionpinger is licensed under MIT, see `LICENSE` file.


TODO
----

- [x] Add logging
- [x] Implement in Goroutine
- [x] Daemonize https://github.com/takama/daemon / https://github.com/sevlyar/go-daemon
- [ ] Implement checkPing method

