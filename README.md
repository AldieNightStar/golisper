# Go(Lisp)er

## Golang Lisp syntax parser

# Syntax demo
```lisp
; Comment string
(program "Server"
    (def init (args 'a' 'b' 'c') (body
        (log "Prepare to init")
        (set self.isReady true)
        (self.start)
        (log "Init done")
    ))
    (def start abstract)
    (def stop abstract)
    (def read-log (returns 'string') abstract)
)
```

# Supports tokens
- `String`: `` `str` 'str' "str" ``
- `Number`: `1 -10 144.22 -99.18`
- `Etc`: `hello abc test`
- `Tag`: `(name val1 val2 val3 (subtag a b c))`
- `Comment`: `; this is the comment`

# API
```go
// Import
import "github.com/AldieNightStar/golisper"

// Lets assume there are loadFile func which loads the file as a string
SourceCode, _ := loadFile("file.txt")

// Parse source into atoken list
tokens, err := golisper.Parse(fileContent)

/* 
    Print each tag
    ==============
    Tag:
        Name string
        Values []*Value
    Value:
        Type ValueType
        StringVal string
        NumberVal float64
        TagVal *Tag
*/
for _, tag := range tags {
    fmt.Println(tag)
}
```
* Output
```
TAG[program](
    STR:'Server' 
    TAG[def](
        ETC(init) 
        TAG[args](
            STR:'a' 
            STR:'b' 
            STR:'c' 
        )
        TAG[body](
            TAG[log](
                STR:'Prepare to init' 
            )
            TAG[set](
                ETC(self.isReady) 
                ETC(true) 
            )
            TAG[self.start](
            )
            TAG[log](
                STR:'Init done' 
            )
        )
    )
    TAG[def](
        ETC(start) 
        ETC(abstract) 
    )
    TAG[def](
        ETC(stop) 
        ETC(abstract) 
    )
    TAG[def](
        ETC(read-log) 
        TAG[returns](
            STR:'string' 
        )
        ETC(abstract) 
    )
)
```

