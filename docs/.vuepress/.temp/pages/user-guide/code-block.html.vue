<template><div><h1 id="code-block-parsing" tabindex="-1"><a class="header-anchor" href="#code-block-parsing" aria-hidden="true">#</a> Code Block Parsing</h1>
<blockquote>
<p>Overview of how code blocks are parsed</p>
</blockquote>
<p>The murex parser creates ASTs ahead of interpreting each block of code. However
the AST is only generated for a block at a time. Take this sample code:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>function example {
    # An example function
    if { $ENVVAR } then {
        out: 'foobar'
    }
    out: 'Finished!'
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>When that code is run <code v-pre>function</code> is executed with the parameters <code v-pre>example</code> and
<code v-pre>{ ... }</code> but the contents of <code v-pre>{ ... }</code> isn't converted into ASTs until someone
calls <code v-pre>example</code> elsewhere in the shell.</p>
<p>When <code v-pre>example</code> (the Murex function defined above) is executed the parser will
then generate AST of the commands inside said function but not any blocks that
are associated with those functions. eg the AST would look something like this:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>[
    {
        "Command": "if",
        "Parameters": [
            "{ $ENVVAR }",
            "then",
            "{\n        out: 'foobar'\n    }"
        ]
    },
    {
        "Command": "out",
        "Parameters": [
            "Finished!"
        ]
    }
]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><blockquote>
<p>Please note this is a mock JSON structure rather than a representation of the
actual AST that would be created. Parameters are stored differently to allow
infixing of variables; and there also needs to be data shared about how
pipelining (eg STDOUT et al) is chained. What is being captured above is only
the command name and parameters.</p>
</blockquote>
<p>So when <code v-pre>if</code> executes, the conditional (the first parameter) is then parsed and
turned into ASTs and executed. Then the last parameter (the <strong>then</strong> block) is
parsed and turned into ASTs, if the first conditional is true.</p>
<p>This sequence of parsing is defined within the <code v-pre>if</code> builtin rather than
Murex's parser. That means any code blocks are parsed only when a builtin
specifically requests that they are executed.</p>
<p>With murex, there's no distinction between text and code. It's up to commands
to determine if they want to execute a parameter as code or not (eg a curly
brace block might be JSON).</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/ansi.html">ANSI Constants</RouterLink>:
Infixed constants that return ANSI escape sequences</li>
<li><RouterLink to="/parser/curly-brace.html">Curly Brace (<code v-pre>{</code>, <code v-pre>}</code>) Tokens</RouterLink>:
Initiates or terminates a code block</li>
<li><RouterLink to="/user-guide/pipeline.html">Pipeline</RouterLink>:
Overview of what a &quot;pipeline&quot; is</li>
<li><RouterLink to="/user-guide/schedulers.html">Schedulers</RouterLink>:
Overview of the different schedulers (or 'run modes') in Murex</li>
</ul>
</div></template>


