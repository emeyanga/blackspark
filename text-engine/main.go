package textengine
/*
BLACKSPARK: PHASE 1 — TEXT ENGINE CORE

Operation Location
All work happens inside the Blackspark repository. No side folders. Everything must stay structured and
clean.
Mission Brief

Build a CLI-based text engine in Go that reads raw text and converts it into structured tokens. No
transformations yet.
Tech Stack

Go language. Standard library only. Terminal execution.
Folder Structure
blackspark/
■■■ text-engine/
■■■ main.go
■■■ input.txt
■■■ output.txt
■■■ internal/
■ ■■■ reader/
■ ■ ■■■ reader.go
■ ■■■ tokenizer/
■ ■ ■■■ tokenizer.go
■ ■■■ scanner/
■ ■ ■■■ scanner.go
■ ■■■ types/
■ ■■■ token.go
■■■ tests/
■ ■■■ cases.txt
■■■ README.md
Execution Command
go run . input.txt output.txt
Sample Input
Hello, world!
'we move in silence'
convert (up) this now...

Expected Output
[0] Hello -> WORD
[1] , -> PUNCT
[2] world -> WORD
[3] ! -> PUNCT
[4] ' -> QUOTE
[5] we move in silence -> TEXT
[6] ' -> QUOTE
[7] convert -> WORD
[8] (up) -> COMMAND
[9] this -> WORD
[10] now -> WORD
[11] . -> PUNCT
[12] . -> PUNCT
[13] . -> PUNCT
Rules

Separate punctuation. Preserve quotes. Keep commands intact. No empty tokens.
Success Criteria

Tokenizer handles all edge cases. Tokens are clean.
*/

package main

import (
	"fmt"
)

func main