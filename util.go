package main

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/text"
    "golang.org/x/image/font"
    "golang.org/x/image/font/basicfont"
    "image/color"
    "math"
)

var standardFont = basicfont.Face7x13

func rainbowColor(frameCount int) color.Color {
    r := math.Sin(float64(frameCount)*0.01 + 0) * 127 + 128
    g := math.Sin(float64(frameCount)*0.01 + 2*math.Pi/3) * 127 + 128
    b := math.Sin(float64(frameCount)*0.01 + 4*math.Pi/3) * 127 + 128
    return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}


func createTextImage(textStr string, clr color.Color, fnt font.Face) *ebiten.Image {
    // Measure the size of the text
    bound, _ := font.BoundString(fnt, textStr)
    w, h := (bound.Max.X - bound.Min.X).Ceil(), (bound.Max.Y - bound.Min.Y).Ceil()

    // Create a new image with the measured size
    img := ebiten.NewImage(w, h)

    // Draw the text onto the image
    text.Draw(img, textStr, fnt, 0, h - int(fnt.Metrics().Descent.Ceil()), clr)

    return img
}

func drawSquare(screen *ebiten.Image, p Point, clr color.Color) {
    rect := ebiten.NewImage(gridSize-1, gridSize-1)
    rect.Fill(clr)
    opts := &ebiten.DrawImageOptions{}
    opts.GeoM.Translate(float64(p.X*gridSize), float64(p.Y*gridSize))
    screen.DrawImage(rect, opts)
}