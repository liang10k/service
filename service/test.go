// package main

// import (
// 	"fmt" 
// 	"os"
// )
// func main () {
// 	file err := os.Open();
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	var output string
// 	for scanner.Scan() {
// 		output += scanner.Text() + " "
// 	}
// 	fmt.Println(output)
// }