package main

import "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth = 1000
	screnHeight = 480
)
var (
	running = true
	bkgColor = rl.NewColor(147, 211, 196, 255)

	grassSprite rl.Texture2D
	playerSprite rl.Texture2D

	playerSrc rl.Rectangle
	playerDest rl. rectangle

	playerSpeed float32 = 3
)

func drawscene() {
	rl.DrawTexture(grassSprite, 100, 50, rl.White)
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White)
}

func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		playerDest.Y -= playerSpeed
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		playerDest.Y += playerSpeed
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		playerDest.X -= playerSpeed
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		playerDest.X += playerSpeed
	}
}
func update() {
	running = !rl.WindowShouldClose()
}
func render() {
		rl.BeginDrawing()
		rl.ClearBackground(bkgColor)

		drawScene()

		rl.EndDrawing()
}

func init() {
		rl.InitWindow(screenWidth, screenHeight, "Sproutlings")
		rl.SetExitKey(0)
		rl.SetTargetFPS(60)

		grassSprite = rl.LoadTexture("res/Tilesets/Grass.png")
		playerSprite = rl.LoadTexture("res/Character/BasicCharakterSpritesheet.png")

		playerSrc = rl.NewRectangle(0, 0, 48, 48)
		playerDest = rl.NewRectangle(200, 200, 100, 100)
}
func quit() {
	rl.UnloadTexture(GrassSprite)
	rl.Unloadtexture(PlayerSprite)
	rl.CloseWindow()
}
func main() {
	

	for running() {
		input()
		update()
		render()
		}
	quit()
}
