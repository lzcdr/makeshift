## Hey everyone!

I just wanted to share a little project I've been working on: a simple library for Go called Makeshift. The idea behind it was to create something really easy to use for task automation without all the complexity that often comes with other build tools.

# What It Does

Makeshift allows you to:

* Define simple commands in your Go code
* Execute them easily from the command line
* Keep everything lightweight and straightforward

# Getting Started

You can install it with:

```bash
go get github.com/lzcdr/makeshift
```

# Here’s a quick example of how it works:

```go
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
	return makeshift.ExecCommand("rm", "-rf", "myapp")
	return nil
}

func main() {
	task := Task{}
	makeshift.DoIt(task)
}
```

To run a specific command, execute your Go program with the command name as an argument:

```bash
go run main.go Build
```

If you run the program without any arguments, it will list the available commands:

```bash
go run main.go
```

# Why Use It?

I designed Makeshift because I wanted something that requires minimal setup and still gets the job done. If you're looking for a simple way to organize your Go tasks without the overhead of more complex systems, give it a try!

# Feedback?
I would love any feedback you have or any suggestions for features! Let me know what you think or if you have any questions.

Thanks for checking it out! 😊
