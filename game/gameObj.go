package game
import(
	"game2D/texture"
	"game2D/sprite"
	"github.com/go-gl/mathgl/mgl32"
)
type GameObj struct{
	texture *texture.Texture2D
	x float32
	y float32
	size mgl32.Vec2
	rotate float32
	color mgl32.Vec3
}

func(gameObj *GameObj) Draw(renderer *sprite.SpriteRenderer){
	renderer.DrawSprite(gameObj.texture, mgl32.Vec2{gameObj.x,gameObj.y}, gameObj.size, gameObj.rotate, gameObj.color)
}