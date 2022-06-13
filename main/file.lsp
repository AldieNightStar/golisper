; Function (parse text)
(set parse (def t (do
    (set arr (list))
    (set acc (cell (list)))

    (iterate (str-iterate (concat t " ")) c (do
        (if (or (eq c " ") (eq c "\t")) (do
            (list-add arr (str-join (cell-get acc) ""))
            (cell-set acc (list))
        ) (do
            (list-add (cell-get acc) c)
        ))
    ))

    (return arr)
)))

; Call the parser (parse text)
(print (parse "Hello world and all inside"))

; Iterate over parser's result. Will iterate every word
(iterate (parse "This is the parsed array") word (do
    (print word)
))