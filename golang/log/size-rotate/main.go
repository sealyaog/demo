package main

import "time"
import log "github.com/cihub/seelog"

func main() {
    logger, err := log.LoggerFromConfigAsFile("seelog.xml")
    if err != nil {
        panic(err)
    }
    log.ReplaceLogger(logger);
    defer logger.Flush()
    //log.Errorf("Hello I am error)
    for j := 1; j <= 10; j++ { 
        log.Debugf("Hello I am debug %d", j)
        log.Errorf("Hello I am debug %d", j)
        logger.Flush()
        time.Sleep(1000000000)
    }
}
