package main

import (
	"fmt"

	"github.com/lzcdr/makeshift"
)

// Task embeds gomake.Task to define targets
type Task struct {
	makeshift.Task
}

// Build builds the project
func (t Task) Build() error {
	fmt.Println("Building the project...")
	return makeshift.ExecCommand("go", "build", "-o", "myapp", ".")
}

// Test runs tests
func (t Task) Test() error {
	fmt.Println("Running tests...")
	return makeshift.ExecCommand("go", "test", "./...")
}

// Clean cleans build artifacts
func (t Task) Clean() error {
	fmt.Println("Cleaning build artifacts...")
	return makeshift.Remove("-r", "-i", "myapp")
}

func main() {
	task := Task{}
	makeshift.DoIt(task)
}
