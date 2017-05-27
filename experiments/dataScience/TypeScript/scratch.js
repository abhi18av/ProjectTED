/*
var html2json = require('html2json').html2json;
var json2html = require('html2json').json2html;



var str = `< div id = "1" class = "foo" > <h2>sample text with
                <code>inline tag</code>
</h2> < pre id = "demo" class = "foo bar" > foo < /pre>
<pre id="output" class="goo">goo</pre > <input id="execute" type="button" value="execute"/> < /div>`

console.log(str)

var json = html2json(str)
console.log(json)

*/


var parser = require('xml2json');

var xml = "<foo attr=\"value\">bar</foo>";
console.log("input -> %s", xml)

// xml to json
var json = parser.toJson(xml);
console.log("to json -> %s", json);

// json to xml
var xml = parser.toXml(json);
console.log("back to xml -> %s", xml)