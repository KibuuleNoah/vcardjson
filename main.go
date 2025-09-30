package main

import (
	"fmt"
	"os"
	"strings"
	"vcardjson/utils"
)


func main() {
	if len(os.Args) < 1{
    fmt.Println("Usage: go run main.go <path/input.vcf>")
		return
	}

	readPath := os.Args[len(os.Args)-1]

	writePath := strings.Replace(readPath,".","_",1)+".json"

	var contacts []utils.Contact

	if err := utils.ReadVcard(&contacts,readPath); err != nil{
		panic(err)
	}

	if err := utils.WriteContactsToJson(contacts,writePath); err != nil{
		panic(err)
	}
    
	fmt.Println("Contacts saved to ", writePath, "Successfully")

}
