package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "sort"
)

func sort_sli(sli *[]int, c *chan[]int){
    fmt.Print(*sli)
    fmt.Print(" is sorted to ")
    sort.Ints(*sli)
    fmt.Println(*sli)
    *c <- *sli
}

func main() {
    // read data
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("Please enter a sequence of integers separated by spaces.\n>")
    scanner.Scan()
    line := scanner.Text()
    tokens := strings.Split(line, " ")

    sli := make([]int, 0, 3)
    for _, v := range tokens {

        number, err := strconv.Atoi(v)
        if err != nil {
            fmt.Println("Please enter a sequence of integers only!")
            return
        }
        sli = append(sli, number)
    }
    // split data
    p := len(sli) % 4
    lens := []int{len(sli) / 4, len(sli) / 4, len(sli) / 4, len(sli) / 4}
    for i := 0; i < p; i++ {
        lens[i] += 1
    }
    sli1 := sli[0:lens[0]]
    sli2 := sli[lens[0]:lens[0] + lens[1]]
    sli3 := sli[lens[0] + lens[1]:lens[0] + lens[1] + lens[2]]
    sli4 := sli[lens[0] + lens[1] + lens[2]:len(sli)]
    fmt.Print("Sub array 1: ")
    fmt.Println(sli1)
    fmt.Print("Sub array 2: ")
    fmt.Println(sli2)
    fmt.Print("Sub array 3: ")
    fmt.Println(sli3)
    fmt.Print("Sub array 4: ")
    fmt.Println(sli4)
    // sort data
    c := make(chan []int, 4)
    go sort_sli(&sli1, &c)
    go sort_sli(&sli2, &c)
    go sort_sli(&sli3, &c)
    go sort_sli(&sli4, &c)
    sorted1 := <-c
    sorted2 := <-c
    sorted3 := <-c
    sorted4 := <-c
    // merge
    merged := make([]int, 0, 10)
    merged = append(merged, sorted1...)
    merged = append(merged, sorted2...)
    merged = append(merged, sorted3...)
    merged = append(merged, sorted4...)
    sort.Ints(merged)
    fmt.Print("Final result is: ")
    fmt.Println(merged)
}
