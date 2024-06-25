The **hint text** is a (typically) blue status line that appears directly below
your prompt. The idea behind the **hint text** is to provide clues to you as
type instructions into the prompt; but without adding distractions. It is there
to be used if you want it while keeping out of the way when you don't want it.

{{ if env "DOCGEN_TARGET=vuepress" }}
<!-- markdownlint-disable -->
<figure>
    <img src="/screenshot-hint-text-rsync.png?v={{ env "COMMITHASHSHORT" }}" class="centre-image"/>
    <figcaption>`rsync` flag, with example, pulled from `man` pages ({{link "read more" "man-pages"}})</figcaption>
</figure>
<figure>
    <img src="/screenshot-hint-text-egrep.png?v={{ env "COMMITHASHSHORT" }}" class="centre-image"/>
    <figcaption>`egrep` is an {{link "alias" "alias"}} so also show the destination command</figcaption>
</figure>    
<!-- markdownlint-restore -->
{{ end }}
