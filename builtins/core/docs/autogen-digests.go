package docs

var digests map[string]string = map[string]string{
	`out`:             `'echo' a string to the STDOUT`,
	`if`:              `Conditional statement to execute different blocks of code depending on the result of the condition`,
	`unset`:           `Deallocates an environmental variable (aliased to '!export')`,
	`swivel-table`:    `Rotates a table by 90 degrees`,
	`pt`:              `Pipe telemetry. Writes data-types and bytes written`,
	`trypipe`:         `Checks state of each function in a pipeline and exits block on error`,
	`set`:             `Define a variable and set it's value`,
	`alter`:           `Change a value within a structured data-type and pass that change along the pipeline without altering the original source input`,
	`prepend`:         `Add data to the start of an array`,
	`print`:           `Write a string to the OS STDOUT (bypassing _murex_ pipelines)`,
	`catch`:           `Handles the exception code raised by 'try' or 'trypipe'`,
	`try`:             `Handles errors inside a block of code`,
	`append`:          `Add data to the end of an array`,
	`swivel-datatype`: `Converts tabulated data into a map of values for serialised data-types such as JSON and YAML`,
	`err`:             `'echo' a string to the STDERR`,
	`tout`:            `'echo' a string to the STDOUT and set it's data-type`,
}
