# Comparisons with other shells

This is just a bit of fun inspired from a [recent comment Hacker News](https://news.ycombinator.com/item?id=14700307)
where members compared different shell one liners.

## Disclaimer

While the intention of this document is to example the expressiveness
and real world applications of _murex_, I want to stress that it is _not_
an attempt to argue that _murex_ is better than any of the other shells
exampled.

Further to that point it is worth remembering that some of the other
shells described will have different goals to _murex_, such as POSIX
compatibility or a different approach to syntax. Or they might be a
single domain utility like `jq`.

Lastly in all of examples provided there will be a multitude of ways of
writing the code. This is true for both _murex_ and the other tools too.

So please treat these examples as a fun comparison between different
tools to help demonstrate using _murex_ on real world problems rather
than a recommendation nor "flamewar" about which method is best.

## Github issue parsing

Bash + jq:

    curl -s https://api.github.com/repos/lmorg/murex/issues | jq -r  '.[] | [(.number|tostring), .title] | join(": ")'

Elvish:

    curl -s https://api.github.com/repos/lmorg/murex/issues | from-json | each explode | each [issue]{ echo $issue[number]: $issue[title] }

Murex:

    open https://api.github.com/repos/lmorg/murex/issues -> foreach { -> [ number title ] -> sprintf "%2s: %s\n" }

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

    2darray: { git log --pretty=format:"%s" } { git log --pretty=format:"%H" -> htmlesc } -> foreach rec {
        out "<tr> <td>$rec[1]</td> <td>$rec[0]</td> <tr>"
    }

What this does is create a two dimensional JSON array using the two
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
    2darray: { git log --pretty=format:"%s" } { git log --pretty=format:"%H" -> htmlesc} -> set gitlog

    # For each grouped record create a HTML row.
    # Use the $row variable defined above as a template.
    $gitlog -> foreach rec {
        printf "$row" $rec[1] $rec[1] $rec[0]
    } -> set rows

    # Merge the rows into the HTML template
    printf $table $rows
