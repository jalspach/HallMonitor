package leds

// returns status of the LED's in bianry
// want to know whether an LED is on, off or flashing
// binary (LSB) Red, Yellow Green, Blink
// 0 all off, 1 Red solid, 2 yellow solid, 3 red and yellow solid, 4 green solid, 5 red and green solid, 6 green and yellow solid, 7 red yellow and green solid
// 8 NA (All off and flashing), 9 Red Flashing, A yellow flashing, B red and yellow flashing, C green flashing, D Green and Red flashing, E green and yellow flashing, green yellow and red flashing

func checkstatus() int {
	//returns the status of the LED's
	return 0
}
// Set the status of the LED's based on the the Hex value. If an invalid number is passed will return 0X_BEEF if everything worked it will return 0x_FFFF
func setstatus(hex int) int {
	// Set the status of the leds
	switch hex {
	case 0x0:
		//all off
	case 0x1:
		//Red solid
	case 0x2:
		//Yellow solid
	case 0x3:
		//Red and Yellow solid
	case 0x4:
		//Green solid
	case 0x5:
		//Red and Green solid
	case 0x6:
		//Yellow and Green solid
	case 0x7:
		//Red, Green and Yellow solid
	case 0x8:
		//Not Valid (all off but still flashing)
	case 0x9:
		//Red flashing
	case 0xA:
		//Yellow flashing
	case 0xB:
		//Red and Yellow flashing
	case 0xC:
		//Green flashing
	case 0xD:
		//Red Green flashing
	case 0xE:
		//Yellow and Green flashing
	case 0xF:
		//Red, Yellow and Green flashing
	default:
		return 0xBEEF
	}
	return 0xFFFF
}
func setthinking() int {
	//build these to call setstatus with a specific hex value instead of setting LED's manually
	// set alert, set bootup, etc...
	return 0
}
