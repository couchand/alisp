package main

import "os"
import "fmt"
import "bufio"
import "github.com/couchand/alisp/run"
import "github.com/couchand/alisp/repl"

func main() {

    if len(os.Args) == 1 {
        in := bufio.NewReader(os.Stdin)
        out := bufio.NewWriter(os.Stdout)

        repl.Start(in, out)
    } else {
        infile := os.Args[1]

        f, err := os.Open(infile)
        if err != nil {
            panic(err)
        }

        in := bufio.NewReader(f)
        out := bufio.NewWriter(os.Stdout)

        res := run.Run(in, out)
        fmt.Printf("%v\n", res)
    }
}
