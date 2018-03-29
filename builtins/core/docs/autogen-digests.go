package docs

var digests map[string]string = map[string]string{
	`pt`:              `Pipe telemetry. Writes data-types and bytes written`,
	`set`:             `Define a variable and set it's value`,
	`swivel-table`:    `Rotates a table by 90 degrees`,
	`err`:             `'echo' a string to the STDERR`,
	`out`:             `'echo' a string to the STDOUT`,
	`if`:              `Conditional statement to execute different blocks of code depending on the result of the condition`,
	`alter`:           `Change a value within a structured data-type and pass that change along the pipeline without altering the original source input`,
	`append`:          `Add data to the end of an array`,
	`catch`:           `Handles the exception code raised by 'try' or 'trypipe'`,
	`try`:             `Handles errors inside a block of code`,
	`unset`:           `Deallocates an environmental variable (aliased to '!export')`,
	`prepend`:         `Add data to the start of an array`,
	`tout`:            `'echo' a string to the STDOUT and set it's data-type`,
	`trypipe`:         `Checks state of each function in a pipeline and exits block on error`,
	`swivel-datatype`: `Converts tabulated data into a map of values for serialised data-types such as JSON and YAML`,
	`print`:           `Write a string to the OS STDOUT (bypassing _murex_ pipelines)`,
}
