package main

import (
	s "github.com/inancgumus/prettyslice"
)

// === Notes === //

// make() function
// func main() {
// 	s.PrintBacking = true
// 	s.MaxPerLine = 10

// 	tasks := []string{"jump", "run", "read"}

// 	upTasks := make([]string, 0, len(tasks))
// 	s.Show("upTasks", upTasks)

// 	for _, task := range tasks {
// 		upTasks = append(upTasks, strings.ToUpper(task))
// 		s.Show("upTasks", upTasks)
// 	}
// }

// copy() function
func main() {
	s.PrintBacking = true
	s.MaxPerLine = 10

	// tasks := []string{"jump", "run", "read"}

	// upTasks := make([]string, 0, len(tasks))
	// s.Show("upTasks", upTasks)

	// for _, task := range tasks {
	// 	upTasks = append(upTasks, strings.ToUpper(task))
	// 	s.Show("upTasks", upTasks)
	// }
}
