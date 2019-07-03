package camera
import(
	"github.com/go-gl/mathgl/mgl32"
)
type Direction int
const (
    UP        Direction = 0 	// 摄像机移动状态:上
    DOWN      Direction = 1     // 下
    LEFT      Direction = 2     // 左
    RIGHT     Direction = 3     // 右
)
type Camera2D struct{
	position    mgl32.Vec3
	front       mgl32.Vec3
	up          mgl32.Vec3
	right       mgl32.Vec3
	movementSpeed float32
}
func NewDefaultCamera() *Camera2D{
	position := mgl32.Vec3{0, 0, 0}
	front    := mgl32.Vec3{0, 0, -1}
	up		 := mgl32.Vec3{0, 1, 0}
	right    := mgl32.Vec3{1, 0, 0}
	movementSpeed := float32(2.5)
	return &Camera2D{position:position, front:front, up:up, right:right, movementSpeed:movementSpeed}
}
//获取view
func (camera *Camera2D) GetViewMatrix() *float32{
	target := camera.position.Add(camera.front)
	view := mgl32.LookAtV(camera.position,target, camera.up)
	return &view[0]
}
//键盘回调
func (camera *Camera2D) ProcessKeyboard(direction Direction, deltaTime float32){
	velocity := camera.movementSpeed * deltaTime;
	if (direction == UP){
		camera.position = camera.position.Add(camera.up.Mul(velocity))
	}
	if (direction == DOWN){
		camera.position = camera.position.Sub(camera.up.Mul(velocity))
	}
	if (direction == LEFT){
		camera.position = camera.position.Sub(camera.right.Mul(velocity))
	}
	if (direction == RIGHT){
		camera.position = camera.position.Add(camera.right.Mul(velocity))
	}
}