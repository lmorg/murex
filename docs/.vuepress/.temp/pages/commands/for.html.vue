<template><div><h1 id="for" tabindex="-1"><a class="header-anchor" href="#for" aria-hidden="true">#</a> <code v-pre>for</code></h1>
<blockquote>
<p>A more familiar iteration loop to existing developers</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>This <code v-pre>for</code> loop is fills a small niche where <code v-pre>foreach</code> or <code v-pre>formap</code> are
inappropiate in your script. It's generally not recommended to use <code v-pre>for</code>
because it performs slower and doesn't adhere to Murex's design
philosophy. However it does offer additional flexibility around recursion.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>for ( variable; conditional; incrementation ) { code-block } -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>    » for ( i=1; i&lt;6; i++ ) { echo $i }
    1
    2
    3
    4
    5
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="syntax" tabindex="-1"><a class="header-anchor" href="#syntax" aria-hidden="true">#</a> Syntax</h3>
<p><code v-pre>for</code> is a little naughty in terms of breaking Murex's style guidelines due
to the first parameter being entered as one string treated as 3 separate code
blocks. The syntax is like this for two reasons:</p>
<ol>
<li>readability (having multiple <code v-pre>{ blocks }</code> would make scripts unsightly</li>
<li>familiarity (for those using to <code v-pre>for</code> loops in other languages</li>
</ol>
<p>The first parameter is: <code v-pre>( i=1; i&lt;6; i++ )</code>, but it is then converted into the
following code:</p>
<ol>
<li><code v-pre>let i=0</code> - declare the loop iteration variable</li>
<li><code v-pre>= i&lt;0</code> - if the condition is true then proceed to run the code in
the second parameter - <code v-pre>{ echo $i }</code></li>
<li><code v-pre>let i++</code> - increment the loop iteration variable</li>
</ol>
<p>The second parameter is the code to execute upon each iteration</p>
<h3 id="better-for-loops" tabindex="-1"><a class="header-anchor" href="#better-for-loops" aria-hidden="true">#</a> Better <code v-pre>for</code> loops</h3>
<p>Because each iteration of a <code v-pre>for</code> loop reruns the 2nd 2 parts in the first
parameter (the conditional and incrementation), <code v-pre>for</code> is very slow. Plus the
weird, non-idiomatic, way of writing the 3 parts, it's fair to say <code v-pre>for</code> is
not the recommended method of iteration and in fact there are better functions
to achieve the same thing...most of the time at least.</p>
<p>For example:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>    a: [1..5] -> foreach: i { echo $i }
    1
    2
    3
    4
    5
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>The different in performance can be measured. eg:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>    » time { a: [1..9999] -> foreach: i { out: &lt;null> $i } }
    0.097643108

    » time { for ( i=1; i&lt;10000; i=i+1 ) { out: &lt;null> $i } }
    0.663812496
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>You can also do step ranges with <code v-pre>foreach</code>:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>    » time { for ( i=10; i&lt;10001; i=i+2 ) { out: &lt;null> $i } }
    0.346254973

    » time { a: [1..999][0,2,4,6,8],10000 -> foreach i { out: &lt;null> $i } }
    0.053924326
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>...though granted the latter is a little less readable.</p>
<p>The big catch with using <code v-pre>a</code> piped into <code v-pre>foreach</code> is that values are passed
as strings rather than numbers.</p>
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
<li><RouterLink to="/commands/a.html"><code v-pre>a</code> (mkarray)</RouterLink>:
A sophisticated yet simple way to build an array or list</li>
<li><RouterLink to="/commands/break.html"><code v-pre>break</code></RouterLink>:
Terminate execution of a block within your processes scope</li>
<li><RouterLink to="/commands/foreach.html"><code v-pre>foreach</code></RouterLink>:
Iterate through an array</li>
<li><RouterLink to="/commands/formap.html"><code v-pre>formap</code></RouterLink>:
Iterate through a map or other collection of data</li>
<li><RouterLink to="/commands/if.html"><code v-pre>if</code></RouterLink>:
Conditional statement to execute different blocks of code depending on the result of the condition</li>
<li><RouterLink to="/commands/ja.html"><code v-pre>ja</code> (mkarray)</RouterLink>:
A sophisticated yet simply way to build a JSON array</li>
<li><RouterLink to="/commands/let.html"><code v-pre>let</code></RouterLink>:
Evaluate a mathematical function and assign to variable (deprecated)</li>
<li><RouterLink to="/commands/set.html"><code v-pre>set</code></RouterLink>:
Define a local variable and set it's value</li>
<li><RouterLink to="/commands/while.html"><code v-pre>while</code></RouterLink>:
Loop until condition false</li>
</ul>
</div></template>


