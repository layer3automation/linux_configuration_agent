package utils

import (
	"fmt"
	"log"
	"os/exec"
)

func concat(args ...string) string {
	result := ""
	for _, s := range args {
		result = result + " " + s
	}
	return result
}

// ExecuteCommand runs a command and return the error or nil.
func ExecuteCommand(name string, args ...string) error {
	log.Printf("executing [%s %s]", name, concat(args...))
	_, err := exec.Command(name, args...).Output()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	fmt.Printf("command %s successfully executed", name)
	return nil
}
