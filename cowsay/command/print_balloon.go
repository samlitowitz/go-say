package command

import (
	"fmt"
	"github.com/samlitowitz/go-say/cowsay"
	"io"
	"strings"
)

func PrintBalloon(w io.Writer, opts *cowsay.Options, message string) error {
	processedMessages := make([]string, 0, 1)
	switch opts.NewLines {
	case true:
		processedMessages = append(processedMessages, strings.Split(strings.ReplaceAll(message, "\t", "        "), "\n")...)
	case false:
		processedMessages = append(processedMessages, strings.Split(strings.TrimSpace(formatMessage(message, opts.WrapWidth)), "\n")...)
	}

	maxMessageLength := -1
	for _, processedMessage := range processedMessages {
		if len(processedMessage) > maxMessageLength {
			maxMessageLength = len(processedMessage)
		}
	}

	err := writeBalloonTop(w, maxMessageLength)
	if err != nil {
		return err
	}

	processedMessageFormat := fmt.Sprintf("%%s %%-%ds %%s\n", maxMessageLength)

	err = writeMessage(w, processedMessageFormat, processedMessages, maxMessageLength)
	if err != nil {
		return err
	}

	err = writeBalloonBottom(w, maxMessageLength)
	if err != nil {
		return err
	}

	return nil
}

func formatMessage(message string, wrapWidth int) string {
	buf := []byte(message)
	for i := len(message) - 1; i > 0; i-- {
		// replace new lines with tabs
		if buf[i] == '\n' {
			buf[i] = ' '
		}
		// replace all sequences of multiple spaces with a single space
		if buf[i] == ' ' && buf[i-1] == ' ' {
			buf = append(buf[:i-1], buf[i:]...)
			continue
		}
		// replace multiple tabs with a single tab
		if buf[i] == '	' && buf[i-1] == '	' {
			buf = append(buf[:i-1], buf[i:]...)
			continue
		}
	}
	max := len(buf)
	start := 0
	lastSpaceAt := -1
	for i := 0; i < max-1; i++ {
		if buf[i] == '	' {
			buf[i] = ' '
		}

		switch buf[i] {
		case ' ': // space
			lastSpaceAt = i
		case '\n': // new line
			start = i + 1
			lastSpaceAt = -1
			continue
		}

		// not time to wrap yet
		if i-start < wrapWidth-1 {
			continue
		}
		// wrapping on a space, easy
		if lastSpaceAt == i {
			// replace with new line
			buf[i] = '\n'
			// reset
			start = i + 1
			lastSpaceAt = -1
			continue
		}
		// wrapping in the middle of a string
		// last space is real, break there
		if lastSpaceAt > start {
			// replace with new line
			buf[lastSpaceAt] = '\n'
			// reset
			start = lastSpaceAt + 1
			lastSpaceAt = -1
			continue
		}
		// no space, break right here
		// inject a new line
		buf = append(buf[:i], append([]byte("\n"), buf[i:]...)...)
		// increment max and i to accommodate
		max++
		i++
		// reset
		start = i + 1
		lastSpaceAt = -1
	}

	return string(buf)
}

func writeBalloonTop(w io.Writer, maxMessageLength int) error {
	buf := make([]byte, maxMessageLength+2)
	for i := 0; i < maxMessageLength+2; i++ {
		buf[i] = '_'
	}

	_, err := w.Write([]byte(" "))
	if err != nil {
		return err
	}

	_, err = w.Write(buf)
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(" \n"))
	if err != nil {
		return err
	}
	return nil
}

func writeMessage(w io.Writer, format string, processedMessages []string, maxMessageLength int) error {
	processedMessages = trimAndStripEmptyMessages(processedMessages)
	if len(processedMessages) < 2 {
		return writeShortMessage(w, format, processedMessages, maxMessageLength)
	}
	return writeLongMessage(w, format, processedMessages, maxMessageLength)
}

func trimAndStripEmptyMessages(processedMessages []string) []string {
	toRemove := make([]int, 0, len(processedMessages)/2)
	for i, processedMessage := range processedMessages {
		processedMessages[i] = strings.TrimSpace(processedMessage)
		if len(processedMessages[i]) == 0 {
			toRemove = append(toRemove, i)
		}
	}
	for i := len(toRemove) - 1; i >= 0; i-- {
		j := toRemove[i]
		processedMessages = append(processedMessages[:j], processedMessages[j+1:]...)
	}
	return processedMessages
}

func writeShortMessage(w io.Writer, format string, processedMessages []string, maxMessageLength int) error {
	left, right := "<", ">"

	_, err := w.Write([]byte(fmt.Sprintf(format, left, processedMessages[0], right)))

	return err
}

func writeLongMessage(w io.Writer, format string, processedMessages []string, maxMessageLength int) error {
	upLeft, upRight, downLeft, downRight, left, right := "/", "\\", "\\", "/", "|", "|"

	// first line
	_, err := w.Write([]byte(fmt.Sprintf(format, upLeft, processedMessages[0], upRight)))
	if err != nil {
		return nil
	}

	for i := 1; i <= len(processedMessages)-2; i++ {
		_, err = w.Write([]byte(fmt.Sprintf(format, left, processedMessages[i], right)))
		if err != nil {
			return nil
		}
	}

	// last line
	n := len(processedMessages) - 1
	_, err = w.Write([]byte(fmt.Sprintf(format, downLeft, processedMessages[n], downRight)))
	if err != nil {
		return nil
	}
	return nil
}

func writeBalloonBottom(w io.Writer, maxMessageLength int) error {
	buf := make([]byte, maxMessageLength+2)
	for i := 0; i < maxMessageLength+2; i++ {
		buf[i] = '-'
	}

	_, err := w.Write([]byte(" "))
	if err != nil {
		return err
	}

	_, err = w.Write(buf)
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(" \n"))
	if err != nil {
		return err
	}
	return nil
}
