import{_ as i}from"./plugin-vue_export-helper-c27b6911.js";import{r as a,o as r,c as l,d as e,b as d,w as t,e as n,f as o}from"./app-45f7c304.js";const u={},c=o(`<h1 id="args" tabindex="-1"><a class="header-anchor" href="#args" aria-hidden="true">#</a> <code>args</code></h1><blockquote><p>Command line flag parser for Murex shell scripting</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>One of the nuisances of shell scripts is handling flags. More often than not your script will be littered with <code>$1</code> still variables and not handle flags shifting in placement amongst a group of parameters. <code>args</code> aims to fix that by providing a common tool for parsing flags.</p><p><code>args</code> takes a name of a variable to assign the result of the parsed parameters as well as a JSON structure containing the result. It also returns a non-zero exit number if there is an error when parsing.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>args var-name { json-block } -&gt; \`&lt;stdout&gt;\`
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>#!/usr/bin/env murex

# First we define what parameters to accept:
# Pass the \`args\` function a JSON string (because JSON objects share the same braces as murex block, you can enter JSON
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

out &quot;The structure of \\$args is: \${$args-&gt;pretty}\\n\\n&quot;


# Some example usage:
# -------------------

!if { $(args.Flags.--bool) } {
    out &quot;Flag \`--bool\` was not set.&quot;
}

# \`&lt;!null&gt;\` redirects the STDERR to a named pipe. In this instance it&#39;s the &#39;null&#39; pipe so equivalent to 2&gt;/dev/null
# thus we are just suppressing any error messages.
try &lt;!null&gt; {
    $(args.Flags.--str) -&gt; set fStr
    $(args.Flags.--num) -&gt; set fNum

    out &quot;Defined Flags:&quot;
    out &quot;  --str == $(fStr)&quot;
    out &quot;  --num == $(fNum)&quot;
}

catch {
    err &quot;Missing \`--str\` and/or \`--num\` flags.&quot;
}

$args[Additional] -&gt; foreach flag {
    out &quot;Additional argument (ie not assigned to a flag): \`$(flag)\`.&quot;
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,10);function v(m,b){const s=a("RouterLink");return r(),l("div",null,[c,e("ul",null,[e("li",null,[d(s,{to:"/user-guide/reserved-vars.html"},{default:t(()=>[n("Reserved Variables")]),_:1}),n(": Special variables reserved by Murex")])])])}const p=i(u,[["render",v],["__file","args.html.vue"]]);export{p as default};
