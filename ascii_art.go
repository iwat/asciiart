package main

import (
	"strings"
)

/*
 * Transform a 'Plane' of ASCII characters to an equivalent plane where the
 * ASCII box drawings have been replaced by their Unicode counterpart.
 */
func renderAsciiToUnicode(input string) string {
	lines := strings.Split(input, "\n")

	rendered := make([]string, len(lines))
	for row := 0; row < len(lines); row++ {
		runes := []rune(lines[row])
		for col := 0; col < len(runes); col++ {
			rendered[row] += string([]rune{lookupPattern(patternFromString(lines, row, col))})
		}
	}
	return strings.Join(rendered, "\n")
}

/*
 * Find the 'Char' to replace the center of a 'Pattern'.
 */
func lookupPattern(pattern string) rune {
	for _, patternDefinition := range patternDefinitions {
		satisfied := true
		for i := 0; i < 9; i++ {
			if !connectsLike([]rune(pattern)[i], []rune(patternDefinition.pattern)[i]) {
				satisfied = false
				break
			}
		}
		if !satisfied {
			continue
		}
		return patternDefinition.char
	}
	return []rune(pattern)[4]
}

func patternFromString(lines []string, row, col int) string {
	result := [9]rune{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}

	result[0] = runeAt(lines, row-1, col-1)
	result[1] = runeAt(lines, row-1, col)
	result[2] = runeAt(lines, row-1, col+1)
	result[3] = runeAt(lines, row, col-1)
	result[4] = runeAt(lines, row, col)
	result[5] = runeAt(lines, row, col+1)
	result[6] = runeAt(lines, row+1, col-1)
	result[7] = runeAt(lines, row+1, col)
	result[8] = runeAt(lines, row+1, col+1)

	return string(result[:])
}

type pattern struct {
	pattern string
	char    rune
}

/*
 * The actual pattern definitions. For convenience, the simple patterns are at
 * the top, and more complex ones at the bottom. 'lookupPattern' will first try
 * the most complex pattern and work its way to the simpler patterns, thus
 * avoiding to choose a simpler pattern and forgetting some connection.
 */
var patternDefinitions = []pattern{
	{" | " + " v " + "   ", '▽'},
	{"   " + " ^ " + " | ", '△'},
	{"   " + " <-" + "   ", '◁'},
	{"   " + "-> " + "   ", '▷'},
	{" # " + " # " + "   ", '█'},
	{"   " + " ##" + "   ", '█'},
	{"   " + "## " + "   ", '█'},
	{"   " + " # " + " # ", '█'},
	{" | " + " '-" + "   ", '╰'},
	{" | " + "-' " + "   ", '╯'},
	{"   " + "-. " + " | ", '╮'},
	{"   " + " .-" + " | ", '╭'},
	{" | " + "-+-" + " | ", '┼'},
	{" | " + "-+-" + "   ", '┴'},
	{"   " + "-+-" + " | ", '┬'},
	{" | " + "-+ " + " | ", '┤'},
	{" | " + " +-" + " | ", '├'},
	{"   " + "-+ " + " | ", '┐'},
	{"   " + " +-" + " | ", '┌'},
	{" | " + " +-" + "   ", '└'},
	{" | " + "-+ " + "   ", '┘'},
	{" | " + "=+=" + " | ", '╪'},
	{" | " + "=+=" + "   ", '╧'},
	{"   " + "=+=" + " | ", '╤'},
	{" | " + "=+ " + " | ", '╡'},
	{" | " + " +=" + " | ", '╞'},
	{"   " + "=+ " + " | ", '╕'},
	{"   " + " +=" + " | ", '╒'},
	{" | " + " +=" + "   ", '╘'},
	{" | " + "=+ " + "   ", '╛'},
	{" : " + " : " + "   ", '┆'},
	{"   " + " : " + " : ", '┆'},
	{" | " + " : " + "   ", '┆'},
	{"   " + " : " + " | ", '┆'},
	{" | " + " | " + "   ", '│'},
	{"   " + " | " + " | ", '│'},
	{"   " + "== " + "   ", '═'},
	{"   " + " ==" + "   ", '═'},
	{"   " + "-- " + "   ", '─'},
	{"   " + " --" + "   ", '─'},
}

/*
 * Whether a character can connect to another character. For example, @+@
 * connects both horizontally (like @-@) and vertically (like @|@), so it
 * 'connectsLike' @-@, @|@, and of course like itself.
 */
func connectsLike(char, pattern rune) bool {
	switch pattern {
	case '-':
		return containsElem(char, []rune{'-', '>', '<', '─'}) || connectsLike(char, '+')
	case '=':
		return containsElem(char, []rune{'=', '>', '<', '═'}) || connectsLike(char, '+')
	case '|':
		return containsElem(char, []rune{'|', '^', 'v', '│'}) || connectsLike(char, ':') || connectsLike(char, '+')
	case ':':
		return containsElem(char, []rune{':', '┆'})
	case '+':
		return containsElem(char, []rune{'+', '└', '┘', '┌', '┐', '╘', '╛', '╒', '╕', '├', '┤', '┬', '┴', '┼', '╞', '╡', '╤', '╧', '╪'}) || connectsLike(char, '.')
	case '.':
		return containsElem(char, []rune{'\'', '.', '╭', '╮', '╯', '╰'})
	case '\'':
		return connectsLike(char, '.')
	case ' ':
		return true
	default:
		return char == pattern
	}
}

func runeAt(lines []string, row, col int) rune {
	if row < 0 || row >= len(lines) {
		return ' '
	}

	if col < 0 || col >= len([]rune(lines[row])) {
		return ' '
	}

	return []rune(lines[row])[col]
}

func containsElem(char rune, elems []rune) bool {
	for _, elem := range elems {
		if char == elem {
			return true
		}
	}
	return false
}
