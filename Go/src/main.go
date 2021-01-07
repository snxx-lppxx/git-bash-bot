package main

import (
	"os"
	"sys"
	"io"
	"fmt"
)

func main() {
	add := 12
	fmt.Println("%d", add)
}
func osOn() {
	intChan := make(chan int, 10)
	intChan <- 1
	fmt.Println(<- intChan)
}
func sysOn() {

}
func ioOn() {

}
func 












//-----------------------------------------------
/*"git"
	"init", "version", "clone",
"git-remote-add-origin"

//---- Function for push 
[["add"],
	["commit"],
["push"]]
//----==================

["ls", "dir"]
"mkdir"
"rm"

//---- Creating new files
["nano",
"nul>"]
//----===================

["ren", "rename"] 
*/