package command

import (
	"fmt"
	"io"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
)

func PrintCowList(w io.Writer, cowPath string) error {
	cowFiles := make([]string, 0)
	err := filepath.WalkDir(cowPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(d.Name(), ".cow") {
			return nil
		}
		cowFiles = append(cowFiles, strings.TrimSuffix(d.Name(), ".cow"))
		return nil
	})
	if err != nil {
		return err
	}
	sort.Strings(cowFiles)

	_, err = w.Write([]byte(fmt.Sprintf("Cow files in %s:\n", cowPath)))
	if err != nil {
		return err
	}
	linePos := 0
	for _, cowFile := range cowFiles {
		// too large for this line, write newline and reset line pos
		if linePos+1+len(cowFile) > 76 {
			_, err = w.Write([]byte("\n"))
			if err != nil {
				return err
			}
			linePos = 0
		}
		// write to occupied line, prefix with space for separation
		if linePos > 0 {
			_, err = w.Write([]byte(" " + cowFile))
			if err != nil {
				return err
			}
			linePos += 1 + len(cowFile)
			continue
		}
		// write to new line
		_, err = w.Write([]byte(cowFile))
		if err != nil {
			return err
		}
		linePos += len(cowFile)
	}
	_, err = w.Write([]byte("\n"))
	return err
}
