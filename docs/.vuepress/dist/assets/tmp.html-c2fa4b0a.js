import{_ as o}from"./plugin-vue_export-helper-c27b6911.js";import{r as p,o as c,c as i,d as n,b as t,w as e,e as s,f as l}from"./app-45f7c304.js";const u={},r=l(`<h1 id="tmp" tabindex="-1"><a class="header-anchor" href="#tmp" aria-hidden="true">#</a> <code>tmp</code></h1><blockquote><p>Create a temporary file and write to it</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p><code>tmp</code> creates a temporary file, writes the contents of STDIN to it then returns its filename to STDOUT.</p><p>You can optionally specify a file extension, for example if the temporary file needs to be read by <code>open</code> or an editor which uses extensions to define syntax highlighting.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>\`&lt;stdin&gt;\` -&gt; tmp [ file-extension ] -&gt; \`&lt;stdout&gt;\`
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><pre><code>» out: &quot;Hello, world!&quot; -&gt; set: tmp

» out: $tmp
/var/folders/3t/267q_b0j27d29bnf6pf7m7vm0000gn/T/murex838290600/8ec6936c1ac1c347bf85675eab4a0877-13893

» open: $tmp
Hello, world!
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><p>The temporary file name is a base64 encoded md5 hash of the time plus Murex function ID with Murex process ID appended:</p><div class="language-go line-numbers-mode" data-ext="go"><pre class="language-go"><code><span class="token keyword">package</span> io

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">&quot;crypto/md5&quot;</span>
	<span class="token string">&quot;encoding/hex&quot;</span>
	<span class="token string">&quot;io&quot;</span>
	<span class="token string">&quot;os&quot;</span>
	<span class="token string">&quot;strconv&quot;</span>
	<span class="token string">&quot;time&quot;</span>

	<span class="token string">&quot;github.com/lmorg/murex/lang&quot;</span>
	<span class="token string">&quot;github.com/lmorg/murex/lang/types&quot;</span>
	<span class="token string">&quot;github.com/lmorg/murex/utils/consts&quot;</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">init</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	lang<span class="token punctuation">.</span><span class="token function">DefineMethod</span><span class="token punctuation">(</span><span class="token string">&quot;tmp&quot;</span><span class="token punctuation">,</span> cmdTempFile<span class="token punctuation">,</span> types<span class="token punctuation">.</span>Any<span class="token punctuation">,</span> types<span class="token punctuation">.</span>String<span class="token punctuation">)</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">cmdTempFile</span><span class="token punctuation">(</span>p <span class="token operator">*</span>lang<span class="token punctuation">.</span>Process<span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	p<span class="token punctuation">.</span>Stdout<span class="token punctuation">.</span><span class="token function">SetDataType</span><span class="token punctuation">(</span>types<span class="token punctuation">.</span>String<span class="token punctuation">)</span>

	ext<span class="token punctuation">,</span> <span class="token boolean">_</span> <span class="token operator">:=</span> p<span class="token punctuation">.</span>Parameters<span class="token punctuation">.</span><span class="token function">String</span><span class="token punctuation">(</span><span class="token number">0</span><span class="token punctuation">)</span>
	<span class="token keyword">if</span> ext <span class="token operator">!=</span> <span class="token string">&quot;&quot;</span> <span class="token punctuation">{</span>
		ext <span class="token operator">=</span> <span class="token string">&quot;.&quot;</span> <span class="token operator">+</span> ext
	<span class="token punctuation">}</span>

	fileId <span class="token operator">:=</span> time<span class="token punctuation">.</span><span class="token function">Now</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">String</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token operator">+</span> <span class="token string">&quot;:&quot;</span> <span class="token operator">+</span> strconv<span class="token punctuation">.</span><span class="token function">Itoa</span><span class="token punctuation">(</span><span class="token function">int</span><span class="token punctuation">(</span>p<span class="token punctuation">.</span>Id<span class="token punctuation">)</span><span class="token punctuation">)</span>

	h <span class="token operator">:=</span> md5<span class="token punctuation">.</span><span class="token function">New</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	<span class="token boolean">_</span><span class="token punctuation">,</span> err <span class="token operator">:=</span> h<span class="token punctuation">.</span><span class="token function">Write</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token function">byte</span><span class="token punctuation">(</span>fileId<span class="token punctuation">)</span><span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> err
	<span class="token punctuation">}</span>

	name <span class="token operator">:=</span> consts<span class="token punctuation">.</span>TempDir <span class="token operator">+</span> hex<span class="token punctuation">.</span><span class="token function">EncodeToString</span><span class="token punctuation">(</span>h<span class="token punctuation">.</span><span class="token function">Sum</span><span class="token punctuation">(</span><span class="token boolean">nil</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token operator">+</span> <span class="token string">&quot;-&quot;</span> <span class="token operator">+</span> strconv<span class="token punctuation">.</span><span class="token function">Itoa</span><span class="token punctuation">(</span>os<span class="token punctuation">.</span><span class="token function">Getpid</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token operator">+</span> ext

	file<span class="token punctuation">,</span> err <span class="token operator">:=</span> os<span class="token punctuation">.</span><span class="token function">Create</span><span class="token punctuation">(</span>name<span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> err
	<span class="token punctuation">}</span>

	<span class="token keyword">defer</span> file<span class="token punctuation">.</span><span class="token function">Close</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

	<span class="token boolean">_</span><span class="token punctuation">,</span> err <span class="token operator">=</span> io<span class="token punctuation">.</span><span class="token function">Copy</span><span class="token punctuation">(</span>file<span class="token punctuation">,</span> p<span class="token punctuation">.</span>Stdin<span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> err
	<span class="token punctuation">}</span>

	<span class="token boolean">_</span><span class="token punctuation">,</span> err <span class="token operator">=</span> p<span class="token punctuation">.</span>Stdout<span class="token punctuation">.</span><span class="token function">Write</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token function">byte</span><span class="token punctuation">(</span>name<span class="token punctuation">)</span><span class="token punctuation">)</span>
	<span class="token keyword">return</span> err
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>This should should provide enough distance to run <code>tmp</code> in parallel....should you ever want to.</p><p><code>tmp</code> files are also located inside a unique per-process Murex temp directory which itself is located in the appropriate temp directory for the host OS (eg <code>$TMPDIR</code> on macOS).</p><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,15),d=n("code",null,">>",-1),k=n("code",null,">",-1),m=n("code",null,"open",-1),v=n("code",null,"pipe",-1);function h(b,f){const a=p("RouterLink");return c(),i("div",null,[r,n("ul",null,[n("li",null,[t(a,{to:"/commands/greater-than-greater-than.html"},{default:e(()=>[d,s(" (append file)")]),_:1}),s(": Writes STDIN to disk - appending contents if file already exists")]),n("li",null,[t(a,{to:"/commands/greater-than.html"},{default:e(()=>[k,s(" (truncate file)")]),_:1}),s(": Writes STDIN to disk - overwriting contents if file already exists")]),n("li",null,[t(a,{to:"/commands/open.html"},{default:e(()=>[m]),_:1}),s(": Open a file with a preferred handler")]),n("li",null,[t(a,{to:"/commands/pipe.html"},{default:e(()=>[v]),_:1}),s(": Manage Murex named pipes")])])])}const y=o(u,[["render",h],["__file","tmp.html.vue"]]);export{y as default};
