# obj-loader
A simple Go (w/ GopherJS and three.js) webapp that allows users to upload .objs and have them displayed in a funky way!
The .obj Loader takes the vertices specified in the .obj file, normalizes the scale of the model and renders cubes of random colour in place of the vertices.
You can also fly around and explore the model! (with very bad controls!)

Just don't upload any obj over 1mb or so (depending on your hardware). The more vertices the more resource intensive it becomes.

Test it out [here](https://www.kouhai.world/)!

Made this for fun and to test out GopherJS, as well as dip my feet in WebGL.

<img src="https://www.kouhai.world/photos/cow.png">
