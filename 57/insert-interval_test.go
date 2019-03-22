package leetcode

import "testing"

func TestNormal(t *testing.T) {
	input := []Interval{{1, 3}}
	insertInterval := Interval{1, 3}
	output := insert(input, insertInterval)
	expect := []Interval{{1, 3}}
	if equal(expect, output) == false {
		t.Fail()
	}

}

func TestNormal1(t *testing.T) {
	input := []Interval{
		{1, 3},
		{4, 5},
		{5, 7},
		{9, 10}}
	insertInterval := Interval{1, 4}
	output := insert(input, insertInterval)
	expect := []Interval{
		{1, 5},
		{5, 7},
		{9, 10},
	}
	if equal(expect, output) == false {
		t.Fail()
	}

}

func TestNormal2(t *testing.T) {
	input := []Interval{
		{1, 3},
		{6, 9}}
	insertInterval := Interval{2, 5}
	output := insert(input, insertInterval)
	expect := []Interval{
		{1, 5},
		{6, 9},
	}
	if equal(expect, output) == false {
		t.Fail()
	}

}

func TestNormal3(t *testing.T) {
	input := []Interval{}
	insertInterval := Interval{5, 7}
	output := insert(input, insertInterval)
	expect := []Interval{
		{5, 7},
	}
	if equal(expect, output) == false {
		t.Fail()
	}

}

func TestNormal4(t *testing.T) {
	input := []Interval{
		{1, 5},
	}
	insertInterval := Interval{2, 7}
	output := insert(input, insertInterval)
	expect := []Interval{
		{1, 7},
	}
	if equal(expect, output) == false {
		t.Fail()
	}

}

func TestNormal5(t *testing.T) {
	input := []Interval{
		{1, 5},
	}
	insertInterval := Interval{6, 7}
	output := insert(input, insertInterval)
	expect := []Interval{
		{1, 5},
		{6, 7},
	}
	if equal(expect, output) == false {
		t.Fail()
	}

}

func TestNormal6(t *testing.T) {
	input := []Interval{
		{1, 5},
	}
	insertInterval := Interval{0, 0}
	output := insert(input, insertInterval)
	expect := []Interval{
		{0, 0},
		{1, 5},
	}
	if equal(expect, output) == false {
		t.Fail()
	}

}

func TestNormal7(t *testing.T) {
	input := []Interval{
		{1, 5},
		{13, 15},
	}
	insertInterval := Interval{7, 7}
	output := insert(input, insertInterval)
	expect := []Interval{
		{1, 5},
		{7, 7},
		{13, 15},
	}
	if equal(expect, output) == false {
		t.Fail()
	}

}

func equal(a, b []Interval) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}
