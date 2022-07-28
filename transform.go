package gocolor

import (
	"fmt"
	"strconv"
	"strings"
)

// RgbToRGB rgb => r,g,b
func RgbToRGB(rgb int64) (r, g, b int64) {
	r, g, b = rgb&0xff, (rgb&0xff00)>>8, (rgb&0xff0000)>>16
	return
}

// RGBToHex r,g,b => hex
func RGBToHex(r, g, b int64) string {
	R := strconv.FormatInt(r, 16)
	if len(R) < 2 {
		R = "0" + R
	}
	G := strconv.FormatInt(g, 16)
	if len(G) < 2 {
		G = "0" + G
	}
	B := strconv.FormatInt(b, 16)
	if len(B) < 2 {
		B = "0" + B
	}
	return fmt.Sprintf("#%s%s%s", R, G, B)
}

// RgbToHex rgb => hex
func RgbToHex(rgb int64) string {
	r, g, b := rgb&0xff, (rgb&0xff00)>>8, (rgb&0xff0000)>>16
	R := strconv.FormatInt(r, 16)
	if len(R) < 2 {
		R = "0" + R
	}
	G := strconv.FormatInt(g, 16)
	if len(G) < 2 {
		G = "0" + G
	}
	B := strconv.FormatInt(b, 16)
	if len(B) < 2 {
		B = "0" + B
	}
	return fmt.Sprintf("#%s%s%s", R, G, B)
}

// HexToRgb 16进制颜色代码转为rgb int64
func HexToRgb(hexStr string) int64 {
	hexStr = strings.TrimLeft(hexStr, "#")
	R, _ := strconv.ParseInt(hexStr[0:2], 16, 10)
	G, _ := strconv.ParseInt(hexStr[2:4], 16, 10)
	B, _ := strconv.ParseInt(hexStr[4:], 16, 10)
	r, g, b := R, G<<8, B<<16
	return r + g + b
}
