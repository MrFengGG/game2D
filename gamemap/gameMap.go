package gamemap
import(
	"game2D/game"
)

type GameMap struct{
	Height float32
	Width float32
}
func NewGameMap(height,width float32) *GameMap{
	return &GameMap{Height:height,Width:width}
}
func (gameMap GameMap) InOut(gameObj game.GameObj) bool{
	x, y := gameObj.GetPosition()
	height, width := gameObj.GetSize()
	harfHeight := height / 2;
	harfWidth := width / 2
	return x + harfWidth / 2 > gameMap.Width || x - harfWidth < 0 || y + harfHeight > gameMap.Height || y - harfHeight < 0
}
func (gameMap GameMap) InOut(gameObj game.GameObj) bool{
	x, y := gameObj.GetPosition()
	height, width := gameObj.GetSize()
	harfHeight := height / 2;
	harfWidth := width / 2
	return x + harfWidth / 2 > gameMap.Width || x - harfWidth < 0 || y + harfHeight > gameMap.Height || y - harfHeight < 0
}