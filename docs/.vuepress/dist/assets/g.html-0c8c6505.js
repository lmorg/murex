import{_ as s}from"./plugin-vue_export-helper-c27b6911.js";import{r as i,o as r,c as d,d as e,b as o,w as n,e as t,f as c}from"./app-45f7c304.js";const l={},h=c(`<h1 id="g" tabindex="-1"><a class="header-anchor" href="#g" aria-hidden="true">#</a> <code>g</code></h1><blockquote><p>Glob pattern matching for file system objects (eg <code>*.txt</code>)</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>Returns a list of files and directories that match a glob pattern.</p><p>Output is a JSON list.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>g: pattern -&gt; \`&lt;stdout&gt;\`

[ \`&lt;stdin&gt;\` -&gt; ] @g command pattern [ -&gt; \`&lt;stdout&gt;\` ]

!g: pattern -&gt; \`&lt;stdout&gt;\`

\`&lt;stdin&gt;\` -&gt; g: pattern -&gt; \`&lt;stdout&gt;\`

\`&lt;stdin&gt;\` -&gt; !g: pattern -&gt; \`&lt;stdout&gt;\`
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><p>Inline globbing:</p><pre><code>cat: @{ g: *.txt }
</code></pre><p>Writing a JSON array of files to disk:</p><pre><code>g: *.txt |&gt; filelist.json
</code></pre><p>Writing a list of files to disk:</p><pre><code>g: *.txt -&gt; format str |&gt; filelist.txt
</code></pre><p>Checking if a file exists:</p><pre><code>if { g: somefile.txt } then {
    # file exists
}
</code></pre><p>Checking if a file does not exist:</p><pre><code>!if { g: somefile.txt } then {
    # file does not exist
}
</code></pre><p>Return all files apart from text files:</p><pre><code>!g: *.txt
</code></pre><p>Filtering a file list based on glob matches:</p><pre><code>f: +f -&gt; g: *.md
</code></pre><p>Remove any glob matches from a file list:</p><pre><code>f: +f -&gt; !g: *.md
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="pattern-reference" tabindex="-1"><a class="header-anchor" href="#pattern-reference" aria-hidden="true">#</a> Pattern Reference</h3><ul><li><code>*</code> matches any number of (including zero) characters</li><li><code>?</code> matches any single character</li></ul><h3 id="inverse-matches" tabindex="-1"><a class="header-anchor" href="#inverse-matches" aria-hidden="true">#</a> Inverse Matches</h3><p>If you want to exclude any matches based on wildcards, rather than include them, then you can use the bang prefix. eg</p><pre><code>» g: READ*
[
    &quot;README.md&quot;
]

» !g: *
Error in \`!g\` (1,1): No data returned.
</code></pre><h3 id="when-used-as-a-method" tabindex="-1"><a class="header-anchor" href="#when-used-as-a-method" aria-hidden="true">#</a> When Used As A Method</h3><p><code>!g</code> first looks for files that match its pattern, then it reads the file list from STDIN. If STDIN contains contents that are not files then <code>!g</code> might not handle those list items correctly. This shouldn&#39;t be an issue with <code>frx</code> in its normal mode because it is only looking for matches however when used as <code>!g</code> any items that are not files will leak through.</p><p>This is its designed feature and not a bug. If you wish to remove anything that also isn&#39;t a file then you should first pipe into either <code>g: *</code>, <code>rx: .*</code>, or <code>f +f</code> and then pipe that into <code>!g</code>.</p><p>The reason for this behavior is to separate this from <code>!regexp</code> and <code>!match</code>.</p><h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2><ul><li><code>g</code></li><li><code>!g</code></li></ul><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,37),f=e("code",null,"f",-1),p=e("code",null,"match",-1),g=e("code",null,"regexp",-1),u=e("code",null,"rx",-1),m=e("code",null,".*\\\\.txt",-1);function x(b,_){const a=i("RouterLink");return r(),d("div",null,[h,e("ul",null,[e("li",null,[o(a,{to:"/commands/f.html"},{default:n(()=>[f]),_:1}),t(": Lists or filters file system objects (eg files)")]),e("li",null,[o(a,{to:"/commands/match.html"},{default:n(()=>[p]),_:1}),t(": Match an exact value in an array")]),e("li",null,[o(a,{to:"/commands/regexp.html"},{default:n(()=>[g]),_:1}),t(": Regexp tools for arrays / lists of strings")]),e("li",null,[o(a,{to:"/commands/rx.html"},{default:n(()=>[u]),_:1}),t(": Regexp pattern matching for file system objects (eg "),m,t(")")])])])}const v=s(l,[["render",x],["__file","g.html.vue"]]);export{v as default};
