<template><div><h1 id="open" tabindex="-1"><a class="header-anchor" href="#open" aria-hidden="true">#</a> <code v-pre>open</code></h1>
<blockquote>
<p>Open a file with a preferred handler</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>open</code> is a smart tool for reading files:</p>
<ol>
<li>It will read a file from disk or a HTTP(S) endpoints</li>
<li>Detect the file type via file extension or HTTP header <code v-pre>Content-Type</code></li>
<li>It intelligently writes to STDOUT</li>
</ol>
<ul>
<li>If STDOUT is a TTY it will perform any transformations to render to the
terminal (eg using inlining images)</li>
<li>If STDOUT is a pipe then it will write a byte stream with the relevant
data-type</li>
</ul>
<ol start="4">
<li>If there are no open handlers then it will fallback to the systems default.
eg <code v-pre>open</code> (on macOS, Linux), <code v-pre>open-xdg</code> (X11), etc.</li>
</ol>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>open filename[.gz]|uri -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>» open https://api.github.com/repos/lmorg/murex/issues -&gt; foreach issue { out: &quot;$issue[number]: $issue[title]&quot; }
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="file-extensions" tabindex="-1"><a class="header-anchor" href="#file-extensions" aria-hidden="true">#</a> File Extensions</h3>
<p>Supported file extensions are listed in <code v-pre>config</code> under the app and key names of
<strong>shell</strong>, <strong>extensions</strong>.</p>
<p>Unsupported file extensions are defaulted to generic, <code v-pre>*</code>.</p>
<p>Files with a <code v-pre>.gz</code> extension are assumed to be gzipped and thus are are
automatically expanded.</p>
<h3 id="mime-types" tabindex="-1"><a class="header-anchor" href="#mime-types" aria-hidden="true">#</a> MIME Types</h3>
<p>The <code v-pre>Content-Type</code> HTTP header is compared against a list of MIME types, which
are stored in <code v-pre>config</code> under the app and key names of <strong>shell</strong>, <strong>mime-types</strong>.</p>
<p>There is a little bit of additional logic to determine the Murex data-type to
use should the MIME type not appear in <code v-pre>config</code>, as seen in the following code:</p>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token keyword">package</span> lang

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">"regexp"</span>
	<span class="token string">"strings"</span>

	<span class="token string">"github.com/lmorg/murex/lang/types"</span>
<span class="token punctuation">)</span>

<span class="token keyword">var</span> rxMimePrefix <span class="token operator">=</span> regexp<span class="token punctuation">.</span><span class="token function">MustCompile</span><span class="token punctuation">(</span><span class="token string">`^([-0-9a-zA-Z]+)/.*$`</span><span class="token punctuation">)</span>

<span class="token comment">// MimeToMurex gets the murex data type for a corresponding MIME</span>
<span class="token keyword">func</span> <span class="token function">MimeToMurex</span><span class="token punctuation">(</span>mimeType <span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token builtin">string</span> <span class="token punctuation">{</span>
	mime <span class="token operator">:=</span> strings<span class="token punctuation">.</span><span class="token function">Split</span><span class="token punctuation">(</span>mimeType<span class="token punctuation">,</span> <span class="token string">";"</span><span class="token punctuation">)</span><span class="token punctuation">[</span><span class="token number">0</span><span class="token punctuation">]</span>
	mime <span class="token operator">=</span> strings<span class="token punctuation">.</span><span class="token function">TrimSpace</span><span class="token punctuation">(</span>mime<span class="token punctuation">)</span>
	mime <span class="token operator">=</span> strings<span class="token punctuation">.</span><span class="token function">ToLower</span><span class="token punctuation">(</span>mime<span class="token punctuation">)</span>

	<span class="token comment">// Find a direct match. This is only used to pick up edge cases, eg text files used as images.</span>
	dt <span class="token operator">:=</span> mimes<span class="token punctuation">[</span>mime<span class="token punctuation">]</span>
	<span class="token keyword">if</span> dt <span class="token operator">!=</span> <span class="token string">""</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> dt
	<span class="token punctuation">}</span>

	<span class="token comment">// No direct match found. Fall back to prefix.</span>
	prefix <span class="token operator">:=</span> rxMimePrefix<span class="token punctuation">.</span><span class="token function">FindStringSubmatch</span><span class="token punctuation">(</span>mime<span class="token punctuation">)</span>
	<span class="token keyword">if</span> <span class="token function">len</span><span class="token punctuation">(</span>prefix<span class="token punctuation">)</span> <span class="token operator">!=</span> <span class="token number">2</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> types<span class="token punctuation">.</span>Generic
	<span class="token punctuation">}</span>

	<span class="token keyword">switch</span> prefix<span class="token punctuation">[</span><span class="token number">1</span><span class="token punctuation">]</span> <span class="token punctuation">{</span>
	<span class="token keyword">case</span> <span class="token string">"text"</span><span class="token punctuation">,</span> <span class="token string">"i-world"</span><span class="token punctuation">,</span> <span class="token string">"message"</span><span class="token punctuation">:</span>
		<span class="token keyword">return</span> types<span class="token punctuation">.</span>String

	<span class="token keyword">case</span> <span class="token string">"audio"</span><span class="token punctuation">,</span> <span class="token string">"music"</span><span class="token punctuation">,</span> <span class="token string">"video"</span><span class="token punctuation">,</span> <span class="token string">"image"</span><span class="token punctuation">,</span> <span class="token string">"model"</span><span class="token punctuation">:</span>
		<span class="token keyword">return</span> types<span class="token punctuation">.</span>Binary

	<span class="token keyword">case</span> <span class="token string">"application"</span><span class="token punctuation">:</span>
		<span class="token keyword">if</span> strings<span class="token punctuation">.</span><span class="token function">HasSuffix</span><span class="token punctuation">(</span>mime<span class="token punctuation">,</span> <span class="token string">"+json"</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
			<span class="token keyword">return</span> types<span class="token punctuation">.</span>Json
		<span class="token punctuation">}</span>
		<span class="token keyword">return</span> types<span class="token punctuation">.</span>Generic

	<span class="token keyword">default</span><span class="token punctuation">:</span>
		<span class="token comment">// Mime type not recognized so lets just make it a generic.</span>
		<span class="token keyword">return</span> types<span class="token punctuation">.</span>Generic
	<span class="token punctuation">}</span>

<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="http-user-agent" tabindex="-1"><a class="header-anchor" href="#http-user-agent" aria-hidden="true">#</a> HTTP User Agent</h3>
<p><code v-pre>open</code>'s user agent is the same as <code v-pre>get</code> and <code v-pre>post</code> and is configurable via
<code v-pre>config</code> under they app <strong>http</strong></p>
<pre><code>» config -&gt; [http]
{
    &quot;cookies&quot;: {
        &quot;Data-Type&quot;: &quot;json&quot;,
        &quot;Default&quot;: {
            &quot;example.com&quot;: {
                &quot;name&quot;: &quot;value&quot;
            },
            &quot;www.example.com&quot;: {
                &quot;name&quot;: &quot;value&quot;
            }
        },
        &quot;Description&quot;: &quot;Defined cookies to send, ordered by domain.&quot;,
        &quot;Dynamic&quot;: false,
        &quot;Global&quot;: false,
        &quot;Value&quot;: {
            &quot;example.com&quot;: {
                &quot;name&quot;: &quot;value&quot;
            },
            &quot;www.example.com&quot;: {
                &quot;name&quot;: &quot;value&quot;
            }
        }
    },
    &quot;default-https&quot;: {
        &quot;Data-Type&quot;: &quot;bool&quot;,
        &quot;Default&quot;: false,
        &quot;Description&quot;: &quot;If true then when no protocol is specified (`http://` nor `https://`) then default to `https://`.&quot;,
        &quot;Dynamic&quot;: false,
        &quot;Global&quot;: false,
        &quot;Value&quot;: false
    },
    &quot;headers&quot;: {
        &quot;Data-Type&quot;: &quot;json&quot;,
        &quot;Default&quot;: {
            &quot;example.com&quot;: {
                &quot;name&quot;: &quot;value&quot;
            },
            &quot;www.example.com&quot;: {
                &quot;name&quot;: &quot;value&quot;
            }
        },
        &quot;Description&quot;: &quot;Defined HTTP request headers to send, ordered by domain.&quot;,
        &quot;Dynamic&quot;: false,
        &quot;Global&quot;: false,
        &quot;Value&quot;: {
            &quot;example.com&quot;: {
                &quot;name&quot;: &quot;value&quot;
            },
            &quot;www.example.com&quot;: {
                &quot;name&quot;: &quot;value&quot;
            }
        }
    },
    &quot;insecure&quot;: {
        &quot;Data-Type&quot;: &quot;bool&quot;,
        &quot;Default&quot;: false,
        &quot;Description&quot;: &quot;Ignore certificate errors.&quot;,
        &quot;Dynamic&quot;: false,
        &quot;Global&quot;: false,
        &quot;Value&quot;: false
    },
    &quot;redirect&quot;: {
        &quot;Data-Type&quot;: &quot;bool&quot;,
        &quot;Default&quot;: true,
        &quot;Description&quot;: &quot;Automatically follow redirects.&quot;,
        &quot;Dynamic&quot;: false,
        &quot;Global&quot;: false,
        &quot;Value&quot;: true
    },
    &quot;timeout&quot;: {
        &quot;Data-Type&quot;: &quot;int&quot;,
        &quot;Default&quot;: 10,
        &quot;Description&quot;: &quot;Timeout in seconds for `get` and `getfile`.&quot;,
        &quot;Dynamic&quot;: false,
        &quot;Global&quot;: false,
        &quot;Value&quot;: 10
    },
    &quot;user-agent&quot;: {
        &quot;Data-Type&quot;: &quot;str&quot;,
        &quot;Default&quot;: &quot;murex/1.7.0000 BETA&quot;,
        &quot;Description&quot;: &quot;User agent string for `get` and `getfile`.&quot;,
        &quot;Dynamic&quot;: false,
        &quot;Global&quot;: false,
        &quot;Value&quot;: &quot;murex/1.7.0000 BETA&quot;
    }
}
</code></pre>
<h3 id="open-flags" tabindex="-1"><a class="header-anchor" href="#open-flags" aria-hidden="true">#</a> Open Flags</h3>
<p>If the <code v-pre>open</code> builtin falls back to using the systems default (like <code v-pre>open-xdg</code>)
then the only thing that gets passed is the path being opened. If the path is
stdin then a temporary file will be created. If you want to pass command line
flags to <code v-pre>open-xdg</code> (for example), then you need to call that command directly.
In the case of macOS and some Linux systems, that might look like:</p>
<pre><code>exec open --flags filename
</code></pre>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/types/generic.html"><code v-pre>*</code> (generic) </RouterLink>:
generic (primitive)</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/exec.html"><code v-pre>exec</code></RouterLink>:
Runs an executable</li>
<li><RouterLink to="/commands/fexec.html"><code v-pre>fexec</code> </RouterLink>:
Execute a command or function, bypassing the usual order of precedence.</li>
<li><RouterLink to="/commands/foreach.html"><code v-pre>foreach</code></RouterLink>:
Iterate through an array</li>
<li><RouterLink to="/commands/get.html"><code v-pre>get</code></RouterLink>:
Makes a standard HTTP request and returns the result as a JSON object</li>
<li><RouterLink to="/commands/getfile.html"><code v-pre>getfile</code></RouterLink>:
Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.</li>
<li><RouterLink to="/commands/openagent.html"><code v-pre>openagent</code></RouterLink>:
Creates a handler function for `open</li>
<li><RouterLink to="/commands/out.html"><code v-pre>out</code></RouterLink>:
Print a string to the STDOUT with a trailing new line character</li>
<li><RouterLink to="/commands/post.html"><code v-pre>post</code></RouterLink>:
HTTP POST request with a JSON-parsable return</li>
</ul>
</div></template>


