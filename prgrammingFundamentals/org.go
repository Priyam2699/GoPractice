package main

import "fmt"

const a int64 = 42
const b float32 = 42.555
const c string = "James Bond"

// or you can write this

const (
	d int32   = 41
	e float64 = 41.555
	f string  = "Indiana Jones"
)

func main() {

	fmt.Println(a, b, c, d, e, f)
}
