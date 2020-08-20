
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



func output(msg string, indent int) {
  fmt.Print("+")
  for x:=0; x < indent; x++ {
    fmt.Print("\t")
  }
  fmt.Println("|-",msg)

}



func recursivePrint(files []os.FileInfo, level int, dirname string) {

	fmt.Println("Entered recursive function: dirname: ", dirname)

	for _, f := range files {

		if(f.IsDir() == false) {
			output(f.Name(), level)

		}else if f.IsDir() == true && f.Name() != ".git"{
			output("[" + f.Name() + "]", level)

			// dirname, err := os.Getwd()
			// if err != nil {log.Fatal(err)}
			// dirname = dirname + "/" + f.Name() 
			// fmt.Println("Checking subdirectory: ", dirname)


			// folder, err := ioutil.ReadDir(dirname)
			// if err != nil {log.Fatal(err)}
			// recursivePrint(folder, level + 1, dirname)
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
