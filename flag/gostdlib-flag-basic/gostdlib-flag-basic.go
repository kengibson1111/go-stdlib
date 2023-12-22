package main

import (
	"flag"
	"fmt"
)

const (
	defaultintvar = -1234
	defaultinthelp = "help message for intname2"

	defaultint64var = -12345
	defaultint64help = "help message for int64name2"

	defaultuintvar = 1234
	defaultuinthelp = "help message for uintname2"

	defaultuint64var = 12345
	defaultuint64help = "help message for uint64name2"

	defaultfloat64var = 42.1
	defaultfloat64help = "help message for float64name2"

	defaultstringvar = "foobar"
	defaultstringhelp = "help message for stringname2"

	defaultboolvar = false
	defaultboolhelp = "help message for boolname2"
)

var intptr = flag.Int("intname", defaultintvar, "help message for intname")
var intvar int

var int64ptr = flag.Int64("int64name", defaultint64var, "help message for int64name")
var int64var int64

var uintptr = flag.Uint("uintname", defaultuintvar, "help message for uintname")
var uintvar uint

var uint64ptr = flag.Uint64("uint64name", defaultuint64var, "help message for uint64name")
var uint64var uint64

var float64ptr = flag.Float64("float64name", defaultfloat64var, "help message for float64name")
var float64var float64

var stringptr = flag.String("stringname", defaultstringvar, "help message for stringname")
var stringvar string

var boolptr = flag.Bool("boolname", defaultboolvar, "help message for boolname")
var boolvar bool

func init() {
	flag.IntVar(&intvar, "intname2", defaultintvar, defaultinthelp)
	flag.IntVar(&intvar, "i2", defaultintvar, defaultinthelp + " (shorthand)")

	flag.Int64Var(&int64var, "int64name2", defaultint64var, defaultint64help)
	flag.Int64Var(&int64var, "i642", defaultint64var, defaultint64help + " (shorthand)")

	flag.UintVar(&uintvar, "uintname2", defaultuintvar, defaultuinthelp)
	flag.UintVar(&uintvar, "u2", defaultuintvar, defaultuinthelp + " (shorthand)")

	flag.Uint64Var(&uint64var, "uint64name2", defaultuint64var, defaultuint64help)
	flag.Uint64Var(&uint64var, "u642", defaultuint64var, defaultuint64help + " (shorthand)")

	flag.Float64Var(&float64var, "float64name2", defaultfloat64var, defaultfloat64help)
	flag.Float64Var(&float64var, "f642", defaultfloat64var, defaultfloat64help + " (shorthand)")

	flag.StringVar(&stringvar, "stringname2", defaultstringvar, defaultstringhelp)
	flag.StringVar(&stringvar, "s2", defaultstringvar, defaultstringhelp + " (shorthand)")

	flag.BoolVar(&boolvar, "boolname2", defaultboolvar, defaultboolhelp)
	flag.BoolVar(&boolvar, "b2", defaultboolvar, defaultboolhelp + " (shorthand)")
}

func main() {
	flag.Parse()

	fmt.Println("intptr has value ", *intptr)
	fmt.Println("intvar has value ", intvar)
	fmt.Println("int64ptr has value ", *int64ptr)
	fmt.Println("int64var has value ", int64var)
	fmt.Println("uintptr has value ", *uintptr)
	fmt.Println("uintvar has value ", uintvar)
	fmt.Println("uint64ptr has value ", *uint64ptr)
	fmt.Println("uint64var has value ", uint64var)
	fmt.Println("float64ptr has value ", *float64ptr)
	fmt.Println("float64var has value ", float64var)
	fmt.Println("stringptr has value ", *stringptr)
	fmt.Println("stringvar has value ", stringvar)
	fmt.Println("boolptr has value ", *boolptr)
	fmt.Println("boolvar has value ", boolvar)
}
