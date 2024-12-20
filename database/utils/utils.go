package utils

import (
	"fmt"
	"log"
	"math/rand"
	"os"
)

func SaveData1(path string, data []byte) error {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}

	defer func(fp *os.File) {
		err := fp.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(fp)

	_, err = fp.Write(data)
	if err != nil {
		return err
	}
	return fp.Sync()
}

// Replacing data atomically by renaming files
func SaveData2(path string, data []byte) error {
	tmp := fmt.Sprintf("%s.tmp.%d", path, rand.Int())
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}
	defer func(fp *os.File) { // 4. discard the temporary file if it still exists
		err := fp.Close() // not expected to fail
		if err == nil {
			return
		}
		err = os.Remove(tmp)
		if err != nil {
			log.Fatal(err)
		}
	}(fp)
	if _, err = fp.Write(data); err != nil { // 1. save to the temporary file
		return err
	}
	if err = fp.Sync(); err != nil { // 2. fsync
		return err
	}
	err = os.Rename(tmp, path) // 3. replace the target
	return err
}
