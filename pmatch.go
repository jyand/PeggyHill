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

// Permutations of length k where k = 2^n... be careful
func ExpTuples(seq []string, n uint8) []string {
        s := []string{}
        for j := 0 ; j < len(seq) - 1 ; j++ {
                for i := j + 1 ; i < len(seq) ; i++ {
                        s = append(s, seq[j] + seq[i], seq[i] + seq[j])
                }
        }
        if n < 2 {
                return s
        }
        return ExpTuples(s, n-1)
}

// Permutations of length k where k = n + 1
func LinTuples(alph []string, seq []string, n uint8) []string {
        s := []string{}
        for j := 0 ; j < len(alph) ; j++ {
                for i := j + 1 ; i < len(seq) ; i++ {
                        s = append(s, alph[j] + seq[i], seq[i] + alph[j])
                }
        }
        if n < 2 {
                return s
        }
        return LinTuples(alph, s, n-1)
}

// match character sequences with word list
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

func Echo(s []string, n uint64) {
        if n == 0 {
                fmt.Println(len(s))
                return
        }
        fmt.Println(s[n-1])
        Echo(s, n-1)
}

func main() {
        a := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
        exp := MatchTuples(ExpTuples(a, 1), StringsFromFile("dict/web2.dat"))
        lin := LinTuples(a, a, 3)
        Echo(lin, uint64(len(lin)))
        Echo(exp, uint64(len(exp)))
}
