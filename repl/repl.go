package repl

import (
    "bufio"
    "fmt"
    "strings"

    "github.com/couchand/alisp/tree"
    "github.com/couchand/alisp/lexer"
    "github.com/couchand/alisp/parser"
    "github.com/couchand/alisp/eval"
)

func Start(input *bufio.Reader, output *bufio.Writer) {
    log := func(format string, params ...interface{}) {
        fmt.Fprintf(output, format + "\n", params...)
        output.Flush()
    }

    prompt := func(p string) {
        fmt.Fprintf(output, p)
        output.Flush()
    }

    run := func(t tree.SyntaxTree) {
        defer func() {
            if r := recover(); r != nil {
                log("Eval Error: %s", r)
            }
        }()

        log("%v", eval.Eval(t))
    }

    parse := func(l string) tree.SyntaxTree {
        var parsed tree.SyntaxTree

        defer func() {
            if r := recover(); r != nil {
                log("Parse Error: %s", r)
            } else {
                //log("Parsed: %v", parsed)

                run(parsed)
            }
        }()

        lex := lexer.MakeLexer(l)
        parsed = parser.Parse(lex)
        return parsed
    }

    log("Welcome to the alisp read, evaluate, print loop!")
    log("  (ctrl-D to exit)")

    for {
        log("")
        prompt("$ ")

        read, err := input.ReadString('\n')
        if err != nil {
            log("")
            break
        }

        line := strings.Trim(read[0:len(read)-1], " ")

        parse(line)
    }
}
