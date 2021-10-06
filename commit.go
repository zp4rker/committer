package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var finalCommitMsg *string

func main() {
	amount := flag.Int("amount", 10000000, "the amount of commits to go up to")
	finalCommitMsg = flag.String("final-commit", "default", "the message for the final commit")
	flag.Parse()

	if *finalCommitMsg == "default" {
		msg := fmt.Sprintf("We did it! %v commits! 🎉", *amount)
		finalCommitMsg = &msg
	}

	i := 1

	count := 0
	start := time.Now()

	if _, err := os.Stat(".git/"); os.IsNotExist(err) {
		// init repo
		fmt.Print("Initialising git repo...")
		if err := exec.Command("git", "init").Run(); err != nil {
			panic("Failed to execute git init!")
		}
	}

	// check if started before
	started := false
	fmt.Print("\033[2K\rChecking if already started...")
	if bytes, err := exec.Command("git", "log", "-1", "--pretty=%B").Output(); err == nil {
		str := strings.TrimSpace(string(bytes))
		re := regexp.MustCompile(`Commit (\d+) of \d+`)

		if re.MatchString(str) {
			started = true
			match := re.FindStringSubmatch(str)[1]
			i, _ = strconv.Atoi(match)
			i++
			fmt.Printf("\033[2K\rResuming from commit #%v", i)
		}
	}
	// check current commit count
	if !started {
		fmt.Print("\033[2K\rCounting commits...")
		cmd := exec.Command("git", "rev-list", "--count", "HEAD")
		if bytes, err := cmd.Output(); err == nil {
			if parsed, err := strconv.Atoi(strings.TrimSpace(string(bytes))); err == nil {
				i = parsed
				fmt.Printf("\033[2K\rCounted %v commits!", i)
				i++
			}
		}
	}

	// check if already complete
	if i > *amount - 1 {
		finalCommit()
		return
	}

	// add files
	fmt.Print("\033[2K\rStaging all files in case of any changes...")
	if err := exec.Command("git", "add", "-A").Run(); err != nil {
		panic("Failed to execute git add!")
	}

	for; i < *amount; i++ {
		if count == 0 {
			fmt.Print("\033[2K\rStarting commits now...")
		}

		msg := fmt.Sprintf("Commit %v of %v", i, *amount)
		if err := exec.Command("git", "commit", "--allow-empty", "-m", msg).Run(); err != nil {
			fmt.Print("\033[2K\rEncountered an error, waiting 5 seconds...")
			time.Sleep(5 * time.Second) // wait 5 seconds
			i-- // repeat this commit
			continue
		}
		count++
		percentage := i * 100 / *amount
		now := time.Now()
		secondsElapsed := now.Sub(start).Seconds()

		fmt.Printf("\033[2K\r%v%% (%v)", percentage, i)

		if secondsElapsed > 0 {
			commitsPerSecond := int(float64(count) / secondsElapsed)
			fmt.Printf(" %v commit/s", commitsPerSecond)
		}
	}

	if i == *amount {
		finalCommit()
	}
}

func finalCommit() {
	if bytes, err := exec.Command("git", "log", "-1", "--pretty=%B").Output(); err == nil {
		str := strings.TrimSpace(string(bytes))

		if str != *finalCommitMsg {
			if err := exec.Command("git", "add", "-A").Run(); err != nil {
				panic("Failed to execute git add!")
			}
			msg := *finalCommitMsg
			if err := exec.Command("git", "commit", "--allow-empty", "-m", msg).Run(); err != nil {
				panic("Failed to execute git commit!")
			}

			fmt.Println("\033[2K\r" + msg)
		} else {
			fmt.Println("\033[2K\r" + "We already reached our goal! 😁")
		}
	}
}