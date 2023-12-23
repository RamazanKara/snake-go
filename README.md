# Snake Game in Go

This is a simple implementation of the classic Snake game written in Go, using the Ebiten 2D game library. Navigate the snake around the screen, collect food, and avoid colliding with yourself!

## Features

- Classic Snake gameplay
- Score tracking
- Flashing rainbow text effect
- "Game Over" screen with restart capability

## Installation

To run this game, you need to have Go installed on your system. If you don't have Go installed, you can download it from [The Go Programming Language website](https://golang.org/dl/).

Once Go is installed, you can clone this repository using:

\```bash
git clone https://github.com/RamazanKara/snake-go.git
cd snake-go
\```

## Running the Game

To run the game, navigate to the cloned directory and use the Go run command:

\```bash
go run main.go game.go types.go util.go
\```

## Controls

- Use the arrow keys (↑, ↓, ←, →) to change the direction of the snake.
- Press the spacebar to restart the game after a "Game Over".

## Acknowledgments

- Snake game logic and rendering implemented using the [Ebiten](https://ebiten.org/) gaming library for Go.

## License

This project is open-source and available under the [MIT License](LICENSE).
