package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



/*
 * Complete the 'getMinLength' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER k
 */

func getMinLength(a []int64, k int64) int64 {
    var h int64 = 1
    var t int64 = a[0]
    var l int64 = int64(len(a)-1)
    var i int64 = 1
    
    for ;i<=l; i++ {
        if t*a[i] <= k {
            t *= a[i]
        } else {
            h++
            t = a[i]
        }
    }
    return h
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    aCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var a []int64

    for i := 0; i < int(aCount); i++ {
        aItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        aItem := int64(aItemTemp)
        a = append(a, aItem)
    }

    kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    k := int64(kTemp)

    result := getMinLength(a, k)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
