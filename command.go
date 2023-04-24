package main

import (
	"fmt"
	"strconv"
)

func handleLibCommands(token []string) {
	switch token[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)

		}
	case "add":
		if len(token) == 6 {
			id++
			lib.Add(&Music{
				Id:     id,
				Name:   token[2],
				Artist: token[3],
				Source: token[4],
				Type:   token[5],
			})
		} else {
			fmt.Println("USAGE: lib add <name><artist><source><type>")
		}
	case "remove":
		if len(token) == 3 {
			removeByid, _ := strconv.Atoi(token[2])
			lib.Remove(removeByid)
		} else {
			fmt.Println("USAGE: lib remove <name>")
		}
	default:
		fmt.Println("Unrecognized lib command:", token[1])
	}
}

func handlePlayCommand(token []string) {
	if len(token) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}

	e := lib.Find(token[1])
	if e == nil {
		fmt.Println("The music", token[1], "does not exist.")
		return
	}

	Play(e.Source, e.Type)
}
