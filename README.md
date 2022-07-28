### Example

```go
package main

import "github.com/Sorks/gocolor"


func main() {
    // 1. RGB Int 转换为 (R, G, B)
    r, g, b := gocolor.RgbToRGB(RgbInt)
    // 2. RGB 转16进制
    hexColor := gocolor.RgbToHex(RgbInt)
    // 3. r,g,b => hex
    hex := gocolor.RGBToHex(r, g, b)
    // 4. 16进制颜色代码转为rgb int	
    rgb := gocolor.HexToRgb(hexStr)
}
```