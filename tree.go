/*
Alex Shelton
tree file diplay console tool

recursively lists directory files and folders 

*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
)





func recursivePrint(files []os.FileInfo, level int, dirname string) {

	fmt.Println("Entered recursive function: dirname: ", dirname)

	for _, f := range files {

		for indent := 0; indent < level; indent++ {
			fmt.Println("\t")
		}

		if(f.IsDir() == false) {
			fmt.Println(f.Name())

		}else if f.IsDir() == true && f.Name() != ".git"{
			fmt.Println("[" + f.Name() + "]")

			dirname, err := os.Getwd()
			if err != nil {log.Fatal(err)}
			dirname = dirname + "/" + f.Name() 
			fmt.Println("Checking subdirectory: ", dirname)


			folder, err := ioutil.ReadDir(dirname)
			if err != nil {log.Fatal(err)}
			recursivePrint(folder, level + 1, dirname)
		}
	}
}




func main() {

	dirname, err := os.Getwd()
	if err != nil {
			log.Println(err)
	}


	folder, err := ioutil.ReadDir(dirname)
	if err != nil {log.Fatal(err)}

	// fmt.Println(len(folder))
	
	fmt.Println("Checking files from: ", dirname)
	recursivePrint(folder, 0, dirname)



}
