package main

import (
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

const (
	windowWidth  = 900
	windowHeight = 900
)

var (
	grid            [30][30]Block = CreateGrid()
	pointImage      *ebiten.Image = nil
	snakeImage      *ebiten.Image = nil
	snakeBodyImage  *ebiten.Image = nil
	length          int           = 0
	playerDirection string        = ""
	endText         string        = ""
	bodyPositions   [][2]int
)

func init() {
	var err error
	if pointImage, _, err = ebitenutil.NewImageFromFile("images/point.png"); err != nil {
		log.Fatal(err)
	}
	if snakeBodyImage, _, err = ebitenutil.NewImageFromFile("images/snakeBody.png"); err != nil {
		log.Fatal(err)
	}
	if snakeImage, _, err = ebitenutil.NewImageFromFile("images/snakeHead.png"); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
	CreateBody(&grid, &bodyPositions, snakeBodyImage, &length)
	SpawnPoints(&grid, pointImage)
	SetDirection(&grid, &playerDirection)
	Move(&grid, playerDirection, snakeImage, &bodyPositions, &length)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{40, 40, 50, 255})
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col].Image != nil {
				grid[row][col].Draw(screen)
			}
		}
	}
	//text.Draw(screen, endText, font, windowWidth/2, windowHeight / 2, color.RGBA{230, 240, 240, 255})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}

func main() {
	SpawnPlayer(&grid, snakeImage)
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Snake")
	ebiten.SetMaxTPS(8)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
