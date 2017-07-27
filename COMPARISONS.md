# Comparisons with other shells

This is just a bit of fun inspired from a [recent comment Hacker News](https://news.ycombinator.com/item?id=14700307)
where members compared different shell one liners.

## Github issue parsing

Bash + jq:

    curl -s https://api.github.com/repos/lmorg/murex/issues | jq -r  '.[] | [(.number|tostring), .title] | join(": ")'

Elvish:

    curl https://api.github.com/repos/lmorg/murex/issues | from-json | each explode | each [issue]{ echo $issue[number]: $issue[title] }

Murex:

    get https://api.github.com/repos/lmorg/murex/issues -> [ Body ] -> foreach { -> [ number title ] -> sprintf "%2s: %s\n" }