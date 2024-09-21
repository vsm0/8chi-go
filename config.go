package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type Config struct {
	Pc uint16
	File string
	FontAddr uint16
	Seed int64
	Tps, Frequency, Cycles int
	Limited, QuirkShift, QuirkJump, QuirkLoad bool
}

func NewConfig() *Config {
	if len(os.Args) < 2 {
		fmt.Println("use flag -h for usage help")
		os.Exit(0)
	}

	pc := flag.Int("pc", 0x200, "where to load and start rom in ram")
	file := flag.String("f", "", "filename of the rom to load")
	fontAddr := flag.Int("fnt", 0x50, "where to load the font in ram")
	seed := flag.Int64("s", time.Now().UnixNano(), "seed for RNG")
	tps := flag.Int("t", 60, "ticks per second for general simulation")
	frequency := flag.Int("hz", 700, "cycle frequency / number of instructions per second")
	cycles := flag.Int("c", 0, "number of cycles to run; unused if unlimited; see -l")
	limited := flag.Bool("l", false, "whether or not to run a limited number of cycles; see -c")
	quirkShift := flag.Bool("qs", false, "whether or not to enable SHIFT quirk")
	quirkJump := flag.Bool("qj", false, "whether or not to enable JUMP quirk")
	quirkLoad := flag.Bool("ql", false, "whether or not to enable LOAD quirk")

	usage := flag.Bool("h", false, "prints usable flags")

	flag.Parse()

	if *usage {
		fmt.Println("Flags:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	return &Config{
		Pc: uint16(*pc),
		File: *file,
		FontAddr: uint16(*fontAddr),
		Seed: *seed,
		Tps: *tps,
		Frequency: *frequency,
		Cycles: *cycles,
		Limited: *limited,
		QuirkShift: *quirkShift,
		QuirkJump: *quirkJump,
		QuirkLoad: *quirkLoad,
	}
}
