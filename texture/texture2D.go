package texture
import(
	"os"
	"image"
	"image/png"
	"image/draw"
	"errors"
	"github.com/go-gl/gl/v4.1-core/gl"
	"fmt"
)

type Texture2D struct{
	ID uint32
	TEXTUREINDEX uint32
}
func NewTexture2D(file string, TEXTUREINDEX uint32) *Texture2D{
	imgFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	img, err := png.Decode(imgFile)
	if err != nil {
		panic(err)
	}
	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		panic(errors.New("unsupported stride"))
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
	
	var textureID uint32
	fmt.Println("hahahah")
	gl.GenTextures(1, &textureID)
	fmt.Println(textureID)
	gl.BindTexture(gl.TEXTURE_2D, textureID)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))
	gl.BindTexture(gl.TEXTURE_2D, 0);
	return &Texture2D{ ID: textureID,TEXTUREINDEX:TEXTUREINDEX}
}

func (texture *Texture2D) Use(){
	gl.ActiveTexture(texture.TEXTUREINDEX)
	gl.BindTexture(gl.TEXTURE_2D, texture.ID)
}