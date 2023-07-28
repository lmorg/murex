import{_ as o}from"./plugin-vue_export-helper-c27b6911.js";import{r as s,o as l,c as d,d as e,b as a,w as n,e as t,f as r}from"./app-45f7c304.js";const u={},c=r(`<h1 id="foreach" tabindex="-1"><a class="header-anchor" href="#foreach" aria-hidden="true">#</a> <code>foreach</code></h1><blockquote><p>Iterate through an array</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p><code>foreach</code> reads an array or map from STDIN and iterates through it, running a code block for each iteration with the value of the iterated element passed to it.</p><p>By default <code>foreach</code>&#39;s output data type is inherieted from its input data type. For example is STDIN is <code>yaml</code> then so will STDOUT. The only exception to this is if STDIN is <code>json</code> in which case STDOUT will be jsonlines (<code>jsonl</code>), or when additional flags are used such as <code>--jmap</code>.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><p><code>{ code-block }</code> reads from a variable and writes to an array / unbuffered STDOUT:</p><pre><code>\`&lt;stdin&gt;\` -&gt; foreach variable { code-block } -&gt; \`&lt;stdout&gt;\`
</code></pre><p><code>{ code-block }</code> reads from STDIN and writes to an array / unbuffered STDOUT:</p><pre><code>\`&lt;stdin&gt;\` -&gt; foreach { -&gt; code-block } -&gt; \`&lt;stdout&gt;\`
</code></pre><p><code>foreach</code> writes to a buffered JSON map:</p><pre><code>\`&lt;stdin&gt;\` -&gt; foreach --jmap variable { code-block (map key) } { code-block (map value) } -&gt; \`&lt;stdout&gt;\`
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><p>There are two basic ways you can write a <code>foreach</code> loop depending on how you want the iterated element passed to the code block.</p><p>The first option is to specify a temporary variable which can be read by the code block:</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>    » a [1..3] -&gt; foreach i { out $i }
    1
    2
    3
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><blockquote><p>Please note that the variable is specified <strong>without</strong> the dollar prefix, then used in the code block <strong>with</strong> the dollar prefix.</p></blockquote><p>The second option is for the code block&#39;s STDIN to read the element:</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>    » a [1..3] -&gt; foreach { -&gt; cat }
    1
    2
    3
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><blockquote><p>STDIN can only be read as the first command. If you cannot process the element on the first command then it is recommended you use the first option (passing a variable) instead.</p></blockquote><h3 id="writing-json-maps" tabindex="-1"><a class="header-anchor" href="#writing-json-maps" aria-hidden="true">#</a> Writing JSON maps</h3><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>    » ja [Monday..Friday] -&gt; foreach --jmap day { out $day -&gt; left 3 } { $day }
    {
        &quot;Fri&quot;: &quot;Friday&quot;,
        &quot;Mon&quot;: &quot;Monday&quot;,
        &quot;Thu&quot;: &quot;Thursday&quot;,
        &quot;Tue&quot;: &quot;Tuesday&quot;,
        &quot;Wed&quot;: &quot;Wednesday&quot;
    }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="using-steps-to-jump-iterations-by-more-than-1-one" tabindex="-1"><a class="header-anchor" href="#using-steps-to-jump-iterations-by-more-than-1-one" aria-hidden="true">#</a> Using steps to jump iterations by more than 1 (one)</h3><p>You can step through an array, list or table in jumps of user definable quantities. The value passed in STDIN and $VAR will be an array of all the records within that step range. For example:</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>    » %[1..10] -&gt; foreach --step 3 value { out &quot;Iteration $.i: $value&quot; }
    Iteration 1: [
        1,
        2,
        3
    ]
    Iteration 2: [
        4,
        5,
        6
    ]
    Iteration 3: [
        7,
        8,
        9
    ]
    Iteration 4: [
        10
    ]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2><ul><li><code>--jmap</code> Write a <code>json</code> map to STDOUT instead of an array</li><li><code>--step</code><code>&lt;int&gt;</code> Iterates in steps. Value passed to block is an array of items in the step range. Not (yet) supported with \`--jmap</li></ul><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="meta-values" tabindex="-1"><a class="header-anchor" href="#meta-values" aria-hidden="true">#</a> Meta values</h3><p>Meta values are a JSON object stored as the variable <code>$.</code>. The meta variable will get overwritten by any other block which invokes meta values. So if you wish to persist meta values across blocks you will need to reassign <code>$.</code>, eg</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>    %[1..3] -&gt; foreach {
        meta_parent = $.
        %[7..9] -&gt; foreach {
            out &quot;$(meta_parent.i): $.i&quot;
        }
    }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>The following meta values are defined:</p><ul><li><code>i</code>: iteration number</li></ul><h3 id="preserving-the-data-type-when-no-flags-used" tabindex="-1"><a class="header-anchor" href="#preserving-the-data-type-when-no-flags-used" aria-hidden="true">#</a> Preserving the data type (when no flags used)</h3><p><code>foreach</code> will preserve the data type read from STDIN in all instances where data is being passed along the pipeline and push that data type out at the other end:</p><ul><li>The temporary variable will be created with the same data-type as <code>foreach</code>&#39;s STDIN, or the data type of the array element (eg if it is a string or number)</li><li>The code block&#39;s STDIN will have the same data-type as <code>foreach</code>&#39;s STDIN</li><li><code>foreeach</code>&#39;s STDOUT will also be the same data-type as it&#39;s STDIN (or <code>jsonl</code> (jsonlines) where STDIN was <code>json</code> because <code>jsonl</code> better supports streaming)</li></ul><p>This last point means you may need to <code>cast</code> your data if you&#39;re writing data in a different format. For example the following is creating a YAML list however the data-type is defined as <code>json</code>:</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>    » ja [1..3] -&gt; foreach i { out &quot;- $i&quot; }
    - 1
    - 2
    - 3

    » ja [1..3] -&gt; foreach i { out &quot;- $i&quot; } -&gt; debug -&gt; [[ /Data-Type/Murex ]]
    json
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Thus any marshalling or other data-type-aware API&#39;s would fail because they are expecting <code>json</code> and receiving an incompatible data format.</p><p>This can be resolved via <code>cast</code>:</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>    » ja [1..3] -&gt; foreach i { out &quot;- $i&quot; } -&gt; cast yaml
    - 1
    - 2
    - 3

    » ja [1..3] -&gt; foreach i { out &quot;- $i&quot; } -&gt; cast yaml -&gt; debug -&gt; [[ /Data-Type/Murex ]]
    yaml
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>The output is the same but now it&#39;s defined as <code>yaml</code> so any further pipelined processes will now automatically use YAML marshallers when reading that data.</p><h3 id="tips-when-writing-json-inside-for-loops" tabindex="-1"><a class="header-anchor" href="#tips-when-writing-json-inside-for-loops" aria-hidden="true">#</a> Tips when writing JSON inside for loops</h3><p>One of the drawbacks (or maybe advantages, depending on your perspective) of JSON is that parsers generally expect a complete file for processing in that the JSON specification requires closing tags for every opening tag. This means it&#39;s not always suitable for streaming. For example</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>    » ja [1..3] -&gt; foreach i { out ({ &quot;$i&quot;: $i }) }
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
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,57),v=e("code",null,"ReadArrayWithType()",-1),m=e("code",null,"[[",-1),h=e("code",null,"a",-1),b=e("code",null,"break",-1),p=e("code",null,"cast",-1),f=e("code",null,"debug",-1),q=e("code",null,"for",-1),g=e("code",null,"formap",-1),y=e("code",null,"format",-1),x=e("code",null,"if",-1),w=e("code",null,"ja",-1),_=e("code",null,"json",-1),T=e("code",null,"jsonl",-1),D=e("code",null,"left",-1),j=e("code",null,"out",-1),S=e("code",null,"while",-1),k=e("code",null,"yaml",-1);function N(I,O){const i=s("RouterLink");return l(),d("div",null,[c,e("ul",null,[e("li",null,[a(i,{to:"/apis/ReadArrayWithType.html"},{default:n(()=>[v,t(" (type)")]),_:1}),t(": Read from a data type one array element at a time and return the elements contents and data type")]),e("li",null,[a(i,{to:"/commands/element.html"},{default:n(()=>[m,t(" (element)")]),_:1}),t(": Outputs an element from a nested structure")]),e("li",null,[a(i,{to:"/commands/a.html"},{default:n(()=>[h,t(" (mkarray)")]),_:1}),t(": A sophisticated yet simple way to build an array or list")]),e("li",null,[a(i,{to:"/commands/break.html"},{default:n(()=>[b]),_:1}),t(": Terminate execution of a block within your processes scope")]),e("li",null,[a(i,{to:"/commands/cast.html"},{default:n(()=>[p]),_:1}),t(": Alters the data type of the previous function without altering it's output")]),e("li",null,[a(i,{to:"/commands/debug.html"},{default:n(()=>[f]),_:1}),t(": Debugging information")]),e("li",null,[a(i,{to:"/commands/for.html"},{default:n(()=>[q]),_:1}),t(": A more familiar iteration loop to existing developers")]),e("li",null,[a(i,{to:"/commands/formap.html"},{default:n(()=>[g]),_:1}),t(": Iterate through a map or other collection of data")]),e("li",null,[a(i,{to:"/commands/format.html"},{default:n(()=>[y]),_:1}),t(": Reformat one data-type into another data-type")]),e("li",null,[a(i,{to:"/commands/if.html"},{default:n(()=>[x]),_:1}),t(": Conditional statement to execute different blocks of code depending on the result of the condition")]),e("li",null,[a(i,{to:"/commands/ja.html"},{default:n(()=>[w,t(" (mkarray)")]),_:1}),t(": A sophisticated yet simply way to build a JSON array")]),e("li",null,[a(i,{to:"/types/json.html"},{default:n(()=>[_]),_:1}),t(": JavaScript Object Notation (JSON)")]),e("li",null,[a(i,{to:"/types/jsonl.html"},{default:n(()=>[T]),_:1}),t(": JSON Lines")]),e("li",null,[a(i,{to:"/commands/left.html"},{default:n(()=>[D]),_:1}),t(": Left substring every item in a list")]),e("li",null,[a(i,{to:"/commands/out.html"},{default:n(()=>[j]),_:1}),t(": Print a string to the STDOUT with a trailing new line character")]),e("li",null,[a(i,{to:"/commands/while.html"},{default:n(()=>[S]),_:1}),t(": Loop until condition false")]),e("li",null,[a(i,{to:"/types/yaml.html"},{default:n(()=>[k]),_:1}),t(": YAML Ain't Markup Language (YAML)")])])])}const J=o(u,[["render",N],["__file","foreach.html.vue"]]);export{J as default};
