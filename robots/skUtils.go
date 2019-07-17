package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"gitlab.bluecatlabs.net/dns-platform/BSL/test-fixtures/iotmockserver/config"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("fatal error in test %v\n", err)
		panic(err)
	}
}

func copy(fileName, tmpFolder string) {
	src := "ssl/" + fileName
	dst := tmpFolder + fileName
	// Read all content of src to data
	data, err := ioutil.ReadFile(src)
	checkErr(err)
	// Write data to dst
	err = ioutil.WriteFile(dst, data, 0644)
	checkErr(err)
}

func CopyKeysToTmpFolder() string {
	dir, error := os.Getwd()
	checkErr(error)
	fmt.Printf("current dir %s\n", dir)
	tmpFolder := "/temp/" + strconv.Itoa(os.Getpid()) + "/"
	err := os.Mkdir(tmpFolder, 0755)
	checkErr(err)
	copy(config.SERVER_PUBLIC_KEYFILE, tmpFolder)
	copy(config.CLIENT_PUBLIC_KEYFILE, tmpFolder)
	copy(config.CLIENT_PRIVATE_KEYFILE, tmpFolder)
	return tmpFolder
}
