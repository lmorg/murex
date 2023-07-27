<template><div><h1 id="autocomplete" tabindex="-1"><a class="header-anchor" href="#autocomplete" aria-hidden="true">#</a> <code v-pre>autocomplete</code></h1>
<blockquote>
<p>Set definitions for tab-completion in the command line</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>autocomplete</code> digests a JSON schema and uses that to define the tab-
completion rules for suggestions in the interactive command line.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>autocomplete get [ command ] -&gt; `&lt;stdout&gt;`

autocomplete set command { mxjson }
</code></pre>
<h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2>
<ul>
<li><code v-pre>get</code>
output all autocompletion schemas</li>
<li><code v-pre>set</code>
define a new autocompletion schema</li>
</ul>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="undefining-autocomplete" tabindex="-1"><a class="header-anchor" href="#undefining-autocomplete" aria-hidden="true">#</a> Undefining autocomplete</h3>
<p>Currently there is no support for undefining an autocompletion rule however
you can overwrite existing rules.</p>
<h2 id="directives" tabindex="-1"><a class="header-anchor" href="#directives" aria-hidden="true">#</a> Directives</h2>
<p>The directives are listed below. Headings are formatted as follows:</p>
<pre><code>&quot;DirectiveName&quot;: json data-type (default value)
</code></pre>
<p>Where &quot;default value&quot; is what will be auto-populated at run time if you don't
define an autocomplete schema manually. <strong>zls</strong> stands for zero-length string
(ie: &quot;&quot;).</p>
<h3 id="alias-string-zls" tabindex="-1"><a class="header-anchor" href="#alias-string-zls" aria-hidden="true">#</a> &quot;Alias&quot;: string (zls)</h3>
<p>Aliases are used inside <strong>FlagValues</strong> as a way of pointing one flag to another
without duplicating code. eg <code v-pre>-v</code> and <code v-pre>--version</code> might be the same flag. Or
<code v-pre>-?</code>, <code v-pre>-h</code> and <code v-pre>--help</code>. With <strong>Alias</strong> you can write the definitions for one
flag and then point all the synonyms as an alias to that definition.</p>
<h3 id="allowany-boolean-false" tabindex="-1"><a class="header-anchor" href="#allowany-boolean-false" aria-hidden="true">#</a> &quot;AllowAny&quot;: boolean (false)</h3>
<p>The way autocompletion works in Murex is the suggestion engine looks for
matches and if it fines one, it then moves onto the next index in the JSON
schema. This means unexpected values typed in the interactive terminal will
break the suggestion engine's ability to predict what the next expected
parameter should be. Setting <strong>AllowAny</strong> to <code v-pre>true</code> tells the suggestion
engine to accept any value as the next parameter thus allowing it to then
predict the next parameter afterwards.</p>
<p>This directive isn't usually necessary because such fields are often the last
parameter or most parameters can be detectable with a reasonable amount of
effort. However <strong>AllowAny</strong> is often required for more complex command line
tools.</p>
<h3 id="allowmultiple-boolean-false" tabindex="-1"><a class="header-anchor" href="#allowmultiple-boolean-false" aria-hidden="true">#</a> &quot;AllowMultiple&quot;: boolean (false)</h3>
<p>Set to <code v-pre>true</code> to enable multiple parameters following the same rules as defined
in this index. For example the following will suggest directories on each tab
for multiple parameters:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>autocomplete set example { [{
    "IncDirs": true,
    "AllowMultiple": true
}] }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="anyvalue-boolean-false" tabindex="-1"><a class="header-anchor" href="#anyvalue-boolean-false" aria-hidden="true">#</a> &quot;AnyValue&quot;: boolean (false)</h3>
<p>Deprecated. Please use <strong>AllowAny</strong> instead.</p>
<h3 id="autobranch-boolean-false" tabindex="-1"><a class="header-anchor" href="#autobranch-boolean-false" aria-hidden="true">#</a> &quot;AutoBranch&quot;: boolean (false)</h3>
<p>Use this in conjunction with <strong>Dynamic</strong>. If the return is an array of paths,
for example <code v-pre>[ &quot;/home/foo&quot;, &quot;/home/bar&quot; ]</code> then <strong>AutoBranch</strong> will return
the following patterns in the command line:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» example [tab]
# suggests "/home/"

» example /home/[tab]
# suggests "/home/foo" and "/home/bar"
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Please note that <strong>AutoBranch</strong>'s behavior is also dependant on a &quot;shell&quot;
<code v-pre>config</code> setting, recursive-enabled&quot;:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» config get shell recursive-enabled
true
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="cachettl-int-5" tabindex="-1"><a class="header-anchor" href="#cachettl-int-5" aria-hidden="true">#</a> &quot;CacheTTL&quot;: int (5)</h3>
<p>Dynamic autocompletions (via <strong>Dynamic</strong> or <strong>DynamicDesc</strong>) are cached to
improve interactivity performance. By default the cache is very small but you
can increase that cache or even disable it entirely. Setting this value will
define the duration (in seconds) to cache that autocompletion.</p>
<p>If you wish to disable this then set <strong>CacheTTL</strong> to <code v-pre>-1</code>.</p>
<p>This directive needs to live in the very first definition and affects all
autocompletes within the rest of the command. For example</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>autocomplete set foobar { [
    {
        "Flags": [ "--foo", "--bar" ],
        "CacheTTL": 60
    },
    {
        "Dynamic": ({
            a: [Monday..Friday]
            sleep: 3
        })
    }
] }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Here the days of the week take 3 seconds to show up as autocompletion
suggestions the first time and instantly for the next 60 seconds after.</p>
<h3 id="dynamic-string-zls" tabindex="-1"><a class="header-anchor" href="#dynamic-string-zls" aria-hidden="true">#</a> &quot;Dynamic&quot;: string (zls)</h3>
<p>This is a Murex block which returns an array of suggestions.</p>
<p>Code inside that block are executed like a function and the parameters will
mirror the same as those parameters entered in the interactive terminal.</p>
<p>Two variables are created for each <strong>Dynamic</strong> function:</p>
<ul>
<li>
<p><code v-pre>ISMETHOD</code>: <code v-pre>true</code> if the command being autocompleted is going to run as a
pipelined method. <code v-pre>false</code> if it isn't.</p>
</li>
<li>
<p><code v-pre>PREFIX</code>: contains the partial term. For example if you typed <code v-pre>hello wor[tab]</code>
then <code v-pre>$PREFIX</code> would be set to <strong>wor</strong> for <strong>hello</strong>'s
autocompletion.</p>
</li>
</ul>
<p>The expected STDOUT should be an array (list) of any data type. For example:</p>
<pre><code>[
    &quot;Monday&quot;,
    &quot;Tuesday&quot;,
    &quot;Wednesday&quot;,
    &quot;Thursday&quot;,
    &quot;Friday&quot;
]
</code></pre>
<p>You can additionally include suggestions if any of the array items exactly
matches any of the following strings:</p>
<ul>
<li><code v-pre>@IncFiles</code> (<a href="(#incfiles-boolean-false)">read more</a>)</li>
<li><code v-pre>@IncDirs</code> (<a href="(#incdirs-boolean-false)">read more</a>)</li>
<li><code v-pre>@IncExePath</code> (<a href="(#incexepath-boolean-false)">read more</a>)</li>
<li><code v-pre>@IncExeAll</code> (<a href="(#incexeall-boolean-false)">read more</a>)</li>
<li><code v-pre>@IncManPage</code> (<a href="(#incmanpage-boolean-false)">read more</a>)</li>
</ul>
<h3 id="dynamicdesc-string-zls" tabindex="-1"><a class="header-anchor" href="#dynamicdesc-string-zls" aria-hidden="true">#</a> &quot;DynamicDesc&quot;: string (zls)</h3>
<p>This is very similar to <strong>Dynamic</strong> except your function should return a
map instead of an array. Where each key is the suggestion and the value is
a description.</p>
<p>The description will appear either in the hint text or alongside the
suggestion - depending on which suggestion &quot;popup&quot; you define (see
<strong>ListView</strong>).</p>
<p>Two variables are created for each <strong>Dynamic</strong> function:</p>
<ul>
<li>
<p><code v-pre>ISMETHOD</code>: <code v-pre>true</code> if the command being autocompleted is going to run as a
pipelined method. <code v-pre>false</code> if it isn't.</p>
</li>
<li>
<p><code v-pre>PREFIX</code>: contains the partial term. For example if you typed <code v-pre>hello wor[tab]</code>
then <code v-pre>$PREFIX</code> would be set to <strong>wor</strong> for <strong>hello</strong>'s
autocompletion.</p>
</li>
</ul>
<p>The expected STDOUT should be an object (map) of any data type. The key is the
autocompletion suggestion, with the value being the description. For example:</p>
<pre><code>{
    &quot;Monday&quot;: &quot;First day of the week&quot;,
    &quot;Tuesday&quot;: &quot;Second day of the week&quot;,
    &quot;Wednesday&quot;: &quot;Third day of the week&quot;
    &quot;Thursday&quot;: &quot;Forth day of the week&quot;,
    &quot;Friday&quot;: &quot;Fifth day of the week&quot;,
}
</code></pre>
<h3 id="execcmdline-boolean-false" tabindex="-1"><a class="header-anchor" href="#execcmdline-boolean-false" aria-hidden="true">#</a> &quot;ExecCmdline&quot;: boolean (false)</h3>
<p>Sometimes you'd want your autocomplete suggestions to aware of the output
returned from the commands that preceded it. For example the suggestions
for <code v-pre>[</code> (index) will depend entirely on what data is piped into it.</p>
<p><strong>ExecCmdline</strong> tells Murex to run the commandline up until the command
which your cursor is editing and pipe that output to the STDIN of that
commands <strong>Dynamic</strong> or <strong>DynamicDesc</strong> code block.</p>
<blockquote>
<p>This is a dangerous feature to enable so <strong>ExecCmdline</strong> is only honoured
if the commandline is considered &quot;safe&quot;. <strong>Dynamic</strong> / <strong>DynamicDesc</strong>
will still be executed however if the commandline is &quot;unsafe&quot; then your
dynamic autocompletion blocks will have no STDIN.</p>
</blockquote>
<p>Because this is a dangerous feature, your partial commandline will only
execute if the following conditions are met:</p>
<ul>
<li>the commandline must be one pipeline (eg <code v-pre>;</code> tokens are not allowed)</li>
<li>the commandline must not have any new line characters</li>
<li>there must not be any redirection, including named pipes
(eg <code v-pre>cmd &lt;namedpipe&gt;</code>) and the STDOUT/STDERR switch token (<code v-pre>?</code>)</li>
<li>the commandline doesn't inline any variables (<code v-pre>$strings</code>, <code v-pre>@arrays</code>) or
functions (<code v-pre>${subshell}</code>, <code v-pre>$[index]</code>)</li>
<li>lastly all commands are whitelisted in &quot;safe-commands&quot;
(<code v-pre>config get shell safe-commands</code>)</li>
</ul>
<p>If these criteria are met, the commandline is considered &quot;safe&quot;; if any of
those conditions fail then the commandline is considered &quot;unsafe&quot;.</p>
<p>Murex will come with a number of sane commands already included in its
<code v-pre>safe-commands</code> whitelist however you can add or remove them using <code v-pre>config</code></p>
<pre><code>» function: foobar { -&gt; match foobar }
» config: eval shell safe-commands { -&gt; append foobar }
</code></pre>
<p>Remember that <strong>ExecCmdline</strong> is designed to be included with either
<strong>Dynamic</strong> or <strong>DynamicDesc</strong> and those code blocks would need to read
from STDIN:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>autocomplete set "[" { [{
    "AnyValue": true,
    "AllowMultiple": true,
    "ExecCmdline": true,
    "Dynamic": ({
        switch ${ get-type: stdin } {
            case * {
                `&lt;stdin>` -> [ 0: ] -> format json -> [ 0 ]
            }

            catch {
                `&lt;stdin>` -> formap k v { out $k } -> cast str -> append "]"
            }
        }
    })
}] }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="fileregexp-string-zls" tabindex="-1"><a class="header-anchor" href="#fileregexp-string-zls" aria-hidden="true">#</a> &quot;FileRegexp&quot;: string (zls)</h3>
<p>When set in conjunction with <strong>IncFiles</strong>, this directive will filter on files
files which match the regexp string. eg to only show &quot;.txt&quot; extensions you can
use the following:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>autocomplete set notepad.exe { [{
    "IncFiles": true,
    "FileRegexp": (\.txt)
}] }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><blockquote>
<p>Please note that you may need to double escape any regexp strings: escaping
the <code v-pre>.</code> match and then also escaping the escape character in JSON. It is
recommended you use the <code v-pre>mxjson</code> method of quoting using parentheses as this
will compile that string into JSON, automatically adding additional escaping
where required.</p>
</blockquote>
<h3 id="flagvalues-map-of-arrays-null" tabindex="-1"><a class="header-anchor" href="#flagvalues-map-of-arrays-null" aria-hidden="true">#</a> &quot;FlagValues&quot;: map of arrays (null)</h3>
<p>This is a map of the flags with the values being the same array of directive
as the top level.</p>
<p>This allows you to nest operations by flags. eg when a flag might accept
multiple parameters.</p>
<p><strong>FlagValues</strong> takes a map of arrays, eg</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>autocomplete set example { [{
    "Flags": [ "add", "delete" ],
    "FlagValues": {
        "add": [{
            "Flags": [ "foo" ]
        }],
        "delete": [{
            "Flags": [ "bar" ]
        }]
    }
}] }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>...will provide &quot;foo&quot; as a suggestion to <code v-pre>example add</code>, and &quot;bar&quot; as a
suggestion to <code v-pre>example delete</code>.</p>
<h4 id="defaults-for-matched-flags" tabindex="-1"><a class="header-anchor" href="#defaults-for-matched-flags" aria-hidden="true">#</a> Defaults for matched flags</h4>
<p>You can set default properties to all matched flags by using <code v-pre>*</code> as a
<strong>FlagValues</strong> value. To expand the above example...</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>autocomplete set example { [{
    "Flags": [ "add", "delete" ],
    "FlagValues": {
        "add": [{
            "Flags": [ "foo" ]
        }],
        "delete": [{
            "Flags": [ "bar" ]
        }],
        "*": [{
            "IncFiles"
        }]
    }
}] }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>...in this code we are saying not only does &quot;add&quot; support &quot;foo&quot; and &quot;delete&quot;
supports &quot;bar&quot;, but both &quot;add&quot; and &quot;delete&quot; also supports any filesystem files.</p>
<p>This default only applies if there is a matched <strong>Flags</strong> or <strong>FlagValues</strong>.</p>
<h4 id="defaults-for-any-flags-including-unmatched" tabindex="-1"><a class="header-anchor" href="#defaults-for-any-flags-including-unmatched" aria-hidden="true">#</a> Defaults for any flags (including unmatched)</h4>
<p>If you wanted a default which applied to all <strong>FlagValues</strong>, even when the flag
wasn't matched, then you can use a zero length string (&quot;&quot;). For example</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>autocomplete set example { [{
    "Flags": [ "add", "delete" ],
    "FlagValues": {
        "add": [{
            "Flags": [ "foo" ]
        }],
        "delete": [{
            "Flags": [ "bar" ]
        }],
        "": [{
            "IncFiles"
        }]
    }
}] }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="flags-array-of-strings-auto-populated-from-man-pages" tabindex="-1"><a class="header-anchor" href="#flags-array-of-strings-auto-populated-from-man-pages" aria-hidden="true">#</a> &quot;Flags&quot;: array of strings (auto-populated from man pages)</h3>
<p>Setting <strong>Flags</strong> is the fastest and easiest way to populate suggestions
because it is just an array of strings. eg</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>autocomplete set example { [{
    "Flags": [ "foo", "bar" ]
}] }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>If a command doesn't <strong>Flags</strong> already defined when you request a completion
suggestion but that command does have a man page, then <strong>Flags</strong> will be
automatically populated with any flags identified from an a quick parse of
the man page. However because man pages are written to be human readable
rather than machine parsable, there may not be a 100% success rate with the
automatic man page parsing.</p>
<h3 id="flagsdesc-map-of-strings-null" tabindex="-1"><a class="header-anchor" href="#flagsdesc-map-of-strings-null" aria-hidden="true">#</a> &quot;FlagsDesc&quot;: map of strings (null)</h3>
<p>This is the same concept as <strong>Flags</strong> except it is a map with the suggestion
as a key and description as a value. This distinction is the same as the
difference between <strong>Dynamic</strong> and <strong>DynamicDesc</strong>.</p>
<p>Please note that currently man page parsing cannot provide a description so
only <strong>Flags</strong> get auto-populated.</p>
<h3 id="goto-string-zls" tabindex="-1"><a class="header-anchor" href="#goto-string-zls" aria-hidden="true">#</a> &quot;Goto&quot;: string (zls)</h3>
<p>This is a <code v-pre>goto</code> in programming terms. While &quot;ugly&quot; it does allow for quick and
easy structural definitions without resorting to writing the entire
autocomplete in code.</p>
<p><strong>Goto</strong> takes a string which represents the path to jump to from the top level
of that autocomplete definition. The path should look something like:
<code v-pre>/int/string/int/string....</code> where</p>
<ul>
<li>
<p>the first character is the separator,</p>
</li>
<li>
<p>the first value is an integer that relates to the index in your autocomplete
array,</p>
</li>
<li>
<p>the second value is a string which points to the flag value map (if you
defined <strong>FlagValues</strong>),</p>
</li>
<li>
<p>the third value is the integer of the autocomplete array inside that
<strong>FlagValues</strong> map,</p>
</li>
<li>
<p>...and so on as necessary.</p>
</li>
</ul>
<p>An example of a really simple <strong>Goto</strong>:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>autocomplete set dd { [
    {
        "Flags": [ "if=", "of=", "bs=", "iflag=", "oflag=", "count=", "status=" ],
        "FlagValues": {
            "if": [{
                "IncFiles": true
            }],
            "of": [{
                "IncFiles": true
            }],
            "*": [{
                "AllowAny": true
            }]
        }
    },
    {
        "Goto": "/0"
    }
] }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p><strong>Goto</strong> is given precedence over any other directive. So ensure it's the only
directive in it's group.</p>
<h3 id="ignoreprefix-boolean-false" tabindex="-1"><a class="header-anchor" href="#ignoreprefix-boolean-false" aria-hidden="true">#</a> &quot;IgnorePrefix&quot;: boolean (false)</h3>
<p>When set to <code v-pre>true</code>, this allows <strong>Dynamic</strong> and <strong>DynamicDesc</strong> functions to
return every result and not just those that match the partial term (as would
normally be the default).</p>
<h3 id="incdirs-boolean-false" tabindex="-1"><a class="header-anchor" href="#incdirs-boolean-false" aria-hidden="true">#</a> &quot;IncDirs&quot;: boolean (false)</h3>
<p>Enable to include directories.</p>
<p>Not needed if <strong>IncFiles</strong> is set to <code v-pre>true</code>.</p>
<p>Behavior of this directive can be altered with <code v-pre>config set shell recursive-enabled</code></p>
<h3 id="incexeall-boolean-false" tabindex="-1"><a class="header-anchor" href="#incexeall-boolean-false" aria-hidden="true">#</a> &quot;IncExeAll&quot;: boolean (false)</h3>
<p>Enable this to any executables. Suggestions will include aliases, functions
builtins and any executables in <code v-pre>$PATH</code>. It will not include private functions.</p>
<h3 id="incexepath-boolean-false" tabindex="-1"><a class="header-anchor" href="#incexepath-boolean-false" aria-hidden="true">#</a> &quot;IncExePath&quot;: boolean (false)</h3>
<p>Enable this to include any executables in <code v-pre>$PATH</code>. Suggestions will not include
aliases, functions nor privates.</p>
<h3 id="incfiles-boolean-true" tabindex="-1"><a class="header-anchor" href="#incfiles-boolean-true" aria-hidden="true">#</a> &quot;IncFiles&quot;: boolean (true)</h3>
<p>Include files and directories. This is enabled by default for any commands
that don't have autocomplete defined but you will need to manually enable
it in any <code v-pre>autocomplete</code> schemas you create and want files as part of the
suggestions.</p>
<p>If you want to filter files based on file name then you can set a regexp
string to match to using <strong>FileRegexp</strong>.</p>
<h3 id="incmanpage-boolean-false" tabindex="-1"><a class="header-anchor" href="#incmanpage-boolean-false" aria-hidden="true">#</a> &quot;IncManPage&quot;: boolean (false)</h3>
<p>The default behavior for commands with no autocomplete defined is to parse the
man page and use those results. If a custom autocomplete is defined then that
man page parser is disabled by default. You can re-enable it and include its
results with other flags and behaviors you define by using this directive.</p>
<h3 id="listview-boolean-false" tabindex="-1"><a class="header-anchor" href="#listview-boolean-false" aria-hidden="true">#</a> &quot;ListView&quot;: boolean (false)</h3>
<p>This alters the appearance of the autocompletion suggestions &quot;popup&quot;. Rather
than suggestions being in a grid layout (with descriptions overwriting the
hint text) the suggestions are in a list view with the descriptions next to
them on the same row (similar to how an IDE might display it's suggestions).</p>
<h3 id="nestedcommand-boolean-false" tabindex="-1"><a class="header-anchor" href="#nestedcommand-boolean-false" aria-hidden="true">#</a> &quot;NestedCommand&quot;: boolean (false)</h3>
<p>Only enable this if the command you are autocompleting is a nested parameter
of the parent command you have types. For example with <code v-pre>sudo</code>, once you've
typed the command name you wish to elivate, then you would want suggestions
for that command rather than for <code v-pre>sudo</code> itself.</p>
<h3 id="optional-boolean-false" tabindex="-1"><a class="header-anchor" href="#optional-boolean-false" aria-hidden="true">#</a> &quot;Optional&quot;: boolean (false)</h3>
<p>Specifies if a match is required for the index in this schema. ie optional
flags.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/stdin.html"><code v-pre>&lt;stdin&gt;</code> </RouterLink>:
Read the STDIN belonging to the parent code block</li>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/alias.html"><code v-pre>alias</code></RouterLink>:
Create an alias for a command</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/function.html"><code v-pre>function</code></RouterLink>:
Define a function block</li>
<li><RouterLink to="/commands/get-type.html"><code v-pre>get-type</code></RouterLink>:
Returns the data-type of a variable or pipe</li>
<li><RouterLink to="/commands/private.html"><code v-pre>private</code></RouterLink>:
Define a private function block</li>
<li><RouterLink to="/commands/summary.html"><code v-pre>summary</code> </RouterLink>:
Defines a summary help text for a command</li>
<li><RouterLink to="/commands/switch.html"><code v-pre>switch</code></RouterLink>:
Blocks of cascading conditionals</li>
<li><RouterLink to="/types/mxjson.html">mxjson</RouterLink>:
Murex-flavoured JSON (deprecated)</li>
</ul>
</div></template>


