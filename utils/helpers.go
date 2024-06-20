package utils

func All(x [3]rune) bool {
    if len(x) == 0 {
        return true
    }

    first := x[0]

    if first == ' ' {
        return false
    }

    for _, a := range x {
        if first != a {
            return false
        }
    }

    return true
}

func ForEach[K any, B any](x []K, f func(K) B) []B {
    ret := make([]B, len(x))

    for i, a := range x {
        ret[i] = f(a)
    }

    return ret
}
