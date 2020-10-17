package main

import (
    "fmt"
    "os"
    "text/tabwriter"
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



	w := tabwriter.NewWriter(os.Stdout, 10, 1, 1, ' ', 0)
	fmt.Fprintln(w,c1+"usr:" + rs + "  " + envars["usr"]+c1+"@"+rs+ envars["hos"]+"\t")
	fmt.Fprintln(w,c1+"shl:" + rs + "  " + envars["shl"])
	w.Flush()
}
