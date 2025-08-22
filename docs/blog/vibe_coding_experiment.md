# My Vibe Coding Experiment

> My personal thoughts after trying out "vibe coding" for the first time.

So I tried out vibe coding for the first time the other day. If you've not heard of "vibe coding" before, and I hadn't either, basically it's where you describe your requirements and let AI write the code for you.

My problem: I wanted code that could parse a `plist` file. For anyone unfamiliar with plists: it is a horrible macOS-specific file format based on XML. This specific file is iTerm's colour theme config. I wanted to enable support for iTerm2 themes in my own terminal emulator.

Given I have little love for XML and this routine was just to allow importing a very specific type of config file, it's fair to say this wasn't an essential routine nor something I wanted to waste hours writing myself.

So I thought to myself "_I wonder if I can get ChatGPT to write the code for me?_"

Overall, it took a couple of hours and a few refinements where the original code produced didn't match specifications. In fact the entire process was a complete departure to my preferred approach to programming. Rather than debugging specific parts of each function call, it was a lot more trial and error. Albeit with the AI performing all of the "error" part and me "trialing" then feeding back to the AI where it errored.

That all said, the final code was still generated quicker than it likely would have taken me to write the parser by hand.

So I can see the appeal of this type of development for some people...and particularly if you're not already comfortable writing code yourself. It's a massive productive gain. However I wouldn't yet trust it for anything important just yet. A little like how you'd have lesser experienced developers working on less mission critical code while you're training them on the job.

However, for me, I cannot see "vibe coding" becoming anything more than a fun novelty that I seldom fallback on. For me, it felt a little too heavy on the "_let's see if this black box can randomly throw the correct sequence of characters_" to consider using this on anything that actually mattered.

<hr>

Published: 21.03.2025 at 21:11

## See Also

* [iTerm2 Integrations](../integrations/iterm2.md):
  Get more out of iTerm2 terminal emulator

<hr/>

This document was generated from [gen/blog/vibe_coding_experiment_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/blog/vibe_coding_experiment_doc.yaml).