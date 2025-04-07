//go:generate go run ./cmd/codegen_assets/main.go -in ./assets -g *.png -out ./assets/assets.go
//go:generate go fmt ./assets

package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/mtratsiuk/yt/001_go_raylib_embed_assets/assets"
)

const W = 640
const H = 360

func main() {
	rl.SetConfigFlags(rl.FlagWindowMinimized | rl.FlagWindowHighdpi)
	rl.InitWindow(W, H, "Assets Embed Example")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	assets := assets.NewAssets()
	defer assets.Close()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		{
			rl.ClearBackground(rl.White)

			rl.DrawTexture(
				assets.SmileyFace,
				W/2-assets.SmileyFace.Width/2, H/2-assets.SmileyFace.Height/2,
				rl.White,
			)
		}
		rl.EndDrawing()
	}
}
