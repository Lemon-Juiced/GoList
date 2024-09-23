package main

import (
	"fmt"
	"reflect"
)

func itemToString[T any](item T) string {
	return fmt.Sprintf("%v", item)
}

/**
 * A generic list implementation in Go.
 * This implementation uses the `any` type constraint to create a generic list.
 * The `any` type constraint allows the list to store any type of data.
 *
 * @author: Lemon_Juiced
 */
type List[T any] struct {
	items []T
}

/**
 * Creates a new list.
 *
 * @return: A new list.
 */
func newList[T any]() *List[T] {
	return &List[T]{items: []T{}}
}

/**
 * Adds an item to the list.
 *
 * @param item: The item to add.
 */
func (l *List[T]) add(item T) {
	l.items = append(l.items, item)
}

/**
 * Adds multiple items to the list.
 *
 * @param items: The items to add.
 */
func (l *List[T]) addAll(items []T) {
	l.items = append(l.items, items...)
}

/**
 * Gets the size of the list.
 *
 * @return: The size of the list.
 */
func (l *List[T]) size() int {
	return len(l.items)
}

/**
 * Gets an item from the list at the given index.
 *
 * @param index: The index of the item to get.
 * @return: The item at the given index, or throws an error if the index is out of bounds.
 */
func (l *List[T]) get(index int) T {
	if index < 0 || index >= len(l.items) {
		panic(formatError("Index out of bounds."))
	}
	return l.items[index]
}

/**
 * Removes an item from the list at the given index.
 *
 * @param index: The index of the item to remove.
 * @return: The item that was removed, or throws an error if the index is out of bounds.
 */
func (l *List[T]) remove(index int) T {
	if index < 0 || index >= len(l.items) {
		panic(formatError("Index out of bounds."))
	}
	item := l.items[index]
	l.items = append(l.items[:index], l.items[index+1:]...)
	return item
}

/**
 * Clears the list. (Removes all items from the list)
 */
func (l *List[T]) clear() {
	l.items = []T{}
}

/**
 * Checks if the list contains the given item.
 *
 * @param item: The item to check for.
 * @return: True if the list contains the item, false otherwise.
 */
func (l *List[T]) contains(item T) bool {
	for _, i := range l.items {
		if reflect.DeepEqual(i, item) {
			return true
		}
	}
	return false
}

/**
 * Checks if the list is empty.
 *
 * @return: True if the list is empty, false otherwise.
 */
func (l *List[T]) isEmpty() bool {
	return len(l.items) == 0
}

/**
 * Gets the head of the list as if it were treated as a stack/queue.
 *
 * @return: The head of the list, or throws an error if the list is empty.
 */
func (l *List[T]) head() T {
	if len(l.items) == 0 {
		panic(formatError("List is empty."))
	}
	return l.items[0]
}

/**
 * Gets the tail of the list as if it were treated as a stack/queue.
 *
 * @return: The tail of the list, or throws an error if the list is empty.
 */
func (l *List[T]) tail() T {
	if len(l.items) == 0 {
		panic(formatError("List is empty."))
	}
	return l.items[len(l.items)-1]
}

/**
 * Removes the head of the list (pops) as if it were treated as a stack/queue.
 */
func (l *List[T]) pop() {
	if len(l.items) == 0 {
		panic(formatError("List is empty."))
	}
	l.items = l.items[1:]
}

/**
 * Removes the tail of the list (dequeues) as if it were treated as a stack/queue.
 */
func (l *List[T]) dequeue() {
	if len(l.items) == 0 {
		panic(formatError("List is empty."))
	}
	l.items = l.items[:len(l.items)-1]
}

/**
 * Converts the list to a string.
 *
 * @return: The string representation of the list.
 */
func (l *List[T]) asString() string {
	str := "["
	for i, item := range l.items {
		str += itemToString(item)
		if i < len(l.items)-1 {
			str += ", "
		}
	}
	str += "]"
	return str
}

/**
 * Converts the list to a string with a custom separator.
 *
 * @param separator: The separator to use.
 * @return: The string representation of the list.
 */
func (l *List[T]) asStringWithSeparator(separator string) string {
	str := ""
	for i, item := range l.items {
		str += itemToString(item)
		if i < len(l.items)-1 {
			str += separator
		}
	}
	return str
}

/**
 * Sets the item at the given index to the given item.
 *
 * @param index: The index of the item to set.
 * @param item: The item to set.
 */
func (l *List[T]) set(index int, item T) {
	if index < 0 || index >= len(l.items) {
		panic(formatError("Index out of bounds."))
	}
	l.items[index] = item
}

/**
 * Swaps the items at the given indices.
 *
 * @param index1: The first index.
 * @param index2: The second index.
 */
func (l *List[T]) swap(index1 int, index2 int) {
	if index1 < 0 || index1 >= len(l.items) || index2 < 0 || index2 >= len(l.items) {
		panic(formatError("Index out of bounds."))
	}
	l.items[index1], l.items[index2] = l.items[index2], l.items[index1]
}

/**
 * Reverses the list.
 */
func (l *List[T]) reverse() {
	for i := 0; i < len(l.items)/2; i++ {
		l.items[i], l.items[len(l.items)-i-1] = l.items[len(l.items)-i-1], l.items[i]
	}
}

/**
 * Sorts the list using the given comparison function.
 *
 * @param compare: The comparison function.
 */
func (l *List[T]) sort(compare func(T, T) bool) {
	for i := 0; i < len(l.items); i++ {
		for j := i + 1; j < len(l.items); j++ {
			if compare(l.items[j], l.items[i]) {
				l.items[i], l.items[j] = l.items[j], l.items[i]
			}
		}
	}
}

/**
 * Appends the given list to the current list.
 *
 * @param list: The list to append.
 */
func (l *List[T]) append(list *List[T]) {
	l.items = append(l.items, list.items...)
}

/**
 * Gets the union of the current list and the given list.
 *
 * @param list: The list to union with.
 * @return: The union of the two lists.
 */
func (l *List[T]) union(list *List[T]) *List[T] {
	union := newList[T]()
	union.addAll(l.items)
	union.addAll(list.items)
	return union
}

/**
 * Gets the intersection of the current list and the given list.
 *
 * @param list: The list to intersect with.
 * @return: The intersection of the two lists.
 */
func (l *List[T]) intersection(list *List[T]) *List[T] {
	intersection := newList[T]()
	for _, item := range l.items {
		if list.contains(item) {
			intersection.add(item)
		}
	}
	return intersection
}

/**
 * Formats an error message using red text with the ANSI escape codes.
 *
 * @param message: The message to format.
 * @return: The formatted error message.
 */
func formatError(message string) string {
	return "\033[31mError: " + message + "\033[0m\n"
}
