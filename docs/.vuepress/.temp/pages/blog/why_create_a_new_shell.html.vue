<template><div><h1 id="why-create-a-new-shell-blog" tabindex="-1"><a class="header-anchor" href="#why-create-a-new-shell-blog" aria-hidden="true">#</a> Why Create A New Shell? - Blog</h1>
<blockquote>
<p>This article discuses the motivation behind creating a new shell</p>
</blockquote>
<h2 id="reading-lines" tabindex="-1"><a class="header-anchor" href="#reading-lines" aria-hidden="true">#</a> Reading lines</h2>
<p>Bash shell is a popular command-line interface for Unix and Linux operating systems. One of its many useful features is the ability to read files line by line. This can be helpful for processing large files or performing repetitive tasks on a file's contents. The most basic way to read a file line by line is to use a while loop with the <code v-pre>read</code> command:</p>
<pre><code>#!/bin/bash

while read line; do
    echo $line
done &lt; file.txt
</code></pre>
<p>In this example, the <code v-pre>while</code> loop reads each line of the <code v-pre>file.txt</code> file and stores it in the <code v-pre>$line</code> variable. The <code v-pre>echo</code> command then prints the contents of the <code v-pre>$line</code> variable to the console. The <code v-pre>&lt;</code> symbol tells Bash to redirect the contents of the file into the loop.</p>
<p>The <code v-pre>read</code> command is what actually reads each line of the file. By default, it reads one line at a time and stores it in the variable specified. You can also use the <code v-pre>-r</code> option with the <code v-pre>read</code> command to disable backslash interpretation, which can be useful when dealing with files that contain backslashes.</p>
<p>Another useful feature of Bash is the ability to perform operations on each line of a file before processing it. For example, you can use <code v-pre>sed</code> to replace text within each line of a file:</p>
<pre><code>#!/bin/bash

while read line; do
    new_line=$(echo $line | sed 's/foo/bar/g')
    echo $new_line
done &lt; file.txt
</code></pre>
<p>In this example, <code v-pre>sed</code> replaces all instances of &quot;foo&quot; with &quot;bar&quot; in each line of the file. The modified line is then stored in the <code v-pre>$new_line</code> variable and printed to the console.</p>
<p>Of course you could just run</p>
<pre><code>sed 's/foo/bar/g' file.txt
</code></pre>
<p>...but the reasons for the for this contrived example will follow.</p>
<h2 id="but-what-if-my-files-aren-t-line-delimited" tabindex="-1"><a class="header-anchor" href="#but-what-if-my-files-aren-t-line-delimited" aria-hidden="true">#</a> But what if my files aren't line delimited?</h2>
<p>The problem with Bash, and all traditional Linux or UNIX shells, is that they operate on byte streams. To be fair, this isn't so much a fault of Bash <em>per se</em> but more a result of the design of UNIX where (almost) everything is a file, including pipes. This means everything is treated as bytes. Unlike, for example, Powershell which passes .NET objects around. Byte streams make complete sense when you're working on '70s or '80s mainframes but it is a little less productive in the modern world of structured formats like JSON.</p>
<p>So how do you read lists from objects in, for example, JSON? In Bash, this isn't so easy. You need to rely on third party tools like <code v-pre>jq</code> and often end up throwing out all of the older core utilities, like <code v-pre>sed</code>, that have become muscle memory. In Powershell it is a lot easier but you then need .NET wrappers around your existing command line tools. In Murex you can have your proverbial cake and eat it.</p>
<h2 id="structured-iteration-in-murex" tabindex="-1"><a class="header-anchor" href="#structured-iteration-in-murex" aria-hidden="true">#</a> Structured iteration in Murex</h2>
<p>The following examples will use JSON as the input format, however Murex natively supports other structured data formats too.</p>
<h3 id="the-foreach-builtin" tabindex="-1"><a class="header-anchor" href="#the-foreach-builtin" aria-hidden="true">#</a> The <code v-pre>foreach</code> builtin</h3>
<p>Lets say you have an array that looks something like the following:</p>
<pre><code>» %[January..December]
[
    &quot;January&quot;,
    &quot;February&quot;,
    &quot;March&quot;,
    &quot;April&quot;,
    &quot;May&quot;,
    &quot;June&quot;,
    &quot;July&quot;,
    &quot;August&quot;,
    &quot;September&quot;,
    &quot;October&quot;,
    &quot;November&quot;,
    &quot;December&quot;
]
</code></pre>
<p>And lets say you wanted to only return months that ended in <strong>ber</strong> (excuse the following video but I find this track be B.E.R. to be a particularly effective earworm)</p>
<iframe width="560" height="315" src="https://www.youtube-nocookie.com/embed/MKtW-k8za7I?controls=0" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>
<p>Sure you could <code v-pre>grep</code> that in Bash but what you're left with isn't JSON. And what if that JSON is minified?</p>
<pre><code>[&quot;January&quot;,&quot;February&quot;,&quot;March&quot;,&quot;April&quot;,&quot;May&quot;,&quot;June&quot;,&quot;July&quot;,&quot;August&quot;,&quot;September&quot;,&quot;October&quot;,&quot;November&quot;,&quot;December&quot;]
</code></pre>
<p>Now reading each line isn't going to work.</p>
<p>Murex doesn't just treat files as byte streams, it passes type annotations too. And it uses those annotations to dynamically alter how to read files. So you could grep only <strong>ber</strong> from minified JSON with a simple <code v-pre>foreach { | grep &quot;ber&quot; }</code>:</p>
<pre><code>$months = %[January..December]
$months | foreach {
    | grep ber
}
</code></pre>
<p>The best thing is <code v-pre>foreach</code> will understand arrays from JSON, YAML, S-Expressions, CSV, and others... as well as regular terminal output. So it is one syntax to learn regardless of the input file format.</p>
<h3 id="the-formap-builtin" tabindex="-1"><a class="header-anchor" href="#the-formap-builtin" aria-hidden="true">#</a> The <code v-pre>formap</code> builtin</h3>
<p>Arrays are easy though. What about maps? Or dictionaries, objects, hashes... as other languages might refer to them. Key value pairs are always going to be harder to parse because even with file formats like YAML, arrays are line delimited where as splitting key value pairs requires extra grokking.</p>
<p>Murex also has an iterator for such constructs: <code v-pre>formap</code>:</p>
<pre><code>» echo '{&quot;a&quot;:1,&quot;b&quot;:2,&quot;c&quot;:3}' | :json: formap key value { echo &quot;key=$key, value=$value&quot; }
key=a, value=1
key=b, value=2
key=c, value=3
</code></pre>
<p>Here we are using <code v-pre>:json:</code> to cast / annotate the STDIN stream for <code v-pre>formap</code>.</p>
<h3 id="lambdas" tabindex="-1"><a class="header-anchor" href="#lambdas" aria-hidden="true">#</a> Lambdas</h3>
<p>Lambdas were introduced in version 4.0 of Murex. Lambdas offer shortcuts around common matching problems with structured data. For example:</p>
<pre><code>» $example = %{a:1, b:2, c:3}
» @example[{$.val == 2}]
{
    &quot;b&quot;: 2
}
</code></pre>
<p>The way this particular lambda works is that for each element in <code v-pre>$example</code> the code block <code v-pre>{$.val == 2}</code> runs with the <code v-pre>$.</code> variable is assigned with a structure that looks like this:</p>
<pre><code># first iteration
{
    &quot;key&quot;: &quot;a&quot;,
    &quot;val&quot;: 1
}
# second iteration
{
    &quot;key&quot;: &quot;b&quot;,
    &quot;val&quot;: 2
}
# third iteration
{
    &quot;key&quot;: &quot;c&quot;,
    &quot;val&quot;: 3
}
</code></pre>
<p>Thus <code v-pre>$.val</code> (eg <strong>2</strong>) would be compared to the number <strong>2</strong>. If the result is true, that element is output. If the result is false then that element is ignored.</p>
<h2 id="conclusion" tabindex="-1"><a class="header-anchor" href="#conclusion" aria-hidden="true">#</a> Conclusion</h2>
<p>There are a multitude of ways to iterate through structured data on Linux and UNIX systems.</p>
<hr>
<p>Published: 06.05.2021 at 08:24</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/parser/create-array.html">Create array (<code v-pre>%[]</code>) constructor</RouterLink>:
Quickly generate arrays</li>
<li><RouterLink to="/commands/cast.html"><code v-pre>cast</code></RouterLink>:
Alters the data type of the previous function without altering it's output</li>
<li><RouterLink to="/commands/foreach.html"><code v-pre>foreach</code></RouterLink>:
Iterate through an array</li>
<li><RouterLink to="/commands/formap.html"><code v-pre>formap</code></RouterLink>:
Iterate through a map or other collection of data</li>
<li><RouterLink to="/commands/mkarray/">mkarray</RouterLink>:</li>
</ul>
</div></template>


