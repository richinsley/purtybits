package purtybits

const (
	ANSIReset      string = "\033[0m"
	ANSIUnderlined string = "\u001b[4m"
	ANSIReversed   string = "\u001b[7m"
)

const (
	ForegroundBlack         string = "\033[30m"
	ForgroundBrightBlack    string = "\033[30;1m"
	ForgroundRed            string = "\033[31m"
	ForegroundBrightRed     string = "\033[31;1m"
	ForegroundGreen         string = "\033[32m"
	ForegroundBrightGreen   string = "\033[32;1m"
	ForegroundYellow        string = "\033[33m"
	ForegroundBrightYellow  string = "\033[33;1m"
	ForegroundBlue          string = "\033[34m"
	ForegroundBrightBlue    string = "\033[34;1m"
	ForegroundMagenta       string = "\033[35m"
	ForegroundBrightMagenta string = "\033[35;1m"
	ForegroundCyan          string = "\033[36m"
	ForegroundBrightCyan    string = "\033[36;1m"
	ForegroundWhite         string = "\033[37m"
	ForegroundBrightWhite   string = "\033[37;1m"

	BackgroundBlack         string = "\033[40m"
	BackgroundBrightBlack   string = "\033[40;1m"
	BackgroundRed           string = "\033[41m"
	BackgroundBrightRed     string = "\033[41;1m"
	BackgroundGreen         string = "\033[42m"
	BackgroundBrightGreen   string = "\033[42;1m"
	BackgroundYellow        string = "\033[43m"
	BackgroundBrightYellow  string = "\033[43;1m"
	BackgroundBlue          string = "\033[44m"
	BackgroundBrightBlue    string = "\033[44;1m"
	BackgroundMagenta       string = "\033[45m"
	BackgroundBrightMagenta string = "\033[45;1m"
	BackgroundCyan          string = "\033[46m"
	BackgroundBrightCyan    string = "\033[46;1m"
	BackgroundWhite         string = "\033[47m"
	BackgroundBrightWhite   string = "\033[47;1m"
)

type PurtyBitColor struct {
	ForegroundColor string
	BackgroundColor string
	Underline       bool
	Character       string `default:""`
}

func (me *PurtyBitColor) SetCharacter(c string) *PurtyBitColor {
	me.Character = c
	return me
}

func (me *PurtyBitColor) SetUnderlined(u bool) *PurtyBitColor {
	me.Underline = u
	return me
}

// GetColorCodes returns a string with ANSI codes defining the coloring and formatting rules
func (me *PurtyBitColor) GetColorCodes() string {
	retv := me.ForegroundColor + me.BackgroundColor
	if me.Underline {
		retv = ANSIUnderlined + retv
	}
	return retv + me.Character
}

// PurtyBitColorBlackOnYellow returns a PurtyBitColor specifying White characters on a Black background
func PurtyBitColorDefault(character string) *PurtyBitColor {
	return PurtyBitColorWhiteOnBlack(character)
}

// PurtyBitColorBlackOnYellow returns a PurtyBitColor specifying White characters on a Black background
func PurtyBitColorWhiteOnBlack(character string) *PurtyBitColor {
	return &PurtyBitColor{
		ForegroundColor: ForegroundBrightWhite,
		BackgroundColor: BackgroundBrightBlack,
		Character:       character,
	}
}

// PurtyBitColorBlackOnYellow returns a PurtyBitColor specifying White characters on a Red background
func PurtyBitColorWhiteOnRed(character string) *PurtyBitColor {
	return &PurtyBitColor{
		ForegroundColor: ForegroundBrightWhite,
		BackgroundColor: BackgroundBrightRed,
		Character:       character,
	}
}

// PurtyBitColorBlackOnYellow returns a PurtyBitColor specifying White characters on a Green background
func PurtyBitColorWhiteOnGreen(character string) *PurtyBitColor {
	return &PurtyBitColor{
		ForegroundColor: ForegroundBrightWhite,
		BackgroundColor: BackgroundBrightGreen,
		Character:       character,
	}
}

// PurtyBitColorBlackOnYellow returns a PurtyBitColor specifying White characters on a Yellow background
func PurtyBitColorWhiteOnYellow(character string) *PurtyBitColor {
	return &PurtyBitColor{
		ForegroundColor: ForegroundBrightWhite,
		BackgroundColor: BackgroundBrightYellow,
		Character:       character,
	}
}

// PurtyBitColorBlackOnYellow returns a PurtyBitColor specifying White characters on a Blue background
func PurtyBitColorWhiteOnBlue(character string) *PurtyBitColor {
	return &PurtyBitColor{
		ForegroundColor: ForegroundBrightWhite,
		BackgroundColor: BackgroundBrightBlue,
		Character:       character,
	}
}

// PurtyBitColorBlackOnYellow returns a PurtyBitColor specifying White characters on a Magenta background
func PurtyBitColorWhiteOnMagenta(character string) *PurtyBitColor {
	return &PurtyBitColor{
		ForegroundColor: ForegroundBrightWhite,
		BackgroundColor: BackgroundBrightMagenta,
		Character:       character,
	}
}

// PurtyBitColorBlackOnYellow returns a PurtyBitColor specifying White characters on a Cyan background
func PurtyBitColorWhiteOnCyan(character string) *PurtyBitColor {
	return &PurtyBitColor{
		ForegroundColor: ForegroundBrightWhite,
		BackgroundColor: BackgroundBrightCyan,
		Character:       character,
	}
}

// PurtyBitColorBlackOnYellow returns a PurtyBitColor specifying Black characters on a White background
func PurtyBitColorBlackOnWhite(character string) *PurtyBitColor {
	return &PurtyBitColor{
		ForegroundColor: ForgroundBrightBlack,
		BackgroundColor: BackgroundBrightWhite,
		Character:       character,
	}
}

// PurtyBitColorBlackOnYellow returns a PurtyBitColor specifying Black characters on a Red background
func PurtyBitColorBlackOnRed(character string) *PurtyBitColor {
	return &PurtyBitColor{
		ForegroundColor: ForgroundBrightBlack,
		BackgroundColor: BackgroundBrightRed,
		Character:       character,
	}
}

// PurtyBitColorBlackOnYellow returns a PurtyBitColor specifying Black characters on a Green background
func PurtyBitColorBlackOnGreen(character string) *PurtyBitColor {
	return &PurtyBitColor{
		ForegroundColor: ForgroundBrightBlack,
		BackgroundColor: BackgroundBrightGreen,
		Character:       character,
	}
}

// PurtyBitColorBlackOnYellow returns a PurtyBitColor specifying Black characters on a Yello background
func PurtyBitColorBlackOnYellow(character string) *PurtyBitColor {
	return &PurtyBitColor{
		ForegroundColor: ForgroundBrightBlack,
		BackgroundColor: BackgroundBrightYellow,
		Character:       character,
	}
}

// PurtyBitColorBlackOnBlue returns a PurtyBitColor specifying Black characters on a Blue background
func PurtyBitColorBlackOnBlue(character string) *PurtyBitColor {
	return &PurtyBitColor{
		ForegroundColor: ForgroundBrightBlack,
		BackgroundColor: BackgroundBrightBlue,
		Character:       character,
	}
}

// PurtyBitColorBlackOnMagenta returns a PurtyBitColor specifying Black characters on a Magenta background
func PurtyBitColorBlackOnMagenta(character string) *PurtyBitColor {
	return &PurtyBitColor{
		ForegroundColor: ForgroundBrightBlack,
		BackgroundColor: BackgroundBrightMagenta,
		Character:       character,
	}
}

// PurtyBitColorBlackOnCyan returns a PurtyBitColor specifying Black characters on a Cyan background
func PurtyBitColorBlackOnCyan(character string) *PurtyBitColor {
	return &PurtyBitColor{
		ForegroundColor: ForgroundBrightBlack,
		BackgroundColor: BackgroundBrightCyan,
		Character:       character,
	}
}
