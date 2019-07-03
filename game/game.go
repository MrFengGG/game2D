package game
import(
	"game2D/resource"
	"game2D/sprite"
	"game2D/camera"
	"game2D/gamemap"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type GameState int
const(
	GAME_ACTIVE GameState = 0
	GAME_MENU   GameState = 1
)
type Game struct{
	//游戏状态
	state GameState
	width, height int32
	renderer *sprite.SpriteRenderer
	gameMap *gamemap.GameMap
	camera *camera.Camera2D
	Keys [1024]bool
}
func NewGame(width,height int32) *Game{
	game := Game{width:width, height:height,state:GAME_ACTIVE,camera:camera.NewDefaultCamera(2000000,10000,width,height)}
	return &game
}
func (game *Game) Init(){
	resource.LoadShader("./glsl/shader.vs", "./glsl/shader.fs", "sprite")
	projection := mgl32.Ortho(0, float32(game.width),float32(game.height),0, -1, 1)
	shader := resource.GetShader("sprite")
	shader.Use()
	shader.SetInt("image", 0)
	shader.SetMatrix4fv("projection", &projection[0])
	game.renderer = sprite.NewSpriteRenderer(shader)
	resource.LoadTexture(gl.TEXTURE0,"./image/stone.png","stone")
	resource.LoadTexture(gl.TEXTURE0,"./image/soil.png","soil")
	resource.LoadTexture(gl.TEXTURE0,"./image/man-stand.png","man-stand")
	game.gameMap = gamemap.NewGameMap(2000000,10000,"123")
}
func (game *Game) ProcessInput(delta float64){
	if(game.state == GAME_ACTIVE){
		if(game.Keys[glfw.KeyA]){
			game.camera.ProcessKeyboard(camera.LEFT,float32(delta))
		}
		if(game.Keys[glfw.KeyD]){
			game.camera.ProcessKeyboard(camera.RIGHT,float32(delta))
		}
		if(game.Keys[glfw.KeyW]){
			game.camera.ProcessKeyboard(camera.UP,float32(delta))
		}
		if(game.Keys[glfw.KeyS]){
			game.camera.ProcessKeyboard(camera.DOWN,float32(delta))
		}
	}
}
func (game *Game) Update(delta float64){

}
func (game *Game) Render(){
	resource.GetShader("sprite").SetMatrix4fv("view",game.camera.GetViewMatrix())
	game.gameMap.Draw(game.camera.GetPosition(),mgl32.Vec2{float32(game.width),float32(game.height)},game.renderer)
}