package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	home, _ := os.UserHomeDir()
	LOGFILE := path.Join(home, "mGo.log")
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	iLog := log.New(f, "iLog ", log.LstdFlags)
	iLog.Println("Hello there!")
	iLog.Println("General Kenobi...")
}
