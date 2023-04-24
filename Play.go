package main

import "fmt"

type Player interface {
	Play(source string)
}

func Play(source string, mtype string) {
	var p Player

	switch mtype {
	case "MP3":
		p = &MP3Player{}
	//case "WAV"
	//	p = &WVAPlayer{}
	default:
		fmt.Println("没有与歌曲类型相匹配的播放器")
		return
	}
	p.Play(source)
}
