package main

import "testing"

func TestFirstRow(t *testing.T) {
	row := "FBFBBFF"

	result := FindRow(row)

	if result != 44 {
		t.Fatal("Row should be 44 got", result)
	}
}

func TestSecondRow(t *testing.T) {
	row := "BFFFBBF"

	result := FindRow(row)

	if result != 70 {
		t.Fatal("Row should be 70 got", result)
	}

}

func TestThirdRow(t *testing.T) {
	row := "FFFBBBF"

	result := FindRow(row)

	if result != 14 {
		t.Fatal("Row should be 14 got", result)
	}

}

func TestFourthRow(t *testing.T) {
	row := "BBFFBBF"

	result := FindRow(row)

	if result != 102 {
		t.Fatal("Row should be 102 got", result)
	}

}
