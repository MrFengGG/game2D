package physic

//检测两个矩形是否发生碰撞
func IsCollidingAABB(thisGameObj,anotherObj React) bool{
	tX,tY := thisGameObj.GetPosition()
	tW,tH := thisGameObj.GetSize()

	aX,aY := anotherObj.GetPosition()
	aW,aH := anotherObj.GetSize()
	return isCollidingReact(tX,tY,tW,tH,aX,aY,aW,aH);
}
type React interface{
	GetPosition() (float32,float32)
	GetSize() (float32,float32)
}
func isCollidingReact(x1,y1,w1,h1,x2,y2,w2,h2 float32) bool{
	// x轴方向碰撞？
	collisionX := x1 + w1 >= x2 && x2 + w2 >= x1
	// y轴方向碰撞？
	collisionY := y1 + h1 >= y2 && y2 + h2 >= y1
	return collisionX && collisionY
}
//检测两个矩形运动后是否会发生碰撞
func WillCollidingAABB(thisGameObj,anotherObj React,dtX,dtY float32) bool{
	tX,tY := thisGameObj.GetPosition()
	tX += dtX
	tY += dtY
	tW,tH := thisGameObj.GetSize()
	aX,aY := anotherObj.GetPosition()
	aW,aH := anotherObj.GetSize()
	return isCollidingReact(tX,tY,tW,tH,aX,aY,aW,aH);
}