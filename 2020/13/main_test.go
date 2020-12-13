package main

import "testing"

func TestFoldPair(t *testing.T) {
	a := &Bus{0, 2}
	b := &Bus{1, 5}
	newBus := FoldPair(a, b)
	if newBus.FirstTime != 4 {
		t.Fatal("Should be 4 got", newBus.FirstTime)
	}
	if newBus.RepeatsEvery != 10 {
		t.Fatal("Should be 10 got", newBus.RepeatsEvery)
	}
}

func TestFoldPairLarger(t *testing.T) {
	a := &Bus{0, 42}
	b := &Bus{6, 13}
	newBus := FoldPair(a, b)
	if newBus.FirstTime != 462 {
		t.Fatal("Should be 462 got", newBus.FirstTime)
	}
	if newBus.RepeatsEvery != 546 {
		t.Fatal("Should be 546 got", newBus.RepeatsEvery)
	}
}

func TestFoldPairSame(t *testing.T) {
	a := &Bus{0, 42}
	b := &Bus{0, 42}
	newBus := FoldPair(a, b)
	if newBus.FirstTime != 0 {
		t.Fatal("Should be 0 got", newBus.FirstTime)
	}
	if newBus.RepeatsEvery != 42 {
		t.Fatal("Should be 42 got", newBus.RepeatsEvery)
	}
}

func TestFoldPairBuss(t *testing.T) {
	a := &Bus{0, 42}
	b := &Bus{6, 3}
	newBus := FoldPair(a, b)
	if newBus.FirstTime != 0 {
		t.Fatal("Should be 0 got", newBus.FirstTime)
	}
	if newBus.RepeatsEvery != 42 {
		t.Fatal("Should be 42 got", newBus.RepeatsEvery)
	}
}

func TestFoldTwice(t *testing.T) {
	a := &Bus{0, 2}
	b := &Bus{1, 5}
	c := &Bus{5, 17}

	interrim := FoldPair(a, b)
	result := FoldPair(interrim, c)

	if result.FirstTime%2 != 0 {
		t.Fatal("Should be modulo 2", result.FirstTime, result.FirstTime%2)
	}

	if (result.FirstTime+1)%5 != 0 {
		t.Fatal("result.FirstTime+1 should be modulo 5:", result.FirstTime, result.FirstTime%5)
	}
	if (result.FirstTime+5)%17 != 0 {
		t.Fatal("result.FirstTime+5 should be modulo 17:", result.FirstTime, result.FirstTime%17)
	}

}

func TestWithInput1(t *testing.T) {
	_, buses := Parse([]string{"", "7,13,x,x,59,x,31,19"})
	expected := 1068781
	out := Part2(buses)

	if out != expected {
		t.Fatalf("Expected %v got %v", expected, out)
	}
}

func TestWithInput2(t *testing.T) {
	_, buses := Parse([]string{"", "17,x,13,19"})

	expected := 3417
	out := Part2(buses)

	if out != expected {
		t.Fatalf("Expected %v got %v", expected, out)
	}
}

func TestWithInput3(t *testing.T) {
	_, buses := Parse([]string{"", "67,7,59,61"})
	expected := 754018
	out := Part2(buses)

	if out != expected {
		t.Fatalf("Expected %v got %v", expected, out)
	}
}

func TestWithInput4(t *testing.T) {
	_, buses := Parse([]string{"", "67,x,7,59,61"})
	expected := 779210
	out := Part2(buses)

	if out != expected {
		t.Fatalf("Expected %v got %v", expected, out)
	}
}

func TestWithInput5(t *testing.T) {
	_, buses := Parse([]string{"", "67,7,x,59,61"})
	expected := 1261476
	out := Part2(buses)

	if out != expected {
		t.Fatalf("Expected %v got %v", expected, out)
	}
}

func TestWithInput6(t *testing.T) {
	_, buses := Parse([]string{"", "1789,37,47,1889"})
	expected := 1202161486
	out := Part2(buses)

	if out != expected {
		t.Fatalf("Expected %v got %v", expected, out)
	}
}
