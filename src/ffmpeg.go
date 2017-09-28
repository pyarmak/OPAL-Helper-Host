package main

import (
    "fmt"
    "strconv"
    "strings"
	"os"
	"os/exec"
	"runtime"
	"time"
	"io/ioutil"
	"io"
	"regexp"
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

func FfmpegDownload(url string, username string, output string, errorMsg string) {
	os.Remove(output)
	ffmpeg := "lib/ffmpeg/ffmpeg"
	if runtime.GOOS == "windows" {
		ffmpeg = "lib\\ffmpeg\\ffmpeg.exe"
	}
	file, err := ioutil.ReadFile("copyright.txt")
	if err != nil {
		Send(true, fmt.Sprintf(errorMsg, err))
		os.Exit(1)
	}
	disclaimer := fmt.Sprintf(string(file), username, time.Now().Year())
	re := regexp.MustCompile(`\r?\n`)
	disclaimer = re.ReplaceAllString(disclaimer, " ")
	watermark := fmt.Sprintf("drawtext='text=%v:expansion=normal:fontfile=lib/roboto/Roboto-Regular.ttf:y=h-line_h-10:x=w-mod(max(t-4.5\\,0)*(w+tw)/255.5\\,(w+tw)):fontcolor=white:fontsize=40:shadowx=2:shadowy=2'", disclaimer)
	args := strings.Fields(fmt.Sprintf("-i %v -vf -bsf:a aac_adtstoasc -acodec copy %v", url, output))
	args = append(args[:3], append([]string{watermark}, args[3:]...)...)
	cmd := exec.Command(ffmpeg, args...)
	//Send(true, strings.Join(cmd.Args, " "))
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