package main

import (
	"fmt"
	"os"
	"sort"
	"testing"
)

func setup() {
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}

func Test1(t *testing.T) {
	fmt.Println("I'm test1")
	ints := []int{1, 2, 4, 3}
	sort.Slice(ints, func(i, j int) bool {
		if ints[i] > ints[j] {
			return false
		}
		return true
	})
	t.Log(ints)

	t.Log("test1")
	t.Log("test2")
	t.Log("test3")
	t.Log("test5")
	t.Log("test4")
<<<<<<< HEAD
	t.Log("test8")
	t.Log("test6")
	t.Log("test7")
}

func Test2(t *testing.T) {
	fmt.Println("I'm test2")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
