package main

import (
	"fmt"
	"time"
)

func (p *MP3Player) Play(source string) {
	fmt.Println("正在播放", source)

	p.process = 0

	for p.process < 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(".")
		p.process += 10
	}
	fmt.Println("播放结束")
}
