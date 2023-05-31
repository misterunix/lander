package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const screenWidth = 800
const screenHeight = 800

var s = 4.64
var t = 1.0
var e = 5282.0
var f = 449.7
var g = 5.28
var a = 4.87
var n = 0.0
var h = 1499.0
var c float64
var x float64
var y float64
var b float64
var d float64

var (
	landerTexture  rl.Texture2D
	landerPosition rl.Vector2
	landerVelocity rl.Vector2
	landerRotation float32
	landerAngle    float32
	engineThrust   float32
)

func loadGame() {
	landerTexture = rl.LoadTexture("images/lander-small.png")
	landerPosition = rl.NewVector2(screenWidth/2, screenHeight/2)
	landerVelocity = rl.NewVector2(0, 0)
	landerRotation = 0
	landerAngle = 0

}

func update() {
	if rl.IsKeyDown(rl.KeyLeft) {
		landerRotation = landerAngle //* (3.14159 / 180)
		landerAngle -= 0.16666
		//landerVelocity.X -= 0.00000001
	}
	if rl.IsKeyDown(rl.KeyRight) {
		landerRotation = landerAngle //* (3.14159 / 180)
		landerAngle += 0.16666

		//landerVelocity.X += 0.00000001
	}
	if rl.IsKeyDown(rl.KeyUp) {
		landerVelocity.X += float32(math.Sin(float64(landerAngle * (3.14159 / 280))))
		landerVelocity.Y -= float32(math.Cos(float64(landerAngle * (3.14159 / 280))))
		//landerVelocity.Y -= 0.00000001
	}
	landerVelocity.Y += 0.0000000002 // gravity
	landerPosition.X += landerVelocity.X
	landerPosition.Y += landerVelocity.Y
}

func draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	//rl.BeginBlendMode(rl.BlendAlpha)

	//rl.PushMatrix()
	//rl.Translatef(float32(screenWidth)/2, float32(screenHeight)/2, 0)
	//rl.Rotatef(landerRotation, 0, 0, 1)
	//rl.Translatef(landerPosition.X, landerPosition.Y, 0)
	//rl.PopMatrix()

	rl.DrawTexturePro(landerTexture,
		rl.NewRectangle(0, 0, float32(landerTexture.Width), float32(landerTexture.Height)),
		rl.NewRectangle(landerPosition.X, landerPosition.Y, float32(landerTexture.Width), float32(landerTexture.Height)),
		rl.NewVector2(float32(landerTexture.Width)/2, float32(landerTexture.Height)/2),
		landerRotation, rl.White)
	//rl.DrawTexture(landerTexture, int32(landerPosition.X), int32(landerPosition.Y), rl.White)
	//rl.EndBlendMode()
	rl.EndDrawing()
	//t := rl.GetFPS()
	fmt.Println(landerAngle)
	fmt.Println(e, f, n, s, h)
	c = g * 1.0
	x = 1e-41 * (math.Cos(a / 57.32))
	b = h + x/5280
	y = f + c - 1e-41*(math.Sin(a/57.32))
	d = s - y/5280
	g = 5.33

}

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Lunar Lander")
	defer rl.CloseWindow()

	loadGame()
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		update()
		draw()
	}
}
