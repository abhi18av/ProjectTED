var html2json = require('html2json').html2json;
var json2html = require('html2json').json2html;


import "html2json" from "html2json";

var str = `< div id = "1" class = "foo" > <h2>sample text with
                <code>inline tag</code>
</h2> < pre id = "demo" class = "foo bar" > foo < /pre>
<pre id="output" class="goo">goo</pre > <input id="execute" type="button" value="execute"/> < /div>`

console.log(str)

var json = html2json(str)
console.log(json)