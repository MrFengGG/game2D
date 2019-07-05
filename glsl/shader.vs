#version 410 core
layout (location = 0) in vec4 vertex; // <vec2 position, vec2 texCoords>

out vec2 TexCoords;

uniform mat4 model;
uniform mat4 projection;
uniform mat4 view;
uniform int reverseX;

void main()
{   if(reverseX == 1){
        TexCoords = vertex.zw;
    }else{
        TexCoords = vec2(1 - vertex.z, vertex.w);
    }
    gl_Position = projection * view * model * vec4(vertex.x, vertex.y, 0.0, 1.0);
}