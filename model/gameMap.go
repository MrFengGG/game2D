package model
import(
	"game2D/resource"
	"game2D/sprite"
	"math"
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
)

type GameMap struct{
	Height float32
	Width float32
	blocks [][] *Block
	heightBlockNum int
	widthBlockNum int
	
}
func NewGameMap(height,width float32, mapFile string) *GameMap{
	heightBlockNum := int(math.Ceil(float64(height / BlockHeight)))
	widthBlockNum := int(math.Ceil(float64(width / BlockWidth)))
	grounds := heightBlockNum / 3 * 2
	fmt.Println("size:",heightBlockNum,widthBlockNum)
	blocks := make([][]*Block,heightBlockNum)
	for i := 0;i < heightBlockNum;i++{
		rowBlocks := make([]*Block,widthBlockNum)
		if(i > grounds){
			for j := 0; j < widthBlockNum;j++{
				gameObj := NewGameObj(resource.GetTexture("soil"),float32(j) * BlockWidth,float32(i)*BlockHeight,&mgl32.Vec2{BlockWidth,BlockHeight},0,&mgl32.Vec3{1,1,1})
				rowBlocks[j] = &Block{GameObj:*gameObj}
			}
		}
		blocks[i] = rowBlocks
	}
	return &GameMap{Height:height,Width:width,blocks:blocks,heightBlockNum:heightBlockNum,widthBlockNum:widthBlockNum}
}
func (gameMap *GameMap) CheckObjectStock(moveObj *MoveObj)(bool,bool,bool,bool){
	xPos := x / gameMap.Width * float32(gameMap.widthBlockNum)
	yPos := y / gameMap.Height * float32(gameMap.heightBlockNum)

}
func (gameMap *GameMap) checkObjectStock(moveObj *MoveObj,block *Block)(int,int,int,int){
	objX,objY := block.GetPosition()
	upStoke := x > block.
}
func (gameMap *GameMap) Draw(position mgl32.Vec2, zoom mgl32.Vec2, renderer *sprite.SpriteRenderer){
	startY := int(math.Floor(float64((position[0]) / gameMap.Width * float32(gameMap.widthBlockNum))) - 1);
	if(startY <= 0){
		startY = 0
	}
	endY := int(math.Ceil(float64((position[0] + zoom[0]) / gameMap.Width * float32(gameMap.widthBlockNum))))
	if(endY >= gameMap.widthBlockNum){
		endY = gameMap.widthBlockNum - 1
	}
	startX := int(math.Floor(float64((position[1]) / gameMap.Height * float32(gameMap.heightBlockNum))))
	if(startX < 0){
		startX = 0
	}
	endX := int(math.Ceil(float64((position[1] + zoom[1]) / gameMap.Height * float32(gameMap.heightBlockNum))))
	if(endX >= gameMap.heightBlockNum){
		endX = gameMap.heightBlockNum - 1
	}
	for i:=startX;i<=endX;i++{
		for j := startY;j<endY;j++{
			block := gameMap.blocks[int(i)][int(j)]
			if(block != nil){
				block.Draw(renderer)
			}
		}
	}
}
func (gameMap GameMap) InOut(gameObj GameObj) bool{
	x, y := gameObj.GetPosition()
	height, width := gameObj.GetSize()
	harfHeight := height / 2;
	harfWidth := width / 2
	return x + harfWidth / 2 > gameMap.Width || x - harfWidth < 0 || y + harfHeight > gameMap.Height || y - harfHeight < 0
} 