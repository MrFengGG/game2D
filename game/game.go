package game
import(
	"game2D/resource"
	"game2D/sprite"
	"game2D/camera"
	"game2D/model"
	"game2D/texture"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type GameState int
const(
	GAME_ACTIVE GameState = 0
	GAME_MENU   GameState = 1
	WORD_WIDTH  float32     = 1500
	WORD_HEIGHT float32     = 1000
)
type Game struct{
	//游戏状态
	state GameState
	width, height float32
	renderer *sprite.SpriteRenderer
	gameMap *model.GameMap
	camera *camera.Camera2D
	player *model.MoveObj
	Keys [1024]bool
}
func NewGame(width,height float32) *Game{
	game := Game{width:width, 
				height:height,
				state:GAME_ACTIVE,
				camera:camera.NewDefaultCamera(WORD_WIDTH,WORD_HEIGHT,width,height)}
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
	resource.LoadTexture(gl.TEXTURE0,"./image/1.png","1")
	resource.LoadTexture(gl.TEXTURE0,"./image/2.png","2")
	resource.LoadTexture(gl.TEXTURE0,"./image/3.png","3")
	resource.LoadTexture(gl.TEXTURE0,"./image/4.png","4")
	resource.LoadTexture(gl.TEXTURE0,"./image/5.png","5")
	resource.LoadTexture(gl.TEXTURE0,"./image/6.png","6")
	game.gameMap = model.NewGameMap(WORD_WIDTH,WORD_HEIGHT,"123")
	gameObj := model.NewGameObj(resource.GetTexture("man-stand"),WORD_WIDTH/2,WORD_HEIGHT/2,&mgl32.Vec2{70,100},0,&mgl32.Vec3{1,1,1})
	game.player = model.NewMoveObject(*gameObj,100,30,[]*texture.Texture2D{resource.GetTexture("1"),
																			resource.GetTexture("2"),
																			resource.GetTexture("3"),
																			resource.GetTexture("4"),
																			resource.GetTexture("5"),
																			resource.GetTexture("6"),},game.gameMap)
}
func (game *Game) ProcessInput(delta float64){
	if(game.state == GAME_ACTIVE){
		if(game.Keys[glfw.KeyA]){
			game.camera.ProcessKeyboard(camera.LEFT,float32(delta))
			game.player.Move(camera.LEFT,float32(delta))
		}
		if(game.Keys[glfw.KeyD]){
			game.camera.ProcessKeyboard(camera.RIGHT,float32(delta))
			game.player.Move(camera.RIGHT,float32(delta))
		}
		if(game.Keys[glfw.KeyW]){
			game.camera.ProcessKeyboard(camera.UP,float32(delta))
			game.player.Move(camera.UP,float32(delta))
		}
		if(game.Keys[glfw.KeyS]){
			game.camera.ProcessKeyboard(camera.DOWN,float32(delta))
			game.player.Move(camera.DOWN,float32(delta))
		}
	}
}
func (game *Game) Update(delta float64){

}
func (game *Game) Render(delta float64){
	resource.GetShader("sprite").SetMatrix4fv("view",game.camera.GetViewMatrix())
	game.player.MoveBy(float32(delta))
	game.player.Draw(game.renderer)
	game.gameMap.Draw(game.camera.GetPosition(),mgl32.Vec2{float32(game.width),float32(game.height)},game.renderer)
}