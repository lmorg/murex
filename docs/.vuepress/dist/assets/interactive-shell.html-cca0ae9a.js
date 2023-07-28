import{_ as r}from"./plugin-vue_export-helper-c27b6911.js";import{r as l,o as d,c as h,d as e,b as i,e as t,w as n,f as a}from"./app-45f7c304.js";const c={},u=e("h1",{id:"murex-s-interactive-shell-user-guide",tabindex:"-1"},[e("a",{class:"header-anchor",href:"#murex-s-interactive-shell-user-guide","aria-hidden":"true"},"#"),t(" Murex's Interactive Shell - User Guide")],-1),p=e("blockquote",null,[e("p",null,"What's different about Murex's interactive shell?")],-1),g=e("h2",{id:"overview",tabindex:"-1"},[e("a",{class:"header-anchor",href:"#overview","aria-hidden":"true"},"#"),t(" Overview")],-1),f=e("p",null,"Aside from Murex being carefully designed with scripting in mind, the interactive shell itself is also built around productivity. To achieve this we wrote our own readline library. Below is an example of that library in use:",-1),m={href:"https://asciinema.org/a/232714",target:"_blank",rel:"noopener noreferrer"},_=e("img",{src:"https://asciinema.org/a/232714.svg",alt:"asciicast",loading:"lazy"},null,-1),b=a('<p>The above demo includes the following features of Murex&#39;s bespoke readline library:</p><ul><li>hint text - blue status text below the prompt (the colour is configurable)</li><li>syntax highlighting (albeit there isn’t much syntax to highlight in the example). This can also be turned off if your preference is to have colours disabled</li><li>tab-completion in gridded mode (seen when typing <code>cd</code>)</li><li>tab-completion in list view (seen when selecting a process name to <code>kill</code> where the process ID was substituted when selected)</li><li>searching through the tab-completion suggestions (seen in both <code>cd</code> and <code>kill</code> - enabled by pressing <code>[ctrl]</code>+<code>[f]</code>)</li><li>line editing using $EDITOR (<code>vi</code> in the example - enabled by pressing <code>[esc]</code> followed by <code>[v]</code>)</li><li>readline’s warning before pasting multiple lines of data into the buffer and the preview option that’s available as part of the aforementioned warning</li><li>and VIM keys (enabled by pressing <code>[esc]</code>)</li></ul><h2 id="readline" tabindex="-1"><a class="header-anchor" href="#readline" aria-hidden="true">#</a> readline</h2><p>Murex uses a custom <code>readline</code> library to enable support for new features on in addition to the existing uses you&#39;d normally expect from a shell. It is because of this Murex provides one of the best user experiences of any of the shells available today.</p><h3 id="hotkeys" tabindex="-1"><a class="header-anchor" href="#hotkeys" aria-hidden="true">#</a> Hotkeys</h3>',5),x=e("h3",{id:"autocompletion",tabindex:"-1"},[e("a",{class:"header-anchor",href:"#autocompletion","aria-hidden":"true"},"#"),t(" Autocompletion")],-1),y=e("code",null,"[tab]",-1),w=e("code",null,"autocomplete",-1),k=e("code",null,"|",-1),v=e("code",null,"->",-1),T=a(`<p>The <code>|</code> token will behave much like any other shell however <code>-&gt;</code> will offer suggestions with matching data types (as seen in <code>runtime --methods</code>). This is a way of helping highlight commands that naturally follow after another in a pipeline. Which is particularly important in Murex as it introduces data types and dozens of new builtins specifically for working with data structures in an intelligent and readable yet succinct way.</p><p>You can add your own commands and functions to Murex as methods by defining them with <code>method</code>. For example if we were to add <code>jq</code> as a method:</p><pre><code>method: define jq {
    &quot;Stdin&quot;:  &quot;json&quot;,
    &quot;Stdout&quot;: &quot;@Any&quot;
}
</code></pre><h3 id="syntax-completion" tabindex="-1"><a class="header-anchor" href="#syntax-completion" aria-hidden="true">#</a> Syntax Completion</h3><p>Like with most IDEs, Murex will auto close brackets et al.</p>`,5),S={href:"https://asciinema.org/a/408029",target:"_blank",rel:"noopener noreferrer"},I=e("img",{src:"https://asciinema.org/a/408029.svg",alt:"asciicast",loading:"lazy"},null,-1),q=e("h3",{id:"syntax-highlighting",tabindex:"-1"},[e("a",{class:"header-anchor",href:"#syntax-highlighting","aria-hidden":"true"},"#"),t(" Syntax Highlighting")],-1),M=e("p",null,"Pipelines in the interactive terminal are syntax highlighted. This is similar to what one expects from an IDE.",-1),A=e("p",null,"Syntax highlighting can be disabled by running:",-1),D=e("pre",null,[e("code",null,`config: set shell syntax-highlighting off
`)],-1),N=e("h3",{id:"spellchecker",tabindex:"-1"},[e("a",{class:"header-anchor",href:"#spellchecker","aria-hidden":"true"},"#"),t(" Spellchecker")],-1),C=e("p",null,"Murex supports inline spellchecking, where errors are underlined. For example",-1),H={href:"https://asciinema.org/a/408024",target:"_blank",rel:"noopener noreferrer"},B=e("img",{src:"https://asciinema.org/a/408024.svg",alt:"asciicast",loading:"lazy"},null,-1),P=a(`<h3 id="hint-text" tabindex="-1"><a class="header-anchor" href="#hint-text" aria-hidden="true">#</a> Hint Text</h3><p>The <strong>hint text</strong> is a (typically) blue status line that appears directly below your prompt. The idea behind the <strong>hint text</strong> is to provide clues to you as type instructions into the prompt; but without adding distractions. It is there to be used if you want it while keeping out of the way when you don&#39;t want it.</p><h4 id="configuring-hint-text-colour" tabindex="-1"><a class="header-anchor" href="#configuring-hint-text-colour" aria-hidden="true">#</a> Configuring Hint Text Colour</h4><p>By default the <strong>hint text</strong> will appear blue. This is also customizable:</p><pre><code>» config get shell hint-text-formatting
{BLUE}
</code></pre>`,5),E=a(`<p>It is also worth noting that if colour is disabled then the <strong>hint text</strong> will not be coloured even if <strong>hint-text-formatting</strong> includes colour codes:</p><pre><code>» config: set shell color false
</code></pre><p>(please note that <strong>syntax highlighting</strong> is unaffected by the above config)</p><h3 id="custom-hint-text-statuses" tabindex="-1"><a class="header-anchor" href="#custom-hint-text-statuses" aria-hidden="true">#</a> Custom Hint Text Statuses</h3><p>There is a lot of behavior hardcoded into Murex like displaying the full path to executables and the values of variables. However if there is no status to be displayed then Murex can fallback to a default <strong>hint text</strong> status. This default is a user defined function. At time of writing this document the author has the following function defined:</p><pre><code>config: set shell hint-text-func {
    trypipe &lt;!null&gt; {
        git status --porcelain -b -&gt; set gitstatus
        $gitstatus -&gt; head -n1 -&gt; regexp &#39;s/^## //&#39; -&gt; regexp &#39;s/\\.\\.\\./ =&gt; /&#39;
    }
    catch {
        out &quot;Not a git repository.&quot;
    }
}
</code></pre><p>...which produces a colorized status that looks something like the following:</p><pre><code>develop =&gt; origin/develop
</code></pre><h4 id="disabling-hint-text" tabindex="-1"><a class="header-anchor" href="#disabling-hint-text" aria-hidden="true">#</a> Disabling Hint Text</h4><p>It is enabled by default but can be disabled if you prefer a more minimal prompt:</p><pre><code>» config: set shell hint-text-enabled false
</code></pre><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,12),O=e("code",null,"->",-1),z=e("code",null,"{",-1),L=e("code",null,"}",-1),V=e("code",null,"|",-1),R=e("code",null,"autocomplete",-1),U=e("code",null,"config",-1),j=e("code",null,"method",-1),F=e("code",null,"runtime",-1);function W(X,$){const s=l("ExternalLinkIcon"),o=l("RouterLink");return d(),h("div",null,[u,p,g,f,e("p",null,[e("a",m,[_,i(s)])]),b,e("p",null,[t("A full breakdown of supported hotkeys is available at "),i(o,{to:"/user-guide/terminal-keys.html"},{default:n(()=>[t("terminal-keys.md")]),_:1}),t(".")]),x,e("p",null,[t("Autocompletion happen when you press "),y,t(" and will differ slightly depending on what is defined in "),w,t(" and whether you use the traditional "),i(o,{to:"/parser/pipe-posix.html"},{default:n(()=>[t("POSIX pipe token")]),_:1}),t(", "),k,t(", or the "),i(o,{to:"/parser/pipe-arrow.html"},{default:n(()=>[t("arrow pipe")]),_:1}),t(", "),v,t(".")]),T,e("p",null,[e("a",S,[I,i(s)])]),q,M,A,D,N,C,e("p",null,[e("a",H,[B,i(s)])]),e("p",null,[t("This might require some manual steps to enable, please see the "),i(o,{to:"/user-guide/spellcheck.html"},{default:n(()=>[t("spellcheck user guide")]),_:1}),t(" for more details.")]),P,e("p",null,[t("The formatting config takes a string and supports "),i(o,{to:"/user-guide/ansi.html"},{default:n(()=>[t("ANSI constants")]),_:1}),t(".")]),E,e("ul",null,[e("li",null,[i(o,{to:"/user-guide/ansi.html"},{default:n(()=>[t("ANSI Constants")]),_:1}),t(": Infixed constants that return ANSI escape sequences")]),e("li",null,[i(o,{to:"/parser/pipe-arrow.html"},{default:n(()=>[t("Arrow Pipe ("),O,t(") Token")]),_:1}),t(": Pipes STDOUT from the left hand command to STDIN of the right hand command")]),e("li",null,[i(o,{to:"/user-guide/code-block.html"},{default:n(()=>[t("Code Block Parsing")]),_:1}),t(": Overview of how code blocks are parsed")]),e("li",null,[i(o,{to:"/parser/curly-brace.html"},{default:n(()=>[t("Curly Brace ("),z,t(", "),L,t(") Tokens")]),_:1}),t(": Initiates or terminates a code block")]),e("li",null,[i(o,{to:"/parser/pipe-posix.html"},{default:n(()=>[t("POSIX Pipe ("),V,t(") Token")]),_:1}),t(": Pipes STDOUT from the left hand command to STDIN of the right hand command")]),e("li",null,[i(o,{to:"/user-guide/spellcheck.html"},{default:n(()=>[t("Spellcheck")]),_:1}),t(": How to enable inline spellchecking")]),e("li",null,[i(o,{to:"/user-guide/terminal-keys.html"},{default:n(()=>[t("Terminal Hotkeys")]),_:1}),t(": A list of all the terminal hotkeys and their uses")]),e("li",null,[i(o,{to:"/commands/autocomplete.html"},{default:n(()=>[R]),_:1}),t(": Set definitions for tab-completion in the command line")]),e("li",null,[i(o,{to:"/commands/config.html"},{default:n(()=>[U]),_:1}),t(": Query or define Murex runtime settings")]),e("li",null,[i(o,{to:"/commands/method.html"},{default:n(()=>[j]),_:1}),t(": Define a methods supported data-types")]),e("li",null,[i(o,{to:"/commands/runtime.html"},{default:n(()=>[F]),_:1}),t(": Returns runtime information on the internal state of Murex")])])])}const Y=r(c,[["render",W],["__file","interactive-shell.html.vue"]]);export{Y as default};