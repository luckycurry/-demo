package main

type Music struct {
	Id     int
	Name   string
	Artist string
	Source string
	Type   string
}

type MusicManager struct {
	Musics []Music
}

type MP3Player struct {
	stat    int
	process int
}
