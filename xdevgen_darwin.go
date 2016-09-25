package main

import (
	"os"
	"fmt"
	"path/filepath"
	"strings"
	"io"
)

func main()  {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	fkDirPath := "/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX.sdk/System/Library/Frameworks"
	fkDir, err := os.Open(fkDirPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fkDirInfo, err := fkDir.Stat()
	if err != nil || !fkDirInfo.IsDir() {
		return
	}
	fkDirSubNames, err := fkDir.Readdirnames(0)
	if err != nil {
		return
	}
	for _, fkDirSubName := range fkDirSubNames {
		strs := strings.Split(fkDirSubName, ".")
		if len(strs) == 2 && strs[1] == "framework" {
			hDirPath := filepath.Join(fkDirPath, fkDirSubName, "Headers")
			hDir, err := os.Open(hDirPath)
			if err != nil {
				continue
			}
			hDirInfo, err := hDir.Stat()
			if err != nil || !hDirInfo.IsDir() {
				continue
			}

			newHDirPath := filepath.Join(wd, strs[0])
			os.Mkdir(newHDirPath, os.ModePerm)

			hDirSubNames, err := hDir.Readdirnames(0)
			if err != nil {
				continue
			}
			for _, hDirSubName := range hDirSubNames {
				strs := strings.Split(hDirSubName, ".")
				if len(strs) == 2 && strs[1] == "h" {
					src, err := os.Open(filepath.Join(hDirPath, hDirSubName))
					if err != nil {
						continue
					}
					dest, err := os.Create(filepath.Join(newHDirPath, hDirSubName))
					if err != nil {
						continue
					}
					io.Copy(dest, src)
					src.Close()
					dest.Close()
				}
			}
		}
	}
}
