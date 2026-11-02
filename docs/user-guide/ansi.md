# ANSI Constants

> Infixed constants that return ANSI escape sequences

## Description

ANSI Constants is a Murex convention of passing ANSI escape sequences into
strings. It uses the `{}` notation with the constant name placed in between two
curly braces. eg `{GREEN}`. 

## Constants

Rather than duplicate the constants from source, and risk the documentation and
implementation drifting, this document will embed the source directly below.

```go
package ansi

var constants = map[string][]byte{
	// ASCII control keys
	"^@": {0},
	"^A": {1},
	"^B": {2},
	"^C": {3},
	"^D": {4},
	"^E": {5},
	"^F": {6},
	"^G": {7},
	"^H": {8},
	"^I": {9},
	"^J": {10},
	"^K": {11},
	"^L": {12},
	"^M": {13},
	"^N": {14},
	"^O": {15},
	"^P": {16},
	"^Q": {17},
	"^R": {18},
	"^S": {19},
	"^T": {20},
	"^U": {21},
	"^V": {22},
	"^W": {23},
	"^X": {24},
	"^Y": {25},
	"^Z": {26},
	"^[": {27},
	`^/`: {28},
	"^]": {29},
	"^^": {30},
	"^_": {31},

	"^?": {127},

	// ASCII control codes
	"EOF":      {4},
	"EOT":      {4},
	"BELL":     {7},
	"BS-ISO":   {8},
	"LF":       {10},
	"CR":       {13},
	"CRLF":     {13, 10},
	"ESC":      {27},
	"ESCAPE":   {27},
	"BS-ASCII": {127},

	// ASCII escape sequences
	"CURSOR-UP":         {27, '[', 'A'},
	"CURSOR-DOWN":       {27, '[', 'B'},
	"CURSOR-FORWARDS":   {27, '[', 'C'},
	"CURSOR-BACKWARDS":  {27, '[', 'D'},
	"CURSOR-HOME":       {27, '[', 'H'},
	"CURSOR-HOME-VT100": {27, '[', '1', '~'},
	"CURSOR-END":        {27, '[', 'F'},
	"CURSOR-END-VT100":  {27, '[', '4', '~'},

	"INS":    {27, '[', '2', '~'},
	"INSERT": {27, '[', '2', '~'},
	"DEL":    {27, '[', '3', '~'},
	"DELETE": {27, '[', '3', '~'},

	// Function keys
	"F1-VT100": {27, 'O', 'P'},
	"F2-VT100": {27, 'O', 'Q'},
	"F3-VT100": {27, 'O', 'R'},
	"F4-VT100": {27, 'O', 'S'},
	"F1-VT220": {27, '[', '1', '1', '~'},
	"F2-VT220": {27, '[', '1', '2', '~'},
	"F3-VT220": {27, '[', '1', '3', '~'},
	"F4-VT220": {27, '[', '1', '4', '~'},
	"F1":       {27, 'O', 'P'},
	"F2":       {27, 'O', 'Q'},
	"F3":       {27, 'O', 'R'},
	"F4":       {27, 'O', 'S'},
	"F5":       {27, '[', '1', '5', '~'},
	"F6":       {27, '[', '1', '7', '~'},
	"F7":       {27, '[', '1', '8', '~'},
	"F8":       {27, '[', '1', '9', '~'},
	"F9":       {27, '[', '2', '0', '~'},
	"F10":      {27, '[', '2', '1', '~'},
	"F11":      {27, '[', '2', '3', '~'},
	"F12":      {27, '[', '2', '4', '~'},

	// alt-num
	"ALT-0": {27, '0'},
	"ALT-1": {27, '1'},
	"ALT-2": {27, '2'},
	"ALT-3": {27, '3'},
	"ALT-4": {27, '4'},
	"ALT-5": {27, '5'},
	"ALT-6": {27, '6'},
	"ALT-7": {27, '7'},
	"ALT-8": {27, '8'},
	"ALT-9": {27, '9'},

	// control sequence
	"CSI": {27, '['},
}

var sgr = map[string][]byte{
	// text effects
	"RESET":      {27, '[', '0', 'm'},
	"BOLD":       {27, '[', '1', 'm'},
	"ITALIC":     {27, '[', '3', 'm'}, // Not commonly supported in terminals
	"UNDERSCORE": {27, '[', '4', 'm'},
	"UNDERLINE":  {27, '[', '4', 'm'},
	"UNDEROFF":   {27, '[', '2', '4', 'm'},
	"BLINK":      {27, '[', '5', 'm'},
	"INVERT":     {27, '[', '7', 'm'},

	"ALT-FONT-1": {27, '[', '1', '1', 'm'}, // Not commonly supported in terminals
	"ALT-FONT-2": {27, '[', '1', '2', 'm'}, // Not commonly supported in terminals
	"ALT-FONT-3": {27, '[', '1', '3', 'm'}, // Not commonly supported in terminals
	"ALT-FONT-4": {27, '[', '1', '4', 'm'}, // Not commonly supported in terminals
	"ALT-FONT-5": {27, '[', '1', '5', 'm'}, // Not commonly supported in terminals
	"ALT-FONT-6": {27, '[', '1', '6', 'm'}, // Not commonly supported in terminals
	"ALT-FONT-7": {27, '[', '1', '7', 'm'}, // Not commonly supported in terminals
	"ALT-FONT-8": {27, '[', '1', '8', 'm'}, // Not commonly supported in terminals
	"ALT-FONT-9": {27, '[', '1', '9', 'm'}, // Not commonly supported in terminals
	"FRAKTUR":    {27, '[', '2', '0', 'm'}, // Not commonly supported in terminals

	// text colours
	"BLACK":   {27, '[', '3', '0', 'm'},
	"RED":     {27, '[', '3', '1', 'm'},
	"GREEN":   {27, '[', '3', '2', 'm'},
	"YELLOW":  {27, '[', '3', '3', 'm'},
	"BLUE":    {27, '[', '3', '4', 'm'},
	"MAGENTA": {27, '[', '3', '5', 'm'},
	"CYAN":    {27, '[', '3', '6', 'm'},
	"WHITE":   {27, '[', '3', '7', 'm'},

	"BLACK-BRIGHT":   {27, '[', '9', '0', 'm'},
	"RED-BRIGHT":     {27, '[', '9', '1', 'm'},
	"GREEN-BRIGHT":   {27, '[', '9', '2', 'm'},
	"YELLOW-BRIGHT":  {27, '[', '9', '3', 'm'},
	"BLUE-BRIGHT":    {27, '[', '9', '4', 'm'},
	"MAGENTA-BRIGHT": {27, '[', '9', '5', 'm'},
	"CYAN-BRIGHT":    {27, '[', '9', '6', 'm'},
	"WHITE-BRIGHT":   {27, '[', '9', '7', 'm'},

	// background colours
	"BG-BLACK":   {27, '[', '4', '0', 'm'},
	"BG-RED":     {27, '[', '4', '1', 'm'},
	"BG-GREEN":   {27, '[', '4', '2', 'm'},
	"BG-YELLOW":  {27, '[', '4', '3', 'm'},
	"BG-BLUE":    {27, '[', '4', '4', 'm'},
	"BG-MAGENTA": {27, '[', '4', '5', 'm'},
	"BG-CYAN":    {27, '[', '4', '6', 'm'},
	"BG-WHITE":   {27, '[', '4', '7', 'm'},

	"BG-BLACK-BRIGHT":   {27, '[', '1', '0', '0', 'm'},
	"BG-RED-BRIGHT":     {27, '[', '1', '0', '1', 'm'},
	"BG-GREEN-BRIGHT":   {27, '[', '1', '0', '2', 'm'},
	"BG-YELLOW-BRIGHT":  {27, '[', '1', '0', '3', 'm'},
	"BG-BLUE-BRIGHT":    {27, '[', '1', '0', '4', 'm'},
	"BG-MAGENTA-BRIGHT": {27, '[', '1', '0', '5', 'm'},
	"BG-CYAN-BRIGHT":    {27, '[', '1', '0', '6', 'm'},
	"BG-WHITE-BRIGHT":   {27, '[', '1', '0', '7', 'm'},
}
```

### How To Read The Code Above

Each line will look something a little like

```
"GREEN":   {27, 91, 51, 50, 109},
```

The part within quotes is the constant name, and the part that follows is the
sequence of bytes that are infixed.

So the example above will replace `{GREEN}` from within a string with the
byte values of 27, 91, 51, 50 and 109 (in that order).

## Unsupported Constants

If a constant does not exist in the above code, then the infix string is left
unedited.

```
# Green spelt correctly
» out "{GREEN}PASSED{RESET}"
PASSED

# Green spelt incorrectly (ie so that it doesn't exist as a valid constant)
» out "{GREEEEN}PASSED{RESET}"
{GREEEEN}PASSED
```

## Enabling / Disabling ANSI Escape Sequences

These sequences are enabled by default. To disable run the following:

```
» config set shell color false
```

You will need to add this to your Murex profile, `~/.murex_profile` to make
it persistent.

## See Also

* [Error String, strerr: `err`](../commands/err.md):
  Print a line to the stderr
* [Output String, stdout: `out`](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [Output With Type Annotation: `tout`](../commands/tout.md):
  Print a string to the stdout and set it's data-type
* [Profile Files](../user-guide/profile.md):
  A breakdown of the different files loaded on start up
* [`%(Brace Quote)`](../parser/brace-quote.md):
  Initiates or terminates a string (variables expanded)
* [`(brace quote)`](../parser/brace-quote-func.md):
  Write a string to the stdout without new line (deprecated)

<hr/>

This document was generated from [gen/user-guide/ansi_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/user-guide/ansi_doc.yaml).