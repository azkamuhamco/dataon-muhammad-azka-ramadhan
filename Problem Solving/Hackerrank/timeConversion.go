package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "strconv"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
    cekAmPm          := s[8:]
    clockTwoDigit, _ := strconv.Atoi(s[:2])
    clockOneDigit, _ := strconv.Atoi(s[1:2])
    
    if cekAmPm == "PM" {
        if clockTwoDigit == 12 {
            return s[:8]
        } else if clockTwoDigit == 10 || clockTwoDigit == 11 {
            return strconv.Itoa(clockTwoDigit+12) + s[2:8]
        }
        return strconv.Itoa(clockOneDigit+12) + s[2:8]
    }
    
    if clockTwoDigit == 12 {
        return "00" + s[2:8]
    }
    return s[:8]
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := timeConversion(s)

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
