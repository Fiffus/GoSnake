package main

import (
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
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
	ttf             *sfnt.Font
	textFont        font.Face
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
	if ttf, err = opentype.Parse(fonts.MPlus1pRegular_ttf); err != nil {
		log.Fatal(err)
	}
	textFont, err = opentype.NewFace(ttf, &opentype.FaceOptions{
		Size:    100.0,
		DPI:     75,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (game *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
	if endText != "You lost!" {
		CreateBody(&grid, &bodyPositions, snakeBodyImage, &length)
		SpawnPoints(&grid, pointImage)
		SetDirection(&grid, &playerDirection)
		Move(&grid, playerDirection, snakeImage, &bodyPositions, &length)
	}
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{40, 40, 50, 255})
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col].Image != nil {
				grid[row][col].Draw(screen)
			}
		}
	}
	if endText != "" {
		text.Draw(screen, endText, textFont, windowWidth/2-9*25, windowHeight/2+25, color.RGBA{220, 255, 255, 255})
	}
	if length+1 == len(grid)*len(grid[0]) {
		endText = "You won!"
		text.Draw(screen, endText, textFont, windowWidth/2-8*25, windowHeight/2+25, color.RGBA{220, 255, 255, 255})
	}
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
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
