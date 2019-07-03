package game
import(
	"game2D/resource"
	"game2D/sprite"
	"game2D/camera"
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
	gameMap *[]*GameObj
	camera *camera.Camera2D
	Keys [1024]bool
}
func NewGame(width,height int32) *Game{
	game := Game{width:width, height:height,state:GAME_ACTIVE,camera:camera.NewDefaultCamera(height + 100,width + 100,width,height)}
	return &game
}
func (game *Game) Init(){
	resource.LoadShader("./glsl/shader.vs", "./glsl/shader.fs", "sprite")
	projection := mgl32.Ortho(0, float32(game.width+100),float32(game.height+100),0, -1, 1)
	shader := resource.GetShader("sprite")
	shader.Use()
	shader.SetInt("image", 0)
	shader.SetMatrix4fv("projection", &projection[0])
	game.renderer = sprite.NewSpriteRenderer(shader)

	resource.LoadTexture(gl.TEXTURE0,"./image/stone.png","stone")
	resource.LoadTexture(gl.TEXTURE0,"./image/soil.png","soil")
	game.gameMap = getObj()
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
	for _,gameobj := range *game.gameMap{
		gameobj.Draw(game.renderer)
	}
}

func getObj() *[]*GameObj{
	blockWidth := 20;
	slice1 := make([]*GameObj, 1000)
	for i := 0; i < 1000; i++ {
		slice1[i] = &GameObj{texture:resource.GetTexture("soil"),x:float32(i * blockWidth) ,y:float32(0),size:mgl32.Vec2{20,15},rotate:0,color:mgl32.Vec3{1,1,1}}
	}
	return &slice1
}