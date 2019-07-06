package model
import(
	"game2D/resource"
	"game2D/sprite"
	"github.com/go-gl/mathgl/mgl32"
)
type GameObj struct{
	texture *resource.Texture2D
	x float32
	y float32
	size *mgl32.Vec2
	rotate float32
	color *mgl32.Vec3
	isXReverse int32
}
func(gameObj GameObj) GetPosition()mgl32.Vec2{
	return mgl32.Vec2{gameObj.x, gameObj.y}
}
func(gameObj *GameObj) SetPosition(position mgl32.Vec2){
	gameObj.x = position[0]
	gameObj.y = position[1]
}
func(gameObj GameObj) GetSize()mgl32.Vec2{
	return mgl32.Vec2{gameObj.size[0], gameObj.size[1]}
}
func(gameObj *GameObj) Draw(renderer *sprite.SpriteRenderer){
	renderer.DrawSprite(gameObj.texture, &mgl32.Vec2{gameObj.x,gameObj.y}, gameObj.size, gameObj.rotate, gameObj.color,gameObj.isXReverse)
}
func(gameObj *GameObj) ReverseX(){
	gameObj.isXReverse = -1
}
func(gameObj *GameObj) ForWardX(){
	gameObj.isXReverse = 1
}
func NewGameObj(texture *resource.Texture2D, x, y float32, size *mgl32.Vec2, rotate float32, color *mgl32.Vec3) *GameObj{
	return &GameObj{texture:texture,
					x:x,
					y:y,
					size:size,
					rotate:rotate,
					color:color,
					isXReverse:1}
}