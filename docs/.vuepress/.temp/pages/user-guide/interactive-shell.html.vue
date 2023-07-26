<template><div><h1 id="interactive-shell" tabindex="-1"><a class="header-anchor" href="#interactive-shell" aria-hidden="true">#</a> Interactive Shell</h1>
<blockquote>
<p>What's different about Murex's interactive shell?</p>
</blockquote>
<h2 id="overview" tabindex="-1"><a class="header-anchor" href="#overview" aria-hidden="true">#</a> Overview</h2>
<p>Aside from Murex being carefully designed with scripting in mind, the
interactive shell itself is also built around productivity. To achieve this
we wrote our own readline library. Below is an example of that library in use:</p>
<p><a href="https://asciinema.org/a/232714" target="_blank" rel="noopener noreferrer"><img src="https://asciinema.org/a/232714.svg" alt="asciicast" loading="lazy"><ExternalLinkIcon/></a></p>
<p>The above demo includes the following features of Murex's bespoke readline
library:</p>
<ul>
<li>hint text - blue status text below the prompt (the colour is configurable)</li>
<li>syntax highlighting (albeit there isn’t much syntax to highlight in the
example). This can also be turned off if your preference is to have colours
disabled</li>
<li>tab-completion in gridded mode (seen when typing <code v-pre>cd</code>)</li>
<li>tab-completion in list view (seen when selecting a process name to <code v-pre>kill</code>
where the process ID was substituted when selected)</li>
<li>searching through the tab-completion suggestions (seen in both <code v-pre>cd</code> and
<code v-pre>kill</code> - enabled by pressing <code v-pre>[ctrl]</code>+<code v-pre>[f]</code>)</li>
<li>line editing using $EDITOR (<code v-pre>vi</code> in the example - enabled by pressing <code v-pre>[esc]</code>
followed by <code v-pre>[v]</code>)</li>
<li>readline’s warning before pasting multiple lines of data into the buffer and
the preview option that’s available as part of the aforementioned warning</li>
<li>and VIM keys (enabled by pressing <code v-pre>[esc]</code>)</li>
</ul>
<h2 id="readline" tabindex="-1"><a class="header-anchor" href="#readline" aria-hidden="true">#</a> readline</h2>
<p>Murex uses a custom <code v-pre>readline</code> library to enable support for new features on
in addition to the existing uses you'd normally expect from a shell. It is
because of this Murex provides one of the best user experiences of any of the
shells available today.</p>
<h3 id="hotkeys" tabindex="-1"><a class="header-anchor" href="#hotkeys" aria-hidden="true">#</a> Hotkeys</h3>
<p>A full breakdown of supported hotkeys is available at <RouterLink to="/user-guide/terminal-keys.html">terminal-keys.md</RouterLink>.</p>
<h3 id="autocompletion" tabindex="-1"><a class="header-anchor" href="#autocompletion" aria-hidden="true">#</a> Autocompletion</h3>
<p>Autocompletion happen when you press <code v-pre>[tab]</code> and will differ slightly depending
on what is defined in <code v-pre>autocomplete</code> and whether you use the traditional
<RouterLink to="/parser/pipe-posix.html">POSIX pipe token</RouterLink>, <code v-pre>|</code>, or the <RouterLink to="/parser/pipe-arrow.html">arrow pipe</RouterLink>,
<code v-pre>-&gt;</code>.</p>
<p>The <code v-pre>|</code> token will behave much like any other shell however <code v-pre>-&gt;</code> will offer
suggestions with matching data types (as seen in <code v-pre>runtime --methods</code>). This is
a way of helping highlight commands that naturally follow after another in a
pipeline. Which is particularly important in Murex as it introduces data
types and dozens of new builtins specifically for working with data structures
in an intelligent and readable yet succinct way.</p>
<p>You can add your own commands and functions to Murex as methods by defining
them with <code v-pre>method</code>. For example if we were to add <code v-pre>jq</code> as a method:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>method: define jq {
    "Stdin":  "json",
    "Stdout": "@Any"
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="syntax-completion" tabindex="-1"><a class="header-anchor" href="#syntax-completion" aria-hidden="true">#</a> Syntax Completion</h3>
<p>Like with most IDEs, Murex will auto close brackets et al.</p>
<p><a href="https://asciinema.org/a/408029" target="_blank" rel="noopener noreferrer"><img src="https://asciinema.org/a/408029.svg" alt="asciicast" loading="lazy"><ExternalLinkIcon/></a></p>
<h3 id="syntax-highlighting" tabindex="-1"><a class="header-anchor" href="#syntax-highlighting" aria-hidden="true">#</a> Syntax Highlighting</h3>
<p>Pipelines in the interactive terminal are syntax highlighted. This is similar
to what one expects from an IDE.</p>
<p>Syntax highlighting can be disabled by running:</p>
<pre><code>config: set shell syntax-highlighting off
</code></pre>
<h3 id="spellchecker" tabindex="-1"><a class="header-anchor" href="#spellchecker" aria-hidden="true">#</a> Spellchecker</h3>
<p>Murex supports inline spellchecking, where errors are underlined. For example</p>
<p><a href="https://asciinema.org/a/408024" target="_blank" rel="noopener noreferrer"><img src="https://asciinema.org/a/408024.svg" alt="asciicast" loading="lazy"><ExternalLinkIcon/></a></p>
<p>This might require some manual steps to enable, please see the <RouterLink to="/user-guide/spellcheck.html">spellcheck user guide</RouterLink>
for more details.</p>
<h3 id="hint-text" tabindex="-1"><a class="header-anchor" href="#hint-text" aria-hidden="true">#</a> Hint Text</h3>
<p>The <strong>hint text</strong> is a (typically) blue status line that appears directly below
your prompt. The idea behind the <strong>hint text</strong> is to provide clues to you as
type instructions into the prompt; but without adding distractions. It is there
to be used if you want it while keeping out of the way when you don't want it.</p>
<h4 id="configuring-hint-text-colour" tabindex="-1"><a class="header-anchor" href="#configuring-hint-text-colour" aria-hidden="true">#</a> Configuring Hint Text Colour</h4>
<p>By default the <strong>hint text</strong> will appear blue. This is also customizable:</p>
<pre><code>» config get shell hint-text-formatting
{BLUE}
</code></pre>
<p>The formatting config takes a string and supports <RouterLink to="/user-guide/ansi.html">ANSI constants</RouterLink>.</p>
<p>It is also worth noting that if colour is disabled then the <strong>hint text</strong> will
not be coloured even if <strong>hint-text-formatting</strong> includes colour codes:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» config: set shell color false
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>(please note that <strong>syntax highlighting</strong> is unaffected by the above config)</p>
<h3 id="custom-hint-text-statuses" tabindex="-1"><a class="header-anchor" href="#custom-hint-text-statuses" aria-hidden="true">#</a> Custom Hint Text Statuses</h3>
<p>There is a lot of behavior hardcoded into Murex like displaying the full path
to executables and the values of variables. However if there is no status to be
displayed then Murex can fallback to a default <strong>hint text</strong> status. This
default is a user defined function. At time of writing this document the author
has the following function defined:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>    config: set shell hint-text-func {
        trypipe &lt;!null> {
            git status --porcelain -b -> set gitstatus
            $gitstatus -> head -n1 -> regexp 's/^## //' -> regexp 's/\.\.\./ => /'
        }
        catch {
            out "Not a git repository."
        }
    }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>...which produces a colorized status that looks something like the following:</p>
<pre><code>develop =&gt; origin/develop
</code></pre>
<h4 id="disabling-hint-text" tabindex="-1"><a class="header-anchor" href="#disabling-hint-text" aria-hidden="true">#</a> Disabling Hint Text</h4>
<p>It is enabled by default but can be disabled if you prefer a more minimal
prompt:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» config: set shell hint-text-enabled false
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/ansi.html">ANSI Constants</RouterLink>:
Infixed constants that return ANSI escape sequences</li>
<li><RouterLink to="/parser/pipe-arrow.html">Arrow Pipe (<code v-pre>-&gt;</code>) Token</RouterLink>:
Pipes STDOUT from the left hand command to STDIN of the right hand command</li>
<li><RouterLink to="/user-guide/code-block.html">Code Block Parsing</RouterLink>:
Overview of how code blocks are parsed</li>
<li><RouterLink to="/parser/curly-brace.html">Curly Brace (<code v-pre>{</code>, <code v-pre>}</code>) Tokens</RouterLink>:
Initiates or terminates a code block</li>
<li><RouterLink to="/parser/pipe-posix.html">POSIX Pipe (<code v-pre>|</code>) Token</RouterLink>:
Pipes STDOUT from the left hand command to STDIN of the right hand command</li>
<li><RouterLink to="/user-guide/spellcheck.html">Spellcheck</RouterLink>:
How to enable inline spellchecking</li>
<li><RouterLink to="/user-guide/terminal-keys.html">Terminal Hotkeys</RouterLink>:
A list of all the terminal hotkeys and their uses</li>
<li><RouterLink to="/commands/autocomplete.html"><code v-pre>autocomplete</code></RouterLink>:
Set definitions for tab-completion in the command line</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/method.html"><code v-pre>method</code></RouterLink>:
Define a methods supported data-types</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>runtime</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
</ul>
</div></template>


