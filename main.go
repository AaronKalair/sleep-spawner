package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	MAX_LEVEL := 4

	level, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	// We'll have a bunch of processes that immediatly exit at the max level
	if level == MAX_LEVEL {
		return
	}

	// Need the top level to outlive the others, otherwise the container would
	// exit and you wouldn't be able to inspect the process tree
	sleepTime := 0
	if level == 1 {
		sleepTime = 20000000
	} else {
		// Generate proceses where children sleep for longer than there parents
		// so parents exit first without waiting on the children showing
		// what happens to orphan / zombie processes
		sleepTime = level * 1000
	}

	level += 1
	for i := 0; i < 2; i++ {
		// Spawn a command and intentionally dont wait on it
		log.Println("Spawning a process at level", level, "for", sleepTime, "ms")
		err := exec.Command("/srv/sleep-spawner", strconv.Itoa(level)).Start()
		if err != nil {
			panic(err)
		}
	}
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
}
