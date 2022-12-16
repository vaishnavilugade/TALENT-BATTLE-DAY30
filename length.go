package main
import "fmt"
func main(){
    var a string
    fmt.Print("Enter string:")
    fmt.Scan(&a)
    var length=len([]rune(a))
    fmt.Print("number of characters in string : ",length)
}
