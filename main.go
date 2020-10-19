package main

import (
    "fmt"
	"os"
	"os/exec"
	"strings"
	"strconv"
)
func tostring (data []byte) string {
	return string(data[:])
}
func main() {

var c1 ="\033[35m" 
var rs ="\033[0m"

var envars = map[string]string{}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	krn, err :=  exec.Command("uname","-sr").Output()
	if err != nil{
		fmt.Println(err)
	}
	pkg, err :=  exec.Command("pacman","-Q").Output()
	if err != nil{
		fmt.Println(err)
	}

envars["usr"] = os.Getenv("USER")
envars["hos"] = hostname
envars["sh"] = os.Getenv("SHELL")
envars["os"] = "artix"
envars["krn"] = strings.TrimSuffix(tostring(krn),"\n")
envars["pkg"] = strconv.Itoa(strings.Count(tostring(pkg),"\n"))

info := `
%s• usr:%s	%s@%s
%s• os :%s	%s
%s• shl:%s	%s
%s• krn:%s	%s
%s• pkg:%s	%s 

 %s
 %s
`
blocks := "[41m  [m  [42m  [m  [43m  [m  [44m  [m  [45m  [m  [46m  [m "
blocks2 := "[101m  [m  [102m  [m  [103m  [m  [104m  [m  [105m  [m  [105m  [m "

t := fmt.Sprintf(info,
	c1,rs,envars["usr"],envars["hos"],
	c1,rs,envars["os"],
	c1,rs,envars["sh"],
	c1,rs,envars["krn"],
	c1,rs,envars["pkg"],blocks,blocks2)

fmt.Println(t)
}
