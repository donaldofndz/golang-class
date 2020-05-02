// Every Go program is made up of packages
package main

// The name of the package is the same as the last element of the import path
import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/**
There is another way to import

import "fmt"
import "math"
import "time"
...
*/

/**
In go a name is export if it beign with capital letter,
for example is we do

fmt.Println(math.pi)

it will throw an error

fmt.Println(math.Pi)
*/

/**
To run this we can use the following commands:

- go run hello.go (we are only going to exacute the code)
- go build hello.go (this will create the executable)
	hello
	* to run this we only need to run: ./hello

what is fmt on golang?

fmt mean "from the docs"

Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.
*/

// Programs start running in package main
func main() {
	fmt.Printf("Hello World!\n")
	fmt.Printf("The time is: ", time.Now())
	fmt.Print("\nMy favorite number is: ", rand.Intn(10))
	fmt.Printf("\nNow you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
}
