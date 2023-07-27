<template><div><h1 id="what-s-new-in-murex-v2-6-change-log" tabindex="-1"><a class="header-anchor" href="#what-s-new-in-murex-v2-6-change-log" aria-hidden="true">#</a> What's new in murex v2.6 - Change Log</h1>
<p>This update has introduced a potential breaking change: variables now need to be defined before usage otherwise the commandline will fail. Read notes to learn how to disable this feature where needed. Also included in this release is the <code v-pre>select</code> command as part of the standard build.</p>
<p>Breaking Change:</p>
<ul>
<li><code v-pre>strict-vars</code> is enabled by default. This means if a variable doesn't exist / hasn't been instantiated then your command will fail. You can disable this by putting <code v-pre>config: set proc strict-vars false</code> inside any affected <code v-pre>function</code>s and/or in your <code v-pre>~/.murex_profile</code></li>
</ul>
<p>Features:</p>
<ul>
<li>
<p><code v-pre>select</code> builtin now added to the default build. This built in will stream tabulated data into an in memory sqlite3 database and take SQL arguments to query that data. Thus allowing you to inline SQL queries as part of your native command line. This addition will add complications for anyone building from source on Windows. If you are a Windows user you are recommended to use WSL or, if you desperately need to use Murex natively on Windows, download the appropriate per-compiled binary from the downloads page</p>
</li>
<li>
<p>When used in WSL, any Windows mounts included in <code v-pre>$PATH</code> will now be subject to the same inclusion rules as Windows file systems when Murex is used in Windows. This means only executables (ie files with <code v-pre>.exe</code>, amongst other, extensions are present). This has dramatically cleaned up autocomplete suggestions when running Murex in WSL.</p>
</li>
<li>
<p>New <code v-pre>autocomplete</code> directive: <code v-pre>FileRegexp</code>. Use this if you want to return only files where the file name matches a defined pattern (eg only <strong>.txt</strong> files: <code v-pre>\.txt$</code>)</p>
</li>
<li>
<p>HTTP client (as used by <code v-pre>get</code>, <code v-pre>getfile</code>, <code v-pre>post</code>, <code v-pre>open</code>, etc) supports additional logic with dynamically guessing the correct data type to assign to STDOUT where the <code v-pre>Content-Type</code> hasn't been mapped. Any <code v-pre>Content-Type</code> with a <code v-pre>+json</code> suffix will be assumed to be <code v-pre>json</code> data type</p>
</li>
<li>
<p><code v-pre>getfile</code> output has been revamped. If STDOUT is a TTY then you'll now see a progress bar</p>
</li>
<li>
<p><code v-pre>.tfstate</code> (Terraform state files) have been added to supported JSON extensions</p>
</li>
<li>
<p><code v-pre>~/.ssh/config</code> hosts have been added to SSH (et al)'s autocomplete suggestions</p>
</li>
<li>
<p><code v-pre>murex-package: reload</code> now only reloads packages, not <code v-pre>.murex_preload nor </code>.murex_profile`</p>
</li>
<li>
<p>website tweaks: updated INSTALL doc, CSS fixes</p>
</li>
<li>
<p>debug message added to warn when fork FIDs are being killed</p>
</li>
</ul>
<p>Non-user facing changes (internal changes to the Murex code base):</p>
<ul>
<li>
<p>Additional tests</p>
</li>
<li>
<p>Older function definitions replaced with the newer (correct) way to define functions and methods. This allows for greater discovery in the interactive terminal</p>
</li>
<li>
<p>Root unit tests finish quicker</p>
</li>
</ul>
<p>Bug fixes:</p>
<ul>
<li>
<p>HTTP client (as used by <code v-pre>get</code>, <code v-pre>getfile</code>, <code v-pre>post</code>, <code v-pre>open</code>, etc) wasn't breaking down the <code v-pre>Content-Type</code> header correctly. This only affected unknown MIME types where Murex had to guess the correct data type to assign to STDOUT. This is an edge case bug.</p>
</li>
<li>
<p><code v-pre>while</code>, when used without a separate parameter for the conditional logic, was buffering the STDOUT for each loop and outputting it after that iteration had ended. <code v-pre>while</code> now uses a &quot;tee&quot; stdio writer which means it can stream STDOUT while the iteration is in progress Synonyms:</p>
</li>
</ul>
<hr>
<p>Published: 26.02.2022 at 18:33</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/autocomplete.html"><code v-pre>autocomplete</code></RouterLink>:
Set definitions for tab-completion in the command line</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/optional/select.html"><code v-pre>select</code> </RouterLink>:
Inlining SQL into shell pipelines</li>
<li><RouterLink to="/commands/while.html"><code v-pre>while</code></RouterLink>:
Loop until condition false</li>
</ul>
</div></template>


