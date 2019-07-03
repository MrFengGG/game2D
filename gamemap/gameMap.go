package gamemap
import(
	"game2D/resource"
	"game2D/model"
	"game2D/sprite"
	"math"
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
)

type GameMap struct{
	Height float32
	Width float32
	blocks [][] *model.Block
	heightBlockNum int
	widthBlockNum int
	
}
func NewGameMap(height,width float32, mapFile string) *GameMap{
	heightBlockNum := int(math.Ceil(float64(height / model.BlockHeight)))
	widthBlockNum := int(math.Ceil(float64(width / model.BlockWidth)))
	fmt.Println("size:",heightBlockNum,widthBlockNum)
	blocks := make([][]*model.Block,widthBlockNum)
	for i := 0;i < widthBlockNum;i++{
		rowBlocks := make([]*model.Block,heightBlockNum)
		for j := 0; j < heightBlockNum;j++{
			gameObj := model.NewGameObj(resource.GetTexture("soil"),float32(i) * model.BlockWidth,float32(j)*model.BlockHeight,&mgl32.Vec2{20,15},0,&mgl32.Vec3{0,1,1})
			rowBlocks[j] = &model.Block{GameObj:*gameObj}
		}
		blocks[i] = rowBlocks
	}
	return &GameMap{Height:height,Width:width,blocks:blocks,heightBlockNum:heightBlockNum,widthBlockNum:widthBlockNum}
}
func (gameMap *GameMap) Draw(position mgl32.Vec2, zoom mgl32.Vec2, renderer *sprite.SpriteRenderer){
	startX := int(math.Ceil(float64((position[0]) / gameMap.Width * float32(gameMap.widthBlockNum))) - 1);
	if(startX <= 0){
		startX = 0
	}
	endX := int(math.Floor(float64((position[0] + zoom[0]) / gameMap.Width * float32(gameMap.widthBlockNum))))
	if(endX >= gameMap.widthBlockNum){
		endX = gameMap.widthBlockNum - 1
	}
	startY := math.Ceil(float64((position[1]) / gameMap.Height * float32(gameMap.heightBlockNum)))
	endY := math.Floor(float64((position[1] + zoom[1]) / gameMap.Height * float32(gameMap.heightBlockNum)))
	fmt.Println("startX endX",startX,endX)
	for i:=startX;i<=endX;i++{
		for j := startY;j<endY;j++{
			gameMap.blocks[int(i)][int(j)].Draw(renderer)
		}
	}
}
func (gameMap GameMap) InOut(gameObj model.GameObj) bool{
	x, y := gameObj.GetPosition()
	height, width := gameObj.GetSize()
	harfHeight := height / 2;
	harfWidth := width / 2
	return x + harfWidth / 2 > gameMap.Width || x - harfWidth < 0 || y + harfHeight > gameMap.Height || y - harfHeight < 0
}