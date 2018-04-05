package weather

import (
	"strconv"
)

var top = [10]string{" _ ", "   ", " _ ", " _ ", "   ", " _ ", " _ ", " _ ", " _ ", " _ "}
var mid = [10]string{"| |", "  |", " _|", " _|", "|_|", "|_ ", "|_ ", "  |", "|_|", "|_|"}
var bot = [10]string{"|_|", "  |", "|_ ", " _|", "  |", " _|", "|_|", "  |", "|_|", " _|"}
var fourtyeight = 48

func toDigital(number int) string {
	str := strconv.Itoa(number)
	palldigit := finddigit(str)
	return palldigit
}

func finddigit(str string) string {
	var digit int
	var alltop string
	var allmid string
	var allbot string

	for i := range str {
		digit = int(str[i])
		alltop += top[digit-fourtyeight]
		allmid += mid[digit-fourtyeight]
		allbot += bot[digit-fourtyeight]
	}
	alldigit := alltop + "\n" + allmid + "\n" + allbot
	return alldigit
}
