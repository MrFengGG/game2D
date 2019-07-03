package game
import(
	"game2D/resource"
	"game2D/sprite"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type GameState int
const(
	GAME_ACTIVE GameState = 0
	GAME_MENU   GameState = 1
)
type Game struct{
	//游戏状态
	state GameState
	width, height uint32
	renderer *sprite.SpriteRenderer

}
func NewGame(width,height uint32) *Game{
	game := Game{width:width, height:height}
	return &game
}
func (game *Game) Init(){
	
	resource.LoadShader("./glsl/shader.vs", "./glsl/shader.fs", "sprite")
	projection := mgl32.Ortho(0, float32(game.width), float32(game.height), 0, -1, 1)
	shader := resource.GetShader("sprite")
	shader.Use()
	shader.SetInt("image", 0)
	shader.SetMatrix4fv("projection", &projection[0])
	game.renderer = sprite.NewSpriteRenderer(shader)

	resource.LoadTexture(gl.TEXTURE0,"./image/face.jpg","face")
	resource.LoadTexture(gl.TEXTURE0,"./image/wood.jpg","wood")
	resource.LoadTexture(gl.TEXTURE0,"./image/te.jpg","te")
}
func (game *Game) ProcessInput(delta float64){

}
func (game *Game) Update(delta float64){

}
func (game *Game) Render(){
	gameObj := GameObj{texture:resource.GetTexture("wood"),x:0,y:float32(game.height)-40,size:mgl32.Vec2{30,40},rotate:0,color:mgl32.Vec3{0,1,0}}
	gameObjt := GameObj{texture:resource.GetTexture("te"),x:30,y:float32(game.height)-40,size:mgl32.Vec2{30,40},rotate:0,color:mgl32.Vec3{1,0,0}}
	gameObj.Draw(game.renderer)
	gameObjt.Draw(game.renderer)
}