package basics

import (
	"fmt"
	"strconv"
)

// Function which are defined with starting letter as Caps can be called from outside the package.

func VariablesInitialize() {
	//3 different way of defining
	//first method
	var ival int
	ival = 21
	//2nd method
	var ival2 int = 42
	//3rd method := can be used create and initialize together.
	ival3 := 84

	var ival4 int
	fmt.Println(ival)
	fmt.Println(ival2)
	fmt.Println(ival3)
	fmt.Println(ival4) //always defaults to zero - no need for initialization
}

func VariablesInt() {
	//int is same as int32
	var var1 int
	var var2 int8
	var var3 int16
	var var4 int32
	var var5 int64

	// can set multiple variables in one shot
	var1, var2, var3, var4, var5 = 1, 2, 3, 4, 5

	//%v - can be used to print any data type - no need to worry about the data type
	//%T - tells the type of the variable
	fmt.Printf("%v, %T\n", var1, var1)
	fmt.Printf("%v, %T\n", var2, var2)
	fmt.Printf("%v, %T\n", var3, var3)
	fmt.Printf("%v, %T\n", var4, var4)
	fmt.Printf("%v, %T\n", var5, var5)

	// one variable of a type cannot be assigned to another. The below will give an compilation error
	//Errors will be like this - cannot use var4 (type int32) as type int in assignment;  cannot use var1 (type int) as type int16 in assignment
	//var1 = var4
	//var3 = var1
}

func VariablesBool() {
	var isValid bool

	// default value is zero again. Therefore, it should print alse.
	fmt.Printf("\n%v, %T", isValid, isValid)

	// comparison operations also return a bool. See examples below
	isEqual := (5 * 2) == 10
	fmt.Printf("\n%v, %T", isEqual, isEqual)

	isNotEqual := (5 * 20) == 99
	fmt.Printf("\n%v, %T", isNotEqual, isNotEqual)
}

func VariablesConv() {
	var var1 int = 90
	var var2 float32

	// this assignment of a integer variable to float will not work and give a compilation error.
	// Error would be somethin like - cannot use var1 (type int) as type float32 in assignment
	// uncomment and test
	//var2 = var1
	//therefore, it needs to be done via calling the appropriate conversion functions - in this case float32()
	var2 = float32(var1)
	fmt.Printf("\n%v, %T", var2, var2)

	//converting a string to int.
	// Atoi function returns 2 details - the converted value and error.
	// if error is nil, then conversion was sucessful
	var var3 string = "9"
	var4, e := strconv.Atoi(var3)
	if e == nil {
		fmt.Printf("\nSuccessful conversion - %v, %T", var4, var4)
	}
	var5, e := strconv.Atoi("a30")
	if e == nil {
		fmt.Printf("\nSuccessful conversion - %v, %T", var5, var5)
	} else {
		//print the error as well
		fmt.Printf("\nUnsuccessful conversion - %v, %T, %v", var5, var5, e)
	}

	// conversion of integer to string is via Itoa function.
	var6 := 40
	var varstr2 string
	varstr2 = strconv.Itoa(var6)
	fmt.Printf("\nSuccessful conversion - %v, %T", varstr2, varstr2)
}

func OperationsBasic() {

	fmt.Println("Addition :", 10+5)
	fmt.Println("Subtraction :", 10-5)
	fmt.Println("Multiplication :", 10*5)
	// division returns an integer again. not a float. so, it will lose data if there is a reminder.
	fmt.Println("Division :", 11/5)
	//modulus operator returns the remainder.
	fmt.Println("Modulus :", 11%5)
}

func OperationsBitwise() {
	var var1 int = 10
	var var2 int = 2

	fmt.Printf("\n base 2 values of %v and %v are - %b, %b", var1, var2, var1, var2)

	fmt.Printf("\n Bitwise Or on  %v and %v is %b. Value is %v", var1, var2, var1|var2, var1|var2)
	fmt.Printf("\n Bitwise And on %v and %v is %b. Value is %v", var1, var2, var1&var2, var1&var2)

	fmt.Printf("\n Bitwise Xor on %v and %v is %b. Value is %v", var1, var2, var1^var2, var1^var2)

	fmt.Printf("\n Bitwise right shift operation on %v by 1 makes it %v. Binary representation is %b", var2, var2>>1, var2>>1)

	fmt.Printf("\n Bitwise left shift operation on %v by 1 makes it %v. Binary representation is %b", var2, var2<<1, var2<<1)

}

func VariablesFloat() {
	// if := is used to initialize a float, it is always considered as float64
	f1 := 3.14
	f2 := 3e50
	fmt.Printf("\n %v, %T", f1, f1)
	fmt.Printf("\n %v, %T", f2, f2)
	// cannot assign float32 to float64
	var f3 float32 = 3.11
	// this assignment below will give a compilation error
	// error is - cannot use f2 (type float64) as type float32 in assignment
	// uncomment and test
	//f3 = f2
	fmt.Printf("\n %v, %T", f3, f3)

	// use conversion function to convert from float64 to float32
	var f4 float32
	var f5 float32
	// f1 can be converted.
	// but f2 converted to float32 will give Inf - as it has larger value than float32
	f4 = float32(f1)
	f5 = float32(f2)
	fmt.Printf("\n %v, %T", f4, f4)
	fmt.Printf("\n %v, %T", f5, f5)
}

func VariablesComplex() {
	// complex variables are primitive data types in golang.
	// complex64 and complex128
	var c complex64 = 3 + 5i
	fmt.Printf("\n%v, %T", c, c)
	fmt.Printf("\n%v, %T", real(c), real(c))
	fmt.Printf("\n%v, %T", imag(c), imag(c))

	// we can also initialize just the complex part
	// real part will be considered as 0
	var c1 complex64 = 5i
	fmt.Printf("\n%v, %T", c1, c1)
	// or initialized with only the real part.
	var c2 complex64 = 3
	fmt.Printf("\n%v, %T", c2, c2)

	// we can do add, subtract, divide and multiply of complex variables.
	fmt.Printf("\n%v", c1+c2)
	fmt.Printf("\n%v", c1-c2)
	fmt.Printf("\n%v", c1*c2)
	fmt.Printf("\n%v", c1/c2)
}

func VariablesStrings() {

	var s string = "hello there"
	fmt.Printf("\n%v, %T", s, s)

	// if we try to get a particular character in the string,
	//it will be of type byte and return the ascii value
	// This can be a bit confusing for folks coming from other languages.
	fmt.Printf("\n%v, %T", s[0], s[0])

	// strings are immutable; therefore, we cannot assign any value to a particular position or change it.
	// for example, below case will fail with error - cannot assign to s[1] (strings are immutable)
	// uncomment and test
	//s[1] = 'o'

	var s1 string = "Gaël"
	fmt.Printf("\n%v, %T", s1, s1)

	// if you take this chinese string and use it as is, it can print the value
	var s2 string = "世界"
	fmt.Printf("\n%v, %T", s2, s2)
	// check the size and it will not return 2. it will return 6. That is 3 bytes per character. 3*2
	fmt.Printf("\n %v", len(s2))
	// if we try to print the s2[1], it will return the ascii code of the first byte of the 3 byte chinese character
	fmt.Printf("\n%v, %T", s2[1], s2[1])

	// compare the length of strings of english, french and chinese of 2 characters
	// it will return 2 , 3 and 6. the ë occupies 2 bytes; and each chinese character occupies 3 characters
	fmt.Printf("\n %v, %v, %v", len("el"), len("ël"), len("世界"))
	//therefore, be careful about how to use the character when looping through characters in a string.

	// looping through character and using it as a character directly will be fine
	// but if we try to convert it to a string based on position in the array it can return odd values.
	// see examples below.
	// the following 3 loops show data properly.
	fmt.Println("")
	for _, c := range "el" {
		fmt.Print(string(c))
	}
	fmt.Println("")
	for _, c := range "世界" {
		fmt.Print(string(c))
	}
	fmt.Println("")
	for _, c := range "ël" {
		fmt.Print(string(c))
	}

	fmt.Println("")
	// the following 3 loops will not show data properly. English string will be fine.
	var str1 string = "el"
	for i := range str1 {
		fmt.Print(string(str1[i]))
	}

	fmt.Println("")
	// the first character will not print properly.
	var str2 string = "ël"
	for i := range str2 {
		fmt.Print(string(str2[i]))
	}

	fmt.Println("")
	// both characters will not be printed properly
	var str3 string = "世界"
	for i := range str3 {
		fmt.Print(str3[i])
	}
	// some references - https://coderwall.com/p/k7zvyg/dealing-with-unicode-in-go
}

func PointersBasics() {
	//definition of a pointer
	var firstname *string = new(string)
	//assign a value to the pointer
	*firstname = "Akash!"
	//get the address
	fmt.Printf("Address of the pointer: - %v\n", firstname)
	//get the value
	fmt.Printf("Value pointed by the pointer: - %v\n", *firstname)
	// actions like +1 on pointers are not allowed on pointers(what we can do in c/c++)
	// and it will give an compilation error as given below.
	// "invalid operation: firstname + 1 (mismatched types *string and int).""
	// To test this, uncomment the below line to verify
	// fmt.Printf("Pointer addresss + 1 : %v", firstname+1)
}

func PointersBasics2() {
	var FirstName = "Akash"
	var ptr *string = &FirstName
	//ptr references the address; *ptr shows the value at that address.
	fmt.Printf("The pointer and value are %v, %v\n", ptr, *ptr)
	FirstName = "Lokesh"
	// when the value at address is changed in this case FirstName, the pointer value remains the same
	fmt.Printf("The pointer and value are %v, %v\n", ptr, *ptr)
}

func ConstantsBasics() {
	// can define without a type - called untyped consts
	const a = 42
	fmt.Printf("\n %v, %T", a, a)
	// can define with a type - so, it follow all rules of a type - cannot be used with other types
	const a1 int = 10
	var b1 int32 = 90
	// this below addition operation will not work as the a1 is of type int and b1 is of type int32
	// uncomment check the compilation error - "invalid operation: a1 + b1 (mismatched types int and int32)"
	//fmt.Printf("\n%v %T", a1+b1, a1+b1)
	// in this case the addition will work as a2 is untyped
	const a2 = 10
	fmt.Printf("\n%v %T", a2+b1, a2+b1)

	// consts are immutable - so, it cannot be changed later on.
	// the below line will give an error saying cannot assign to a (untyped int constant 42)
	//a = 99

	const (
		x = iota
		y = iota
		z = iota
	)
	// iota initializes it to 0 and every other occurence of iota is +1.
	fmt.Printf("\n%v, %T", x, x)
	fmt.Printf("\n%v, %T", y, y)
	fmt.Printf("\n%v, %T", z, z)

	// we do not have to explicitly set the iota for consequent ones as well.
	const (
		x1 = iota
		y1
		z1
	)
	// x1 will be 0 and other will be +1 the previous value.
	fmt.Printf("\n%v, %T", x1, x1)
	fmt.Printf("\n%v, %T", y1, y1)
	fmt.Printf("\n%v, %T", z1, z1)

	// we can also set it some fixed value by doing some numeric calculation or skip a particular value
	const (
		x2 = iota + 5
		y2
		z2
	)
	// x1 will be 5 and other will be +1 the previous value.
	fmt.Printf("\n%v, %T", x2, x2)
	fmt.Printf("\n%v, %T", y2, y2)
	fmt.Printf("\n%v, %T", z2, z2)

	// ignore or do not set value for one of them.
	const (
		_ = iota + 10
		y3
		z3
	)
	fmt.Printf("\n%v, %T", y3, y3)
	fmt.Printf("\n%v, %T", z3, z3)

	// can be used to set value based on bit wise operations
	const (
		isUser = 1 << iota
		isAdmin
		isLocationAdmin
	)
	// first one is 1 left moved on zero; 2nd is 1 left moved of 1 value - that will be 2 and so on.
	fmt.Printf("\n%v, %T", isUser, isUser)
	fmt.Printf("\n%v, %T", isAdmin, isAdmin)
	fmt.Printf("\n%v, %T", isLocationAdmin, isLocationAdmin)

	// we can use it for comparison like this
	var userType byte = isAdmin
	fmt.Printf("\nIs the user an admin ? - %v", isAdmin&userType == isAdmin)
	fmt.Printf("\nIs the user an Location admin ? - %v", isLocationAdmin&userType == isLocationAdmin)
}

func IfCalls() {

	// if like other languages should return a true for execution
	// else will be called if the value returned is false.

	a := 5
	b := 6
	if a == b {
		fmt.Println("a is equal to b")
	} else {
		fmt.Println("a is not equal to b")
	}

	//else should be immediately after the } , it cannot be in the next line as shown below.
	// uncomment and test and you will see a compilation error.
	/* 	if a != b {
	   		fmt.Println("a is not equal to b")
	   	}
	   	else {
	   		fmt.Println("a is  equal to b")
	   	} */

	if a != b {
		fmt.Println("a is not equal to b")
	} else {
		fmt.Println("a is  equal to b")
	}

	// we can also include a statement and then check for the value.
	// it first executes the statement and then checks for true or false
	if a := 5 + 10; a > 15 {
		fmt.Println("a is greater than 15")
	} else {
		fmt.Println("a is less than or equal to 15")
	}

	// else if can also be added.
	if a := 5 + 10; a > 15 {
		fmt.Println("a is greater than 15")
	} else if a == 15 {
		fmt.Println("a is equal to 15")
	} else {
		fmt.Println("a is less than 15")
	}
}

func SwitchCalls() {

	// basic switch is similar to other languages
	// tagged switch - where we compare a variables in switch with various variable values in case.
	var1 := 1
	switch var1 {
	case 1:
		fmt.Println("value is 1")
	case 2:
		fmt.Println("value is 2 ")
	default:
		fmt.Println("value is something else")
	}

	// note that there is no break after a case as in other languages. go compiler adds it automatically
	// for multiple checks, we do not need to add multipole case statement. we can include in one
	// as shown below.
	var2 := 5
	switch var2 {
	case 1, 2, 3:
		fmt.Println("value is in 1, 2, 3")
	case 4, 5, 6:
		fmt.Println("value is 4, 5, 6 ")
	default:
		fmt.Println("value is something else")
	}

	// there is also a tagless switch. There is no variable selected there.
	switch {
	case 1 < 5:
		fmt.Println("1 is less than 5")
	case 1 < 4:
		fmt.Println("1 is less than 4")
	}
	// in this case above, all though both case statements will return true, after the first case statement
	// the system exits. It does not check for the 2nd case.

	// we can check for the type of the variable as well as shown below.
	var v interface{} = 1.0
	switch v.(type) {
	case int:
		fmt.Println("The type of variable is int")
	case float64:
		fmt.Println("The type of variable is float64")
	case string:
		fmt.Println("The type of variable is string")
	default:
		fmt.Println("some other type")
	}

	var v1 interface{} = [3]int{}
	switch v1.(type) {
	case int:
		fmt.Println("The type of variable is int")
	case float64:
		fmt.Println("The type of variable is float64")
	case string:
		fmt.Println("The type of variable is string")
	case [3]int:
		fmt.Println("it is an array of 3 ints")
	case []int:
		fmt.Println("it is a slice")
	default:
		fmt.Println("some other type")
	}

	var v2 interface{} = []int{}
	switch v2.(type) {
	case int:
		fmt.Println("The type of variable is int")
	case float64:
		fmt.Println("The type of variable is float64")
	case string:
		fmt.Println("The type of variable is string")
	case [3]int:
		fmt.Println("it is an array of 3 ints")
	case []int:
		fmt.Println("it is a slice")
	default:
		fmt.Println("some other type")
	}

	// note that type [3]int is different from [2]int from []int (slice)

}
