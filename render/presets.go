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
	VertexShader3D = `#version 460 core
		layout (location = 0) in vec3 aPos;
		layout (location = 1) in vec3 aColor;
		out vec3 color;

		uniform mat4 model;
		uniform mat4 view;
		uniform mat4 projection;

		void main() {
			// CRITICAL ORDER: Projection * View * Model * Vertex
			gl_Position = projection * view * model * vec4(aPos, 1.0);
			color = aColor;
		}`
)

var (
	TriangleVertices = []float32{
		-0.5, -0.5, 0.0, 1.0, 0.0, 0.0,
		0.5, -0.5, 0.0, 0.0, 1.0, 0.0,
		0.0, 0.5, 0.0, 0.0, 0.0, 1.0,
	}
	CubeVertices = []float32{
		-1, -1, -1, 1, 0, 0,
		1, -1, -1, 1, 0, 0,
		1, 1, -1, 1, 0, 0,
		-1, 1, -1, 1, 0, 0,
		-1, -1, 1, 0, 1, 0,
		1, -1, 1, 0, 1, 0,
		1, 1, 1, 0, 1, 0,
		-1, 1, 1, 0, 1, 0,
	}
	CubeIndices = []uint32{
		0, 1, 2, 0, 2, 3,
		4, 5, 6, 4, 6, 7,
		0, 3, 7, 0, 7, 4,
		1, 5, 6, 1, 6, 2,
		0, 4, 5, 0, 5, 1,
		3, 2, 6, 3, 6, 7,
	}
)
