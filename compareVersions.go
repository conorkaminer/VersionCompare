package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	compareVal := compareVersions(os.Args[1], os.Args[2])

	fmt.Println(compareVal)

	os.Exit(compareVal)
}

func compareVersions(a string, b string) int {
	aSplit := strings.Split(a, ".")
	bSplit := strings.Split(b, ".")

	aLen := len(aSplit)
	bLen := len(bSplit)
	length := min(aLen, bLen)

	for i := 0; i < length; i++ {
		intA, _ := strconv.ParseInt(aSplit[i], 10, 32)
		intB, _ := strconv.ParseInt(bSplit[i], 10, 32)
		if intA > intB {
			return 1
		} else if intB > intA {
			return -1
		}
	}
	if aLen > bLen {
		return 1
	} else if bLen > aLen {
		return -1
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
