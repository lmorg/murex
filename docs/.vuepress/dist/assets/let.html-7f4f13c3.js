import{_ as d}from"./plugin-vue_export-helper-c27b6911.js";import{r as i,o as l,c,d as e,e as o,b as a,w as n,f as r}from"./app-45f7c304.js";const h={},u=r('<h1 id="let" tabindex="-1"><a class="header-anchor" href="#let" aria-hidden="true">#</a> <code>let</code></h1><blockquote><p>Evaluate a mathematical function and assign to variable (deprecated)</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p><code>let</code> evaluates a mathematical function and then assigns it to a locally scoped variable (like <code>set</code>)</p>',4),p=e("code",null,"expr",-1),b=r(`<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>let var_name=evaluation

let var_name++

let var_name--
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><pre><code>» let: age=18
» $age
18

» let: age++
» $age
19

» let: under18=age&lt;18
» $under18
false

» let: under21 = age &lt; 21
» $under21
true
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="other-operators" tabindex="-1"><a class="header-anchor" href="#other-operators" aria-hidden="true">#</a> Other Operators</h3><p><code>let</code> also supports the following operators (substitute <strong>VAR</strong> with your variable name, and <strong>NUM</strong> with a number):</p><ul><li><code>VAR--</code>, subtract 1 from VAR</li><li><code>VAR++</code>, add 1 to VAR</li><li><code>VAR -= NUM</code>, subtract NUM from VAR</li><li><code>VAR += NUM</code>, add NUM to VAR</li><li><code>VAR /= NUM</code>, divide VAR by NUM</li><li><code>VAR *= NUM</code>, multiply VAR by NUM</li></ul><p>eg</p><pre><code>» let: i=0
» let: i++
» $i
1

» let: i+=8
» $i
9

» let: i/=3
3
</code></pre><p>Please note these operators are not supported by <code>=</code>.</p><h3 id="variables" tabindex="-1"><a class="header-anchor" href="#variables" aria-hidden="true">#</a> Variables</h3><p>There are two ways you can use variables with the math functions. Either by string interpolation like you would normally with any other function, or directly by name.</p><p>String interpolation:</p><pre><code>» set abc=123
» = $abc==123
true
</code></pre><p>Directly by name:</p><pre><code>» set abc=123
» = abc==123
false
</code></pre><p>To understand the difference between the two, you must first understand how string interpolation works; which is where the parser tokenised the parameters like so</p><pre><code>command line: = $abc==123
token 1: command (name: &quot;=&quot;)
token 2: parameter 1, string (content: &quot;&quot;)
token 3: parameter 1, variable (name: &quot;abc&quot;)
token 4: parameter 1, string (content: &quot;==123&quot;)
</code></pre><p>Then when the command line gets executed, the parameters are compiled on demand similarly to this crude pseudo-code</p><pre><code>command: &quot;=&quot;
parameters 1: concatenate(&quot;&quot;, GetValue(abc), &quot;==123&quot;)
output: &quot;=&quot; &quot;123==123&quot;
</code></pre><p>Thus the actual command getting run is literally <code>123==123</code> due to the variable being replace <strong>before</strong> the command executes.</p><p>Whereas when you call the variable by name it&#39;s up to <code>=</code> or <code>let</code> to do the variable substitution.</p><pre><code>command line: = abc==123
token 1: command (name: &quot;=&quot;)
token 2: parameter 1, string (content: &quot;abc==123&quot;)

command: &quot;=&quot;
parameters 1: concatenate(&quot;abc==123&quot;)
output: &quot;=&quot; &quot;abc==123&quot;
</code></pre><p>The main advantage (or disadvantage, depending on your perspective) of using variables this way is that their data-type is preserved.</p><pre><code>» set str abc=123
» = abc==123
false

» set int abc=123
» = abc==123
true
</code></pre><p>Unfortunately is one of the biggest areas in Murex where you&#39;d need to be careful. The simple addition or omission of the dollar prefix, <code>$</code>, can change the behavior of <code>=</code> and <code>let</code>.</p><h3 id="strings" tabindex="-1"><a class="header-anchor" href="#strings" aria-hidden="true">#</a> Strings</h3><p>Because the usual Murex tools for encapsulating a string (<code>&quot;</code>, <code>&#39;</code> and <code>()</code>) are interpreted by the shell language parser, it means we need a new token for handling strings inside <code>=</code> and <code>let</code>. This is where backtick comes to our rescue.</p><pre><code>» set str abc=123
» = abc==\`123\`
true
</code></pre><p>Please be mindful that if you use string interpolation then you will need to instruct <code>=</code> and <code>let</code> that your field is a string</p><pre><code>» set str abc=123
» = \`$abc\`==\`123\`
true
</code></pre><h3 id="best-practice-recommendation" tabindex="-1"><a class="header-anchor" href="#best-practice-recommendation" aria-hidden="true">#</a> Best practice recommendation</h3><p>As you can see from the sections above, string interpolation offers us some conveniences when comparing variables of differing data-types, such as a <code>str</code> type with a number (eg <code>num</code> or <code>int</code>). However it makes for less readable code when just comparing strings. Thus the recommendation is to avoid using string interpolation except only where it really makes sense (ie use it sparingly).</p><h3 id="non-boolean-logic" tabindex="-1"><a class="header-anchor" href="#non-boolean-logic" aria-hidden="true">#</a> Non-boolean logic</h3><p>Thus far the examples given have been focused on comparisons however <code>=</code> and <code>let</code> supports all the usual arithmetic operators:</p><pre><code>» = 10+10
20

» = 10/10
1

» = (4 * (3 + 2))
20

» = \`foo\`+\`bar\`
foobar
</code></pre><h3 id="read-more" tabindex="-1"><a class="header-anchor" href="#read-more" aria-hidden="true">#</a> Read more</h3>`,38),m={href:"https://github.com/Knetic/govaluate",target:"_blank",rel:"noopener noreferrer"},f=r(`<h3 id="type-annotations" tabindex="-1"><a class="header-anchor" href="#type-annotations" aria-hidden="true">#</a> Type Annotations</h3><p>When <code>set</code> or <code>global</code> are used as a function, the parameters are passed as a string which means the variables are defined as a <code>str</code>. If you wish to define them as an alternate data type then you should add type annotations:</p><pre><code>» set: int age = 30
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
</code></pre><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,44),g=e("code",null,"(",-1),v=e("code",null,"=",-1),x=e("code",null,"[[",-1),y=e("code",null,"[",-1),w=e("code",null,"export",-1),_=e("code",null,"expr",-1),k=e("code",null,"global",-1),q=e("code",null,"if",-1),$=e("code",null,"set",-1);function V(A,T){const t=i("RouterLink"),s=i("ExternalLinkIcon");return l(),c("div",null,[u,e("p",null,[e("strong",null,[o("This is a deprecated feature. Please refer to "),a(t,{to:"/commands/expr.html"},{default:n(()=>[p]),_:1}),o(" instead.")])]),b,e("p",null,[o("Murex uses the "),e("a",m,[o("govaluate package"),a(s)]),o(". More information can be found in it's manual.")]),f,e("ul",null,[e("li",null,[a(t,{to:"/user-guide/reserved-vars.html"},{default:n(()=>[o("Reserved Variables")]),_:1}),o(": Special variables reserved by Murex")]),e("li",null,[a(t,{to:"/user-guide/scoping.html"},{default:n(()=>[o("Variable and Config Scoping")]),_:1}),o(": How scoping works within Murex")]),e("li",null,[a(t,{to:"/commands/brace-quote.html"},{default:n(()=>[g,o(" (brace quote)")]),_:1}),o(": Write a string to the STDOUT without new line")]),e("li",null,[a(t,{to:"/commands/equ.html"},{default:n(()=>[v,o(" (arithmetic evaluation)")]),_:1}),o(": Evaluate a mathematical function (deprecated)")]),e("li",null,[a(t,{to:"/commands/element.html"},{default:n(()=>[x,o(" (element)")]),_:1}),o(": Outputs an element from a nested structure")]),e("li",null,[a(t,{to:"/commands/index2.html"},{default:n(()=>[y,o(" (index)")]),_:1}),o(": Outputs an element from an array, map or table")]),e("li",null,[a(t,{to:"/commands/export.html"},{default:n(()=>[w]),_:1}),o(": Define an environmental variable and set it's value")]),e("li",null,[a(t,{to:"/commands/expr.html"},{default:n(()=>[_]),_:1}),o(": Expressions: mathematical, string comparisons, logical operators")]),e("li",null,[a(t,{to:"/commands/global.html"},{default:n(()=>[k]),_:1}),o(": Define a global variable and set it's value")]),e("li",null,[a(t,{to:"/commands/if.html"},{default:n(()=>[q]),_:1}),o(": Conditional statement to execute different blocks of code depending on the result of the condition")]),e("li",null,[a(t,{to:"/commands/set.html"},{default:n(()=>[$]),_:1}),o(": Define a local variable and set it's value")])])])}const N=d(h,[["render",V],["__file","let.html.vue"]]);export{N as default};
