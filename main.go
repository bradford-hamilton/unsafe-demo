package main

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/bradford-hamilton/uintptr-demo/pkg/rand"
)

type user struct {
	name string
	age  int
}

func main() {
	u := user{}
	fmt.Println(u) // {name: age:0}

	namePtr := (*string)(unsafe.Pointer(&u))
	fmt.Println(namePtr) // 0xc00000c080

	*namePtr = "bradford" // update the value pointed at by namePtr

	name := *(*string)(unsafe.Pointer(&u))
	fmt.Println(name) // bradford

	d := rand.NewDog("brown", 40, "mrDog") // create new dog with all private fields
	dogNamePtr := (*string)(unsafe.Pointer(
		uintptr(unsafe.Pointer(&d)) + unsafe.Sizeof(int(0)) + unsafe.Sizeof(string("")),
	))

	fmt.Println(*dogNamePtr) // mrDog
	*dogNamePtr = "someOtherDog"
	fmt.Println(d) // {brown 40 someOtherDog}

	pup := rand.NewPuppy("red", 80, "toby")
	fmt.Println(pup) // {{red 80 toby} carol baskin}

	dogNameBytes := (*[]byte)(unsafe.Pointer(
		uintptr(unsafe.Pointer(&pup)) + uintptr(unsafe.Sizeof(rand.Dog{})),
	))
	fmt.Println(string(*dogNameBytes)) // carol baskin

	*dogNameBytes = []byte("bradford")
	fmt.Println(pup) // {{red 80 toby} bradford}

	b := []byte("josce lamson-scribner")
	fmt.Println(b) // [106 111 115 99 101 32 108 97 109 115 111 110 45 115 99 114 105 98 110 101 114]

	s := ""
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	hdr.Data = uintptr(unsafe.Pointer(&b[0]))
	hdr.Len = len(b)

	fmt.Println(s) // josce lamson-scribner
}

// - A pointer value of any type can be converted to a Pointer.
// - A Pointer can be converted to a pointer value of any type.
// - A uintptr can be converted to a Pointer.
// - A Pointer can be converted to a uintptr.
