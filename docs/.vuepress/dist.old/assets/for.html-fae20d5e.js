import{_ as a}from"./plugin-vue_export-helper-c27b6911.js";import{r as s,o as l,c as d,d as e,b as n,w as o,e as i,f as r}from"./app-45f7c304.js";const u={},c=r(`<h1 id="for" tabindex="-1"><a class="header-anchor" href="#for" aria-hidden="true">#</a> <code>for</code></h1><blockquote><p>A more familiar iteration loop to existing developers</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>This <code>for</code> loop is fills a small niche where <code>foreach</code> or <code>formap</code> are inappropiate in your script. It&#39;s generally not recommended to use <code>for</code> because it performs slower and doesn&#39;t adhere to Murex&#39;s design philosophy. However it does offer additional flexibility around recursion.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>for ( variable; conditional; incrementation ) { code-block } -&gt; \`&lt;stdout&gt;\`
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>    » for ( i=1; i&lt;6; i++ ) { echo $i }
    1
    2
    3
    4
    5
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="syntax" tabindex="-1"><a class="header-anchor" href="#syntax" aria-hidden="true">#</a> Syntax</h3><p><code>for</code> is a little naughty in terms of breaking Murex&#39;s style guidelines due to the first parameter being entered as one string treated as 3 separate code blocks. The syntax is like this for two reasons:</p><ol><li>readability (having multiple <code>{ blocks }</code> would make scripts unsightly</li><li>familiarity (for those using to <code>for</code> loops in other languages</li></ol><p>The first parameter is: <code>( i=1; i&lt;6; i++ )</code>, but it is then converted into the following code:</p><ol><li><code>let i=0</code> - declare the loop iteration variable</li><li><code>= i&lt;0</code> - if the condition is true then proceed to run the code in the second parameter - <code>{ echo $i }</code></li><li><code>let i++</code> - increment the loop iteration variable</li></ol><p>The second parameter is the code to execute upon each iteration</p><h3 id="better-for-loops" tabindex="-1"><a class="header-anchor" href="#better-for-loops" aria-hidden="true">#</a> Better <code>for</code> loops</h3><p>Because each iteration of a <code>for</code> loop reruns the 2nd 2 parts in the first parameter (the conditional and incrementation), <code>for</code> is very slow. Plus the weird, non-idiomatic, way of writing the 3 parts, it&#39;s fair to say <code>for</code> is not the recommended method of iteration and in fact there are better functions to achieve the same thing...most of the time at least.</p><p>For example:</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>    a: [1..5] -&gt; foreach: i { echo $i }
    1
    2
    3
    4
    5
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>The different in performance can be measured. eg:</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>    » time { a: [1..9999] -&gt; foreach: i { out: &lt;null&gt; $i } }
    0.097643108

    » time { for ( i=1; i&lt;10000; i=i+1 ) { out: &lt;null&gt; $i } }
    0.663812496
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>You can also do step ranges with <code>foreach</code>:</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>    » time { for ( i=10; i&lt;10001; i=i+2 ) { out: &lt;null&gt; $i } }
    0.346254973

    » time { a: [1..999][0,2,4,6,8],10000 -&gt; foreach i { out: &lt;null&gt; $i } }
    0.053924326
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>...though granted the latter is a little less readable.</p><p>The big catch with using <code>a</code> piped into <code>foreach</code> is that values are passed as strings rather than numbers.</p><h3 id="tips-when-writing-json-inside-for-loops" tabindex="-1"><a class="header-anchor" href="#tips-when-writing-json-inside-for-loops" aria-hidden="true">#</a> Tips when writing JSON inside for loops</h3><p>One of the drawbacks (or maybe advantages, depending on your perspective) of JSON is that parsers generally expect a complete file for processing in that the JSON specification requires closing tags for every opening tag. This means it&#39;s not always suitable for streaming. For example</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>    » ja [1..3] -&gt; foreach i { out ({ &quot;$i&quot;: $i }) }
    { &quot;1&quot;: 1 }
    { &quot;2&quot;: 2 }
    { &quot;3&quot;: 3 }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p><strong>What does this even mean and how can you build a JSON file up sequentially?</strong></p><p>One answer if to write the output in a streaming file format and convert back to JSON</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>    » ja [1..3] -&gt; foreach i { out (- &quot;$i&quot;: $i) }
    - &quot;1&quot;: 1
    - &quot;2&quot;: 2
    - &quot;3&quot;: 3

    » ja [1..3] -&gt; foreach i { out (- &quot;$i&quot;: $i) } -&gt; cast yaml -&gt; format json
    [
        {
            &quot;1&quot;: 1
        },
        {
            &quot;2&quot;: 2
        },
        {
            &quot;3&quot;: 3
        }
    ]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p><strong>What if I&#39;m returning an object rather than writing one?</strong></p><p>The problem with building JSON structures from existing structures is that you can quickly end up with invalid JSON due to the specifications strict use of commas.</p><p>For example in the code below, each item block is it&#39;s own object and there are no <code>[ ... ]</code> encapsulating them to denote it is an array of objects, nor are the objects terminated by a comma.</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>    » config -&gt; [ shell ] -&gt; formap k v { $v -&gt; alter /Foo Bar }
    {
        &quot;Data-Type&quot;: &quot;bool&quot;,
        &quot;Default&quot;: true,
        &quot;Description&quot;: &quot;Display the interactive shell&#39;s hint text helper. Please note, even when this is disabled, it will still appear when used for regexp searches and other readline-specific functions&quot;,
        &quot;Dynamic&quot;: false,
        &quot;Foo&quot;: &quot;Bar&quot;,
        &quot;Global&quot;: true,
        &quot;Value&quot;: true
    }
    {
        &quot;Data-Type&quot;: &quot;block&quot;,
        &quot;Default&quot;: &quot;{ progress $PID }&quot;,
        &quot;Description&quot;: &quot;Murex function to execute when an \`exec\` process is stopped&quot;,
        &quot;Dynamic&quot;: false,
        &quot;Foo&quot;: &quot;Bar&quot;,
        &quot;Global&quot;: true,
        &quot;Value&quot;: &quot;{ progress $PID }&quot;
    }
    {
        &quot;Data-Type&quot;: &quot;bool&quot;,
        &quot;Default&quot;: true,
        &quot;Description&quot;: &quot;ANSI escape sequences in Murex builtins to highlight syntax errors, history completions, {SGR} variables, etc&quot;,
        &quot;Dynamic&quot;: false,
        &quot;Foo&quot;: &quot;Bar&quot;,
        &quot;Global&quot;: true,
        &quot;Value&quot;: true
    }
    ...
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Luckily JSON also has it&#39;s own streaming format: JSON lines (<code>jsonl</code>). We can <code>cast</code> this output as <code>jsonl</code> then <code>format</code> it back into valid JSON:</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>    » config -&gt; [ shell ] -&gt; formap k v { $v -&gt; alter /Foo Bar } -&gt; cast jsonl -&gt; format json
    [
        {
            &quot;Data-Type&quot;: &quot;bool&quot;,
            &quot;Default&quot;: true,
            &quot;Description&quot;: &quot;Write shell history (interactive shell) to disk&quot;,
            &quot;Dynamic&quot;: false,
            &quot;Foo&quot;: &quot;Bar&quot;,
            &quot;Global&quot;: true,
            &quot;Value&quot;: true
        },
        {
            &quot;Data-Type&quot;: &quot;int&quot;,
            &quot;Default&quot;: 4,
            &quot;Description&quot;: &quot;Maximum number of lines with auto-completion suggestions to display&quot;,
            &quot;Dynamic&quot;: false,
            &quot;Foo&quot;: &quot;Bar&quot;,
            &quot;Global&quot;: true,
            &quot;Value&quot;: &quot;6&quot;
        },
        {
            &quot;Data-Type&quot;: &quot;bool&quot;,
            &quot;Default&quot;: true,
            &quot;Description&quot;: &quot;Display some status information about the stop process when ctrl+z is pressed (conceptually similar to ctrl+t / SIGINFO on some BSDs)&quot;,
            &quot;Dynamic&quot;: false,
            &quot;Foo&quot;: &quot;Bar&quot;,
            &quot;Global&quot;: true,
            &quot;Value&quot;: true
        },
    ...
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h4 id="foreach-will-automatically-cast-it-s-output-as-jsonl-if-it-s-stdin-type-is-json" tabindex="-1"><a class="header-anchor" href="#foreach-will-automatically-cast-it-s-output-as-jsonl-if-it-s-stdin-type-is-json" aria-hidden="true">#</a> <code>foreach</code> will automatically cast it&#39;s output as <code>jsonl</code> <em>if</em> it&#39;s STDIN type is <code>json</code></h4><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>    » ja: [Tom,Dick,Sally] -&gt; foreach: name { out Hello $name }
    Hello Tom
    Hello Dick
    Hello Sally

    » ja [Tom,Dick,Sally] -&gt; foreach name { out Hello $name } -&gt; debug -&gt; [[ /Data-Type/Murex ]]
    jsonl

    » ja: [Tom,Dick,Sally] -&gt; foreach: name { out Hello $name } -&gt; format: json
    [
        &quot;Hello Tom&quot;,
        &quot;Hello Dick&quot;,
        &quot;Hello Sally&quot;
    ]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,40),v=e("code",null,"a",-1),m=e("code",null,"break",-1),h=e("code",null,"foreach",-1),b=e("code",null,"formap",-1),p=e("code",null,"if",-1),q=e("code",null,"ja",-1),f=e("code",null,"let",-1),g=e("code",null,"set",-1),x=e("code",null,"while",-1);function y(w,_){const t=s("RouterLink");return l(),d("div",null,[c,e("ul",null,[e("li",null,[n(t,{to:"/commands/a.html"},{default:o(()=>[v,i(" (mkarray)")]),_:1}),i(": A sophisticated yet simple way to build an array or list")]),e("li",null,[n(t,{to:"/commands/break.html"},{default:o(()=>[m]),_:1}),i(": Terminate execution of a block within your processes scope")]),e("li",null,[n(t,{to:"/commands/foreach.html"},{default:o(()=>[h]),_:1}),i(": Iterate through an array")]),e("li",null,[n(t,{to:"/commands/formap.html"},{default:o(()=>[b]),_:1}),i(": Iterate through a map or other collection of data")]),e("li",null,[n(t,{to:"/commands/if.html"},{default:o(()=>[p]),_:1}),i(": Conditional statement to execute different blocks of code depending on the result of the condition")]),e("li",null,[n(t,{to:"/commands/ja.html"},{default:o(()=>[q,i(" (mkarray)")]),_:1}),i(": A sophisticated yet simply way to build a JSON array")]),e("li",null,[n(t,{to:"/commands/let.html"},{default:o(()=>[f]),_:1}),i(": Evaluate a mathematical function and assign to variable (deprecated)")]),e("li",null,[n(t,{to:"/commands/set.html"},{default:o(()=>[g]),_:1}),i(": Define a local variable and set it's value")]),e("li",null,[n(t,{to:"/commands/while.html"},{default:o(()=>[x]),_:1}),i(": Loop until condition false")])])])}const j=a(u,[["render",y],["__file","for.html.vue"]]);export{j as default};
