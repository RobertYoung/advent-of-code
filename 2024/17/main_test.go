package main

import (
	"testing"
)

func equal(result []int64, i []int64) bool {
	if len(result) != len(i) {
		return false
	}

	for j := 0; j < len(result); j++ {
		if result[j] != i[j] {
			return false
		}
	}

	return true
}

func TestCreateGame(t *testing.T) {
	input := `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

	result := CreateGame(input)

	if result.register.a != 729 {
		t.Errorf("TestCreateGame().register.a = %v; want %v", result.register.a, 729)
	}

	if result.register.b != 0 {
		t.Errorf("TestCreateGame().register.b = %v; want %v", result.register.b, 0)
	}

	if result.register.c != 0 {
		t.Errorf("TestCreateGame().register.c = %v; want %v", result.register.c, 0)
	}

	if !equal(result.program, []int64{0, 1, 5, 4, 3, 0}) {
		t.Errorf("TestCreateGame().register.c = %v; want %v", result.program, []int{0, 1, 5, 4, 3, 0})
	}
}

func TestExample1(t *testing.T) {
	game := Game{
		register: Register{
			a: 0,
			b: 0,
			c: 9,
		},
		program: []int64{2, 6},
	}

	game.execute()

	if game.register.b != 1 {
		t.Errorf("TestExample1().register.b = %v; want %v", game.register.b, 3)
	}
}

func TestExample2(t *testing.T) {
	game := Game{
		register: Register{
			a: 10,
			b: 0,
			c: 0,
		},
		program: []int64{5, 0, 5, 1, 5, 4},
	}

	result := game.execute()

	if !equal(result, []int64{0, 1, 2}) {
		t.Errorf("TestExample2() = %v; want %v", result, []int64{0, 1, 2})
	}
}

func TestExample3(t *testing.T) {
	game := Game{
		register: Register{
			a: 2024,
			b: 0,
			c: 0,
		},
		program: []int64{0, 1, 5, 4, 3, 0},
	}

	result := game.execute()

	if !equal(result, []int64{4, 2, 5, 6, 7, 7, 7, 7, 3, 1, 0}) {
		t.Errorf("TestExample3() = %v; want %v", result, []int64{4, 2, 5, 6, 7, 7, 7, 7, 3, 1, 0})
	}

	if game.register.a != 0 {
		t.Errorf("TestExample3().register.a = %v; want %v", game.register.a, 0)
	}
}

func TestExample4(t *testing.T) {
	game := Game{
		register: Register{
			a: 0,
			b: 29,
			c: 0,
		},
		program: []int64{1, 7},
	}

	game.execute()

	if game.register.b != 26 {
		t.Errorf("TestExample4().register.b = %v; want %v", game.register.b, 26)
	}
}

func TestExample5(t *testing.T) {
	game := Game{
		register: Register{
			a: 0,
			b: 2024,
			c: 43690,
		},
		program: []int64{4, 0},
	}

	game.execute()

	if game.register.b != 44354 {
		t.Errorf("TestExample5().register.b = %v; want %v", game.register.b, 44354)
	}
}

func TestSample1(t *testing.T) {
	input := `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

	game := CreateGame(input)
	result := game.execute()

	if !equal(result, []int64{4, 6, 3, 5, 6, 3, 5, 2, 1, 0}) {
		t.Errorf("TestSample1() = %v; want %v", result, []int64{4, 6, 3, 5, 6, 3, 5, 2, 1, 0})
	}
}

func TestPart2Sample1(t *testing.T) {
	input := `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`

	game := CreateGame(input)
	result := game.findCopy()

	if result != 117440 {
		t.Errorf("TestPart2Sample1() = %v; want %v", result, 117440)
	}
}
