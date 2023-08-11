{{ if env "DOCGEN_TARGET=vuepress" }}---
title: {{ md .Title }}
index: true
category:
  - {{ md .ID }}
---

{{ end }}<h1>{{ md .Title }}</h1>{{ if .Description }}

{{ md .Description }}{{ end }}

{{ if env "DOCGEN_TARGET=" }}<h2>Table of Contents</h2>

<div id="toc">

- [Language Tour](#language-tour)
- [User Guides](#user-guides)
- [Builtin Commands](#builtin-commands)
  - [Standard Builtins](#standard-builtins)
  - [Optional Builtins](#optional-builtins)
- [Data Types](#data-types)
- [Events](#events)
- [API Reference](#api-reference)

</div>
{{ end }}
## Language Tour

The [Language Tour](/tour.md) is a great introduction into the Murex language.

## User Guides

{{ if .Documents }}{{ range $i,$a := .Documents }}{{ if gt $i 0 }}
{{ end }}* [{{ md .Title }}](../{{ md .Hierarchy }}.md):
  {{ md .Summary }}{{ end }}{{ else }}No pages currently exist for this category.{{ end }}

{{ if env "DOCGEN_TARGET=" }}## Operators And Tokens

{{ if otherdocs "parser" }}{{ range $i,$a := otherdocs "parser" }}{{ if gt $i 0 }}
{{ end }}* [{{ md .Title }}]({{ md .Hierarchy }}.md):
  {{ md .Summary }}{{ end }}{{ else }}No pages currently exist for this category.{{ end }}

## Builtin Commands

### Standard Builtins

{{ if otherdocs "commands" }}{{ range $i,$a := otherdocs "commands" }}{{ if gt $i 0 }}
{{ end }}* [{{ md .Title }}](../{{ md .Hierarchy }}.md):
  {{ md .Summary }}{{ end }}{{ else }}No pages currently exist for this category.{{ end }}

### Optional Builtins

These builtins are optional. `select` is included as part of the default build
but can be disabled without breaking functionality. The other optional builtins
are only included by default on Windows.

{{ if otherdocs "optional" }}{{ range $i,$a := otherdocs "optional" }}{{ if gt $i 0 }}
{{ end }}* [{{ md .Title }}](../{{ md .Hierarchy }}.md):
  {{ md .Summary }}{{ end }}{{ else }}No pages currently exist for this category.{{ end }}

## Data Types

{{ if otherdocs "types" }}{{ range $i,$a := otherdocs "types" }}{{ if gt $i 0 }}
{{ end }}* [{{ md .Title }}](../{{ md .Hierarchy }}.md):
  {{ md .Summary }}{{ end }}{{ else }}No pages currently exist for this category.{{ end }}

## Events

{{ if otherdocs "events" }}{{ range $i,$a := otherdocs "events" }}{{ if gt $i 0 }}
{{ end }}* [{{ md .Title }}](../{{ md .Hierarchy }}.md):
  {{ md .Summary }}{{ end }}{{ else }}No pages currently exist for this category.{{ end }}

## API Reference

These API docs are provided for any developers wishing to write their own builtins.

{{ if otherdocs "apis" }}{{ range $i,$a := otherdocs "apis" }}{{ if gt $i 0 }}
{{ end }}* [{{ md .Title }}](../{{ md .Hierarchy }}.md):
  {{ md .Summary }}{{ end }}{{ else }}No pages currently exist for this category.{{ end }}
{{ end }}