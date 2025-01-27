package makeshift

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"

	"github.com/fatih/color"
)

// Task is an empty struct that users embed in their make.go file
type Task struct{}

// Remove removes files and directories based on provided flags and paths.
func Remove(args ...string) error {
	recursive := false
	force := false
	interactive := false
	verbose := false

	var paths []string

	// Parse the arguments
	for _, arg := range args {
		switch arg {
		case "-r":
			recursive = true
		case "-f":
			force = true
		case "-i":
			interactive = true
		case "-v":
			verbose = true
		default:
			paths = append(paths, arg) // Collecting paths to remove
		}
	}

	for _, path := range paths {
		// Check if the file or directory exists
		if _, err := os.Stat(path); os.IsNotExist(err) {
			if !force {
				color.Red("rm: %s: No such file or directory\n", path)
			}
			continue // Skip to the next path
		}

		if interactive {
			// Prompt for confirmation
			color.Yellow("rm: remove %s? ", path)
			var response string
			fmt.Scanln(&response)
			if strings.ToLower(response) != "y" {
				continue // Skip removing this path
			}
		}

		// Remove directories recursively or regular files
		if recursive {
			if err := os.RemoveAll(path); err != nil {
				color.Red("rm: cannot remove %s: %s\n", path, err)
			} else if verbose {
				color.Green("removed directory: %s\n", path)
			}
		} else {
			if err := os.Remove(path); err != nil {
				color.Red("rm: cannot remove %s: %s\n", path, err)
			} else if verbose {
				color.Green("removed file: %s\n", path)
			}
		}
	}

	return nil
}

// ExecCommand runs a shell command and logs the output
func ExecCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	color.Yellow("Executing: %s %s\n", command, strings.Join(args, " "))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("command failed: %v", err)
	}

	return nil
}

// Run discovers and executes the specified target using reflection
func Run(task interface{}, targetName string) error {
	taskValue := reflect.ValueOf(task)
	taskType := taskValue.Type()

	// Iterate over all methods of the Task struct
	for i := 0; i < taskType.NumMethod(); i++ {
		method := taskType.Method(i)
		if strings.EqualFold(method.Name, targetName) {
			color.Yellow("Running target: %s\n", method.Name)
			result := method.Func.Call([]reflect.Value{taskValue})
			if len(result) > 0 && !result[0].IsNil() {
				err := result[0].Interface().(error)
				color.Red("Error: %v\n", err)
				return err
			}
			color.Green("Target '%s' completed successfully!\n", method.Name)
			return nil
		}
	}

	return fmt.Errorf("target '%s' not found", targetName)
}

// ListTargets lists all available targets using reflection
func ListTargets(task interface{}) {
	taskValue := reflect.ValueOf(task)
	taskType := taskValue.Type()

	color.Cyan("Available targets:\n")
	for i := 0; i < taskType.NumMethod(); i++ {
		method := taskType.Method(i)
		color.Blue("- %s\n", method.Name)
	}
}

// DoIt handles the main logic: listing targets or running a specific target
func DoIt(task interface{}) {
	if len(os.Args) < 2 {
		ListTargets(task)
		return
	}

	target := os.Args[1]
	if err := Run(task, target); err != nil {
		color.Red("Error: %v\n", err)
		os.Exit(1)
	}
}
