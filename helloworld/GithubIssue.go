package main

import (
	github2 "helloworld/github"
	"os"
)

func main()  {
	github2.SearchIssues(os.Args[1:])

}
