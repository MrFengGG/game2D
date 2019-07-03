package resource

import (
	"io/ioutil"
	"game2D/texture"
	"game2D/shader"
)

var (
	textures = make(map[string]*texture.Texture2D)
	shaders  = make(map[string]*shader.Shader)
)

func LoadShader(vShaderFile, fShaderFile, name string){
	vertexString, err := ioutil.ReadFile(vShaderFile)
	if err != nil{
        panic(err)
	}
	fragmentString, err := ioutil.ReadFile(fShaderFile)
	if err != nil{
        panic(err)
	}
	shaders[name] = shader.Compile(string(vertexString), string(fragmentString))
}
func GetShader(name string) *shader.Shader{
	return shaders[name]
}

func LoadTexture(TEXTUREINDEX uint32, file, name string){
	texture := texture.NewTexture2D(file, TEXTUREINDEX)
	textures[name] = texture
}
func GetTexture(name string) *texture.Texture2D{
	return textures[name]
}