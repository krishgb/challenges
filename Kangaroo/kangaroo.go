package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func showerror(e error) {
	if e != nil {
		fmt.Fprintf(os.Stdout, "%s", e)
	}
}

func kangaroo(x1 int, v1 int, x2 int, v2 int) string {
	if (x1 > x2 && v1 >= v2) || (x1 < x2 && v1 <= v2) {
		return "NO"
	}

	sum1, sum2 := x1 + v1, x2 + v2
    for {
        if sum1 == sum2 {
            return "YES"
        } else if (x1 > x2 && v1 < v2) && sum1 < sum2{
			return "NO"
		} else if (x1 < x2 && v1 > v2) && sum2 < sum1 {
			return "NO"
		}
		sum1 += v1
		sum2 += v2
    }
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	input, _, err := reader.ReadLine()
	showerror(err)

	arr := strings.Split(strings.TrimSpace(string(input)), " ")
	arrInInt := []int{}
	for i := range arr {
		j, err := strconv.Atoi(arr[i])
		showerror(err)
		arrInInt = append(arrInInt, j)
	}

	x1, v1, x2, v2 := arrInInt[0], arrInInt[1], arrInInt[2], arrInInt[3]
	fmt.Println(kangaroo(x1, v1, x2, v2))
}