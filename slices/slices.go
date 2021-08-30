package slices

import "fmt"

func ArrayBasics() {
	// if the count is specified it is an array
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Printf("\n%v, %T", arr, arr)
	// type will say [3]int

	// this will copy the array and not be a reference.
	b := arr
	b[1] = 5 // will update the copied array element and not the original
	fmt.Printf("\n%v, %T", b, b)
	fmt.Printf("\n%v, %T", arr, arr) // this will still show 2 in the 2nd element

	// we can define an array without specifying the size. This is an array not a slice.
	var arr1 [3]int = [...]int{5, 6, 7}
	fmt.Printf("\n%v, %T", arr1, arr1)

	// this is again an array not a slice. Check the type output
	arr2 := [...]int{5, 19, 10}
	fmt.Printf("\n%v, %T", arr2, arr2)
}

func SliceBasics() {
	// it is very simlar to array but points to some part of the array
	var arr [5]int = [...]int{1, 2, 3, 4, 5}
	var slc []int = arr[1:3]                   // create a slice of the original array from position 1 till 3 - (not including 3)
	fmt.Printf("\n%v, %T", arr, arr)           // type will return it as [5]int
	fmt.Printf("\n%v, %T", slc, slc)           // the type will return it as []int
	fmt.Printf("\n%v, %v", len(slc), cap(slc)) // the len will return length of the slice and cap will return the underlying capacity - based on underlying array
}

func SliceBasics2() {
	// it is very simlar to array but points to some part of the array
	var arr [5]int = [...]int{1, 2, 3, 4, 5}
	var slc []int = arr[1:5] // create a slice of the original array from position 1 till 3 - (not including 3)
	fmt.Printf("\nBefore updating the slice, array value is : %v, %T", arr, arr)
	fmt.Printf("\nBefore updating the slice, slice value is : %v, %T", slc, slc)
	// updating the slice will update the array value
	slc[0] = 20
	fmt.Printf("\n After updating the slice, array value is  : %v, %T", arr, arr)
	fmt.Printf("\n After updating the slice, slice value is  : %v, %T", slc, slc)

	// = or := operation will not copy the data and create a new slice (like it did in array)
	// it will be a reference to the same underlying array
	slc1 := slc[0:2]
	fmt.Printf("\n Before updating the slice2, array   value is  : %v, %T", arr, arr)
	fmt.Printf("\n Before updating the slice2, slice 1 value is  : %v, %T", slc, slc)
	fmt.Printf("\n Before updating the slice2, slice 2 value is  : %v, %T", slc1, slc1)

	slc1[0] = 200 // this will update the original array value ; that means slc will also be updated
	fmt.Printf("\n After updating the slice2, array   value is  : %v, %T", arr, arr)
	fmt.Printf("\n After updating the slice2, slice 1 value is  : %v, %T", slc, slc)
	fmt.Printf("\n After updating the slice2, slice 2 value is  : %v, %T", slc1, slc1)
}

func SliceBasics3() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := a[:]   // consists of all elements of slice a
	c := a[2:]  // consists of all elements from position 2 (starting position is 0)
	d := a[:6]  // consists of all elements from position 0 till 6 (not including position 6)
	e := a[2:6] // consists of all elements from position 2, 3, 4, 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}

func SliceBasics4() {
	// create a slice of size 3
	a := make([]int, 3)
	fmt.Printf("\n%v, %T", a, a) // initializes to 0
	fmt.Printf("\n%v, %v", len(a), cap(a))
	a = append(a, 3)
	fmt.Println("\nAfter appending one value to the slice")
	fmt.Printf("\n%v, %T", a, a)
	fmt.Printf("\n%v, %v", len(a), cap(a)) // capacity will become larger than original length now
	// if we try to assign a value more than length, it will fail.
	// it will give an error as - index out of range [6] with length 4
	// uncomment and test
	//a[6] = 10
	//fmt.Printf("\n%v, %T", a, a)

	// append can take multiple values
	a = append(a, 3, 4, 5)
	fmt.Printf("\n%v, %T", a, a)

	// can append other slices as well using the ... operator
	b := []int{10, 11, 12}
	a = append(a, b...)
	fmt.Printf("\n%v, %T", a, a)
	fmt.Printf("\n%v, %v", len(a), cap(a))
}

func SliceBasics5() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	// if we do something like this - that is create a b variable that concatenation first 2 and last 2
	b := append(a[:2], a[3:]...)
	fmt.Println(b)
	fmt.Println(a) // this a will also be updated. So, be careful if we do such things

	// instead try this with slices
	fmt.Println("Try copying with make and copy")
	a2 := []int{1, 2, 3, 4, 5}
	fmt.Println(a2)
	b2 := make([]int, 2)
	copy(b2, a2[:2])
	//b2 = append(b2, a[:2]...)
	b2 = append(b2, a2[3:]...)
	fmt.Println(b2)
	fmt.Println(a2)
}
