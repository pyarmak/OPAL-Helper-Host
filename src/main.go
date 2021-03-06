package main

import (
	"bufio"
	"os"
	"os/exec"
	"fmt"
	"path"
	"os/user"
	"runtime"
	"path/filepath"
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
			content.Player = "lib/mpv/mpv"
			if runtime.GOOS == "windows" {
				content.Player = "lib\\mpv\\mpv.exe"
			}
		}
		cmd := exec.Command(content.Player, content.Url)
		err := cmd.Start()
		if err != nil {
			Send(true, fmt.Sprintf("Command error: %v", err))
			os.Exit(1)
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
			content.Dest = path.Join(homedir, "/Downloads/OPALhelper/videos")
		}
		if _, err := os.Stat(content.Dest); os.IsNotExist(err) {
			os.MkdirAll(content.Dest, os.ModePerm)
		}
		output := path.Join(content.Dest, content.Name)
		FfmpegDownload(content.Url, content.Username, output, "ffmpeg error - %v")
		os.Exit(0)
	case "convert":
		if runtime.GOOS == "windows" {
			ext := filepath.Ext(content.Url)
			filename := content.Url[0:len(content.Url)-len(ext)]
			filename += ".pdf"
			cmd := exec.Command("lib\\converter\\OfficeToPDF.exe", content.Url, filename)
			err := cmd.Start()
			if err != nil {
				Send(true, fmt.Sprintf("Command error: %v", err))
				os.Exit(1)
			}
			err = cmd.Wait()
			os.Remove(content.Url)
			Send(false, filename)
			os.Exit(0)
		} else {
			Send(true, "Converting is only supported on Windows systems")
			os.Exit(1)
		}
	default:
		Send(true, "Received an unknown action type.")
		os.Exit(1)
	}
}

