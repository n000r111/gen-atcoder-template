package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	num := os.Args[1]

	fileInfo, err := os.Lstat("../")
	if err != nil {
		fmt.Println(err)
	}

	fileMode := fileInfo.Mode()
	unixPerms := fileMode & os.ModePerm

	if err := os.Mkdir("../"+num, unixPerms); err != nil {
		fmt.Println(err)
	}

	subDirs := []string{"A", "B", "C", "D", "E"}
	for _, dir := range subDirs {
		problemPath := filepath.Join("..", num, dir)

		if err := os.Mkdir(problemPath, unixPerms); err != nil {
			fmt.Println("Error creating directory:", dir, err)
			return
		}

		createdFile, err := os.Create(filepath.Join(problemPath, "main.py"))
		if err != nil {
			fmt.Println("Error creating main.py:", err)
			return
		}
		defer createdFile.Close()

		originalFilePath := filepath.Join(".", "templates", "main.py")

		templateFile, err := os.Open(originalFilePath)
		if err != nil {
			fmt.Println("Error opening template file:", err)
			return
		}

		_, err = templateFile.WriteTo(createdFile)
		if err != nil {
			fmt.Println("Error writing to main.py:", err)
			return
		}

		srcInfo, err := os.Stat(originalFilePath)
		if err != nil {
			fmt.Println("failed to get source file information: %w", err)
			return
		}
		err = os.Chmod(filepath.Join(problemPath, "main.py"), srcInfo.Mode())
		if err != nil {
			fmt.Println("failed to set file permissions: %w", err)
			return
		}

		if _, err := os.Create(filepath.Join(problemPath, ".tool-versions")); err != nil {
			fmt.Println("Error creating .tools-versions:", err)
			return
		}

		if os.WriteFile(filepath.Join(problemPath, ".tool-versions"), []byte("python 3.12.4\n"), unixPerms); err != nil {
			fmt.Println("Error writing to .tools-versions:", err)
			return
		}

		if err := os.Mkdir(filepath.Join(problemPath, "test"), unixPerms); err != nil {
			fmt.Println("Error creating directory:", dir, err)
			return
		}

		for i := range [3]int{} {
			testFilePath := filepath.Join(problemPath, "test", "test"+strconv.Itoa(i+1)+".txt")
			if _, err := os.Create(testFilePath); err != nil {
				fmt.Println("Error writing to test"+strconv.Itoa(i+1)+".txt:", err)
				return
			}

			if err := os.WriteFile(testFilePath, []byte(strconv.Itoa(i+1)+"\n"), unixPerms); err != nil {
				fmt.Println("Error writing to test"+strconv.Itoa(i+1)+".txt:", err)
				return
			}
		}
	}
}
