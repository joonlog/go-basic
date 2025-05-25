package main

import (
	"fmt"
	"os"
)

func main() {
	info, err := os.Stat("/etc/passwd")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("파일명:", info.Name())
	fmt.Println("크기:", info.Size(), "bytes")
	fmt.Println("퍼미션:", info.Mode())
	fmt.Println("수정시간:", info.ModTime())
}

