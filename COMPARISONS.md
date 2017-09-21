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

## Git format

Example taken from [oilshell.org/blog/2017/09/19.html](http://www.oilshell.org/blog/2017/09/19.html)

Bash / Oil Shell:

    # Escape portions of standard input delimited by special bytes
    escape-segments() {
      python -c '
    import cgi, re, sys

    print re.sub(
      r"\x01(.*)\x02",
      lambda match: cgi.escape(match.group(1)),
      sys.stdin.read())
    '
    }

    # Write an HTML table to stdout
    git-log-html() {
      echo '<table>'

      local format=$'
      <tr>
        <td> <a href="https://example.com/commit/%H">%h</a> </td>
        <td>\x01%s\x02</td>
      </tr>'
      git log -n 5 --pretty="format:$format" | escape-segments

      echo '</table>'
    }

Murex:

This can be done in one line but I'll split it for readability:

    3darray: { git log --pretty=format:"%s" } { git log --pretty=format:"%H" -> htmlesc } -> foreach rec {
        out "<tr> <td>${ $rec->[1] }</td> <td>${ $rec->[0] }</td> <tr>"
    }

What this does is create a three dimensional JSON array using the two
outputs from `git log`. Since we now have the output sorted in a proper
object we can then pull the required values and drop it into the HTML
string with clean escaping. A longer version of the code might read:

    # Define the HTML templates
    set table="<table>\n%s\n</table>"
    set row='
        <tr>
            <td> <a href="https://example.com/commit/%s">%s</a> </td>
            <td>%s</td>
        </tr>'

    # Create array with git logs output. HTML escape the titles.
    3darray: { git log --pretty=format:"%s" } { git log --pretty=format:"%H" -> htmlesc} -> set gitlog

    # For each grouped record create a HTML row.
    # Use the $row variable defined above as a template.
    $gitlog -> foreach rec {
        printf "$row" ${$rec->[1]} ${$rec->[1]} ${$rec->[0]}
    } -> set rows

    # Merge the rows into the HTML template
    printf $table $rows
