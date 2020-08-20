
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



func recursivePrint(files []os.FileInfo, level int, dirname string) {
	for _, f := range files {

		//Outputting file logic
		if f.IsDir() == false {
			output(f.Name(), true, level)

		}else if f.IsDir() == true && f.Name() != ".git"{
			//Sub directory logic
			output(f.Name(), false, level)

			path := dirname + "/" + f.Name()
			folder, err := ioutil.ReadDir(path)
			if err != nil {log.Fatal(err)}

			recursivePrint(folder, level + 1, path)
		}
	}
}




func main() {





	// args := os.Args[1:]

	// if len(args) == 0 {
	// 	fmt.Println("Need to specify path. use '.' for current directory")
	// 	return
	// }


	// dirname, err := os.Getwd()
	// if err != nil {
	// 		log.Println(err)
	// }

	onlyDirectories := flag.Bool("d", false, "Listing Directories only" )
	flag.Parse()


	// folder, err := ioutil.ReadDir(dirname)
	// if err != nil {log.Fatal(err)}

	
	// fmt.Println("[" + dirname + "]")
	// recursivePrint(folder, 0, dirname)

	fmt.Println(*onlyDirectories)
}
