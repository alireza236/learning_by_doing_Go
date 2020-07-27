package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)


var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.59-pipeline-temp")

type FileInfo struct{
   Filepath string // location
   Content []byte // file content
   Sum string // md5 sum content

   IsRenamed bool //indicator whether  the particular file is renamed already or not

}


func readFiles() <-chan FileInfo {
	chanOut := 	make(chan FileInfo)

	go func() {
		err := filepath.Walk(tempPath, func (path string, info os.FileInfo, err error) error  {
			// if theres an error , return immediately
			if err!= nil {
				return err
			}

			//if it is a subdirectory 

			if info.IsDir() {
				return nil
			}

			buf, err := ioutil.ReadFile(path)

			if err != nil {
				return nil
			}

			chanOut <- FileInfo{
				Filepath: path,
				Content: buf,
			}

			return nil
		})

		if err != nil {
			log.Println("ERROR", err.Error())
		}
		close(chanOut)
	}()
	return chanOut
}

func getSum(chanIn <-chan Fileinfo ) <-chan FileInfo  {
	chanOut := make(chan  Filenfo)

	go func() {
		for fileInfo := chanIn {
			fileInfo.Sum = fmt.Sprintf("%x", md5.Sum(fileInfo.Content))
			chanOut <- fileInfo
		}
		close(chanOut)
	}() 
	return chanOut
}

func mergeChanFileInfo(chanInMany ...<-chan FileInfo) <-chan FileInfo  {
	wg := new(sync.WaitGroup)
	chanOut := make(chan FileInfo)

	wg.Add(len(chanInMany))

	for _, eachChan := chanInMany {
		go func(eachChan <- fileIFileInfo) {
			for eachChanData := range eachChan {
				chanOut <- eachChanData
			} 
			wg.Done()

		}(eachChan)
	}

	go func() {
		wg.Wait()
		close(chanOut)
	}() 

	return chanOut
}

func main() {
	log.Println("start")
	start := time.Now()
	
	// pipeline 1 : loop all files and read it

	chanFileContent := readFiles()

	chanFileSum1 := getSum(chanFileContent)
	chanFileSum2 := getSum(chanFileContent)
	chanFileSum3 := getSum(chanFileContent)

	chanFileSum := mergeChanFileInfo(chanFileSum1, chanFileSum2, chanFileSum3)


}