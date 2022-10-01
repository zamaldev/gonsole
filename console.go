package gonsole

import (
	"errors"
	"fmt"
)

type color int

const (
	DEFAULT                             = "\x1b[0m" // All attributes off
	BOLD                                = "\x1b[1m" // As with faint, the color change is a PC (SCO / CGA) invention
	FAINT                               = "\x1b[2m" // May be implemented as a light font weight like bold
	ITALIC                              = "\x1b[3m" // Not widely supported. Sometimes treated as inverse or blink
	UNDERLINED                          = "\x1b[4m" // Style extensions exist for Kitty, VTE, mintty and iTerm2
	BLINKING_SLOW                       = "\x1b[5m" // Sets blinking to less than 150 times per minute
	BLINKING_RAPID                      = "\x1b[6m" // MS-DOS ANSI.SYS, 150+ per minute; not widely supported
	INVERTED                            = "\x1b[7m" // Swap foreground and background colors; inconsistent emulation
	HIDE                                = "\x1b[8m" // Not widely supported
	CROSSED                             = "\x1b[9m" // Characters legible but marked as if for deletion. Not supported in Terminal.app
	PRIMARY                             = "\x1b[10m"
	ALTERNATE_FONT_1                    = "\x1b[11m" // Rarely supported
	ALTERNATE_FONT_2                    = "\x1b[12m"
	ALTERNATE_FONT_3                    = "\x1b[13m"
	ALTERNATE_FONT_4                    = "\x1b[14m"
	ALTERNATE_FONT_5                    = "\x1b[15m"
	ALTERNATE_FONT_6                    = "\x1b[16m"
	ALTERNATE_FONT_7                    = "\x1b[17m"
	ALTERNATE_FONT_8                    = "\x1b[18m"
	ALTERNATE_FONT_9                    = "\x1b[19m"
	FRAKTUR                             = "\x1b[20m"
	DOUBLE_UNDERLINED                   = "\x1b[21m" // Double-underline per ECMA-48, but instead disables bold intensity on several terminals, including in the Linux kernel's console before version 4.17
	NORMAL_INTENSITY                    = "\x1b[22m" // Neither bold nor faint; color changes where intensity is implemented as such.
	NO_ITALIC_NOR_BLACKLETTER           = "\x1b[23m"
	NOT_UNDERLINED                      = "\x1b[24m" // Neither singly nor doubly underlined
	NOT_BLINKING                        = "\x1b[25m" // Turn blinking off
	PROPORTIONAL_SPACING                = "\x1b[26m" // ITU T.61 and T.416, not known to be used on terminals
	NOT_REVERSED                        = "\x1b[27m" // Turn blinking off
	REVEAL                              = "\x1b[28m" // Not concealed
	NOT_CROSSED_OUT                     = "\x1b[29m"
	STD_COLOR_BLACK_FOREGROUND          = "\x1b[30m"
	STD_COLOR_RED_FOREGROUND            = "\x1b[31m"
	STD_COLOR_GREEN_FOREGROUND          = "\x1b[32m"
	STD_COLOR_YELLOW_FOREGROUND         = "\x1b[33m"
	STD_COLOR_BLUE_FOREGROUND           = "\x1b[34m"
	STD_COLOR_MAGENTA_FOREGROUND        = "\x1b[35m"
	STD_COLOR_CYAN_FOREGROUND           = "\x1b[36m"
	STD_COLOR_WHITE_FOREGROUND          = "\x1b[37m"
	DEFAULT_COLOR_FOREGROUND            = "\x1b[39m" // Implementation defined (according to standard)
	STD_COLOR_BLACK_BACKGROUND          = "\x1b[40m"
	STD_COLOR_RED_BACKGROUND            = "\x1b[41m"
	STD_COLOR_GREEN_BACKGROUND          = "\x1b[42m"
	STD_COLOR_YELLOW_BACKGROUND         = "\x1b[43m"
	STD_COLOR_BLUE_BACKGROUND           = "\x1b[44m"
	STD_COLOR_MAGENTA_BACKGROUND        = "\x1b[45m"
	STD_COLOR_CYAN_BACKGROUND           = "\x1b[46m"
	STD_COLOR_WHITE_BACKGROUND          = "\x1b[47m"
	DEFAULT_COLOR_BACKGROUND            = "\x1b[49m" // Implementation defined (according to standard)
	DISABLE_PROPORTIONAL_SPACING        = "\x1b[50m" // T.61 and T.416
	FRAMED                              = "\x1b[51m" // Implemented as "emoji variation selector" in mintty.
	ENCIRCLED                           = "\x1b[52m" // Implemented as "emoji variation selector" in mintty.
	OVERLINED                           = "\x1b[53m" // Not supported in Terminal.app
	NO_FRAMED_NOR_ENCIRCLED             = "\x1b[54m"
	NOT_OVERLINED                       = "\x1b[55m"
	DEFAULT_UNDERLINE_COLOR             = "\x1b[59m" // Not in standard; implemented in Kitty, VTE, mintty, and iTerm2.
	IDEOGRAM_UNDERLINE_OR_RIGHT_LINE    = "\x1b[60m" // Rarely supported
	IDEOGRAM_DOUBLE_UNDERLINE_OR_RIGHT  = "\x1b[61m" // Rarely supported
	IDEOGRAM_OVERLINE_OR_LEFT_LINE      = "\x1b[62m" // Rarely supported
	IDEOGRAM_DOUBLE_OVERLINE_LEFT       = "\x1b[63m" // Rarely supported
	IDEOGRAM_STRESS_MARKING             = "\x1b[64m" // Rarely supported
	IDEOGRAM_RESET                      = "\x1b[65m"
	SUPERSCRIPT                         = "\x1b[73m" // Implemented only in mintty
	SUBSCRIPT                           = "\x1b[74m" // Implemented only in mintty
	NO_SUPERSCRIPT_NOR_SUBSCRIPT        = "\x1b[75m" // Implemented only in mintty
	STD_COLOR_BRIGHT_BLACK_FOREGROUND   = "\x1b[90m"
	STD_COLOR_BRIGHT_RED_FOREGROUND     = "\x1b[91m"
	STD_COLOR_BRIGHT_GREEN_FOREGROUND   = "\x1b[92m"
	STD_COLOR_BRIGHT_YELLOW_FOREGROUND  = "\x1b[93m"
	STD_COLOR_BRIGHT_BLUE_FOREGROUND    = "\x1b[94m"
	STD_COLOR_BRIGHT_MAGENTA_FOREGROUND = "\x1b[95m"
	STD_COLOR_BRIGHT_CYAN_FOREGROUND    = "\x1b[96m"
	STD_COLOR_BRIGHT_WHITE_FOREGROUND   = "\x1b[97m"
	STD_COLOR_BRIGHT_BLACK_BACKGROUND   = "\x1b[100m"
	STD_COLOR_BRIGHT_RED_BACKGROUND     = "\x1b[101m"
	STD_COLOR_BRIGHT_GREEN_BACKGROUND   = "\x1b[102m"
	STD_COLOR_BRIGHT_YELLOW_BACKGROUND  = "\x1b[103m"
	STD_COLOR_BRIGHT_BLUE_BACKGROUND    = "\x1b[104m"
	STD_COLOR_BRIGHT_MAGENTA_BACKGROUND = "\x1b[105m"
	STD_COLOR_BRIGHT_CYAN_BACKGROUND    = "\x1b[106m"
	STD_COLOR_BRIGHT_WHITE_BACKGROUND   = "\x1b[107m"
)

const (
	COLOR_BLACK                      color = iota // #000000
	COLOR_MAROON                                  // #800000
	COLOR_OFFICE_GREEN                            // #008000
	COLOR_OLIVE                                   // #808000
	COLOR_NAVY_BLUE                               // #000080
	COLOR_PURPLE                                  // #800080
	COLOR_TEAL                                    // #008080
	COLOR_SILVER                                  // #C0C0C0
	COLOR_GRAY                                    // #808080
	COLOR_RED                                     // #FF0000
	COLOR_GREEN                                   // #00FF00
	COLOR_YELLOW                                  // #FFFF00
	COLOR_BLUE                                    // #0000FF
	COLOR_MAGENTA                                 // #FF00FF
	COLOR_CYAN                                    // #00FFFF
	COLOR_WHITE                                   // #FFFFFF
	_                                             // #000000
	COLOR_DARK_NAVY_BLUE                          // #00005F
	COLOR_DARK_BLUE                               // #000087
	COLOR_ZAFFRE                                  // #0000AF
	COLOR_MEDIUM_BLUE                             // #0000D7
	_                                             // #0000FF
	COLOR_DARK_GREEN                              // #005F00
	COLOR_CARIBBEAN_CURRENT                       // #005F5F
	COLOR_SEA_BLUE                                // #005F87
	COLOR_LAPIS_LAZULI                            // #005FAF
	COLOR_TANG_BLUE                               // #005FD7
	COLOR_ULTRAMARINE_BLUE                        // #005FFF
	COLOR_IRISH_GREEN                             // #008700
	COLOR_SEA_GREEN                               // #00875F
	COLOR_DARK_CYAN                               // #008787
	COLOR_BLUE_NCS                                // #0087AF
	COLOR_GREEN_BLUE                              // #0087D7
	COLOR_BLEU_DE_FRANCE                          // #0087FF
	COLOR_ISLAMIC_GREEN                           // #00AF00
	COLOR_PIGMENT_GREEN_CMYK_GREEN                // #00AF5F
	COLOR_JUNGLE_GREEN                            // #00AF87
	COLOR_LIGHT_SEA_GREEN                         // #00AFAF
	COLOR_BRIGHT_CERULEAN                         // #00AFD7
	COLOR_DEEP_SKY_BLUE                           // #00AFFF
	COLOR_LIME                                    // #00D700
	COLOR_MALACHITE                               // #00D75F
	COLOR_AQUA_GREEN                              // #00D787
	COLOR_CARIBBEAN_GREEN                         // #00D7AF
	COLOR_DARK_TURQUOISE                          // #00D7D7
	COLOR_BRIGHT_SKY_BLUE                         // #00D7FF
	_                                             // #00FF00
	COLOR_ERIN                                    // #00FF5F
	COLOR_SPRING_GREEN                            // #00FF87
	COLOR_MEDIUM_SPRING_GREEN                     // #00FFAF
	COLOR_BRIGHT_TURQUOISE                        // #00FFD7
	_                                             // #00FFFF
	COLOR_BLOOD_RED                               // #5F0000
	COLOR_TYRIAN_PURPLE                           // #5F005F
	COLOR_INDIGO                                  // #5F0087
	COLOR_DAISY_BUSH                              // #5F00AF
	COLOR_ELECTRIC_ULTRAMARINE                    // #5F00D7
	COLOR_HAN_PURPLE_CHINESE_PURPLE               // #5F00FF
	COLOR_ANTIQUE_BRONZE                          // #5F5F00
	COLOR_STORM_DUST                              // #5F5F5F
	COLOR_PURPLE_NAVY                             // #5F5F87
	COLOR_RICH_BLUE                               // #5F5FAF
	COLOR_SLATE_BLUE                              // #5F5FD7
	COLOR_NEBULA_BLUE                             // #5F5FFF
	COLOR_OLIVE_DRAB                              // #5F8700
	COLOR_RUSSIAN_GREEN                           // #5F875F
	COLOR_STEEL_TEAL                              // #5F8787
	COLOR_AIR_FORCE_BLUE                          // #5F87AF
	COLOR_GLAUCOUS                                // #5F87D7
	COLOR_CORNFLOWER_BLUE                         // #5F87FF
	COLOR_KELLY_GREEN                             // #5FAF00
	COLOR_FERN                                    // #5FAF5F
	COLOR_SHINY_SHAMROCK                          // #5FAF87
	COLOR_VERDIGRIS                               // #5FAFAF
	COLOR_PICTON_BLUE                             // #5FAFD7
	COLOR_FRENCH_SKY_BLUE                         // #5FAFFF
	COLOR_LIME_GREEN                              // #5FD700
	COLOR_PARIS_GREEN                             // #5FD75F
	COLOR_UFO_GREEN                               // #5FD787
	COLOR_MEDIUM_AQUAMARINE                       // #5FD7AF
	COLOR_MEDIUM_TURQUOISE                        // #5FD7D7
	COLOR_VIVID_SKY_BLUE                          // #5FD7FF
	COLOR_BRIGHT_GREEN                            // #5FFF00
	COLOR_SCREAMIN_GREEN                          // #5FFF5F
	COLOR_GUPPIE_GREEN                            // #5FFF87
	COLOR_LIGHT_BLUISH_GREEN                      // #5FFFAF
	COLOR_BLUE_ZIRCON                             // #5FFFD7
	COLOR_AQUA                                    // #5FFFFF
	COLOR_DARK_RED                                // #870000
	COLOR_DARK_RASPBERRY                          // #87005F
	COLOR_MARDI_GRAS_PURPLE                       // #870087
	COLOR_GRAPE                                   // #8700AF
	COLOR_DARK_VIOLET                             // #8700D7
	COLOR_VIOLET_TRADITIONAL                      // #8700FF
	COLOR_GOLDEN_BROWN                            // #875F00
	COLOR_DEEP_TAUPE                              // #875F5F
	COLOR_FRENCH_LILAC                            // #875F87
	COLOR_DEEP_LILAC                              // #875FAF
	COLOR_MEDIUM_PURPLE                           // #875FD7
	COLOR_MEDIUM_SLATE_BLUE                       // #875FFF
	COLOR_SWAMP_GREEN                             // #878700
	COLOR_DARK_TAN                                // #87875F
	COLOR_BATTLESHIP_GRAY                         // #878787
	COLOR_WILD_BLUE_YONDER                        // #8787AF
	COLOR_PORTAGE                                 // #8787D7
	COLOR_LIGHT_SLATE_BLUE                        // #8787FF
	COLOR_APPLE_GREEN                             // #87AF00
	COLOR_OLIVINE                                 // #87AF5F
	COLOR_DARK_SEA_GREEN                          // #87AF87
	COLOR_MORNING_SKY_BLUE                        // #87AFAF
	COLOR_RUDDY_BLUE                              // #87AFD7
	COLOR_JORDY_BLUE                              // #87AFFF
	COLOR_YELLOW_GREEN                            // #87D700
	COLOR_PASTEL_GREEN                            // #87D75F
	COLOR_GOSSIP                                  // #87D787
	COLOR_ALGAE_GREEN                             // #87D7AF
	COLOR_MIDDLE_BLUE_GREEN                       // #87D7D7
	COLOR_BABY_BLUE                               // #87D7FF
	COLOR_CHARTREUSE                              // #87FF00
	COLOR_SCREAMIN_GREEN_ULTRA_GREEN              // #87FF5F
	COLOR_ULTRA_GREEN                             // #87FF87
	COLOR_BRIGHT_MINT                             // #87FFAF
	COLOR_AQUAMARINE                              // #87FFD7
	COLOR_ELECTRIC_BLUE                           // #87FFFF
	COLOR_TURKEY_RED                              // #AF0000
	COLOR_JAZZBERRY_JAM                           // #AF005F
	COLOR_FANDANGO                                // #AF0087
	COLOR_PURPLE_MUNSELL                          // #AF00AF
	COLOR_DARK_ORCHID                             // #AF00D7
	COLOR_VERONICA                                // #AF00FF
	COLOR_GINGER                                  // #AF5F00
	COLOR_MIDDLE_RED_PURPLE                       // #AF5F5F
	COLOR_PEARLY_PURPLE                           // #AF5F87
	COLOR_DEEP_FUCHSIA                            // #AF5FAF
	COLOR_RICH_LILAC                              // #AF5FD7
	COLOR_LAVENDER_INDIGO                         // #AF5FFF
	COLOR_DARK_GOLDENROD                          // #AF8700
	COLOR_LIGHT_TAUPE                             // #AF875F
	COLOR_ROSY_BROWN                              // #AF8787
	COLOR_OPERA_MAUVE                             // #AF87AF
	COLOR_LAVENDER_FLORAL                         // #AF87D7
	COLOR_TROPICAL_INDIGO                         // #AF87FF
	COLOR_OLIVE_YELLOW                            // #AFAF00
	COLOR_OLIVE_GREEN                             // #AFAF5F
	COLOR_MISTY_MOSS                              // #AFAF87
	COLOR_NOBEL                                   // #AFAFAF
	COLOR_MOON_RAKER                              // #AFAFD7
	COLOR_MAXIMUM_BLUE_PURPLE                     // #AFAFFF
	COLOR_INCHWORM                                // #AFD700
	COLOR_JUNE_BUD                                // #AFD75F
	COLOR_GRANNY_SMITH_APPLE                      // #AFD787
	COLOR_CELADON                                 // #AFD7AF
	COLOR_POWDER_BLUE                             // #AFD7D7
	COLOR_PALE_CORNFLOWER_BLUE                    // #AFD7FF
	COLOR_SPRING_BUD                              // #AFFF00
	COLOR_FRENCH_LIME                             // #AFFF5F
	COLOR_MINT_GREEN                              // #AFFF87
	COLOR_PALE_GREEN                              // #AFFFAF
	COLOR_MAGIC_MINT                              // #AFFFD7
	COLOR_CELESTE                                 // #AFFFFF
	COLOR_RACING_RED_ROSSO_CORSA                  // #D70000
	COLOR_DOGWOOD_ROSE                            // #D7005F
	COLOR_VIVID_CERISE                            // #D70087
	COLOR_BYZANTINE                               // #D700AF
	COLOR_STEEL_PINK                              // #D700D7
	COLOR_PSYCHEDELIC_PURPLE                      // #D700FF
	COLOR_COCOA_BROWN                             // #D75F00
	COLOR_INDIAN_RED                              // #D75F5F
	COLOR_CINNAMON_SATIN                          // #D75F87
	COLOR_SKY_MAGENTA                             // #D75FAF
	COLOR_ORCHID                                  // #D75FD7
	COLOR_HELIOTROPE                              // #D75FFF
	COLOR_HARVEST_GOLD                            // #D78700
	COLOR_PALE_COPPER                             // #D7875F
	COLOR_NEW_YORK_PINK                           // #D78787
	COLOR_MIDDLE_PURPLE                           // #D787AF
	COLOR_PLUM                                    // #D787D7
	COLOR_BRIGHT_LILAC                            // #D787FF
	COLOR_NEON_GOLD                               // #D7AF00
	COLOR_EARTH_YELLOW                            // #D7AF5F
	COLOR_TAN                                     // #D7AF87
	COLOR_PALE_CHESTNUT                           // #D7AFAF
	COLOR_LILAC                                   // #D7AFD7
	COLOR_MAUVE_MALLOW                            // #D7AFFF
	COLOR_PERIDOT                                 // #D7D700
	COLOR_STRAW                                   // #D7D75F
	COLOR_GREEN_EARTH_VERONA_GREEN                // #D7D787
	COLOR_PALE_SPRING_BUD                         // #D7D7AF
	COLOR_TIMBERWOLF                              // #D7D7D7
	COLOR_LAVENDER_BLUE                           // #D7D7FF
	COLOR_CHARTREUSE_YELLOW                       // #D7FF00
	COLOR_FLUORESCENT_YELLOW                      // #D7FF5F
	COLOR_KEY_LIME                                // #D7FF87
	COLOR_CANARY                                  // #D7FFAF
	COLOR_TEA_GREEN                               // #D7FFD7
	COLOR_LIGHT_CYAN                              // #D7FFFF
	_                                             // #FF0000
	COLOR_RADICAL_RED                             // #FF005F
	COLOR_ROSE                                    // #FF0087
	COLOR_HOLLYWOOD_CERISE                        // #FF00AF
	COLOR_HOT_MAGENTA                             // #FF00D7
	_                                             // #FF00FF
	COLOR_ORANGE_CRAYOLA                          // #FF5F00
	COLOR_PASTEL_RED                              // #FF5F5F
	COLOR_LIGHT_CRIMSON                           // #FF5F87
	COLOR_HOT_PINK                                // #FF5FAF
	COLOR_ROSE_PINK                               // #FF5FD7
	COLOR_FLUORESCENT_PINK                        // #FF5FFF
	COLOR_DARK_ORANGE                             // #FF8700
	COLOR_CORAL                                   // #FF875F
	COLOR_LIGHT_CORAL                             // #FF8787
	COLOR_TICKLE_ME_PINK                          // #FF87AF
	COLOR_PALE_MAGENTA                            // #FF87D7
	COLOR_FUCHSIA_PINK                            // #FF87FF
	COLOR_BRIGHT_YELLOW                           // #FFAF00
	COLOR_SANDY_BROWN                             // #FFAF5F
	COLOR_LIGHT_SALMON                            // #FFAF87
	COLOR_LIGHT_PINK                              // #FFAFAF
	COLOR_LAVENDER_PINK                           // #FFAFD7
	COLOR_ELECTRIC_LAVENDER                       // #FFAFFF
	COLOR_GOLD                                    // #FFD700
	COLOR_DANDELION                               // #FFD75F
	COLOR_MEDIUM_YELLOW                           // #FFD787
	COLOR_LIGHT_ORANGE                            // #FFD7AF
	COLOR_PALE_PINK                               // #FFD7D7
	COLOR_PINK_LACE                               // #FFD7FF
	_                                             // #FFFF00
	COLOR_LASER_LEMON                             // #FFFF5F
	COLOR_PASTEL_YELLOW                           // #FFFF87
	COLOR_LEMON_YELLOW                            // #FFFFAF
	COLOR_LIGHT_GOLDENROD_YELLOW                  // #FFFFD7
	_                                             // #FFFFFF
	COLOR_ALMOST_BLACK                            // #080808
	COLOR_SMOKY_BLACK                             // #121212
	COLOR_NERO                                    // #1C1C1C
	COLOR_EERIE_BLACK                             // #262626
	COLOR_DARK_CHARCOAL                           // #303030
	COLOR_JET_BLACK                               // #3A3A3A
	COLOR_ONYX                                    // #444444
	COLOR_MATTERHORN                              // #4E4E4E
	COLOR_DAVY_X27_S_GRAY                         // #585858
	COLOR_GRANITE_GRAY                            // #626262
	COLOR_DIM_GRAY                                // #6C6C6C
	COLOR_NICKEL                                  // #767676
	_                                             // #808080
	COLOR_ALUMINIUM                               // #8A8A8A
	COLOR_SUVA_GREY                               // #949494
	COLOR_SPANISH_GRAY                            // #9E9E9E
	COLOR_GRAY_CHATEAU                            // #A8A8A8
	COLOR_DARK_GRAY                               // #B2B2B2
	COLOR_MEDIUM_GRAY                             // #BCBCBC
	COLOR_NEON_SILVER                             // #C6C6C6
	COLOR_LIGHT_GRAY                              // #D0D0D0
	COLOR_GAINSBORO                               // #DADADA
	COLOR_PLATINUM                                // #E4E4E4
	COLOR_ANTI_FLASH_WHITE                        // #EEEEEE
)

// Get text to set specific color as foreground color.
func (c color) Foreground() string {
	return Foreground(c)
}

// Get text to set specific color as background color.
func (c color) Background() string {
	return Background(c)
}

// Get text to set specific color as underline color.
// Not in standard; implemented in Kitty, VTE, mintty, and iTerm2.
func (c color) Underline() string {
	return Underline(c)
}

// Get text to set foreground color with passed color
func Foreground(c color) string {
	return fmt.Sprintf("\x1b[38;5;%dm", int(c))
}

// Get text to set background color with passed color
func Background(c color) string {
	return fmt.Sprintf("\x1b[48;5;%dm", int(c))
}

// Get text to set underline color with passed color
// Not in standard; implemented in Kitty, VTE, mintty, and iTerm2.
func Underline(c color) string {
	return fmt.Sprintf("\x1b[58;5;%dm", int(c))
}

// Get text to set foreground color with passed RGB color. Color must be between 0 and 255 (included)
func RGBForeground(r, g, b int) (string, error) {
	if r < 0 || r > 255 || g < 0 || g > 255 || b < 0 || b > 255 {
		return "", errors.New("Color value must be between 0 and 255")
	}

	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b), nil
}

// Get text to set background color with passed RGB color. Color must be between 0 and 255 (included)
func RGBBackground(r, g, b int) (string, error) {
	if r < 0 || r > 255 || g < 0 || g > 255 || b < 0 || b > 255 {
		return "", errors.New("Color value must be between 0 and 255")
	}

	return fmt.Sprintf("\x1b[48;2;%d;%d;%dm", r, g, b), nil
}

// Get text to set underline color with passed RGB color. Color must be between 0 and 255 (included)
// Not in standard; implemented in Kitty, VTE, mintty, and iTerm2.
func RGBUnderline(r, g, b int) (string, error) {
	if r < 0 || r > 255 || g < 0 || g > 255 || b < 0 || b > 255 {
		return "", errors.New("Color value must be between 0 and 255")
	}

	return fmt.Sprintf("\x1b[58;2;%d;%d;%dm", r, g, b), nil
}

// Get demo string to see what you can do with this package, and what you console support.
func Demo() string {
	result := ""

	result = result + "\nPredefined colors:"
	for c := 0; c < 256; c++ {
		if c%8 == 0 {
			result = result + DEFAULT + "\n"
		}
		result = result + fmt.Sprintf("%s %03d", color(c).Background(), c)
	}
	result = result + DEFAULT + "\n"

	result = result + "\nColor palette:\n"
	k := 3
	for r := 0; r < 256/k; r++ {
		for gb := 0; gb < 256/k; gb++ {
			c, _ := RGBBackground(r*k, gb*k, 255-(gb*4))
			result = result + fmt.Sprintf("%s  ", c)
		}
		result = result + DEFAULT + "\n"
	}
	result = result + DEFAULT + "\n"

	return result
}
