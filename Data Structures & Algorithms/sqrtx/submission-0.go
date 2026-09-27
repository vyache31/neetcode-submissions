func mySqrt(x int) int {
    left := 0
    right := x
    for left <= right {
        middle := (left + right) / 2
        if middle * middle > x {
            right = middle-1
        } else if middle * middle < x {
            left = middle+1
        } else {
            return middle
        }
    }
    return right
}