package testconfig

import (
	"os"
	"strings"
)

type PackageInfo int

// os join paths
func (i PackageInfo) JoinPath(paths []string) string {
	for n, s := range paths {
		paths[n] = i.ConvertOsPath(s)
	}
	return strings.Join(paths, string(os.PathSeparator))
}

// convert path strings in a string provided
func (i PackageInfo) ConvertOsPath(s string) string {
	if string(os.PathSeparator) == "\\" {
		return strings.Replace(s, "/", "\\", -1)
	} else if string(os.PathSeparator) == "/" {
		return strings.Replace(s, "\\", "/", -1)
	}
	return s
}

// check if a dir exist
func (i PackageInfo) DirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// Get the location of a package dir
// mostly useful for test cases
func (i PackageInfo) GetPackageRootDir(package_path string) (exist bool, path string) {
	for _, s := range strings.Split(os.Getenv("GOPATH"), string(os.PathListSeparator)) {
		s = i.JoinPath([]string{s, "src"})
		var p = i.JoinPath([]string{s, package_path})
		if found, _ := i.DirExists(p); found == true {
			return true, p
		}
	}
	return false, ""
}
