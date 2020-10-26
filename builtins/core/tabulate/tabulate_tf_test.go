package tabulate

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test"
)

var (
	inTf = `Usage: terraform apply [options] [DIR-OR-PLAN]

  Builds or changes infrastructure according to Terraform configuration
  files in DIR.

  By default, apply scans the current directory for the configuration
  and applies the changes appropriately. However, a path to another
  configuration or an execution plan can be provided. Execution plans can be
  used to only execute a pre-determined set of actions.

Options:

  -auto-approve          Skip interactive approval of plan before applying.

  -backup=path           Path to backup the existing state file before
                         modifying. Defaults to the "-state-out" path with
                         ".backup" extension. Set to "-" to disable backup.

  -compact-warnings      If Terraform produces any warnings that are not
                         accompanied by errors, show them in a more compact
                         form that includes only the summary messages.

  -lock=true             Lock the state file when locking is supported.

  -lock-timeout=0s       Duration to retry a state lock.

  -input=true            Ask for input for variables if not directly set.

  -no-color              If specified, output won't contain any color.

  -parallelism=n         Limit the number of parallel resource operations.
                         Defaults to 10.

  -refresh=true          Update state prior to checking for differences. This
                         has no effect if a plan file is given to apply.

  -state=path            Path to read and save state (unless state-out
                         is specified). Defaults to "terraform.tfstate".

  -state-out=path        Path to write state to that is different than
                         "-state". This can be used to preserve the old
                         state.

  -target=resource       Resource to target. Operation will be limited to this
                         resource and its dependencies. This flag can be used
                         multiple times.

  -var 'foo=bar'         Set a variable in the Terraform configuration. This
                         flag can be set multiple times.

  -var-file=foo          Set variables in the Terraform configuration from
                         a file. If "terraform.tfvars" or any ".auto.tfvars"
                         files are present, they will be automatically loaded.`

	csvTfNoFlags = `Builds or changes infrastructure according to Terraform configuration
files in DIR.
"By default, apply scans the current directory for the configuration"
"and applies the changes appropriately. However, a path to another"
configuration or an execution plan can be provided. Execution plans can be
used to only execute a pre-determined set of actions.
-auto-approve,Skip interactive approval of plan before applying.
-backup=path,Path to backup the existing state file before
"modifying. Defaults to the ""-state-out"" path with"
""".backup"" extension. Set to ""-"" to disable backup."
-compact-warnings,If Terraform produces any warnings that are not
"accompanied by errors, show them in a more compact"
form that includes only the summary messages.
-lock=true,Lock the state file when locking is supported.
-lock-timeout=0s,Duration to retry a state lock.
-input=true,Ask for input for variables if not directly set.
-no-color,"If specified, output won't contain any color."
-parallelism=n,Limit the number of parallel resource operations.
Defaults to 10.
-refresh=true,Update state prior to checking for differences. This
has no effect if a plan file is given to apply.
-state=path,Path to read and save state (unless state-out
"is specified). Defaults to ""terraform.tfstate""."
-state-out=path,Path to write state to that is different than
"""-state"". This can be used to preserve the old"
state.
-target=resource,Resource to target. Operation will be limited to this
resource and its dependencies. This flag can be used
multiple times.
-var 'foo=bar',Set a variable in the Terraform configuration. This
flag can be set multiple times.
-var-file=foo,Set variables in the Terraform configuration from
"a file. If ""terraform.tfvars"" or any "".auto.tfvars"""
"files are present, they will be automatically loaded."
`

	csvTfKeyValue = `-auto-approve,Skip interactive approval of plan before applying.
-backup=path,Path to backup the existing state file before
-compact-warnings,If Terraform produces any warnings that are not
-lock=true,Lock the state file when locking is supported.
-lock-timeout=0s,Duration to retry a state lock.
-input=true,Ask for input for variables if not directly set.
-no-color,"If specified, output won't contain any color."
-parallelism=n,Limit the number of parallel resource operations.
-refresh=true,Update state prior to checking for differences. This
-state=path,Path to read and save state (unless state-out
-state-out=path,Path to write state to that is different than
-target=resource,Resource to target. Operation will be limited to this
-var 'foo=bar',Set a variable in the Terraform configuration. This
-var-file=foo,Set variables in the Terraform configuration from
`

	csvTfColumnWraps = `-auto-approve,Skip interactive approval of plan before applying.
-backup=path,"Path to backup the existing state file before modifying. Defaults to the ""-state-out"" path with "".backup"" extension. Set to ""-"" to disable backup."
-compact-warnings,"If Terraform produces any warnings that are not accompanied by errors, show them in a more compact form that includes only the summary messages."
-lock=true,Lock the state file when locking is supported.
-lock-timeout=0s,Duration to retry a state lock.
-input=true,Ask for input for variables if not directly set.
-no-color,"If specified, output won't contain any color."
-parallelism=n,Limit the number of parallel resource operations. Defaults to 10.
-refresh=true,Update state prior to checking for differences. This has no effect if a plan file is given to apply.
-state=path,"Path to read and save state (unless state-out is specified). Defaults to ""terraform.tfstate""."
-state-out=path,"Path to write state to that is different than ""-state"". This can be used to preserve the old state."
-target=resource,Resource to target. Operation will be limited to this resource and its dependencies. This flag can be used multiple times.
-var 'foo=bar',Set a variable in the Terraform configuration. This flag can be set multiple times.
-var-file=foo,"Set variables in the Terraform configuration from a file. If ""terraform.tfvars"" or any "".auto.tfvars"" files are present, they will be automatically loaded."
`

	jsonTfNoFlags = `{"-auto-approve":"Skip interactive approval of plan before applying.","-backup=path":"Path to backup the existing state file before","-compact-warnings":"If Terraform produces any warnings that are not","-input=true":"Ask for input for variables if not directly set.","-lock-timeout=0s":"Duration to retry a state lock.","-lock=true":"Lock the state file when locking is supported.","-no-color":"If specified, output won't contain any color.","-parallelism=n":"Limit the number of parallel resource operations.","-refresh=true":"Update state prior to checking for differences. This","-state-out=path":"Path to write state to that is different than","-state=path":"Path to read and save state (unless state-out","-target=resource":"Resource to target. Operation will be limited to this","-var 'foo=bar'":"Set a variable in the Terraform configuration. This","-var-file=foo":"Set variables in the Terraform configuration from"}`

	jsonTfColumnWraps = `{"-auto-approve":"Skip interactive approval of plan before applying.","-backup=path":"Path to backup the existing state file before modifying. Defaults to the \"-state-out\" path with \".backup\" extension. Set to \"-\" to disable backup.","-compact-warnings":"If Terraform produces any warnings that are not accompanied by errors, show them in a more compact form that includes only the summary messages.","-input=true":"Ask for input for variables if not directly set.","-lock-timeout=0s":"Duration to retry a state lock.","-lock=true":"Lock the state file when locking is supported.","-no-color":"If specified, output won't contain any color.","-parallelism=n":"Limit the number of parallel resource operations. Defaults to 10.","-refresh=true":"Update state prior to checking for differences. This has no effect if a plan file is given to apply.","-state-out=path":"Path to write state to that is different than \"-state\". This can be used to preserve the old state.","-state=path":"Path to read and save state (unless state-out is specified). Defaults to \"terraform.tfstate\".","-target=resource":"Resource to target. Operation will be limited to this resource and its dependencies. This flag can be used multiple times.","-var 'foo=bar'":"Set a variable in the Terraform configuration. This flag can be set multiple times.","-var-file=foo":"Set variables in the Terraform configuration from a file. If \"terraform.tfvars\" or any \".auto.tfvars\" files are present, they will be automatically loaded."}`
)

func TestTabulateTerraform(t *testing.T) {
	test.RunMethodTest(t,
		cmdTabulate, "tabulate",
		inTf,
		types.Generic,
		[]string{},
		csvTfNoFlags,
		nil,
	)

	test.RunMethodTest(t,
		cmdTabulate, "tabulate",
		inTf,
		types.Generic,
		[]string{fKeyVal},
		csvTfKeyValue,
		nil,
	)

	test.RunMethodTest(t,
		cmdTabulate, "tabulate",
		inTf,
		types.Generic,
		[]string{fColumnWraps, fKeyVal},
		csvTfColumnWraps,
		nil,
	)

	test.RunMethodTest(t,
		cmdTabulate, "tabulate",
		inTf,
		types.Generic,
		[]string{fMap},
		jsonTfNoFlags,
		nil,
	)

	test.RunMethodTest(t,
		cmdTabulate, "tabulate",
		inTf,
		types.Generic,
		[]string{fColumnWraps, fMap},
		jsonTfColumnWraps,
		nil,
	)
}
