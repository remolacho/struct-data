package arrayInteger

import (
	"math/rand"
)

type Array struct {
	object []int
}

func New() *Array {
	var initArray []int
	return &Array{object: initArray}
}

func InitValues(array []int) *Array {
	return &Array{object: array}
}

func (a *Array) Rand(size int) {
	a.object = rand.Perm(size)
}

func (a *Array) Clone() []int {
	newArray := make([]int, a.Size())
	copy(newArray, a.object)
	return newArray
}

func (a *Array) Add(value int) *Array {
	a.object = append(a.object, value)
	return a
}

func (a *Array) Delete(index int) []int {
	a.object = append(a.object[:index], a.object[index+1:]...)
	return a.object
}

func (a *Array) Shift() []int {
	return a.Delete(0)
}

func (a *Array) Pop() []int {
	return a.Delete(a.Size() - 1)
}

func (a *Array) DeleteInCopy(index int) []int {
	newArray := a.Clone()
	newArray = append(newArray[:index], newArray[index+1:]...)
	return newArray
}

func (a *Array) Values() []int {
	return a.object
}

func (a *Array) Size() int {
	return len(a.object)
}

func (a *Array) Min() int {
	min := a.object[0]

	for i := 0; i <= a.Size()-1; i++ {
		if a.object[i] < min {
			min = a.object[i]
		}
	}

	return min
}

func (a *Array) Max() int {
	var max int

	for i := 0; i <= a.Size()-1; i++ {
		if a.object[i] > max {
			max = a.object[i]
		}
	}

	return max
}

func (a *Array) Avg() float64 {
	var avg int

	for i := 0; i <= a.Size()-1; i++ {
		avg += a.object[i]
	}

	return float64(avg) / float64(a.Size())
}

func (a *Array) Sum() int {
	var acum int

	for i := 0; i <= a.Size()-1; i++ {
		acum += a.object[i]
	}

	return acum
}

func (a *Array) FindAll(value int) []int {
	var newArray []int

	for i := 0; i <= a.Size()-1; i++ {
		if a.object[i] == value {
			newArray = append(newArray, value)
		}
	}

	return newArray
}

func (a *Array) Find(value int) int {
	for i := 0; i <= a.Size()-1; i++ {
		if a.object[i] == value {
			return value
		}
	}

	return -1
}

func (a *Array) Detect(value int) bool {
	for i := 0; i <= a.Size()-1; i++ {
		if a.object[i] == value {
			return true
		}
	}

	return false
}

func (a *Array) SortAsc() []int {
	newArray := a.Clone()
	sw := false
	aux := 0
	rounds := 0
	size := a.Size()

	for !sw {
		sw = true

		for i := 0; i < (size-1)-rounds; i++ {
			if newArray[i] > newArray[i+1] {
				aux = newArray[i+1]
				newArray[i+1] = newArray[i]
				newArray[i] = aux
				sw = false
			}
		}

		rounds++
	}

	return newArray
}

func (a *Array) SortDesc() []int {
	newArray := a.Clone()
	sw := false
	aux := 0
	rounds := 0
	size := a.Size()

	for !sw {
		sw = true

		for i := 0; i < (size-1)-rounds; i++ {
			if newArray[i] < newArray[i+1] {
				aux = newArray[i+1]
				newArray[i+1] = newArray[i]
				newArray[i] = aux
				sw = false
			}
		}

		rounds++
	}

	return newArray
}

func (a *Array) SortSelection() []int {
	newArray := a.Clone()
	size := a.Size()
	posMin := 0
	aux := 0

	for i := 0; i < size; i++ {
		posMin = i
		for j := i + 1; j < size; j++ {
			if newArray[posMin] > newArray[j] {
				posMin = j
			}
		}

		aux = newArray[i]
		newArray[i] = newArray[posMin]
		newArray[posMin] = aux
	}

	return newArray
}

func (a *Array) SortSelectionMe() []int {
	newArray := a.Clone()
	posLeft := 0
	posRight := 0
	size := a.Size()
	min := 0
	changed := false

	for posLeft < size {
		min = newArray[posLeft]
		changed = false

		for i := posLeft; i < size; i++ {
			if min > newArray[i] {
				min = newArray[i]
				posRight = i
				changed = true
			}
		}

		if changed {
			min = newArray[posLeft]
			newArray[posLeft] = newArray[posRight]
			newArray[posRight] = min
		}

		posLeft++
	}

	return newArray
}

func (a *Array) SortByInsertion() []int {
	newArray := a.Clone()
	size := a.Size()
	var aux int

	for i := 1; i < size; i++ {
		aux = newArray[i]
		var j int

		for j = i - 1; j >= 0 && newArray[j] > aux; j-- {
			newArray[j+1] = newArray[j]
		}

		newArray[j+1] = aux
	}

	return newArray
}

func (a *Array) SortByDoubleRun() []int {
	newArray := a.Clone()
	size := a.Size()
	var aux int
	var posMin int
	var posMax int

	posTop := size - 1

	for i := 0; i < size/2; i++ {
		posMin = i
		posMax = i

		for j := i; j <= posTop; j++ {
			if newArray[j] > newArray[posMax] {
				posMax = j
			}
			if newArray[j] < newArray[posMin] {
				posMin = j
			}
		}

		aux = newArray[posMin]
		newArray[posMin] = newArray[i]
		newArray[i] = aux

		if posMax == i {
			posMax = posMin
		}

		aux = newArray[posTop]
		newArray[posTop] = newArray[posMax]
		newArray[posMax] = aux

		posTop--
	}

	return newArray
}
