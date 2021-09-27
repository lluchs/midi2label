package main

import (
	"fmt"
	"os"

	"gitlab.com/gomidi/midi/reader"
)

func main() {
	if len(os.Args) != 2 || os.Args[1][0] == '-' {
		fmt.Printf("Usage: %s <file.mid>\n", os.Args[0])
		fmt.Println("Convert markers in a midi file to Audacity labels.")
		os.Exit(1)
	}
	filename := os.Args[1]

	var rd *reader.Reader
	rd = reader.New(reader.NoLogger(),
		reader.Marker(func(p reader.Position, text string) {
			secs := reader.TimeAt(rd, p.AbsoluteTicks).Seconds()
			fmt.Printf("%f\t%f\t%s\n", secs, secs, text)
		}),
	)

	err := reader.ReadSMFFile(rd, filename)

	if err != nil {
		fmt.Printf("could not read SMF file: %w\n", err)
		os.Exit(1)
	}
}
