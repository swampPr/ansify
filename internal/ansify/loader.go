package ansify

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"

	_ "golang.org/x/image/webp"
)

// LoadFile function  INFO:  Loads the src image given by user
func LoadFile(filePath string) (image.Image, error) {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, err
	}

	_, err = os.Stat(absPath)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}

	imgObj, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return imgObj, nil
}
