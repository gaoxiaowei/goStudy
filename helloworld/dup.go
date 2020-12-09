package main

import (
  "bufio"
  "fmt"
  "io/ioutil"
  "os"
  "strings"
)
type gxw int
//func main()  {
//  counts := make(map[string]int)
//  input:=bufio.NewScanner(os.Stdin)
//  for input.Scan(){
//    counts[input.Text()]++
//  }
//  // 注意：忽略input.Err()中可能的错误
//  for line,n := range counts{
//    if n>1 {
//      fmt.Printf("%d\t%s\n",n,line)
//    }
//  }
//}
func countLines(f*os.File,counts map[string]int){
  input := bufio.NewScanner(f)
  for input.Scan(){
    counts[input.Text()]++
  }
}
//func main(){
//  counts:=make(map[string]int)
//  files:=os.Args[1:]
//  if len(files) ==0{
//    countLines(os.Stdin,counts)
//  }else{
//    for _, arg := range  files{
//      f,err:=os.Open(arg)
//      if err !=nil{
//         fmt.Fprintf(os.Stderr,"dup2:%v\n",err)
//         continue
//      }
//      countLines(f,counts)
//      f.Close()
//    }
//  }
//  for line,n := range counts{
//    if n>=1{
//      fmt.Printf("%d\t%s\n",n,line)
//    }
//  }
//}
var array [10]string
func init(){
  for i :=0;i<10;i++ {
    array[i] =fmt.Sprintf("%d",i)
  }
}
func main(){
  //r :=[...]int {99:-1}
  //r[0]=10
  //test := 0.18+22.91+22.88+22.96+23.04+0.05+0.96+1.66+0.92+0.87+1.02+1.62+0.9+0.03
  //test :=0.96+1.66+0.92+0.87+1.02+1.62+0.9
  test :=22.91+22.88+22.96+23.04
  fmt.Println(test)
  counts:=make(map[string]int)
  files:=os.Args[1:]
  for _, filename := range  files{
    data,err:=ioutil.ReadFile(filename)
    if err !=nil{
      fmt.Fprintf(os.Stderr,"dup2:%v\n",err)
      continue
    }

    for _,line :=range  strings.Split(string(data),"\n"){
      counts[line]++
    }
  }
  for line,n := range counts{
    if n>=1{
      fmt.Printf("%d\t%s\n",n,line)
    }
  }
}