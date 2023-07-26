<template><div><h1 id="onprompt" tabindex="-1"><a class="header-anchor" href="#onprompt" aria-hidden="true">#</a> <code v-pre>onPrompt</code></h1>
<blockquote>
<p>Events triggered by changes in state of the interactive shell</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>onPrompt</code> events are triggered by changes in state of the interactive shell
(often referred to as <em>readline</em>). Those states are defined in the interrupts
section below.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>event: onPrompt name=[before|after|abort|eof] { code block }

!event: onPrompt [before_|after_|abort_|eof_]name
</code></pre>
<h2 id="valid-interrupts" tabindex="-1"><a class="header-anchor" href="#valid-interrupts" aria-hidden="true">#</a> Valid Interrupts</h2>
<ul>
<li><code v-pre>abort</code>
Triggered if <code v-pre>ctrl</code>+<code v-pre>c</code> pressed while in the interactive prompt</li>
<li><code v-pre>after</code>
Triggered after user has written a command into the interactive prompt and then hit `enter</li>
<li><code v-pre>before</code>
Triggered before readline displays the interactive prompt</li>
<li><code v-pre>eof</code>
Triggered if <code v-pre>ctrl</code>+<code v-pre>d</code> pressed while in the interactive prompt</li>
</ul>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p><strong>Interrupt 'before':</strong></p>
<pre><code>event: onPrompt example=before {
    out: &quot;This will appear before your command prompt&quot;
}
</code></pre>
<p><strong>Interrupt 'after':</strong></p>
<pre><code>event: onPrompt example=after {
    out: &quot;This will appear after you've hit [enter] on your command prompt&quot;
    out: &quot;...but before the command executes&quot;
}
</code></pre>
<p><strong>Echo the command line:</strong></p>
<pre><code>» event: onPrompt echo=after { -&gt; set event; out $event.Interrupt.CmdLine }
» echo hello world
echo hello world
hello world
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="payload" tabindex="-1"><a class="header-anchor" href="#payload" aria-hidden="true">#</a> Payload</h3>
<p>The following payload is passed to the function via STDIN:</p>
<pre><code>{
    &quot;Name&quot;: &quot;&quot;,
    &quot;Interrupt&quot;: {
        &quot;Name&quot;: &quot;&quot;,
        &quot;Operation&quot;: &quot;&quot;,
        &quot;CmdLine&quot;: &quot;&quot;
    }
}
</code></pre>
<h4 id="name" tabindex="-1"><a class="header-anchor" href="#name" aria-hidden="true">#</a> Name</h4>
<p>This is the <strong>namespaced</strong> name -- ie the name and operation.</p>
<h4 id="interrupt-name" tabindex="-1"><a class="header-anchor" href="#interrupt-name" aria-hidden="true">#</a> Interrupt/Name</h4>
<p>This is the name you specified when defining the event.</p>
<h4 id="operation" tabindex="-1"><a class="header-anchor" href="#operation" aria-hidden="true">#</a> Operation</h4>
<p>This is the interrupt you specified when defining the event.</p>
<p>Valid interrupt operation values are specified below.</p>
<h4 id="cmdline" tabindex="-1"><a class="header-anchor" href="#cmdline" aria-hidden="true">#</a> CmdLine</h4>
<p>This is the commandline you typed in the prompt.</p>
<p>Please note this is only populated if the interrupt is <strong>after</strong>.</p>
<h3 id="stdout" tabindex="-1"><a class="header-anchor" href="#stdout" aria-hidden="true">#</a> Stdout</h3>
<p>Stdout is written to the terminal. So this can be used to provide multiple
additional lines to the prompt since readline only supports one line for the
prompt itself and three extra lines for the hint text.</p>
<h3 id="order-of-execution" tabindex="-1"><a class="header-anchor" href="#order-of-execution" aria-hidden="true">#</a> Order of execution</h3>
<p>Interrupts are run in alphabetical order. So an event named &quot;alfa&quot; would run
before an event named &quot;zulu&quot;. If you are writing multiple events and the order
of execution matters, then you can prefix the names with a number, eg <code v-pre>10_jump</code></p>
<h3 id="namespacing" tabindex="-1"><a class="header-anchor" href="#namespacing" aria-hidden="true">#</a> Namespacing</h3>
<p>The <code v-pre>onPrompt</code> event differs a little from other events when it comes to the
namespacing of interrupts. Typically you cannot have multiple interrupts with
the same name for an event. However with <code v-pre>onPrompt</code> their names are further
namespaced by the interrupt name. In layman's terms this means <code v-pre>example=before</code>
wouldn't overwrite <code v-pre>example=after</code>.</p>
<p>The reason for this namespacing is because, unlike other events, you might
legitimately want the same name for different interrupts (eg a smart prompt
that has elements triggered from different interrupts).</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/interactive-shell.html">Murex's Interactive Shell</RouterLink>:
What's different about Murex's interactive shell?</li>
<li><RouterLink to="/user-guide/terminal-keys.html">Terminal Hotkeys</RouterLink>:
A list of all the terminal hotkeys and their uses</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/event.html"><code v-pre>event</code></RouterLink>:
Event driven programming for shell scripts</li>
<li><RouterLink to="/events/oncommandcompletion.html"><code v-pre>onCommandCompletion</code></RouterLink>:
Trigger an event upon a command's completion</li>
<li><RouterLink to="/events/onkeypress.html">onkeypress</RouterLink>:</li>
</ul>
</div></template>


