package ansify

import (
	"image"
	"os"

	"golang.org/x/term"
)

// ResizeSrc function  INFO:  Resizes the SRC into a smaller dst image
func ResizeSrc(src image.Image) (*image.RGBA, error) {
	srcHeight := src.Bounds().Dy()
	srcWidth := src.Bounds().Dx()

	termWidth, termHeight, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return nil, err
	}

	targetWidth := termWidth - 2
	targetHeight := (termHeight - 2) * 2

	scaleX := float64(targetWidth) / float64(srcWidth)
	scaleY := float64(targetHeight) / float64(srcHeight)

	scale := min(scaleX, scaleY)

	newWidth := int(float64(srcWidth) * scale)
	newHeight := int(float64(srcHeight) * scale)

	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	for y := range newHeight {
		for x := range newWidth {
			srcX := x * srcWidth / newWidth
			srcY := y * srcHeight / newHeight

			dst.Set(x, y, src.At(srcX, srcY))
		}
	}

	return dst, nil
}
