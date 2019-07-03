package model


import(
	"game2D/camera"
)

type Player struct{
	GameObj
	stock bool
	movementSpeed float32
}

func(player *Player) Move(direction camera.Direction){

}