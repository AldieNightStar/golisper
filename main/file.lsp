(program "Server"
    ; some sort of comment
    ; dsadsadsadads
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