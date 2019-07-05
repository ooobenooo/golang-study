package main

import "fmt"

func main() {
    c := make(chan int, 2)
    // loop(3, c) // 没有使用线程，超出缓存channel容量会抛出异常
     go loop(3, c)
     x := <-c
     fmt.Println(x)
     
     fmt.Println("----")
     c = make(chan int, 2)
     go loop(5, c)
     for v := range c {
         fmt.Println(v)
     }
     
     fmt.Println("----")
     c = make(chan int, 2)
     go loop(5, c)
     for {
         select {
             case z := <-c:
             fmt.Println(z)
             if z >= 4 {
                 return
             }
         }
    }

     fmt.Println("--=--")
}

func loop(l int, c chan int) {
    for i := 0; i < l; i++ {
        c <- i
    }
    close(c)
}