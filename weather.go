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
	for i := range str {
		ten = int(str[i])
		ptop = ptop + top[ten-48]
		pmid = pmid + mid[ten-48]
		pbot = pbot + bot[ten-48]
	}
	return ptop + "\n" + pmid + "\n" + pbot

}
