# Makeshift

Makeshift is a simple library for Go that makes task automation super easy. Think of it as a lightweight alternative to Makefiles, but without all the extra complexity. It’s perfect for when you just want to define some tasks, run them from the command line, and keep things straightforward.

## What It Does

Makeshift lets you:

* Define tasks in Go: Write your commands directly in your Go code.
* Run tasks easily: Execute them from the command line with no fuss.
* Keep it simple: No heavy build tools or complicated setups—just lightweight and easy.

## Why It’s Cool

* Works everywhere: Runs on any platform that supports Go.
* Written in Go: It’s fast, reliable, and feels right at home in your Go projects.
* Super simple: No steep learning curve—just define tasks and go!

Makeshift is for developers who want a no-nonsense way to automate tasks without dealing with the overhead of tools like Make or CMake.
    
## Installation

To include the `makeshift` library in your Go project, use the following command:

```bash
go get github.com/lzcdr/makeshift
```

## Usage

Defining Commands
To define commands, embed the Task struct in your own struct. Each method you define on that struct corresponds to a command that can be executed.
### File "./example/make/make.go":

```go
package main

import (
	"fmt"

	"github.com/lzcdr/makeshift"
)

// Task embeds makeshift.Task to define targets
type Task struct {
	makeshift.Task
}

// Build builds the project
func (t Task) Build() error {
	fmt.Println("Building the project...")
	return makeshift.ExecCommand("go", "build", "-o", "myapp.exe", "./example/cmd")
}

// Test runs tests
func (t Task) Test() error {
	fmt.Println("Running tests...")
	return makeshift.ExecCommand("go", "test", "./...")
}

// Clean cleans build artifacts
func (t Task) Clean() error {
	fmt.Println("Cleaning build artifacts...")
	return makeshift.Remove("-r", "-f", "myapp.exe")
}

func main() {
	task := Task{}
	makeshift.DoIt(task)
}
```
## Running Commands

To run a specific command, execute your Go program with the command name as an argument:

```bash
go run ./example/make/make.go Build
```

If you run the program without any arguments, it will list the available commands:

```bash
go run ./example/make/make.go
```

## Available Functions

* ExecCommand(command string, args ...string) error: Executes a shell command with the given arguments and logs the output.

* Remove(args ...string) error: Removes files and directories based on provided flags and paths. Supported flags: -r (recursive), -f (force - do not print or interrupt on errors), -i (interactive - ask user confirmations), -v (verbose). 

* DoIt(task interface{}): Handles the logic for listing or running a specific command.
