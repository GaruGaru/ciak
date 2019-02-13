package files

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func HashFile(file string) (string, error) {
	sha := sha1.New()

	data, err := ioutil.ReadFile(file)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", sha.Sum(data)), nil
}

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
