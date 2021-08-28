package purtybits

import (
	"errors"
	"fmt"
	"math"
)

var (
	defaultcolor PurtyBitColor = PurtyBitColor{
		BackgroundColor: BackgroundBrightBlack,
		ForegroundColor: ForegroundBrightWhite,
		Underline:       false,
	}
)

type HexCodePlacement int

const (
	// Do not include hex code
	HexCodeNone HexCodePlacement = iota
	// Display hex code to left of binary value
	HexCodeLeftOfValue
	// Display Hex code to right of binary value
	HexCodeRightOfValue
	// Group all hex codes to left of values
	HexCodeGroupToLeft
	// Group all hex codes to right of values
	HexCodeGroupToRight
)

type BitRange struct {
	StartBitPosition int
	EndBitPosition   int
	OneColor         *PurtyBitColor
	ZeroColor        *PurtyBitColor
}

type PurtyBits struct {
	ColumnCount           int
	HexCodePlacement      HexCodePlacement
	ByteSeperator         string
	HexCodeColor          *PurtyBitColor
	BinarySeperatorColor  *PurtyBitColor
	OnesColor             *PurtyBitColor
	ZedsColor             *PurtyBitColor
	ByteOpenBracketColor  *PurtyBitColor
	ByteCloseBracketColor *PurtyBitColor
	ColorRanges           []*BitRange

	validated bool
}

// NewPurtyBits returns a PurtyBits with default formatting rules
func NewPurtyBits(column_count int, hexcodeplacement HexCodePlacement) *PurtyBits {
	retv := &PurtyBits{
		ColumnCount:      column_count,
		HexCodePlacement: hexcodeplacement,
	}
	retv.validate()
	return retv
}

func (me *PurtyBits) getByteSeperator() string {
	return me.BinarySeperatorColor.GetColorCodes()
}

func (me *PurtyBits) get_one_bit(absoluteBitPos int) string {
	color := me.OnesColor
	for _, v := range me.ColorRanges {
		if absoluteBitPos >= v.StartBitPosition && absoluteBitPos <= v.EndBitPosition {
			color = v.OneColor
		}
	}
	return color.GetColorCodes() + ANSIReset
}

func (me *PurtyBits) get_zero_bit(absoluteBitPos int) string {
	color := me.ZedsColor
	for _, v := range me.ColorRanges {
		if absoluteBitPos >= v.StartBitPosition && absoluteBitPos <= v.EndBitPosition {
			color = v.ZeroColor
		}
	}
	return color.GetColorCodes() + ANSIReset
}

func (me *PurtyBits) gen_purty_bits(b uint8, byte_index int) string {
	var retv string
	for i := 0; i < 8; i++ {
		if b&0x01 == 0x01 {
			retv = me.get_one_bit(byte_index*8+(7-i)) + retv
		} else {
			retv = me.get_zero_bit(byte_index*8+(7-i)) + retv
		}
		b = b >> 1
	}
	return retv
}

func (me *PurtyBits) gen_purty_row(buffer []uint8, buffer_index int) string {
	var retv string
	for i, v := range buffer {
		if me.HexCodePlacement == HexCodeLeftOfValue {
			retv += fmt.Sprintf("%s%s%02x%s", me.ByteOpenBracketColor.GetColorCodes(), me.HexCodeColor.GetColorCodes(), v, me.ByteCloseBracketColor.GetColorCodes()) + me.HexCodeColor.GetColorCodes() + ANSIReset
		}
		retv = retv + me.gen_purty_bits(v, buffer_index+i)
		if me.HexCodePlacement == HexCodeRightOfValue {
			retv += fmt.Sprintf("%s%s%02x%s", me.ByteOpenBracketColor.GetColorCodes(), me.HexCodeColor.GetColorCodes(), v, me.ByteCloseBracketColor.GetColorCodes()) + ANSIReset
		}
		if i != (len(buffer))-1 {
			retv += me.getByteSeperator()
		}
	}
	return retv
}

func (me *PurtyBits) gen_purty_bytes(buffer []uint8) string {
	var retv string = me.ByteOpenBracketColor.GetColorCodes() + me.HexCodeColor.GetColorCodes()

	for _, v := range buffer {
		retv += fmt.Sprintf("%02x", v)
	}
	// pad based off of column count
	pcount := me.ColumnCount - len(buffer)
	if pcount != 0 {
		for j := 0; j < pcount; j++ {
			retv += "  "
		}
	}
	retv += me.ByteCloseBracketColor.GetColorCodes() + ANSIReset
	return retv
}

// BufferToStrings formats the given buffer into an of array of strings
// binary representations, applying the ANSI coloring and formatting rules.
func (me *PurtyBits) BufferToStrings(buffer []uint8) []string {
	me.validate()

	var retv []string
	ccount := me.ColumnCount
	if ccount < 1 {
		ccount = 1 // clamp to a min of 1
	}

	rcount := int(math.Ceil(float64(len(buffer)) / float64(ccount)))
	for j := 0; j < rcount; j++ {
		slice_sindex := j * ccount
		slice_eindex_offset := int(math.Min(float64(j*ccount+ccount), float64(len(buffer))))
		rslice := buffer[slice_sindex:slice_eindex_offset]
		purty_row := me.gen_purty_row(rslice, j*ccount)
		if me.HexCodePlacement == HexCodeGroupToLeft {
			purty_row = me.gen_purty_bytes(rslice) + me.getByteSeperator() + purty_row
		} else if me.HexCodePlacement == HexCodeGroupToRight {
			purty_row = purty_row + me.getByteSeperator() + me.gen_purty_bytes(rslice)
		}
		retv = append(retv, purty_row)
	}
	return retv
}

func (me *PurtyBits) validate() {
	if me.validated {
		return
	}

	if me.HexCodeColor == nil {
		v := defaultcolor
		me.HexCodeColor = v.SetCharacter("")
	}

	if me.BinarySeperatorColor == nil {
		v := defaultcolor
		me.BinarySeperatorColor = v.SetCharacter("|")
	}

	if me.OnesColor == nil {
		v := defaultcolor
		me.OnesColor = v.SetCharacter("1")
	}

	if me.ZedsColor == nil {
		v := defaultcolor
		me.ZedsColor = v.SetCharacter("0")
	}

	if me.ByteOpenBracketColor == nil {
		v := defaultcolor
		me.ByteOpenBracketColor = v.SetCharacter("[")
	}

	if me.ByteCloseBracketColor == nil {
		v := defaultcolor
		me.ByteCloseBracketColor = v.SetCharacter("]")
	}
	me.validated = true
}

func (me *PurtyBits) ColorBitRange(start_bit_position int, end_bit_position int, onecolor *PurtyBitColor, zerocolor *PurtyBitColor) error {
	if end_bit_position < start_bit_position {
		return errors.New("invalid bit range")
	}

	if onecolor == nil {
		// copy the default color and set the character value
		v := defaultcolor
		onecolor = v.SetCharacter("1")
	}

	if zerocolor == nil {
		// copy the default color and set the character value
		v := defaultcolor
		zerocolor = v.SetCharacter("1")
	}

	me.ColorRanges = append(me.ColorRanges, &BitRange{
		StartBitPosition: start_bit_position,
		EndBitPosition:   end_bit_position,
		ZeroColor:        zerocolor,
		OneColor:         onecolor,
	})
	return nil
}
