package main
import (
        "testing"
)

func BenachmarkMain(b *testing.B) {
        for i := 0 ; i < b.N ; i++ {
                main()
        }
}
