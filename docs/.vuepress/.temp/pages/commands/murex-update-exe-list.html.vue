<template><div><h1 id="murex-update-exe-list" tabindex="-1"><a class="header-anchor" href="#murex-update-exe-list" aria-hidden="true">#</a> <code v-pre>murex-update-exe-list</code></h1>
<blockquote>
<p>Forces Murex to rescan $PATH looking for executables</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>On application launch, Murex scans and caches all the executables found in
$PATH on your host. Murex then does regular scans there after. However if
you want to force a new scan (for example you've just installed a new
program and you want it to appear in tab completion) then you can run <code v-pre>murex-update-exe-list</code>.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>murex-update-exe-list
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» murex-update-exe-list
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>Murex will automatically update the exe list each time tab completion is
invoked for command name completion via the REPL shell.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/cpuarch.html"><code v-pre>cpuarch</code></RouterLink>:
Output the hosts CPU architecture</li>
<li><RouterLink to="/commands/cpucount.html"><code v-pre>cpucount</code></RouterLink>:
Output the number of CPU cores available on your host</li>
<li><RouterLink to="/commands/os.html"><code v-pre>os</code></RouterLink>:
Output the auto-detected OS name</li>
</ul>
</div></template>


