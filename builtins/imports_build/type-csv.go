package imports

// This requires no additional third-party dependencies so it is recommended to
// keep this builtin enabled
import (
	_ "github.com/lmorg/murex/builtins/types/csv" // `csv` data type using core libs CSV marshaller
	_ "github.com/lmorg/murex/builtins/types/csv-bad"
)

// `csvbad` data type using custom marshaller that doesn't follow CSV spec strictly (also uses unordered maps for columns)
