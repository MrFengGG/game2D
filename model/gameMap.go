package model
import(
	"game2D/resource"
	"game2D/sprite"
	"game2D/physic"
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
//一个简单的测试用的游戏地图生成函数
func NewGameMap(width,height float32, mapFile string) *GameMap{
	heightBlockNum := int(math.Ceil(float64(height / BlockHeight)))
	widthBlockNum := int(math.Ceil(float64(width / BlockWidth)))
	grounds := heightBlockNum / 3 * 2
	fmt.Println("map block size:",heightBlockNum,widthBlockNum)
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
	return &GameMap{Height:height,
					Width:width,
					blocks:blocks,
					heightBlockNum:heightBlockNum,
					widthBlockNum:widthBlockNum}
}
//检测一个物体是否与地图中的方块发生碰撞
func (gameMap *GameMap) IsColl(gameObj GameObj)bool{
	x,y := gameObj.GetPosition();
	w,h := gameObj.GetSize()
	result := false
	startX,endX,startY,endY := gameMap.FetchBox(mgl32.Vec2{x,y},mgl32.Vec2{w,h})
	for i:=startX;i<=endX;i++{
		for j := startY;j<endY;j++{
			block := gameMap.blocks[int(i)][int(j)]
			if(block != nil){
				if(physic.IsCollidingAABB(gameObj,block)){
					gameMap.blocks[int(i)][int(j)] = nil
				}
			}
		}
	}
	return result
}
//将一个物体坐标转换为地图格子坐标范围
func (gameMap *GameMap) FetchBox(position,size mgl32.Vec2)(int,int,int,int){
	startY := int(math.Floor(float64((position[0]) / gameMap.Width * float32(gameMap.widthBlockNum))) - 1);
	if(startY <= 0){
		startY = 0
	}
	endY := int(math.Ceil(float64((position[0] + size[0]) / gameMap.Width * float32(gameMap.widthBlockNum))))
	if(endY >= gameMap.widthBlockNum){
		endY = gameMap.widthBlockNum - 1
	}
	startX := int(math.Floor(float64((position[1]) / gameMap.Height * float32(gameMap.heightBlockNum))))
	if(startX < 0){
		startX = 0
	}
	endX := int(math.Ceil(float64((position[1] + size[1]) / gameMap.Height * float32(gameMap.heightBlockNum))))
	if(endX >= gameMap.heightBlockNum){
		endX = gameMap.heightBlockNum - 1
	}
	return startX,endX,startY,endY
}
//渲染地图
func (gameMap *GameMap) Draw(position mgl32.Vec2, zoom mgl32.Vec2, renderer *sprite.SpriteRenderer){
	startX,endX,startY,endY := gameMap.FetchBox(position,zoom)
	for i:=startX;i<=endX;i++{
		for j := startY;j<endY;j++{
			block := gameMap.blocks[int(i)][int(j)]
			if(block != nil){
				block.Draw(renderer)
			}
		}
	}
}
