<template><div><h1 id="spellcheck" tabindex="-1"><a class="header-anchor" href="#spellcheck" aria-hidden="true">#</a> Spellcheck</h1>
<blockquote>
<p>How to enable inline spellchecking</p>
</blockquote>
<p>Murex supports inline spellchecking, where errors are underlined. For example</p>
<p><a href="https://asciinema.org/a/408024" target="_blank" rel="noopener noreferrer"><img src="https://asciinema.org/a/408024.svg" alt="asciicast" loading="lazy"><ExternalLinkIcon/></a></p>
<p>However to use this there needs to be a few satisfied prerequisites, not all of
which will be enabled by default:</p>
<h2 id="cli-spellchecker-3rd-party-software" tabindex="-1"><a class="header-anchor" href="#cli-spellchecker-3rd-party-software" aria-hidden="true">#</a> CLI Spellchecker (3rd Party Software)</h2>
<p>A CLI spellchecker needs to be installed. The recommendation is <code v-pre>aspell</code>. This
might already be installed by default with your OS or has been included as a
dependency with another application. You can check if <code v-pre>aspell</code> is installed by
running the following:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>which: aspell
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>If that returns no data, then you will need to install <code v-pre>aspell</code> yourself.
Please consult your OS docs for how to install software.</p>
<p>For help debugging issues with <code v-pre>aspell</code>, please see the last section in this
document.</p>
<h2 id="murex-config" tabindex="-1"><a class="header-anchor" href="#murex-config" aria-hidden="true">#</a> Murex Config</h2>
<h3 id="ansi-escape-sequences" tabindex="-1"><a class="header-anchor" href="#ansi-escape-sequences" aria-hidden="true">#</a> ANSI Escape Sequences</h3>
<p>ANSI escape sequences need to be enabled (which they are by default). This
option is found in <code v-pre>config</code> under <strong>shell</strong>, <strong>color</strong>.</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>config: set shell color true
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><h3 id="spellcheck-enable" tabindex="-1"><a class="header-anchor" href="#spellcheck-enable" aria-hidden="true">#</a> Spellcheck Enable</h3>
<p>Spellcheck needs to be enabled. This option can be found in <code v-pre>config</code> under
<strong>shell</strong>, <strong>spellcheck-enabled</strong>.</p>
<p>To enable this run:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>config: set shell spellcheck-enabled true
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>...or add the above line to your Murex profile, <code v-pre>~/.murex_profile</code> to make
the change persistent.</p>
<blockquote>
<p>Please note that this option will automatically be enabled if <code v-pre>aspell</code> is
installed.</p>
</blockquote>
<h3 id="spellcheck-murex-code" tabindex="-1"><a class="header-anchor" href="#spellcheck-murex-code" aria-hidden="true">#</a> Spellcheck Murex Code</h3>
<p>This shouldn't need tweaking if you're running <code v-pre>aspell</code> but other spellcheckers
will require updated code. The default will look something like this:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» config: get shell spellcheck-func
{ -> aspell list }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><p>The default should be good enough for most people but should you want to run an
alternative spellchecker then follow the instructions in the next section:</p>
<h2 id="how-to-write-your-own-spellcheck-func" tabindex="-1"><a class="header-anchor" href="#how-to-write-your-own-spellcheck-func" aria-hidden="true">#</a> How To Write Your Own <code v-pre>spellcheck-func</code></h2>
<p>You might legitimately want to run a different spellchecker and if so you'll
need to write your own <strong>spellcheck-func</strong>. Fortunately this is simple:</p>
<p>The function reads the command line from STDIN, if the spellchecker reads lines
from parameters rather than STDIN you'll need to write something equivalent to
the following</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>{
    # This is a theoretical example. It will not work generically.
    -> set: line
    newspellchecker --check "$line"
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>The output of the function must me an array containing the misspelt words only.
That array can be JSON just as long as you have set STDOUT's data type to
<code v-pre>json</code>. Similarly, other supported Murex data types can be used too. However
in general you might just want to go with a misspelling per line as it's pretty
POSIX friendly and thus most spellcheckers are likely to support it. eg</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» out "a list of misspelt words: qwert fuubar madeupword" -> aspell: list
qwert
fuubar
madeupword
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="user-dictionary" tabindex="-1"><a class="header-anchor" href="#user-dictionary" aria-hidden="true">#</a> User Dictionary</h2>
<p>Murex has it's own user dictionary, which is held as a JSON array:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» config: get shell spellcheck-user-dictionary
["murex"]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><p>You can add words to a user dictionary via:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» config: eval shell spellcheck-user-dictionary { -> append "myword" }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>or</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» config: eval shell spellcheck-user-dictionary { -> alter: --merge / (["word1", "word2", "word3"]) }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><blockquote>
<p>Don't forget to record these in your Murex profile, <code v-pre>~/.murex_profile</code> to
make the changes persistent.</p>
</blockquote>
<h3 id="ignored-by-default" tabindex="-1"><a class="header-anchor" href="#ignored-by-default" aria-hidden="true">#</a> Ignored By Default</h3>
<p>Sometimes commands are not valid words in ones native language. Thus any words
that fall into the following categories are ignored by default:</p>
<ul>
<li>words that are also the names of commands found in <code v-pre>$PATH</code></li>
<li>words that are the names of Murex functions (defined via <code v-pre>function</code>)</li>
<li>words that are builtins (eg <code v-pre>config</code> and <code v-pre>jsplit</code>)</li>
<li>any global aliases</li>
<li>also any words that are also the names of global variables</li>
</ul>
<h2 id="common-problems-with-aspell" tabindex="-1"><a class="header-anchor" href="#common-problems-with-aspell" aria-hidden="true">#</a> Common Problems With <code v-pre>aspell</code></h2>
<h3 id="error-no-word-lists-can-be-found-for-the-language-en-nz" tabindex="-1"><a class="header-anchor" href="#error-no-word-lists-can-be-found-for-the-language-en-nz" aria-hidden="true">#</a> <code v-pre>Error: No word lists can be found for the language &quot;en_NZ&quot;.</code></h3>
<p>The <code v-pre>en_NZ</code> portion of the error will differ depending on your language.</p>
<p>If this error arises then it means <code v-pre>aspell</code> is installed but it doesn't have
the dictionary for your language. This is an easy fix in most OSs. For example
in Ubuntu:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>apt install aspell-en
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>(you may need to change <code v-pre>-en</code> with your specific language code)</p>
<p>Please consult your operating systems documentation for how to install software
and what the package names are for <code v-pre>aspell</code> and its corresponding dictionaries.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/ansi.html">ANSI Constants</RouterLink>:
Infixed constants that return ANSI escape sequences</li>
<li><RouterLink to="/user-guide/code-block.html">Code Block Parsing</RouterLink>:
Overview of how code blocks are parsed</li>
<li><RouterLink to="/parser/curly-brace.html">Curly Brace (<code v-pre>{</code>, <code v-pre>}</code>) Tokens</RouterLink>:
Initiates or terminates a code block</li>
<li><RouterLink to="/user-guide/profile.html">Murex Profile Files</RouterLink>:
A breakdown of the different files loaded on start up</li>
<li><RouterLink to="/user-guide/interactive-shell.html">Murex's Interactive Shell</RouterLink>:
What's different about Murex's interactive shell?</li>
<li><RouterLink to="/commands/alter.html"><code v-pre>alter</code></RouterLink>:
Change a value within a structured data-type and pass that change along the pipeline without altering the original source input</li>
<li><RouterLink to="/commands/append.html"><code v-pre>append</code></RouterLink>:
Add data to the end of an array</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/types/json.html"><code v-pre>json</code> </RouterLink>:
JavaScript Object Notation (JSON)</li>
<li><RouterLink to="/commands/jsplit.html"><code v-pre>jsplit</code> </RouterLink>:
Splits STDIN into a JSON array based on a regex parameter</li>
<li><RouterLink to="/commands/set.html"><code v-pre>set</code></RouterLink>:
Define a local variable and set it's value</li>
</ul>
</div></template>


