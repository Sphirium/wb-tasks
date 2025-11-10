package main

import "fmt"

// Умеет играть все форматы
type AllPlayer interface {
	TrackPlay(filename string)
}

// Реализация старого проигрывателя
type OldMP3Player struct{}

func (p *OldMP3Player) TrackPlay(filename string) {
	fmt.Printf("Playing MP3: %s\n", filename)

}

// Новый проигрыватель — умеет играть только AAC
type ACCPlayer struct{}

func (a *ACCPlayer) PlayAAC(filename string) {
	fmt.Printf("Playing AAC: %s\n", filename)

}

// Адаптер делает AACPlayer похожим на AllPlayer
type AACAdapter struct {
	aacPlayer *ACCPlayer
}

// Адаптер реализует интерфейс AllPlayer
func (a *AACAdapter) TrackPlay(filename string) {
	a.aacPlayer.PlayAAC(filename)
}

// playMusiс работает только с AllPlayer и может проигрывать любой формат
func playMusic(player AllPlayer, filename string) {
	player.TrackPlay(filename)
}

func main() {
	// Старый проигрыватель — работает напрямую
	oldPlayer := &OldMP3Player{}
	playMusic(oldPlayer, "Eminem - Lose Yourself.mp3")

	// Новый проигрыватель — через адаптер
	aacPlayer := &ACCPlayer{}
	adapter := &AACAdapter{aacPlayer: aacPlayer}
	playMusic(adapter, "Eminem - Lose Yourself.aac")
}
