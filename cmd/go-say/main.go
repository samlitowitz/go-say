package main

import (
	"flag"
	"github.com/samlitowitz/go-say/cowsay"
	"github.com/samlitowitz/go-say/cowsay/command"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func getCowPath() string {
	cowPath, exists := os.LookupEnv("COWPATH")
	if exists && len(cowPath) > 0 {
		return cowPath
	}
	return "/usr/local/share/cows"
}

func main() {
	// Predefined eye/tongue combinations
	var borg, dead, greedy, paranoid, stoned, tired, wired, young bool
	flag.BoolVar(&borg, "b", false, "Borg")
	flag.BoolVar(&dead, "d", false, "Dead")
	flag.BoolVar(&greedy, "g", false, "Greedy")
	flag.BoolVar(&paranoid, "p", false, "Paranoid")
	flag.BoolVar(&stoned, "s", false, "Stoned")
	flag.BoolVar(&tired, "t", false, "Tired")
	flag.BoolVar(&wired, "w", false, "Wired")
	flag.BoolVar(&young, "y", false, "Young")

	// options
	var list, newLines bool
	var eyes, cowFile, tongue string
	var wrapWidth int
	flag.BoolVar(&list, "l", false, "List")
	flag.BoolVar(&newLines, "n", false, "Strip tabs and new lines")
	flag.StringVar(&eyes, "e", "oo", "Eyes")
	flag.StringVar(&cowFile, "f", "default.cow", "Cow file")
	flag.StringVar(&tongue, "T", "  ", "Tongue")
	flag.IntVar(&wrapWidth, "W", 40, "Wrap width")

	// new option, helps with delve
	var inputFile string
	flag.StringVar(&inputFile, "if", "", "Input file instead of stdin")

	flag.Parse()

	cowPath := getCowPath()

	if list {
		err := command.PrintCowList(os.Stdout, cowPath)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	message := strings.TrimSpace(strings.Join(flag.Args(), " "))
	if len(inputFile) != 0 {
		tmp, err := ioutil.ReadFile(inputFile)
		if err != nil {
			log.Fatal(err)
		}
		message = string(tmp)
	}
	if len(message) == 0 {
		tmp, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		message = string(tmp)
	}

	_, err := cowsay.IsValidEyes(cowsay.Eyes(eyes))
	if err != nil {
		log.Fatal(err)
	}
	_, err = cowsay.IsValidTongue(cowsay.Tongue(tongue))
	if err != nil {
		log.Fatal(err)
	}

	opts := cowsay.NewOptions()
	opts.Thoughts = cowsay.Say
	opts.Eyes = cowsay.Eyes(eyes)
	opts.Tongue = cowsay.Tongue(tongue)
	opts.File = cowFile
	opts.CowPath = cowPath
	opts.WrapWidth = wrapWidth
	opts.NewLines = newLines

	switch true {
	case borg == true:
		opts.Eyes = cowsay.Borg
	case dead == true:
		opts.Eyes = cowsay.DeadEyes
		opts.Tongue = cowsay.DeadTongue
	case greedy == true:
		opts.Eyes = cowsay.Greedy
	case paranoid == true:
		opts.Eyes = cowsay.Paranoid
	case stoned == true:
		opts.Eyes = cowsay.StonedEyes
		opts.Tongue = cowsay.StonedTongue
	case tired == true:
		opts.Eyes = cowsay.Tired
	case wired == true:
		opts.Eyes = cowsay.Wired
	case young == true:
		opts.Eyes = cowsay.Young
	}

	err = command.PrintBalloon(os.Stdout, opts, message)
	if err != nil {
		log.Fatal(err)
	}

	err = command.PrintCow(os.Stdout, opts)
	if err != nil {
		log.Fatal(err)
	}
}
