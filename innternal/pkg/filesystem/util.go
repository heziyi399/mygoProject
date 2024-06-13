package filesystem

import "os"

func isDirExist(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	return err == nil && s.IsDir()
}
