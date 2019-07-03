package block
import(
	"game2D/game"
)

type Block struct{
	game.GameObj
}

func (block Block) name() string{
	return "block"
}