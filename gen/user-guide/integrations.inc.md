## Description

Murex is highly customizable, as every good shell should be. However the Murex
community believes that extendability and support for customizations are no
substitute for good defaults.
    
Murex aims to provide you with the best support for your wider command line
needs, in its "out-of-the-box" configuration.

So with that goal in mind, the following configurations come pre-installed with
Murex's base install.

## Included Integrations

> Please note that this is not an exhaustive list of all integrations
> precompiled into Murex. A full breakdown can be viewed under `/integrations`
> in the project source directory ([view in Github](https://github.com/lmorg/murex/tree/master/integrations)).

{{ if otherdocs "integrations" }}{{ range $i,$a := otherdocs "integrations" }}{{ if gt $i 0 }}
{{ end }}* [{{ md .Title }}](../{{ md .Hierarchy }}.md):
    {{ md .Summary }}{{ end }}{{ else }}No pages currently exist for this category.{{ end }}
