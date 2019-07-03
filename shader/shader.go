package shader

import(
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"strings"
	"fmt"
)

type Shader struct{
	ID uint32
}

func Compile(vertexString, fragmentString string) *Shader{
	vertexShader,err := compile(vertexString+"\x00", gl.VERTEX_SHADER)
	if err != nil{
        panic(err)
	}
	fragmentShader,err := compile(fragmentString+"\x00", gl.FRAGMENT_SHADER)
	if err != nil{
        panic(err)
	}
	progID := gl.CreateProgram()
	gl.AttachShader(progID, vertexShader)
	gl.AttachShader(progID, fragmentShader)    
	gl.LinkProgram(progID)
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)
	return &Shader{ ID: progID}
}
func (shader *Shader) Use(){
	gl.UseProgram(shader.ID)
}

func (shader *Shader) SetBool(name string, value bool){
	var a int32 = 0;
	if(value){
		a = 1
	}
	gl.Uniform1i(gl.GetUniformLocation(shader.ID, gl.Str(name + "\x00")), a)
}

func (shader *Shader) SetInt(name string, value int32){
	gl.Uniform1i(gl.GetUniformLocation(shader.ID, gl.Str(name + "\x00")), value)
}

func (shader *Shader) SetFloat(name string, value float32){
	gl.Uniform1f(gl.GetUniformLocation(shader.ID, gl.Str(name + "\x00")), value)
}

func (shader *Shader) SetMatrix4fv(name string, value *float32){
	gl.UniformMatrix4fv(gl.GetUniformLocation(shader.ID, gl.Str(name + "\x00")), 1,false,value)
}

func (shader *Shader) SetVector3f(name string, vec3 mgl32.Vec3){
	gl.Uniform3f(gl.GetUniformLocation(shader.ID, gl.Str(name + "\x00")), vec3[0], vec3[1], vec3[2]);
}
func compile(sourceString string, shaderType uint32)(uint32, error){
	shader := gl.CreateShader(shaderType)
	source, free := gl.Strs(sourceString)
	gl.ShaderSource(shader, 1, source, nil)
    free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
        var logLength int32
        gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
        log := strings.Repeat("\x00", int(logLength+1))
        gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))
        return 0, fmt.Errorf("failed to compile %v: %v", source, log)
    }
    return shader, nil
}