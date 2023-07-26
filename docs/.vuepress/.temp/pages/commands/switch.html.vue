<template><div><h1 id="switch" tabindex="-1"><a class="header-anchor" href="#switch" aria-hidden="true">#</a> <code v-pre>switch</code></h1>
<blockquote>
<p>Blocks of cascading conditionals</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>switch</code> is a large block for simplifying cascades of conditional statements.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>switch [value] {
  case | if { conditional } [then] { code-block }
  case | if { conditional } [then] { code-block }
  ...
  [ default { code-block } ]
} -> &lt;stdout>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>The first parameter should be either <strong>case</strong> or <strong>if</strong> -- the statements are
subtly different and thus alter the behavior of <code v-pre>switch</code>.</p>
<p><strong>then</strong> is optional ('then' is assumed even if not explicitly present).</p>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Output an array of editors installed:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>switch {
    if { which: vi    } { out: vi    }
    if { which: vim   } { out: vim   }
    if { which: nano  } { out: nano  }
    if { which: emacs } { out: emacs }
} -> format: json
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>A higher/lower game written using <code v-pre>switch</code>:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>function higherlower {
  try {
    rand: int 100 -> set rand
    while { $rand } {
      read: guess "Guess a number between 1 and 100: "

      switch {
        case: { = $guess &lt; $rand } then {
          out: "Too low"
        }

        case: { = $guess > $rand } then {
          out: "Too high"
        }

        default: {
          out: "Correct"
          let: rand=0
        }
      }
    }
  }
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>String matching with <code v-pre>switch</code>:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>read: name "What is your name? "
switch $name {
    case "Tom"   { out: "I have a brother called Tom" }
    case "Dick"  { out: "I have an uncle called Dick" }
    case "Sally" { out: "I have a sister called Sally" }
    default      { err: "That is an odd name" }
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="comparing-values-vs-boolean-state" tabindex="-1"><a class="header-anchor" href="#comparing-values-vs-boolean-state" aria-hidden="true">#</a> Comparing Values vs Boolean State</h3>
<h4 id="by-values" tabindex="-1"><a class="header-anchor" href="#by-values" aria-hidden="true">#</a> By Values</h4>
<p>If you supply a value with <code v-pre>switch</code>...</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>switch value { ... }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>...then all the conditionals are compared against that value. For example:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>switch foo {
    case bar {
        # not executed because foo != bar
    }
    case foo {
        # executed because foo != foo
    }
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>You can use code blocks to return strings too</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>switch foo {
    case {out: bar} then {
        # not executed because foo != bar
    }
    case {out: foo} then {
        # executed because foo != foo
    }
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h4 id="by-boolean-state" tabindex="-1"><a class="header-anchor" href="#by-boolean-state" aria-hidden="true">#</a> By Boolean State</h4>
<p>This style of syntax could be argued as a prettier counterpart to if/else if.
Only code blocks are support and each block is checked for its boolean state
rather than string matching.</p>
<p>This is simply written as:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>switch { ... }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><h3 id="when-to-use-case-if-and-default" tabindex="-1"><a class="header-anchor" href="#when-to-use-case-if-and-default" aria-hidden="true">#</a> When To Use <code v-pre>case</code>, <code v-pre>if</code> and <code v-pre>default</code>?</h3>
<p>A <code v-pre>switch</code> command may contain multiple <strong>case</strong> and <strong>if</strong> blocks. These
statements subtly alter the behavior of <code v-pre>switch</code>. You can mix and match <strong>if</strong>
and <strong>case</strong> statements within the same <code v-pre>switch</code> block.</p>
<h4 id="case" tabindex="-1"><a class="header-anchor" href="#case" aria-hidden="true">#</a> case</h4>
<p>A <strong>case</strong> statement will only move on to the next statement if the result of
the <strong>case</strong> statement is <strong>false</strong>. If a <strong>case</strong> statement is <strong>true</strong> then
<code v-pre>switch</code> will exit with an exit number of <code v-pre>0</code>.</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>switch {
    case { false } then {
        # ignored because case == false
    }
    case { true } then {
        # executed because case == true
    }
    case { true } then {
        # ignored because a previous case was true
    }
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="if" tabindex="-1"><a class="header-anchor" href="#if" aria-hidden="true">#</a> if</h3>
<p>An <strong>if</strong> statement will proceed to the next statement <em>even</em> if the result of
the <strong>if</strong> statement is <strong>true</strong>.</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>switch {
    if { false } then {
        # ignored because if == false
    }
    if { true } then {
        # executed because if == true
    }
    if { true } then {
        # executed because if == true
    }
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="default" tabindex="-1"><a class="header-anchor" href="#default" aria-hidden="true">#</a> default</h3>
<p><strong>default</strong> statements are only run if <em>all</em> <strong>case</strong> <em>and</em> <strong>if</strong> statements are
false.</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>switch {
    if { false } then {
        # ignored because if == false
    }
    if { true } then {
        # executed because if == true
    }
    if { true } then {
        # executed because if == true
    }
    if { false } then {
        # ignored because if == false
    }
    default {
        # ignored because one or more previous if's were true
    }
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><blockquote>
<p><strong>default</strong> was added in Murex version 3.1</p>
</blockquote>
<h3 id="catch" tabindex="-1"><a class="header-anchor" href="#catch" aria-hidden="true">#</a> catch</h3>
<p><strong>catch</strong> has been deprecated in version 3.1 and replaced with <strong>default</strong>.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/not.html"><code v-pre>!</code> (not)</RouterLink>:
Reads the STDIN and exit number from previous process and not's it's condition</li>
<li><RouterLink to="/commands/and.html"><code v-pre>and</code></RouterLink>:
Returns <code v-pre>true</code> or <code v-pre>false</code> depending on whether multiple conditions are met</li>
<li><RouterLink to="/commands/break.html"><code v-pre>break</code></RouterLink>:
Terminate execution of a block within your processes scope</li>
<li><RouterLink to="/commands/catch.html"><code v-pre>catch</code></RouterLink>:
Handles the exception code raised by <code v-pre>try</code> or <code v-pre>trypipe</code></li>
<li><RouterLink to="/commands/false.html"><code v-pre>false</code></RouterLink>:
Returns a <code v-pre>false</code> value</li>
<li><RouterLink to="/commands/if.html"><code v-pre>if</code></RouterLink>:
Conditional statement to execute different blocks of code depending on the result of the condition</li>
<li><RouterLink to="/commands/let.html"><code v-pre>let</code></RouterLink>:
Evaluate a mathematical function and assign to variable (deprecated)</li>
<li><RouterLink to="/commands/or.html"><code v-pre>or</code></RouterLink>:
Returns <code v-pre>true</code> or <code v-pre>false</code> depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.</li>
<li><RouterLink to="/commands/set.html"><code v-pre>set</code></RouterLink>:
Define a local variable and set it's value</li>
<li><RouterLink to="/commands/true.html"><code v-pre>true</code></RouterLink>:
Returns a <code v-pre>true</code> value</li>
<li><RouterLink to="/commands/try.html"><code v-pre>try</code></RouterLink>:
Handles errors inside a block of code</li>
<li><RouterLink to="/commands/trypipe.html"><code v-pre>trypipe</code></RouterLink>:
Checks state of each function in a pipeline and exits block on error</li>
<li><RouterLink to="/commands/while.html"><code v-pre>while</code></RouterLink>:
Loop until condition false</li>
</ul>
</div></template>


