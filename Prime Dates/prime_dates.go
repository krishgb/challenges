package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func intDays(s []string) []int {

	s_day, _ := strconv.ParseInt(s[:1][0], 10, 64)
	s_month, _ := strconv.ParseInt(s[1:2][0], 10, 64)
	s_year, _ := strconv.ParseInt(s[2:][0], 10, 64)

	return []int{int(s_day), int(s_month), int(s_year)}
}

func intString(s string) int {
	i, _ := strconv.ParseInt(s, 10, 64)
	return int(i)
}

func getLuckyDays(s string, s1 string) ([]string, [][]string, [][][]string) {
	n1 := intDays(strings.Split(s, "-"))
	n2 := intDays(strings.Split(s1, "-"))

	n1_day := n1[0]
	n1_month := n1[1]
	n1_year := n1[2]

	n2_day := n2[0]
	n2_month := n2[1]
	n2_year := n2[2]

	years := []string{}
	months := [][]string{}
	days := [][][]string{}

	monArr := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}
	daysArr := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31"}
	for i := n1_year; i <= n2_year; i++ {

		// months string
		if n1_year == n2_year {
			months = append(months, monArr[n1_month-1:n2_month])
		} else if i == n1_year {
			months = append(months, monArr[n1_month-1:])
		} else if i < n2_year {
			months = append(months, monArr)
		} else {
			months = append(months, monArr[:n2_month])
		}

		// years string
		years = append(years, strconv.Itoa(i))
	}

	m := n1_month

	for j := 0; j < len(years); j++ {
		yearIn, _ := strconv.Atoi(years[j])
		leapYear := false

		if years[j][2:] == "00" {
			firstTwo := intString(years[j][:2])
			if firstTwo%4 == 0 {
				leapYear = true
			}
		}

		y := [][]string{}
		for i := 0; i < len(months[j]); i++ {

			if n1_day == n2_day && n1_year == n2_year && n1_month == n2_month {
				y = append(y, daysArr[n1_day-1:n2_day])
				break
			}
			if yearIn == n1_year && m == n1_month {
				if yearIn%4 == 0 && m == 2 {
					if years[j][2:] == "00" && !leapYear {
						y = append(y, daysArr[:28])
					} else {
						y = append(y, daysArr[n1_day-1:29])
					}
				} else if m == 2 {
					y = append(y, daysArr[n1_day-1:28])
				} else if m == 1 || m == 3 || m == 5 || m == 7 || m == 8 || m == 10 || m == 12 {
					y = append(y, daysArr[n1_day-1:])
				} else {
					y = append(y, daysArr[n1_day-1:30])
				}

			} else if yearIn == n2_year && m == n2_month {

				y = append(y, daysArr[:n2_day])

			} else if yearIn%4 == 0 && m == 2 {
				if years[j][2:] == "00" && !leapYear {
					y = append(y, daysArr[:28])
				} else {
					y = append(y, daysArr[:29])
				}
			} else {
				if m == 2 {
					y = append(y, daysArr[:28])
				} else if m == 1 || m == 3 || m == 5 || m == 7 || m == 8 || m == 10 || m == 12 {
					y = append(y, daysArr)
				} else {
					y = append(y, daysArr[:30])
				}
			}
			m++
			if m == 13 {
				m = 1
			}
		}
		days = append(days, y)
	}

	return years, months, days
}

func giveDatesArr() ([]int, int) {
	reader, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	reader = strings.TrimSpace(reader)
	n := strings.Split(reader, " ")

	y, m, d := getLuckyDays(n[0], n[1])


	dates := []int{}
	luckyDays := 0
	for i, ye := range y {
		for j, mo := range m[i] {
			for _, da := range d[i][j] {
				date := intString(da + mo + ye)
				dates = append(dates, date)
				if date%4 == 0 || date%7 == 0 {
					luckyDays++
				}
			}
		}
	}

	return dates, luckyDays
}

func main() {

	fmt.Print("Enter the dates in the format \"dd-mm-yyyy\" with space between them: ")

	_, i := giveDatesArr()

	fmt.Println("LuckyDays: ",i)
}
