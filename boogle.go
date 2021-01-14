package main
import (
        "fmt"
        "time"
        "math/rand"
)

const SIZE = 4

func InitGame() [4][4]string {
        a := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
        var b [4][4]string
        for j := 0 ; j < 4 ; j++ {
                for i := 0 ; i < 4 ; i++ {
                        rand.Seed(time.Now().UnixNano())
                        b[i][j] = a[rand.Intn(26)]
                }
        }
        return b
}

func Permutations([][]string) []string {
}

func main() {
        game := InitGame()
        for j := 0 ; j < SIZE ; j++ {
                for i := 0 ; i < SIZE ; i++ {
                        fmt.Printf("%s ", game[i][j])
                }
                fmt.Println()
        }
}

