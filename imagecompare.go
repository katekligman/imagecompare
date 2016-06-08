package imagecompare

import (
    "errors"
    "image"
    "image/color"
    "image/png"
    "os"

    "github.com/disintegration/imaging"
)

func CompareImage(a image.Image, b image.Image) (int) {
    a_bounds := a.Bounds()

    // simple image compare
    for y := a_bounds.Min.Y; y < a_bounds.Max.Y; y++ {
        for x := a_bounds.Min.X; x < a_bounds.Max.X; x++ {
            a_pixel := a.At(x, y)
            b_pixel := b.At(x, y)
            ar, ag, ab, aa := a_pixel.RGBA()
            br, bg, bb, ba := b_pixel.RGBA()
            if ar != br || ag != bg || ab != bb || aa != ba {
                return 1
            }
        }
    }
    return 0
}

func GetImageMask(block_size int, a image.Image, b image.Image) (image.Image, error) {
    
    // images must have the same dimensions
    ba := a.Bounds()
    bb := b.Bounds()

    if ba.Max.X != bb.Max.X || ba.Max.Y != bb.Max.Y {
        return nil, errors.New("Images must have the same dimensions")
    }

    // image mask
    c := imaging.New(ba.Max.X, bb.Max.Y, color.RGBA{0, 0, 0, 0})
    //c := imaging.Clone(a)

    // mask block for fill
    block := imaging.New(block_size, block_size, color.RGBA{0, 0, 0, 255})

    for i := 0; i <= ba.Max.X; i += block_size {
        for j := 0; j <= ba.Max.Y; j += block_size {
            rt := image.Pt(i, j)
            rb := image.Pt(i+block_size, j+block_size)
            a_chunk := imaging.Crop(a, image.Rectangle{rt, rb})
            b_chunk := imaging.Crop(b, image.Rectangle{rt, rb})
            if CompareImage(a_chunk, b_chunk) != 0 {
                c = imaging.Paste(c, block, rt)
            } else {
                //c = imaging.Paste(c, a_chunk, rt)
            }
        }
    }

    return c, nil
}

func ThreeWayImageCompare(block_size int, a1_path string, a2_path string, b_path string) (int) {

    a1, err := imaging.Open(a1_path); if err != nil {
        return -1
    }
    a2, err := imaging.Open(a2_path); if err != nil {
        return -1
    }
    b, err := imaging.Open(b_path); if err != nil {
        return -1
    }
    mask, err := GetImageMask(block_size, a1, a2); if err != nil {
        return -1
    }

    c1 := imaging.Overlay(a1, mask, image.Pt(0, 0), 100.0)
    c2 := imaging.Overlay(b, mask, image.Pt(0, 0), 100.0)

    return CompareImage(c1, c2)
}
