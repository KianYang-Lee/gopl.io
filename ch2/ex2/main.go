/*
Exercise 2.2: Write a general-purpose unit-conversion program analogous to cf that reads
numbers from its command-line arguments or from the standard input if there are no arguments,
and converts each number into units like temperature in Celsius and Fahrenheit,
length in feet and meters, weight in pounds and kilograms, and the like.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/tempconv"
)

type Feet float64
type Meter float64
type Pound float64
type Kilogram float64

func (f Feet) String() string      { return fmt.Sprintf("%g ft.", f) }
func (m Meter) String() string     { return fmt.Sprintf("%gm", m) }
func (p Pound) String() string     { return fmt.Sprintf("%g pound", p) }
func (kg Kilogram) String() string { return fmt.Sprintf("%g kg", kg) }

func main() {
	args := os.Args[1:]

	// Take from Stdin if no arguments are provided
	if len(args) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			args = append(args, input.Text())
		}
	}

	fmt.Println(args)
	for _, arg := range args {
		raw, err := strconv.ParseFloat(arg, 64)
		// Exit if error during parsing, CODE 1
		if err != nil {
			fmt.Printf("ex2: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(raw)
		c := tempconv.Celsius(raw)
		ft := Feet(raw)
		m := Meter(raw)
		p := Pound(raw)
		kg := Kilogram(raw)
		fmt.Printf("Raw value of [%g] converted to :\n%v\n%v\n%v\n%v\n%v\n%v\n", raw, f, c, ft, m, p, kg)

	}
}
