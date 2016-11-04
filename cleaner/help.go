package cleaner

const HelpString = `<h2>Safe HTML</h2>
<p>You are allowed to use a safe subset of HTML.</p>
<p>The <code>title</code> attribute is allowed on all elements.</p>
<p><code>&lt;a href&gt;</code> is allowed for HTTP, HTTPS, MAILTO, and DATA links. <code>rel="nofollow"</code> is allowed and will be added automatically for users with less than 10 reputation.</p>
<p><code>&lt;img src&gt;</code> is allowed with the same restrictions for URLs as links. The <code>alt</code>, <code>width</code>, and <code>height</code> attributes are allowed, but optional. The <code>class</code> attribute may be specified with a value of <code>emoji</code> to simulate the appearance of an emoji.</p>
<p><code>&lt;video src&gt;</code> and <code>&lt;audio src&gt;</code> are allowed with the same restrictions for URLs. The <code>poster</code> attribute is allowed on videos, but not required. <code>controls</code> will automatically be added.</p>
<p>Basic formatting tags are allowed: <code>&lt;b&gt;</code>, <code>&lt;i&gt;</code>, <code>&lt;u&gt;</code>, <code>&lt;s&gt;</code>, <code>&lt;em&gt;</code>, <code>&lt;strong&gt;</code>, <code>&lt;strike&gt;</code>, <code>&lt;big&gt;</code>, <code>&lt;small&gt;</code>, <code>&lt;sup&gt;</code>, <code>&lt;sub&gt;</code>, <code>&lt;ins&gt;</code>, <code>&lt;del&gt;</code>, <code>&lt;abbr&gt;</code>, <code>&lt;address&gt;</code>, <code>&lt;cite&gt;</code>, and <code>&lt;q&gt;</code></p>
<p><code>&lt;p&gt;</code>, <code>&lt;blockquote&gt;</code>, <code>&lt;br&gt;</code>, <code>&lt;hr&gt;</code>, and headings <code>&lt;h1&gt;</code> through <code>&lt;h6&gt;</code> are supported.</p>
<p>For code, <code>&lt;pre&gt;</code>, <code>&lt;code&gt;</code>, <code>&lt;kbd&gt;</code>, and <code>&lt;tt&gt;</code> are allowed.</p>
<p><code>&lt;table&gt;</code> is supported with an optional <code>class</code> of any combination of <code>table</code>, <code>table-bordered</code>, and <code>table-striped</code>.</p>
<p><code>&lt;ul&gt;</code> and <code>&lt;ol&gt;</code> are supported with an optional <code>start</code> attribute, and <code>&lt;li&gt;</code> allows an optional <code>value</code> attribute.</p>
<p>For spoilers, use <code>&lt;details&gt;</code> with an optional <code>&lt;summary&gt;</code>. <code>open</code> may be specified to start the spoiler visible.</p>
<p>Definition lists are supported using <code>&lt;dl&gt;</code>, <code>&lt;dt&gt;</code>, and <code>&lt;dd&gt;</code>.</p>`
