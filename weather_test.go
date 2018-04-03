package weather

import "testing"

func TestNumberNine(t *testing.T) {
	expectdigitalnine := " _ \n" +
		"|_|\n" +
		" _|"
	digitalnine := toDigital(9)
	if digitalnine != expectdigitalnine {
		t.Error("expect reruls = \n" + expectdigitalnine + " but actual is " + digitalnine)
	}
}
func TestNumberEight(t *testing.T) {
	expectdigitalnine := " _ \n" +
		"|_|\n" +
		"|_|"
	digitalnine := toDigital(8)
	if digitalnine != expectdigitalnine {
		t.Error("expect reruls = \n" + expectdigitalnine + " but actual is " + digitalnine)
	}
}

func TestNumberEightNine(t *testing.T) {
	expectdigitalnine := " _  _ \n" +
		"|_||_|\n" +
		"|_| _|"
	digitalnine := toDigital(89)
	if digitalnine != expectdigitalnine {
		t.Error("expect reruls = \n" + expectdigitalnine + " but actual is " + digitalnine)
	}
}
func TestNumberEightNineNine(t *testing.T) {
	expectdigitalnine := " _  _  _ \n" +
		"|_||_||_|\n" +
		"|_| _| _|"
	digitalnine := toDigital(899)
	if digitalnine != expectdigitalnine {
		t.Error("expect reruls = \n" + expectdigitalnine + " but actual is " + digitalnine)
	}
}
