<template><div><h1 id="murex-package" tabindex="-1"><a class="header-anchor" href="#murex-package" aria-hidden="true">#</a> <code v-pre>murex-package</code></h1>
<blockquote>
<p>Murex's package manager</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Murex comes with it's own package manager to make managing plugins easier.</p>
<p>The format of the packages is a directory, typically located at <code v-pre>~/.murex_modules</code>,
which contains one or more murex scripts. Each script can be it's own module.
ie there are multiple modules that can be grouped together and distributed as a
single package.</p>
<p>The way packages and modules are represented is as a path:</p>
<pre><code>package/module
</code></pre>
<p><code v-pre>murex-package</code> is a package management tool for administrating murex modules
and packages.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<p>Install a new package</p>
<pre><code>murex-package: install uri -&gt; `&lt;stdout&gt;`
</code></pre>
<p>Update all packages</p>
<pre><code>murex-package: update -&gt; `&lt;stdout&gt;`
</code></pre>
<p>Enable a package or module which had been disabled</p>
<pre><code>murex-package: enable package

murex-package: enable package/module
</code></pre>
<p>Disable a package</p>
<pre><code>murex-package: disable package

murex-package: disable package/module
</code></pre>
<p>Import packages from another package database</p>
<pre><code>murex-package: import [ uri/ | local/path/ ]packages.json -&gt; `&lt;stdout&gt;`
</code></pre>
<p>Check status of murex packages</p>
<pre><code>murex-package: status -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2>
<ul>
<li><code v-pre>cd</code>
Changes working directory to a package's install location</li>
<li><code v-pre>disable</code>
Disables a previously enabled package or module</li>
<li><code v-pre>enable</code>
Enables a previously disabled package or module</li>
<li><code v-pre>import</code>
Import packages described in a backup package DB from user defined URI or local path</li>
<li><code v-pre>install</code>
Installs a package from a user defined URI</li>
<li><code v-pre>list</code>
Returns a list of indexed packages/modules (eg what's enabled or disabled)</li>
<li><code v-pre>reload</code>
Reloads all enabled modules</li>
<li><code v-pre>status</code>
Returns the version status of locally installed packages</li>
<li><code v-pre>update</code>
Updates all installed packages</li>
</ul>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="murex-package-list-enabled-vs-loaded" tabindex="-1"><a class="header-anchor" href="#murex-package-list-enabled-vs-loaded" aria-hidden="true">#</a> <code v-pre>murex-package: list</code>... <code v-pre>enabled</code> vs <code v-pre>loaded</code></h3>
<p><code v-pre>enabled</code> and <code v-pre>disabled</code> reads the package status from disk rather than the
package cache in your current Murex session (like <code v-pre>runtime</code> reports). This
because the typical use for <code v-pre>murex-package list enabled|disabled</code> is to view
which packages and modules will be loaded with any new murex session.</p>
<p>If you wish to view what modules are loaded in a current session then use
<code v-pre>murex-package list loaded</code> instead. This is also equivalent to using
<code v-pre>runtime --modules</code>.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/modules.html">Modules and Packages</RouterLink>:
An introduction to Murex modules and packages</li>
<li><RouterLink to="/user-guide/profile.html">Murex Profile Files</RouterLink>:
A breakdown of the different files loaded on start up</li>
<li><RouterLink to="/commands/alias.html"><code v-pre>alias</code></RouterLink>:
Create an alias for a command</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/function.html"><code v-pre>function</code></RouterLink>:
Define a function block</li>
<li><RouterLink to="/commands/murex-docs.html"><code v-pre>murex-docs</code></RouterLink>:
Displays the man pages for Murex builtins</li>
<li><RouterLink to="/commands/private.html"><code v-pre>private</code></RouterLink>:
Define a private function block</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>runtime</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
</ul>
</div></template>


