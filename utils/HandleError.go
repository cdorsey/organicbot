package utils

import (
    "os"
    "fmt"
)

func HandleError(err error, fatal bool) {
    if (err != nil) {
        fmt.Print(err.Error())
        if fatal {
            os.Exit(1)
        }
    }
}