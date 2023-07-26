<template><div><h1 id="args" tabindex="-1"><a class="header-anchor" href="#args" aria-hidden="true">#</a> <code v-pre>args</code></h1>
<blockquote>
<p>Command line flag parser for Murex shell scripting</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>One of the nuisances of shell scripts is handling flags. More often than not
your script will be littered with <code v-pre>$1</code> still variables and not handle flags
shifting in placement amongst a group of parameters. <code v-pre>args</code> aims to fix that by
providing a common tool for parsing flags.</p>
<p><code v-pre>args</code> takes a name of a variable to assign the result of the parsed parameters
as well as a JSON structure containing the result. It also returns a non-zero
exit number if there is an error when parsing.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>args var-name { json-block } -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>#!/usr/bin/env murex

# First we define what parameters to accept:
# Pass the `args` function a JSON string (because JSON objects share the same braces as murex block, you can enter JSON
# directly as unescaped values as parameters in murex).
#
# --str: str == string data type
# --num: num == numeric data type
# --bool: bool == flag used == true, missing == false
# -b: --bool == alias of --bool flag
args: args %{
    AllowAdditional: true
    Flags: {
        --str:  str
        --num:  num
        --bool: bool
        -b: --bool
    }
}
catch {
    # Lets check for errors in the command line parameters. If they exist then
    # print the error and then exit.
    err $args.error
    exit 1
}

out "The structure of \$args is: ${$args->pretty}\n\n"


# Some example usage:
# -------------------

!if { $(args.Flags.--bool) } {
    out "Flag `--bool` was not set."
}

# `&lt;!null>` redirects the STDERR to a named pipe. In this instance it's the 'null' pipe so equivalent to 2>/dev/null
# thus we are just suppressing any error messages.
try &lt;!null> {
    $(args.Flags.--str) -> set fStr
    $(args.Flags.--num) -> set fNum

    out "Defined Flags:"
    out "  --str == $(fStr)"
    out "  --num == $(fNum)"
}

catch {
    err "Missing `--str` and/or `--num` flags."
}

$args[Additional] -> foreach flag {
    out "Additional argument (ie not assigned to a flag): `$(flag)`."
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/reserved-vars.html">Reserved Variables</RouterLink>:
Special variables reserved by Murex</li>
</ul>
</div></template>


