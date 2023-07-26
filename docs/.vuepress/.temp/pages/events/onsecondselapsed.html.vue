<template><div><h1 id="onsecondselapsed" tabindex="-1"><a class="header-anchor" href="#onsecondselapsed" aria-hidden="true">#</a> <code v-pre>onSecondsElapsed</code></h1>
<blockquote>
<p>Events triggered by time intervals</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>onSecondsElapsed</code> events are triggered every <em>n</em> seconds.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>event: onSecondsElapsed name=seconds { code block }

!event: onSecondsElapsed name
</code></pre>
<h2 id="valid-interrupts" tabindex="-1"><a class="header-anchor" href="#valid-interrupts" aria-hidden="true">#</a> Valid Interrupts</h2>
<ul>
<li><code v-pre>&lt;seconds&gt;</code>
Duration in seconds. eg <code v-pre>60</code> would be 60 seconds / 1 minute</li>
</ul>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>event: onSecondsElapsed example=60 {
    out: &quot;60 seconds has passed&quot;
}
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="payload" tabindex="-1"><a class="header-anchor" href="#payload" aria-hidden="true">#</a> Payload</h3>
<p>The following payload is passed to the function via STDIN:</p>
<pre><code>{
    &quot;Name&quot;: &quot;&quot;,
    &quot;Interrupt&quot;: 0
}
</code></pre>
<h4 id="name" tabindex="-1"><a class="header-anchor" href="#name" aria-hidden="true">#</a> Name</h4>
<p>This is the name you specified when defining the event.</p>
<h4 id="interrupt" tabindex="-1"><a class="header-anchor" href="#interrupt" aria-hidden="true">#</a> Interrupt</h4>
<p>This is the duration you defined the event to wait for.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/event.html"><code v-pre>event</code></RouterLink>:
Event driven programming for shell scripts</li>
</ul>
</div></template>


