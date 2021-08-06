package demo_vector

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func Op1() {
	// create a flat representation of our matrix
	components := []float64{1.2, -5.7, -2.4, 7.3}
	a := mat.NewDense(2, 2, components)

	// As a sanity check, output the matrix to standard out.
	fa := mat.Formatted(a, mat.Prefix(" "))
	fmt.Printf("mat = %v\n\n", fa)
}
