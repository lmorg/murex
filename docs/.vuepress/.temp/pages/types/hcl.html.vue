<template><div><h1 id="hcl" tabindex="-1"><a class="header-anchor" href="#hcl" aria-hidden="true">#</a> <code v-pre>hcl</code></h1>
<blockquote>
<p>HashiCorp Configuration Language (HCL)</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>The description below is taken from the HCL git repository's <a href="https://github.com/hashicorp/hcl" target="_blank" rel="noopener noreferrer">README<ExternalLinkIcon/></a>:</p>
<blockquote>
<p>HCL (HashiCorp Configuration Language) is a configuration language built by
HashiCorp. The goal of HCL is to build a structured configuration language
that is both human and machine friendly for use with command-line tools, but
specifically targeted towards DevOps tools, servers, etc.</p>
<p>HCL is also fully JSON compatible. That is, JSON can be used as completely
valid input to a system expecting HCL. This helps makes systems interoperable
with other systems.</p>
</blockquote>
<p>HCL support within Murex is pretty mature however it is not considered a
primitive. Which means, while it is a recommended builtin which you should
expect in most deployments of Murex, it's still an optional package and
thus may not be present in some edge cases. This is because it relies on
external source packages for the shell to compile.</p>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>terraform {
  required_version = &quot;~&gt; 0.11.5&quot;
}

data &quot;aws_availability_zones&quot; &quot;available&quot; {}

data &quot;aws_vpc&quot; &quot;vpc&quot; {
  id = &quot;${var.vpc_id}&quot;
}

data &quot;aws_route53_zone&quot; &quot;external&quot; {
  zone_id = &quot;${var.external_hosted_zone_id}&quot;
}

data &quot;aws_iam_policy_document&quot; &quot;assume_role_policy&quot; {
  statement {
    actions = [&quot;sts:AssumeRole&quot;]
    effect  = &quot;Allow&quot;

    principals {
      identifiers = [&quot;ec2.amazonaws.com&quot;]
      type        = &quot;Service&quot;
    }
  }
}
</code></pre>
<p>See the HashiCorp's <a href="https://github.com/hashicorp/hcl" target="_blank" rel="noopener noreferrer">documentation<ExternalLinkIcon/></a> for HCL syntax.</p>
<h2 id="default-associations" tabindex="-1"><a class="header-anchor" href="#default-associations" aria-hidden="true">#</a> Default Associations</h2>
<ul>
<li><strong>Extension</strong>: <code v-pre>hcl</code></li>
<li><strong>Extension</strong>: <code v-pre>tf</code></li>
<li><strong>Extension</strong>: <code v-pre>tfvars</code></li>
<li><strong>MIME</strong>: <code v-pre>application/hcl</code></li>
<li><strong>MIME</strong>: <code v-pre>application/x-hcl</code></li>
<li><strong>MIME</strong>: <code v-pre>text/hcl</code></li>
<li><strong>MIME</strong>: <code v-pre>text/x-hcl</code></li>
</ul>
<h2 id="supported-hooks" tabindex="-1"><a class="header-anchor" href="#supported-hooks" aria-hidden="true">#</a> Supported Hooks</h2>
<ul>
<li><code v-pre>Marshal()</code>
Supported via a JSON marshaller because HCL is designed to be written by humans but &quot;compiled&quot; into JSON</li>
<li><code v-pre>ReadArray()</code>
Works with HCL arrays. Maps are converted into arrays</li>
<li><code v-pre>ReadArrayWithType()</code>
Works with HCL arrays. Maps are converted into arrays. Elements data-type in Murex mirrors the HCL type of the element</li>
<li><code v-pre>ReadIndex()</code>
Works against all properties in HCL</li>
<li><code v-pre>ReadMap()</code>
Works with HCL maps</li>
<li><code v-pre>ReadNotIndex()</code>
Works against all properties in HCL</li>
<li><code v-pre>Unmarshal()</code>
Supported</li>
<li><code v-pre>WriteArray()</code>
Works with HCL arrays</li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/apis/Marshal.html"><code v-pre>Marshal()</code> (type)</RouterLink>:
Converts structured memory into a structured file format (eg for stdio)</li>
<li><RouterLink to="/apis/ReadArray.html"><code v-pre>ReadArray()</code> (type)</RouterLink>:
Read from a data type one array element at a time</li>
<li><RouterLink to="/apis/ReadIndex.html"><code v-pre>ReadIndex()</code> (type)</RouterLink>:
Data type handler for the index, <code v-pre>[</code>, builtin</li>
<li><RouterLink to="/apis/ReadMap.html"><code v-pre>ReadMap()</code> (type)</RouterLink>:
Treat data type as a key/value structure and read its contents</li>
<li><RouterLink to="/apis/ReadNotIndex.html"><code v-pre>ReadNotIndex()</code> (type)</RouterLink>:
Data type handler for the bang-prefixed index, <code v-pre>![</code>, builtin</li>
<li><RouterLink to="/apis/Unmarshal.html"><code v-pre>Unmarshal()</code> (type)</RouterLink>:
Converts a structured file format into structured memory</li>
<li><RouterLink to="/apis/WriteArray.html"><code v-pre>WriteArray()</code> (type)</RouterLink>:
Write a data type, one array element at a time</li>
<li><RouterLink to="/commands/element.html"><code v-pre>[[</code> (element)</RouterLink>:
Outputs an element from a nested structure</li>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/cast.html"><code v-pre>cast</code></RouterLink>:
Alters the data type of the previous function without altering it's output</li>
<li><RouterLink to="/commands/format.html"><code v-pre>format</code></RouterLink>:
Reformat one data-type into another data-type</li>
<li><RouterLink to="/types/json.html"><code v-pre>json</code> </RouterLink>:
JavaScript Object Notation (JSON)</li>
<li><RouterLink to="/types/jsonl.html"><code v-pre>jsonl</code> </RouterLink>:
JSON Lines</li>
<li><RouterLink to="/commands/open.html"><code v-pre>open</code></RouterLink>:
Open a file with a preferred handler</li>
<li><RouterLink to="/commands/pretty.html"><code v-pre>pretty</code></RouterLink>:
Prettifies JSON to make it human readable</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>runtime</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
<li><RouterLink to="/types/yaml.html"><code v-pre>yaml</code> </RouterLink>:
YAML Ain't Markup Language (YAML)</li>
<li><RouterLink to="/types/mxjson.html">mxjson</RouterLink>:
Murex-flavoured JSON (deprecated)</li>
</ul>
</div></template>


