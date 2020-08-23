
/*
Alex Shelton
tree file diplay console tool

recursively lists directory files and folders 

*/
package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"os"
	"log"
)



//retrieves all files and folders in dir, currently sorted ABC
//sorting files first then folders
//in for loop if index is not a file, checks list from rear for file to switch to
func returnSortedDir(path string) []os.FileInfo {
	//gathering file contents
	files, err := ioutil.ReadDir(path)
	if err != nil {log.Fatal(err)}

	for x := 0; x < len(files)-1; x++ {
		//if its a directory move to back
		if files[x].IsDir() {
			for i := len(files)-1; i > 0; i-- {
				if !files[i].IsDir() {
					files[i], files[x] = files[x], files[i]
					break
				}
			}
		}
	}
	return files
}


func isIn(list [3]string, value string)bool{
	for _,item := range list {
		if item == value {
			return true
		}
	}
	return false
}



//Displaying file || folder logic
func output(msg string, isFile bool, indent int) {
  fmt.Print(" |")
  for x:=0; x < indent; x++ {
    fmt.Print("\t")
	}
	outputString := ""
	//initial file
	if indent < 1 {
		outputString += "-"
	}else{
		outputString += " ╵-"
	}
	if isFile {
		fmt.Println(outputString,msg)
	}else{
		fmt.Println(outputString + " [" + msg + "]")
	}
}


func recursivePrint(files []os.FileInfo, blackList [3]string, level int, dirname string, dirOnly bool, nFiles *int, nFolders *int) {
	//files loop is in ABC order not files first
	for _, f := range files {
		//Outputting file logic
		if f.IsDir() == false && !dirOnly {
			output(f.Name(), true, level)
			*nFiles += 1
		}else if f.IsDir() == true && !isIn(blackList, f.Name()){
			//Sub directory logic
			output(f.Name(), false, level)
			*nFolders += 1
			//enter next folder path recursively
			path := dirname + "/" + f.Name()
			folder := returnSortedDir(path)
			recursivePrint(folder, blackList, level + 1, path, dirOnly, nFiles, nFolders)
		}
	}
}




func main() {
	
	dirBlackList := [3]string{".git", "node_modules", "__pycache__"}
	//FLAGS
	onlyDirectories := flag.Bool("d", false, "Listing Directories only" )
	pathToSearch := flag.String("p", ".", "Directory to start search from")

	flag.Parse() 
	
	//Building directory path
	//if '.' use current directory
	if *pathToSearch == "." {
		path , err := os.Getwd()
		if err != nil {log.Fatal(err)}
		*pathToSearch = path
	}

	folder := returnSortedDir(*pathToSearch)
	
	//Initializing folder and file counter
	nFiles := 0
	nFolders := 0

	//Recursively pringing directories
	fmt.Println("[" + *pathToSearch + "]")
	recursivePrint(folder, dirBlackList, 0, *pathToSearch, *onlyDirectories, &nFiles, &nFolders)
	fmt.Println("Number of directories: ", nFolders, ", Number of files: " , nFiles)

}

