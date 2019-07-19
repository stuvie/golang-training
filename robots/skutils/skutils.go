package skutils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("fatal error: %v\n", err)
		panic(err)
	}
}

func copy(src, dst string) {
	// Read all content of src to data
	data, err := ioutil.ReadFile(src)
	checkErr(err)
	// Write data to dst
	err = ioutil.WriteFile(dst, data, 0644)
	checkErr(err)
}

func copyKeysToTmpFolder() string {
	dir, error := os.Getwd()
	checkErr(error)
	fmt.Printf("current dir %s\n", dir)
	tmpFolder := "/temp/" + strconv.Itoa(os.Getpid()) + "/"
	err := os.Mkdir(tmpFolder, 0755)
	checkErr(err)
	// copy(config.SERVER_PUBLIC_KEYFILE, tmpFolder)

	return tmpFolder
}
