package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func yeet(err error) {
	if err != nil {
		panic(err)
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func figlet(text string) {
	out, err := exec.Command("figlet", "-t", text).Output()
	yeet(err)
	fmt.Println(string(out))
}

// TODO: convert to ms -> 00:11.53
func toMs(timestamp string) int {
	arr := strings.Split(timestamp, ":")
	mins, err := strconv.Atoi(arr[0])
	yeet(err)
	sec, err := strconv.ParseFloat(arr[1], 32)
	yeet(err)
	return mins*int(time.Minute) + int(math.Round(sec*float64(time.Second)))
}

func sleep(cur string, next string) {
	sleepTime := toMs(next) - toMs(cur)
	time.Sleep(time.Duration(sleepTime))
}

func main() {
	dat, err := os.ReadFile("./lyrics.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.Trim(string(dat), "\n"), "\n")
	exec.Command("ffplay", "./song.m4a", "-nodisp", "-autoexit").Start()
	for idx, line := range lines {
		// TODO: parse [time]lyric
		curLine := line
		nextLine := lines[idx+1]
		curOpenBracketIdx := strings.Index(curLine, "[")
		curCloseBracketIdx := strings.Index(curLine, "]")
		nextOpenBracketIdx := strings.Index(nextLine, "[")
		nextCloseBracketIdx := strings.Index(nextLine, "]")
		curTime := curLine[curOpenBracketIdx+1 : curCloseBracketIdx]
		curLyric := curLine[curCloseBracketIdx+1:]
		nextTime := nextLine[nextOpenBracketIdx+1 : nextCloseBracketIdx]
		clear()
		figlet(curLyric)
		sleep(curTime, nextTime)
	}
}
