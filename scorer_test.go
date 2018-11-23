package main

import "testing"

func TestScoreSimple(t *testing.T) {
	game := []string{"01", "11"}
	result := Score(game)
	if result != 3 {
		t.Errorf("Score should be 3 but was %v", result)
	}
}

func TestScoreSingleSpare(t *testing.T) {
	game := []string{"01", "1/", "53", "40"}
	result := Score(game)
	if result != 28 {
		t.Errorf("Score should be 28 but was %v", result)
	}
}

func TestScoreSingleFinalSpare(t *testing.T) {
	game := []string{"01", "53", "40", "2/", "5"}
	result := Score(game)
	if result != 33 {
		t.Errorf("Score should be 33 but was %v", result)
	}
}

func TestScoreSingleStrike(t *testing.T) {
	game := []string{"X", "12", "34"}
	result := Score(game)
	if result != 23 {
		t.Errorf("Score should be 23 but was %v", result)
	}
}

func TestScoreAllStrikes(t *testing.T) {
	game := []string{"X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X"}
	result := Score(game)
	if result != 300 {
		t.Errorf("Score should be 300 but was %v", result)
	}
}

func TestScoreAllSpares(t *testing.T) {
	game := []string{"5/", "5/", "5/", "5/", "5/", "5/", "5/", "5/", "5/", "5/", "5"}
	result := Score(game)
	if result != 150 {
		t.Errorf("Score should be 150 but was %v", result)
	}
}

func TestScoreAllNines(t *testing.T) {
	game := []string{"9-", "9-", "9-", "9-", "9-", "9-", "9-", "9-", "9-", "9-"}
	result := Score(game)
	if result != 90 {
		t.Errorf("Score should be 90 but was %v", result)
	}
}

func TestScoreMixed(t *testing.T) {
	game := []string{"X", "7/", "9-", "X", "-8", "8/", "-6", "X", "X", "X", "81"}
	result := Score(game)
	if result != 167 {
		t.Errorf("Score should be 167 but was %v", result)
	}
}
