<template><div><h1 id="select" tabindex="-1"><a class="header-anchor" href="#select" aria-hidden="true">#</a> <code v-pre>select</code></h1>
<blockquote>
<p>Inlining SQL into shell pipelines</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>select</code> imports tabulated data into an in memory sqlite3 database and
executes SQL queries against the data. It returns a table of the same
data type as the input type</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>`&lt;stdin&gt;` -&gt; select * | ... WHERE ... -&gt; `&lt;stdout&gt;`

select * | ... FROM file[.gz] WHERE ... -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>List a count of all the processes running against each user ID:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» ps aux -> select count(*), user GROUP BY user ORDER BY 1
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
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="default-table-name" tabindex="-1"><a class="header-anchor" href="#default-table-name" aria-hidden="true">#</a> Default Table Name</h3>
<p>The table created is called <code v-pre>main</code>, however you do not need to include a <code v-pre>FROM</code>
condition in your SQL as Murex will inject <code v-pre>FROM main</code> into your SQL if it is
missing. In fact, it is recommended that you exclude <code v-pre>FROM</code> from your SQL
queries for the sake of brevity.</p>
<h3 id="config-options" tabindex="-1"><a class="header-anchor" href="#config-options" aria-hidden="true">#</a> <code v-pre>config</code> Options</h3>
<p><code v-pre>select</code>'s behavior is configurable:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» config -> [ select ]
{
    "fail-irregular-columns": {
        "Data-Type": "bool",
        "Default": false,
        "Description": "When importing a table into sqlite3, fail if there is an irregular number of columns",
        "Dynamic": false,
        "Global": false,
        "Value": false
    },
    "merge-trailing-columns": {
        "Data-Type": "bool",
        "Default": true,
        "Description": "When importing a table into sqlite3, if `fail-irregular-columns` is set to `false` and there are more columns than headings, then any additional columns are concatenated into the last column (space delimitated). If `merge-trailing-columns` is set to `false` then any trailing columns are ignored",
        "Dynamic": false,
        "Global": false,
        "Value": true
    },
    "print-headings": {
        "Data-Type": "bool",
        "Default": true,
        "Description": "Print headings when writing results",
        "Dynamic": false,
        "Global": false,
        "Value": true
    },
    "table-includes-headings": {
        "Data-Type": "bool",
        "Default": true,
        "Description": "When importing a table into sqlite3, treat the first row as headings (if `false`, headings are Excel style column references starting at `A`)",
        "Dynamic": false,
        "Global": false,
        "Value": true
    }
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>(See below for how to use <code v-pre>config</code>)</p>
<h3 id="read-all-vs-sequential-reads" tabindex="-1"><a class="header-anchor" href="#read-all-vs-sequential-reads" aria-hidden="true">#</a> Read All vs Sequential Reads</h3>
<p>At present, <code v-pre>select</code> only supports reading the entire table from STDIN before
importing that data into sqlite3. There is some prototype code being written to
support sequential imports but this is hugely experimental and not yet enabled.</p>
<p>This might make <code v-pre>select</code> unsuitable for large datasets.</p>
<h3 id="early-release" tabindex="-1"><a class="header-anchor" href="#early-release" aria-hidden="true">#</a> Early Release</h3>
<p>This is a very early release so there almost certainly will be bugs hiding.
Which is another reason why this is currently only an optional builtin.</p>
<p>If you do run into any issues then please raise them on <a href="https://github.com/lmorg/murex/issues" target="_blank" rel="noopener noreferrer">Github<ExternalLinkIcon/></a>.</p>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>select</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/types/generic.html"><code v-pre>*</code> (generic) </RouterLink>:
generic (primitive)</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/types/csv.html"><code v-pre>csv</code> </RouterLink>:
CSV files (and other character delimited tables)</li>
<li><RouterLink to="/types/jsonl.html"><code v-pre>jsonl</code> </RouterLink>:
JSON Lines</li>
<li><RouterLink to="/changelog/v2.1.html">v2.1</RouterLink>:
This release comes with support for inlining SQL and some major bug fixes plus a breaking change for <code v-pre>config</code>. Please read for details.</li>
</ul>
</div></template>


