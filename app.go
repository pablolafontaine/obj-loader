package main

import (
	"bufio"
	"fmt"
	"math"

	"math/rand"
	"strconv"

	"strings"

	"github.com/divan/three"
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

func main() {

    width := 600.0
    height := 600.0

    forward := false
    backward := false
    left := false
    right := false
    up := false
    down := false
    centre := false
   


    renderer := three.NewWebGLRenderer()
    
    renderer.SetSize(width, height, true)
    js.Global.Get("document").Get("body").Call("appendChild", renderer.Get("domElement"))

    // setup camera and scene
    camera := three.NewPerspectiveCamera(120, width/height, 1, 1000)
    camera.Position.Set(0, 100, 175)

    scene := three.NewScene()
    scene.Set("background", three.NewColor("white"))
    vertCubes := make([]three.Mesh, 0)

 
    
   


    // cube object
    geom := three.NewBoxGeometry(&three.BoxGeometryParameters{
        Width:  5,
        Height: 5,
        Depth:  5,
    })
    

    

    js.Global.Get("document").Call("write", `<input class="x" type="file" accept=".obj" style="display:block; margin-left:auto;margin-right:auto;width: 35%;">`)

	root := dom.GetWindow().Document()
	input := root.QuerySelector(".x")
    

    dom.GetWindow().AddEventListener("keydown", false, func(event dom.Event) {
		ke := event.(*dom.KeyboardEvent)
		
        switch(ke.KeyCode){
        case 87:
            forward = true;
        case 83:
            backward = true;
        case 65:
            left = true;
        case 68:
            right = true;
        case 81:
            up = true;
        case 69:
            down = true;
        case 32:
            centre = true;
        }

	})
    dom.GetWindow().AddEventListener("keyup", false, func(event dom.Event) {
		ke := event.(*dom.KeyboardEvent)
		
        switch(ke.KeyCode){
        case 87:
            forward = false;
        case 83:
            backward = false;
        case 65:
            left = false;
        case 68:
            right = false;
        case 81:
            up = false;
        case 69:
            down = false;
        case 32:
            centre = false;
        }
      
	})
     // start animation
    var animate func()
    animate = func() {
        
            
            
        
        js.Global.Call("requestAnimationFrame", animate)
        x,y,z := camera.Position.Coords()
        //println(scene.Object.Length())
        
        if forward{camera.Position.SetZ(z - 0.6)}
        if backward{camera.Position.SetZ(z + 0.6)}
        if right{camera.Position.SetX(x + 0.6)}
        if left{camera.Position.SetX(x - 0.6)}
        if up{camera.Position.SetY(y + 0.6)}
        if down{camera.Position.SetY(y - 0.6)}
        if centre{camera.LookAt(0.0,0.0,0.0)}
        

      //  println(vertCubes)

      //  for _,mesh := range vertCubes{
           // mesh.Rotation.Set("y", mesh.Rotation.Get("y").Float()+0.01)
       // }
        
        renderer.Render(scene, camera)

        
    }

	input.AddEventListener("change", true, func(dom.Event) {
		go func() {
            
            scene = three.NewScene()
            scene.Set("background", three.NewColor("white"))
            // lights
            light := three.NewDirectionalLight(three.NewColor("white"), 1)
            light.Position.Set(0, 256, 256)
            scene.Add(light)
           
            var minX, maxX, minY, maxY, minZ, maxZ float64
  
            vertCubes = nil
            vertStrings := make([]string, 0)
            // lights
            

            verts := make([]three.Vector3, 0)
            vertCubes = make([]three.Mesh, 0)
      
            scanner := bufio.NewScanner(strings.NewReader(string(blobToBytes(input.(*dom.HTMLInputElement).Files()[0].Object))))
            for scanner.Scan() {
                vertStrings = append(vertStrings, scanner.Text())
            }
            
            
    
            for _, str := range vertStrings{
                var firstLetter string
                reader := strings.NewReader(str)
                fmt.Fscanf(reader, "%s", &firstLetter)
               
                if firstLetter == "v"{
                    var x, y, z float64

                    fmt.Fscanf(reader, "%f %f %f\n", &x, &y, &z)
                   
                    if y < minY{
                        minY = y
                    }
                    if y > maxY{
                        maxY = y
                    }
                    if x < minX{
                        minX = x
                    }
                    if x > maxX{
                        maxX = x
                    }
                    if z < minZ{
                        minZ = z
                    }
                    if z > maxZ{
                        maxZ = z
                    }

                    vec3 := three.NewVector3(x,y,z)
                    verts = append(verts, vec3)
                }
        
            }

            scale := (1.0/(math.Max(math.Max(maxX-minX,maxY-minY),maxZ-minZ))*500.0)
            

            //points
            for _,vert := range verts{
                // cube material
                params := three.NewMaterialParameters()

                r := strconv.Itoa(rand.Intn(255))
                g := strconv.Itoa(rand.Intn(255))
                b := strconv.Itoa(rand.Intn(255))
                
                params.Color = three.NewColor("rgb("+r+","+g+","+b+")")
                mat := three.NewMeshLambertMaterial(params)
                mesh := three.NewMesh(geom, mat)
                x, y, z := vert.Coords()
                mesh.Position.Set(x*scale,y*scale,z*scale)
                vertCubes = append(vertCubes, *mesh)
                
                scene.Add(mesh)
                mesh = nil
                params = nil

            }
            verts = nil
            
            animate()
            
		}()
	})

 
    
    
   
  

}



func blobToBytes(blob *js.Object) []byte {
	var b = make(chan []byte)
	fileReader := js.Global.Get("FileReader").New()
	fileReader.Set("onload", func() {
		b <- js.Global.Get("Uint8Array").New(fileReader.Get("result")).Interface().([]byte)
	})
	fileReader.Call("readAsArrayBuffer", blob)
	return <-b
}

