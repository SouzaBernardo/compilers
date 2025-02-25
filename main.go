package main

import (
	"compiladores/utils"
	"flag"
	"fmt"
	"regexp"
)

func main() {
	hashmap := make(map[string]string)
	hashmap["🤩"] = "begin"
	hashmap["🥳"] = "end"

	filename := flag.String("filename", "exemplo.emo", "The file to read")
	destination := flag.String("destination", "string.go", "The file to create")

	content, err := utils.ReadFile(filename)
	utils.CreateFileIfNotExists(destination)

	re := regexp.MustCompile(`\s+`)
	splited := re.Split(content, -1)

	for i, parte := range splited {
		fmt.Println(i, parte, hashmap[parte])	
	}


	
	if err != nil { fmt.Println("Erro ao ler o arquivo: %v", err) }
	
	fmt.Println(content)
}
