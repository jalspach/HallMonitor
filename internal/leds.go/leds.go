package leds

import "encoding/hex"

// returns status of the LED's in bianry
// want to know whether an LED is on, off or flashing
// (LSB) Red, Yellow Green
// 0 all off, 1 Red solid, 2 yellow solid, 3 red and yellow solid, 4 green solid, 5 red and green solid, 6 green and yellow solid, 7 red yellow and green solid
// 8 NA (All off and flashing), 9 Red Flashing, A yellow flashing, B red and yellow flashing, C green flashing, D Green and Red flashing, E green and yellow flashing, green yellow and red flashing
func checkstatus() hex {
	//returns the status of the LED's
}
func setstatus(i hex) int {
	// Set the status of the leds
switch i {
case 0:
	//all off
case 1:
	//Red solid
case 2:
	//Yellow solid
case 3:
	//Red and Yellow solid
case 4:
	//Green solid
case 5:
	//Red and Green solid
case 6:
	//Yellow and Green solid
case 7:
	//Red, Green and Yellow solid
case 8:
	//Not Valid (all off but still flashing)
case 9:
	//Red flashing
case A:
	//Yellow flashing
case B:
	//Red and Yellow flashing
case C:
	//Green flashing	
case D:
	//Red Green flashing
case E:
	//Yellow and Green flashing
case F:
	//Red, Yellow and Green flashing								
default:
	return 1	
}	


}
func setthinking() int {
	//build these to call setstatus with a specific hex value instead of setting LED's manually
	// set alert, set bootup, etc...
}
