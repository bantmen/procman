package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/shirou/gopsutil/process"
)

const zombieStatus = "Z"

func usage() {
	fmt.Println("procman -- monitor and run your command forever")
	fmt.Println("")

	fmt.Println("Usage:")
	fmt.Println("  ./procman [options] command")
	fmt.Println("")

	fmt.Println("Examples:")
	fmt.Println("  ./procman ls -lh")
	fmt.Println("  ./procman -mem 80 python3 server.py")
	fmt.Println("")

	fmt.Println("Options:")
	flag.PrintDefaults()
}

func main() {
	memPtr := flag.Float64("mem", 100, "Max memory percentage threshold. When exceeded, the process is restarted. Default is no threshold.")

	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		usage()
		os.Exit(0)
	}

	flag.Parse()
	memPercentThreshold := float32(*memPtr)

	fmt.Println("Running process: ", flag.Args())

	cmdName := flag.Args()[0]
	cmdArgs := flag.Args()[1:]

	if memPercentThreshold < 100 {
		fmt.Printf("With memory threshold: %.2f percent\n", memPercentThreshold)
	}

	for {
		// Spawn
		cmd := exec.Command(cmdName, cmdArgs...)
		// We want the process logs to be visible.
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Start()
		if err != nil {
			log.Panicln(err)
		}
		p, err := process.NewProcess(int32(cmd.Process.Pid))
		if err != nil {
			log.Panicln(err)
		}

		// Use log.panic instead of os.exit/log.fatal to ensure deferred functions are called.
		defer cmd.Process.Kill()

		// Inspect
		for {
			// NOTE: Process.Status does not support Windows.
			status, err := p.Status()
			if err != nil {
				log.Panicln(err)
			}
			// Terminated child processes have the "zombie" status.
			if status == zombieStatus {
				err = cmd.Wait()
				if err != nil {
					fmt.Printf("Process stopped with error: %v\n", err)
				} else {
					fmt.Println("Process stopped without error")
				}
				break
			}

			percent, err := p.MemoryPercent()
			if err != nil {
				log.Panicln(err)
			}
			if percent > memPercentThreshold {
				if err := cmd.Process.Kill(); err != nil {
					log.Panicln("failed to kill process: ", err)
				}
				fmt.Printf("Restarting: mem percent is at %f\n", percent)
				break
			}
			time.Sleep(1 * time.Second)
		}
	}
}
