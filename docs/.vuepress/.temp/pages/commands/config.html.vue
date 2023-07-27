<template><div><h1 id="config" tabindex="-1"><a class="header-anchor" href="#config" aria-hidden="true">#</a> <code v-pre>config</code></h1>
<blockquote>
<p>Query or define Murex runtime settings</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Rather than Murex runtime settings being definable via obscure environmental
variables, Murex instead supports a registry of config defined via the
<code v-pre>config</code> command. This means any preferences and/or runtime config becomes
centralised and discoverable.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<p>List all settings:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>config -> &lt;stdout>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>Get a setting:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>config get app key -> `&lt;stdout>`
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>Set a setting:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>config set app key value

`&lt;stdin>` -> config set app key

config eval app key { -> code-block }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Define a new config setting:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>config define app key { mxjson }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>Reset a setting to it's default value:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>!config app key

config default app key
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Using <code v-pre>eval</code> to append to an array (in this instance, adding a function
name to the list of &quot;safe&quot; commands)</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» function: foobar { -> match foobar }
» config: eval shell safe-commands { -> append foobar }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>With regards to <code v-pre>config</code>, the following terms are applied:</p>
<h3 id="app" tabindex="-1"><a class="header-anchor" href="#app" aria-hidden="true">#</a> &quot;app&quot;</h3>
<p>This refers to a grouped category of settings. For example the name of a built
in.</p>
<p>Other app names include</p>
<ul>
<li><code v-pre>shell</code>: for &quot;global&quot; (system wide) Murex settings</li>
<li><code v-pre>proc</code>: for scoped Murex settings</li>
<li><code v-pre>http</code>: for settings that are applied to any processes which use the builtin
HTTP user agent (eg <code v-pre>open</code>, <code v-pre>get</code>, <code v-pre>getfile</code>, <code v-pre>post</code>)</li>
<li><code v-pre>test</code>: settings for Murex's test frameworks</li>
<li><code v-pre>index</code>: settings for <code v-pre>[</code> (index)</li>
</ul>
<h3 id="key" tabindex="-1"><a class="header-anchor" href="#key" aria-hidden="true">#</a> &quot;key&quot;</h3>
<p>This refers to the config setting itself. For example the &quot;app&quot; might be <code v-pre>http</code>
but the &quot;key&quot; might be <code v-pre>timeout</code> - where the &quot;key&quot;, in this instance, holds the
value for how long any HTTP user agents might wait before timing out.</p>
<h3 id="value" tabindex="-1"><a class="header-anchor" href="#value" aria-hidden="true">#</a> &quot;value&quot;</h3>
<p>Value is the actual value of a setting. So the value for &quot;app&quot;: <code v-pre>http</code>, &quot;key&quot;:
<code v-pre>timeout</code> might be <code v-pre>10</code>. eg</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» config get http timeout
10
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="scope-scoped" tabindex="-1"><a class="header-anchor" href="#scope-scoped" aria-hidden="true">#</a> &quot;scope&quot; / &quot;scoped&quot;</h3>
<p>Settings in <code v-pre>config</code>, by default, are scoped per function and module. Any
functions called will inherit the settings of it's caller parent. However any
child functions that then change the settings will only change settings for it's
own function and not the parent caller.</p>
<p>Please note that <code v-pre>config</code> settings are scoped differently to local variables.</p>
<h3 id="global" tabindex="-1"><a class="header-anchor" href="#global" aria-hidden="true">#</a> &quot;global&quot;</h3>
<p>Global settings defined inside a function will affect settings queried inside
another executing function (same concept as global variables).</p>
<h2 id="directives" tabindex="-1"><a class="header-anchor" href="#directives" aria-hidden="true">#</a> Directives</h2>
<p>The directives for <code v-pre>config define</code> are listed below. Headings are formatted
as follows:</p>
<pre><code>&quot;DirectiveName&quot;: json data-type (default value)
</code></pre>
<p>Where &quot;default value&quot; is what will be auto-populated if you don't include that
directive (or &quot;required&quot; if the directive must be included).</p>
<h3 id="datatype-string-required" tabindex="-1"><a class="header-anchor" href="#datatype-string-required" aria-hidden="true">#</a> &quot;DataType&quot;: string (required)</h3>
<p>This is the Murex data-type for the value.</p>
<h3 id="description-string-required" tabindex="-1"><a class="header-anchor" href="#description-string-required" aria-hidden="true">#</a> &quot;Description&quot;: string (required)</h3>
<p>Description is a required field to force developers into writing meaning hints
enabling the discoverability of settings within Murex.</p>
<h3 id="global-boolean-false" tabindex="-1"><a class="header-anchor" href="#global-boolean-false" aria-hidden="true">#</a> &quot;Global&quot;: boolean (false)</h3>
<p>This defines whether this setting is global or scoped.</p>
<p>All <strong>Dynamic</strong> settings <em>must</em> also be <strong>Global</strong>. This is because <strong>Dynamic</strong>
settings rely on a state that likely isn't scoped (eg the contents of a config
file).</p>
<h3 id="default-any-required" tabindex="-1"><a class="header-anchor" href="#default-any-required" aria-hidden="true">#</a> &quot;Default&quot;: any (required)</h3>
<p>This is the initialized and default value.</p>
<h3 id="options-array-nil" tabindex="-1"><a class="header-anchor" href="#options-array-nil" aria-hidden="true">#</a> &quot;Options&quot;: array (nil)</h3>
<p>Some suggested options (if known) to provide as autocompletion suggestions in
the interactive command line.</p>
<h3 id="dynamic-map-of-strings-nil" tabindex="-1"><a class="header-anchor" href="#dynamic-map-of-strings-nil" aria-hidden="true">#</a> &quot;Dynamic&quot;: map of strings (nil)</h3>
<p>Only use this if config options need to be more than just static values stored
inside Murex's runtime. Using <strong>Dynamic</strong> means <code v-pre>autocomplete get app key</code>
and <code v-pre>autocomplete set app key value</code> will spawn off a subshell running a code
block defined from the <code v-pre>Read</code> and <code v-pre>Write</code> mapped values. eg</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code># Create the example config file
(this is the default value) -> > example.conf

# mxjson format, so we can have comments and block quotes: #, (, )
config define example test ({
    "Description": "This is only an example",
    "DataType": "str",
    "Global": true,
    "Dynamic": {
        "Read": ({
            open example.conf
        }),
        "Write": ({
            -> > example.conf
        })
    },
    # read the config file to get the default value
    "Default": "${open example.conf}"
})
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>It's also worth noting the different syntax between <strong>Read</strong> and <strong>Default</strong>.
The <strong>Read</strong> code block is being executed when the <strong>Read</strong> directive is being
requested, whereas the <strong>Default</strong> code block is being executed when the JSON
is being read.</p>
<p>In technical terms, the <strong>Default</strong> code block is being executed by Murex
when <code v-pre>config define</code> is getting executed where as the <strong>Read</strong> and <strong>Write</strong>
code blocks are getting stored as a JSON string and then executed only when
those hooks are getting triggered.</p>
<p>See the <code v-pre>mxjson</code> data-type for more details.</p>
<h3 Read:string()="" id="dynamic" tabindex="-1"><a class="header-anchor" href="#dynamic" aria-hidden="true">#</a> &quot;Dynamic&quot;:</h3>
<p>This is executed when <code v-pre>autocomplete get app key</code> is ran. The STDOUT of the code
block is the setting's value.</p>
<h3 Write:string()="" id="dynamic-1" tabindex="-1"><a class="header-anchor" href="#dynamic-1" aria-hidden="true">#</a> &quot;Dynamic&quot;:</h3>
<p>This is executed when <code v-pre>autocomplete</code> is setting a value (eg <code v-pre>set</code>, <code v-pre>default</code>,
<code v-pre>eval</code>). is ran. The STDIN of the code block is the new value.</p>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>config</code></li>
<li><code v-pre>!config</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/element.html"><code v-pre>[[</code> (element)</RouterLink>:
Outputs an element from a nested structure</li>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/append.html"><code v-pre>append</code></RouterLink>:
Add data to the end of an array</li>
<li><RouterLink to="/commands/event.html"><code v-pre>event</code></RouterLink>:
Event driven programming for shell scripts</li>
<li><RouterLink to="/commands/function.html"><code v-pre>function</code></RouterLink>:
Define a function block</li>
<li><RouterLink to="/commands/get.html"><code v-pre>get</code></RouterLink>:
Makes a standard HTTP request and returns the result as a JSON object</li>
<li><RouterLink to="/commands/getfile.html"><code v-pre>getfile</code></RouterLink>:
Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.</li>
<li><RouterLink to="/commands/match.html"><code v-pre>match</code></RouterLink>:
Match an exact value in an array</li>
<li><RouterLink to="/commands/open.html"><code v-pre>open</code></RouterLink>:
Open a file with a preferred handler</li>
<li><RouterLink to="/commands/post.html"><code v-pre>post</code></RouterLink>:
HTTP POST request with a JSON-parsable return</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>runtime</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
<li><RouterLink to="/types/mxjson.html">mxjson</RouterLink>:
Murex-flavoured JSON (deprecated)</li>
</ul>
</div></template>


