package task
import (
	"fmt"
)

var Task = "task"

var Version string

var version string
func GetVersion() string {
	return version
}

func init()  {
	Version = "controller version1.11"
	fmt.Println("Version init")
}