package render

const (
	DefaultVertexShader = `#version 460 core
        layout (location = 0) in vec3 aPos;
        layout (location = 1) in vec3 aColor;
        out vec3 color;
        void main() {
            gl_Position = vec4(aPos, 1.0);
            color = aColor;
        }`

	DefaultFragmentShader = `#version 460 core
        in vec3 color;
        out vec4 fragColor;
        void main() {
            fragColor = vec4(color, 1.0);
        }`
	TransformVertexShader = `#version 460 core
		layout (location = 0) in vec3 aPos;
		layout (location = 1) in vec3 aColor;
		out vec3 color;

		uniform mat4 model;

		void main() {
			gl_Position = model * vec4(aPos, 1.0); // Apply the rotation!
			color = aColor;
		}`
)
