// just run already
package run

import (
    "io"
    "fmt"
    "bufio"

    "github.com/couchand/alisp/token"
    "github.com/couchand/alisp/tree"
    "github.com/couchand/alisp/lexer"
    "github.com/couchand/alisp/parser"
    "github.com/couchand/alisp/types"
    "github.com/couchand/alisp/eval"
    "github.com/couchand/alisp/scope"
)

type fileLexer struct {
    *bufio.Scanner
}

func isWhitespace(d byte) bool {
    return d == ' ' || d == '\t' || d == '\n' || d == '\r'
}

func isToken(d byte) bool {
    return d == '(' || d == ')' || d == '\''
}

func split(data []byte, atEOF bool) (advance int, chunk []byte, err error) {
    //fmt.Println("Splitting:", string(data))

    if atEOF {
        panic("EOF with: " + string(data))
    }

    start := 0
    for advance < len(data) {
        //fmt.Println("WS Check [%v]='%v'", advance, string(data[advance:advance+1]))
        if isWhitespace(data[advance]) {
            advance += 1
            start = advance
        } else {
            break
        }
    }

    if advance == len(data) {
        return
    }

    //fmt.Println("TOK Check [%v]='%v'", start, string(data[start:start+1]))
    if isToken(data[start]) {
        advance += 1
        chunk = data[start:advance]
        return
    }

    for advance < len(data) {
        //fmt.Println("WS Check [%v]='%v'", advance, string(data[advance:advance+1]))
        if isToken(data[advance]) || isWhitespace(data[advance]) {
            chunk = data[start:advance]
            return
        }
        advance += 1
    }
    return
}

func (l fileLexer) GetToken() token.Token {
    l.Split(split)
    ok := l.Scan()
    if !ok {
        if err := l.Err(); err != nil {
            panic(err)
        }
        return token.EOF
    }
    str := l.Text()

    if str == "(" {
        return token.PAREN_OPEN
    }
    if str == ")" {
        return token.PAREN_CLOSE
    }
    if str == "'" {
        return token.QUOTE
    }
    return token.Atom(str)
}

func MakeLexer(reader io.Reader) lexer.Lexer {
    return fileLexer{bufio.NewScanner(reader)}
}

func Run(input *bufio.Reader, output *bufio.Writer) (res types.Value) {
    root := scope.RootScope()

    log := func(format string, params ...interface{}) {
        fmt.Fprintf(output, format + "\n", params...)
    }

    run := func(t tree.SyntaxTree) types.Value {
        var evaled types.Value

        defer func() {
            if r := recover(); r != nil {
                log("Eval Error: %s", r)
            } else {
                res = evaled
            }
        }()

        evaled = eval.EvalScope(t, root)
        return evaled
    }

    parse := func(l lexer.Lexer) (tree.SyntaxTree, bool) {
        var parsed tree.SyntaxTree
        err := false

        defer func() {
            if r := recover(); r != nil {
                log("Parse Error: %s", r)
                err = true
            } else {
                run(parsed)
            }
        }()

        parsed = parser.Parse(l)
        return parsed, err
    }

    inputLexer := MakeLexer(input)

    for {
        t, _ := parse(inputLexer)
        eof := t.Text == "" && len(t.Children) == 0
        if eof {
            return
        }
    }

    return
}
