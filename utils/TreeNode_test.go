package utils

import "testing"

import "fmt"

func TestNotEqual(t *testing.T) {
	l := &TreeNode{}
	l.Val = 100
	l.Left = &TreeNode{}
	l.Left.Val = 200

	r := &TreeNode{}
	r.Val = 200
	r.Right = &TreeNode{}
	r.Right.Val = 300
	if l.Equal(r) {
		t.Fail()
	}
}

func TestEqual(t *testing.T) {
	l := &TreeNode{}
	l.Val = 100
	l.Left = &TreeNode{}
	l.Left.Val = 200

	r := &TreeNode{}
	r.Val = 100
	r.Left = &TreeNode{}
	r.Left.Val = 200
	if !l.Equal(r) {
		t.Fail()
	}
}

func TestEqual2(t *testing.T) {
	l := &TreeNode{}
	l.Val = 100
	l.Left = &TreeNode{}
	l.Left.Val = 200
	l.Right = &TreeNode{}
	l.Right.Val = 300

	r := &TreeNode{}
	r.Val = 100
	r.Left = &TreeNode{}
	r.Left.Val = 200
	r.Right = &TreeNode{}
	r.Right.Val = 300
	if !l.Equal(r) {
		t.Fail()
	}
}

func TestGenerate(t *testing.T) {
	vals := []interface{}{1, 2, 3, 4}
	root := NewTree(vals)

	d := Description(root)
	if !checkEqualTreeValsArray(vals, d) {
		t.Fail()
	}
	fmt.Println(d)
}

func TestDescription1(t *testing.T) {
	l := &TreeNode{}
	l.Val = 100
	l.Left = &TreeNode{}
	l.Left.Val = 200
	l.Right = &TreeNode{}
	l.Right.Val = 300

	d := Description(l)
	fmt.Println(d)
	expect := []interface{}{100, 200, 300}
	if !checkEqualTreeValsArray(d, expect) {
		t.Fail()
	}
}

func TestDescription2(t *testing.T) {
	l := &TreeNode{}
	l.Val = 100
	l.Right = &TreeNode{}
	l.Right.Val = 200
	l.Right.Right = &TreeNode{}
	l.Right.Right.Val = 300

	d := Description(l)
	fmt.Println(d)
	expect := []interface{}{100, null, 200, null, null, null, 300}
	if !checkEqualTreeValsArray(d, expect) {
		t.Fail()
	}
}

func TestDescription3(t *testing.T) {
	l := &TreeNode{}
	l.Val = 100
	l.Right = &TreeNode{}
	l.Right.Val = 200
	l.Right.Left = &TreeNode{}
	l.Right.Left.Val = 300

	d := Description(l)
	fmt.Println(d)
	expect := []interface{}{100, null, 200, null, null, 300}
	if !checkEqualTreeValsArray(d, expect) {
		t.Fail()
	}
}
