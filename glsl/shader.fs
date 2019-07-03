#version 410 core
in vec2 TexCoords;
out vec4 FragColor;

uniform sampler2D image;
uniform vec3 spriteColor;

void main()
{
    vec4 texColor = texture(image, TexCoords);
    if(texColor.a < 0.1)
        discard;
    FragColor = vec4(spriteColor, 1.0) * texColor;
}