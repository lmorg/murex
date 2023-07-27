<template><div><h1 id="reading-lists-from-the-command-line-blog" tabindex="-1"><a class="header-anchor" href="#reading-lists-from-the-command-line-blog" aria-hidden="true">#</a> Reading Lists From The Command Line - Blog</h1>
<blockquote>
<p>How hard can it be to read a list of data from the command line? If your list is line delimited then it should be easy. However what if your list is a JSON array? This post will explore how to work with lists in a different command line environments.</p>
</blockquote>
<h2 id="preface" tabindex="-1"><a class="header-anchor" href="#preface" aria-hidden="true">#</a> Preface</h2>
<p>A common problem we resort to shell scripting for is iterating through lists. This was easy in the days of old when most data was <code v-pre>\n</code> (new line) delimited but these days structured data is common place with formats like JSON, YAML, TOML, XML and even S-Expressions appearing commonly throughout developer and DevOps tooling.</p>
<p>So lets explore a few techniques for iterating through lists.</p>
<h2 id="reading-lines-in-bash-and-similar-shells" tabindex="-1"><a class="header-anchor" href="#reading-lines-in-bash-and-similar-shells" aria-hidden="true">#</a> Reading lines in Bash and similar shells</h2>
<p>Bash shell is a popular command-line interface for Unix and Linux operating systems. One of its many useful features is the ability to read files line by line. This can be helpful for processing large files or performing repetitive tasks on a file's contents. The most basic way to read a file line by line is to use a while loop with the <code v-pre>read</code> command:</p>
<pre><code>while read line; do
    echo $line
done &lt; file.txt
</code></pre>
<p>In this example, the <code v-pre>while</code> loop reads each line of the <code v-pre>file.txt</code> file and stores it in the <code v-pre>$line</code> variable. The <code v-pre>echo</code> command then prints the contents of the <code v-pre>$line</code> variable to the console. The <code v-pre>&lt;</code> symbol tells Bash to redirect the contents of the file into the loop.</p>
<p>The <code v-pre>read</code> command is what actually reads each line of the file. By default, it reads one line at a time and stores it in the variable specified. You can also use the <code v-pre>-r</code> option with the <code v-pre>read</code> command to disable backslash interpretation, which can be useful when dealing with files that contain backslashes.</p>
<p>Another useful feature of Bash is the ability to perform operations on each line of a file before processing it. For example, you can use <code v-pre>sed</code> to replace text within each line of a file:</p>
<pre><code>while read line; do
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
<p>So how do you read lists from objects in, for example, JSON? In Bash, this isn't so easy. You need to rely on third party tools like <code v-pre>jq</code>. However you do have the benefit of compatibility with all of the older core utilities, like <code v-pre>sed</code>, that have become muscle memory by now. This does also come with its own drawbacks as well, which I'll explore in the following section.</p>
<h2 id="iteration-in-bash-via-jq" tabindex="-1"><a class="header-anchor" href="#iteration-in-bash-via-jq" aria-hidden="true">#</a> Iteration in Bash via <code v-pre>jq</code></h2>
<p><code v-pre>jq</code> is a fantastic tool that has become a staple of many a CI/CD pipeline however it is not part of most operating systems base platform, so it would need to be installed separately. This also creates additional complications whereby you end up having a language within a language -- like running <code v-pre>awk</code> or <code v-pre>sed</code> inside Bash, you're now introducing <code v-pre>jq</code> too. Thus its syntax isn't always the easiest to grok when delving deep into nested JSON with conditionals and such like compared with shells that offer first party tools for working with objects. We can delve deeper into the power of <code v-pre>jq</code> in another article but for now we are going to keep things intentionally simple:</p>
<p>Lets create a JSON array:</p>
<pre><code>json='[&quot;Monday&quot;,&quot;Tuesday&quot;,&quot;Wednesday&quot;,&quot;Thursday&quot;,&quot;Friday&quot;,&quot;Saturday&quot;,&quot;Sunday&quot;]'
</code></pre>
<p>Straight away you should be able to see that Bash, with its reliance on whitespace delimitations, couldn't natively parse this. So now lets run it through <code v-pre>jq</code>:</p>
<pre><code>$ echo $json | jq -r '.[]' | while read -r day do; echo &quot;Happy $day&quot;; done
Happy Monday
Happy Tuesday
Happy Wednesday
Happy Thursday
Happy Friday
Happy Saturday
Happy Sunday
</code></pre>
<p>What's happening here is the <code v-pre>jq</code> tool is converting our JSON array into a <code v-pre>\n</code> delimited list. And from there, we can use <code v-pre>while</code> and <code v-pre>read</code> just like we did in our first example at the start of this article.</p>
<p>The <code v-pre>-r</code> flag tells <code v-pre>jq</code> to strip quotation marks around the values. Without <code v-pre>-r</code> you'd see <code v-pre>Happy &quot;Monday&quot;'</code> and so on and so forth.</p>
<p><code v-pre>.[]</code> is <code v-pre>jq</code> syntax for &quot;all elements (<code v-pre>[]</code>) in the root object space (<code v-pre>.</code>).</p>
<h2 id="iteration-in-murex-via-foreach" tabindex="-1"><a class="header-anchor" href="#iteration-in-murex-via-foreach" aria-hidden="true">#</a> Iteration in Murex via <code v-pre>foreach</code></h2>
<p>Murex doesn't just treat files as byte streams, it passes type annotations too. And it uses those annotations to dynamically alter how to read files. The following examples will also use JSON as the input format, however Murex natively supports other structured data formats too, like YAML, CSV and S-Expressions.</p>
<p>Lets use the same JSON array as we did earlier, except use one of Murex's features to generate arrays programmatically:</p>
<pre><code>» %[Monday..Sunday]
[
    &quot;Monday&quot;,
    &quot;Tuesday&quot;,
    &quot;Wednesday&quot;,
    &quot;Thursday&quot;,
    &quot;Friday&quot;,
    &quot;Saturday&quot;,
    &quot;Sunday&quot;
]
</code></pre>
<p>The <code v-pre>jq</code> example rewritten in Murex would look like the following:</p>
<pre><code>%[Monday..Sunday] | foreach day { out &quot;Happy $day&quot; }
</code></pre>
<p>What's happening here is <code v-pre>%[...]</code> creates the JSON array (as described above) and then the <code v-pre>foreach</code> builtin iterates through the array and assigns that element to a variable named <code v-pre>day</code>.</p>
<blockquote>
<p><code v-pre>out</code> in Murex is the equivalent of <code v-pre>echo</code> in Bash. In fact you can still use <code v-pre>echo</code> in Murex albeit that is just aliased to <code v-pre>out</code>.</p>
</blockquote>
<p>It is also worth noting that since Murex version 3.1 lambdas have been available, allowing you to write code that looks a like this:</p>
<pre><code>$json[{out &quot;Hello $.&quot;}]
</code></pre>
<p>But more on that in a different article.</p>
<h2 id="reading-json-arrays-in-powershell" tabindex="-1"><a class="header-anchor" href="#reading-json-arrays-in-powershell" aria-hidden="true">#</a> Reading JSON arrays in PowerShell</h2>
<p>Microsoft PowerShell is a typed shell, like Murex, which was originally built for Windows but has since been ported to macOS and Linux too. Where PowerShell differs is that rather than using byte streams with type annotations, PowerShell passes .NET objects. Thus you'll see a little more boilerplate code in PowerShell where you need to explicitly convert types -- whereas Murex can get away with implicit definitions.</p>
<pre><code>$jsonString = '[&quot;Monday&quot;,&quot;Tuesday&quot;,&quot;Wednesday&quot;,&quot;Thursday&quot;,&quot;Friday&quot;,&quot;Saturday&quot;,&quot;Sunday&quot;]'
$jsonObject = ConvertFrom-Json $jsonString
foreach ($day in $jsonObject) { Write-Host &quot;Hello $day&quot; }
</code></pre>
<p>The first line is just creating a JSON string and allocating that to <code v-pre>$jsonString</code>. We can largely ignore that as it is the same as we've seen in Bash already. The second line is more interesting as that is where the type conversion happens. <code v-pre>ConvertFrom-Json</code> does exactly as it describes -- PowerShell is generally pretty good at having descriptive names for commands,</p>
<p>From there on it looks relatively similar to Murex syntax: <code v-pre>foreach</code> being the statement name, followed by the variable name.</p>
<h2 id="conclusion" tabindex="-1"><a class="header-anchor" href="#conclusion" aria-hidden="true">#</a> Conclusion</h2>
<p>Iterating over a JSON array from the command line can be done in various ways using different shells. PowerShell, <code v-pre>jq</code>, and Murex offer their unique approaches and syntaxes, making it easy to work with JSON data in different environments. Whether you're a Windows user who prefers PowerShell or a Linux user who prefers Bash or Murex, there are many options available to suit your needs. Regardless of the shell you choose, mastering the art of iterating over JSON arrays can greatly enhance your command-line skills and help you work more efficiently with JSON data.</p>
<hr>
<p>Published: 22.04.2023 at 11:43</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/parser/create-array.html">Create array (<code v-pre>%[]</code>) constructor</RouterLink>:
Quickly generate arrays</li>
<li><RouterLink to="/commands/a.html"><code v-pre>a</code> (mkarray)</RouterLink>:
A sophisticated yet simple way to build an array or list</li>
<li><RouterLink to="/commands/cast.html"><code v-pre>cast</code></RouterLink>:
Alters the data type of the previous function without altering it's output</li>
<li><RouterLink to="/commands/foreach.html"><code v-pre>foreach</code></RouterLink>:
Iterate through an array</li>
<li><RouterLink to="/commands/formap.html"><code v-pre>formap</code></RouterLink>:
Iterate through a map or other collection of data</li>
</ul>
</div></template>


