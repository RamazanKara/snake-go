package main

import (
    "fmt"
    "github.com/hajimehoshi/ebiten/v2"
    "image/color"
    "math/rand"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
    screenWidth  = 800
    screenHeight = 600
    gridSize     = 20
)

// Game struct and its methods
type Game struct {
    snake    []Point
    dir      Point
    food     Point
    updating bool
	tickCount int
	nextDir    Point
	score      int
	colorTickCount int
	gameOver bool

}

// Update progresses the game state by one tick. It's called on every frame.
func (g *Game) Update() error {

	if g.gameOver {
        if ebiten.IsKeyPressed(ebiten.KeySpace) { // Reset on spacebar press
            g.reset()
            g.gameOver = false
        }
        return nil
    }
	if g.updating {
        return nil
    }
    g.updating = true
    defer func() { g.updating = false }()
    // Handle input
    if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.dir.X != 1 {
        g.nextDir = Point{-1, 0}
    } else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.dir.X != -1 {
        g.nextDir = Point{1, 0}
    } else if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.dir.Y != 1 {
        g.nextDir = Point{0, -1}
    } else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && g.dir.Y != -1 {
        g.nextDir = Point{0, 1}
    }
	g.colorTickCount++
    // Update game state based on tick count
    g.tickCount++
    if g.tickCount < 10 { // Adjust this number to control speed
        return nil
    }
    g.tickCount = 0
    g.dir = g.nextDir // Update the direction here

    head := g.snake[len(g.snake)-1]
    newHead := Point{head.X + g.dir.X, head.Y + g.dir.Y}

    if newHead == g.food {
        g.snake = append(g.snake, newHead)
        g.spawnFood()
		g.score++
    } else {
        g.snake = append(g.snake[1:], newHead)
    }

	if g.isGameOver() {
        return nil
    }
    return nil
}

// Draw renders the game state to the screen. It's called on every frame after Update.
func (g *Game) Draw(screen *ebiten.Image) {
	if g.gameOver {
        g.drawGameOverScreen(screen)
    } else {
    screen.Fill(color.RGBA{0, 0, 0, 255})

    for _, p := range g.snake {
        drawSquare(screen, p, color.RGBA{0, 255, 0, 255})
    }

    drawSquare(screen, g.food, color.RGBA{255, 0, 0, 255})

	if g.gameOver {
        g.drawGameOverScreen(screen)
    }

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", g.score))

    textStr := "Game created by RamazanKara"
    currentColor := rainbowColor(g.colorTickCount) // Get the current color
    txtImg := createTextImage(textStr, currentColor, standardFont)

    // Get the size of the text image
    txtWidth, txtHeight := txtImg.Size()

    // Position the text in the bottom right corner
    xPosition := screenWidth - txtWidth - 10 // 10 pixels padding from the right edge
    yPosition := screenHeight - txtHeight - 10 // 10 pixels padding from the bottom edge

    textOpts := &ebiten.DrawImageOptions{}
    textOpts.GeoM.Translate(float64(xPosition), float64(yPosition))
    screen.DrawImage(txtImg, textOpts)
}
}

// Layout defines the layout for the game. It's called when the game starts.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return screenWidth, screenHeight
}

func (g *Game) isGameOver() bool {
    head := g.snake[len(g.snake)-1]
    if head.X < 0 || head.Y < 0 || head.X >= screenWidth/gridSize || head.Y >= screenHeight/gridSize {
        g.gameOver = true
        return true
    }
    for _, p := range g.snake[:len(g.snake)-1] {
        if p == head {
            g.gameOver = true
            return true
        }
    }
    return false
}

func (g *Game) spawnFood() {
    g.food = Point{rand.Intn(screenWidth / gridSize), rand.Intn(screenHeight / gridSize)}
}

func (g *Game) reset() {
    g.snake = []Point{{X: screenWidth / gridSize / 2, Y: screenHeight / gridSize / 2}}
    g.dir = Point{0, 0}
    g.nextDir = Point{0, 0}
    g.spawnFood()
    g.score = 0 // Reset the score
}

func (g *Game) drawGameOverScreen(screen *ebiten.Image) {
    gameOverText := "GAME OVER :("
    fnt := standardFont // Or another large font if you have
    clr := color.RGBA{255, 0, 0, 255} // Red color for visibility

    // Create an image with the game over text
    txtImg := createTextImage(gameOverText, clr, fnt)

    // Center the text image on the screen
    txtWidth, txtHeight := txtImg.Size()
    xPosition := (screenWidth - txtWidth) / 2
    yPosition := (screenHeight - txtHeight) / 2

    opts := &ebiten.DrawImageOptions{}
    opts.GeoM.Translate(float64(xPosition), float64(yPosition))
    screen.DrawImage(txtImg, opts)
}