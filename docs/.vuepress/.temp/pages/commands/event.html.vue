<template><div><h1 id="event" tabindex="-1"><a class="header-anchor" href="#event" aria-hidden="true">#</a> <code v-pre>event</code></h1>
<blockquote>
<p>Event driven programming for shell scripts</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Create or destroy an event interrupt,</p>
<p>Each event will have subtilty different behaviour depending on the event itself
due to the differing roles of each event system. Therefore it is recommended
that you read the docs on each event to understand its behaviour.</p>
<p>However while they might differ, the <code v-pre>event</code> API does try to retain a level of
external consistency. For example each event in defined via <code v-pre>name=interrupt</code>
where <strong>name</strong> is a user defined handle (like a variable or function would have
a name) and <strong>interrupt</strong> is a system state you wish the event to be fired on.</p>
<p>Each event function will have a payload sent via STDIN which would look a
little like the following:</p>
<pre><code>{
    &quot;Name&quot;: &quot;&quot;,
    &quot;Interrupt&quot;: {}
}
</code></pre>
<p><strong>Name</strong> will always refer to the name you passed when defining the event. And
<strong>Interrupt</strong> will carry any event specific metadata that might be useful to
the event function. Thus the value of <strong>Interrupt</strong> will vary from one event to
another.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>event: event-type name=interrupt { code block }

!event: event-type name
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Create an event:</p>
<pre><code>event: onSecondsElapsed autoquit=60 {
    out &quot;You're 60 second timeout has elapsed. Quitting murex&quot;
    exit 1
}
</code></pre>
<p>Destroy an event:</p>
<pre><code>!event onSecondsElapsed autoquit
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="supported-events" tabindex="-1"><a class="header-anchor" href="#supported-events" aria-hidden="true">#</a> Supported events</h3>
<ul>
<li><RouterLink to="/events/oncommandcompletion.html"><code v-pre>onCommandCompletion</code></RouterLink>:
Trigger an event upon a command's completion</li>
<li><RouterLink to="/events/onfilesystemchange.html"><code v-pre>onFileSystemChange</code></RouterLink>:
Add a filesystem watch</li>
<li><RouterLink to="/events/onprompt.html"><code v-pre>onPrompt</code></RouterLink>:
Events triggered by changes in state of the interactive shell</li>
<li><RouterLink to="/events/onsecondselapsed.html"><code v-pre>onSecondsElapsed</code></RouterLink>:
Events triggered by time intervals</li>
</ul>
<h3 id="ansi-constants" tabindex="-1"><a class="header-anchor" href="#ansi-constants" aria-hidden="true">#</a> ANSI constants</h3>
<p>The <code v-pre>interrupt</code> field in the CLI supports ANSI constants. eg</p>
<pre><code>event: onKeyPress f1={F1-VT100} {
    tout: qs HintText=&quot;Key F1 Pressed&quot;
}
</code></pre>
<h3 id="compiled-events" tabindex="-1"><a class="header-anchor" href="#compiled-events" aria-hidden="true">#</a> Compiled events</h3>
<p>To list compiled event types:</p>
<pre><code>» runtime --events -&gt; formap event ! { out $event }
onCommandCompletion
onFileSystemChange
onKeyPress
onPrompt
onSecondsElapsed
</code></pre>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>event</code></li>
<li><code v-pre>!event</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/formap.html"><code v-pre>formap</code></RouterLink>:
Iterate through a map or other collection of data</li>
<li><RouterLink to="/commands/function.html"><code v-pre>function</code></RouterLink>:
Define a function block</li>
<li><RouterLink to="/commands/open.html"><code v-pre>open</code></RouterLink>:
Open a file with a preferred handler</li>
<li><RouterLink to="/commands/private.html"><code v-pre>private</code></RouterLink>:
Define a private function block</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>runtime</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
</ul>
</div></template>


