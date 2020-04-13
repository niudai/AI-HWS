package main

import "strconv"

const NumBlock = 21
const Width = 5

type Position struct {
	x int
	y int
}

type State struct {
	numberPos map[int]Position
	posNumber map[Position]int
	zeroPos   []Position
	// sevenPos  int
	heuristicCost int
	uniformCost   int
	parentState   *State
}

type States []*State

func (s *States) Len() int {
	return len(*s)
}

// Less reports whether the element with
// index i should sort before the element with index j.
// Less(i, j int) bool
// Swap swaps the elements with indexes i and j.
// Swap(i, j int)

// Untested
func (s *States) Less(i, j int) bool {
	return (*s)[i].uniformCost+(*s)[i].heuristicCost < (*s)[j].uniformCost+(*s)[j].heuristicCost
}

// Untested
func (s *States) Swap(i, j int) {
	tmp := (*s)[i]
	(*s)[i] = (*s)[j]
	(*s)[j] = tmp
}

// Untested
// Push distinct state
func (ss *States) Push(i interface{}) {
	s := i.(*State)
	for _, _s := range *ss {
		if IsEqual(_s, s) {
			return
		}
	}
	*ss = append(*ss, s)
}

// Untested
func (ss *States) Pop() interface{} {
	r := (*ss)[len(*ss)-1]
	*ss = (*ss)[:len(*ss)-1]
	return r
}

// check two state if they are equal
func IsEqual(m, n *State) bool {
	for i := 1; i <= NumBlock; i++ {
		if m.numberPos[i] != n.numberPos[i] {
			return false
		}
	}
	return true
}

func (s *State) Clone() *State {
	c := State{numberPos: make(map[int]Position), posNumber: make(map[Position]int), zeroPos: make([]Position, 2), heuristicCost: 0, uniformCost: 0}
	for k, v := range s.numberPos {
		c.numberPos[k] = v
	}
	for k, v := range s.posNumber {
		c.posNumber[k] = v
	}
	copy(c.zeroPos, s.zeroPos)
	c.heuristicCost = s.heuristicCost
	c.uniformCost = s.uniformCost
	return &c
}

// Untested
func (ss States) Exist(s *State) bool {
	for _, _s := range ss {
		if _s == s {
			return true
		}
	}
	return false
}

func (s *State) Serilize() string {
	str := ""
	for i := 0; i < Width; i++ {
		// i is row number
		for j := 0; j < Width; j++ {
			// j is column number, say, the x
			str = str + strconv.Itoa(s.posNumber[Position{x: j, y: i}])
			if j == Width-1 {
				str = str + "\n"
			} else {
				str = str + ", "
			}
		}
	}
	return str
}

func HeuristicCost(a, b *State) int {

	if len(a.numberPos) != len(b.numberPos) {
		panic("two pos array not same len")
	}
	sum := 0
	for i := 0; i < NumBlock; i++ {
		sum += manhattenDist(a.numberPos[i], b.numberPos[i])
	}
	return sum
}

func manhattenDist(a, b Position) int {
	return abs(a.x, b.x) + abs(a.y, b.y)
}

func abs(a, b int) int {
	if a >= b {
		return a - b
	} else {
		return b - a
	}
}
