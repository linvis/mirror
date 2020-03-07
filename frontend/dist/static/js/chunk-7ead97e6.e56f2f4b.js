(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-7ead97e6"],{"0bdc":function(t,e,i){},c3fdc:function(t,e,i){"use strict";i.r(e);var n=i("dbe2");i("0bdc");e["default"]=n["a"]},dbe2:function(t,e,i){"use strict";(function(t,n){i("7f7f"),i("28a5"),i("4917");var s=i("7618"),r=(i("a481"),i("17fb")),a=i.n(r),h=i("c155"),o=i.n(h),c=i("27d6"),l=i.n(c);
/** js sequence diagrams 2.0.1
 *  https://bramp.github.io/js-sequence-diagrams/
 *  (c) 2012-2017 Andrew Brampton (bramp.net)
 *  @license Simplified BSD license.
 */
function u(){this.title=void 0,this.actors=[],this.signals=[]}u.prototype.getActor=function(t,e){var i;t=t.trim();var n=this.actors;for(i in n)if(n[i].alias==t)return n[i];return i=n.push(new u.Actor(t,e||t,n.length)),n[i-1]},u.prototype.getActorWithAlias=function(t){t=t.trim();var e,i,n=/([\s\S]+) as (\S+)$/im.exec(t);return n?(i=n[1].trim(),e=n[2].trim()):i=e=t,this.getActor(e,i)},u.prototype.setTitle=function(t){this.title=t},u.prototype.addSignal=function(t){this.signals.push(t)},u.Actor=function(t,e,i){this.alias=t,this.name=e,this.index=i},u.Signal=function(t,e,i,n){this.type="Signal",this.actorA=t,this.actorB=i,this.linetype=3&e,this.arrowtype=e>>2&3,this.message=n},u.Signal.prototype.isSelf=function(){return this.actorA.index==this.actorB.index},u.Note=function(t,e,i){if(this.type="Note",this.actor=t,this.placement=e,this.message=i,this.hasManyActors()&&t[0]==t[1])throw new Error("Note should be over two different actors")},u.Note.prototype.hasManyActors=function(){return a.a.isArray(this.actor)},u.unescape=function(t){return t.trim().replace(/^"(.*)"$/m,"$1").replace(/\\n/gm,"\n")},u.LINETYPE={SOLID:0,DOTTED:1},u.ARROWTYPE={FILLED:0,OPEN:1},u.PLACEMENT={LEFTOF:0,RIGHTOF:1,OVER:2},"function"!==typeof Object.getPrototypeOf&&("object"===Object(s["a"])("test".__proto__)?Object.getPrototypeOf=function(t){return t.__proto__}:Object.getPrototypeOf=function(t){return t.constructor.prototype});var p=function(){function t(){this.yy={}}var e=function(t,e,i,n){for(i=i||{},n=t.length;n--;i[t[n]]=e);return i},i=[5,8,9,13,15,24],n=[1,13],s=[1,17],r=[24,29,30],a={trace:function(){},yy:{},symbols_:{error:2,start:3,document:4,EOF:5,line:6,statement:7,NL:8,participant:9,actor_alias:10,signal:11,note_statement:12,title:13,message:14,note:15,placement:16,actor:17,over:18,actor_pair:19,",":20,left_of:21,right_of:22,signaltype:23,ACTOR:24,linetype:25,arrowtype:26,LINE:27,DOTLINE:28,ARROW:29,OPENARROW:30,MESSAGE:31,$accept:0,$end:1},terminals_:{2:"error",5:"EOF",8:"NL",9:"participant",13:"title",15:"note",18:"over",20:",",21:"left_of",22:"right_of",24:"ACTOR",27:"LINE",28:"DOTLINE",29:"ARROW",30:"OPENARROW",31:"MESSAGE"},productions_:[0,[3,2],[4,0],[4,2],[6,1],[6,1],[7,2],[7,1],[7,1],[7,2],[12,4],[12,4],[19,1],[19,3],[16,1],[16,1],[11,4],[17,1],[10,1],[23,2],[23,1],[25,1],[25,1],[26,1],[26,1],[14,1]],performAction:function(t,e,i,n,s,r,a){var h=r.length-1;switch(s){case 1:return n.parser.yy;case 4:break;case 6:r[h];break;case 7:case 8:n.parser.yy.addSignal(r[h]);break;case 9:n.parser.yy.setTitle(r[h]);break;case 10:this.$=new u.Note(r[h-1],r[h-2],r[h]);break;case 11:this.$=new u.Note(r[h-1],u.PLACEMENT.OVER,r[h]);break;case 12:case 20:this.$=r[h];break;case 13:this.$=[r[h-2],r[h]];break;case 14:this.$=u.PLACEMENT.LEFTOF;break;case 15:this.$=u.PLACEMENT.RIGHTOF;break;case 16:this.$=new u.Signal(r[h-3],r[h-2],r[h-1],r[h]);break;case 17:this.$=n.parser.yy.getActor(u.unescape(r[h]));break;case 18:this.$=n.parser.yy.getActorWithAlias(u.unescape(r[h]));break;case 19:this.$=r[h-1]|r[h]<<2;break;case 21:this.$=u.LINETYPE.SOLID;break;case 22:this.$=u.LINETYPE.DOTTED;break;case 23:this.$=u.ARROWTYPE.FILLED;break;case 24:this.$=u.ARROWTYPE.OPEN;break;case 25:this.$=u.unescape(r[h].substring(1))}},table:[e(i,[2,2],{3:1,4:2}),{1:[3]},{5:[1,3],6:4,7:5,8:[1,6],9:[1,7],11:8,12:9,13:[1,10],15:[1,12],17:11,24:n},{1:[2,1]},e(i,[2,3]),e(i,[2,4]),e(i,[2,5]),{10:14,24:[1,15]},e(i,[2,7]),e(i,[2,8]),{14:16,31:s},{23:18,25:19,27:[1,20],28:[1,21]},{16:22,18:[1,23],21:[1,24],22:[1,25]},e([20,27,28,31],[2,17]),e(i,[2,6]),e(i,[2,18]),e(i,[2,9]),e(i,[2,25]),{17:26,24:n},{24:[2,20],26:27,29:[1,28],30:[1,29]},e(r,[2,21]),e(r,[2,22]),{17:30,24:n},{17:32,19:31,24:n},{24:[2,14]},{24:[2,15]},{14:33,31:s},{24:[2,19]},{24:[2,23]},{24:[2,24]},{14:34,31:s},{14:35,31:s},{20:[1,36],31:[2,12]},e(i,[2,16]),e(i,[2,10]),e(i,[2,11]),{17:37,24:n},{31:[2,13]}],defaultActions:{3:[2,1],24:[2,14],25:[2,15],27:[2,19],28:[2,23],29:[2,24],37:[2,13]},parseError:function(t,e){if(!e.recoverable)throw new Error(t);this.trace(t)},parse:function(t){function e(){var t;return t=d.lex()||p,"number"!=typeof t&&(t=i.symbols_[t]||t),t}var i=this,n=[0],s=[null],r=[],a=this.table,h="",o=0,c=0,l=0,u=2,p=1,y=r.slice.call(arguments,1),d=Object.create(this.lexer),f={yy:{}};for(var g in this.yy)Object.prototype.hasOwnProperty.call(this.yy,g)&&(f.yy[g]=this.yy[g]);d.setInput(t,f.yy),f.yy.lexer=d,f.yy.parser=this,"undefined"==typeof d.yylloc&&(d.yylloc={});var m=d.yylloc;r.push(m);var _=d.options&&d.options.ranges;"function"==typeof f.yy.parseError?this.parseError=f.yy.parseError:this.parseError=Object.getPrototypeOf(this).parseError;for(var w,x,k,v,b,E,S,A,T,O={};;){if(k=n[n.length-1],this.defaultActions[k]?v=this.defaultActions[k]:(null!==w&&"undefined"!=typeof w||(w=e()),v=a[k]&&a[k][w]),"undefined"==typeof v||!v.length||!v[0]){var L="";for(E in T=[],a[k])this.terminals_[E]&&E>u&&T.push("'"+this.terminals_[E]+"'");L=d.showPosition?"Parse error on line "+(o+1)+":\n"+d.showPosition()+"\nExpecting "+T.join(", ")+", got '"+(this.terminals_[w]||w)+"'":"Parse error on line "+(o+1)+": Unexpected "+(w==p?"end of input":"'"+(this.terminals_[w]||w)+"'"),this.parseError(L,{text:d.match,token:this.terminals_[w]||w,line:d.yylineno,loc:m,expected:T})}if(v[0]instanceof Array&&v.length>1)throw new Error("Parse Error: multiple actions possible at state: "+k+", token: "+w);switch(v[0]){case 1:n.push(w),s.push(d.yytext),r.push(d.yylloc),n.push(v[1]),w=null,x?(w=x,x=null):(c=d.yyleng,h=d.yytext,o=d.yylineno,m=d.yylloc,l>0&&l--);break;case 2:if(S=this.productions_[v[1]][1],O.$=s[s.length-S],O._$={first_line:r[r.length-(S||1)].first_line,last_line:r[r.length-1].last_line,first_column:r[r.length-(S||1)].first_column,last_column:r[r.length-1].last_column},_&&(O._$.range=[r[r.length-(S||1)].range[0],r[r.length-1].range[1]]),b=this.performAction.apply(O,[h,c,o,f.yy,v[1],s,r].concat(y)),"undefined"!=typeof b)return b;S&&(n=n.slice(0,-1*S*2),s=s.slice(0,-1*S),r=r.slice(0,-1*S)),n.push(this.productions_[v[1]][0]),s.push(O.$),r.push(O._$),A=a[n[n.length-2]][n[n.length-1]],n.push(A);break;case 3:return!0}}return!0}},h=function(){var t={EOF:1,parseError:function(t,e){if(!this.yy.parser)throw new Error(t);this.yy.parser.parseError(t,e)},setInput:function(t,e){return this.yy=e||this.yy||{},this._input=t,this._more=this._backtrack=this.done=!1,this.yylineno=this.yyleng=0,this.yytext=this.matched=this.match="",this.conditionStack=["INITIAL"],this.yylloc={first_line:1,first_column:0,last_line:1,last_column:0},this.options.ranges&&(this.yylloc.range=[0,0]),this.offset=0,this},input:function(){var t=this._input[0];this.yytext+=t,this.yyleng++,this.offset++,this.match+=t,this.matched+=t;var e=t.match(/(?:\r\n?|\n).*/g);return e?(this.yylineno++,this.yylloc.last_line++):this.yylloc.last_column++,this.options.ranges&&this.yylloc.range[1]++,this._input=this._input.slice(1),t},unput:function(t){var e=t.length,i=t.split(/(?:\r\n?|\n)/g);this._input=t+this._input,this.yytext=this.yytext.substr(0,this.yytext.length-e),this.offset-=e;var n=this.match.split(/(?:\r\n?|\n)/g);this.match=this.match.substr(0,this.match.length-1),this.matched=this.matched.substr(0,this.matched.length-1),i.length-1&&(this.yylineno-=i.length-1);var s=this.yylloc.range;return this.yylloc={first_line:this.yylloc.first_line,last_line:this.yylineno+1,first_column:this.yylloc.first_column,last_column:i?(i.length===n.length?this.yylloc.first_column:0)+n[n.length-i.length].length-i[0].length:this.yylloc.first_column-e},this.options.ranges&&(this.yylloc.range=[s[0],s[0]+this.yyleng-e]),this.yyleng=this.yytext.length,this},more:function(){return this._more=!0,this},reject:function(){return this.options.backtrack_lexer?(this._backtrack=!0,this):this.parseError("Lexical error on line "+(this.yylineno+1)+". You can only invoke reject() in the lexer when the lexer is of the backtracking persuasion (options.backtrack_lexer = true).\n"+this.showPosition(),{text:"",token:null,line:this.yylineno})},less:function(t){this.unput(this.match.slice(t))},pastInput:function(){var t=this.matched.substr(0,this.matched.length-this.match.length);return(t.length>20?"...":"")+t.substr(-20).replace(/\n/g,"")},upcomingInput:function(){var t=this.match;return t.length<20&&(t+=this._input.substr(0,20-t.length)),(t.substr(0,20)+(t.length>20?"...":"")).replace(/\n/g,"")},showPosition:function(){var t=this.pastInput(),e=new Array(t.length+1).join("-");return t+this.upcomingInput()+"\n"+e+"^"},test_match:function(t,e){var i,n,s;if(this.options.backtrack_lexer&&(s={yylineno:this.yylineno,yylloc:{first_line:this.yylloc.first_line,last_line:this.last_line,first_column:this.yylloc.first_column,last_column:this.yylloc.last_column},yytext:this.yytext,match:this.match,matches:this.matches,matched:this.matched,yyleng:this.yyleng,offset:this.offset,_more:this._more,_input:this._input,yy:this.yy,conditionStack:this.conditionStack.slice(0),done:this.done},this.options.ranges&&(s.yylloc.range=this.yylloc.range.slice(0))),n=t[0].match(/(?:\r\n?|\n).*/g),n&&(this.yylineno+=n.length),this.yylloc={first_line:this.yylloc.last_line,last_line:this.yylineno+1,first_column:this.yylloc.last_column,last_column:n?n[n.length-1].length-n[n.length-1].match(/\r?\n?/)[0].length:this.yylloc.last_column+t[0].length},this.yytext+=t[0],this.match+=t[0],this.matches=t,this.yyleng=this.yytext.length,this.options.ranges&&(this.yylloc.range=[this.offset,this.offset+=this.yyleng]),this._more=!1,this._backtrack=!1,this._input=this._input.slice(t[0].length),this.matched+=t[0],i=this.performAction.call(this,this.yy,this,e,this.conditionStack[this.conditionStack.length-1]),this.done&&this._input&&(this.done=!1),i)return i;if(this._backtrack){for(var r in s)this[r]=s[r];return!1}return!1},next:function(){if(this.done)return this.EOF;var t,e,i,n;this._input||(this.done=!0),this._more||(this.yytext="",this.match="");for(var s=this._currentRules(),r=0;r<s.length;r++)if(i=this._input.match(this.rules[s[r]]),i&&(!e||i[0].length>e[0].length)){if(e=i,n=r,this.options.backtrack_lexer){if(t=this.test_match(i,s[r]),!1!==t)return t;if(this._backtrack){e=!1;continue}return!1}if(!this.options.flex)break}return e?(t=this.test_match(e,s[n]),!1!==t&&t):""===this._input?this.EOF:this.parseError("Lexical error on line "+(this.yylineno+1)+". Unrecognized text.\n"+this.showPosition(),{text:"",token:null,line:this.yylineno})},lex:function(){var t=this.next();return t||this.lex()},begin:function(t){this.conditionStack.push(t)},popState:function(){var t=this.conditionStack.length-1;return t>0?this.conditionStack.pop():this.conditionStack[0]},_currentRules:function(){return this.conditionStack.length&&this.conditionStack[this.conditionStack.length-1]?this.conditions[this.conditionStack[this.conditionStack.length-1]].rules:this.conditions.INITIAL.rules},topState:function(t){return t=this.conditionStack.length-1-Math.abs(t||0),t>=0?this.conditionStack[t]:"INITIAL"},pushState:function(t){this.begin(t)},stateStackSize:function(){return this.conditionStack.length},options:{"case-insensitive":!0},performAction:function(t,e,i,n){switch(i){case 0:return 8;case 1:break;case 2:break;case 3:return 9;case 4:return 21;case 5:return 22;case 6:return 18;case 7:return 15;case 8:return 13;case 9:return 20;case 10:return 24;case 11:return 24;case 12:return 28;case 13:return 27;case 14:return 30;case 15:return 29;case 16:return 31;case 17:return 5;case 18:return"INVALID"}},rules:[/^(?:[\r\n]+)/i,/^(?:\s+)/i,/^(?:#[^\r\n]*)/i,/^(?:participant\b)/i,/^(?:left of\b)/i,/^(?:right of\b)/i,/^(?:over\b)/i,/^(?:note\b)/i,/^(?:title\b)/i,/^(?:,)/i,/^(?:[^\->:,\r\n"]+)/i,/^(?:"[^"]+")/i,/^(?:--)/i,/^(?:-)/i,/^(?:>>)/i,/^(?:>)/i,/^(?:[^\r\n]+)/i,/^(?:$)/i,/^(?:.)/i],conditions:{INITIAL:{rules:[0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18],inclusive:!0}}};return t}();return a.lexer=h,t.prototype=a,a.Parser=t,new t}();function y(t,e){a.a.extend(this,e),this.name="ParseError",this.message=t||""}"undefined"!=typeof exports&&(exports.parser=p,exports.Parser=p.Parser,exports.parse=function(){return p.parse.apply(p,arguments)},exports.main=function(e){e[1]||(console.log("Usage: "+e[0]+" FILE"),t.exit(1));var n=i("3e8f").readFileSync(i("df7c").normalize(e[1]),"utf8");return exports.parser.parse(n)},i.c[i.s]===n&&exports.main(t.argv.slice(1))),y.prototype=new Error,u.ParseError=y,u.parse=function(t){p.yy=new u,p.yy.parseError=function(t,e){throw new y(t,e)};var e=p.parse(t);return delete e.parseError,e};var d=10,f=10,g=10,m=5,_=5,w=10,x=5,k=15,v=0,b=5,E=20,S=u.PLACEMENT,A=u.LINETYPE,T=u.ARROWTYPE,O=0,L=1;function N(t){this.message=t}function I(t,e){if(!t)throw new N(e)}function P(t,e){u.themes[t]=e}function M(t){return t.x+t.width/2}function F(t){return t.y+t.height/2}function R(t,e,i){return t<e?e:t>i?i:t}function B(t,e,i,n){I(a.a.all([t,i,e,n],a.a.isFinite),"x1,x2,y1,y2 must be numeric");var s=Math.sqrt((i-t)*(i-t)+(n-e)*(n-e))/25,r=R(Math.random(),.2,.8),h=R(Math.random(),.2,.8),o=Math.random()>.5?s:-s,c=Math.random()>.5?s:-s,l={x:(i-t)*r+t+o,y:(n-e)*r+e+c},u={x:(i-t)*h+t-o,y:(n-e)*h+e-c};return"C"+l.x.toFixed(1)+","+l.y.toFixed(1)+" "+u.x.toFixed(1)+","+u.y.toFixed(1)+" "+i.toFixed(1)+","+n.toFixed(1)}function $(t,e,i,n){return I(a.a.all([t,e,i,n],a.a.isFinite),"x, y, w, h must be numeric"),"M"+t+","+e+B(t,e,t+i,e)+B(t+i,e,t+i,e+n)+B(t+i,e+n,t,e+n)+B(t,e+n,t,e)}function H(t,e,i,n){return I(a.a.all([t,i,e,n],a.a.isFinite),"x1,x2,y1,y2 must be numeric"),"M"+t.toFixed(1)+","+e.toFixed(1)+B(t,e,i,n)}N.prototype.toString=function(){return"AssertException: "+this.message},String.prototype.trim||(String.prototype.trim=function(){return this.replace(/^\s+|\s+$/g,"")}),u.themes={};var C=function(t,e){this.init(t,e)};if(a.a.extend(C.prototype,{init:function(t,e){this.diagram=t,this.actorsHeight_=0,this.signalsHeight_=0,this.title_=void 0},setupPaper:function(t){},draw:function(t){this.setupPaper(t),this.layout();var e=this.title_?this.title_.height:0,i=d+e;this.drawTitle(),this.drawActors(i),this.drawSignals(i+this.actorsHeight_)},layout:function(){var t=this.diagram,e=this.font_,i=t.actors,n=t.signals;if(t.width=0,t.height=0,t.title){var s=this.title_={},r=this.textBBox(t.title,e);s.textBB=r,s.message=t.title,s.width=r.width+2*(b+v),s.height=r.height+2*(b+v),s.x=d,s.y=d,t.width+=s.width,t.height+=s.height}function h(t,e,n){I(t<e,"a must be less than or equal to b"),t<0?(e=i[e],e.x=Math.max(n-e.width/2,e.x)):e>=i.length?(t=i[t],t.paddingRight=Math.max(n,t.paddingRight)):(t=i[t],t.distances[e]=Math.max(n,t.distances[e]?t.distances[e]:0))}a.a.each(i,(function(t){var i=this.textBBox(t.name,e);t.textBB=i,t.x=0,t.y=0,t.width=i.width+2*(g+f),t.height=i.height+2*(g+f),t.distances=[],t.paddingRight=0,this.actorsHeight_=Math.max(t.height,this.actorsHeight_)}),this),a.a.each(n,(function(t){var i,n,s=this.textBBox(t.message,e);t.textBB=s,t.width=s.width,t.height=s.height;var r=0;if("Signal"==t.type)t.width+=2*(m+_),t.height+=2*(m+_),t.isSelf()?(i=t.actorA.index,n=i+1,t.width+=E):(i=Math.min(t.actorA.index,t.actorB.index),n=Math.max(t.actorA.index,t.actorB.index));else{if("Note"!=t.type)throw new Error("Unhandled signal type:"+t.type);if(t.width+=2*(w+x),t.height+=2*(w+x),r=2*f,t.placement==S.LEFTOF)n=t.actor.index,i=n-1;else if(t.placement==S.RIGHTOF)i=t.actor.index,n=i+1;else if(t.placement==S.OVER&&t.hasManyActors())i=Math.min(t.actor[0].index,t.actor[1].index),n=Math.max(t.actor[0].index,t.actor[1].index),r=-(2*x+2*k);else if(t.placement==S.OVER)return i=t.actor.index,h(i-1,i,t.width/2),h(i,i+1,t.width/2),void(this.signalsHeight_+=t.height)}h(i,n,t.width+r),this.signalsHeight_+=t.height}),this);var o=0;return a.a.each(i,(function(t){t.x=Math.max(o,t.x),a.a.each(t.distances,(function(e,n){"undefined"!=typeof e&&(n=i[n],e=Math.max(e,t.width/2,n.width/2),n.x=Math.max(n.x,t.x+t.width/2+e-n.width/2))})),o=t.x+t.width+t.paddingRight}),this),t.width=Math.max(o,t.width),t.width+=2*d,t.height+=2*d+2*this.actorsHeight_+this.signalsHeight_,this},textBBox:function(t,e){},drawTitle:function(){var t=this.title_;t&&this.drawTextBox(t,t.message,v,b,this.font_,O)},drawActors:function(t){var e=t;a.a.each(this.diagram.actors,(function(t){this.drawActor(t,e,this.actorsHeight_),this.drawActor(t,e+this.actorsHeight_+this.signalsHeight_,this.actorsHeight_);var i=M(t);this.drawLine(i,e+this.actorsHeight_-f,i,e+this.actorsHeight_+f+this.signalsHeight_)}),this)},drawActor:function(t,e,i){t.y=e,t.height=i,this.drawTextBox(t,t.name,f,g,this.font_,L)},drawSignals:function(t){var e=t;a.a.each(this.diagram.signals,(function(t){"Signal"==t.type?t.isSelf()?this.drawSelfSignal(t,e):this.drawSignal(t,e):"Note"==t.type&&this.drawNote(t,e),e+=t.height}),this)},drawSelfSignal:function(t,e){I(t.isSelf(),"signal must be a self signal");var i=t.textBB,n=M(t.actorA),s=n+E+_,r=e+_+t.height/2+i.y;this.drawText(s,r,t.message,this.font_,O);var a=e+m+_,h=a+t.height-2*m-_;this.drawLine(n,a,n+E,a,t.linetype),this.drawLine(n+E,a,n+E,h,t.linetype),this.drawLine(n+E,h,n,h,t.linetype,t.arrowtype)},drawSignal:function(t,e){var i=M(t.actorA),n=M(t.actorB),s=(n-i)/2+i,r=e+m+2*_;this.drawText(s,r,t.message,this.font_,L),r=e+t.height-m-_,this.drawLine(i,r,n,r,t.linetype,t.arrowtype)},drawNote:function(t,e){t.y=e;var i=t.hasManyActors()?t.actor[0]:t.actor,n=M(i);switch(t.placement){case S.RIGHTOF:t.x=n+f;break;case S.LEFTOF:t.x=n-f-t.width;break;case S.OVER:if(t.hasManyActors()){var s=M(t.actor[1]),r=k+x;t.x=Math.min(n,s)-r,t.width=Math.max(n,s)+r-t.x}else t.x=n-t.width/2;break;default:throw new Error("Unhandled note placement: "+t.placement)}return this.drawTextBox(t,t.message,w,x,this.font_,O)},drawTextBox:function(t,e,i,n,s,r){var a=t.x+i,h=t.y+i,o=t.width-2*i,c=t.height-2*i;return this.drawRect(a,h,o,c),r==L?(a=M(t),h=F(t)):(a+=n,h+=n),this.drawText(a,h,e,s,r)}}),"undefined"!=typeof o.a){var G="http://www.w3.org/2000/svg",D={stroke:"#000000","stroke-width":2,fill:"none"},j={stroke:"#000000","stroke-width":2,fill:"#fff"},W={},z=function(t,e,i){a.a.defaults(e,{"css-class":"simple","font-size":16,"font-family":"Andale Mono, monospace"}),this.init(t,e,i)};a.a.extend(z.prototype,C.prototype,{init:function(t,e,i){C.prototype.init.call(this,t),this.paper_=void 0,this.cssClass_=e["css-class"]||void 0,this.font_={"font-size":e["font-size"],"font-family":e["font-family"]};var n=this.arrowTypes_={};n[T.FILLED]="Block",n[T.OPEN]="Open";var s=this.lineTypes_={};s[A.SOLID]="",s[A.DOTTED]="6,2";var r=this;this.waitForFont((function(){i(r)}))},waitForFont:function(t){var e=this.font_["font-family"];if("undefined"==typeof l.a)throw new Error("WebFont is required (https://github.com/typekit/webfontloader).");W[e]?t():l.a.load({custom:{families:[e]},classes:!1,active:function(){W[e]=!0,t()},inactive:function(){W[e]=!0,t()}})},addDescription:function(t,e){var i=document.createElementNS(G,"desc");i.appendChild(document.createTextNode(e)),t.appendChild(i)},setupPaper:function(t){var e=document.createElementNS(G,"svg");t.appendChild(e),this.addDescription(e,this.diagram.title||""),this.paper_=o()(e),this.paper_.addClass("sequence"),this.cssClass_&&this.paper_.addClass(this.cssClass_),this.beginGroup();var i=this.arrowMarkers_={},n=this.paper_.path("M 0 0 L 5 2.5 L 0 5 z");i[T.FILLED]=n.marker(0,0,5,5,5,2.5).attr({id:"markerArrowBlock"}),n=this.paper_.path("M 9.6,8 1.92,16 0,13.7 5.76,8 0,2.286 1.92,0 9.6,8 z"),i[T.OPEN]=n.marker(0,0,9.6,16,9.6,8).attr({markerWidth:"4",id:"markerArrowOpen"})},layout:function(){C.prototype.layout.call(this),this.paper_.attr({width:this.diagram.width+"px",height:this.diagram.height+"px"})},textBBox:function(t,e){var i=this.createText(t,e),n=i.getBBox();return i.remove(),n},pushToStack:function(t){return this._stack.push(t),t},beginGroup:function(){this._stack=[]},finishGroup:function(){var t=this.paper_.group.apply(this.paper_,this._stack);return this.beginGroup(),t},createText:function(t,e){t=a.a.invoke(t.split("\n"),"trim");var i=this.paper_.text(0,0,t);return i.attr(e||{}),t.length>1&&i.selectAll("tspan:nth-child(n+2)").attr({dy:"1.2em",x:0}),i},drawLine:function(t,e,i,n,s,r){var a=this.paper_.line(t,e,i,n).attr(D);return void 0!==s&&a.attr("strokeDasharray",this.lineTypes_[s]),void 0!==r&&a.attr("markerEnd",this.arrowMarkers_[r]),this.pushToStack(a)},drawRect:function(t,e,i,n){var s=this.paper_.rect(t,e,i,n).attr(j);return this.pushToStack(s)},drawText:function(t,e,i,n,s){var r=this.createText(i,n),a=r.getBBox();return s==L&&(t-=a.width/2,e-=a.height/2),r.attr({x:t-a.x,y:e-a.y}),r.selectAll("tspan").attr({x:t}),this.pushToStack(r),r},drawTitle:function(){return this.beginGroup(),C.prototype.drawTitle.call(this),this.finishGroup().addClass("title")},drawActor:function(t,e,i){return this.beginGroup(),C.prototype.drawActor.call(this,t,e,i),this.finishGroup().addClass("actor")},drawSignal:function(t,e){return this.beginGroup(),C.prototype.drawSignal.call(this,t,e),this.finishGroup().addClass("signal")},drawSelfSignal:function(t,e){return this.beginGroup(),C.prototype.drawSelfSignal.call(this,t,e),this.finishGroup().addClass("signal")},drawNote:function(t,e){return this.beginGroup(),C.prototype.drawNote.call(this,t,e),this.finishGroup().addClass("note")}});var Y=function(t,e,i){a.a.defaults(e,{"css-class":"hand","font-size":16,"font-family":"danielbd"}),this.init(t,e,i)};a.a.extend(Y.prototype,z.prototype,{drawLine:function(t,e,i,n,s,r){var a=this.paper_.path(H(t,e,i,n)).attr(D);return void 0!==s&&a.attr("strokeDasharray",this.lineTypes_[s]),void 0!==r&&a.attr("markerEnd",this.arrowMarkers_[r]),this.pushToStack(a)},drawRect:function(t,e,i,n){var s=this.paper_.path($(t,e,i,n)).attr(j);return this.pushToStack(s)}}),P("snapSimple",z),P("snapHand",Y)}if("undefined"==typeof Raphael&&"undefined"==typeof o.a)throw new Error("Raphael or Snap.svg is required to be included.");if(a.a.isEmpty(u.themes))throw new Error("No themes were registered. Please call registerTheme(...).");u.themes.hand=u.themes.snapHand||u.themes.raphaelHand,u.themes.simple=u.themes.snapSimple||u.themes.raphaelSimple,u.prototype.drawSVG=function(t,e){var i={theme:"hand"};if(e=a.a.defaults(e||{},i),!(e.theme in u.themes))throw new Error("Unsupported theme: "+e.theme);var n=a.a.isString(t)?document.getElementById(t):t;if(null===n||!n.tagName)throw new Error("Invalid container: "+t);var s=u.themes[e.theme];new s(this,e,(function(t){t.draw(n)}))},e["a"]=u}).call(this,i("4362"),i("dd40")(t))}}]);