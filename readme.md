# Makeshift

`makeshift` is a simple and flexible Go library that serves as an alternative to Makefiles, allowing you to define tasks and execute command-line instructions in a straightforward manner. It is designed to be cross-platform, leveraging the simplicity and portability of Go. With reflection-based target execution and colorful output logging, `makeshift` makes command execution easy and accessible without the need for additional tools like Make, CMake, or similar build systems.

## Key Features

- **Cross-Platform**: As portable as Go itself, running seamlessly on any platform that supports Go.
- **Written in Go**: The entire library is implemented in Go, ensuring performance and reliability.
- **Super Simple**: Designed for developers looking for a straightforward replacement for Makefiles without the complexity of additional build systems.
    
## Installation

To include the `makeshift` library in your Go project, use the following command:

```bash
go get github.com/lzcdr/makeshift
```

## Usage

Defining Commands
To define commands, embed the Task struct in your own struct. Each method you define on that struct corresponds to a command that can be executed.

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
	return makeshift.Remove("-r", "-f", "myapp")
	return nil
}

func main() {
	task := Task{}
	makeshift.DoIt(task)
}
```
## Running Commands

To run a specific command, execute your Go program with the command name as an argument:

```bash
go run main.go Build
```

If you run the program without any arguments, it will list the available commands:

```bash
go run main.go
```

## Available Functions

* ExecCommand(command string, args ...string) error: Executes a shell command with the given arguments and logs the output.

* Remove(args ...string) error: Removes files and directories based on provided flags and paths. Supported flags: -r (recursive), -f (force - do not print or interrupt on errors), -i (interactive - ask user confirmations), -v (verbose). 

* Run(task interface{}, targetName string) error: Discovers and executes the specified command method using reflection.

* ListTargets(task interface{}): Lists all available commands (methods) for the task.

* DoIt(task interface{}): Handles the logic for listing or running a specific command.
