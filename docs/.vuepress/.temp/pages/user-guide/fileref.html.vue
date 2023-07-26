<template><div><h1 id="fileref" tabindex="-1"><a class="header-anchor" href="#fileref" aria-hidden="true">#</a> FileRef</h1>
<blockquote>
<p>How to track what code was loaded and from where</p>
</blockquote>
<p>Every function, event, autocompletion and even variable is stored with which
file it was sourced, when it was loaded and what module it was loaded from.
This makes it trivial to identify buggy 3rd party code, malicious libraries, or
even just bugs in your own profiles and/or modules.</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» runtime: --functions -> [[ /agent/FileRef/ ]]
{
    "Column": 5,
    "Line": 5,
    "Source": {
        "DateTime": "2021-03-28T09:10:53.572197+01:00",
        "Filename": "/home/lmorg/.murex_modules/murex-dev/murex-dev.mx",
        "Module": "murex-dev/murex-dev"
    }
}

» runtime --globals -> [[ /DEVOPSBIN/FileRef ]]
{
    "Column": 1,
    "Line": 0,
    "Source": {
        "DateTime": "2021-03-28T09:10:53.541952+01:00",
        "Filename": "/home/lmorg/.murex_modules/devops/global.mx",
        "Module": "devops/global"
    }
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="module-strings-for-non-module-code" tabindex="-1"><a class="header-anchor" href="#module-strings-for-non-module-code" aria-hidden="true">#</a> Module Strings For Non-Module Code</h3>
<h4 id="source" tabindex="-1"><a class="header-anchor" href="#source" aria-hidden="true">#</a> Source</h4>
<p>A common shell idiom is to load shell script files via <code v-pre>source</code> / <code v-pre>.</code>. When
this is done the module string (as seen in the <code v-pre>FileRef</code> structures described
above) will be <code v-pre>source/hash</code> where <strong>hash</strong> will be a unique hash of the file
path and load time.</p>
<p>Thus no two sourced files will share the same module string. Even the same file
but modified and sourced twice (before and after the edit) will have different
module strings due to the load time being part of the hashed data.</p>
<h4 id="repl" tabindex="-1"><a class="header-anchor" href="#repl" aria-hidden="true">#</a> REPL</h4>
<p>Any functions, variables, events, auto-completions, etc created manually,
directly, in the interactive shell will have a module string of <code v-pre>murex</code> and an
empty Filename string.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/modules.html">Modules and Packages</RouterLink>:
An introduction to Murex modules and packages</li>
<li><RouterLink to="/commands/element.html"><code v-pre>[[</code> (element)</RouterLink>:
Outputs an element from a nested structure</li>
<li><RouterLink to="/commands/murex-package.html"><code v-pre>murex-package</code></RouterLink>:
Murex's package manager</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>runtime</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
<li><RouterLink to="/commands/source.html"><code v-pre>source</code> </RouterLink>:
Import Murex code from another file of code block</li>
</ul>
</div></template>


