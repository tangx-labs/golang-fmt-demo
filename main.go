package main

import (
	"fmt"
)

type Person struct {
}

func (p *Person) String() string {
	return "oh my god: interface -> Stringer "
}

func (p *Person) Error() string {
	return "oops, i'm sick interface-> Error "
}

func (p *Person) Format(s fmt.State, verb rune) {
	str := []byte("holy , interface Formatter")
	_, _ = s.Write(str)
}

func main() {
	p := &Person{}

	fmt.Println("fmt.Println=", p)
	fmt.Printf("fmt.Printf /v: %v\n", p)
	fmt.Printf("fmt.Printf /s: %s\n", p)
	fmt.Printf("fmt.Printf /+v: %+v\n", p)

	fmt.Println(p.String())
	fmt.Println(p.Error())

}
