// Package ansify provides ansify  INFO:  ANSIfy the image inputted
package ansify

import (
	"fmt"
	"log"
)

// ANSIfy function  INFO:  ANSIfies a given image
func ANSIfy(filePath string) {
	srcObj, err := LoadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	dst, err := ResizeSrc(srcObj)
	if err != nil {
		log.Fatal(err)
	}

	for y := 0; y < dst.Rect.Dy(); y += 2 {
		for x := range dst.Rect.Dx() {
			topR, topG, topB, _ := dst.At(x, y).RGBA()
			bottomR, bottomG, bottomB, _ := dst.At(x, y+1).RGBA()

			topR8, topG8, topB8 := topR>>8, topG>>8, topB>>8
			bottomR8, bottomG8, bottomB8 := bottomR>>8, bottomG>>8, bottomB>>8

			fmt.Printf(
				"\033[38;2;%d;%d;%dm\033[48;2;%d;%d;%dm▀",
				topR8, topG8, topB8,
				bottomR8, bottomG8, bottomB8,
			)

		}

		fmt.Print("\033[0m\n")

	}
}
