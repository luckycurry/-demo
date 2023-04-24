package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var lib *MusicManager
var id int = 1

//var ctrl, signal chan int

func main() {
	fmt.Println(` 
		 Enter following commands to control the player: 
		 lib list -- View the existing music lib 
		 lib add <name><artist><source><type> -- Add a music to the music lib 
		 lib remove <name> -- Remove the specified music from the lib 
		 play <name> -- Play the specified music 
 `)

	lib = NewMusicManager()
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command->")
		rawline, _, _ := r.ReadLine()
		line := string(rawline)

		if line == "q" || line == "e" {
			break
		}

		token := strings.Split(line, " ")
		if token[0] == "lib" {
			handleLibCommands(token)
		} else if token[0] == "play" {
			handlePlayCommand(token)
		} else {
			fmt.Println("Unrecognized command:", token[0])
		}
	}

}
