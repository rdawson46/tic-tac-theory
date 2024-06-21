package utils

import "math"

func max(x[] int) int {
    ret := math.MinInt

    for _, v := range x {
        if v > ret {
            ret = v
        }
    }

    return ret
}

func Check(x [9]rune, idx ...int) (bool, rune) {
    if len(idx) == 0 || max(idx) > 8 {
        return false, ' '
    }

    val := x[idx[0]]

    for _, i := range idx {
        if x[i] != val || x[i] == ' ' {
            return false, ' '
        }
    }
    return true, val
}

func ForEach[K any, B any](x []K, f func(K) B) []B {
    ret := make([]B, len(x))

    for i, a := range x {
        ret[i] = f(a)
    }

    return ret
}
