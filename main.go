package main

import "os"
import "bufio"
import "github.com/couchand/alisp/repl"

func main() {
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)

    repl.Start(in, out)
}
