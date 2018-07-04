package main

import (
    "github.com/aellwein/slf4go"
    "github.com/aellwein/slf4go/_example/modules"
)

// doesn't need initialize

// use slf4go everywhere
func main() {
    logger := slf4go.GetLogger("main")
    logger.Debugf("I want %s", "Cycle Import")
    logger.Errorf("please support it, in %02d second!", 1)
    modules.Login()
}
