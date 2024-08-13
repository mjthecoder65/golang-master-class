package files

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
os provides platform independent interface to operating system functionalities
such as reading or writing to the file system. Here are examples of important
methods and functionalities.
// Design is unix like but the interface is Golike.
 1. os.Getwd()
 2. os.Chdir()
 3. os.ReadDir()
 4. os.Mkdir()
 5. os.Remove()
 5. os.RemoveAll()

// Working with files

 1. os.Create(name string) (os.file, error)
    Creates of truncates the named file
    file, err := os.Create("/path/to/file")
    if err != nil {
    fmt.Println("Failed to ")
    }
    defer file.Close()

 2. os.Open(name string) (os.file, error)
    Opens a named file for reading
    file, err := os.Open("/path/to/file")
    defer file.Close()

    3.os.OpenFile(name string, flag int, perm os.FileMode) (os.File, error)

    file, err := os.OpenFile("/path/to/file", os.O_RDWR | os.O_CREATE, 0755)

    defer file.Close()

 4. os.Rename("/path/to/old_name", "path/to/new_name")

 5. os.Stat("/path/to/file")
    info, err := os.Stat("/path/to/file")
    fmt.Println(info.Name(), info.Size(), info.Mode())

Working With Environment Variables:
os.Getenv(key string) string
os.Setenv(key, value string) error
os.Unsetenv(key string) error

Process Management.
os.Getpid() int
os.Getppid() int
os.Exit(code int) // Terminate the process with the specified exit code.

Path Manipulations.
*/
func ReadFileLines(filename string) {
	// This method if memory efficient for large files, as it process the file line by line.
	file, err := os.Open("example.txt")

	if err != nil {
		log.Fatalf("failed to open files: %s\n", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		log.Fatalf("failed to scan file: %v", err)
	}

}

func LearnAboutFiles() {
	// Reading file is quite easy and straight forward.
	data, err := os.ReadFile("files.go")

	if err != nil {
		log.Fatalf("failed to read file: %s\n", err)
	}

	fmt.Println(string(data))
}
