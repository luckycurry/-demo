package main

import (
	"errors"
	"fmt"
)

// NewMusicManager 获取操作切片
func NewMusicManager() *MusicManager {
	return &MusicManager{make([]Music, 0)}
}

// Len 获取切片的长度
func (m *MusicManager) Len() int {
	return len(m.Musics)
}

// Get 获取曲目的基本消息
func (m *MusicManager) Get(index int) (Music *Music, err error) {
	if index > m.Len() || index < 0 {
		return nil, errors.New("曲库不存在此Id的曲目")
	}
	return &m.Musics[index], nil
}

// Find 通过名字寻找歌曲信息
func (m *MusicManager) Find(name string) *Music {
	if m.Len() == 0 {
		return nil
	}

	for _, v := range m.Musics {
		if v.Name == name {
			return &v
		}
	}
	return nil
}

func (m *MusicManager) Add(Music *Music) {
	m.Musics = append(m.Musics, *Music)
}

func (m *MusicManager) Remove(index int) *Music {
	if index < 0 || index > m.Len() {
		fmt.Println("The index is not in scope")
		return nil
	}

	removeMusic := &m.Musics[index]

	if index < m.Len()-1 {
		m.Musics = append(m.Musics[0:index-1], m.Musics[index+1])
	} else if index == 0 {
		m.Musics = make([]Music, 0)
	} else {
		m.Musics = m.Musics[:index-1]
	}
	return removeMusic
}
