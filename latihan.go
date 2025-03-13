package main
import "fmt"

func palindrom(i int, hasil *int, o int) bool {
    if i == 0 {
        return *hasil == o
    }
    n := i % 10
    *hasil = (*hasil * 10) + n
    return palindrom(i/10, hasil, o)
}

func main() {
    var masukan int
    fmt.Print("Masukkan bilangan: ")
    fmt.Scan(&masukan)
    
    var hasil int
    if palindrom(masukan, &hasil, masukan) {
        fmt.Println("Bilangan tersebut adalah palindrom.")
    } else {
        fmt.Println("Bilangan tersebut bukan palindrom.")
    }
}