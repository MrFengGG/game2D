package game
import(
	"game2D/resource"
	"game2D/sprite"
	"game2D/camera"
	"game2D/model"
	"game2D/constant"
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
	//屏幕大小
	screenWidth, screenHeight float32
	//世界大小
	worldWidth, worldHeight float32
	//精灵渲染器
	renderer *sprite.SpriteRenderer
	//游戏地图
	gameMap *model.GameMap
	//摄像头
	camera *camera.Camera2D
	//玩家
	player *model.MoveObj
	//按键状态
	Keys [1024]bool
}
func NewGame(screenWidth, screenHeight, wordWidth, wordHeight float32) *Game{
	game := Game{screenWidth:screenWidth, 
				 screenHeight:screenHeight,
				 worldWidth:wordWidth,
				 worldHeight:wordHeight,
				 state:GAME_ACTIVE}
	return &game
}
func (game *Game) Init(){
	//初始化着色器
	resource.LoadShader("./glsl/shader.vs", "./glsl/shader.fs", "sprite")
	shader := resource.GetShader("sprite")
	shader.Use()
	shader.SetInt("image", 0)
	//设置投影
	projection := mgl32.Ortho(0, float32(game.screenWidth),float32(game.screenHeight),0, -1, 1)
	shader.SetMatrix4fv("projection", &projection[0])
	//初始化精灵渲染器
	game.renderer = sprite.NewSpriteRenderer(shader)
	//加载资源
	resource.LoadTexture(gl.TEXTURE0,"./image/stone.png","stone")
	resource.LoadTexture(gl.TEXTURE0,"./image/soil.png","soil")
	resource.LoadTexture(gl.TEXTURE0,"./image/man-stand.png","man-stand")
	resource.LoadTexture(gl.TEXTURE0,"./image/1.png","1")
	resource.LoadTexture(gl.TEXTURE0,"./image/2.png","2")
	resource.LoadTexture(gl.TEXTURE0,"./image/3.png","3")
	resource.LoadTexture(gl.TEXTURE0,"./image/4.png","4")
	resource.LoadTexture(gl.TEXTURE0,"./image/5.png","5")
	resource.LoadTexture(gl.TEXTURE0,"./image/6.png","6")
	//创建游戏地图
	game.gameMap = model.NewGameMap(game.worldWidth, game.worldHeight,"testMapFile")
	//创建测试游戏人物
	gameObj := model.NewGameObj(resource.GetTexture("man-stand"),
								game.worldWidth/2,
								game.worldHeight/2,
								&mgl32.Vec2{70,100},
								0,
								&mgl32.Vec3{1,1,1})
	//创建摄像头,将摄像头同步到玩家位置
	game.camera = camera.NewDefaultCamera(game.worldHeight, 
										  game.worldWidth,
										  game.screenWidth,
										  game.screenHeight,
										  mgl32.Vec2{game.worldWidth/2 - game.screenWidth/2, game.worldHeight/2 - game.screenHeight/2})

	game.player = model.NewMoveObject(*gameObj,1000,1000,[]*resource.Texture2D{resource.GetTexture("1"),
																			resource.GetTexture("2"),
																			resource.GetTexture("3"),
																			resource.GetTexture("4"),
																			resource.GetTexture("5"),
																			resource.GetTexture("6"),},game.gameMap)
}
//处理输入
func (game *Game) ProcessInput(delta float64){
	if(game.state == GAME_ACTIVE){
		playerMove := false
		if(game.Keys[glfw.KeyA]){
			playerMove = true
			game.player.Move(constant.LEFT,float32(delta))
		}
		if(game.Keys[glfw.KeyD]){
			playerMove = true
			game.player.Move(constant.RIGHT,float32(delta))
		}
		if(game.Keys[glfw.KeyW]){
			playerMove = true
			game.player.Move(constant.UP,float32(delta))
		}
		if(game.Keys[glfw.KeyS]){
			playerMove = true
			game.player.Move(constant.DOWN,float32(delta))
		}
		if(!playerMove){
			game.player.Stand()
		}
	}
}
func (game *Game) Update(delta float64){

}
//渲染每一帧
func (game *Game) Render(delta float64){
	resource.GetShader("sprite").SetMatrix4fv("view",game.camera.GetViewMatrix())
	//game.player.MoveBy(float32(delta))
	game.player.Draw(game.renderer)
	//摄像头跟随
	position := game.player.GetPosition()
	size := game.player.GetSize()
	game.camera.InPosition(position[0] - game.screenWidth /2 + size[0], position[1] - game.screenHeight / 2 + size[1])
	game.gameMap.Draw(game.camera.GetPosition(),
					  mgl32.Vec2{game.screenWidth,game.screenHeight},
					  game.renderer)
}