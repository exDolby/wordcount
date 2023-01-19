package main

import (
	"os"
	"fmt"
	"strings"
)
func main() {
	imputArgs := os.Args[1]
	if imputArgs == "" {
		fmt.Println(0)
		return
	} 
	s := strings.Split(strings.ReplaceAll(imputArgs, "\r\n", "\n"), " ")
	fmt.Println(len(s))


}