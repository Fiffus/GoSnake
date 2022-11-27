package main

import (
	"fmt"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Position struct {
	X int
	Y int
}

type Size struct {
	width  int
	height int
}

type Block struct {
	position  Position
	size      Size
	Image     *ebiten.Image
	ImageName string
}

func (block *Block) calculateScale() [2]float64 {
	return [2]float64{float64(block.size.width / block.Image.Bounds().Max.X), float64(block.size.height / block.Image.Bounds().Max.Y)}
}

func CreateGrid() [30][30]Block {
	var grid [30][30]Block
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			grid[row][col] = Block{
				position: Position{
					X: row * 30,
					Y: col * 30,
				},
				size: Size{
					width:  30,
					height: 30,
				},
				Image:     nil,
				ImageName: "void",
			}
		}
	}
	return grid
}

func SpawnPoints(grid *[30][30]Block, pointImage *ebiten.Image) {
	var points int = 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col].ImageName == "point" {
				points++
			}
		}
	}
	if points == 0 {
		var pointPosition [2]int = [2]int{rand.Intn(len(grid)), rand.Intn(len(grid[0]))}
		if grid[pointPosition[0]][pointPosition[1]].Image == nil {
			grid[pointPosition[0]][pointPosition[1]].Image = pointImage
			grid[pointPosition[0]][pointPosition[1]].ImageName = "point"
		}
	}
}

func randomInt(number1 int, number2 int) int {
	return number1 + rand.Intn(number2)
}

func SpawnPlayer(grid *[30][30]Block, snakeImage *ebiten.Image) {
	var playerPosition [2]int = [2]int{randomInt(10, 20), randomInt(10, 20)}
	grid[playerPosition[0]][playerPosition[1]].Image = snakeImage
	grid[playerPosition[0]][playerPosition[1]].ImageName = "head"
}

func (block *Block) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(block.calculateScale()[0], block.calculateScale()[1])
	options.GeoM.Translate(float64(block.position.X), float64(block.position.Y))
	screen.DrawImage(block.Image, options)
}

func SetDirection(grid *[30][30]Block, playerDirection *string) {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		*playerDirection = "Up"
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		*playerDirection = "Down"
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		*playerDirection = "Left"
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		*playerDirection = "Right"
	}
}

func CreateBody(grid *[30][30]Block, bodyPositions *[][2]int, snakeImage *ebiten.Image, length *int) {
	for _, position := range *bodyPositions {
		grid[position[0]][position[1]].ImageName = "body"
		grid[position[0]][position[1]].Image = snakeImage
	}
	if len(*bodyPositions) >= *length {
		if *length > 0 {
			pos := *bodyPositions
			grid[pos[0][0]][pos[0][1]].ImageName = "void"
			grid[pos[0][0]][pos[0][1]].Image = nil
			pos = pos[1:]
			*bodyPositions = pos
		}
	}
}

func Move(grid *[30][30]Block, playerDirection string, snakeImage *ebiten.Image, bodyPositions *[][2]int, length *int) {
	if playerDirection != "" {
		for row := 0; row < len(grid); row++ {
			for col := 0; col < len(grid[row]); col++ {
				if playerDirection == "Left" {
					if grid[row][col].ImageName == "head" {
						if row-1 > -1 {
							if grid[row-1][col].ImageName == "body" {
								fmt.Println("lost")
								return
							}
							if *length < 1 {
								grid[row][col].ImageName = "void"
								grid[row][col].Image = nil
							} else {
								grid[row][col].ImageName = "body"
								grid[row][col].Image = snakeBodyImage
							}
							if grid[row-1][col].ImageName == "point" {
								*length++
							}
							grid[row-1][col].ImageName = "head"
							grid[row-1][col].Image = snakeImage
							if *length > len(*bodyPositions) {
								*bodyPositions = append(*bodyPositions, [2]int{row, col})
							}
						} else {
							grid[row][col].ImageName = "void"
							grid[row][col].Image = snakeImage
							fmt.Println("lost")
						}
						return
					}
				}
				if playerDirection == "Right" {
					if grid[row][col].ImageName == "head" {
						if row+1 < len(grid) {
							if grid[row+1][col].ImageName == "body" {
								fmt.Println("lost")
								return
							}
							if *length < 1 {
								grid[row][col].ImageName = "void"
								grid[row][col].Image = nil
							} else {
								grid[row][col].ImageName = "body"
								grid[row][col].Image = snakeBodyImage
							}
							if grid[row+1][col].ImageName == "point" {
								*length++
							}
							grid[row+1][col].ImageName = "head"
							grid[row+1][col].Image = snakeImage
							if *length > len(*bodyPositions) {
								*bodyPositions = append(*bodyPositions, [2]int{row, col})
							}
						} else {
							grid[row][col].ImageName = "void"
							grid[row][col].Image = snakeImage
							fmt.Println("lost")
						}
						return
					}
				}
				if playerDirection == "Up" {
					if grid[row][col].ImageName == "head" {
						if col-1 > -1 {
							if grid[row][col-1].ImageName == "body" {
								fmt.Println("lost")
								return
							}
							if *length < 1 {
								grid[row][col].ImageName = "void"
								grid[row][col].Image = nil
							} else {
								grid[row][col].ImageName = "body"
								grid[row][col].Image = snakeBodyImage
							}
							if grid[row][col-1].ImageName == "point" {
								*length++
							}
							grid[row][col-1].ImageName = "head"
							grid[row][col-1].Image = snakeImage
							if *length > len(*bodyPositions) {
								*bodyPositions = append(*bodyPositions, [2]int{row, col})
							}
						} else {
							grid[row][col].ImageName = "head"
							grid[row][col].Image = snakeImage
							fmt.Println("lost")
						}
						return
					}
				}
				if playerDirection == "Down" {
					if grid[row][col].ImageName == "head" {
						if col+1 < len(grid[0]) {
							if grid[row][col+1].ImageName == "body" {
								fmt.Println("lost")
								return
							}
							if *length < 1 {
								grid[row][col].ImageName = "void"
								grid[row][col].Image = nil
							} else {
								grid[row][col].ImageName = "body"
								grid[row][col].Image = snakeBodyImage
							}
							if grid[row][col+1].ImageName == "point" {
								*length++
							}
							grid[row][col+1].ImageName = "head"
							grid[row][col+1].Image = snakeImage
							if *length > len(*bodyPositions) {
								*bodyPositions = append(*bodyPositions, [2]int{row, col})
							}
						} else {
							grid[row][col].ImageName = "head"
							grid[row][col].Image = snakeImage
							fmt.Println("lost")
						}
						return
					}
				}
			}
		}
	}
}
