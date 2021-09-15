package main

import (
	"fmt"
	"math"
)


func Frog(jumps []int) int {

	// your code here

	// karena N max 10^5
	var td [100000]float64
	td[0] = 0
	td[1] = math.Abs(float64(jumps[2] - jumps[3]))

	for i := 2; i< len(jumps);i++ {
		temp1 := math.Abs(float64(jumps[i] - jumps[i-1]))
		temp2 := math.Abs(float64(jumps[i] - jumps[i-2]))
		td[i] = math.Min(float64(td[i-1])+temp1, float64(td[i-2])+temp2)
	}
	return int(td[len(jumps)-1])
	//td[1] = 20
	//
	//iterasi 2 :
	//temp 1 : 40 - 30 = 10
	//temp 2 : 40 - 10 = 30
	//td[2] : td[1]+10 , td[0] + 30 ?? 30, td[2] = 30
	//
	//iterasi 3 :
	//temp1 : 20 - 40 = 20
	//temp2 : 20 - 30 = 10
	//td[3] : td[2]+20, td[1]+10 , 30 juga
	//
	//td[0,20,30,30]
}


func main() {

fmt.Println(Frog([]int{10, 30, 40, 20})) // 30
fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40

}