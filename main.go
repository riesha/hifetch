package main

import (
    "fmt"
	"os"
)

func main() {

var c1 ="\033[35m" 
var rs ="\033[0m"

var envars = map[string]string{}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}

envars["usr"] = os.Getenv("USER")
envars["hos"] = hostname
envars["shl"] = os.Getenv("SHELL")

info := `
	%susr:%s	%s@%s
	%sshl:%s	%s
	`
t := fmt.Sprintf(info,c1,rs,envars["usr"],envars["hos"],c1,rs,envars["shl"])
fmt.Println(t)
}
