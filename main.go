package main

import (
	"log"

	"clinics-apis/conf"
	"clinics-apis/srv"
)

func main() {

	if err := srv.Start(conf.Config.Addr); err != nil {
		log.Println("start: ", err)
	}
}
