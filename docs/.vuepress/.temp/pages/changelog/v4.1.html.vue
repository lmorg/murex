<template><div><h1 id="what-s-new-in-murex-v4-1-x-change-log" tabindex="-1"><a class="header-anchor" href="#what-s-new-in-murex-v4-1-x-change-log" aria-hidden="true">#</a> What's new in murex v4.1.x - Change Log</h1>
<p>The previous releases have brought significant advancements to Murex's syntax but at the cost of longer gaps between releases. So the 4.1.x versions will be shorter releases but focusing on bug fixes. The 4.1.x release notes will be appended to <a href="https://murex.rocks/docs/changelog/v4.1.html" target="_blank" rel="noopener noreferrer">murex.rocks changelog<ExternalLinkIcon/></a> and available on <a href="https://github.com/lmorg/murex/releases" target="_blank" rel="noopener noreferrer">Github releases<ExternalLinkIcon/></a> too</p>
<h2 id="v4-1-7300" tabindex="-1"><a class="header-anchor" href="#v4-1-7300" aria-hidden="true">#</a> v4.1.7300</h2>
<h3 id="breaking-changes" tabindex="-1"><a class="header-anchor" href="#breaking-changes" aria-hidden="true">#</a> Breaking Changes</h3>
<p>None</p>
<h3 id="features" tabindex="-1"><a class="header-anchor" href="#features" aria-hidden="true">#</a> Features</h3>
<ul>
<li>
<p>generic <code v-pre>*</code> data type marshaller can now marshal maps (<a href="https://github.com/lmorg/murex/issues/593" target="_blank" rel="noopener noreferrer">#593<ExternalLinkIcon/></a>)</p>
</li>
<li>
<p>range can count from end when supplied with a negative index (<a href="https://murex.rocks/docs/commands/range.html" target="_blank" rel="noopener noreferrer">read more<ExternalLinkIcon/></a>)</p>
</li>
</ul>
<h3 id="bug-fixes" tabindex="-1"><a class="header-anchor" href="#bug-fixes" aria-hidden="true">#</a> Bug Fixes</h3>
<ul>
<li>
<p><code v-pre>fexec private</code> should inherit module string from the owners module (<a href="https://github.com/lmorg/murex/issues/600" target="_blank" rel="noopener noreferrer">#600<ExternalLinkIcon/></a>)</p>
</li>
<li>
<p><code v-pre>source</code>-ed files should be loaded into a module of a similar name to the parent module</p>
</li>
<li>
<p>new code blocks, <code v-pre>{</code>, should force a new parameter when following a closing block, <code v-pre>}</code> (<a href="https://github.com/lmorg/murex/issues/603" target="_blank" rel="noopener noreferrer">#603<ExternalLinkIcon/></a>)</p>
</li>
<li>
<p>fixed problem with encapsulated variables, eg <code v-pre>$(PWD)</code>, not getting parsed correctly by syntax highlighter</p>
</li>
</ul>
<h2 id="v4-1-6140" tabindex="-1"><a class="header-anchor" href="#v4-1-6140" aria-hidden="true">#</a> v4.1.6140</h2>
<h3 id="breaking-changes-1" tabindex="-1"><a class="header-anchor" href="#breaking-changes-1" aria-hidden="true">#</a> Breaking Changes</h3>
<p>None</p>
<h3 id="features-1" tabindex="-1"><a class="header-anchor" href="#features-1" aria-hidden="true">#</a> Features</h3>
<ul>
<li>
<p>readline: blinking cursor added to fuzzy find. This is just a very minor UI change to make it easier to tell where keyboard input would be updating</p>
</li>
<li>
<p><code v-pre>git</code> autocomplete now includes git aliases</p>
</li>
</ul>
<h3 id="bug-fixes-1" tabindex="-1"><a class="header-anchor" href="#bug-fixes-1" aria-hidden="true">#</a> Bug Fixes</h3>
<ul>
<li>
<p>merge operator (<code v-pre>&lt;~</code>) should unmarshal strings</p>
</li>
<li>
<p>scalars available in all assigns (eg <code v-pre>+=</code>, <code v-pre>-=</code>, etc)</p>
</li>
<li>
<p>lazy assigns allowed when strict types is set to <code v-pre>false</code>. Meaning merge operations and <code v-pre>+=</code> (et al) will work even if the left scalar is unset</p>
</li>
<li>
<p>suppress <code v-pre>!!!</code> style warnings upon <code v-pre>^C</code></p>
</li>
<li>
<p>strip <code v-pre>:</code> colon from command name when command string is passed from the pre-parser (<a href="https://github.com/lmorg/murex/issues/601" target="_blank" rel="noopener noreferrer">#601<ExternalLinkIcon/></a>)</p>
</li>
</ul>
<h2 id="v4-1-5200" tabindex="-1"><a class="header-anchor" href="#v4-1-5200" aria-hidden="true">#</a> v4.1.5200</h2>
<h3 id="breaking-changes-2" tabindex="-1"><a class="header-anchor" href="#breaking-changes-2" aria-hidden="true">#</a> Breaking Changes</h3>
<p>None</p>
<h3 id="features-2" tabindex="-1"><a class="header-anchor" href="#features-2" aria-hidden="true">#</a> Features</h3>
<ul>
<li>
<p>Merge into operator added: <code v-pre>~&gt;</code></p>
</li>
<li>
<p><code v-pre>f</code>: new flag <code v-pre>i</code> -- deprecates <code v-pre>?</code> flag due to <code v-pre>?</code> being a glob pattern. However <code v-pre>?</code> will remain supported</p>
</li>
<li>
<p><code v-pre>$PREFIX</code> variable passed to <strong>Dynamic</strong> and <strong>DynamicDesc</strong> functions in <code v-pre>autocomplete</code> -- includes the partial term (<a href="https://murex.rocks/docs/commands/autocomplete.html#directives" target="_blank" rel="noopener noreferrer">read more<ExternalLinkIcon/></a>)</p>
</li>
<li>
<p><code v-pre>user@</code> completion support for <code v-pre>ssh</code></p>
</li>
<li>
<p>Updated <code v-pre>autocomplete</code> docs (<a href="https://murex.rocks/docs/commands/autocomplete.html" target="_blank" rel="noopener noreferrer">read more<ExternalLinkIcon/></a>)</p>
</li>
<li>
<p>command line is included in error message when error running inside another module</p>
</li>
<li>
<p>errors within other modules now include the command line string to make debugging even easier</p>
</li>
<li>
<p>lambdas success/fail criteria is now driven by the exitnum rather than stdout data type. This allows it to be much more robust. In addition <code v-pre>$.</code> is now mutable</p>
</li>
</ul>
<h3 id="bug-fixes-2" tabindex="-1"><a class="header-anchor" href="#bug-fixes-2" aria-hidden="true">#</a> Bug Fixes</h3>
<ul>
<li>
<p><code v-pre>f</code> improved help message (previously some items displayed code numbers instead of descriptions)</p>
</li>
<li>
<p>readline: hint text flicker removed</p>
</li>
</ul>
<h2 id="v4-1-4200" tabindex="-1"><a class="header-anchor" href="#v4-1-4200" aria-hidden="true">#</a> v4.1.4200</h2>
<h3 id="breaking-changes-3" tabindex="-1"><a class="header-anchor" href="#breaking-changes-3" aria-hidden="true">#</a> Breaking Changes</h3>
<p>None</p>
<h3 id="features-3" tabindex="-1"><a class="header-anchor" href="#features-3" aria-hidden="true">#</a> Features</h3>
<ul>
<li>Merge operator added: <code v-pre>&lt;~</code></li>
</ul>
<h3 id="bug-fixes-3" tabindex="-1"><a class="header-anchor" href="#bug-fixes-3" aria-hidden="true">#</a> Bug Fixes</h3>
<ul>
<li>
<p>readline: fixed out of bounds errors when calculating unicode cell positions</p>
</li>
<li>
<p>readline: fixed panic raised when vim key <code v-pre>p</code> is used on an empty buffer</p>
</li>
<li>
<p>readline: fixed bug where vim mode became unavailable / <code v-pre>esc</code> didn't work as expected. This was due to autocompletion find (<code v-pre>^f</code>, <code v-pre>^r</code>) being launched before <code v-pre>tab</code> completions and state not being handled correctly</p>
</li>
</ul>
<h2 id="v4-1-3010" tabindex="-1"><a class="header-anchor" href="#v4-1-3010" aria-hidden="true">#</a> v4.1.3010</h2>
<h3 id="breaking-changes-4" tabindex="-1"><a class="header-anchor" href="#breaking-changes-4" aria-hidden="true">#</a> Breaking Changes</h3>
<p>None</p>
<h3 id="features-4" tabindex="-1"><a class="header-anchor" href="#features-4" aria-hidden="true">#</a> Features</h3>
<ul>
<li>
<p>lambda support added to piped indexes, ie <code v-pre>example -&gt; [{ lambda }]</code>, to mirror existing <code v-pre>@example[{ lambda }]</code> functionality</p>
</li>
<li>
<p><code v-pre>--autocomplete</code> flag added to <code v-pre>read</code></p>
</li>
</ul>
<h3 id="bug-fixes-4" tabindex="-1"><a class="header-anchor" href="#bug-fixes-4" aria-hidden="true">#</a> Bug Fixes</h3>
<ul>
<li>
<p>application specific environmental variables (eg <code v-pre>&lt;env:FOO=BAR&gt;</code>) were not getting set</p>
</li>
<li>
<p><code v-pre>cast</code> now checks if it is a method. Previously was piping regardless</p>
</li>
<li>
<p>fixed by where nested maps would overwrite when altered with <code v-pre>--merge</code>. Now they merge as expected</p>
</li>
<li>
<p><code v-pre>alter</code> panics are caught issue (<a href="https://github.com/lmorg/murex/issues/536" target="_blank" rel="noopener noreferrer">#536<ExternalLinkIcon/></a>)</p>
</li>
<li>
<p>remove debugging symbol in deprecated math builtin</p>
</li>
</ul>
<h2 id="v4-1-1100" tabindex="-1"><a class="header-anchor" href="#v4-1-1100" aria-hidden="true">#</a> v4.1.1100</h2>
<h3 id="breaking-changes-5" tabindex="-1"><a class="header-anchor" href="#breaking-changes-5" aria-hidden="true">#</a> Breaking Changes</h3>
<ul>
<li>Any code that expects <code v-pre>formap</code> values to be strings despite that value being a different primitive from the originating object. However this type of code would be an anti-pattern so we shouldn't expect much, if any, code written like this</li>
</ul>
<h3 id="features-5" tabindex="-1"><a class="header-anchor" href="#features-5" aria-hidden="true">#</a> Features</h3>
<ul>
<li><code v-pre>formap</code> values are now data type aware. You no longer need to cast values. Keys are still string only and the intention is to keep them this way</li>
</ul>
<h3 id="bug-fixes-5" tabindex="-1"><a class="header-anchor" href="#bug-fixes-5" aria-hidden="true">#</a> Bug Fixes</h3>
<ul>
<li>
<p>improved error messages around a few internal APIs</p>
</li>
<li>
<p>micro optimisations in string parsing (minor order to operations, now converting points at a slightly later stage. This adds zero impact on code complexity and almost zero code refactoring but results in slightly less memory getting copied about)</p>
</li>
</ul>
<h2 id="v4-1-10" tabindex="-1"><a class="header-anchor" href="#v4-1-10" aria-hidden="true">#</a> v4.1.10</h2>
<h3 id="breaking-changes-6" tabindex="-1"><a class="header-anchor" href="#breaking-changes-6" aria-hidden="true">#</a> Breaking Changes</h3>
<p>None</p>
<h3 id="features-6" tabindex="-1"><a class="header-anchor" href="#features-6" aria-hidden="true">#</a> Features</h3>
<p>No new features</p>
<h3 id="bug-fixes-6" tabindex="-1"><a class="header-anchor" href="#bug-fixes-6" aria-hidden="true">#</a> Bug Fixes</h3>
<ul>
<li>
<p>code references in error messages are cropped to terminal width</p>
</li>
<li>
<p>error handling around multiple chained but incompatible tokens (<a href="https://github.com/lmorg/murex/issues/566" target="_blank" rel="noopener noreferrer">#566<ExternalLinkIcon/></a>)</p>
</li>
</ul>
<hr>
<p>Published: 04.06.2023 at 10:44</p>
</div></template>


