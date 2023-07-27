<template><div><h1 id="profile-files" tabindex="-1"><a class="header-anchor" href="#profile-files" aria-hidden="true">#</a> Profile Files</h1>
<blockquote>
<p>A breakdown of the different files loaded on start up</p>
</blockquote>
<p>Murex has several profile files which are loaded in the following order of
execution:</p>
<ol>
<li><code v-pre>~/.murex_preload</code></li>
<li><code v-pre>~/.murex_modules/*/</code></li>
<li><code v-pre>~/.murex_profile</code></li>
</ol>
<h3 id="murex-preload" tabindex="-1"><a class="header-anchor" href="#murex-preload" aria-hidden="true">#</a> <code v-pre>.murex_preload</code></h3>
<p>This file should only used to define any environmental variables that might
need to be set before the modules are loaded (eg including directories in
<code v-pre>$PATH</code> if you have anything installed in non-standard locations).</p>
<p>Most of the time this file will be empty bar the standard warning message:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code># This file is loaded before any murex modules. It should only contain
# environmental variables required for the modules to work eg:
#
#     export PATH=...
#
# Any other profile config belongs in your profile script instead:
# /home/$USER/.murex_profile
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>This file is created upon the first run of Murex.</p>
<h3 id="murex-modules" tabindex="-1"><a class="header-anchor" href="#murex-modules" aria-hidden="true">#</a> <code v-pre>.murex_modules/</code></h3>
<p>Murex's module directory - where all the modules are installed
to. This directory is managed by <code v-pre>murex-package</code> builtin.</p>
<h3 id="murex-profile" tabindex="-1"><a class="header-anchor" href="#murex-profile" aria-hidden="true">#</a> <code v-pre>.murex_profile</code></h3>
<p>This file is comparable to <code v-pre>.bash_profile</code>, <code v-pre>.bashrc</code> and <code v-pre>.zshrc</code> etc. It
is the standard place to put all user and/or machine specific config in.</p>
<p><code v-pre>.murex_profile</code> is only read from the users home directory. Unlike bash et
al, profiles will not be read from <code v-pre>/etc/profile.d</code> nor similar. Modules
should be used in its place.</p>
<h2 id="overriding-the-default-paths-xdg" tabindex="-1"><a class="header-anchor" href="#overriding-the-default-paths-xdg" aria-hidden="true">#</a> Overriding The Default Paths (XDG)</h2>
<p>Some individuals, particularly those running Linux, follow a standard called
<a href="https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html" target="_blank" rel="noopener noreferrer">XDG Base Directory Specification<ExternalLinkIcon/></a>.
While Murex does not adhere to this standard, instead conforming to the
<em>de facto</em> standard defined by the past precedents of previous shells, in
order to offer flexibility for those who do prefer the XDG specification
Murex does support overriding its own default paths via special environmental
variables.</p>
<ul>
<li><code v-pre>MUREX_PRELOAD</code> defines the preload path (and file name)</li>
<li><code v-pre>MUREX_MODULES</code> defines the module path (only)</li>
<li><code v-pre>MUREX_PROFILE</code> defines the profile path (and file name)</li>
</ul>
<p>Where <code v-pre>MUREX_PRELOAD</code> and/or <code v-pre>MUREX_PROFILE</code> are directories rather than
absolute file names, the path is appended with the default file names as
named above.</p>
<p>For people who wish to use XDG paths, in many instances you can get away
with setting the follow prior to launching Murex (eg in <code v-pre>/etc/profile.d</code>):</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>MUREX_PRELOAD="$XDG_CONFIG_HOME/murex/"
MUREX_MODULES="$XDG_CONFIG_HOME/murex/"
MUREX_PROFILE="$XDG_CONFIG_HOME/murex/"
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>This, however, depends on <code v-pre>$XDG_CONFIG_HOME</code> pointing to a single path rather
than an array of paths (like <code v-pre>$PATH</code>). In that instance you can still use
custom paths in Murex but you might need to get a little more creative in
how you define that value.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/modules.html">Modules and Packages</RouterLink>:
An introduction to Murex modules and packages</li>
<li><RouterLink to="/commands/export.html"><code v-pre>export</code></RouterLink>:
Define an environmental variable and set it's value</li>
<li><RouterLink to="/commands/murex-package.html"><code v-pre>murex-package</code></RouterLink>:
Murex's package manager</li>
</ul>
</div></template>


