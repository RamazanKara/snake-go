package main

import (
    "testing"
)

func TestRainbowColor(t *testing.T) {
    // Test for a few different frame counts
    frameCounts := []int{0, 10, 20, 30}
    for _, frameCount := range frameCounts {
        got := rainbowColor(frameCount)
        _, _, _, a := got.RGBA()
        if a != 0xffff {
            t.Errorf("rainbowColor(%v) = %v; want alpha value 0xffff", frameCount, got)
        }
    }
}
