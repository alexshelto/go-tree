
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



func getDirName(path string)string {
	if path == "." {
		dirname, err := os.Getwd()
		if err != nil {
				log.Println(err)
		}
		return dirname
	}else{
		return path
	}
}



func output(msg string, isFile bool, indent int) {
  fmt.Print("  |")
  for x:=0; x < indent; x++ {
    fmt.Print("\t")
	}

	outputString := ""
	//initial file
	if indent < 1 {
		outputString += "-"
	}else{
		outputString += "╵-"
	}
	if isFile {
		fmt.Println(outputString,msg)
	}else{
		fmt.Println(outputString + "[" + msg + "]")
	}
}



func recursivePrint(files []os.FileInfo, level int, dirname string, dirOnly bool, nFiles *int, nFolders *int) {
	for _, f := range files {

		//Outputting file logic
		if f.IsDir() == false && !dirOnly {
			output(f.Name(), true, level)
			*nFiles += 1

		}else if f.IsDir() == true && f.Name() != ".git"{
			//Sub directory logic
			output(f.Name(), false, level)
			*nFolders += 1

			path := dirname + "/" + f.Name()
			folder, err := ioutil.ReadDir(path)
			if err != nil {log.Fatal(err)}

			recursivePrint(folder, level + 1, path, dirOnly, nFiles, nFolders)
		}
	}
}




func main() {
	onlyDirectories := flag.Bool("d", false, "Listing Directories only" )

	var dirname string
	dirname = getDirName(".")

	folder, err := ioutil.ReadDir(dirname)
	if err != nil {log.Fatal(err)}
	flag.Parse()



	nFiles := 0
	nFolders := 0
	fmt.Println("[" + dirname + "]")
	recursivePrint(folder, 0, dirname, *onlyDirectories, &nFiles, &nFolders)
	fmt.Println("Number of directories: ", nFolders, ", Number of files: " , nFiles)

}
