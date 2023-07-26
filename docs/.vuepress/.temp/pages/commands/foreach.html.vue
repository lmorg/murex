<template><div><h1 id="foreach" tabindex="-1"><a class="header-anchor" href="#foreach" aria-hidden="true">#</a> <code v-pre>foreach</code></h1>
<blockquote>
<p>Iterate through an array</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>foreach</code> reads an array or map from STDIN and iterates through it, running
a code block for each iteration with the value of the iterated element passed
to it.</p>
<p>By default <code v-pre>foreach</code>'s output data type is inherieted from its input data type.
For example is STDIN is <code v-pre>yaml</code> then so will STDOUT. The only exception to this
is if STDIN is <code v-pre>json</code> in which case STDOUT will be jsonlines (<code v-pre>jsonl</code>), or when
additional flags are used such as <code v-pre>--jmap</code>.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<p><code v-pre>{ code-block }</code> reads from a variable and writes to an array / unbuffered STDOUT:</p>
<pre><code>`&lt;stdin&gt;` -&gt; foreach variable { code-block } -&gt; `&lt;stdout&gt;`
</code></pre>
<p><code v-pre>{ code-block }</code> reads from STDIN and writes to an array / unbuffered STDOUT:</p>
<pre><code>`&lt;stdin&gt;` -&gt; foreach { -&gt; code-block } -&gt; `&lt;stdout&gt;`
</code></pre>
<p><code v-pre>foreach</code> writes to a buffered JSON map:</p>
<pre><code>`&lt;stdin&gt;` -&gt; foreach --jmap variable { code-block (map key) } { code-block (map value) } -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>There are two basic ways you can write a <code v-pre>foreach</code> loop depending on how you
want the iterated element passed to the code block.</p>
<p>The first option is to specify a temporary variable which can be read by the
code block:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>    » a [1..3] -> foreach i { out $i }
    1
    2
    3
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><blockquote>
<p>Please note that the variable is specified <strong>without</strong> the dollar prefix,
then used in the code block <strong>with</strong> the dollar prefix.</p>
</blockquote>
<p>The second option is for the code block's STDIN to read the element:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>    » a [1..3] -> foreach { -> cat }
    1
    2
    3
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><blockquote>
<p>STDIN can only be read as the first command. If you cannot process the
element on the first command then it is recommended you use the first
option (passing a variable) instead.</p>
</blockquote>
<h3 id="writing-json-maps" tabindex="-1"><a class="header-anchor" href="#writing-json-maps" aria-hidden="true">#</a> Writing JSON maps</h3>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>    » ja [Monday..Friday] -> foreach --jmap day { out $day -> left 3 } { $day }
    {
        "Fri": "Friday",
        "Mon": "Monday",
        "Thu": "Thursday",
        "Tue": "Tuesday",
        "Wed": "Wednesday"
    }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="using-steps-to-jump-iterations-by-more-than-1-one" tabindex="-1"><a class="header-anchor" href="#using-steps-to-jump-iterations-by-more-than-1-one" aria-hidden="true">#</a> Using steps to jump iterations by more than 1 (one)</h3>
<p>You can step through an array, list or table in jumps of user definable
quantities. The value passed in STDIN and $VAR will be an array of all
the records within that step range. For example:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>    » %[1..10] -> foreach --step 3 value { out "Iteration $.i: $value" }
    Iteration 1: [
        1,
        2,
        3
    ]
    Iteration 2: [
        4,
        5,
        6
    ]
    Iteration 3: [
        7,
        8,
        9
    ]
    Iteration 4: [
        10
    ]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2>
<ul>
<li><code v-pre>--jmap</code>
Write a <code v-pre>json</code> map to STDOUT instead of an array</li>
<li><code v-pre>--step</code>
<code v-pre>&lt;int&gt;</code> Iterates in steps. Value passed to block is an array of items in the step range. Not (yet) supported with `--jmap</li>
</ul>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="meta-values" tabindex="-1"><a class="header-anchor" href="#meta-values" aria-hidden="true">#</a> Meta values</h3>
<p>Meta values are a JSON object stored as the variable <code v-pre>$.</code>. The meta variable
will get overwritten by any other block which invokes meta values. So if you
wish to persist meta values across blocks you will need to reassign <code v-pre>$.</code>, eg</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>    %[1..3] -> foreach {
        meta_parent = $.
        %[7..9] -> foreach {
            out "$(meta_parent.i): $.i"
        }
    }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>The following meta values are defined:</p>
<ul>
<li><code v-pre>i</code>: iteration number</li>
</ul>
<h3 id="preserving-the-data-type-when-no-flags-used" tabindex="-1"><a class="header-anchor" href="#preserving-the-data-type-when-no-flags-used" aria-hidden="true">#</a> Preserving the data type (when no flags used)</h3>
<p><code v-pre>foreach</code> will preserve the data type read from STDIN in all instances where
data is being passed along the pipeline and push that data type out at the
other end:</p>
<ul>
<li>The temporary variable will be created with the same data-type as
<code v-pre>foreach</code>'s STDIN, or the data type of the array element (eg if it is a
string or number)</li>
<li>The code block's STDIN will have the same data-type as <code v-pre>foreach</code>'s STDIN</li>
<li><code v-pre>foreeach</code>'s STDOUT will also be the same data-type as it's STDIN (or <code v-pre>jsonl</code>
(jsonlines) where STDIN was <code v-pre>json</code> because <code v-pre>jsonl</code> better supports streaming)</li>
</ul>
<p>This last point means you may need to <code v-pre>cast</code> your data if you're writing
data in a different format. For example the following is creating a YAML list
however the data-type is defined as <code v-pre>json</code>:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>    » ja [1..3] -> foreach i { out "- $i" }
    - 1
    - 2
    - 3

    » ja [1..3] -> foreach i { out "- $i" } -> debug -> [[ /Data-Type/Murex ]]
    json
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Thus any marshalling or other data-type-aware API's would fail because they
are expecting <code v-pre>json</code> and receiving an incompatible data format.</p>
<p>This can be resolved via <code v-pre>cast</code>:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>    » ja [1..3] -> foreach i { out "- $i" } -> cast yaml
    - 1
    - 2
    - 3

    » ja [1..3] -> foreach i { out "- $i" } -> cast yaml -> debug -> [[ /Data-Type/Murex ]]
    yaml
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>The output is the same but now it's defined as <code v-pre>yaml</code> so any further pipelined
processes will now automatically use YAML marshallers when reading that data.</p>
<h3 id="tips-when-writing-json-inside-for-loops" tabindex="-1"><a class="header-anchor" href="#tips-when-writing-json-inside-for-loops" aria-hidden="true">#</a> Tips when writing JSON inside for loops</h3>
<p>One of the drawbacks (or maybe advantages, depending on your perspective) of
JSON is that parsers generally expect a complete file for processing in that
the JSON specification requires closing tags for every opening tag. This means
it's not always suitable for streaming. For example</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>    » ja [1..3] -> foreach i { out ({ "$i": $i }) }
    { "1": 1 }
    { "2": 2 }
    { "3": 3 }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p><strong>What does this even mean and how can you build a JSON file up sequentially?</strong></p>
<p>One answer if to write the output in a streaming file format and convert back
to JSON</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>    » ja [1..3] -> foreach i { out (- "$i": $i) }
    - "1": 1
    - "2": 2
    - "3": 3

    » ja [1..3] -> foreach i { out (- "$i": $i) } -> cast yaml -> format json
    [
        {
            "1": 1
        },
        {
            "2": 2
        },
        {
            "3": 3
        }
    ]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p><strong>What if I'm returning an object rather than writing one?</strong></p>
<p>The problem with building JSON structures from existing structures is that you
can quickly end up with invalid JSON due to the specifications strict use of
commas.</p>
<p>For example in the code below, each item block is it's own object and there are
no <code v-pre>[ ... ]</code> encapsulating them to denote it is an array of objects, nor are
the objects terminated by a comma.</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>    » config -> [ shell ] -> formap k v { $v -> alter /Foo Bar }
    {
        "Data-Type": "bool",
        "Default": true,
        "Description": "Display the interactive shell's hint text helper. Please note, even when this is disabled, it will still appear when used for regexp searches and other readline-specific functions",
        "Dynamic": false,
        "Foo": "Bar",
        "Global": true,
        "Value": true
    }
    {
        "Data-Type": "block",
        "Default": "{ progress $PID }",
        "Description": "Murex function to execute when an `exec` process is stopped",
        "Dynamic": false,
        "Foo": "Bar",
        "Global": true,
        "Value": "{ progress $PID }"
    }
    {
        "Data-Type": "bool",
        "Default": true,
        "Description": "ANSI escape sequences in Murex builtins to highlight syntax errors, history completions, {SGR} variables, etc",
        "Dynamic": false,
        "Foo": "Bar",
        "Global": true,
        "Value": true
    }
    ...
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Luckily JSON also has it's own streaming format: JSON lines (<code v-pre>jsonl</code>). We can
<code v-pre>cast</code> this output as <code v-pre>jsonl</code> then <code v-pre>format</code> it back into valid JSON:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>    » config -> [ shell ] -> formap k v { $v -> alter /Foo Bar } -> cast jsonl -> format json
    [
        {
            "Data-Type": "bool",
            "Default": true,
            "Description": "Write shell history (interactive shell) to disk",
            "Dynamic": false,
            "Foo": "Bar",
            "Global": true,
            "Value": true
        },
        {
            "Data-Type": "int",
            "Default": 4,
            "Description": "Maximum number of lines with auto-completion suggestions to display",
            "Dynamic": false,
            "Foo": "Bar",
            "Global": true,
            "Value": "6"
        },
        {
            "Data-Type": "bool",
            "Default": true,
            "Description": "Display some status information about the stop process when ctrl+z is pressed (conceptually similar to ctrl+t / SIGINFO on some BSDs)",
            "Dynamic": false,
            "Foo": "Bar",
            "Global": true,
            "Value": true
        },
    ...
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h4 id="foreach-will-automatically-cast-it-s-output-as-jsonl-if-it-s-stdin-type-is-json" tabindex="-1"><a class="header-anchor" href="#foreach-will-automatically-cast-it-s-output-as-jsonl-if-it-s-stdin-type-is-json" aria-hidden="true">#</a> <code v-pre>foreach</code> will automatically cast it's output as <code v-pre>jsonl</code> <em>if</em> it's STDIN type is <code v-pre>json</code></h4>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>    » ja: [Tom,Dick,Sally] -> foreach: name { out Hello $name }
    Hello Tom
    Hello Dick
    Hello Sally

    » ja [Tom,Dick,Sally] -> foreach name { out Hello $name } -> debug -> [[ /Data-Type/Murex ]]
    jsonl

    » ja: [Tom,Dick,Sally] -> foreach: name { out Hello $name } -> format: json
    [
        "Hello Tom",
        "Hello Dick",
        "Hello Sally"
    ]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/apis/ReadArrayWithType.html"><code v-pre>ReadArrayWithType()</code> (type)</RouterLink>:
Read from a data type one array element at a time and return the elements contents and data type</li>
<li><RouterLink to="/commands/element.html"><code v-pre>[[</code> (element)</RouterLink>:
Outputs an element from a nested structure</li>
<li><RouterLink to="/commands/a.html"><code v-pre>a</code> (mkarray)</RouterLink>:
A sophisticated yet simple way to build an array or list</li>
<li><RouterLink to="/commands/break.html"><code v-pre>break</code></RouterLink>:
Terminate execution of a block within your processes scope</li>
<li><RouterLink to="/commands/cast.html"><code v-pre>cast</code></RouterLink>:
Alters the data type of the previous function without altering it's output</li>
<li><RouterLink to="/commands/debug.html"><code v-pre>debug</code></RouterLink>:
Debugging information</li>
<li><RouterLink to="/commands/for.html"><code v-pre>for</code></RouterLink>:
A more familiar iteration loop to existing developers</li>
<li><RouterLink to="/commands/formap.html"><code v-pre>formap</code></RouterLink>:
Iterate through a map or other collection of data</li>
<li><RouterLink to="/commands/format.html"><code v-pre>format</code></RouterLink>:
Reformat one data-type into another data-type</li>
<li><RouterLink to="/commands/if.html"><code v-pre>if</code></RouterLink>:
Conditional statement to execute different blocks of code depending on the result of the condition</li>
<li><RouterLink to="/commands/ja.html"><code v-pre>ja</code> (mkarray)</RouterLink>:
A sophisticated yet simply way to build a JSON array</li>
<li><RouterLink to="/types/json.html"><code v-pre>json</code> </RouterLink>:
JavaScript Object Notation (JSON)</li>
<li><RouterLink to="/types/jsonl.html"><code v-pre>jsonl</code> </RouterLink>:
JSON Lines</li>
<li><RouterLink to="/commands/left.html"><code v-pre>left</code></RouterLink>:
Left substring every item in a list</li>
<li><RouterLink to="/commands/out.html"><code v-pre>out</code></RouterLink>:
Print a string to the STDOUT with a trailing new line character</li>
<li><RouterLink to="/commands/while.html"><code v-pre>while</code></RouterLink>:
Loop until condition false</li>
<li><RouterLink to="/types/yaml.html"><code v-pre>yaml</code> </RouterLink>:
YAML Ain't Markup Language (YAML)</li>
</ul>
</div></template>


