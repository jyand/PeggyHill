package main
import (
        "fmt"
        "os"
        "bufio"
        "regexp"
)

func StringsFromFile(fpath string) []string {
        f, _ := os.Open(fpath)
        defer f.Close()

        var s []string
        sc := bufio.NewScanner(f)
        for sc.Scan() {
                s = append(s, sc.Text())
        }

        return s
}

func CharTuples(seq []string, n uint8) []string {
        s := []string{}
        for j := 0 ; j < len(seq) - 1 ; j++ {
                for i := j + 1 ; i < len(seq) ; i++ {
                        s = append(s, seq[j] + seq[i], seq[i] + seq[j])
                }
        }

        if n > 1 {
                n = n - 1
                return CharTuples(s, n)
        } else {
                return s
        }
}

func MatchTuples(tuples []string, words []string) []string {
        s := []string{}

        for j := 0 ; j < len(tuples) ; j++ {
                re := regexp.MustCompile(tuples[j])
                for i := 0 ; i < len(words) ; i++ {
                        if re.MatchString(words[i]) {
                                s = append(s, tuples[j])
                                break
                        }
                }
        }

        return s
}

func main() {
        a := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
        str := MatchTuples(CharTuples(a, 1), StringsFromFile("corncob.dat"))
        for i := 0 ; i < len(str) ; i++ {
                fmt.Println(str[i])
        }
        fmt.Println(len(str))
}
