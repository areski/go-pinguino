package main

import (
	"github.com/areski/go-pinguino"
)

func main() {
	pinguino.LoadConfig(pinguino.Prod_conf)
	pinguino.StartDaemon()
}
