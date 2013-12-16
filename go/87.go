package main

import "fmt"
import "math"

func nextPrime(n float64) float64 {
    for i := n; i < 7080; i++ {
        prime := true
        var two float64 = 2
        for j := two; j < i; j++ {
            if int(i) % int(j) == 0 {
                prime = false
            }
        }
        if prime {
            return i
        }
    }
    return 0
}

func main() {
    a := make(map[float64]bool)
    var counter int = 0
    var z float64 = 7072 // square root of 50 million
    var q float64 = 369 // cube root of 50 million
    var r float64 = 85 // quad root of 50 million
    var b float64 = 2 // start counter for square root
    for b < z {
        var c float64 = 2 // start counter for cube root
        for c < q {
            var d float64 = 2 // start counter for quad root
            for d < r {
                d = nextPrime(d)
                if a[math.Pow(b,2)+math.Pow(c,3)+math.Pow(d,4)] != true {
                    if math.Pow(b,2)+math.Pow(c,3)+math.Pow(d,4) < 50000000 {
                        fmt.Println(math.Pow(b,2)+math.Pow(c,3)+math.Pow(d,4), b, c, d)
                        counter += 1
                        a[math.Pow(b,2)+math.Pow(c,3)+math.Pow(d,4)] = true
                    }
                }
                d += 1
            }
            c = nextPrime(c+1)
        }
        b = nextPrime(b+1)
    }
    fmt.Println(counter)
}
