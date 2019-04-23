package main

import "fmt"

func myAtoi(str string) int {

    var start int = -1
    var num int = 1
    var pos int = 1
    force := false
    
    // find the index where is started to be a number
    for i := 0; i < len(str); i++ {
        c := int(str[i]) - 48
        
        if (c >= 0 && c <= 9) {
            start = i + 1
            num *= c
            break
        } else {
            
            if (force) {
                return 0
            }
            
            switch str[i] {
                case ' ':
                    continue
                case '-':
                    num = -1
                    pos = -1
                    force = true
                case '+':
                    force = true
                default:
                    return 0
            }
            
        }
    }
    
    if start == -1 {
        return 0
    }
    
    // parse number
    for i := start; i < len(str); i++ {
        c := int(str[i]) - 48
        if (c >= 0 && c <= 9) {
            num = num*10 + c*pos
            if (pos > 0) {
                if num > 2147483647 {
                    return 2147483647
                }
            } else {
                if num < -2147483648 {
                    return -2147483648
                }
            }
        } else {
            break
        }
    }
    
    return num
}

func main() {
    input := "42"
    fmt.Println(myAtoi(input))
}