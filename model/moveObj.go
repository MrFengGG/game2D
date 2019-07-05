package model

import(
	"game2D/resource"
)

type Direction int
const (
    UP        Direction = 0 	// 摄像机移动状态:上
    DOWN      Direction = 1     // 下
    LEFT      Direction = 2     // 左
    RIGHT     Direction = 3     // 右
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
func(moveObj *MoveObj) Move(direction Direction, delta float32){
	if(!moveObj.stockDown && moveObj.y + moveObj.size[1] < moveObj.gameMap.Height){
		if(direction == DOWN){
			moveObj.y += moveObj.flySpeed * delta
		}
	}
	if(direction == UP){
		if(!moveObj.stockUp && moveObj.y > 0){
			moveObj.y -= moveObj.flySpeed * delta
		}
	}
	if(direction == LEFT){
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
			moveObj.x -= moveObj.movementSpeed * delta
		}
	}
	if(direction == RIGHT){
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
			moveObj.x += moveObj.movementSpeed * delta
		}
	}
}
//被动的运动,下坠等
func(moveObj *MoveObj) MoveBy(delta float32){
	if(!moveObj.stockDown && moveObj.y + moveObj.size[1] < moveObj.gameMap.Height){
		moveObj.y += moveObj.fallSpeed * delta
	}
}