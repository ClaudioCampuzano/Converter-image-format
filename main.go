package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"strings"
	"image"
	"image/png"
	"image/jpeg"
	"golang.org/x/image/bmp"
)

func main() {
	archivos, err := ioutil.ReadDir(os.Args[1])
	if err != nil {
        log.Fatal(err)
    }
	for _, archivo := range archivos {

		imageinput, err := os.Open(os.Args[1]+archivo.Name())
		if err != nil {
			log.Fatal(err)
		}
		defer imageinput.Close()

		filenameSplit := strings.Split(archivo.Name(), ".")
		format := filenameSplit[len(filenameSplit)-1]

		if (strings.ToLower(format) != strings.ToLower(os.Args[2])){
			var src image.Image
			switch strings.ToLower(format) {
			case "png":
				src, err = png.Decode(imageinput)
			case "jpg", "jpeg":
				src, err = jpeg.Decode(imageinput)
			case "bmp":
				src, err = bmp.Decode(imageinput)
			default:
				fmt.Println("The " + format + " we don't support to convert")
				os.Exit(1)
			}
		
			if err != nil {
				log.Fatal(err)
			}
		
			outfile, err := os.Create(os.Args[1]+filenameSplit[0]+"."+strings.ToLower(os.Args[2]))
			if err != nil {
				log.Fatal(err)
			}
			defer outfile.Close()
		
			switch strings.ToLower(os.Args[2]) {
			case "png":
				err = png.Encode(outfile, src)
			case "jpg", "jpeg":
				err = jpeg.Encode(outfile, src, nil)
			case "bmp":
				err = bmp.Encode(outfile, src)
			default:
				fmt.Println("The " + format + " we don't support to convert")
				os.Exit(1)
			}
			if err != nil {
				log.Fatal(err)
			}
			
			// Borramos original
			err = os.Remove(os.Args[1]+archivo.Name())
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}