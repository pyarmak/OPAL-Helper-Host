package main

import (
	"bufio"
	"os"
	"os/exec"
	"fmt"
	"path"
	"os/user"
	"runtime"
)

func main() {
	read()
}

func read() {
	for {
		s := bufio.NewReader(os.Stdin)
		length := make([]byte, 4)
		s.Read(length)
		lengthNum := readMessageLength(length)
		content := make([]byte, lengthNum)
		s.Read(content)
		processMessage(content)
	}
}

func processMessage(msg []byte) {
	switch 	content := DecodeMessage(msg); content.Type {
	case "play":
		if content.Player == "DEFAULT" {
			if runtime.GOOS == "windows" {
				content.Player = "mpv\\mpv.exe"
			}
		}
		cmd := exec.Command(content.Player, content.Url)
		err := cmd.Start()
		if err != nil {
			panic(fmt.Sprintf("Error sending message: %v", err))
		}
		Send(false, content.Url)
		os.Exit(0)
	case "download":
		if content.Dest == "DEFAULT" {
			myself, err := user.Current()
			if err != nil {
				Send(true, fmt.Sprintf("failed to get default download location - %v", err))
			}
			homedir := myself.HomeDir
			content.Dest = path.Join(homedir, "/Downloads/OPALhelper")
		}
		if _, err := os.Stat(content.Dest); os.IsNotExist(err) {
			os.MkdirAll(content.Dest, os.ModePerm)
		}
		output := path.Join(content.Dest, content.Name)
		FfmpegDownload(content.Url, output, "ffmpeg error - %v")
		os.Exit(0)
	default:
		Send(true, "Received an unknown action type.")
		os.Exit(1)
	}
}

