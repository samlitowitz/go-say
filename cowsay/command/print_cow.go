package command

import (
	"github.com/samlitowitz/go-say/cowsay"
	"io"
	"io/ioutil"
	"strings"
)

func PrintCow(w io.Writer, opts *cowsay.Options) error {
	buf, err := ioutil.ReadFile(opts.CowPath + "/" + opts.File)
	if err != nil {
		return err
	}
	var inCowDef bool
	template := strings.Split(string(buf), "\n")
	for _, line := range template {
		// skip comments
		if !inCowDef && strings.HasPrefix(line, "##") {
			continue
		}
		// first line of cow definitions
		if !inCowDef && strings.TrimSpace(line) == "$the_cow = <<EOC;" {
			inCowDef = true
			continue
		}

		if !inCowDef {
			continue
		}

		// last line of cow definition
		if line == "EOC" {
			break
		}

		// Add end of line
		line = line + "\n"

		// Remove unnecessarily escaping of underscores
		line = strings.ReplaceAll(line, "\\_", "_")

		// Remove unnecessarily escaping of backslashes
		line = strings.ReplaceAll(line, "\\\\", "\\")

		// Replace $eyes
		line = strings.ReplaceAll(line, "$eyes", string(opts.Eyes))

		// Replace $thoughts
		line = strings.ReplaceAll(line, "$tongue", string(opts.Tongue))

		// Replace $tongue
		line = strings.ReplaceAll(line, "$thoughts", string(opts.Thoughts))

		// replace $eyes, $thoughts, and $tongue
		_, err = w.Write([]byte(line))
		if err != nil {
			return err
		}
	}

	return nil
}
