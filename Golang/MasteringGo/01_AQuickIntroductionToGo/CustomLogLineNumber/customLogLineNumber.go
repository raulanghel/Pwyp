package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	home, _ := os.UserHomeDir()
	LOGFILE := path.Join(home, "mGoN.log")
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	LstdFlags := log.Ldate | log.Lshortfile
	ilog := log.New(f, "LNum", LstdFlags)
	ilog.Println("Hello from another log...")
	ilog.SetFlags(log.Lshortfile | log.LstdFlags)
	ilog.Println("Another log entry...")
}
