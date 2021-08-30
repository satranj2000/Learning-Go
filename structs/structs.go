package structs

import (
	"fmt"
	"strings"
)

type Identifiable interface {
	ID() string
}

type Person struct {
	Firstname     string
	Lastname      string
	twitterHandle string
}

//
// type Person struct {
// 	firstname string // small first character means this attribute can be called only within this file (private) and not outside
// 	lastname  string
// }

type name struct {
	first string
	last  string
}

type Person2 struct {
	name
	twitterHandle string
}

func (p Person) ID() string {
	return "1111"
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.Firstname, p.Lastname)
}

// use the substructure or embedded struct variables directly from the main struct
func (p *Person2) FullName() string {
	return fmt.Sprintf("%s %s", p.first, p.last)
}

// this function won't set the Person's twitter handle as it is a copy of the person struct. It should instead be a pointer.
// see the same function below with a pointer.

// func (p Person) SetTwitterHandle(handle string) {
// 	p.twitterHandle = handle
// }

func (p *Person) SetTwitterHandle(handle string) {
	p.twitterHandle = handle
}

func (p *Person) GetTwitterHandle() string {
	return p.twitterHandle
}

// use type definition. Define TwitterHandle as string.
type TwitterHandle string

type Person3 struct {
	name
	twitterHandle TwitterHandle
}

func (p *Person3) SetTwitterHandle(th TwitterHandle) {
	p.twitterHandle = th
}

func (p *Person3) GetTwitterHandle() TwitterHandle {
	return p.twitterHandle
}

func (th *TwitterHandle) GetRedirectUrl() string {
	//url := strings.Trim(*th, "@") - cannot be used as Trim takes a string parameter ; therefore, need to convert
	url := strings.Trim(string(*th), "@")
	fullurl := fmt.Sprintf("https://twitter.com/%v", url)
	return fullurl
}

func CallPerson() {
	// since
	p := Person{Firstname: "sathish", Lastname: "Ram"}
	fmt.Println(p.Firstname)
}

func CallPerson2() {
	p := Person{Firstname: "ranjani", Lastname: "Sathish"}
	p.SetTwitterHandle("@ranj")
	fmt.Println(p.GetTwitterHandle()) // will this show the twitter handler that was set.
}

func CallPersonwithembeddedStruct() {
	var p Person2
	p.first = "Jack" // directly set the value for name sub structure or embedded structure
	p.last = "darcy"
	fmt.Println(p.FullName())
}

func CallPersonUseNewType() {
	var p Person2
	p.first = "Jack" // directly set the value for name sub structure or embedded structure
	p.last = "darcy"
	p.twitterHandle = "@jack"
	fmt.Println(p.twitterHandle)
}

func CallPersonwithNewType2() {
	var p Person3
	p.first = "Jack" // directly set the value for name sub structure or embedded structure
	p.last = "darcy"
	p.SetTwitterHandle("@jdarcy") // the parameter here is Twitterhandler and not string.
	fmt.Println(p.GetTwitterHandle())
	twithand := p.GetTwitterHandle()
	fmt.Println(twithand.GetRedirectUrl())

}
