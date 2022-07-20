package main

func main() {
}

func finalValueAfterOperations(operations []string) int {
    var r int
    for _, op := range operations {
        if op == "X++" || op == "++X" {
            r++
        } else if op == "X--" || op == "--X" {
            r--
        }
    }
    return r
}
