<template><div><h1 id="what-s-new-in-murex-v4-3-change-log" tabindex="-1"><a class="header-anchor" href="#what-s-new-in-murex-v4-3-change-log" aria-hidden="true">#</a> What's new in murex v4.3 - Change Log</h1>
<p>This brings improved support on Windows plus one breaking change from the previous release (v4.2)</p>
<h3 id="breaking-changes" tabindex="-1"><a class="header-anchor" href="#breaking-changes" aria-hidden="true">#</a> Breaking Changes</h3>
<ul>
<li><code v-pre>pre-prompt-func</code> and <code v-pre>post-prompt-func</code> have been replaced with <code v-pre>onPrompt</code> events: <code v-pre>event onKeyPress example=before { ... }</code></li>
</ul>
<h3 id="features" tabindex="-1"><a class="header-anchor" href="#features" aria-hidden="true">#</a> Features</h3>
<ul>
<li>
<p>New events, <code v-pre>onPrompt</code>, triggered when the interactive prompt is at various stages</p>
</li>
<li>
<p><strong>Dynamic</strong> and <strong>DynamicDesc</strong> autocompletions can have their partial term prefix filter disabled ([read more])(https://murex.rocks/docs/commands/autocomplete.html#ignoreprefix-boolean-false) (<a href="%5B#630%5D(https://github.com/lmorg/murex/issues/635)">discussion</a>)</p>
</li>
<li>
<p>New autocompletions: <code v-pre>gping</code></p>
</li>
</ul>
<h3 id="bug-fixes" tabindex="-1"><a class="header-anchor" href="#bug-fixes" aria-hidden="true">#</a> Bug Fixes:</h3>
<ul>
<li>
<p>Improved Windows support. Arrow keys and other special keys are now handled correctly (<a href="https://github.com/lmorg/murex/issues/630" target="_blank" rel="noopener noreferrer">#630<ExternalLinkIcon/></a>)</p>
</li>
<li>
<p>Fix <code v-pre>ctrl</code>+<code v-pre>a</code> offset by 1 error (<a href="https://github.com/lmorg/murex/pull/628" target="_blank" rel="noopener noreferrer">#628<ExternalLinkIcon/></a>)</p>
</li>
</ul>
<hr>
<p>Published: 02.07.2023 at 09:55</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/autocomplete.html"><code v-pre>autocomplete</code></RouterLink>:
Set definitions for tab-completion in the command line</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/event.html"><code v-pre>event</code></RouterLink>:
Event driven programming for shell scripts</li>
<li><RouterLink to="/events/onprompt.html"><code v-pre>onPrompt</code></RouterLink>:
Changes in state of the interactive shell</li>
</ul>
</div></template>


