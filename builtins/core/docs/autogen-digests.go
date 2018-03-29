package docs

var digests map[string]string = map[string]string{
	`tout`: `'echo' a string to the STDOUT and set it's data-type`,
	`if`: `Conditional statement to execute different blocks of code depending on the result of the condition`,
	`alter`: `Change a value within a structured data-type and pass that change along the pipeline without altering the original source input`,
	`swivel-table`: `Rotates a table by 90 degrees`,
	`out`: `'echo' a string to the STDOUT`,
	`set`: `Define a variable and set it's value`,
	`append`: `Add data to the end of an array`,
	`print`: `Write a string to the OS STDOUT (bypassing _murex_ pipelines)`,
	`trypipe`: `Checks state of each function in a pipeline and exits block on error`,
	`try`: `Handles errors inside a block of code`,
	`swivel-datatype`: `Converts tabulated data into a map of values for serialised data-types such as JSON and YAML`,
	`pt`: `Pipe telemetry. Writes data-types and bytes written`,
	`catch`: `Handles the exception code raised by 'try' or 'trypipe'`,
	`prepend`: `Add data to the start of an array`,
	`err`: `'echo' a string to the STDERR`,
	`unset`: `Deallocates an environmental variable (aliased to '!export')`,
}
