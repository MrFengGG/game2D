package model

import(
	"game2D/resource"
	"game2D/constant"
	"github.com/go-gl/mathgl/mgl32"
)

//可移动的游戏对象
type MoveObj struct{
	GameObj
	//在上下左右方向是否可移动
	stockUp,stockDown,stockLeft,stockRight bool
	//水平移动速度
	movementSpeed float32
	//飞行速度
	fallSpeed float32
	//下坠速度
	flySpeed float32
	//移动时的动画纹理
	moveTextures []*resource.Texture2D
	//游戏地图
	gameMap  *GameMap
	//当前运动帧
	moveIndex int
	//运动帧之间的切换阈值
	moveDelta float32
}

func NewMoveObject(gameObj GameObj,movementSpeed,flySpeed float32, moveTextures []*resource.Texture2D,gameMap *GameMap) *MoveObj{
	moveObj := &MoveObj{GameObj:gameObj,movementSpeed:movementSpeed,fallSpeed:100,gameMap:gameMap,moveTextures:moveTextures,flySpeed:flySpeed,moveIndex:0,moveDelta:0}
	return moveObj
}

func (moveObj *MoveObj) CheckStock(){
	
}
//由用户主动发起的运动
func(moveObj *MoveObj) Move(direction constant.Direction, delta float32){
	shift := mgl32.Vec2{0,0}
	if(direction ==constant. DOWN){
		if(!moveObj.stockDown && moveObj.y + moveObj.size[1] < moveObj.gameMap.Height){
			shift[1] += moveObj.flySpeed * delta
		}
	}
	if(direction == constant.UP){
		if(!moveObj.stockUp && moveObj.y > 0){
			shift[1] -= moveObj.flySpeed * delta
		}
	}
	if(direction == constant.LEFT){
		moveObj.ReverseX()
		if(moveObj.moveIndex >= len(moveObj.moveTextures)){
			moveObj.moveIndex = 0
		}
		moveObj.moveDelta += delta
		if(moveObj.moveDelta > 0.1){
			moveObj.moveDelta = 0
			moveObj.texture = moveObj.moveTextures[moveObj.moveIndex]
			moveObj.moveIndex += 1
		}
		if(!moveObj.stockLeft && moveObj.x > 0){
			shift[0] -= moveObj.movementSpeed * delta
		}
	}
	if(direction == constant.RIGHT){
		moveObj.ForWardX()
		if(moveObj.moveIndex >= len(moveObj.moveTextures)){
			moveObj.moveIndex = 0
		}
		moveObj.moveDelta += delta
		if(moveObj.moveDelta > 0.1){
			moveObj.moveDelta = 0
			moveObj.texture = moveObj.moveTextures[moveObj.moveIndex]
			moveObj.moveIndex += 1
		}
		if(!moveObj.stockRight && moveObj.x + moveObj.size[0] < moveObj.gameMap.Width){
			shift[0] += moveObj.movementSpeed * delta
		}
	}
	isCol,position := moveObj.gameMap.IsColl(moveObj.GameObj,shift)
	if(isCol){
		moveObj.SetPosition(position)
	}else{
		moveObj.x += shift[0]
		moveObj.y += shift[1]
	}
}
//被动的运动,下坠等
func(moveObj *MoveObj) MoveBy(delta float32){
	if(!moveObj.stockDown && moveObj.y + moveObj.size[1] < moveObj.gameMap.Height){
		moveObj.y += moveObj.fallSpeed * delta
	}
}