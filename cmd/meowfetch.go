package cmd

import (
	"fmt"
	"log"
	config "meowfetch/configs"
	"os"

	"github.com/akamensky/argparse"
)

var version_str = `
   /| ､         Author : Itsnexn
  (°､ ｡ 7       Github : https://github.com/itsnexn/meowfetch
   |､  ~ヽ      Version: 1.0.0
   じしf_,)/    License: MIT (https://opensource.org/licenses/MIT)
`

func Run() {
	parser  := argparse.NewParser("meowfetch", "little kitten will fetch your system's information!")
    mode    := parser.Selector("m", "mode", []string{"pride", "color"}, &argparse.Options{Help: "Ascii art mode"})
    border  := parser.Selector("b", "border", []string{"round", "single", "double", "none"},&argparse.Options{Help: "Border style"})
    color   := parser.Int("c", "color", &argparse.Options{Help: "Ascii art color"})
    version := parser.Flag("v", "version", &argparse.Options{Help: "Print Program version"})

    e := parser.Parse(os.Args)

    if e != nil {
        log.Fatal(e)
    }

    if len(*mode) > 0 { *config.Mode = *mode }
    if len(*border) > 0 { *config.BorderStyle = *border }
    if *color < 255 && *color > 0 { *config.Color = *color }
    if *version { fmt.Println(version_str); os.Exit(0) }
}

