package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"time"
)

// Because gxt is used to run other commands. Those commands themself
// could have various kinds of parameters. To avoid conflict, gxt can not
// use command parameters, instead it use environment variable.
// Retrieve environment variable(interger type) by name
func getIntConf(name string, default_val int) int {
	str_val, found := os.LookupEnv(name)
	if found {
		int_val, err := strconv.Atoi(str_val)
		if err != nil {
			panic(err)
		} else {
			return int_val
		}
	} else {
		return default_val
	}
}

const COLOR_RESET = "\x1b[0m"
const COLOR_RED = "\x1b[31m"
const COLOR_GREEN = "\x1b[32m"
const COLOR_YELLOW = "\x1b[33m"

// Wrap os.Stdout/Stderr, prefix their output with the command name and highlight
// the characters with different colors.
type ContextStdStream struct {
	fil *os.File
	buf *bytes.Buffer
	cmd string
}

// Factory function for ContextStdStream
func NewContextStdStream(fil *os.File, cmd string) *ContextStdStream {
	s := new(ContextStdStream)
	s.fil = fil
	s.buf = bytes.NewBuffer([]byte(""))
	s.cmd = cmd
	return s
}


func (s *ContextStdStream) Write(p []byte) (n int, err error) {
	if n, err = s.buf.Write(p); err != nil {
		return
	}
	err = s.flushBufIfNeed()
	return
}

// Flush all content to the file
func (s *ContextStdStream) flush() (err error) {
	line := s.buf.String()
	if len(line) > 0 {
		_, err = s.fil.WriteString(s.formatLine(line))
	}
	return nil
}


// Flush buffered characters before the last line breaker(\n) to the file
func (s *ContextStdStream) flushBufIfNeed() (err error) {
	for {
		line, err := s.buf.ReadString('\n')
		if err == io.EOF {
			s.buf.WriteString(line)
			break
		}
		if err != nil {
			return err
		}

		s.fil.WriteString(s.formatLine(line))
	}
	return nil
}

func (s *ContextStdStream) formatLine(line string) string{
	// prefix
	var color string
	if s.fil.Name() == "/dev/stderr" {
		color = COLOR_RED
	} else {
		color = COLOR_GREEN
	}
	return fmt.Sprintf("[%s%s%s] %s", color, s.cmd, COLOR_RESET, line)
}

func main() {
	max_retries := getIntConf("GXT_MAX_RETRY", 999)
	retry_wait := getIntConf("GXT_RETRY_WAIT", 3)

	cmdArgs := os.Args[1:]
	outStream := NewContextStdStream(os.Stdout, cmdArgs[0])
	errStream := NewContextStdStream(os.Stderr, cmdArgs[0])
	for times := 0; times <= max_retries; times++ {
		if times > 0 {
			fmt.Printf(
				"\n[%sgxt%s] [%s] Retrying %s, the %dth time.\n",
				COLOR_YELLOW, COLOR_RESET, time.Now().Format(time.RFC822), cmdArgs[0], times)
		}

		cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
		cmd.Stdin = os.Stdin

		cmd.Stdout = outStream
		cmd.Stderr = errStream

		err := cmd.Run()
		if err != nil {
			fmt.Printf("[%sgxt%s] Error: %s", COLOR_RED, COLOR_RESET, err)

		} else {
			break
		}
		outStream.flush()
		errStream.flush()
		time.Sleep(time.Duration(retry_wait) * time.Second)
	}
}
