package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func showerror(e error) {
	if e != nil {
		panic(e)
	}
}

func twod(a [][]int) int {
	sum := int(math.Inf(-1))
	for i:= range a {
		for j := range a[i] {
			old := 0
			if j+2 < len(a[i]) && i+2 < len(a[i]) {
				old += a[i][j] + a[i][j+1] + a[i][j+2] + a[i+1][j+1] + a[i+2][j] + a[i+2][j+1] + a[i+2][j+2]
				if old > sum {
					sum = old
				}
			}
		}
	}
	return sum

}

func getInput(limit int) []int{
	arr := []int{}

	input := bufio.NewReader(os.Stdin)
	values, err := input.ReadString('\n')
	showerror(err)
	values = strings.TrimSpace(values)
	v := strings.Split(values, " ")[:limit]
	for _, j := range v{
		valuesInInt, err := strconv.Atoi(j)
		showerror(err)
		arr = append(arr, valuesInInt)

	}
	return arr
}

func main() {

	twoDArray := [][]int{}

	fmt.Print("Enter your array dimension if(6x6) then enter 6: ")
	input := bufio.NewReader(os.Stdin)

	value, err := input.ReadString('\n')
	showerror(err)
	value = strings.TrimSpace(value)

	if value == "" {
		fmt.Println("------ no numbers detected ------")
		os.Exit(1)
	}
	
	limit, err := strconv.ParseInt(value, 10, 64) 
	showerror(err)

	for i:=0; i < int(limit); i++{
		arr := getInput(int(limit))
		twoDArray = append(twoDArray, arr)
	}

	result := twod(twoDArray)
	fmt.Println("The result is", result)
}
  