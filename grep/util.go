package grep

// CanonicalizeFilePath expands a file path to its canonical version
// and return true if it is a file and false if it's a directory
func CanonicalizeFilePath(dir string) bool {
	return true
}

// CollectDescendants collects all files in the given directory
func CollectDescendants(dirs []string) []string {
	return []string{}
}
