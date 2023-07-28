import{_ as l}from"./plugin-vue_export-helper-c27b6911.js";import{r as s,o as r,c as u,d as e,e as t,b as o,w as a,f as d}from"./app-45f7c304.js";const c={},h=d(`<h1 id="select" tabindex="-1"><a class="header-anchor" href="#select" aria-hidden="true">#</a> <code>select</code></h1><blockquote><p>Inlining SQL into shell pipelines</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p><code>select</code> imports tabulated data into an in memory sqlite3 database and executes SQL queries against the data. It returns a table of the same data type as the input type</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>\`&lt;stdin&gt;\` -&gt; select * | ... WHERE ... -&gt; \`&lt;stdout&gt;\`

select * | ... FROM file[.gz] WHERE ... -&gt; \`&lt;stdout&gt;\`
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><p>List a count of all the processes running against each user ID:</p><pre><code>» ps aux -&gt; select count(*), user GROUP BY user ORDER BY 1
count(*) USER
1       _analyticsd
1       _applepay
1       _atsserver
1       _captiveagent
1       _cmiodalassistants
1       _ctkd
1       _datadetectors
1       _displaypolicyd
1       _distnote
1       _gamecontrollerd
1       _hidd
1       _iconservices
1       _installcoordinationd
1       _mdnsresponder
1       _netbios
1       _networkd
1       _reportmemoryexception
1       _timed
1       _usbmuxd
2       _appleevents
3       _assetcache
3       _fpsd
3       _nsurlsessiond
3       _softwareupdate
4       _windowserver
5       _coreaudiod
6       _spotlight
7       _locationd
144     root
308     foobar


select count(*)
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="default-table-name" tabindex="-1"><a class="header-anchor" href="#default-table-name" aria-hidden="true">#</a> Default Table Name</h3><p>The table created is called <code>main</code>, however you do not need to include a <code>FROM</code> condition in your SQL as Murex will inject <code>FROM main</code> into your SQL if it is missing. In fact, it is recommended that you exclude <code>FROM</code> from your SQL queries for the sake of brevity.</p><h3 id="config-options" tabindex="-1"><a class="header-anchor" href="#config-options" aria-hidden="true">#</a> <code>config</code> Options</h3><p><code>select</code>&#39;s behavior is configurable:</p><pre><code>» config -&gt; [ select ]
{
    &quot;fail-irregular-columns&quot;: {
        &quot;Data-Type&quot;: &quot;bool&quot;,
        &quot;Default&quot;: false,
        &quot;Description&quot;: &quot;When importing a table into sqlite3, fail if there is an irregular number of columns&quot;,
        &quot;Dynamic&quot;: false,
        &quot;Global&quot;: false,
        &quot;Value&quot;: false
    },
    &quot;merge-trailing-columns&quot;: {
        &quot;Data-Type&quot;: &quot;bool&quot;,
        &quot;Default&quot;: true,
        &quot;Description&quot;: &quot;When importing a table into sqlite3, if \`fail-irregular-columns\` is set to \`false\` and there are more columns than headings, then any additional columns are concatenated into the last column (space delimitated). If \`merge-trailing-columns\` is set to \`false\` then any trailing columns are ignored&quot;,
        &quot;Dynamic&quot;: false,
        &quot;Global&quot;: false,
        &quot;Value&quot;: true
    },
    &quot;print-headings&quot;: {
        &quot;Data-Type&quot;: &quot;bool&quot;,
        &quot;Default&quot;: true,
        &quot;Description&quot;: &quot;Print headings when writing results&quot;,
        &quot;Dynamic&quot;: false,
        &quot;Global&quot;: false,
        &quot;Value&quot;: true
    },
    &quot;table-includes-headings&quot;: {
        &quot;Data-Type&quot;: &quot;bool&quot;,
        &quot;Default&quot;: true,
        &quot;Description&quot;: &quot;When importing a table into sqlite3, treat the first row as headings (if \`false\`, headings are Excel style column references starting at \`A\`)&quot;,
        &quot;Dynamic&quot;: false,
        &quot;Global&quot;: false,
        &quot;Value&quot;: true
    }
}
</code></pre><p>(See below for how to use <code>config</code>)</p><h3 id="read-all-vs-sequential-reads" tabindex="-1"><a class="header-anchor" href="#read-all-vs-sequential-reads" aria-hidden="true">#</a> Read All vs Sequential Reads</h3><p>At present, <code>select</code> only supports reading the entire table from STDIN before importing that data into sqlite3. There is some prototype code being written to support sequential imports but this is hugely experimental and not yet enabled.</p><p>This might make <code>select</code> unsuitable for large datasets.</p><h3 id="early-release" tabindex="-1"><a class="header-anchor" href="#early-release" aria-hidden="true">#</a> Early Release</h3><p>This is a very early release so there almost certainly will be bugs hiding. Which is another reason why this is currently only an optional builtin.</p>`,21),p={href:"https://github.com/lmorg/murex/issues",target:"_blank",rel:"noopener noreferrer"},f=e("h2",{id:"synonyms",tabindex:"-1"},[e("a",{class:"header-anchor",href:"#synonyms","aria-hidden":"true"},"#"),t(" Synonyms")],-1),q=e("ul",null,[e("li",null,[e("code",null,"select")])],-1),m=e("h2",{id:"see-also",tabindex:"-1"},[e("a",{class:"header-anchor",href:"#see-also","aria-hidden":"true"},"#"),t(" See Also")],-1),g=e("code",null,"*",-1),_=e("code",null,"config",-1),b=e("code",null,"csv",-1),y=e("code",null,"jsonl",-1),x=e("code",null,"config",-1);function D(v,w){const i=s("ExternalLinkIcon"),n=s("RouterLink");return r(),u("div",null,[h,e("p",null,[t("If you do run into any issues then please raise them on "),e("a",p,[t("Github"),o(i)]),t(".")]),f,q,m,e("ul",null,[e("li",null,[o(n,{to:"/types/generic.html"},{default:a(()=>[g,t(" (generic) ")]),_:1}),t(": generic (primitive)")]),e("li",null,[o(n,{to:"/commands/config.html"},{default:a(()=>[_]),_:1}),t(": Query or define Murex runtime settings")]),e("li",null,[o(n,{to:"/types/csv.html"},{default:a(()=>[b]),_:1}),t(": CSV files (and other character delimited tables)")]),e("li",null,[o(n,{to:"/types/jsonl.html"},{default:a(()=>[y]),_:1}),t(": JSON Lines")]),e("li",null,[o(n,{to:"/changelog/v2.1.html"},{default:a(()=>[t("v2.1")]),_:1}),t(": This release comes with support for inlining SQL and some major bug fixes plus a breaking change for "),x,t(". Please read for details.")])])])}const k=l(c,[["render",D],["__file","select.html.vue"]]);export{k as default};
