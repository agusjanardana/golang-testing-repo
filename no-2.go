package main


import "fmt"


// buttom up
func fibo(n int) int {
	// your code here
	if n<= 1 {
		return n
	}
	bot_up := make([]int, n+1)
	bot_up[1] = 1
	bot_up[2] = 1
	for i:=3 ;i<=n;i++ {
		bot_up[i] = bot_up[i-1] + bot_up[i-2]
		bot_up = append(bot_up,bot_up[i])
	}
	return bot_up[n]
}


func main() {
	fmt.Println(fibo(0))  // 0
	fmt.Println(fibo(1))  // 1
	fmt.Println(fibo(2))  // 1
	fmt.Println(fibo(3))  // 2
	fmt.Println(fibo(5))  // 5
	fmt.Println(fibo(6))  // 8
	fmt.Println(fibo(7))  // 13
	fmt.Println(fibo(9))  // 13
	fmt.Println(fibo(100))// 55

}