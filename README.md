go-actionpinger
===============

Go service that run an action when pinging to a specific source is off/on.

This application will run as Go Service (daemon).
The aim of actionpinger is to perform certain actions on an internal network

We are not trying to replace nagios or any monitoring platform,
Actionpinger is focused for use on personal Desktop.
You can find this useful if you want to activate/deactivate certain services on untrusted network,
or according to the output of external webservices.


TODO
----

- [x] Add logging
- [ ] Implement in Goroutine
- [ ] Daemonize https://github.com/takama/daemon / https://github.com/sevlyar/go-daemon
