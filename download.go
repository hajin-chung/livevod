package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func downloadHLSFromUri(uri string, dest string, sizeLimit int64, c chan string) {
	cmd := exec.Command("ffmpeg", "-y", "-progress", "pipe:2", "-hide_banner", "-loglevel", "error", "-i", uri, "-movflags", "+frag_keyframe", "-c", "copy", dest)

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("Failed piping command: ", err)
		c <- "ERR"
	}
	defer stderrPipe.Close()
	stderrReader := bufio.NewReader(stderrPipe)

	if err := cmd.Start(); err != nil {
		fmt.Println("Failed starting command: ", err)
		c <- "ERR"
	}

	go handleReader(cmd, stderrReader, c, sizeLimit)

	err = cmd.Wait()
	if err != nil {
		fmt.Println("Failed waiting command: ", err)
		c <- "ERR"
	}

	c <- "ERR"
}

func handleReader(cmd *exec.Cmd, reader *bufio.Reader, c chan string, sizeLimit int64) {
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		c <- str
		kv := strings.Split(str, "=")
		if len(kv) >= 2 {
			key := kv[0]
			val := strings.Trim(kv[1], "\n")

			if key == "total_size" {
				totalSize, err := strconv.ParseInt(val, 10, 0)
				if err != nil {
					fmt.Println(err)
					continue
				}
				fmt.Printf("total_size=%d\n", totalSize)

				if totalSize > sizeLimit {
					fmt.Println("size limit")
					cmd.Process.Kill()
				}
			}
		}
	}
}
