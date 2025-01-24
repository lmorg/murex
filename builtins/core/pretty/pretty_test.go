package pretty_test

import (
	"fmt"
	"strings"
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

const (
	svgPretty = `
<svg aria-label="Version: 6.3.4247" height="20" role="img" width="122" xlink="http://www.w3.org/1999/xlink" xmlns="http://www.w3.org/2000/svg">
    <clipPath id="r">
        <rect fill="#fff" height="20" rx="3" width="122"/>
    </clipPath>
    <g clip-path="url(#r)">
        <rect fill="#555" height="20" width="51"/>
        <rect fill="#2ea44f" height="20" width="71" x="51"/>
        <rect fill="url(#s)" height="20" width="122"/>
    </g>
    <g fill="#fff" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" font-size="110" text-anchor="middle" text-rendering="geometricPrecision">
        <text aria-hidden="true" fill="#010101" fill-opacity="0.3" textLength="410" transform="scale(.1)" x="265" y="150">Version</text>
        <text fill="#fff" textLength="410" transform="scale(.1)" x="265" y="140">Version</text>
        <text aria-hidden="true" fill="#010101" fill-opacity="0.3" textLength="610" transform="scale(.1)" x="855" y="150">6.3.4247</text>
        <text fill="#fff" textLength="610" transform="scale(.1)" x="855" y="140">6.3.4247</text>
    </g>
    <linearGradient id="s" x2="0" y2="100%">
        <stop offset="0" stop-color="#bbb" stop-opacity="0.1"/>
        <stop offset="1" stop-opacity="0.1"/>
    </linearGradient>
    <title>Version: 6.3.4247</title>
</svg>`

	svgMinified = `<svg aria-label="Version: 6.3.4247" height="20" role="img" width="122" xlink="http://www.w3.org/1999/xlink" xmlns="http://www.w3.org/2000/svg"><clipPath id="r"><rect fill="#fff" height="20" rx="3" width="122"/></clipPath><g clip-path="url(#r)"><rect fill="#555" height="20" width="51"/><rect fill="#2ea44f" height="20" width="71" x="51"/><rect fill="url(#s)" height="20" width="122"/></g><g fill="#fff" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" font-size="110" text-anchor="middle" text-rendering="geometricPrecision"><text aria-hidden="true" fill="#010101" fill-opacity="0.3" textLength="410" transform="scale(.1)" x="265" y="150">Version</text><text fill="#fff" textLength="410" transform="scale(.1)" x="265" y="140">Version</text><text aria-hidden="true" fill="#010101" fill-opacity="0.3" textLength="610" transform="scale(.1)" x="855" y="150">6.3.4247</text><text fill="#fff" textLength="610" transform="scale(.1)" x="855" y="140">6.3.4247</text></g><linearGradient id="s" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity="0.1"/><stop offset="1" stop-opacity="0.1"/></linearGradient><title>Version: 6.3.4247</title></svg>`
)

func TestMarshallers(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  fmt.Sprintf(`%%(%s) -> :xml: format json -> format xml`, svgMinified),
			Stdout: strings.TrimSpace(svgMinified),
		},
		{
			Block:  fmt.Sprintf(`%%(%s) -> :xml: format json -> format xml -> pretty`, svgMinified),
			Stdout: strings.TrimSpace(svgPretty),
		},
	}

	test.RunMurexTests(tests, t)
}
