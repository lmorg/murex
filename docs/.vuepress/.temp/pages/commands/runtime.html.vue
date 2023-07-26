<template><div><h1 id="runtime" tabindex="-1"><a class="header-anchor" href="#runtime" aria-hidden="true">#</a> <code v-pre>runtime</code></h1>
<blockquote>
<p>Returns runtime information on the internal state of Murex</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>runtime</code> is a tool for querying the internal state of Murex. It's output
will be JSON dumps.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>runtime: flags -&gt; `&lt;stdout&gt;`
</code></pre>
<p><code v-pre>builtins</code> is an alias for <code v-pre>runtime: --builtins</code>:</p>
<pre><code>builtins -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>List all the builtin data-types that support WriteArray()</p>
<pre><code>» runtime: --writearray
[
    &quot;*&quot;,
    &quot;commonlog&quot;,
    &quot;csexp&quot;,
    &quot;hcl&quot;,
    &quot;json&quot;,
    &quot;jsonl&quot;,
    &quot;qs&quot;,
    &quot;sexp&quot;,
    &quot;str&quot;,
    &quot;toml&quot;,
    &quot;yaml&quot;
]
</code></pre>
<p>List all the functions</p>
<pre><code>» runtime: --functions -&gt; [ agent aliases ]
[
    {
        &quot;Block&quot;: &quot;\n    # Launch ssh-agent\n    ssh-agent -\u003e head -n2 -\u003e [ :0 ] -\u003e prefix \&quot;export \&quot; -\u003e source\n    ssh-add: @{g \u003c!null\u003e ~/.ssh/*.key} @{g \u003c!null\u003e ~/.ssh/*.pem}\n&quot;,
        &quot;FileRef&quot;: {
            &quot;Column&quot;: 1,
            &quot;Line&quot;: 149,
            &quot;Source&quot;: {
                &quot;DateTime&quot;: &quot;2019-07-07T14:06:11.05581+01:00&quot;,
                &quot;Filename&quot;: &quot;/home/lau/.murex_profile&quot;,
                &quot;Module&quot;: &quot;profile/.murex_profile&quot;
            }
        },
        &quot;Summary&quot;: &quot;Launch ssh-agent&quot;
    },
    {
        &quot;Block&quot;: &quot;\n\t# Output the aliases in human readable format\n\truntime: --aliases -\u003e formap name alias {\n        $name -\u003e sprintf: \&quot;%10s =\u003e ${esccli @alias}\\n\&quot;\n\t} -\u003e cast str\n&quot;,
        &quot;FileRef&quot;: {
            &quot;Column&quot;: 1,
            &quot;Line&quot;: 6,
            &quot;Source&quot;: {
                &quot;DateTime&quot;: &quot;2019-07-07T14:06:10.886706796+01:00&quot;,
                &quot;Filename&quot;: &quot;(builtin)&quot;,
                &quot;Module&quot;: &quot;source/builtin&quot;
            }
        },
        &quot;Summary&quot;: &quot;Output the aliases in human readable format&quot;
    }
]
</code></pre>
<p>To get a list of every flag supported by <code v-pre>runtime</code></p>
<pre><code>» runtime: --help
[
    &quot;--aliases&quot;,
    &quot;--astcache&quot;,
    &quot;--config&quot;,
    &quot;--debug&quot;,
    &quot;--events&quot;,
    &quot;--fids&quot;,
    &quot;--flags&quot;,
    &quot;--functions&quot;,
    &quot;--help&quot;,
    &quot;--indexes&quot;,
    &quot;--marshallers&quot;,
    &quot;--memstats&quot;,
    &quot;--modules&quot;,
    &quot;--named-pipes&quot;,
    &quot;--open-agents&quot;,
    &quot;--pipes&quot;,
    &quot;--privates&quot;,
    &quot;--readarray&quot;,
    &quot;--readmap&quot;,
    &quot;--sources&quot;,
    &quot;--test-results&quot;,
    &quot;--tests&quot;,
    &quot;--unmarshallers&quot;,
    &quot;--variables&quot;,
    &quot;--writearray&quot;
]
</code></pre>
<p>Please also note that you can supply more than one flag. However when you
do use multiple flags the top level of the JSON output will be a map of the
flag names. eg</p>
<pre><code>» runtime: --pipes --tests
{
    &quot;pipes&quot;: [
        &quot;file&quot;,
        &quot;std&quot;,
        &quot;tcp-dial&quot;,
        &quot;tcp-listen&quot;,
        &quot;udp-dial&quot;,
        &quot;udp-listen&quot;
    ],
    &quot;tests&quot;: {
        &quot;state&quot;: {},
        &quot;test&quot;: []
    }
}

» runtime: --pipes
[
    &quot;file&quot;,
    &quot;std&quot;,
    &quot;tcp-dial&quot;,
    &quot;tcp-listen&quot;,
    &quot;udp-dial&quot;,
    &quot;udp-listen&quot;
]

» runtime: --tests
{
    &quot;state&quot;: {},
    &quot;test&quot;: []
}
</code></pre>
<h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2>
<ul>
<li><code v-pre>--aliases</code>
Lists all aliases</li>
<li><code v-pre>--astcache</code>
Lists some data about cached ASTs</li>
<li><code v-pre>--autocomplete</code>
Lists all <code v-pre>autocomplete</code> schemas - both user defined and automatically generated one</li>
<li><code v-pre>--builtins</code>
Lists all builtin commands, compiled into Murex</li>
<li><code v-pre>--config</code>
Lists all properties available to `config</li>
<li><code v-pre>--debug</code>
Outputs the state of debug and inspect mode</li>
<li><code v-pre>--events</code>
Lists all builtin event types and any defined events</li>
<li><code v-pre>--exports</code>
Outputs environmental variables. For Murex variables (<code v-pre>global</code> and <code v-pre>set</code>/<code v-pre>let</code>) use `--variables</li>
<li><code v-pre>--fids</code>
Lists all running processes / functions</li>
<li><code v-pre>--functions</code>
Lists all Murex global functions</li>
<li><code v-pre>--globals</code>
Lists all global variables</li>
<li><code v-pre>--help</code>
Outputs a list of <code v-pre>runtimes</code>'s flags</li>
<li><code v-pre>--indexes</code>
Lists all builtin data-types which are supported by index (<code v-pre>[</code>)</li>
<li><code v-pre>--marshallers</code>
Lists all builtin data-types with marshallers (eg required for <code v-pre>format</code>)</li>
<li><code v-pre>--memstats</code>
Outputs the running state of Go's runtime</li>
<li><code v-pre>--methods</code>
Lists all commands with a defined STDOUT and STDIN data type. This is used to generate smarter autocompletion suggestions with `-&gt;</li>
<li><code v-pre>--modules</code>
Lists all installed modules</li>
<li><code v-pre>--named-pipes</code>
Lists all named pipes defined</li>
<li><code v-pre>--not-indexes</code>
Lists all builtin data-types which are supported by index (<code v-pre>![</code>)</li>
<li><code v-pre>--open-agents</code>
Lists all registered <code v-pre>open</code> handlers</li>
<li><code v-pre>--pipes</code>
Lists builtin pipes compiled into Murex. These can be then be defined as named-pipes</li>
<li><code v-pre>--privates</code>
Lists all Murex private functions</li>
<li><code v-pre>--readarray</code>
Lists all builtin data-types which support ReadArray()</li>
<li><code v-pre>--readarraywithtype</code>
Lists all builtin data-types which support ReadArrayWithType()</li>
<li><code v-pre>--readmap</code>
Lists all builtin data-types which support ReadMap()</li>
<li><code v-pre>--sources</code>
Lists all loaded murex sources</li>
<li><code v-pre>--summaries</code>
Outputs all the override summaries</li>
<li><code v-pre>--test-results</code>
A dump of any unreported test results</li>
<li><code v-pre>--tests</code>
Lists defined tests</li>
<li><code v-pre>--unmarshallers</code>
Lists all builtin data-types with unmarshallers (eg required for <code v-pre>format</code>)</li>
<li><code v-pre>--variables</code>
Lists all local Murex variables which doesn't include environmental nor global variables</li>
<li><code v-pre>--writearray</code>
Lists all builtin data-types which support WriteArray()</li>
</ul>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="usage-in-scripts" tabindex="-1"><a class="header-anchor" href="#usage-in-scripts" aria-hidden="true">#</a> Usage in scripts</h3>
<p><code v-pre>runtime</code> should not be used in scripts because the output of <code v-pre>runtime</code> may
be subject to change as and when the internal mechanics of Murex change.
The purpose behind <code v-pre>runtime</code> is not to provide an API but rather to provide
a verbose &quot;dump&quot; of the internal running state of Murex.</p>
<p>If you require a stable API to script against then please use the respective
command line tool. For example <code v-pre>fid-list</code> instead of <code v-pre>runtime --fids</code>. Some
tools will provide a human readable output when STDOUT is a TTY but output
a script parsable version when STDOUT is not a terminal.</p>
<pre><code>» fid-list
    FID   Parent    Scope  State         Run Mode  BG   Out Pipe    Err Pipe    Command     Parameters
      0        0        0  Executing     Shell     no                           -murex
 265499        0        0  Executing     Normal    no   out         err         fid-list

» fid-list -&gt; pretty
[
    {
        &quot;FID&quot;: 0,
        &quot;Parent&quot;: 0,
        &quot;Scope&quot;: 0,
        &quot;State&quot;: &quot;Executing&quot;,
        &quot;Run Mode&quot;: &quot;Shell&quot;,
        &quot;BG&quot;: false,
        &quot;Out Pipe&quot;: &quot;&quot;,
        &quot;Err Pipe&quot;: &quot;&quot;,
        &quot;Command&quot;: &quot;-murex&quot;,
        &quot;Parameters&quot;: &quot;&quot;
    },
    {
        &quot;FID&quot;: 265540,
        &quot;Parent&quot;: 0,
        &quot;Scope&quot;: 0,
        &quot;State&quot;: &quot;Executing&quot;,
        &quot;Run Mode&quot;: &quot;Normal&quot;,
        &quot;BG&quot;: false,
        &quot;Out Pipe&quot;: &quot;out&quot;,
        &quot;Err Pipe&quot;: &quot;err&quot;,
        &quot;Command&quot;: &quot;fid-list&quot;,
        &quot;Parameters&quot;: &quot;&quot;
    },
    {
        &quot;FID&quot;: 265541,
        &quot;Parent&quot;: 0,
        &quot;Scope&quot;: 0,
        &quot;State&quot;: &quot;Executing&quot;,
        &quot;Run Mode&quot;: &quot;Normal&quot;,
        &quot;BG&quot;: false,
        &quot;Out Pipe&quot;: &quot;out&quot;,
        &quot;Err Pipe&quot;: &quot;err&quot;,
        &quot;Command&quot;: &quot;pretty&quot;,
        &quot;Parameters&quot;: &quot;&quot;
    }
]
</code></pre>
<h3 id="file-reference" tabindex="-1"><a class="header-anchor" href="#file-reference" aria-hidden="true">#</a> File reference</h3>
<p>Some of the JSON dumps produced from <code v-pre>runtime</code> will include a map called
<code v-pre>FileRef</code>. This is a trace of the source file that defined it. It is used
by Murex to help provide meaningful errors (eg with line and character
positions) however it is also useful for manually debugging user-defined
properties such as which module or script defined an <code v-pre>autocomplete</code> schema.</p>
<h3 id="debug-mode" tabindex="-1"><a class="header-anchor" href="#debug-mode" aria-hidden="true">#</a> Debug mode</h3>
<p>When <code v-pre>debug</code> is enabled garbage collection is disabled for variables and
FIDs. This means the output of <code v-pre>runtime --variables</code> and <code v-pre>runtime --fids</code>
will contain more than just the currently defined variables and running
functions.</p>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>runtime</code></li>
<li><code v-pre>builtins</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/autocomplete.html"><code v-pre>autocomplete</code></RouterLink>:
Set definitions for tab-completion in the command line</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/debug.html"><code v-pre>debug</code></RouterLink>:
Debugging information</li>
<li><RouterLink to="/commands/event.html"><code v-pre>event</code></RouterLink>:
Event driven programming for shell scripts</li>
<li><RouterLink to="/commands/export.html"><code v-pre>export</code></RouterLink>:
Define an environmental variable and set it's value</li>
<li><RouterLink to="/commands/fid-list.html"><code v-pre>fid-list</code></RouterLink>:
Lists all running functions within the current Murex session</li>
<li><RouterLink to="/commands/foreach.html"><code v-pre>foreach</code></RouterLink>:
Iterate through an array</li>
<li><RouterLink to="/commands/formap.html"><code v-pre>formap</code></RouterLink>:
Iterate through a map or other collection of data</li>
<li><RouterLink to="/commands/format.html"><code v-pre>format</code></RouterLink>:
Reformat one data-type into another data-type</li>
<li><RouterLink to="/commands/function.html"><code v-pre>function</code></RouterLink>:
Define a function block</li>
<li><RouterLink to="/commands/global.html"><code v-pre>global</code></RouterLink>:
Define a global variable and set it's value</li>
<li><RouterLink to="/commands/let.html"><code v-pre>let</code></RouterLink>:
Evaluate a mathematical function and assign to variable (deprecated)</li>
<li><RouterLink to="/commands/method.html"><code v-pre>method</code></RouterLink>:
Define a methods supported data-types</li>
<li><RouterLink to="/commands/open.html"><code v-pre>open</code></RouterLink>:
Open a file with a preferred handler</li>
<li><RouterLink to="/commands/openagent.html"><code v-pre>openagent</code></RouterLink>:
Creates a handler function for `open</li>
<li><RouterLink to="/commands/pipe.html"><code v-pre>pipe</code></RouterLink>:
Manage Murex named pipes</li>
<li><RouterLink to="/commands/pretty.html"><code v-pre>pretty</code></RouterLink>:
Prettifies JSON to make it human readable</li>
<li><RouterLink to="/commands/private.html"><code v-pre>private</code></RouterLink>:
Define a private function block</li>
<li><RouterLink to="/commands/set.html"><code v-pre>set</code></RouterLink>:
Define a local variable and set it's value</li>
<li><RouterLink to="/commands/source.html"><code v-pre>source</code> </RouterLink>:
Import Murex code from another file of code block</li>
<li><RouterLink to="/commands/test.html"><code v-pre>test</code></RouterLink>:
Murex's test framework - define tests, run tests and debug shell scripts</li>
</ul>
</div></template>


