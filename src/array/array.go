package array

import "fmt"

type array[T comparable] struct {
	array *[]T
}

func NewArray[T comparable](arr *[]T) *array[T] {
	return &array[T]{
		array: arr,
	}
}

func (a *array[T]) ForEach(callback func(item *T, index int)) {
	for i, item := range *a.array {
		callback(&item, i)
	}
}

func (a *array[T]) Map(callback func(item T, index int) any) []any {
	result := make([]any, len(*a.array))
	for i, item := range *a.array {
		result[i] = callback(item, i)
	}
	return result
}

func (a *array[T]) Filter(callback func(item T, index int) bool) []T {
	result := make([]T, 0)
	for i, item := range *a.array {
		if callback(item, i) {
			result = append(result, item)
		}
	}
	return result
}

func (a *array[T]) Reduce(callback func(accumulator any, item T, index int) any, initialValue any) any {
	accumulator := initialValue
	for i, item := range *a.array {
		accumulator = callback(accumulator, item, i)
	}
	return accumulator
}

func (a *array[T]) Find(callback func(item T, index int) bool) (T, bool) {
	for i, item := range *a.array {
		if callback(item, i) {
			return item, true
		}
	}
	var zeroValue T
	return zeroValue, false
}

func (a *array[T]) FindIndex(callback func(item T, index int) bool) int {
	for i, item := range *a.array {
		if callback(item, i) {
			return i
		}
	}
	return -1
}

func (a *array[T]) Some(callback func(item T, index int) bool) bool {
	for i, item := range *a.array {
		if callback(item, i) {
			return true
		}
	}
	return false
}

func (a *array[T]) Every(callback func(item T, index int) bool) bool {
	for i, item := range *a.array {
		if !callback(item, i) {
			return false
		}
	}
	return true
}

func (a *array[T]) Length() int {
	return len(*a.array)
}

func (a *array[T]) Clear() {
	*a.array = []T{}
}

func (a *array[T]) ToArray() []T {
	return *a.array
}

func (a *array[T]) Copy() *array[T] {
	newArray := make([]T, len(*a.array))
	copy(newArray, *a.array)
	return NewArray(&newArray)
}

func (a *array[T]) Get(index int) (T, bool) {
	if index < 0 || index >= len(*a.array) {
		var zeroValue T
		return zeroValue, false
	}
	return (*a.array)[index], true
}

func (a *array[T]) Set(index int, value T) bool {
	if index < 0 || index >= len(*a.array) {
		return false
	}
	(*a.array)[index] = value
	return true
}

func (a *array[T]) Push(value T) {
	*a.array = append(*a.array, value)
}

func (a *array[T]) Pop() (T, bool) {
	if len(*a.array) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	value := (*a.array)[len(*a.array)-1]
	*a.array = (*a.array)[:len(*a.array)-1]
	return value, true
}

func (a *array[T]) Shift() (T, bool) {
	if len(*a.array) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	value := (*a.array)[0]
	*a.array = (*a.array)[1:]
	return value, true
}

func (a *array[T]) Unshift(value T) {
	*a.array = append([]T{value}, *a.array...)
}

func (a *array[T]) IndexOf(value T) int {
	for i, item := range *a.array {
		if item == value {
			return i
		}
	}
	return -1
}

func (a *array[T]) LastIndexOf(value T) int {
	for i := len(*a.array) - 1; i >= 0; i-- {
		if (*a.array)[i] == value {
			return i
		}
	}
	return -1
}

func (a *array[T]) Slice(start, end int) []T {
	if start < 0 {
		start = 0
	}
	if end > len(*a.array) || end < 0 {
		end = len(*a.array)
	}
	if start >= end {
		return []T{}
	}
	return (*a.array)[start:end]
}

func (a *array[T]) Reverse() {
	n := len(*a.array)
	for i := 0; i < n/2; i++ {
		(*a.array)[i], (*a.array)[n-i-1] = (*a.array)[n-i-1], (*a.array)[i]
	}
}

func (a *array[T]) Sort(comparator func(a, b T) int) {
	if len(*a.array) <= 1 {
		return
	}
	for i := 0; i < len(*a.array)-1; i++ {
		for j := 0; j < len(*a.array)-i-1; j++ {
			if comparator((*a.array)[j], (*a.array)[j+1]) > 0 {
				(*a.array)[j], (*a.array)[j+1] = (*a.array)[j+1], (*a.array)[j]
			}
		}
	}
}

func (a *array[T]) Join(separator string) string {
	result := ""
	for i, item := range *a.array {
		if i > 0 {
			result += separator
		}
		result += anyToString(item)
	}
	return result
}

func anyToString(value any) string {
	switch v := value.(type) {
	case string:
		return v
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", v)
	case float32, float64:
		return fmt.Sprintf("%f", v)
	case bool:
		return fmt.Sprintf("%t", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}
