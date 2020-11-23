/*
Let us assume the following formula for displacement s as a function of time t, acceleration a, initial velocity vo, and initial displacement so.

s =½ a t2 + vot + so

Write a program which first prompts the user to enter values for acceleration, initial velocity, and initial displacement. Then the program should prompt the user to enter a value for time and the program should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn() which takes three float64 arguments, acceleration a, initial velocity vo, and initial displacement so. GenDisplaceFn() should return a function which computes displacement as a function of time, assuming the given values acceleration, initial velocity, and initial displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume the following values for acceleration, initial velocity, and initial displacement: a = 10, vo = 2, so = 1. I can use the following statement to call GenDisplaceFn() to generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print the displacement after 5 seconds.

fmt.Println(fn(5))
*/

package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Enter a, v0 and s0: ")
	in := bufio.NewReader(os.Stdin)
	input, _ := in.ReadString('\n')
	input = input[:len(input)-1]
	numbers := strings.Fields(input)
	sli := make([]float64, 0, len(numbers))
	for i := 0; i < 3; i++ {
		f, err := strconv.ParseFloat(numbers[i], 64)
		if err == nil {
			sli = append(sli, f)
		} else {
			os.Exit(1);
		}
	}
	fn := GenDisplaceFn(sli[0], sli[1], sli[2])
	fmt.Printf("Enter t: ")
	input, _ = in.ReadString('\n')
	input = input[:len(input)-1]
	numbers = strings.Fields(input)
	f, _ := strconv.ParseFloat(numbers[0], 64)
	fmt.Printf("fn(%f) = %f\n",f, fn(f))
}

func GenDisplaceFn(a, v, s float64) func (float64)  float64 {
	f := func (t float64) float64 {
		fmt.Printf("Function : s = 1/2*%f*t^2 + %ft + %f\n", a,v,s)
		return 0.5 * a * t * t + v * t + s
	}
	return f
}
