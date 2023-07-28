import{_ as o}from"./plugin-vue_export-helper-c27b6911.js";import{r as i,o as u,c as l,d as n,b as s,w as a,e,f as p}from"./app-45f7c304.js";const c={},r=p(`<h1 id="open" tabindex="-1"><a class="header-anchor" href="#open" aria-hidden="true">#</a> <code>open</code></h1><blockquote><p>Open a file with a preferred handler</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p><code>open</code> is a smart tool for reading files:</p><ol><li>It will read a file from disk or a HTTP(S) endpoints</li><li>Detect the file type via file extension or HTTP header <code>Content-Type</code></li><li>It intelligently writes to STDOUT</li></ol><ul><li>If STDOUT is a TTY it will perform any transformations to render to the terminal (eg using inlining images)</li><li>If STDOUT is a pipe then it will write a byte stream with the relevant data-type</li></ul><ol start="4"><li>If there are no open handlers then it will fallback to the systems default. eg <code>open</code> (on macOS, Linux), <code>open-xdg</code> (X11), etc.</li></ol><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>open filename[.gz]|uri -&gt; \`&lt;stdout&gt;\`
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><pre><code>» open https://api.github.com/repos/lmorg/murex/issues -&gt; foreach issue { out: &quot;$issue[number]: $issue[title]&quot; }
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="file-extensions" tabindex="-1"><a class="header-anchor" href="#file-extensions" aria-hidden="true">#</a> File Extensions</h3><p>Supported file extensions are listed in <code>config</code> under the app and key names of <strong>shell</strong>, <strong>extensions</strong>.</p><p>Unsupported file extensions are defaulted to generic, <code>*</code>.</p><p>Files with a <code>.gz</code> extension are assumed to be gzipped and thus are are automatically expanded.</p><h3 id="mime-types" tabindex="-1"><a class="header-anchor" href="#mime-types" aria-hidden="true">#</a> MIME Types</h3><p>The <code>Content-Type</code> HTTP header is compared against a list of MIME types, which are stored in <code>config</code> under the app and key names of <strong>shell</strong>, <strong>mime-types</strong>.</p><p>There is a little bit of additional logic to determine the Murex data-type to use should the MIME type not appear in <code>config</code>, as seen in the following code:</p><div class="language-go line-numbers-mode" data-ext="go"><pre class="language-go"><code><span class="token keyword">package</span> lang

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">&quot;regexp&quot;</span>
	<span class="token string">&quot;strings&quot;</span>

	<span class="token string">&quot;github.com/lmorg/murex/lang/types&quot;</span>
<span class="token punctuation">)</span>

<span class="token keyword">var</span> rxMimePrefix <span class="token operator">=</span> regexp<span class="token punctuation">.</span><span class="token function">MustCompile</span><span class="token punctuation">(</span><span class="token string">\`^([-0-9a-zA-Z]+)/.*$\`</span><span class="token punctuation">)</span>

<span class="token comment">// MimeToMurex gets the murex data type for a corresponding MIME</span>
<span class="token keyword">func</span> <span class="token function">MimeToMurex</span><span class="token punctuation">(</span>mimeType <span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token builtin">string</span> <span class="token punctuation">{</span>
	mime <span class="token operator">:=</span> strings<span class="token punctuation">.</span><span class="token function">Split</span><span class="token punctuation">(</span>mimeType<span class="token punctuation">,</span> <span class="token string">&quot;;&quot;</span><span class="token punctuation">)</span><span class="token punctuation">[</span><span class="token number">0</span><span class="token punctuation">]</span>
	mime <span class="token operator">=</span> strings<span class="token punctuation">.</span><span class="token function">TrimSpace</span><span class="token punctuation">(</span>mime<span class="token punctuation">)</span>
	mime <span class="token operator">=</span> strings<span class="token punctuation">.</span><span class="token function">ToLower</span><span class="token punctuation">(</span>mime<span class="token punctuation">)</span>

	<span class="token comment">// Find a direct match. This is only used to pick up edge cases, eg text files used as images.</span>
	dt <span class="token operator">:=</span> mimes<span class="token punctuation">[</span>mime<span class="token punctuation">]</span>
	<span class="token keyword">if</span> dt <span class="token operator">!=</span> <span class="token string">&quot;&quot;</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> dt
	<span class="token punctuation">}</span>

	<span class="token comment">// No direct match found. Fall back to prefix.</span>
	prefix <span class="token operator">:=</span> rxMimePrefix<span class="token punctuation">.</span><span class="token function">FindStringSubmatch</span><span class="token punctuation">(</span>mime<span class="token punctuation">)</span>
	<span class="token keyword">if</span> <span class="token function">len</span><span class="token punctuation">(</span>prefix<span class="token punctuation">)</span> <span class="token operator">!=</span> <span class="token number">2</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> types<span class="token punctuation">.</span>Generic
	<span class="token punctuation">}</span>

	<span class="token keyword">switch</span> prefix<span class="token punctuation">[</span><span class="token number">1</span><span class="token punctuation">]</span> <span class="token punctuation">{</span>
	<span class="token keyword">case</span> <span class="token string">&quot;text&quot;</span><span class="token punctuation">,</span> <span class="token string">&quot;i-world&quot;</span><span class="token punctuation">,</span> <span class="token string">&quot;message&quot;</span><span class="token punctuation">:</span>
		<span class="token keyword">return</span> types<span class="token punctuation">.</span>String

	<span class="token keyword">case</span> <span class="token string">&quot;audio&quot;</span><span class="token punctuation">,</span> <span class="token string">&quot;music&quot;</span><span class="token punctuation">,</span> <span class="token string">&quot;video&quot;</span><span class="token punctuation">,</span> <span class="token string">&quot;image&quot;</span><span class="token punctuation">,</span> <span class="token string">&quot;model&quot;</span><span class="token punctuation">:</span>
		<span class="token keyword">return</span> types<span class="token punctuation">.</span>Binary

	<span class="token keyword">case</span> <span class="token string">&quot;application&quot;</span><span class="token punctuation">:</span>
		<span class="token keyword">if</span> strings<span class="token punctuation">.</span><span class="token function">HasSuffix</span><span class="token punctuation">(</span>mime<span class="token punctuation">,</span> <span class="token string">&quot;+json&quot;</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
			<span class="token keyword">return</span> types<span class="token punctuation">.</span>Json
		<span class="token punctuation">}</span>
		<span class="token keyword">return</span> types<span class="token punctuation">.</span>Generic

	<span class="token keyword">default</span><span class="token punctuation">:</span>
		<span class="token comment">// Mime type not recognized so lets just make it a generic.</span>
		<span class="token keyword">return</span> types<span class="token punctuation">.</span>Generic
	<span class="token punctuation">}</span>

<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="http-user-agent" tabindex="-1"><a class="header-anchor" href="#http-user-agent" aria-hidden="true">#</a> HTTP User Agent</h3><p><code>open</code>&#39;s user agent is the same as <code>get</code> and <code>post</code> and is configurable via <code>config</code> under they app <strong>http</strong></p><pre><code>» config -&gt; [http]
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
        &quot;Description&quot;: &quot;If true then when no protocol is specified (\`http://\` nor \`https://\`) then default to \`https://\`.&quot;,
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
        &quot;Description&quot;: &quot;Timeout in seconds for \`get\` and \`getfile\`.&quot;,
        &quot;Dynamic&quot;: false,
        &quot;Global&quot;: false,
        &quot;Value&quot;: 10
    },
    &quot;user-agent&quot;: {
        &quot;Data-Type&quot;: &quot;str&quot;,
        &quot;Default&quot;: &quot;murex/1.7.0000 BETA&quot;,
        &quot;Description&quot;: &quot;User agent string for \`get\` and \`getfile\`.&quot;,
        &quot;Dynamic&quot;: false,
        &quot;Global&quot;: false,
        &quot;Value&quot;: &quot;murex/1.7.0000 BETA&quot;
    }
}
</code></pre><h3 id="open-flags" tabindex="-1"><a class="header-anchor" href="#open-flags" aria-hidden="true">#</a> Open Flags</h3><p>If the <code>open</code> builtin falls back to using the systems default (like <code>open-xdg</code>) then the only thing that gets passed is the path being opened. If the path is stdin then a temporary file will be created. If you want to pass command line flags to <code>open-xdg</code> (for example), then you need to call that command directly. In the case of macOS and some Linux systems, that might look like:</p><pre><code>exec open --flags filename
</code></pre><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,27),d=n("code",null,"*",-1),m=n("code",null,"config",-1),q=n("code",null,"exec",-1),h=n("code",null,"fexec",-1),k=n("code",null,"foreach",-1),f=n("code",null,"get",-1),g=n("code",null,"getfile",-1),v=n("code",null,"openagent",-1),b=n("code",null,"out",-1),y=n("code",null,"post",-1);function x(w,T){const t=i("RouterLink");return u(),l("div",null,[r,n("ul",null,[n("li",null,[s(t,{to:"/types/generic.html"},{default:a(()=>[d,e(" (generic) ")]),_:1}),e(": generic (primitive)")]),n("li",null,[s(t,{to:"/commands/config.html"},{default:a(()=>[m]),_:1}),e(": Query or define Murex runtime settings")]),n("li",null,[s(t,{to:"/commands/exec.html"},{default:a(()=>[q]),_:1}),e(": Runs an executable")]),n("li",null,[s(t,{to:"/commands/fexec.html"},{default:a(()=>[h]),_:1}),e(": Execute a command or function, bypassing the usual order of precedence.")]),n("li",null,[s(t,{to:"/commands/foreach.html"},{default:a(()=>[k]),_:1}),e(": Iterate through an array")]),n("li",null,[s(t,{to:"/commands/get.html"},{default:a(()=>[f]),_:1}),e(": Makes a standard HTTP request and returns the result as a JSON object")]),n("li",null,[s(t,{to:"/commands/getfile.html"},{default:a(()=>[g]),_:1}),e(": Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.")]),n("li",null,[s(t,{to:"/commands/openagent.html"},{default:a(()=>[v]),_:1}),e(": Creates a handler function for `open")]),n("li",null,[s(t,{to:"/commands/out.html"},{default:a(()=>[b]),_:1}),e(": Print a string to the STDOUT with a trailing new line character")]),n("li",null,[s(t,{to:"/commands/post.html"},{default:a(()=>[y]),_:1}),e(": HTTP POST request with a JSON-parsable return")])])])}const M=o(c,[["render",x],["__file","open.html.vue"]]);export{M as default};
