package grep

import (
	"net/rpc"
)

// Grepper that greps like the UNIX utility
type Grepper interface {
	// Grep a keyword in given directories, output all lines matching the keyword.
	// Include line numbers if specified.
	Grep(keyword string, dirs []string, includeLineNum bool) []string
}

// LocalGrepper that greps in local files.
type LocalGrepper struct {
}

// NewLocalGrepper creates a local Grepper.
func NewLocalGrepper() LocalGrepper {
	return LocalGrepper{}
}

// Grep local files.
func (l LocalGrepper) Grep(keyword string, dirs []string, includeLineNum bool) []string {

	// for nil, dir := range dirs {
	// 	file, err := os.Open(dir)
	// 	scanner :=
	// }

	return []string{""}
}

// DistributedGrepper that greps in both local files and remote files
type DistributedGrepper struct {
	peers []rpc.Client
	// self string
}

// NewDistributedGrepper creates a Distributed Grepper.
func NewDistributedGrepper(id int, idByServerAddr map[int]string) DistributedGrepper {
	return DistributedGrepper{}
}

// Grep both local and remote files.
func (l DistributedGrepper) Grep(keyword string, dirs []string, includeLineNum bool) []string {

	return []string{""}
}
