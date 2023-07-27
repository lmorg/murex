<template><div><h1 id="test" tabindex="-1"><a class="header-anchor" href="#test" aria-hidden="true">#</a> <code v-pre>test</code></h1>
<blockquote>
<p>Murex's test framework - define tests, run tests and debug shell scripts</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>test</code> is used to define tests, run tests and debug Murex shell scripts.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<p>Define an inlined test</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>test: define test-name { json-properties }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>Define a state report</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>test: state name { code block }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>Define a unit test</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>test: unit function|private|open|event test-name { json-properties }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>Enable or disable boolean test states (more options available in <code v-pre>config</code>)</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>test: config [ enable|!enable ] [ verbose|!verbose ] [ auto-report|!auto-report ]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>Disable test mode</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>!test
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>Execute a function with testing enabled</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>test: run { code-block }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>Execute unit test(s)</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>test: run package/module/test-name|*
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>Write report</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>test: report
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Inlined test</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>function: hello-world {
    test: define example {
        "StdoutRegex": (^Hello World$)
    }

    out: &lt;test_example> "Hello Earth"
}

test: run { hello-world }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Unit test</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>test: unit function aliases {
    "PreBlock": ({
        alias ALIAS_UNIT_TEST=example param1 param2 param3
    }),
    "StdoutRegex": "([- _0-9a-zA-Z]+ => .*?\n)+",
    "StdoutType": "str",
    "PostBlock": ({
        !alias ALIAS_UNIT_TEST
    })
}

function: aliases {
    # Output the aliases in human readable format
    runtime: --aliases -> formap: name alias {
        $name -> sprintf: "%10s => ${esccli @alias}\n"
    } -> cast: str
}

test: run aliases
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="report" tabindex="-1"><a class="header-anchor" href="#report" aria-hidden="true">#</a> Report</h3>
<p><code v-pre>test: report</code> is only needed if <code v-pre>config: test auto-report</code> is set false.
However <code v-pre>test: run</code> automatically enables <strong>auto-report</strong>.</p>
<p>When the report is generated, be it automatically or manually triggered, it
flushes the table of pending reports.</p>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>test</code></li>
<li><code v-pre>!test</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/namedpipe.html"><code v-pre>&lt;&gt;</code> / <code v-pre>read-named-pipe</code></RouterLink>:
Reads from a Murex named pipe</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
</ul>
</div></template>


