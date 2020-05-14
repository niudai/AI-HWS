package main

import (
	"container/heap"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"ai.com/exp2/p2"
)

var goalState *State
var closeStates States
var openStates States

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
	state := State{numberPos: make(map[int]Position), heuristicCost: 0, uniformCost: 0, posNumber: make(map[Position]int)}
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
func (s *State) findNextStates() []*State {
	states := make([]*State, 0)
	// not 7 part:
	for _, p := range s.zeroPos {
		// for every blank space
		pos := findAjacentPos(s, p)
		for _, p_ := range pos {
			// for every ajacent space
			// p is blank, p_ is candidate
			num := s.posNumber[p_.Position]
			// if p is not 7 or blank, exchange happens
			if num != 7 && num != 0 {
				c := s.Clone()
				c.parentState = s

				c.posNumber[p] = num
				c.posNumber[p_.Position] = 0
				c.numberPos[num] = p
				c.transition = p_.Transition
				// delete zero pos
				for i, p2 := range c.zeroPos {
					if p2 == p {
						c.zeroPos = append(c.zeroPos[:i], append(c.zeroPos[i+1:], p_.Position)...)
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

func findAjacentPos(s *State, p Position) []PositionWithT {
	pos := make([]PositionWithT, 0)
	var des Position
	if p.x > 0 {
		des = Position{p.x - 1, p.y}
		pos = append(pos, PositionWithT{des, Transition{number: s.posNumber[des], direction: RIGHT}})
	}
	if p.y > 0 {
		des = Position{p.x, p.y - 1}
		pos = append(pos, PositionWithT{des, Transition{number: s.posNumber[des], direction: DOWN}})
	}
	if p.x < Width-1 {
		des = Position{p.x + 1, p.y}
		pos = append(pos, PositionWithT{des, Transition{number: s.posNumber[des], direction: LEFT}})
	}
	if p.y < Width-1 {
		des = Position{p.x, p.y + 1}
		pos = append(pos, PositionWithT{des, Transition{number: s.posNumber[des], direction: UP}})
	}
	return pos
}

// TODO: not tested
func findSevenAjacentPos(state *State) []*State {
	seven := state.numberPos[7]
	ajacentPos := make([]*State, 0)
	var c *State = state.Clone()
	c.parentState = state
	// go left
	if seven.x-2 >= 0 {
		l1 := Position{x: seven.x - 1, y: seven.y}
		l2 := Position{x: seven.x - 2, y: seven.y - 1}
		if state.posNumber[l1] == 0 && state.posNumber[l2] == 0 {
			// update new 7
			c.posNumber[l1], c.posNumber[l2] = 7, 7
			// update new 0
			c.posNumber[seven], c.posNumber[Position{seven.x, seven.y - 1}] = 0, 0
			// update new 7 position
			c.numberPos[7] = Position{seven.x - 1, seven.y}
			// update uniformCost
			c.uniformCost++
			c.transition = Transition{number: 7, direction: LEFT}
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
			// update new 7
			c.posNumber[r1], c.posNumber[r2] = 7, 7
			// update new 0
			c.posNumber[seven], c.posNumber[Position{seven.x - 1, seven.y - 1}] = 0, 0
			// update new 7 position
			c.numberPos[7] = Position{seven.x + 1, seven.y}
			// update uniformCost
			c.uniformCost++
			// update zero pos
			c.transition = Transition{number: 7, direction: RIGHT}
			c.zeroPos = []Position{Position{seven.x, seven.y}, Position{seven.x - 1, seven.y - 1}}
			ajacentPos = append(ajacentPos, c)
		}
	}
	// go up
	if seven.y-2 >= 0 {
		u1 := Position{x: seven.x, y: seven.y - 2}
		u2 := Position{x: seven.x - 1, y: seven.y - 2}
		if state.posNumber[u1] == 0 && state.posNumber[u2] == 0 {
			// update new 7
			c.posNumber[u1], c.posNumber[u2] = 7, 7
			// update new 0
			c.posNumber[seven], c.posNumber[Position{seven.x - 1, seven.y - 1}] = 0, 0
			// update new 7 position
			c.numberPos[7] = Position{seven.x, seven.y - 1}
			// update uniformCost
			c.uniformCost++
			// update zero pos
			c.transition = Transition{number: 7, direction: UP}
			c.zeroPos = []Position{Position{seven.x, seven.y}, Position{seven.x - 1, seven.y - 1}}
			ajacentPos = append(ajacentPos, c)
		}
	}
	// go down
	if seven.y+1 < Width {
		d1 := Position{x: seven.x - 1, y: seven.y}
		d2 := Position{x: seven.x, y: seven.y + 1}
		if state.posNumber[d1] == 0 && state.posNumber[d2] == 0 {
			// update new 7
			c.posNumber[d1], c.posNumber[d2] = 7, 7
			// update new 0
			c.posNumber[Position{seven.x - 1, seven.y - 1}], c.posNumber[Position{seven.x, seven.y - 1}] = 0, 0
			// update new 7 position
			c.numberPos[7] = Position{seven.x, seven.y + 1}
			// update uniformCost
			c.uniformCost++
			// update zero pos
			c.transition = Transition{number: 7, direction: DOWN}
			c.zeroPos = []Position{Position{seven.x - 1, seven.y - 1}, Position{seven.x, seven.y - 1}}
			ajacentPos = append(ajacentPos, c)
		}
	}
	return ajacentPos
}

func PrintPath(p *State, result io.Writer) {
	transitionSeq := make([]Transition, 1)
	for p != nil {
		transitionSeq = append(transitionSeq, p.transition)
		fmt.Fprintln(result, p.Serilize())
		// fmt.Println(p.transition)
		p = p.parentState
	}
	for i := len(transitionSeq) - 1; i > 0; i-- {
		fmt.Println(transitionSeq[i])
	}
}

type StrState struct {
	state string
}

func P1(w http.ResponseWriter, r *http.Request) {
	goal := []int{1, 2, 3, 4, 5, 7, 7, 8, 9, 10, 6, 7, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 0, 0}
	// f, _ := os.Open("p1.txt")
	defer r.Body.Close()
	// buf := make([]byte, 200)
	// buf, err := ioutil.ReadAll(r.Body)
	str := r.URL.Query().Get("initial")
	// n, err := r.Body.Read(buf)
	r.ParseForm()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// str := string(buf)
	// str = strings.Replace(str, "\n", ",", -1)
	var numStr []string = strings.Split(str, ",")
	var num []int = make([]int, len(numStr))
	for i, s := range numStr {
		n, _ := strconv.ParseInt(s, 10, 32)
		num[i] = int(n)
	}
	initialState := buildState(num)
	goalState = buildState(goal)
	initialState.heuristicCost = HeuristicCost(initialState, goalState)
	initialState.uniformCost = 0
	closeStates = States{}
	openStates = States{initialState}
	heap.Init(&openStates)
	if r.Method == "GET" {
		log.Println("SOLOVING P1...")
		for {
			cur := heap.Pop(&openStates).(*State)
			// fmt.Println(cur.Serilize())
			closeStates.Push(cur)
			if IsEqual(cur, goalState) {
				fmt.Println("I have reached to the Goal!!!")
				PrintPath(cur, w)
				break
			}
			for _, n := range cur.findNextStates() {
				if !closeStates.Exist(n) {
					// if not reached before
					// fmt.Println("Push new state!")
					heap.Push(&openStates, n)
				} else {
					// fmt.Println("Already Met!")
				}
			}
		}

	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func Js(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.js")
}

func main() {
	p2.Exe()

	goal := []int{1, 2, 3, 4, 5, 7, 7, 8, 9, 10, 6, 7, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 0, 0}
	f, _ := os.Open("p1.txt")
	resultF, _ := os.OpenFile("result.txt", os.O_RDWR, 0755)
	defer f.Close()
	defer resultF.Close()
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
	initialState := buildState(num)
	goalState = buildState(goal)
	initialState.heuristicCost = HeuristicCost(initialState, goalState)
	initialState.uniformCost = 0
	closeStates = States{}
	openStates = States{initialState}

	p2.P2()
	heap.Init(&openStates)
	http.HandleFunc("/ai/exp1", P1)
	http.HandleFunc("/", Index)
	http.HandleFunc("/index.js", Js)
	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
