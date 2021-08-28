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
<pre>HexCodeGroupToRight:
<span style="background-color:#333333"><font color="#D3D7CF"><b>000</b></font></span><span style="background-color:#CC0000"><font color="#333333"><b>0</b></font></span><span style="background-color:#333333"><font color="#D3D7CF"><b>0001|00000000|0000000</b></font></span><span style="background-color:#CC0000"><font color="#D3D7CF"><b>1</b></font></span><span style="background-color:#333333"><font color="#D3D7CF"><b>|</b></font></span><span style="background-color:#CC0000"><font color="#D3D7CF"><b>1</b></font></span><span style="background-color:#CC0000"><font color="#333333"><b>0</b></font></span><span style="background-color:#CC0000"><font color="#D3D7CF"><b>1</b></font></span><span style="background-color:#CC0000"><font color="#333333"><b>0</b></font></span><span style="background-color:#CC0000"><font color="#D3D7CF"><b>1</b></font></span><span style="background-color:#CC0000"><font color="#333333"><b>0</b></font></span><span style="background-color:#CC0000"><font color="#D3D7CF"><b>1</b></font></span><span style="background-color:#333333"><font color="#D3D7CF"><b>0|</b></font></span><span style="background-color:#4E9A06"><font color="#D3D7CF"><u style="text-decoration-style:single"><b>&lt;010001aa&gt;</b></u></font></span>
<span style="background-color:#333333"><font color="#D3D7CF"><b>10101010|11111111|11111111|00000000|</b></font></span><span style="background-color:#4E9A06"><font color="#D3D7CF"><u style="text-decoration-style:single"><b>&lt;aaffff00&gt;</b></u></font></span>

HexCodeLeftOfValue:
<span style="background-color:#4E9A06"><font color="#D3D7CF"><u style="text-decoration-style:single"><b>&lt;01&gt;</b></u></font></span><span style="background-color:#333333"><font color="#D3D7CF"><b>000</b></font></span><span style="background-color:#CC0000"><font color="#333333"><b>0</b></font></span><span style="background-color:#333333"><font color="#D3D7CF"><b>0001|</b></font></span><span style="background-color:#4E9A06"><font color="#D3D7CF"><u style="text-decoration-style:single"><b>&lt;00&gt;</b></u></font></span><span style="background-color:#333333"><font color="#D3D7CF"><b>00000000|</b></font></span><span style="background-color:#4E9A06"><font color="#D3D7CF"><u style="text-decoration-style:single"><b>&lt;01&gt;</b></u></font></span><span style="background-color:#333333"><font color="#D3D7CF"><b>0000000</b></font></span><span style="background-color:#CC0000"><font color="#D3D7CF"><b>1</b></font></span><span style="background-color:#333333"><font color="#D3D7CF"><b>|</b></font></span><span style="background-color:#4E9A06"><font color="#D3D7CF"><u style="text-decoration-style:single"><b>&lt;aa&gt;</b></u></font></span><span style="background-color:#CC0000"><font color="#D3D7CF"><b>1</b></font></span><span style="background-color:#CC0000"><font color="#333333"><b>0</b></font></span><span style="background-color:#CC0000"><font color="#D3D7CF"><b>1</b></font></span><span style="background-color:#CC0000"><font color="#333333"><b>0</b></font></span><span style="background-color:#CC0000"><font color="#D3D7CF"><b>1</b></font></span><span style="background-color:#CC0000"><font color="#333333"><b>0</b></font></span><span style="background-color:#CC0000"><font color="#D3D7CF"><b>1</b></font></span><span style="background-color:#333333"><font color="#D3D7CF"><b>0</b></font></span>
<span style="background-color:#4E9A06"><font color="#D3D7CF"><u style="text-decoration-style:single"><b>&lt;aa&gt;</b></u></font></span><span style="background-color:#333333"><font color="#D3D7CF"><b>10101010|</b></font></span><span style="background-color:#4E9A06"><font color="#D3D7CF"><u style="text-decoration-style:single"><b>&lt;ff&gt;</b></u></font></span><span style="background-color:#333333"><font color="#D3D7CF"><b>11111111|</b></font></span><span style="background-color:#4E9A06"><font color="#D3D7CF"><u style="text-decoration-style:single"><b>&lt;ff&gt;</b></u></font></span><span style="background-color:#333333"><font color="#D3D7CF"><b>11111111|</b></font></span><span style="background-color:#4E9A06"><font color="#D3D7CF"><u style="text-decoration-style:single"><b>&lt;00&gt;</b></u></font></span><span style="background-color:#333333"><font color="#D3D7CF"><b>00000000</b></font></span>
</pre>
