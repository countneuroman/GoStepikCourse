// Часовая стрелка повернулась с начала суток на d градусов. Определите, сколько сейчас целых часов h и целых минут m.
package main

import "fmt"

func main() {
	var input int
	fmt.Scan(&input)
	hours := input / 30
	minutes := (input % 30) * 2
	fmt.Println("It is", hours, "hours", minutes, "minutes.")
}
