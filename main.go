package main

import "log"

const tcpPort = "9000"

func main() {
	log.SetFlags(0)
	log.Printf(bannerFormat, tcpPort)

	server, err := NewServer(":" + tcpPort)
	if err != nil {
		log.Fatalf(listenFailedFormat, err)
	}

	log.Printf(listeningFormat, tcpPort)
	server.Run()
}
