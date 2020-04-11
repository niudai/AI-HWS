package main

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

// TODO: not tested
func (s *State) Serilize() string {
	str := make([]byte, 0)
	for i := 0; i < Width; i++ {
		for j := 0; j < Width; j++ {
			s.posNumber[Position{x: i, y: j}]
		}
	}
	return string(str)
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
