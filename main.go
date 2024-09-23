package main

import (
	"fmt"
)

func blueText(text string) string {
	return fmt.Sprintf("\033[34m%s\033[0m", text)
}

func greenText(text string) string {
	return fmt.Sprintf("\033[32m%s\033[0m", text)
}

func main() {
	// Test newList()
	fmt.Println(blueText("Running Test newList()"))
	list := newList[int]()
	if list == nil {
		fmt.Println("Test newList(): Expected newList() to return a non-nil list")
	} else {
		fmt.Printf("%s: %v\n", greenText("Test newList() passed"), list.asString())
	}
	fmt.Println()

	// Test add() and size()
	fmt.Println(blueText("Running Test add(1) and size()"))
	list.add(1)
	if list.size() != 1 {
		fmt.Printf("Test add(1) and size(): Expected size() to be 1, got %d\n", list.size())
	} else {
		fmt.Printf("%s: %v\n", greenText("Test add(1) and size() passed"), list.asString())
	}
	fmt.Println()

	// Test addAll()
	fmt.Println(blueText("Running Test addAll([]int{2, 3})"))
	list.addAll([]int{2, 3})
	if list.size() != 3 {
		fmt.Printf("Test addAll([]int{2, 3}): Expected size() to be 3, got %d\n", list.size())
	} else {
		fmt.Printf("%s: %v\n", greenText("Test addAll([]int{2, 3}) passed"), list.asString())
	}
	fmt.Println()

	// Test get()
	fmt.Println(blueText("Running Test get(0)"))
	if list.get(0) != 1 {
		fmt.Printf("Test get(0): Expected get(0) to return 1, got %d\n", list.get(0))
	} else {
		fmt.Printf("%s: %v\n", greenText("Test get(0) passed"), list.get(0))
	}
	fmt.Println()

	// Test remove()
	fmt.Println(blueText("Running Test remove(1)"))
	removedItem := list.remove(1)
	if removedItem != 2 {
		fmt.Printf("Test remove(1): Expected removed item to be 2, got %d\n", removedItem)
	} else {
		fmt.Printf("%s: %v\n", greenText("Test remove(1) passed"), list.asString())
	}
	if list.size() != 2 {
		fmt.Printf("Test remove(1): Expected size() to be 2, got %d\n", list.size())
	} else {
		fmt.Printf("%s: %v\n", greenText("Test remove(1) size check passed"), list.asString())
	}
	fmt.Println()

	// Test clear()
	fmt.Println(blueText("Running Test clear()"))
	list.clear()
	if list.size() != 0 {
		fmt.Printf("Test clear(): Expected size() to be 0, got %d\n", list.size())
	} else {
		fmt.Printf("%s: %v\n", greenText("Test clear() passed"), list.asString())
	}
	fmt.Println()

	// Test contains()
	fmt.Println(blueText("Running Test contains(1) and contains(2)"))
	list.add(1)
	if !list.contains(1) {
		fmt.Println("Test contains(1): Expected list to contain 1")
	} else {
		fmt.Printf("%s: %v\n", greenText("Test contains(1) passed"), list.asString())
	}
	if list.contains(2) {
		fmt.Println("Test contains(2): Expected list to not contain 2")
	} else {
		fmt.Printf("%s: %v\n", greenText("Test contains(2) passed"), list.asString())
	}
	fmt.Println()

	// Test isEmpty()
	fmt.Println(blueText("Running Test isEmpty()"))
	if list.isEmpty() {
		fmt.Println("Test isEmpty(): Expected list to not be empty")
	} else {
		fmt.Printf("%s: %v\n", greenText("Test isEmpty() passed for non-empty list"), list.asString())
	}
	list.clear()
	if !list.isEmpty() {
		fmt.Println("Test isEmpty(): Expected list to be empty")
	} else {
		fmt.Printf("%s: %v\n", greenText("Test isEmpty() passed for empty list"), list.asString())
	}
	fmt.Println()

	// Test head() and tail()
	fmt.Println(blueText("Running Test head() and tail()"))
	list.addAll([]int{1, 2, 3})
	if list.head() != 1 {
		fmt.Printf("Test head(): Expected head() to be 1, got %d\n", list.head())
	} else {
		fmt.Printf("%s: %v\n", greenText("Test head() passed"), list.head())
	}
	if list.tail() != 3 {
		fmt.Printf("Test tail(): Expected tail() to be 3, got %d\n", list.tail())
	} else {
		fmt.Printf("%s: %v\n", greenText("Test tail() passed"), list.tail())
	}
	fmt.Println()

	// Test pop() and dequeue()
	fmt.Println(blueText("Running Test pop() and dequeue()"))
	list.pop()
	if list.head() != 2 {
		fmt.Printf("Test pop(): Expected head() to be 2 after pop(), got %d\n", list.head())
	} else {
		fmt.Printf("%s: %v\n", greenText("Test pop() passed"), list.asString())
	}
	list.dequeue()
	if list.tail() != 2 {
		fmt.Printf("Test dequeue(): Expected tail() to be 2 after dequeue(), got %d\n", list.tail())
	} else {
		fmt.Printf("%s: %v\n", greenText("Test dequeue() passed"), list.asString())
	}
	fmt.Println()

	// Test asString() and asStringWithSeparator()
	fmt.Println(blueText("Running Test asString() and asStringWithSeparator()"))
	list.clear()
	list.addAll([]int{1, 2, 3})
	if list.asString() != "[1, 2, 3]" {
		fmt.Printf("Test asString(): Expected asString() to return '[1, 2, 3]', got %s\n", list.asString())
	} else {
		fmt.Printf("%s: %v\n", greenText("Test asString() passed"), list.asString())
	}
	if list.asStringWithSeparator("; ") != "1; 2; 3" {
		fmt.Printf("Test asStringWithSeparator(): Expected asStringWithSeparator() to return '1; 2; 3', got %s\n", list.asStringWithSeparator("; "))
	} else {
		fmt.Printf("%s: %v\n", greenText("Test asStringWithSeparator() passed"), list.asStringWithSeparator("; "))
	}
	fmt.Println()

	// Test set()
	fmt.Println(blueText("Running Test set(1, 5)"))
	list.set(1, 5)
	if list.get(1) != 5 {
		fmt.Printf("Test set(1, 5): Expected set(1, 5) to set value at index 1 to 5, got %d\n", list.get(1))
	} else {
		fmt.Printf("%s: %v\n", greenText("Test set(1, 5) passed"), list.asString())
	}
	fmt.Println()

	// Test swap()
	fmt.Println(blueText("Running Test swap(0, 2)"))
	list.swap(0, 2)
	if list.get(0) != 3 || list.get(2) != 1 {
		fmt.Printf("Test swap(0, 2): Expected swap(0, 2) to swap values at indices 0 and 2, got %d and %d\n", list.get(0), list.get(2))
	} else {
		fmt.Printf("%s: %v\n", greenText("Test swap(0, 2) passed"), list.asString())
	}
	fmt.Println()

	// Test reverse()
	fmt.Println(blueText("Running Test reverse()"))
	list.reverse()
	if list.get(0) != 1 || list.get(2) != 3 {
		fmt.Printf("Test reverse(): Expected reverse() to reverse the list, got %d and %d\n", list.get(0), list.get(2))
	} else {
		fmt.Printf("%s: %v\n", greenText("Test reverse() passed"), list.asString())
	}
	fmt.Println()

	// Test sort()
	fmt.Println(blueText("Running Test sort()"))
	list.sort(func(a, b int) bool { return a < b })
	if list.get(0) != 1 || list.get(2) != 5 {
		fmt.Printf("Test sort(): Expected sort() to sort the list, got %d and %d\n", list.get(0), list.get(2))
	} else {
		fmt.Printf("%s: %v\n", greenText("Test sort() passed"), list.asString())
	}
	fmt.Println()

	// Test append()
	fmt.Println(blueText("Running Test append(anotherList)"))
	anotherList := newList[int]()
	anotherList.addAll([]int{6, 7})
	list.append(anotherList)
	if list.size() != 5 || list.get(4) != 7 {
		fmt.Printf("Test append(anotherList): Expected append() to append another list, got size %d and last element %d\n", list.size(), list.get(4))
	} else {
		fmt.Printf("%s: %v\n", greenText("Test append(anotherList) passed"), list.asString())
	}
	fmt.Println()

	// Test union()
	fmt.Println(blueText("Running Test union()"))
	list1 := newList[int]()
	list1.addAll([]int{1, 2, 3})
	list2 := newList[int]()
	list2.addAll([]int{3, 4, 5})
	unionList := list1.union(list2)
	expectedUnion := []int{1, 2, 3, 3, 4, 5}
	for i, item := range expectedUnion {
		if unionList.get(i) != item {
			fmt.Printf("Test union(): Expected union() list to have item %d at index %d, got %d\n", item, i, unionList.get(i))
		} else {
			fmt.Printf("%s for index %d: %v\n", greenText("Test union() passed"), i, unionList.asString())
		}
	}
	fmt.Println()

	// Test intersection()
	fmt.Println(blueText("Running Test intersection()"))
	intersectionList := list1.intersection(list2)
	expectedIntersection := []int{3}
	for i, item := range expectedIntersection {
		if intersectionList.get(i) != item {
			fmt.Printf("Test intersection(): Expected intersection() list to have item %d at index %d, got %d\n", item, i, intersectionList.get(i))
		} else {
			fmt.Printf("%s for index %d: %v\n", greenText("Test intersection() passed"), i, intersectionList.asString())
		}
	}
	fmt.Println()
}
