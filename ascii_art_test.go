package main

import (
	"testing"
)

func TestPatternFromString(t *testing.T) {
	lines := []string{
		"abc",
		"def",
		"ghi",
	}
	if patternFromString(lines, 1, 1) != "abcdefghi" {
		t.Error("1,1 should be abcdefghi")
	}
	if patternFromString(lines, 0, 0) != "    ab de" {
		t.Error("0,0 should be ____ab_de")
	}
	if patternFromString(lines, 2, 2) != "ef hi    " {
		t.Error("1,1 should be ef_hi____")
	}
}

func TestBasicElementsSimpleBoxes(t *testing.T) {
	output := renderASCIIToUnicode("" +
		"++  +-----+  +--+--+-----+  \n" +
		"++  +--+  |  |  |  |     |  \n" +
		"       |  |  +--+--+--+--+  \n" +
		"+---+  |  |  |     |  |  |  \n" +
		"+---+  +--+  +-----+--+--+  ")
	expected := "" +
		"┌┐  ┌─────┐  ┌──┬──┬─────┐  \n" +
		"└┘  └──┐  │  │  │  │     │  \n" +
		"       │  │  ├──┴──┼──┬──┤  \n" +
		"┌───┐  │  │  │     │  │  │  \n" +
		"└───┘  └──┘  └─────┴──┴──┘  "
	if output != expected {
		t.Error("bad render result\n" + output + "\n" + expected)
	}
}

func TestBasicElementsRoundedBoxes(t *testing.T) {
	output := renderASCIIToUnicode("" +
		"..  .-----.  .--+--+-----.  \n" +
		"''  '--.  |  |  |  |     |  \n" +
		"       |  |  +--+--+--+--+  \n" +
		".---.  |  |  |     |  |  |  \n" +
		"'---'  '--'  '-----+--+--'  ")
	expected := "" +
		"╭╮  ╭─────╮  ╭──┬──┬─────╮  \n" +
		"╰╯  ╰──╮  │  │  │  │     │  \n" +
		"       │  │  ├──┴──┼──┬──┤  \n" +
		"╭───╮  │  │  │     │  │  │  \n" +
		"╰───╯  ╰──╯  ╰─────┴──┴──╯  "
	if output != expected {
		t.Error("bad render result\n" + output)
	}
}

func TestBasicElementsDottedAndDoubleStrokes(t *testing.T) {
	output := renderASCIIToUnicode("" +
		"++  .-----.  +==+==+=====+  \n" +
		"++  +==+  :  |  :  |     |  \n" +
		"       :  :  +==+==+==+==+  \n" +
		"+===+  :  :  |     |  :  |  \n" +
		"+---+  '--'  +=====+==+==+  ")
	expected := "" +
		"┌┐  ╭─────╮  ╒══╤══╤═════╕  \n" +
		"└┘  ╘══╕  ┆  │  ┆  │     │  \n" +
		"       ┆  ┆  ╞══╧══╪══╤══╡  \n" +
		"╒═══╕  ┆  ┆  │     │  ┆  │  \n" +
		"└───┘  ╰──╯  ╘═════╧══╧══╛  "
	if output != expected {
		t.Error("bad render result\n" + output)
	}
}

func TestBasicElementsCastShadows(t *testing.T) {
	output := renderASCIIToUnicode("" +
		"+-------------+   \n" +
		"|             |   \n" +
		"+---+     +---+#  \n" +
		"  ##|     |#####  \n" +
		"    |     |#      \n" +
		"    +-----+#      \n" +
		"      ######      ")
	expected := "" +
		"┌─────────────┐   \n" +
		"│             │   \n" +
		"└───┐     ┌───┘█  \n" +
		"  ██│     │█████  \n" +
		"    │     │█      \n" +
		"    └─────┘█      \n" +
		"      ██████      "
	if output != expected {
		t.Error("bad render result\n" + output)
	}
}

// Already rendered portions are not affected.
func TestPropertiesIdempotent(t *testing.T) {
	output := renderASCIIToUnicode("" +
		"┌┐  ╭─────╮  ╒══╤══╤═════╕  ┌───┐   \n" +
		"└┘  ╘══╕  ┆  │  ┆  │     │  │   │   \n" +
		"       ┆  ┆  ╞══╧══╪══╤══╡  │   │█  \n" +
		"╒═══╕  ┆  ┆  │     │  ┆  │  └───┘█  \n" +
		"└───┘  ╰──╯  ╘═════╧══╧══╛    ████  ")
	expected := "" +
		"┌┐  ╭─────╮  ╒══╤══╤═════╕  ┌───┐   \n" +
		"└┘  ╘══╕  ┆  │  ┆  │     │  │   │   \n" +
		"       ┆  ┆  ╞══╧══╪══╤══╡  │   │█  \n" +
		"╒═══╕  ┆  ┆  │     │  ┆  │  └───┘█  \n" +
		"└───┘  ╰──╯  ╘═════╧══╧══╛    ████  "
	if output != expected {
		t.Error("bad render result\n" + output)
	}
}

// Existing characters can be removed.
func TestPropertiesIncrementalRemove(t *testing.T) {
	output := renderASCIIToUnicode("" +
		"┌──┬ ┬──┐\n" +
		"   │ │  │\n" +
		"╞══╪ ╪══╡\n" +
		"│        \n" +
		"├──┼ ┼──┤\n" +
		"   │ │  │\n" +
		"└──┴ ┴──┘")
	expected := "" +
		"───┐ ┌──┐\n" +
		"   │ │  │\n" +
		"╒══╛ ╘══╛\n" +
		"│        \n" +
		"└──┐ ┌──┐\n" +
		"   │ │  │\n" +
		"───┘ └──┘"
	if output != expected {
		t.Error("bad render result\n" + output)
	}
}

// New connections can be added.
func TestPropertiesIncrementalAdd(t *testing.T) {
	output := renderASCIIToUnicode("" +
		"┌──┐-┌──┐\n" +
		"│  │ │  │\n" +
		"╘══╛=╘══╛\n" +
		"|  : :  |\n" +
		"┌──┐-┌──┐\n" +
		"│  │ │  │\n" +
		"└──┘-└──┘")
	expected := "" +
		"┌──┬─┬──┐\n" +
		"│  │ │  │\n" +
		"╞══╪═╪══╡\n" +
		"│  ┆ ┆  │\n" +
		"├──┼─┼──┤\n" +
		"│  │ │  │\n" +
		"└──┴─┴──┘"
	if output != expected {
		t.Error("bad render result\n" + output)
	}
}

// Existing connections can be altered by replacing/adding characters.
func TestPropertiesIncrementalReplace(t *testing.T) {
	output := renderASCIIToUnicode("" +
		"┌──+──┐  .─────+─────.  ┌────┐    \n" +
		"│  +==+  │     |     │  │####│    \n" +
		"+==+  |  │     |     │  │#   │█   \n" +
		"└──+─-┘  │     +-----+  └────┘█#  \n" +
		"         +=====+     │    █████#  \n" +
		"╭──+─-╮  │     |     │     #####  \n" +
		"│  |  |  │     |     │            \n" +
		"╰──+──╯  '───────────'            ")
	expected := "" +
		"┌──┬──┐  ╭─────┬─────╮  ┌────┐    \n" +
		"│  ╞══╡  │     │     │  │████│    \n" +
		"╞══╡  │  │     │     │  │█   │█   \n" +
		"└──┴──┘  │     ├─────┤  └────┘██  \n" +
		"         ╞═════╡     │    ██████  \n" +
		"╭──┬──╮  │     │     │     █████  \n" +
		"│  │  │  │     │     │            \n" +
		"╰──┴──╯  ╰───────────╯            "
	if output != expected {
		t.Error("bad render result\n" + output)
	}
}

// Some connections do not work as expected (mostly because the corresponding
// Unicode characters do not exist), e.g. rounded corners with double-stroke
// lines, or connection pieces connecting horizontal single- and double-stroke
// lines.
func TestLimitations(t *testing.T) {
	output := renderASCIIToUnicode("" +
		"--+==  .==.  .--.--.   \n" +
		"  |    |  |  |  |  |   \n" +
		"==+--  '=='  .--+--'   \n" +
		"  |          |  |  |   \n" +
		"--+==  --==  '--'--'   ")
	expected := "" +
		"──┐══  ╒══╕  ╭──╮──╮   \n" +
		"  │    │  │  │  │  │   \n" +
		"══├──  ╘══╛  ╰──┼──╯   \n" +
		"  │          │  │  │   \n" +
		"──┘══  ──══  ╰──╰──╯   "
	if output != expected {
		t.Error("bad render result\n" + output)
	}
}
