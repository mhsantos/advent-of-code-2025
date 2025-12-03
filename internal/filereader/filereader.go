package filereader

import (
	"bufio"
	"log/slog"
	"os"
)

func ReadInput(filename string) []string {
	file, err := os.Open("./input/" + filename)
	if err != nil {
		slog.Error("error opening file", slog.String("filename", filename), slog.Any("error", err))
	}
	defer file.Close()

	var lines []string
	reader := bufio.NewReader(file)
	var lineBuffer, part []byte
	var isPrefix bool

	for {
		if part, isPrefix, err = reader.ReadLine(); err != nil {
			break
		}
		lineBuffer = append(lineBuffer, part...)
		if !isPrefix {
			lines = append(lines, string(lineBuffer[:]))
			lineBuffer = make([]byte, 0)
		}
	}
	return lines
}
