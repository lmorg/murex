//go:build !windows
// +build !windows

package man

import (
	"embed"
)

//go:embed *.1.gz
var manPages embed.FS

/*func TestMan(t *testing.T) {
	files, err := manPages.ReadDir(".")
	if err != nil {
		t.Error(err.Error())
	}

	count.Tests(t, len(files)*2)

	for _, entry := range files {
		file, err := manPages.Open(entry.Name())
		if err != nil {
			t.Errorf("%s: %s", entry.Name(), err.Error())
		}

		gz, err := gzip.NewReader(file)
		if err != nil {
			t.Errorf("%s: %s", entry.Name(), err.Error())
		}

		flags, descs := ParseByStdio(gz)
		if len(flags) == 0 {
			t.Errorf("%d flags returned for '%s'", len(flags), entry.Name())
		}
		if len(descs) == 0 {
			t.Errorf("%d descriptions returned for '%s'", len(descs), entry.Name())
		}

		err = gz.Close()
		if err != nil {
			t.Errorf("%s: %s", entry.Name(), err.Error())
		}
		err = file.Close()
		if err != nil {
			t.Errorf("%s: %s", entry.Name(), err.Error())
		}
	}
}
*/
