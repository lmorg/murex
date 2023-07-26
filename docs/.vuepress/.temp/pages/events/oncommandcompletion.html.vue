<template><div><h1 id="oncommandcompletion" tabindex="-1"><a class="header-anchor" href="#oncommandcompletion" aria-hidden="true">#</a> <code v-pre>onCommandCompletion</code></h1>
<blockquote>
<p>Trigger an event upon a command's completion</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>onCommandCompletion</code> events are triggered after a command has finished
executing in the interactive terminal.</p>
<p>Background processes or commands ran from inside aliases, functions, nested
blocks or from shell scripts cannot trigger this event. This is to protect
against accidental race conditions, infinite loops and breaking expected
behaviour / the portability of Murex scripts. On those processes directly ran
from the prompt can trigger this event.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>event: onCommandCompletion name=command { code block }

!event: onCommandCompletion name
</code></pre>
<h2 id="valid-interrupts" tabindex="-1"><a class="header-anchor" href="#valid-interrupts" aria-hidden="true">#</a> Valid Interrupts</h2>
<ul>
<li><code v-pre>&lt;command&gt;</code>
Name of command that triggers this event</li>
</ul>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p><strong>Read STDERR:</strong></p>
<p>In this example we check the output from <code v-pre>pacman</code>, which is ArchLinux's package
management tool, to see if you have accidentally ran it as a non-root user. If
the STDERR contains a message saying you are no root, then this event function
will re-run <code v-pre>pacman</code> with <code v-pre>sudo</code>.</p>
<pre><code>event: onCommandCompletion sudo-pacman=pacman {
    `&lt;stdin&gt;` -&gt; set event
    read-named-pipe: $event.Interrupt.Stderr \
    -&gt; regexp 'm/error: you cannot perform this operation unless you are root/' \
    -&gt; if {
          sudo pacman @event.Interrupt.Parameters
       }
}
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="payload" tabindex="-1"><a class="header-anchor" href="#payload" aria-hidden="true">#</a> Payload</h3>
<p>The following payload is passed to the function via STDIN:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>    {
        "Name": "",
        "Interrupt": {
            "Command": "",
            "Parameters": [],
            "Stdout": "",
            "Stderr": "",
            "ExitNum": 0
        }
    }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h4 id="name" tabindex="-1"><a class="header-anchor" href="#name" aria-hidden="true">#</a> Name</h4>
<p>This is the name you specified when defining the event.</p>
<h4 id="command" tabindex="-1"><a class="header-anchor" href="#command" aria-hidden="true">#</a> Command</h4>
<p>Name of command executed prior to this event being triggered</p>
<h4 id="operation" tabindex="-1"><a class="header-anchor" href="#operation" aria-hidden="true">#</a> Operation</h4>
<p>The commandline parameters of the aforementioned command</p>
<h4 id="stdout" tabindex="-1"><a class="header-anchor" href="#stdout" aria-hidden="true">#</a> Stdout</h4>
<p>This is the name of the Murex named pipe which contains a copy of the STDOUT
from the command which executed prior to this event.</p>
<p>You can read this with <code v-pre>read-named-pipe</code>. eg</p>
<pre><code>» `&lt;stdin&gt;` -&gt; set: event
» read-named-pipe: $event.Interrupt.Stdout -&gt; ...
</code></pre>
<h4 id="stderr" tabindex="-1"><a class="header-anchor" href="#stderr" aria-hidden="true">#</a> Stderr</h4>
<p>This is the name of the Murex named pipe which contains a copy of the STDERR
from the command which executed prior to this event.</p>
<p>You can read this with <code v-pre>read-named-pipe</code>. eg</p>
<pre><code>» `&lt;stdin&gt;` -&gt; set: event
» read-named-pipe: $event.Interrupt.Stderr -&gt; ...
</code></pre>
<h4 id="exitnum" tabindex="-1"><a class="header-anchor" href="#exitnum" aria-hidden="true">#</a> ExitNum</h4>
<p>This is the exit number returned from the executed command.</p>
<h3 id="stdout-1" tabindex="-1"><a class="header-anchor" href="#stdout-1" aria-hidden="true">#</a> Stdout</h3>
<p>Stdout is written to the terminal. So this can be used to provide multiple
additional lines to the prompt since readline only supports one line for the
prompt itself and three extra lines for the hint text.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/namedpipes.html">Murex Named Pipes</RouterLink>:
A detailed breakdown of named pipes in Murex</li>
<li><RouterLink to="/commands/stdin.html"><code v-pre>&lt;stdin&gt;</code> </RouterLink>:
Read the STDIN belonging to the parent code block</li>
<li><RouterLink to="/commands/alias.html"><code v-pre>alias</code></RouterLink>:
Create an alias for a command</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/event.html"><code v-pre>event</code></RouterLink>:
Event driven programming for shell scripts</li>
<li><RouterLink to="/commands/function.html"><code v-pre>function</code></RouterLink>:
Define a function block</li>
<li><RouterLink to="/commands/if.html"><code v-pre>if</code></RouterLink>:
Conditional statement to execute different blocks of code depending on the result of the condition</li>
<li><RouterLink to="/events/onprompt.html"><code v-pre>onPrompt</code></RouterLink>:
Events triggered by changes in state of the interactive shell</li>
<li><RouterLink to="/commands/regexp.html"><code v-pre>regexp</code></RouterLink>:
Regexp tools for arrays / lists of strings</li>
<li><RouterLink to="/commands/namedpipe.html">read-named-pipe</RouterLink>:
Reads from a Murex named pipe</li>
</ul>
</div></template>


