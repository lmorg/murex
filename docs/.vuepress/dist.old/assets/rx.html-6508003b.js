import{_ as s}from"./plugin-vue_export-helper-c27b6911.js";import{r as n,o as i,c as d,d as e,b as o,w as r,e as t,f as c}from"./app-45f7c304.js";const h={},l=c(`<h1 id="rx" tabindex="-1"><a class="header-anchor" href="#rx" aria-hidden="true">#</a> <code>rx</code></h1><blockquote><p>Regexp pattern matching for file system objects (eg <code>.*\\\\.txt</code>)</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>Returns a list of files and directories that match a regexp pattern.</p><p>Output is a JSON list.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>rx: pattern -&gt; \`&lt;stdout&gt;\`

!rx: pattern -&gt; \`&lt;stdout&gt;\`

\`&lt;stdin&gt;\` -&gt; rx: pattern -&gt; \`&lt;stdout&gt;\`

\`&lt;stdin&gt;\` -&gt; !rx: pattern -&gt; \`&lt;stdout&gt;\`
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><p>Inline regex file matching:</p><pre><code>cat: @{ rx: &#39;.*\\.txt&#39; }
</code></pre><p>Writing a list of files to disk:</p><pre><code>rx: &#39;.*\\.go&#39; |&gt; filelist.txt
</code></pre><p>Checking if files exist:</p><pre><code>if { rx: somefiles.* } then {
    # files exist
}
</code></pre><p>Checking if files do not exist:</p><pre><code>!if { rx: somefiles.* } then {
    # files do not exist
}
</code></pre><p>Return all files apart from text files:</p><pre><code>!g: &#39;\\.txt$&#39;
</code></pre><p>Filtering a file list based on regexp matches file:</p><pre><code>f: +f -&gt; rx: &#39;.*\\.txt&#39;
</code></pre><p>Remove any regexp file matches from a file list:</p><pre><code>f: +f -&gt; !rx: &#39;.*\\.txt&#39;
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="traversing-directories" tabindex="-1"><a class="header-anchor" href="#traversing-directories" aria-hidden="true">#</a> Traversing Directories</h3><p>Unlike globbing (<code>g</code>) which can traverse directories (eg <code>g: /path/*</code>), <code>rx</code> is only designed to match file system objects in the current working directory.</p><p><code>rx</code> uses Go (lang)&#39;s standard regexp engine.</p><h3 id="inverse-matches" tabindex="-1"><a class="header-anchor" href="#inverse-matches" aria-hidden="true">#</a> Inverse Matches</h3><p>If you want to exclude any matches based on wildcards, rather than include them, then you can use the bang prefix. eg</p><pre><code>» rx: READ*
[
    &quot;README.md&quot;
]

murex-dev» !rx: .*
Error in \`!rx\` (1,1): No data returned.
</code></pre><h3 id="when-used-as-a-method" tabindex="-1"><a class="header-anchor" href="#when-used-as-a-method" aria-hidden="true">#</a> When Used As A Method</h3><p><code>!rx</code> first looks for files that match its pattern, then it reads the file list from STDIN. If STDIN contains contents that are not files then <code>!rx</code> might not handle those list items correctly. This shouldn&#39;t be an issue with <code>rx</code> in its normal mode because it is only looking for matches however when used as <code>!rx</code> any items that are not files will leak through.</p><p>This is its designed feature and not a bug. If you wish to remove anything that also isn&#39;t a file then you should first pipe into either <code>g: *</code>, <code>rx: .*</code>, or <code>f +f</code> and then pipe that into <code>!rx</code>.</p><p>The reason for this behavior is to separate this from <code>!regexp</code> and <code>!match</code>.</p><h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2><ul><li><code>rx</code></li><li><code>!rx</code></li></ul><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,36),p=e("code",null,"f",-1),f=e("code",null,"g",-1),x=e("code",null,"*.txt",-1),u=e("code",null,"match",-1),m=e("code",null,"regexp",-1);function g(_,b){const a=n("RouterLink");return i(),d("div",null,[l,e("ul",null,[e("li",null,[o(a,{to:"/commands/f.html"},{default:r(()=>[p]),_:1}),t(": Lists or filters file system objects (eg files)")]),e("li",null,[o(a,{to:"/commands/g.html"},{default:r(()=>[f]),_:1}),t(": Glob pattern matching for file system objects (eg "),x,t(")")]),e("li",null,[o(a,{to:"/commands/match.html"},{default:r(()=>[u]),_:1}),t(": Match an exact value in an array")]),e("li",null,[o(a,{to:"/commands/regexp.html"},{default:r(()=>[m]),_:1}),t(": Regexp tools for arrays / lists of strings")])])])}const k=s(h,[["render",g],["__file","rx.html.vue"]]);export{k as default};
