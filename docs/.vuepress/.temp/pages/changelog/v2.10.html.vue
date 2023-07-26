<template><div><h1 id="what-s-new-in-murex-v2-10-change-log" tabindex="-1"><a class="header-anchor" href="#what-s-new-in-murex-v2-10-change-log" aria-hidden="true">#</a> What's new in murex v2.10 - Change Log</h1>
<p>This release brings a few minor improvements and bug fixes rather than big new headline features.</p>
<p>Breaking Changes:</p>
<ul>
<li>None</li>
</ul>
<p>Features:</p>
<ul>
<li>
<p>Added support for ranges to be used as variables. eg</p>
  <pre><code>f: +f -> set v
  echo: @v[2..]</code></pre>
</li>
<li>
<p>Complete re-write of <code v-pre>f</code>. It is now the tool I had always intended when I first created that builtin</p>
</li>
<li>
<p><code v-pre>g</code> now supports being run as a method (eg to filter items from a file list)</p>
</li>
<li>
<p><code v-pre>rx</code> now supports being run as a method (eg to filter items from a file list)</p>
</li>
<li>
<p>Updated autocompletes for <code v-pre>git</code></p>
</li>
<li>
<p>Updated autocompletes for <code v-pre>terraform</code></p>
</li>
</ul>
<p>Minor Changes:</p>
<ul>
<li>
<p>Added support for profiling</p>
</li>
<li>
<p>Removed unused data types</p>
</li>
<li>
<p>Dropped <code v-pre>profile/</code> prefix from the module string for modules and profiles (as seen in <code v-pre>FileRef</code>)</p>
</li>
</ul>
<p>Bug Fixes:</p>
<ul>
<li>
<p>Lots of improvements to how <code v-pre>||</code> and <code v-pre>&amp;&amp;</code> are used in other <code v-pre>runmode</code>'s, <code v-pre>try</code>, and <code v-pre>trypipe</code></p>
</li>
<li>
<p>autoglob prompt now has a mutex wrapped around it to prevent concurrent processes sending the prompt haywire</p>
</li>
<li>
<p>Spellchecker no longer underlines partial words</p>
</li>
<li>
<p>Better removal of \x16 and other non-printable characters leaking into readline</p>
</li>
</ul>
<hr>
<p>Published: 01.08.2022 at 20:10</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/parser/logical-and.html">And (<code v-pre>&amp;&amp;</code>) Logical Operator</RouterLink>:
Continues next operation if previous operation passes</li>
<li><RouterLink to="/user-guide/fileref.html">FileRef</RouterLink>:
How to track what code was loaded and from where</li>
<li><RouterLink to="/user-guide/modules.html">Modules and Packages</RouterLink>:
An introduction to Murex modules and packages</li>
<li><RouterLink to="/user-guide/profile.html">Murex Profile Files</RouterLink>:
A breakdown of the different files loaded on start up</li>
<li><RouterLink to="/parser/logical-or.html">Or (<code v-pre>||</code>) Logical Operator</RouterLink>:
Continues next operation only if previous operation fails</li>
<li><RouterLink to="/user-guide/schedulers.html">Schedulers</RouterLink>:
Overview of the different schedulers (or 'run modes') in Murex</li>
<li><RouterLink to="/user-guide/spellcheck.html">Spellcheck</RouterLink>:
How to enable inline spellchecking</li>
<li><RouterLink to="/commands/range.html"><code v-pre>[</code> (range) </RouterLink>:
Outputs a ranged subset of data from STDIN</li>
<li><RouterLink to="/commands/autocomplete.html"><code v-pre>autocomplete</code></RouterLink>:
Set definitions for tab-completion in the command line</li>
<li><RouterLink to="/commands/f.html"><code v-pre>f</code></RouterLink>:
Lists or filters file system objects (eg files)</li>
<li><RouterLink to="/commands/g.html"><code v-pre>g</code></RouterLink>:
Glob pattern matching for file system objects (eg <code v-pre>*.txt</code>)</li>
<li><RouterLink to="/commands/runmode.html"><code v-pre>runmode</code></RouterLink>:
Alter the scheduler's behaviour at higher scoping level</li>
<li><RouterLink to="/commands/rx.html"><code v-pre>rx</code></RouterLink>:
Regexp pattern matching for file system objects (eg <code v-pre>.*\\.txt</code>)</li>
<li><RouterLink to="/commands/try.html"><code v-pre>try</code></RouterLink>:
Handles errors inside a block of code</li>
<li><RouterLink to="/commands/trypipe.html"><code v-pre>trypipe</code></RouterLink>:
Checks state of each function in a pipeline and exits block on error</li>
</ul>
</div></template>


