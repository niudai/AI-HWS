package p2

import (
	. "fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var state [][]int

type Assign struct {
	x      int
	y      int
	number int
}

const WIDTH = 9

func P2() string {
	Println("P2 IS WORKING...")
	return "p2 package"
}

func TrivialSearch() error {
	initA := Assign{x: -1, y: WIDTH - 1}
	initA, _ = nextAssign(initA)
	for i := 1; i <= WIDTH; i++ {
		initA.number = i
		if _Search(initA) {
			return Errorf("Solution Not Found")
			// search success, print state
		} else {
			// Printf("init number %d failed. \n%v\n", i, SerializeState(state))
		}
	}
	return nil
}

// Untested
// for every position, try every value iteratively, if consistent, assign it, try next position,
// until reach the last element, goal reached, or failed (inconsistent), need recover.
// do it recursively
// search receives the position to assign, and value to assign
// if failed, recover, return false, let the caller know it failed.
// caller try next available value
// if all value are tried, return false
// if reaches the last element, and consistent, return true, let the caller know it return true
func _Search(a Assign) bool {
	//
	if isConsistent(state, a) {
		// if consistent, assign it.
		assign(state, a)
		nextA, err := nextAssign(a)
		if err != nil {
			// err != nil indicates there is not next assign, so return true
			return true
		}
		for k := 1; k <= WIDTH; k++ {
			nextA.number = k
			if _Search(nextA) {
				return true
				// search success, just return true to propagate the success
			}
		}
		// every chance not ok, indicate current assign is not ok, so unassign to recover
		unAssign(state, a)
		// if tried all val, all failed, just return false to propagate the failure to the upper caller.
		return false
	} else {
		// not consistent, should recover from current assign
		// unAssign(state, a)
		return false
	}
}

// tested
// skip the assigned value
func nextAssign(assign Assign) (Assign, error) {
	cov := assign.x*WIDTH + assign.y
	for i := cov + 1; i < WIDTH*WIDTH; i++ {
		// iterate from cov to the last one
		if state[i/WIDTH][i%WIDTH] == 0 {
			// find not assigned val
			return Assign{x: i / WIDTH, y: i % WIDTH, number: 0}, nil
		}
	}
	return Assign{}, Errorf("reaches the last position!")
}

// assign value to the state
func assign(state [][]int, assign Assign) {
	if state[assign.x][assign.y] != 0 {
		panic(Sprintf("cannot assign to an already assign val"))
	}
	state[assign.x][assign.y] = assign.number
}

// Tested
func unAssign(state [][]int, assign Assign) {
	if state[assign.x][assign.y] == 0 {
		panic(Sprintf("cannot unassign to an not assigned val. A: %v\n%v", assign, SerializeState(state)))
	}
	state[assign.x][assign.y] = 0
}

// Tested
func isConsistent(state [][]int, assign Assign) bool {
	if state[assign.x][assign.y] != 0 {
		panic(Sprintf("%+v Cannot assign to an already assigned position\n%v", assign, SerializeState(state)))
	}
	// check row
	for i := 0; i < WIDTH; i++ {
		// Println(state[assign.x][i])
		if state[assign.x][i] != 0 {
			if state[assign.x][i] == assign.number {
				return false
			}
		}
	}
	// check column
	for i := 0; i < WIDTH; i++ {
		// Println(state[i][assign.y])
		if state[i][assign.y] != 0 {
			if state[i][assign.y] == assign.number {
				return false
			}
		}
	}
	// check nine block consistency
	// TODO
	blockX, blockY := assign.x-assign.x%3, assign.y-assign.y%3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			v := state[blockX+i][blockY+j]
			// Println(v)
			if v != 0 && v == assign.number {
				return false
			}
		}
	}

	return true
}

func SerializeState(state [][]int) string {

	str := ""
	for _, v := range state {
		str = str + Sprintln(v)
	}
	return str
}

func Exe() {
	f, _ := os.Open("p2.txt")
	resultF, _ := os.OpenFile("result.txt", os.O_RDWR, 0755)
	defer f.Close()
	defer resultF.Close()
	buf := make([]byte, 300)
	n, err := f.Read(buf)
	if err != nil {
		Println(err)
	}
	str := string(buf[:n])
	str = strings.Replace(str, "\n", " ", -1)
	var numStr []string = strings.Split(str, " ")
	state = make([][]int, WIDTH)
	for i, _ := range state {
		state[i] = make([]int, WIDTH)
	}
	for i, s := range numStr {
		n, _ := strconv.ParseInt(s, 10, 32)
		state[i/WIDTH][i%WIDTH] = int(n)
	}
	Println(SerializeState(state))
	// initA := Assign{x: 8, y: 7}
	// nxt, _ := nextAssign(initA)
	// unAssign(state, initA)
	start := time.Now()
	TrivialSearch()
	duration := time.Now().Sub(start)
	Println(SerializeState(state))
	Printf("Trivial Search Cost %v\n", duration)
	// ass := Assign{
	// 	number: 6,
	// 	x:      4,
	// 	y:      4,
	// }
	// Println(isConsistent(state, ass))
	// Search()
}
