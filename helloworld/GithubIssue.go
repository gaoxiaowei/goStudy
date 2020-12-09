package main

import (
	"fmt"
	"helloworld/github"
	"log"
	"os"
)

func main(){
	result, error := github.SearchIssues(os.Args[1:])
	if error !=nil{
		log.Fatal(error)
	}
	fmt.Printf("%d issues:\n",result.TotalCount)
	for _,item :=range result.Items{
		fmt.Printf("#%-5d %0.9s %.55s\n",item.Number,item.User.Login,item.Title)

	}

}
