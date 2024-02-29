# todo: create a provisioner

# Introduction / Fundamentals

## Go Runtime

- Allows build, compile and execute. 
- Package == Project == Workspace

## Go CLI main commands

- `go build`   - compiles and generate a binary.
- `go run`     - can take one or more files, compile and run.
- `go fmt`     - formats code in different files.
- `go install` - compiles and installs a package.
- `go get`     - downloads raw source code.
- `go test`    - run tests associated with current project.

## Package

- Is a collection of `.go` source code.
- Package can have many files.
- Every file must declare in which package it is contained, on the first line.
- There are two types of packages: executable and reusable
- The name of the package determines if you are making an executable or dependency type package
- Files in the same package can freely call functions defined in other files.

```go
package main

import "fmt" // get all code and functionalities from fmt package

func main() {
	fmt.Println("hi")
}
```

> Why main? because it creates an executable type package!

## Executable Package

Generates a runnable file.

Must have always a function called `func main()`.

## Reusable Package

Used as "helpers", good to put reusable logic.

## Basic Types

- bool
- string
- int
- float64
> type conversion: []byte("hi there")
> zero value of strnig "", int 0, float 0, bool false

## Array and Slices

Array - fixed length
Slice - can grow or shrink

Defined with a data type and all records must have that type.

To create a slice you can use:

```go
cards := []string{"card0", "card1"}
```

To add use `cards = append(cards, "new_element")`. It does not modify, it returns a new slice.

To loop use

```go
for i,element := range cards {
    fmt.Println(i, element)
}
```
> range is a keyword used to iterate

Using [startIndexIncluding:upToNotIncluding]

# Card Deck Project

- newDeck         -> creates a list of cards
- print           -> log out the contents of the deck
- shuffle         -> shuffles all card in the deck
- deal            -> create a hand of cards
- saveToFile      -> save list of card to disk
- newDeckFromFile -> load list of cards from disk

OOP approach: class deck, instance contains methods and attributes
GO approach: we can define a new type called deck, and it will be a slice of strings. We can attach functions to the custom deck, called *functions witime.Now().UnixNano()th a receiver*.

`main.go` - create and manipulate deck
`deck.go` - describes what a deck is and how it works
`deck_test.go` - automated tests


Function with receiver example:
```go
func (d deck) print() {

}
```
> A receiver tells that all variables of type deck, gets access to the print method. Usually a two/one variable letter that matches the type of the receiver.

# Tests

To enable tests we must create a `go.mod` file. we can do this using `go mod init cards`

Is not like using testing frameworks, like mocha. We get a small set of functions, usually we just use go language to test and make sure the code is running correctly.

# Pointer

a pointer on type is different than a pointer to a pointer.
a star where a type should be, means it is a pointer to the type
a start on a field, is dereferencing, turns the pointer into value

```go
func (pointerToPerson *person) updateName() {
    *pointerToPerson // operator, it means we manipulate the value the pointer is referencing
}
```

Go has a shortcut you can use

```go
func (p *person) updateName(name string) {
    p.firstName = name
}
```

If we define a receiver with pointer to some type, when we call the function we can call with the pointer (&) or the root type.

example:

jimPointer = &jim
jimPointer.updateName

or

jim.updateName


## Gotchas

Struct and Slices behave differently when passing as paramaters

```go
func updateSlice(s []string) {
    s[0] = "c"
}
slice := []string{"a", "b"}
updateSlice(slice)
// slice[0] would be "c"
```

Since an slice is a ptr to a head, a cap and length, when we pass it as parameter, we are passing the pointer (which will be copied) but the pointer will point to the same array.

We call this a `reference type`.

Other reference type are: Maps, Channels, Pointers, Functions
Value types: int, float, string, bool, structs.

# Map

Collection of key:value

all keys must have the same type
all values must have the same type but can be different than key type


Some differences from sstruct:

You can iterate over keys as they are indexed.
Reference type.
Dont need to know all keys at compile time.
Represent collection of related properties.

# Interfaces

Form of reusing code

We create a new custom type. Other subtypes can use this interface (new type) to share common logic with function that have the same name and return type. In the example at `interfaces` dir,  english bot satisfies bot interface, so it is accepted into printGreeting using bot as type.
It will look at receviver functions and try to match the name with return type, if it matches, than you can consider that type part of interface.

```go
type <type_name> interface {
    <func_name>(arg_type_1, arg_type_2..) (return_type1, return_type2..)
}
```

> **Concrete Type vs Interface Type**: concrete type we can create a value out directly, like map of struct. with interface type you can't create a value out of it directly.

**Interfaces are not generic types**. Go does not have a generic type;
**Interfaces are implicit**. We don't have to say custom types satisfies the interface.
**Interfaces are a contract**. It helps us manage types.

Reader interface takes a byte slice and modifies it, then returns err and n
Writer interfaces takes byte slice as input and sources it to a destination

io.Copy(dst Writer, src Reader) , means it receives something that implemest writer interface and something that implements reader interface, for example can be used as io.Copy(os.Stdout, resp.Body) so the types we ssent must implement writer/reader interface.

os.Stdout is a value of type file, file has  a function called Write and implements the interface.

# Go Routine

When we run a go process, we automatically create a Go routine. Go routine executes our code line by line.
Using "go" keyword will create a new go routine.

Go has its own scheduler, and can run on a single core CPU.

Scheduler runs a routine until it finishes or makes a blocking call (like http request)

By default go tries to use one core, in multi-core ambient, go scheduler will run one thread on each logical core.

The main routine is the only one who exits. If main routine finishes before child routines, child routines won't be executed.

## Channels

Channels are used to communicate between routines. Channels can have type and you can only send messages of that type on the channel.

```go
c := make(chan string)
```

When passing a channel to a function, you must specify the type of data on channel, example: `checkLink(link string, c chan string)`



### Sending Data

- channel <- 5 - sned value 5 to channel
- myNumber <- channel - wait for a value to be send into the channel, assign to mynumber
- fmt.Println(<- channel) - wait for a value to be sent in the channel and log it

### Repeating Routines

wait for the channel c to return a value, assign to l and run the body
```go
for l := range c {
   go checkLink(l, c)
}
```

or

```go
for {
	go checkLink(<-c, c)
}
```


