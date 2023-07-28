import{_ as i}from"./plugin-vue_export-helper-c27b6911.js";import{r,o as l,c as s,d as e,b as t,w as n,e as a,f as d}from"./app-45f7c304.js";const c={},h=d(`<h1 id="export" tabindex="-1"><a class="header-anchor" href="#export" aria-hidden="true">#</a> <code>export</code></h1><blockquote><p>Define an environmental variable and set it&#39;s value</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>Defines, updates or deallocates an environmental variable.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>\`&lt;stdin&gt;\` -&gt; export var_name

export var_name=data
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><p>As a method:</p><pre><code>» out &quot;Hello, world!&quot; -&gt; export hw
» out &quot;$hw&quot;
Hello, World!
</code></pre><p>As a function:</p><pre><code>» export hw=&quot;Hello, world!&quot;
» out &quot;$hw&quot;
Hello, World!
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="deallocation" tabindex="-1"><a class="header-anchor" href="#deallocation" aria-hidden="true">#</a> Deallocation</h3><p>You can unset variable names with the bang prefix:</p><pre><code>!export var_name
</code></pre><p>For compatibility with other shells, <code>unset</code> is also supported but it&#39;s really not an idiomatic method of deallocation since it&#39;s name is misleading and suggests it is a deallocator for local Murex variables defined via <code>set</code>.</p><h3 id="exporting-a-local-or-global-variable" tabindex="-1"><a class="header-anchor" href="#exporting-a-local-or-global-variable" aria-hidden="true">#</a> Exporting a local or global variable</h3><p>You can also export a local or global variable of the same name by specifying that variable name without a following value. For example</p><pre><code># Create a local variable called &#39;foo&#39;:
» set: foo=bar
» env -&gt; grep: foo

# Export that local variable as an environmental variable:
» export: foo
» env -&gt; grep: foo
foo=bar

# Changing the value of the local variable doesn&#39;t alter the value of the environmental variable:
» set: foo=rab
» env -&gt; grep: foo
foo=bar
» out: $foo
rab
</code></pre><h3 id="type-annotations" tabindex="-1"><a class="header-anchor" href="#type-annotations" aria-hidden="true">#</a> Type Annotations</h3><p>When <code>set</code> or <code>global</code> are used as a function, the parameters are passed as a string which means the variables are defined as a <code>str</code>. If you wish to define them as an alternate data type then you should add type annotations:</p><pre><code>» set: int age = 30
(\`$age\` is an integer, \`int\`)

» global: bool dark_theme = true
</code></pre><p>(<code>$dark_theme</code> is a boolean, <code>bool</code>)</p><p>When using <code>set</code> or <code>global</code> as a method, by default they will define the variable as the data type of the pipe:</p><pre><code>» open: example.json -&gt; set: file
</code></pre><p>(<code>$file</code> is defined a <code>json</code> type because <code>open</code> wrote to <code>set</code>&#39;s pipe with a <code>json</code> type)</p><p>You can also annotate <code>set</code> and <code>global</code> when used as a method too:</p><pre><code>out: 30 -&gt; set: int age
</code></pre><p>(<code>$age</code> is an integer, <code>int</code>, despite <code>out</code> writing a string, \`str, to the pipe)</p><blockquote><p><code>export</code> does not support type annotations because environmental variables must always be strings. This is a limitation of the current operating systems.</p></blockquote><h3 id="scoping" tabindex="-1"><a class="header-anchor" href="#scoping" aria-hidden="true">#</a> Scoping</h3><p>Variable scoping is simplified to three layers:</p><ol><li>Local variables (<code>set</code>, <code>!set</code>, <code>let</code>)</li><li>Global variables (<code>global</code>, <code>!global</code>)</li><li>Environmental variables (<code>export</code>, <code>!export</code>, <code>unset</code>)</li></ol><p>Variables are looked up in that order of too. For example a the following code where <code>set</code> overrides both the global and environmental variable:</p><pre><code>» set:    foobar=1
» global: foobar=2
» export: foobar=3
» out: $foobar
1
</code></pre><h4 id="local-variables" tabindex="-1"><a class="header-anchor" href="#local-variables" aria-hidden="true">#</a> Local variables</h4><p>These are defined via <code>set</code> and <code>let</code>. They&#39;re variables that are persistent across any blocks within a function. Functions will typically be blocks encapsulated like so:</p><pre><code>function example {
    # variables scoped inside here
}
</code></pre><p>...or...</p><pre><code>private example {
    # variables scoped inside here
}
</code></pre><p>...however dynamic autocompletes, events, unit tests and any blocks defined in <code>config</code> will also be triggered as functions.</p><p>Code running inside any control flow or error handing structures will be treated as part of the same part of the same scope as the parent function:</p><pre><code>» function example {
»     try {
»         # set &#39;foobar&#39; inside a \`try\` block
»         set: foobar=example
»     }
»     # &#39;foobar&#39; exists outside of \`try\` because it is scoped to \`function\`
»     out: $foobar
» }
example
</code></pre><p>Where this behavior might catch you out is with iteration blocks which create variables, eg <code>for</code>, <code>foreach</code> and <code>formap</code>. Any variables created inside them are still shared with any code outside of those structures but still inside the function block.</p><p>Any local variables are only available to that function. If a variable is defined in a parent function that goes on to call child functions, then those local variables are not inherited but the child functions:</p><pre><code>» function parent {
»     # set a local variable
»     set: foobar=example
»     child
» }
»
» function child {
»     # returns the \`global\` value, &quot;not set&quot;, because the local \`set\` isn&#39;t inherited
»     out: $foobar
» }
»
» global: $foobar=&quot;not set&quot;
» parent
not set
</code></pre><p>It&#39;s also worth remembering that any variable defined using <code>set</code> in the shells FID (ie in the interactive shell) is localised to structures running in the interactive, REPL, shell and are not inherited by any called functions.</p><h4 id="global-variables" tabindex="-1"><a class="header-anchor" href="#global-variables" aria-hidden="true">#</a> Global variables</h4><p>Where <code>global</code> differs from <code>set</code> is that the variables defined with <code>global</code> will be scoped at the global shell level (please note this is not the same as environmental variables!) so will cascade down through all scoped code-blocks including those running in other threads.</p><h4 id="environmental-variables" tabindex="-1"><a class="header-anchor" href="#environmental-variables" aria-hidden="true">#</a> Environmental variables</h4><p>Exported variables (defined via <code>export</code>) are system environmental variables. Inside Murex environmental variables behave much like <code>global</code> variables however their real purpose is passing data to external processes. For example <code>env</code> is an external process on Linux (eg <code>/usr/bin/env</code> on ArchLinux):</p><pre><code>» export foo=bar
» env -&gt; grep foo
foo=bar
</code></pre><h3 id="function-names" tabindex="-1"><a class="header-anchor" href="#function-names" aria-hidden="true">#</a> Function Names</h3><p>As a security feature function names cannot include variables. This is done to reduce the risk of code executing by mistake due to executables being hidden behind variable names.</p><p>Instead Murex will assume you want the output of the variable printed:</p><pre><code>» out &quot;Hello, world!&quot; -&gt; set hw
» $hw
Hello, world!
</code></pre><p>On the rare occasions you want to force variables to be expanded inside a function name, then call that function via <code>exec</code>:</p><pre><code>» set cmd=grep
» ls -&gt; exec: $cmd main.go
main.go
</code></pre><p>This only works for external executables. There is currently no way to call aliases, functions nor builtins from a variable and even the above <code>exec</code> trick is considered bad form because it reduces the readability of your shell scripts.</p><h3 id="usage-inside-quotation-marks" tabindex="-1"><a class="header-anchor" href="#usage-inside-quotation-marks" aria-hidden="true">#</a> Usage Inside Quotation Marks</h3><p>Like with Bash, Perl and PHP: Murex will expand the variable when it is used inside a double quotes but will escape the variable name when used inside single quotes:</p><pre><code>» out &quot;$foo&quot;
bar

» out &#39;$foo&#39;
$foo

» out %($foo)
bar
</code></pre><h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2><ul><li><code>export</code></li><li><code>!export</code></li><li><code>unset</code></li></ul><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,65),p=e("code",null,"(",-1),u=e("code",null,"=",-1),b=e("code",null,"expr",-1),f=e("code",null,"global",-1),m=e("code",null,"let",-1),v=e("code",null,"set",-1);function g(x,w){const o=r("RouterLink");return l(),s("div",null,[h,e("ul",null,[e("li",null,[t(o,{to:"/user-guide/reserved-vars.html"},{default:n(()=>[a("Reserved Variables")]),_:1}),a(": Special variables reserved by Murex")]),e("li",null,[t(o,{to:"/user-guide/scoping.html"},{default:n(()=>[a("Variable and Config Scoping")]),_:1}),a(": How scoping works within Murex")]),e("li",null,[t(o,{to:"/commands/brace-quote.html"},{default:n(()=>[p,a(" (brace quote)")]),_:1}),a(": Write a string to the STDOUT without new line")]),e("li",null,[t(o,{to:"/commands/equ.html"},{default:n(()=>[u,a(" (arithmetic evaluation)")]),_:1}),a(": Evaluate a mathematical function (deprecated)")]),e("li",null,[t(o,{to:"/commands/expr.html"},{default:n(()=>[b]),_:1}),a(": Expressions: mathematical, string comparisons, logical operators")]),e("li",null,[t(o,{to:"/commands/global.html"},{default:n(()=>[f]),_:1}),a(": Define a global variable and set it's value")]),e("li",null,[t(o,{to:"/commands/let.html"},{default:n(()=>[m]),_:1}),a(": Evaluate a mathematical function and assign to variable (deprecated)")]),e("li",null,[t(o,{to:"/commands/set.html"},{default:n(()=>[v]),_:1}),a(": Define a local variable and set it's value")])])])}const k=i(c,[["render",g],["__file","export.html.vue"]]);export{k as default};
