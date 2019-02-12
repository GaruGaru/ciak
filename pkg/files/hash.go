package files

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func HashDir(root string) (string, error) {
	sha := sha1.New()

	checksum := sha.Sum(nil)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		checksum = sha.Sum(data)
		return nil
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", checksum), nil
}
