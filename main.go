package main

import (
    "os"
    "fmt"
    "log"
    "net/http"
    "image"
    "strings"
    "image/png"
    "image/draw"
    "path/filepath"
    "github.com/mkideal/cli"
    "github.com/mandarl/go-selfupdate/selfupdate"
)

var VERSION string = "dev"

type argT struct {
    cli.Helper
    InputImage string `cli:"*i,input-image" usage:"input screenshot image"`
    Device string `cli:"d,device" usage:"the device skin to use" dft:"nexus_6"`
    Orientation string `cli:"o,orientation" usage:"device orientation; can be port or land" dft:"port"`
    OutputFile string `cli:"p,output-file" usage:"output file path" dft:"output.png"`
    Version bool `cli:"!v,version" usage:"print the current version"`
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
    
    if(args.Version) {
        fmt.Printf("device-art: Verison: %s\n", VERSION)
        os.Exit(0)
    }
    
    runUpdate()
    
    coordinates := map[string][2]int{		//map[filename] [2]int[x,y]
	"nexus_6_port":  [2]int{229, 239},
	"nexus_6_land":  [2]int{318, 77},
	"nexus_6p_port":  [2]int{312, 579},
	"nexus_6p_land":  [2]int{579, 320},
	"nexus_5_port":  [2]int{306, 436},
	"nexus_5_land":  [2]int{436, 306},
	"nexus_5x_port":  [2]int{305, 485},
	"nexus_5x_land":  [2]int{484, 313},
	"iphone_6_port":  [2]int{218, 320},
	"iphone_6_land":  [2]int{320, 185},
	"iphone_6_plus_port":  [2]int{380, 420},
	"iphone_6_plus_land":  [2]int{420, 339},
	"iphone_5_port":  [2]int{219, 339},
	"iphone_5_land":  [2]int{339, 221},
	"ipad_air_2_port":  [2]int{203, 314},
	"ipad_air_2_land":  [2]int{314, 205},
    }
    // coordinates := map[string]int{
    //     "nexus_6_port": 1
    // }

    deviceBackImage := readImageWeb(imageBackPath(args))
    screenshotImage := readImage(args.InputImage)
    
    //starting position of the second image (bottom left)
    x := coordinates[fmt.Sprintf("%s_%s", args.Device, args.Orientation)][0]
    y := coordinates[fmt.Sprintf("%s_%s", args.Device, args.Orientation)][1]
    sp2 := image.Point{x, y}
    
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
    fmt.Printf("Success - output file is : %s\n", args.OutputFile)
}

func imageBackPath (args *argT) string {
    return fmt.Sprintf("%s_%s_back.png", args.Device, args.Orientation)
}

func imageBasePath (args *argT) string {
    path := ""//args.ImageDirectory
    // if (path == "") {
    //     path = "./data"
    // }
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

func readImageWeb(path string) image.Image{
    url := "https://raw.githubusercontent.com/mandarl/device-art/master/assets/" + path

    response, e := http.Get(url)
    if e != nil {
        log.Fatal(e)
    }

    defer response.Body.Close()
    
    // infile, err := Asset(path) //os.Open(path)
    // if err != nil {
    //     // replace this with real error handling
    //     log.Fatal("can't open file", err)
    // }
    // r := bytes.NewReader(infile)
    //defer infile.Close()
    
    // Decode will figure out what type of image is in the file on its own.
    // We just have to be sure all the image packages we want are imported.
    src, _, err := image.Decode(response.Body)
    if err != nil {
        // replace this with real error handling
        log.Fatal("can't read image", path, err)
    }
    return src
}

func runUpdate() {
    var updater = &selfupdate.Updater{
        CurrentVersion: VERSION,
        ApiURL:         "http://dipoletech.com/projects/", //u.fetch(u.ApiURL + u.CmdName + "/" + plat + ".json")
        BinURL:         "http://dipoletech.com/projects/", //u.BinURL + u.CmdName + "/" + u.Info.Version + "/" + plat + ".gz"
        DiffURL:        "", //u.fetch(u.DiffURL + u.CmdName + "/" + u.CurrentVersion + "/" + u.Info.Version + "/" + plat)
        Dir:            "update/",
        CmdName:        "device-art", // this is added to apiurl to get json
    }
    
    if updater != nil {
        go updater.BackgroundRun()
    }
}