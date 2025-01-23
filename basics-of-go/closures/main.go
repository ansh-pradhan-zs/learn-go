package main

import "fmt"

func fib() func() int {
	a,b := 0,1
	return func() int {
		a,b = b, a+b
		return b
	}
}

func main()  {
	fmt.Println("Learning Closures")

	f := fib()

	// for x := f(); x < 100; x = f(){
	// 	fmt.Println(x)
	// }

	g := f // ? both the f and the g functions are pointing to the same environment pointer of the function, hence they share the variables a and b

	fmt.Println(f(), f(), f(), f(), f())
	fmt.Println(g(), g(), g(), g(), g())


}