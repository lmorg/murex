<template><div><h1 id="terminal-hotkeys" tabindex="-1"><a class="header-anchor" href="#terminal-hotkeys" aria-hidden="true">#</a> Terminal Hotkeys</h1>
<blockquote>
<p>A list of all the terminal hotkeys and their uses</p>
</blockquote>
<h2 id="command-prompt" tabindex="-1"><a class="header-anchor" href="#command-prompt" aria-hidden="true">#</a> Command Prompt</h2>
<h3 id="autocomplete" tabindex="-1"><a class="header-anchor" href="#autocomplete" aria-hidden="true">#</a> Autocomplete</h3>
<p>Pressing <code v-pre>tab</code> provides autocompletion suggestions. Suggestions can come in one
of two formats:</p>
<ol>
<li>a gridded view where the hint text (the, typically blue, text under the
prompt) provides the description</li>
<li>a list view where the description is printed alongside the completion
suggestion.</li>
</ol>
<p>While the autocompletion suggestions are open, the following keys are assigned
roles:</p>
<ul>
<li>
<p>arrow keys (<code v-pre>left</code>, <code v-pre>right</code>, <code v-pre>up</code>, <code v-pre>down</code>): highlight different suggestions</p>
</li>
<li>
<p><code v-pre>tab</code>: highlight the next suggestion</p>
</li>
<li>
<p><code v-pre>shift</code>+<code v-pre>tab</code>: highlight the previous suggestion</p>
</li>
<li>
<p><code v-pre>enter</code> / <code v-pre>return</code>: this selects the highlighted autocompletion</p>
</li>
<li>
<p><code v-pre>esc</code>: closes the suggestions without selecting one</p>
</li>
<li>
<p><code v-pre>ctrl</code>+<code v-pre>f</code>: fuzzy find in the suggestions</p>
</li>
<li>
<p><code v-pre>f1</code>: enable / disable preview box. This will hide your terminal output while
enabled. The preview box supports additional key bindings (see below)</p>
</li>
</ul>
<h3 id="fuzzy-find-autocomplete" tabindex="-1"><a class="header-anchor" href="#fuzzy-find-autocomplete" aria-hidden="true">#</a> Fuzzy Find Autocomplete</h3>
<p>Pressing <code v-pre>ctrl</code>+<code v-pre>f</code> either from the prompt, or while the autocomplete
suggestions are open, will open up the fuzzy find dialog to search through
available suggestions. This can also be used to quickly jump to specific
sub-directories.</p>
<p>Your typed search terms will appear in the hint text.</p>
<p>By default the fuzzy finder will look for any item that includes <em>all</em> of the
search words. However the search behavior can be changed if the first search
term is any of the following:</p>
<ul>
<li>
<p><code v-pre>or</code>: show results that match <em>any</em> of the search terms. eg <code v-pre>or .md .txt</code>
will match both markdown and txt files (when finding files in completion
suggestions).</p>
</li>
<li>
<p><code v-pre>!</code>: only show suggestions that do not match any of the search terms. eg
<code v-pre>! .md .txt</code> will match all files except markdown and txt files (when finding
files in completion suggestions).</p>
</li>
<li>
<p><code v-pre>g</code>: show only results that match a shell glob. eg <code v-pre>*.txt</code>. This mode is
automatically assumed if you include an abstricts in your search term.</p>
</li>
<li>
<p><code v-pre>rx</code>: use a regexp pattern matcher instead of any fuzzy search. Expressions
will be case insensitive and non-greedy by default.</p>
</li>
</ul>
<p>Aside from globbing matches, searching in fuzzy finder is not case sensitive.</p>
<p>While the fuzzy finder is open, the following keys are assigned roles:</p>
<ul>
<li>
<p>arrow keys (<code v-pre>left</code>, <code v-pre>right</code>, <code v-pre>up</code>, <code v-pre>down</code>): highlight different suggestions</p>
</li>
<li>
<p><code v-pre>tab</code>: highlight the next suggestion</p>
</li>
<li>
<p><code v-pre>shift</code>+<code v-pre>tab</code>: highlight the previous suggestion</p>
</li>
<li>
<p><code v-pre>enter</code> / <code v-pre>return</code>: this selects the highlighted autocompletion</p>
</li>
<li>
<p><code v-pre>esc</code>: cancel search</p>
</li>
<li>
<p><code v-pre>f1</code>: enable / disable preview box. This will hide your terminal output while
enabled. The preview box supports additional key bindings (see below)</p>
</li>
</ul>
<h3 id="preview-box" tabindex="-1"><a class="header-anchor" href="#preview-box" aria-hidden="true">#</a> Preview Box</h3>
<p>The preview box is a way of quickly examining the contents of a function, man
page, text file or even image, based on what autocomplete suggestion is
highlighted.</p>
<p>While the preview box is open, the rest of your terminal output will be hidden.
However once you close it, that output will reappear.</p>
<p>While the preview box is open, the following keys are assigned roles:</p>
<ul>
<li>
<p><code v-pre>f1</code>: closes the preview box</p>
</li>
<li>
<p><code v-pre>page up</code> scroll up the contents of the preview box, one page at a time</p>
</li>
<li>
<p><code v-pre>ctrl</code>+<code v-pre>arrow up</code> scroll up the contents of the preview box, one page at a
time (IBM keyboard layouts)</p>
</li>
<li>
<p><code v-pre>option</code>+<code v-pre>arrow up</code> scroll up the contents of the preview box, one page at a
time (Apple keyboard layouts)</p>
</li>
<li>
<p><code v-pre>page down</code> scroll down the contents of the preview box, one page at a time</p>
</li>
<li>
<p><code v-pre>ctrl</code>+<code v-pre>arrow down</code> scroll down the contents of the preview box, one page at
a time (IBM keyboard layouts)</p>
</li>
<li>
<p><code v-pre>option</code>+<code v-pre>arrow down</code> scroll down the contents of the preview box, one page
at a time (Apple keyboard layouts)</p>
</li>
</ul>
<h3 id="search-shell-history" tabindex="-1"><a class="header-anchor" href="#search-shell-history" aria-hidden="true">#</a> Search Shell History</h3>
<p>This displays up your timestamped shell history as an autocomplete list with
fuzzy find activated. Using <code v-pre>ctrl</code>+<code v-pre>r</code> you can rapidly rerun previous
command lines.</p>
<p>From here, the usual autocomplete / fuzzy find hotkeys apply. Such as pressing
<code v-pre>esc</code> to cancel history completion.</p>
<p>If the prompt line is not empty, then the current line is included in the
history search.</p>
<h3 id="line-editing" tabindex="-1"><a class="header-anchor" href="#line-editing" aria-hidden="true">#</a> Line Editing</h3>
<p>These are the various hotkeys and editing modes available in Murex's
interactive command prompt.</p>
<h4 id="navigation" tabindex="-1"><a class="header-anchor" href="#navigation" aria-hidden="true">#</a> Navigation</h4>
<ul>
<li>
<p>Arrow keys, <code v-pre>left</code> and <code v-pre>right</code>: move the cursor forwards or backwards in line</p>
</li>
<li>
<p>Arrow keys, <code v-pre>up</code> and <code v-pre>down</code>: If the command line spans multiple lines on the
screen then this will jump up or down to the next/previous line. When at the
top or bottom line, the <code v-pre>up</code> or <code v-pre>down</code> keys will search through your history
of past command lines</p>
</li>
<li>
<p><code v-pre>alt</code>+<code v-pre>b</code>: jump backwards a word at a time (Emacs compatibility)</p>
</li>
<li>
<p><code v-pre>ctrl</code>+<code v-pre>left</code>: jump backwards a word at a time (IBM keyboard layouts)</p>
</li>
<li>
<p><code v-pre>option</code>+<code v-pre>left</code>: jump backwards a word at a time (Apple keyboard layouts)</p>
</li>
<li>
<p><code v-pre>alt</code>+<code v-pre>f</code>: jump forwards a word at a time (Emacs compatibility)</p>
</li>
<li>
<p><code v-pre>ctrl</code>+<code v-pre>right</code>: jump forwards a word at a time (IBM keyboard layouts)</p>
</li>
<li>
<p><code v-pre>option</code>+<code v-pre>right</code>: jump forwards a word at a time (Apple keyboard layouts)</p>
</li>
<li>
<p><code v-pre>ctrl</code>+<code v-pre>a</code>: jump to beginning of line</p>
</li>
<li>
<p><code v-pre>home</code>: jump to beginning of line</p>
</li>
<li>
<p><code v-pre>ctrl</code>+<code v-pre>e</code>: jump to end of line</p>
</li>
<li>
<p><code v-pre>end</code>: jump to end of line</p>
</li>
</ul>
<h4 id="general-hotkeys" tabindex="-1"><a class="header-anchor" href="#general-hotkeys" aria-hidden="true">#</a> General Hotkeys</h4>
<ul>
<li><code v-pre>ctrl</code>+<code v-pre>k</code>: clears line after cursor</li>
<li><code v-pre>ctrl</code>+<code v-pre>u</code>: clears the whole line</li>
</ul>
<h3 id="vim-keys" tabindex="-1"><a class="header-anchor" href="#vim-keys" aria-hidden="true">#</a> Vim Keys</h3>
<p>Pressing <code v-pre>esc</code> while no autocomplete suggestions are shown will switch the
line editor into <strong>vim keys</strong> mode.</p>
<p>Press <code v-pre>i</code> to return to normal editing mode.</p>
<h4 id="supported-keys" tabindex="-1"><a class="header-anchor" href="#supported-keys" aria-hidden="true">#</a> Supported Keys</h4>
<ul>
<li><code v-pre>a</code>: insert after current character</li>
<li><code v-pre>A</code>: insert at end of line</li>
<li><code v-pre>b</code>: jump to beginning of word</li>
<li><code v-pre>B</code>: jump to previous whitespace</li>
<li><code v-pre>d</code>: delete mode</li>
<li><code v-pre>D</code>: delete characters</li>
<li><code v-pre>e</code>: jump to end of word</li>
<li><code v-pre>E</code>: jump to next whitespace</li>
<li><code v-pre>h</code>: previous character (like <code v-pre>left</code>)</li>
<li><code v-pre>i</code>: insert mode</li>
<li><code v-pre>I</code>: insert at beginning of line</li>
<li><code v-pre>l</code>: next character (like <code v-pre>right</code>)</li>
<li><code v-pre>p</code>: paste after</li>
<li><code v-pre>P</code>: paste before</li>
<li><code v-pre>r</code>: replace character (replace once)</li>
<li><code v-pre>R</code>: replace many characters</li>
<li><code v-pre>u</code>: undo</li>
<li><code v-pre>v</code>: visual editor (opens line in <code v-pre>$EDITOR</code>)</li>
<li><code v-pre>w</code>: jump to end of word</li>
<li><code v-pre>W</code>: jump to next whitespace</li>
<li><code v-pre>x</code>: delete character</li>
<li><code v-pre>y</code>: yank (copy line)</li>
<li><code v-pre>Y</code>: same as <code v-pre>y</code></li>
<li><code v-pre>[</code>: jump to previous brace</li>
<li><code v-pre>]</code>: jump to next brace</li>
<li><code v-pre>$</code>: jump to end of line</li>
<li><code v-pre>%</code>: jump to either end of matching bracket</li>
<li><code v-pre>0</code> to <code v-pre>9</code>: repeat action <em>n</em> times. eg <code v-pre>5x</code> would delete (<code v-pre>x</code>) five (<code v-pre>5</code>)
characters</li>
</ul>
<h4 id="full-screen-editing-via-editor" tabindex="-1"><a class="header-anchor" href="#full-screen-editing-via-editor" aria-hidden="true">#</a> Full Screen Editing via <code v-pre>$EDITOR</code></h4>
<p>When in &quot;vim keys&quot; mode, press <code v-pre>v</code> to bring up the visual editor. The editor
will be whichever command is stored in the <code v-pre>$EDITOR</code> environmental variable.</p>
<h3 id="recalling-previous-words" tabindex="-1"><a class="header-anchor" href="#recalling-previous-words" aria-hidden="true">#</a> Recalling Previous Words</h3>
<ul>
<li><code v-pre>shift</code>+<code v-pre>f1</code> recalls the first word</li>
<li><code v-pre>shift</code>+<code v-pre>f2</code> recalls the second word</li>
<li>...</li>
<li><code v-pre>shift</code>+<code v-pre>f12</code> recalls the twelfth word</li>
</ul>
<p>In the following example, code inside square brackets represent key presses
rather than text:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» echo two three four five six seven eight nine
two three four five six seven eight nine
» [shift+f1]echo [shift+f5]five
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="job-control" tabindex="-1"><a class="header-anchor" href="#job-control" aria-hidden="true">#</a> Job Control</h2>
<p>While processes are running, the following keys are assigned roles:</p>
<ul>
<li>
<p><code v-pre>ctrl</code>+<code v-pre>c</code>: kill foreground process. Pressing this will send a kill (SIGINT)
request to the foreground process</p>
</li>
<li>
<p><code v-pre>ctrl</code>+<code v-pre>\</code>: kill all running processes in current shell session, including
any background processes too. This hotkey is a effectively an emergency kill
switch to bring you back to the command prompt should <code v-pre>ctrl</code>+<code v-pre>c</code> prove
ineffective. Use this sparingly because it doesn't allow processes to end
gracefully</p>
</li>
<li>
<p><code v-pre>ctrl</code>+<code v-pre>z</code>: suspend foreground process. This will take you back to the prompt
and from there you can then use job control to resume execution in either the
foreground or background. (<RouterLink to="/commands/fid-list.html">read more</RouterLink>)</p>
</li>
</ul>
<h2 id="miscellaneous" tabindex="-1"><a class="header-anchor" href="#miscellaneous" aria-hidden="true">#</a> Miscellaneous</h2>
<h3 id="cancel-prompt" tabindex="-1"><a class="header-anchor" href="#cancel-prompt" aria-hidden="true">#</a> Cancel Prompt</h3>
<p>Pressing <code v-pre>ctrl</code>+<code v-pre>c</code> while on the prompt will clear the prompt. This is similar
to <code v-pre>ctrl</code>+<code v-pre>u</code>.</p>
<h3 id="end-of-file" tabindex="-1"><a class="header-anchor" href="#end-of-file" aria-hidden="true">#</a> End Of File</h3>
<p>Pressing <code v-pre>ctrl</code>+<code v-pre>d</code> on an empty prompt will send EOF (end of file). This will
exit that running shell session.</p>
<h3 id="alternative-cancel-key" tabindex="-1"><a class="header-anchor" href="#alternative-cancel-key" aria-hidden="true">#</a> Alternative Cancel Key</h3>
<p><code v-pre>ctrl</code>+<code v-pre>g</code> performs the same action as <code v-pre>esc</code> at all states of the interactive
shell.</p>
<h3 id="clear-screen" tabindex="-1"><a class="header-anchor" href="#clear-screen" aria-hidden="true">#</a> Clear Screen</h3>
<p>Pressing <code v-pre>ctrl</code>+<code v-pre>l</code> will clear the screen.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/interactive-shell.html">Murex's Interactive Shell</RouterLink>:
What's different about Murex's interactive shell?</li>
<li><a href="/rosetta">Rosetta Stone</a>:
A tabulated list of Bashism's and their equivalent Murex syntax</li>
<li><RouterLink to="/user-guide/spellcheck.html">Spellcheck</RouterLink>:
How to enable inline spellchecking</li>
</ul>
</div></template>


