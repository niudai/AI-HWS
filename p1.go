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
	sevenPos := findSevenAjacentPos(s)
	states = append(states, sevenPos...)
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
func findSevenAjacentPos(state *State) []*State {
	seven := state.numberPos[7]
	ajacentPos := make([]*State, 0)
	// go left
	if seven.x-2 >= 0 {
		l1 := Position{x: seven.x - 1, y: seven.y}
		l2 := Position{x: seven.x - 2, y: seven.y - 1}
		if state.posNumber[l1] == 0 && state.posNumber[l2] == 0 {
			c := state.Clone()
			// update new 7
			c.posNumber[l1], c.posNumber[l2] = 7, 7
			// update new 0
			c.posNumber[seven], c.posNumber[Position{seven.x, seven.y - 1}] = 0, 0
			// update new 7 position
			c.numberPos[7] = Position{seven.x - 1, seven.y}
			// update uniformCost
			c.uniformCost++
			// update zero pos
			c.zeroPos = []Position{Position{seven.x, seven.y}, Position{seven.x, seven.y - 1}}
			ajacentPos = append(ajacentPos, c)
		}
	}
	// go right
	if seven.x+1 < Width {
		r1 := Position{x: seven.x + 1, y: seven.y}
		r2 := Position{x: seven.x + 1, y: seven.y - 1}
		if state.posNumber[r1] == 0 && state.posNumber[r2] == 0 {
			c := state.Clone()
			// update new 7
			c.posNumber[r1], c.posNumber[r2] = 7, 7
			// update new 0
			c.posNumber[seven], c.posNumber[Position{seven.x - 1, seven.y - 1}] = 0, 0
			// update new 7 position
			c.numberPos[7] = Position{seven.x + 1, seven.y}
			// update uniformCost
			c.uniformCost++
			// update zero pos
			c.zeroPos = []Position{Position{seven.x, seven.y}, Position{seven.x - 1, seven.y - 1}}
			ajacentPos = append(ajacentPos, c)
		}
	}
	// go up
	if seven.y-2 >= 0 {
		u1 := Position{x: seven.x, y: seven.y - 2}
		u2 := Position{x: seven.x - 1, y: seven.y - 2}
		if state.posNumber[u1] == 0 && state.posNumber[u2] == 0 {
			c := state.Clone()
			// update new 7
			c.posNumber[u1], c.posNumber[u2] = 7, 7
			// update new 0
			c.posNumber[seven], c.posNumber[Position{seven.x - 1, seven.y - 1}] = 0, 0
			// update new 7 position
			c.numberPos[7] = Position{seven.x, seven.y - 1}
			// update uniformCost
			c.uniformCost++
			// update zero pos
			c.zeroPos = []Position{Position{seven.x, seven.y}, Position{seven.x - 1, seven.y - 1}}
			ajacentPos = append(ajacentPos, c)
		}
	}
	// go down
	if seven.y+1 < Width {
		d1 := Position{x: seven.x - 1, y: seven.y}
		d2 := Position{x: seven.x, y: seven.y + 1}
		if state.posNumber[d1] == 0 && state.posNumber[d2] == 0 {
			c := state.Clone()
			// update new 7
			c.posNumber[d1], c.posNumber[d2] = 7, 7
			// update new 0
			c.posNumber[Position{seven.x - 1, seven.y - 1}], c.posNumber[Position{seven.x, seven.y - 1}] = 0, 0
			// update new 7 position
			c.numberPos[7] = Position{seven.x, seven.y + 1}
			// update uniformCost
			c.uniformCost++
			// update zero pos
			c.zeroPos = []Position{Position{seven.x - 1, seven.y - 1}, Position{seven.x, seven.y - 1}}
			ajacentPos = append(ajacentPos, c)
		}
	}
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
