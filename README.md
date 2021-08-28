<h1 align="center">
  <br>
  Purty Bits
  <br>
</h1>

# Package purtybits is a simple tool to format bytes into ANSI colorized binary


Download:
```bash
go get github.com/richinsley/purtybits
```

Usage:
```go
package main

import (
	"fmt"

	"github.com/richinsley/purtybits"
)

func main() {
	bytesPerRow := 4
	purty := purtybits.NewPurtyBits(bytesPerRow, purtybits.HexCodeGroupToRight)

	// create a custom color to color the hex values and open/close brackets
	purty.HexCodeColor = &purtybits.PurtyBitColor{
		BackgroundColor: purtybits.BackgroundBrightGreen,
		ForegroundColor: purtybits.ForegroundBrightWhite,
		Underline:       true,
	}

	purty.ByteOpenBracketColor = &purtybits.PurtyBitColor{
		BackgroundColor: purtybits.BackgroundBrightGreen,
		ForegroundColor: purtybits.ForegroundBrightWhite,
		Underline:       true,
		Character:       "<",
	}

	purty.ByteCloseBracketColor = &purtybits.PurtyBitColor{
		BackgroundColor: purtybits.BackgroundBrightGreen,
		ForegroundColor: purtybits.ForegroundBrightWhite,
		Underline:       true,
		Character:       ">",
	}

	// buffer of arbitrary bytes we'll purty print
	var buffer []uint8 = []uint8{0x01, 0x00, 0x01, 0xaa, 0xaa, 0xff, 0xff, 0x00}

	// we want to color bit 3 and bits 23-30 and red
	// v                        v-------v
	// 00000001|00000000|00000001|10101010|10101010|11111111|11111111|00000000
	purty.ColorBitRange(23, 30, purtybits.PurtyBitColorWhiteOnRed("1"), purtybits.PurtyBitColorBlackOnRed("0"))
	purty.ColorBitRange(3, 3, purtybits.PurtyBitColorWhiteOnRed("1"), purtybits.PurtyBitColorBlackOnRed("0"))

	// generate and display the purty printed rows
	fmt.Println("HexCodeGroupToRight:")
	purtyrows := purty.BufferToStrings(buffer)
	for _, s := range purtyrows {
		fmt.Println(s)
	}
	fmt.Println()

	// switch hex code placement to HexCodeLeftOfValue and generate again
	purty.HexCodePlacement = purtybits.HexCodeLeftOfValue
	fmt.Println("HexCodeLeftOfValue:")
	purtyrows = purty.BufferToStrings(buffer)
	for _, s := range purtyrows {
		fmt.Println(s)
	}
}
```

Output:

![bash output](/images/purtybits_output.png)