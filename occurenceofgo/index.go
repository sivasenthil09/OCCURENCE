package main

import (
    "bufio"
    "fmt"
    "os"
)

func main(){
    var occurence int
    fmt.Print("Enter a sentence : ")
    reader:=bufio.NewReader(os.Stdin)
    input,_:=reader.ReadString('\n')
    size:=len(input)
    for i:=0;i<size;i++{
        if input[i]=='g'||input[i]=='G'{
                if input[i+1]=='o'||input[i+1]=='O'{
                    occurence++;
                }
            } 
        }
    
    fmt.Print(occurence)
}