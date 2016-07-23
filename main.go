package main

import (
    "os"
    "fmt"
    "log"
    "image"
    "strings"
    "image/png"
    "image/draw"
    "path/filepath"
    "github.com/mkideal/cli"
)

var VERSION string = "DEV"

type argT struct {
    cli.Helper
    InputImage string `cli:"*i,input-image" usage:"input screenshot image"`
    Device string `cli:"d,device" usage:"the device skin to use" dft:"nexus_6"`
    Orientation string `cli:"o,orientation" usage:"device orientation; can be port or land" dft:"port"`
    OutputFile string `cli:"p,output-file" usage:"output file path" dft:"output.png"`
    //ImageDirectory string `cli:"a,device-art-dir" usage:"directory containing device art images"`
}

// Validate implements cli.Validator interface
func (argv *argT) Validate(ctx *cli.Context) error {
	if !strings.EqualFold(argv.Orientation, "port") && !strings.EqualFold(argv.Orientation, "land") {
		return fmt.Errorf("orientation must be either port or land - %s is not a valid value", ctx.Color().Yellow(argv.Orientation))
	}
	return nil
}


func main() {
    cli.Run(&argT{}, func(ctx *cli.Context) error {
        argv := ctx.Argv().(*argT)
        run(argv)
        return nil
    })
}

func run(args *argT) {

    deviceBackImage := readImage(imageBackPath(args))
    screenshotImage := readImage(args.InputImage)
    
    //starting position of the second image (bottom left)
    sp2 := image.Point{229, 239}
    
    //new rectangle for the second image
    r2 := image.Rectangle{sp2, deviceBackImage.Bounds().Max}
    
    //rectangle for the big image
    r := image.Rectangle{image.ZP, r2.Max}
    
    rgba := image.NewRGBA(r)
    
    draw.Draw(rgba, deviceBackImage.Bounds(), deviceBackImage, image.ZP, draw.Src)
    draw.Draw(rgba, r2, screenshotImage, image.ZP, draw.Src)
    
    out, err := os.Create(args.OutputFile)
    if err != nil {
        log.Fatal("can't create file", err)
    }
    
    png.Encode(out, rgba)
}

func imageBackPath (args *argT) string {
     return pathJoin(imageBasePath(args), 
                fmt.Sprintf("%s_%s_back.png", args.Device, args.Orientation))
}

func imageBasePath (args *argT) string {
    path := ""//args.ImageDirectory
    if (path == "") {
        path = "./data"
    }
    return path
}

func pathJoin (path1 string, path2 string) string{
    path := filepath.Join(path1, path2)
    if _, err := os.Stat(path); os.IsNotExist(err) {
      log.Panicf("Path %s does not exist", path)
    }
    return path
}

func readImage(path string) image.Image{
    infile, err := os.Open(path)
    if err != nil {
        // replace this with real error handling
        log.Fatal("can't open file", err)
    }
    defer infile.Close()
    
    // Decode will figure out what type of image is in the file on its own.
    // We just have to be sure all the image packages we want are imported.
    src, _, err := image.Decode(infile)
    if err != nil {
        // replace this with real error handling
        log.Fatal("can't read image", path, err)
    }
    return src
}