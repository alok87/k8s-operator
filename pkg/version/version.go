package version

import "fmt"

// Version for the binary
var Version = "none"

// Format Version
func Format() string {
	return fmt.Sprintf("Version %s\n", Version)
}
