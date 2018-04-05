package weather

import (
	"strconv"
)

func toDigital(number int) string {
	top := [10]string{" _ ", "   ", " _ ", " _ ", "   ", " _ ", " _ ", " _ ", " _ ", " _ "}
	mid := [10]string{"| |", "  |", " _|", " _|", "|_|", "|_ ", "|_ ", "  |", "|_|", "|_|"}
	bot := [10]string{"|_|", "  |", "|_ ", " _|", "  |", " _|", "|_|", "  |", "|_|", " _|"}
	str := strconv.Itoa(number)

	var ten int
	var ptop string
	var pmid string
	var pbot string
	fourtyeight := 48
	for i := range str {
		ten = int(str[i])
		ptop = ptop + top[ten-fourtyeight]
		pmid = pmid + mid[ten-fourtyeight]
		pbot = pbot + bot[ten-fourtyeight]
	}
	return ptop + "\n" + pmid + "\n" + pbot

}
