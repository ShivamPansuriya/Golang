package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	createFile("new.txt")

	openFile("new.txt")

	openFileWithOptions("new.txt")

	fileInfo("new.txt")

	renameFile("new.txt", "new_name.txt")

	deleteFile("new_name.txt")

	writeFile("new.txt")

	bfioWriteFile("new.txt")

	readFile("new1.txt")

}

func createFile(name string) {
	newFile, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("file created: ", newFile.Name())
	}

	// TRUNCATING A FILE
	err = os.Truncate("new.txt", 0) //0 means completely empty the file.

	// error handling
	if err != nil {
		log.Fatal(err)
	}

	// CLOSING THE FILE
	newFile.Close()
}

func openFile(name string) {
	file, err := os.Open(name) // open in read-only mode
	if err != nil {
		fmt.Println(err)
	}
	file.Close()
}

func openFileWithOptions(name string) {
	file, err := os.OpenFile(name, os.O_APPEND, 0644)
	// We can Use opening attributes individually or combined
	// using an OR between them
	// e.g. os.O_CREATE|os.O_APPEND
	// or os.O_CREATE|os.O_TRUNC|os.O_WRONLY
	// os.O_RDONLY // Read only
	// os.O_WRONLY // Write only
	// os.O_RDWR // Read and write
	// os.O_APPEND // Append to end of file
	// os.O_CREATE // Create is none exist
	// os.O_TRUNC // Truncate file when opening

	if err != nil {
		fmt.Println(err)
	}

	file.Close()
}

func fileInfo(filename string) {
	fileInfo, err := os.Stat(filename)
	p := fmt.Println
	p("File Name:", fileInfo.Name())        // => File Name: a.txt
	p("Size in bytes:", fileInfo.Size())    // => Size in bytes: 0
	p("Last modified:", fileInfo.ModTime()) // => Last modified: 2019-10-21 16:16:00.325037748 +0300 EEST
	p("Is Directory? ", fileInfo.IsDir())   // => Is Directory?  false
	p("Pemissions:", fileInfo.Mode())       // => Pemissions: -rw-r-----

	if err != nil {
		fmt.Println(err)
	}

	// CHECKING IF FILE EXISTS
	fileInfo, err = os.Stat("b.txt")
	// error handling
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("The file does not exist")
		}
	}

}

func renameFile(oldName, newName string) {
	err := os.Rename(oldName, newName)
	// error handling
	if err != nil {
		log.Fatal(err)
	}
}

func deleteFile(name string) {
	err := os.Remove(name)
	// error handling
	if err != nil {
		log.Fatal(err)
	}
}

func writeFile(name string) {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := []byte("Hello World")
	_, err = file.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

func bfioWriteFile(name string) {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	bufferwriter := bufio.NewWriter(file)
	bufferwriter.WriteString("\nHello World from anonymus")

	log.Printf("available buffer size %d \n", bufferwriter.Available())
	log.Printf("available filled %d \n", bufferwriter.Buffered())

	bufferwriter.Flush()
}

func readFile(name string) {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

	fmt.Println(len(data)) // 113

	fmt.Printf("data type %T\n", data) // data type []uint8

}
