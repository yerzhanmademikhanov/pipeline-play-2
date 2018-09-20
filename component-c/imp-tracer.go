package main


import (
    "fmt"
    "io"
    "os"
    "os/exec"
    "bufio"
    "sync"
)

func main() {
    args := os.Args
    if len(args) < 2 {
        fmt.Println("Need only one argument")
        os.Exit(1)
    }

    script := args[1]
    cmd := exec.Command(script, args[2:]...)

    stdoutPipeReader, stdoutPipeWriter, err := os.Pipe()



	cmd.Stdout = stdoutPipeWriter
	stderrPipeReader, stderrPipeWriter, err := os.Pipe()

	cmd.Stderr = stderrPipeWriter

	// start command
	err = cmd.Start()
	stdoutPipeWriter.Close()
	stderrPipeWriter.Close()
	defer stdoutPipeReader.Close()
	defer stderrPipeReader.Close()
	if err != nil {
		fmt.Println("error")
	}

    var piperGroup sync.WaitGroup
    piperGroup.Add(1)
    go func() {

        defer piperGroup.Done()
        stdout := os.Stdout
		writeLines(stdoutPipeReader, stdout, "stdout")
	}()

    piperGroup.Wait()
}

func writeLines(reader io.Reader, writer io.Writer, channelName string) error {
	lines := bufio.NewScanner(reader)
	for lines.Scan() {
		rawLine := lines.Text()
		// Print the line, unchanged, for the CI product to harvest it as normal, without context information.
		fmt.Fprintln(writer, rawLine)
	}
	return lines.Err()
}
