package main

import (
    "fmt"
    "strconv"
    "strings"
	"os"
	"os/exec"
	"io"
	"runtime"
)

var duration = 0
var allRes = ""
var lastPer = -1

func durToSec(dur string) (sec int) {
    durAry := strings.Split(dur, ":")
    if len(durAry) != 3 {
        return
    }
    hr, _ := strconv.Atoi(durAry[0])
    sec = hr * (60 * 60)
    min, _ := strconv.Atoi(durAry[1])
    sec += min * (60)
    second, _ := strconv.Atoi(durAry[2])
    sec += second
    return
}
func getRatio(res string) {
    i := strings.Index(res, "Duration")
    if i >= 0 {

        dur := res[i+10:]
        if len(dur) > 8 {
            dur = dur[0:8]

            duration = durToSec(dur)
            //Send(false, fmt.Sprintln("duration:", duration))
            allRes = ""
        }
    }
    if duration == 0 {
        return
    }
    i = strings.Index(res, "time=")
    if i >= 0 {

        time := res[i+5:]
        if len(time) > 8 {
            time = time[0:8]
            sec := durToSec(time)
            per := (sec * 100) / duration
            if lastPer != per {
                lastPer = per
                Send(false, strconv.Itoa(per))
            }

            allRes = ""
        }
    }
}

func FfmpegDownload(url string, output string, errorMsg string) {
	os.Remove(output)
	ffmpeg := "lib/ffmpeg/ffmpeg"
	if runtime.GOOS == "windows" {
		ffmpeg = "lib\\ffmpeg\\ffmpeg.exe"
	}
	cmd := exec.Command(ffmpeg, "-i", url, "-bsf:a", "aac_adtstoasc", "-acodec", "copy", "-vcodec", "copy", output)
	stderr, _ := cmd.StderrPipe()
	cmd.Start()
	oneByte := make([]byte, 8)
	for {
		_, err := stderr.Read(oneByte)
		if err != nil {
			if err == io.EOF {
				Send(false, "DONE")
			} else {
				Send(true, fmt.Sprintf(errorMsg, err))
				os.Exit(1)
			}
			break
		}
		allRes += string(oneByte)
		getRatio(allRes)
	}
}