package main
import (
        "fmt"
        "time"
        "math/rand"
)

const SIZE = 10

func InitGame() [SIZE][SIZE]string {
        a := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
        var b [SIZE][SIZE]string
        for j := 0 ; j < SIZE ; j++ {
                for i := 0 ; i < SIZE ; i++ {
                        rand.Seed(time.Now().UnixNano())
                        b[i][j] = a[rand.Intn(26)]
                }
        }
        return b
}


func Permutations(b [SIZE][SIZE]string) []string {
        p := []string{}
        N := SIZE - 1

        for j := 0 ; j < N ; j++ {
                for i := 0 ; i < N ; i++ {
                        p = append(p, b[i][j] + b[i][j+1], b[i][j] + b[i+1][j+1], b[i][j] + b[i+1][j])
                        p = append(p, b[i][j+1] + b[i][j], b[i+1][j+1] + b[i][j], b[i+1][j] + b[i][j])
                        if i > 0 {
                                p = append(p, b[i][j] + b[i-1][j+1])
                                p = append(p, b[i-1][j+1] + b[i][j])
                        }
                }
                p = append(p, b[N][j] + b[N-1][j+1], b[N][j] + b[N][j+1])
                p = append(p, b[N-1][j+1] + b[N][j], b[N][j+1] + b[N][j])
        }
        for i := 0 ; i < N ; i++ {
                p = append(p, b[i][N] + b[i+1][N])
                p = append(p, b[i+1][N] + b[i][N])
        }

        return p
}

func main() {
        b := InitGame()
        for j := 0 ; j < SIZE ; j++ {
                for i := 0 ; i < SIZE ; i++ {
                        fmt.Printf("%s ", b[i][j])
                }
                fmt.Println()
        }
        p := Permutations(b)
        for i := 0 ; i < len(p) ; i++ {
                fmt.Println(p[i])
        }
        fmt.Println("Number of Permutations:")
        fmt.Println(len(p))
}
