# obj-loader
A simple Go (w/ GopherJS and three.js) webapp that allows users to upload .objs and have them displayed in a funky way!
The .obj Loader takes the verticies specified in the .obj file, normalizes the scale of the model and renders cubes of random colour in place of the verticies.
You can also fly around and explore the model! (with very bad controls!)

Just don't upload any obj over 1mb or so (depending on your hardware). The more verticies the more resource intensive it becomes.

Made this for fun and to test out GopherJS, as well as dip my feet in WebGL.

<img src="https://www.kouhai.world/photos/cow.png">
