package main

import (
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var goalState *State

func isSameState(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func buildState(nums []int) *State {
	state := State{numberPos: make(map[int]Position), heuristicCost: 0, posNumber: make(map[Position]int)}
	state.zeroPos = make([]Position, 0)
	for i, v := range nums {
		if v != 0 {
			state.posNumber[Position{x: i % Width, y: i / Width}] = v
			state.numberPos[v] = Position{x: i % Width, y: i / Width}
		} else {
			state.zeroPos = append(state.zeroPos, Position{x: i % Width, y: i / Width})
		}
	}
	return &state
}

// TODO: not tested yet
func findNextStates(s *State) []*State {
	states := make([]*State, 0)
	// not 7 part:
	for _, p := range s.zeroPos {
		// for every blank space
		pos := findAjacentPos(p)
		for _, p_ := range pos {
			// for every ajacent space
			// p is blank, p_ is candidate
			num := s.posNumber[p_]
			// if p is not 7 or blank, exchange happens
			if num != 7 && num != 0 {
				c := s.Clone()
				c.posNumber[p] = num
				c.posNumber[p_] = 0
				c.numberPos[num] = p
				// delete zero pos
				for i, p2 := range c.zeroPos {
					if p2 == p {
						c.zeroPos = append(c.zeroPos[:i], append(c.zeroPos[i+1:], p_)...)
					}
				}
				c.uniformCost++
				c.heuristicCost = HeuristicCost(goalState, c)
				states = append(states, c)
			}
		}
	}
	// 7 part:
	// sevenPos := findSevenAjacentPos(s.numberPos[7])
	// var isValid bool = true
	// s.posNumber[]
	return states
}

func findAjacentPos(p Position) []Position {
	pos := make([]Position, 0)
	if p.x > 0 {
		pos = append(pos, Position{p.x - 1, p.y})
	}
	if p.y > 0 {
		pos = append(pos, Position{p.x, p.y - 1})
	}
	if p.x < Width-1 {
		pos = append(pos, Position{p.x + 1, p.y})
	}
	if p.y < Width-1 {
		pos = append(pos, Position{p.x, p.y + 1})
	}
	return pos
}

// TODO: not tested
func findSevenAjacentPos(seven Position) [][]Position {
	ajacentPos := make([][]Position, 0)
	l1 := Position{x: seven.x - 1, y: seven.y}
	l2 := Position{x: seven.x - 2, y: seven.y - 1}
	lc := append(make([]Position, 0), l1, l2)
	ajacentPos = append(ajacentPos, lc)
	r1 := Position{x: seven.x + 1, y: seven.y}
	r2 := Position{x: seven.x + 1, y: seven.y - 1}
	rc := append(make([]Position, 0), r1, r2)
	ajacentPos = append(ajacentPos, rc)
	u1 := Position{x: seven.x - 1, y: seven.y - 2}
	u2 := Position{x: seven.x - 1, y: seven.y - 2}
	uc := append(make([]Position, 0), u1, u2)
	ajacentPos = append(ajacentPos, uc)
	d1 := Position{x: seven.x - 1, y: seven.y}
	d2 := Position{x: seven.x, y: seven.y + 1}
	dc := append(make([]Position, 0), d1, d2)
	ajacentPos = append(ajacentPos, dc)
	return ajacentPos
}

func main() {
	goal := []int{1, 2, 3, 4, 5, 7, 7, 8, 9, 10, 6, 7, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}

	f, _ := os.Open("p1.txt")
	defer f.Close()
	buf := make([]byte, 100)
	n, err := f.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	str := string(buf[:n])
	str = strings.Replace(str, "\n", ",", -1)
	var numStr []string = strings.Split(str, ",")
	var num []int = make([]int, len(numStr))
	for i, s := range numStr {
		n, _ := strconv.ParseInt(s, 10, 32)
		num[i] = int(n)
	}
	// fmt.Println(goal)
	// fmt.Println(num)
	l := list.New()
	l.PushBack(goal)
	l.PushBack(num)
	initialState := buildState(num)
	goalState = buildState(goal)
	initialState.heuristicCost = HeuristicCost(initialState, goalState)
	initialState.uniformCost = 0
	// fmt.Println(isSameState(num, num))
	// fmt.Println(initialState.numberPos, "\n", initialState.heuristicCost, "\n", initialState.zeroPos)
	// fmt.Printf("%+v\n", findNextStates(initialState))
	fmt.Println(initialState.Serilize())
	for _, v := range findNextStates(initialState) {
		fmt.Printf("%+v\n", v.Serilize())
	}
}
