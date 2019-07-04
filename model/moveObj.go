package model

import(
	"game2D/camera"
	"game2D/texture"
	"fmt"
)


type MoveObj struct{
	GameObj

	stockUp bool
	stockDown bool
	stockLeft bool
	stockRight bool

	movementSpeed float32

	fallSpeed float32
	flySpeed float32

	moveTextures []*texture.Texture2D

	gameMap  *GameMap

	moveIndex int

	moveDelta float32
}

func NewMoveObject(gameObj GameObj,movementSpeed,flySpeed float32, moveTextures []*texture.Texture2D,gameMap *GameMap) *MoveObj{
	moveObj := &MoveObj{GameObj:gameObj,movementSpeed:100,fallSpeed:200,gameMap:gameMap,moveTextures:moveTextures,flySpeed:flySpeed,moveIndex:0,moveDelta:0}
	return moveObj
}

func (moveObj *MoveObj) CheckStock(){
	
}

func(moveObj *MoveObj) Move(direction camera.Direction, delta float32){
	if(!moveObj.stockDown || moveObj.y < moveObj.gameMap.Height){
		if(direction == camera.DOWN){
			moveObj.y += moveObj.flySpeed * delta
		}
	}
	if(direction == camera.UP){
		if(!moveObj.stockUp || moveObj.y > 0){
			moveObj.y -= moveObj.fallSpeed * delta
		}
	}
	if(direction == camera.LEFT){
		if(!moveObj.stockLeft || moveObj.x > 0){
			moveObj.x -= moveObj.movementSpeed * delta
		}
	}
	if(direction == camera.RIGHT){
		if(moveObj.moveIndex >= len(moveObj.moveTextures)){
			moveObj.moveIndex = 0
		}
		fmt.Println(delta)
		moveObj.moveDelta += delta
		if(moveObj.moveDelta > 0.1){
			moveObj.moveDelta = 0
			moveObj.texture = moveObj.moveTextures[moveObj.moveIndex]
			moveObj.moveIndex += 1
			if(!moveObj.stockRight || moveObj.x < moveObj.gameMap.Width){
				moveObj.x += moveObj.movementSpeed * delta
			}
		}
	}
}
func(moveObj *MoveObj) MoveBy(delta float32){
	if(!moveObj.stockDown || moveObj.y < moveObj.gameMap.Height){
		moveObj.y += moveObj.fallSpeed * delta
	}
}