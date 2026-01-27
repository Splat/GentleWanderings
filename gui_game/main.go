package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	_ "image/png" // This is used to decode PNG images"
	"log"
	"os"
)

type Game struct {
	buttons []*Button
	sprite  *ebiten.Image
	state   GameState
}

// Define a set of game states
type GameState int

const (
	MainMenu GameState = iota
	GamePlay
	Quit
)

type Button struct {
	name     string
	x, y     int
	img      *ebiten.Image
	btnState GameState
}

func NewButton(name string, x, y, w, h int, clr color.Color, btnState GameState) *Button {
	img := ebiten.NewImage(w, h)
	img.Fill(clr)
	return &Button{
		name:     name,
		x:        x,
		y:        y,
		img:      img,
		btnState: btnState,
	}
}

func (b *Button) clicked(mx, my int) bool {
	w, h := b.img.Size()
	return mx >= b.x && mx < b.x+w && my >= b.y && my < b.y+h
}

func (b *Button) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.x), float64(b.y))
	screen.DrawImage(b.img, op)
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		for _, btn := range g.buttons {
			if btn.clicked(x, y) {
				g.state = btn.btnState
				break
			}
		}
	}

	// If the state is Quit, end the game
	if g.state == Quit {
		return fmt.Errorf("game ended by user")
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case MainMenu:
		// draw the buttons
		for _, btn := range g.buttons {
			btn.Draw(screen)
		}

	case GamePlay:
		// draw the game here
		ebiten.SetWindowTitle("Sprites Example")
		// Here, you'd draw the sprite onto the screen at some position depending on your game's state
		// (x, y) could represent the player's position on the map.
		op := &ebiten.DrawImageOptions{}
		x, y := 50.0, 50.0
		op.GeoM.Translate(x, y)
		screen.DrawImage(g.sprite, op)
	}
}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	// Open the sprite image file
	file, err := os.Open("./gui_game/sprite.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	game := &Game{
		buttons: []*Button{
			NewButton("New", 50, 50, 120, 60, color.RGBA{0, 150, 0, 255}, GamePlay),
			NewButton("Load", 50, 150, 120, 60, color.RGBA{0, 0, 150, 255}, GamePlay),
			NewButton("Options", 50, 250, 120, 60, color.RGBA{150, 0, 0, 255}, MainMenu),
			NewButton("Quit", 50, 350, 120, 60, color.RGBA{150, 0, 0, 255}, Quit),
		},
		sprite: ebiten.NewImageFromImage(img), // Create an ebiten.Image from the standard image.Image
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("My Game Menu")
	if err := ebiten.RunGame(game); err != nil && err.Error() != "game ended by user" {
		panic(err)
	}
}
