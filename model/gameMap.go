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
	grounds := heightBlockNum / 4
	xGrounds := widthBlockNum / 4
	fmt.Println("map block size:",heightBlockNum,widthBlockNum)
	blocks := make([][]*Block,heightBlockNum)
	for i := 0;i < heightBlockNum;i++{
		rowBlocks := make([]*Block,widthBlockNum)
		if(i < grounds || i > grounds*3){
			for j := 0; j < widthBlockNum;j++{
				if(j < xGrounds || j > xGrounds * 3 || i == 0){
					gameObj := NewGameObj(resource.GetTexture("soil"),float32(j) * BlockWidth,float32(i)*BlockHeight,&mgl32.Vec2{BlockWidth,BlockHeight},0,&mgl32.Vec3{1,1,1})
					rowBlocks[j] = &Block{GameObj:*gameObj}
				}
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
func (gameMap *GameMap) IsColl(gameObj GameObj,shift mgl32.Vec2)(bool,mgl32.Vec2){
	position := gameObj.GetPosition();
	size := gameObj.GetSize()
	startX,endX,startY,endY := gameMap.FetchBox(mgl32.Vec2{position[0],position[1]},mgl32.Vec2{size[0],size[1]})
	for i:=startX;i<=endX;i++{
		for j := startY;j<endY;j++{
			block := gameMap.blocks[int(i)][int(j)]
			if(block != nil){
				fmt.Println("blockPosition",block.GetPosition())
				isCol,position := physic.ColldingAABBPlace(gameObj,block,shift)
				if(isCol){
					return isCol,position
				}
			}
		}
	}
	return false,gameObj.GetPosition()
}
//将一个物体坐标转换为地图格子坐标范围
func (gameMap *GameMap) FetchBox(position,size mgl32.Vec2)(int,int,int,int){
	startY := int(math.Floor(float64((position[0]) / gameMap.Width * float32(gameMap.widthBlockNum))))-1;
	if(startY <= 0){
		startY = 0
	}
	endY := int(math.Ceil(float64((position[0] + size[0]) / gameMap.Width * float32(gameMap.widthBlockNum)))) + 1
	if(endY >= gameMap.widthBlockNum){
		endY = gameMap.widthBlockNum - 1
	}
	startX := int(math.Floor(float64((position[1]) / gameMap.Height * float32(gameMap.heightBlockNum)))) -1
	if(startX < 0){
		startX = 0
	}
	endX := int(math.Ceil(float64((position[1] + size[1]) / gameMap.Height * float32(gameMap.heightBlockNum)))) +1
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
