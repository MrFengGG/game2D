package resource

import (
	"io/ioutil"
)

var (
	textures = make(map[string]*Texture2D)
	shaders  = make(map[string]*Shader)
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
	shaders[name] = Compile(string(vertexString), string(fragmentString))
}
func GetShader(name string) *Shader{
	return shaders[name]
}

func LoadTexture(TEXTUREINDEX uint32, file, name string){
	texture := NewTexture2D(file, TEXTUREINDEX)
	textures[name] = texture
}
func GetTexture(name string) *Texture2D{
	return textures[name]
}