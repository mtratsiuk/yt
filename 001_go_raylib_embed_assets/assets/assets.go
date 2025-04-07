// Generated code - do not edit manually!

package assets

import (
	_ "embed"
	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed smiley_face.png
var SmileyFace []byte

type Assets struct {
	SmileyFace rl.Texture2D
}

func NewAssets() Assets {
	a := Assets{}
	a.SmileyFace = loadTexture(SmileyFace, ".png")

	return a
}

func (a *Assets) Close() {
	rl.UnloadTexture(a.SmileyFace)
}

func loadTexture(image []byte, ext string) rl.Texture2D {
	img := rl.LoadImageFromMemory(ext, image, int32(len(image)))
	defer rl.UnloadImage(img)

	return rl.LoadTextureFromImage(img)
}
