# Go: The Complete Developer's Guide (Udemy)
[Course Link](https://www.udemy.com/course/go-the-complete-developers-guide)

[Course Diagrams](https://github.com/StephenGrider/GoCasts/tree/master/diagrams)

[Completed Examples](https://github.com/StephenGrider/GoCasts/tree/master/diagrams)
## Section 2: A Simple Start
### 7. Five Important Questions
#### How do we run the code in our project?
```go build```
* Compiles a bunch of go source code files.

```go run```
* Compiles and executes one or two files.

```go fmt```
* Formats all the code in each file in the current directory.

```go install```
* Compiles and `installs` a package.

```go get```
* Downloads the raw source code of someone else's package.

```go test```
* Runs any tests associated with the current project.

### 8. Go Packages
What does `package main` mean?

```package == project == workspace```
* First line must declare the package it belongs to.
* Two types of packages:
    * `Executable` - Generates a file we can run. The `main` package is executable. Must have `func main`.
    * `Reusable` - Code used as `helpers`. Good place to put reusable logic. All other packages are reusable.
### 9. Import Statements
* imports source code from for example a library (i.e. Reusable Packages).
### 10. File Organization
How is `main.go` organized?
1. `package main` - package declaration
2. `import statement` - Import other packages that our program depends on
3. `func main(){}` - Declare functions, tell Go to do something.
## Section 3: Deeper Into Go
### 12. Project Overview
Cards Program
* `newDeck` - Create a list of playing cards. Essentially an array of strings.
* `print`  - Log out the contents of a deck of cards.
* `shuffle` - Shuffles all the cards in a deck.
* `deal` - Create a "hand" of cards.
* `saveToFile` - Save a list of cards to a file on the local machine.
* `newDeckFromFile` - Load a list of cards from the local machine.
### 14. Variable Declaration
#### Declaring
`var card string = "Ace of Spades"`
or
`card := "Ace of Spades"`

Go will _infer_ a type with `:=`. Only use when initializing a new variable, otherwise use `=`

* `var` - We're about to create a new variable.
* `card` - The name of the variable will be card.
* `string` - only a string will ever be assigned to this variable.
  * Go is a `Statically Typed Language` vs `Dynamically Typed Language`.
    * In a `Dynamically Typed Language` we don't care what types are assigned to any given variable. The interpreter will handle that for us.
    * [Statically vs Dynamically Typed Languages](https://android.jlelse.eu/magic-lies-here-statically-typed-vs-dynamically-typed-languages-d151c7f95e2b)
* `= "Ace of Spades"` - Assign `Ace of Spades` as the value of `card`.
* Variables can be initialized outside of a function but cannot be assigned a a variable.

#### Go Fundamental Types
`bool` - true, false

`string` - "Ace of Spades"

`int` - 1, 2, 3

`float64` - 10.0001, 0.0008, -1000.2323

## 15. Functions and Return Types
```$xslt
func newCard() string {
	return "Five of Diamonds"
}
```
* `newCard` - define function called `newCard`
* `string` - When executed, this function will return a value of type `string`.

## 16. Slices and For Loops
* Types of Arrays
    * `Array` - Fixed length list of things.
    * `Slice` - An array that can grow or shrink.
        * Every element in a slice must be of the same data type.
### Syntax: Slices
```$xslt
cards := []string{"Ace of Diamonds", newCard()}
cards = append(cards, "Six of Spades")
```
* `cards := []string{"Ace of Diamonds", newCard()}` assign new slice with string type the values of "Ace of Diamonds" and the value of `newCard` function.

### Syntax: `for` Loops
```$xslt
for i, card := range cards {
	fmt.Println(i, card)
}
```
* `i` - index of this element in array.
* `card` - current card we are iterating over.
* `:= range cards` - Take the slice of cards and loop over it.
* `fmt.Println(card)` - Run this one time for each card in the slice.

## 17. OO Approach vs Go Approach
* We take basic data types and `extend\borrows` those types vs a OO Aproach with classes, methods, properties.
* A Go approach would be to create a new type like `type deck []string`
  * our deck type is a slice of strings.
* Functions with `deck` as a `receiver`.
  * A function with a receiver is like a `method` (function that belongs to a class instance.)

## 18. Custom Type Declarations
* We are declaring a type of `deck` which is a slice of strings.
#### Example
```$xslt
type deck []string
cards := deck{"Ace of Diamonds", newCard()}
```

## 19. Receiver Functions
* In `deck.go`:
```$xslt
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
```
* This line `func (d deck) print() {`
  * Any variable of type `deck` now gets access to the `print method`
    * `d` - The actual copy of the deck we're working with is available in the function as a variable called `d`.
      * you can view this receiver function `d` as python's `self` or JavaScript's `this`.
      * By convention we use a single or abbr. of type (`deck`).
    * Every variable of the type `deck` can call this function on itself.
* In `main.go`:
```$xslt
cards.print()
```

## 20. Creating a New Deck
```$xslt
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}
	return cards
}
```
* replace `index` with `_` telling Go we don't care about the index variable.
## 21. Slice Range Syntax
* Slices are zero-indexed
### Example
```$xslt
fruits[startIndexIncluding : upToNotIncluding]
fruits[0:2] or fruits[:2]
```
## 22. Multiple Return Values
In `deck.go`
```$xslt
func deal(d deck, handsize int) (deck, deck){
	// return two values of type deck.
	return d[:handsize], d[handsize:]
}
```
In `main.go`
```$xslt
func main() {
	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)
	hand.print()
	remainingDeck.print()
}
```
## 24. Deck to String
### Type Conversion
`[]byte("Hi there!")` 
or
```$xslt
greeting := "hi there"
fmt.Println([]byte(greeting))
```


`[]byte` - Type we want
`"Hi there!"` - Value we have

## 25. Joining A Slice of Strings
```$xslt
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
```

## 26. Saving Data to the Hard Drive
[io/util - WriteFile](https://golang.org/pkg/io/ioutil/#example_WriteFile)
```$xslt
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
```
* receiver `(d deck)`

## 27. Reading From the Hard Drive
```$xslt
func ReadFile(filename string) ([]byte, error)
```
* if no error then return value of error will be `nil`.
## 28. Error Handling 
```$xslt
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ", ")
	return deck(s)
}
```
## Section 4: Organizing Data with Structs
### Struct
* Data structure. A collection of properties that are related together.
### 44. Pointer Operations
```
&variable - Give me the memory address of the value this variable is pointing at
*pointer - Give me the value this memory address is pointing at.
```
#### Breakdown
```
func (pointerToPerson *person) updateName(){
    *pointerToPerson
}
*person = This is the type description, it means we are working with a pointer to a person.
*pointerToPerson = this is an operator, it means we want to manipulate the value the pointer is referencing.
```

* Turn `address` into `value` with `*address`
* Turn `value` into `address` with `&value`
### 47. Reference vs. Value Types
#### Value Types
* Use pointers to change these in a function:
  * int
  * float
  * string
  * bool
  * structs
#### Reference Types
* Don't worry about pointers with these types
  * slices
  * maps
  * channels
  * pointers
  * functions