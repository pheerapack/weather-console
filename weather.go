package weather

import (
	"strconv"
)

func toDigital(number int) string {
	top := [10]string{" _ ", "   ", " _ ", " _ ", "   ", " _ ", " _ ", " _ ", " _ ", " _ "}
	mid := [10]string{"| |", "  |", " _|", " _|", "|_|", "|_ ", "|_ ", "  |", "|_|", "|_|"}
	bot := [10]string{"|_|", "  |", "|_ ", " _|", "  |", " _|", "|_|", "  |", "|_|", " _|"}
	str := strconv.Itoa(number)

	var digit int
	var alltop string
	var allmid string
	var allbot string
	fourtyeight := 48
	for i := range str {
		digit = int(str[i])
		alltop = alltop + top[digit-fourtyeight]
		allmid = allmid + mid[digit-fourtyeight]
		allbot = allbot + bot[digit-fourtyeight]
	}
	return alltop + "\n" + allmid + "\n" + allbot

}
