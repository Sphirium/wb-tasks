package main

import (
	"fmt"
	"log"
	"os"

	"github.com/beevik/ntp"
)

const defaultNTPServer = "pool.ntp.org"

func main() {
	time, err := ntp.Time(defaultNTPServer)
	if err != nil {
		log.SetOutput(os.Stderr)
		log.Fatalf("Ошибка получения времени от NTP-сервера %s: %v", defaultNTPServer, err)
	}
	fmt.Println(time.Format("2006-01-02 15:04:05"))
}
