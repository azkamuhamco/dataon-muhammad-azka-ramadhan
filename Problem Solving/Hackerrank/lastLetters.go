package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)


/*
 * Complete the 'lastLetters' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING word as parameter.
 */

func lastLetters(word string) string {
    var cetak2 string
    rns := []rune(word) // convert to rune
    for i, j := 0, len(rns)-1; i<j; i, j = i+1, j-1 {
        rns[i], rns[j] = rns[j], rns[i]
    }
    cetak  := string(rns)
    cetak2 = cetak[0:1] + " " + cetak[1:2] 
    return cetak2
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    word := readLine(reader)

    result := lastLetters(word)

    fmt.Fprintf(writer, "%s\n", result)

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
