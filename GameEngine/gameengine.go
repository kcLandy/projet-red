package main

import "fmt"

type GameEngine struct {
	ScreenWidth int
	ScreenHeight int
}

func (g *GameEngine) PrintScreenSize() {
	fmt.Println(g.ScreenWidth)
	fmt.Println(g.ScreenHeight)
}

func (g *GameEngine) InitGameEngine() {
	g.ScreenWidth = 800
	g.ScreenHeight = 600
}