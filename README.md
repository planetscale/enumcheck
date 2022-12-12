# enumcheck

> PlanetScale: Temporary fork of enumcheck due to Go module cache issues

***This is still a WIP, so exact behavior may change.***

Analyzer for exhaustive enum switches.

To install:

```
go install github.com/planetscale/enumcheck@latest
```

This package reports errors for:

``` go
//enumcheck:exhaustive
type Letter byte

const (
	Alpha Letter = iota
	Beta
	Gamma
)

func Switch(x Letter) {
	switch x { // error: "missing cases Beta, Gamma and default"
	case Alpha:
		fmt.Println("alpha")
	case 4: // error: "implicit conversion of 4 to Letter"
		fmt.Println("beta")
	}
}

func Assignment() {
    var x Letter
    x = 123 // error: "implicit conversion of 123 to Letter
}

```

This can also be used with types:

``` go
//enumcheck:exhaustive
type Expr interface{}

var _ Expr = Add{}
var _ Expr = Mul{}

type Add []Expr
type Mul []Expr

type Invalid []Expr

func Switch(x Expr) {
	switch x.(type) { // error: "missing cases Mul"
	case Add:
		fmt.Println("alpha")
	case Invalid: // error: "implicit conversion of Invalid to Expr"
		fmt.Println("beta")
	default:
		fmt.Println("unknown")
	}
}

func Assignment() {
	var x Expr
	x = 3 // error: "implicit conversion of 3 to Expr
	_ = x
}
```

Or with structs:

``` go
//enumcheck:exhaustive
type Option struct{ value string }

var (
	True  = Option{"true"}
	False = Option{"false"}
	Maybe = Option{"maybe"}
)

func DayNonExhaustive() {
	var day Option

	switch day { // want "missing cases False, Maybe and default"
	case Option{"invalid"}: // want "invalid enum for enumstruct.Option"
		fmt.Println("beta")
	case True:
		fmt.Println("beta")
	}
}
```

Mode `//enumcheck:relaxed` allows to make "default" case optional:

``` go
//enumcheck:relaxed
type Option string

var (
	Alpha = Option("alpha")
	Beta  = Option("beta")
)

func Relaxed() {
	var day Option
	switch day {
	case Alpha:
		fmt.Println("alpha")
	case Beta:
		fmt.Println("beta")
	}
}
```

Mode `//enumcheck:silent` allows to silence reports for switch statements:

``` go
//enumcheck:silent
type Option string

var (
	Alpha = Option("alpha")
	Beta  = Option("beta")
)

func NoErrorHere() {
	var day Option
	switch day {
	case Beta:
		fmt.Println("beta")
	}
}

func EnablePerSwitch() {
	var day Option
	switch day { //enumcheck:exhaustive
	case Beta:
		fmt.Println("beta")
	}
}
```
