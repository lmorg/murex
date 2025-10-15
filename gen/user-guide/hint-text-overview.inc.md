The **hint text** is a (typically) blue status line that appears directly below
your prompt. The idea behind the **hint text** is to provide clues to you as
type instructions into the prompt; but without adding distractions. It is there
to be used if you want it while keeping out of the way when you don't want it.

{{ if env "DOCGEN_TARGET=vuepress" }}
<!-- markdownlint-disable -->
<img class="vhs-hint-text">
<!-- markdownlint-restore -->
{{ else }}![hint-text](/images/vhs-hint-text-dark.gif){{ end }}