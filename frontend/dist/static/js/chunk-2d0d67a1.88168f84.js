(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-2d0d67a1"],{"735d":function(n,a){(function(n){var a=n.languages.javadoclike={parameter:{pattern:/(^\s*(?:\/{3}|\*|\/\*\*)\s*@(?:param|arg|arguments)\s+)\w+/m,lookbehind:!0},keyword:{pattern:/(^\s*(?:\/{3}|\*|\/\*\*)\s*|\{)@[a-z][a-zA-Z-]+\b/m,lookbehind:!0},punctuation:/[{}]/};function e(a,e){var t="doc-comment",o=n.languages[a];if(o){var r=o[t];if(!r){var i={};i[t]={pattern:/(^|[^\\])\/\*\*[^/][\s\S]*?(?:\*\/|$)/,lookbehind:!0,alias:"comment"},o=n.languages.insertBefore(a,"comment",i),r=o[t]}if(r instanceof RegExp&&(r=o[t]={pattern:r}),Array.isArray(r))for(var s=0,p=r.length;s<p;s++)r[s]instanceof RegExp&&(r[s]={pattern:r[s]}),e(r[s]);else e(r)}}function t(n,a){"string"===typeof n&&(n=[n]),n.forEach((function(n){e(n,(function(n){n.inside||(n.inside={}),n.inside.rest=a}))}))}Object.defineProperty(a,"addSupport",{value:t}),a.addSupport(["java","javascript","php"],a)})(Prism)}}]);