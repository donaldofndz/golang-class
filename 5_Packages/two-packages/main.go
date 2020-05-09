package main 

import "fmt"
import "golang-class/5_Packages/two-packages/cryptocurrency"

func main()  {
	fmt.Println("Hey, we are working at two-packages")
	fmt.Println(cryptocurrency.ConvertDollarsToBitcoin(1.02))
}