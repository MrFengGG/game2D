package camera
import(
	"github.com/go-gl/mathgl/mgl32"
)
type Camera2D struct{
	position,front,up     						   mgl32.Vec3
	movementSpeed         						   float32
	wordWidth,wordHeight,screenWidth,screenHeight  float32
}
func NewDefaultCamera(wordHeight ,wordWidth, screenWidth, screenHeight float32, position2D mgl32.Vec2) *Camera2D{
	position := mgl32.Vec3{position2D[0], position2D[1], 0}
	front    := mgl32.Vec3{0, 0, -1}
	up		 := mgl32.Vec3{0, 1, 0}
	movementSpeed := float32(100)

	return &Camera2D{position:position, 
		front:front, 
		up:up, 
		movementSpeed:movementSpeed,
		wordHeight:wordHeight,
		wordWidth:wordWidth,
		screenHeight:screenHeight,
		screenWidth:screenWidth}
}
//获取摄像头位置
func (camera *Camera2D) GetPosition() mgl32.Vec2{
	return mgl32.Vec2{camera.position[0], camera.position[1]}
}
//获取view
func (camera *Camera2D) GetViewMatrix() *float32{
	target := camera.position.Add(camera.front)
	view := mgl32.LookAtV(camera.position,target, camera.up)
	return &view[0]
}
//重置世界边界
func (camera *Camera2D) resetWordSize(width,height float32){
	camera.wordWidth = width
	camera.wordHeight = height
}
//重设屏幕大小
func (camera *Camera2D) resetScreenSize(width,height float32){
	camera.screenWidth = width
	camera.screenHeight = height
}
//根据坐标转换视野
func(camera *Camera2D) InPosition(x,y float32){
	if(x < 0){
		camera.position[0] = 0
	}else if(x + camera.screenWidth > camera.wordWidth){
		x = camera.wordWidth - camera.screenWidth
	}else{
		camera.position[0] = x
	}
	if(y < 0){
		camera.position[1] = 0
	}else if(y + camera.screenHeight < camera.wordHeight){
		camera.position[1] = camera.wordHeight - camera.screenHeight
	}else{
		camera.position[1] = y
	}
}