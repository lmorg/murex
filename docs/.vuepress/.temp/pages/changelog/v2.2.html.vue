<template><div><h1 id="what-s-new-in-murex-v2-2-change-log" tabindex="-1"><a class="header-anchor" href="#what-s-new-in-murex-v2-2-change-log" aria-hidden="true">#</a> What's new in murex v2.2 - Change Log</h1>
<p>This is mainly a bug fix release but it does include one breaking change for <code v-pre>config</code>. Please read for details.</p>
<p>This release sees new optional features plus major bug fixes to the existing
code base:</p>
<h3 id="breaking-changes" tabindex="-1"><a class="header-anchor" href="#breaking-changes" aria-hidden="true">#</a> Breaking Changes:</h3>
<p>One <code v-pre>config</code> <strong>shell</strong> key has changed it's name:</p>
<ul>
<li>spellcheck-block -&gt; <code v-pre>spellcheck-func</code></li>
</ul>
<p>This is to bring spellcheck inline with other keys that contain executable
code.</p>
<p><strong>This change might break some of your existing profile config!</strong></p>
<h3 id="user-facing-changes" tabindex="-1"><a class="header-anchor" href="#user-facing-changes" aria-hidden="true">#</a> User Facing Changes:</h3>
<ul>
<li>
<p>Windows coreutils no longer returns exit code of 13 for aliases</p>
</li>
<li>
<p>Fixed regression bug in fg (added stronger testing here to prevent future
regressions)</p>
</li>
<li>
<p>Generic data type, <code v-pre>*</code>, now uses <code v-pre>tabwriter</code> to align columns</p>
</li>
<li>
<p>Spellchecker error handling tweaked to be more user friendly</p>
</li>
<li>
<p>Environmental variables can been created per command:</p>
<pre><code>&lt;pre&gt;&lt;code&gt;» env &lt;env:foo=bar&gt; | grep foo
foo=bar&lt;/code&gt;&lt;/pre&gt;
</code></pre>
</li>
<li>
<p>stdstream.SetDataType() only works on first call -- ignored thereafter.
This will impact pipes that are written to by many processes but results
in more predictable code due possible timing errors where different
processes are overwriting the stream data type</p>
</li>
<li>
<p>stdstream.SetDataType() only works on first call -- ignored thereafter.
This will impact pipes that are written to by many processes but results
in more predictable code due possible timing errors where different
processes are overwriting the stream data type</p>
</li>
</ul>
<h3 id="non-user-facing-maintenance-changes" tabindex="-1"><a class="header-anchor" href="#non-user-facing-maintenance-changes" aria-hidden="true">#</a> Non-User Facing / Maintenance Changes</h3>
<ul>
<li>
<p>All dependencies have been updated, pinned and the <code v-pre>vendors</code> directory
rebuilt</p>
</li>
<li>
<p>Fixed some issues flagged up in <a href="https://goreportcard.com/report/github.com/lmorg/murex" target="_blank" rel="noopener noreferrer">goreportcard.com<ExternalLinkIcon/></a></p>
</li>
<li>
<p>Some internal API changes that have no UI/UX ramifications but makes the
code more maintainable</p>
</li>
<li>
<p>Lots more unit tests added</p>
</li>
<li>
<p>Lots of code rewritten to avoid potential data races</p>
</li>
</ul>
<hr>
<p>Published: 21.06.2021 at 21:51</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/types/generic.html"><code v-pre>*</code> (generic) </RouterLink>:
generic (primitive)</li>
<li><RouterLink to="/commands/fg.html"><code v-pre>fg</code></RouterLink>:
Sends a background process into the foreground</li>
</ul>
</div></template>


