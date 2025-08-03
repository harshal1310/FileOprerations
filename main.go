package main

import (
	listFiles "TrueCallerAssignment/ListFiles"
	worker "TrueCallerAssignment/Worker"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Expect go run main .go <dir name>")
	}
	dirName := os.Args[1]
	files, err := listFiles.ListFile(dirName)
	//fmt.Println(files)
	if err != nil {
		log.Fatal("Error while listing the files : ", err)
	}
	workerPool := 5
	worker.StartWorker(files, workerPool)
}
