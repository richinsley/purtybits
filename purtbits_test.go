package purtybits

import (
	"fmt"
	"testing"
)

func TestSimple(t *testing.T) {
	purty := NewPurtyBits(8, HexCodeGroupToRight)

	// create a custom color to color the hex values and open/close brackets
	purty.HexCodeColor = &PurtyBitColor{
		BackgroundColor: BackgroundBrightGreen,
		ForegroundColor: ForegroundBrightWhite,
		Underline:       true,
	}

	purty.ByteOpenBracketColor = &PurtyBitColor{
		BackgroundColor: BackgroundBrightGreen,
		ForegroundColor: ForegroundBrightWhite,
		Underline:       true,
		Character:       "<",
	}

	purty.ByteCloseBracketColor = &PurtyBitColor{
		BackgroundColor: BackgroundBrightGreen,
		ForegroundColor: ForegroundBrightWhite,
		Underline:       true,
		Character:       ">",
	}

	// buffer of arbitrary bytes we'll purty print
	var buffer []uint8 = []uint8{0x01, 0x00, 0x01, 0xaa, 0xaa, 0xff, 0xff, 0x00}

	// we want to color bit 3 and bits 23-30 and red
	// v                        v-------v
	// 00000001|00000000|00000001|10101010|10101010|11111111|11111111|00000000
	purty.ColorBitRange(23, 30, PurtyBitColorWhiteOnRed("1"), PurtyBitColorBlackOnRed("0"))
	purty.ColorBitRange(0, 0, PurtyBitColorWhiteOnRed("1"), PurtyBitColorBlackOnRed("0"))

	// generate and display the purty printed rows
	fmt.Println("HexCodeGroupToRight:")
	purtyrows := purty.BufferToStrings(buffer)
	for _, s := range purtyrows {
		fmt.Println(s)
	}
	fmt.Println()

	// switch hex code placement to HexCodeLeftOfValue and generate again
	purty.HexCodePlacement = HexCodeLeftOfValue
	fmt.Println("HexCodeLeftOfValue:")
	purtyrows = purty.BufferToStrings(buffer)
	for _, s := range purtyrows {
		fmt.Println(s)
	}
}
