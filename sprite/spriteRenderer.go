package sprite
import(
	"game2D/shader"
	"game2D/texture"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type SpriteRenderer struct{
	shader *shader.Shader
	vao uint32
}
func NewSpriteRenderer(shader *shader.Shader) *SpriteRenderer{
	spriteRenderer := SpriteRenderer{shader:shader}
	spriteRenderer.initRenderData()
	return &spriteRenderer
}
func(spriteRenderer *SpriteRenderer) DrawSprite(texture *texture.Texture2D, position *mgl32.Vec2, size *mgl32.Vec2, rotate float32, color *mgl32.Vec3){
	model := mgl32.Translate3D(position[0], position[1], 0).Mul4(mgl32.Translate3D(0.5*size[0], 0.5*size[1], 0))
	model = model.Mul4(mgl32.HomogRotate3D(rotate, mgl32.Vec3{0, 0, 1}))
	model = model.Mul4(mgl32.Translate3D(-0.5*size[0], -0.5*size[1], 0))
	model = model.Mul4(mgl32.Scale3D(size[0], size[1], 1))
	
	spriteRenderer.shader.SetMatrix4fv("model", &model[0])
	spriteRenderer.shader.SetVector3f("spriteColor", *color)
	texture.Use()

	gl.BindVertexArray(spriteRenderer.vao);
    gl.DrawArrays(gl.TRIANGLES, 0, 6);
    gl.BindVertexArray(0);
}
func(spriteRenderer *SpriteRenderer) initRenderData(){
	var vbo uint32
	vertices := []float32{
		0.0, 1.0, 0.0, 1.0,
        1.0, 0.0, 1.0, 0.0,
        0.0, 0.0, 0.0, 0.0, 

        0.0, 1.0, 0.0, 1.0,
        1.0, 1.0, 1.0, 1.0,
        1.0, 0.0, 1.0, 0.0,
	}
	gl.GenVertexArrays(1, &spriteRenderer.vao);
    gl.GenBuffers(1, &vbo);

    gl.BindBuffer(gl.ARRAY_BUFFER, vbo);
    gl.BufferData(gl.ARRAY_BUFFER, 4 * len(vertices), gl.Ptr(vertices), gl.STATIC_DRAW);

    gl.BindVertexArray(spriteRenderer.vao);
    gl.EnableVertexAttribArray(0);
	gl.VertexAttribPointer(0, 4, gl.FLOAT, false, 4 * 4, gl.PtrOffset(0));
	gl.BindBuffer(gl.ARRAY_BUFFER, 0);
    gl.BindVertexArray(0);
}