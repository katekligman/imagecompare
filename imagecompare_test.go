package imagecompare

import "testing"
import "github.com/disintegration/imaging"

func TestGetImageMask(t *testing.T) {
    a, err := OpenPng("test/Test_card_a.png"); if err != nil {
        t.Errorf("Can't open test image")
    }
    b, err := OpenPng("test/Test_card_b.png"); if err != nil {
        t.Errorf("Can't open test image")
    }
    mask, err := GetImageMask(16, a, b); if err != nil {
        mask = mask
        t.Errorf("GetImageMask no workie")
    } 
    imaging.Save(mask, "mask.png")
}

func TestCompareImage(t *testing.T) {
    img, err := imaging.Open("test/Test_card_a.png"); if err != nil {
        t.Errorf("Issue opening test image")
    }
    img2, err := imaging.Open("test/Test_card_b.png"); if err != nil {
        t.Errorf("Issue opening test image")
    }
    if CompareImage(img, img2) == 0 {
        t.Errorf("A difference should be detected")
    }
}

func TestThreeWayImageCompare(t *testing.T) {
    a1 := "test/Test_card_a.png"
    a2 := "test/Test_card_b.png"
    a3 := "test/Test_card_c.png"
    //b := "PM5544_with_non-PAL_signals.png"

    if ThreeWayImageCompare(16, a1, a2, a3) != 1 {
        t.Errorf("a1, a2, a3 should return 1")
    }

    if ThreeWayImageCompare(16, a1, a2, a1) != 0 {
        t.Errorf("a1, a2, a1 should return 0")
    }

}
