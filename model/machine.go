package model

import "math/rand"

type Machine struct {
	nums  [75]int
	index int
}

func NewMachine() Machine {
	slice := rand.Perm(75)
	var nums [75]int
	copy(nums[:], slice)
	return Machine{
		nums:  nums,
		index: -1,
	}
}

func (m *Machine) Roll() *int {
	m.index++
	if m.index >= 75 {
		return nil
	}

	num := m.nums[m.index]
	return &num
}
