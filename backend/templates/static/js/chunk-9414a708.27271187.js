(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-9414a708"],{"0b54":function(t,e,a){},"18e9":function(t,e,a){},"3a33":function(t,e,a){"use strict";var r=a("0b54"),n=a.n(r);n.a},9899:function(t,e,a){"use strict";a.r(e);var r=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"mixin-components-container"},[a("el-row",[a("sleep-new")],1),t._v(" "),a("el-row",{staticStyle:{"margin-top":"10px"}},[a("sleep-chart")],1),t._v(" "),a("el-row",{staticStyle:{"margin-top":"10px"}},[a("sleep-analysis")],1)],1)},n=[],i=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("el-card",{staticClass:"box-card",attrs:{shadow:"hover"}},[a("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[a("span",[t._v("睡眠统计")])]),t._v(" "),a("el-tabs",{on:{"tab-click":t.handleTabClick}},[a("el-tab-pane",{attrs:{label:"睡眠时长",name:"first"}}),t._v(" "),a("el-tab-pane",{attrs:{label:"入睡时间",name:"second"}}),t._v(" "),a("el-tab-pane",{attrs:{label:"苏醒时间",name:"third"}})],1),t._v(" "),a("el-row",{attrs:{id:"sleep_chart"}})],1)},s=[],o=(a("6b54"),a("7f7f"),a("ea7f")),l=a.n(o),d=a("37d8"),c=a.n(d),u=a("b775");function h(t){return Object(u["a"])({url:"/record/submit/sleep",method:"post",data:t})}function p(t){return Object(u["a"])({url:"/record/query/sleep/",method:"get",params:{token:t}})}function _(t){return Object(u["a"])({url:"/record/analysis/sleep/",method:"get",params:{token:t}})}var m=a("d8ad5");c()(l.a);var f={data:function(){return{activeTab:"first",sleep_chart:null,dur_data:[],start_time_data:[],end_time_data:[],chart_data:[]}},created:function(){m["a"].$on("update-sleep-record",this.getSleepRecord)},mounted:function(){this.getSleepRecord()},beforeDestroy:function(){m["a"].$off("update-sleep-record",this.getSleepRecord)},methods:{getSleepRecord:function(){var t=this;p().then((function(e){console.log(e.data),t.dur_data=[],t.start_time_data=[],t.end_time_data=[];for(var a=0;a<e.data.date.length;a++)t.dur_data.push([e.data.date[a],e.data.duration[a]]),t.start_time_data.push([e.data.date[a],e.data.start_time[a]+1440]),t.end_time_data.push([e.data.date[a],e.data.end_time[a]+1440]);t.chart_data=t.dur_data,console.log(t.chart_data),t.initChart()}))},handleTabClick:function(t,e){console.log(t.name,this.activeTab),this.activeTab!==t.name&&("first"===t.name?this.chart_data=this.dur_data:"second"===t.name?this.chart_data=this.start_time_data:"third"===t.name&&(this.chart_data=this.end_time_data),this.activeTab=t.name,this.sleep_chart.series[0].update({data:this.chart_data},!1),this.sleep_chart.redraw())},timeToString:function(t){var e=parseInt(t/60),a=parseInt(t%60);return parseInt(e/10).toString()+(e%10).toString()+":"+parseInt(a/10).toString()+(a%10).toString()},initChart:function(){this.sleep_chart=l.a.stockChart("sleep_chart",{rangeSelector:{enabled:!0,allButtonsEnabled:!0,buttons:[{type:"week",count:1,text:"last 7 days"},{type:"week",count:2,text:"last 14 days"},{type:"month",text:"Month"}],buttonTheme:{width:60},selected:0},title:{text:"Sleep Analysis"},yAxis:{labels:{formatter:function(){this.value=this.value%1440;var t=parseInt(this.value/60),e=parseInt(this.value%60);return parseInt(t/10).toString()+(t%10).toString()+":"+parseInt(e/10).toString()+(e%10).toString()}},title:{text:"Hours"}},tooltip:{formatter:function(){var t=this.y%1440,e=parseInt(t/60),a=parseInt(t%60),r=parseInt(e/10).toString()+(e%10).toString()+":"+parseInt(a/10).toString()+(a%10).toString();return"Sleep about "+r}},series:[{name:"sleep time",data:this.chart_data}]})}}},g=f,v=a("2877"),b=Object(v["a"])(g,i,s,!1,null,null,null),S=b.exports,k=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("el-card",{staticClass:"box-card",attrs:{shadow:"hover"}},[a("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[a("span",[t._v("日常活动")])]),t._v(" "),a("el-form",{ref:"form",attrs:{"label-position":"left",model:t.form,"label-width":"80px"}},[a("el-form-item",{attrs:{label:"日期"}},[a("el-date-picker",{attrs:{type:"date",placeholder:"选择日期","picker-options":t.pickerOptions},model:{value:t.record_date,callback:function(e){t.record_date=e},expression:"record_date"}})],1),t._v(" "),a("el-form-item",{attrs:{label:"开始时间"}},[a("el-time-picker",{attrs:{format:"HH:mm",placeholder:"开始时间"},on:{change:t.handleStartTimeChange},model:{value:t.start_utc,callback:function(e){t.start_utc=e},expression:"start_utc"}})],1),t._v(" "),a("el-form-item",{attrs:{label:"结束时间"}},[a("el-time-picker",{attrs:{format:"HH:mm",placeholder:"结束时间"},on:{change:t.handleEndTimeChange},model:{value:t.end_utc,callback:function(e){t.end_utc=e},expression:"end_utc"}})],1),t._v(" "),a("el-form-item",{attrs:{label:"持续时间"}},[a("el-time-select",{attrs:{"picker-options":{start:"00:00",step:"00:10",end:"04:00"},placeholder:"持续时间"},model:{value:t.duration_utc,callback:function(e){t.duration_utc=e},expression:"duration_utc"}})],1),t._v(" "),a("el-form-item",[a("el-button",{attrs:{type:"primary"},on:{click:t.handleSubmit}},[t._v("立即创建")]),t._v(" "),a("el-button",[t._v("取消")])],1)],1)],1)},I=[],x={data:function(){return{start_time:0,end_time:0,duration:0,start_utc:0,end_utc:0,duration_utc:0,pickerOptions:{disabledDate:function(t){return t.getTime()>Date.now()}},record_date:new Date,form:{}}},created:function(){},methods:{handleStartTimeChange:function(t){if(this.start_time=60*this.start_utc.getHours()+this.start_utc.getMinutes(),0!==this.end_time){var e=function(t){var e=parseInt(t/60),a=parseInt(t%60);return parseInt(e/10).toString()+(e%10).toString()+":"+parseInt(a/10).toString()+(a%10).toString()};this.start_time>this.end_time&&(this.start_time=this.start_time-1440),this.duration=this.end_time-this.start_time,this.duration_utc=e(this.duration)}},handleEndTimeChange:function(t){this.end_time=60*this.end_utc.getHours()+this.end_utc.getMinutes();var e=function(t){var e=parseInt(t/60),a=parseInt(t%60);return parseInt(e/10).toString()+(e%10).toString()+":"+parseInt(a/10).toString()+(a%10).toString()};0!==this.start_time&&(this.start_time>this.end_time&&(this.start_time=this.start_time-1440),this.duration=this.end_time-this.start_time,this.duration_utc=e(this.duration))},handleSubmit:function(t){var e=this,a=Date.UTC(this.record_date.getFullYear(),this.record_date.getMonth(),this.record_date.getDate());console.log(this.record_date.getTimezoneOffset());var r={record_date:a,start_time:this.start_time,end_time:this.end_time,duration:this.duration};h(r).then((function(t){e.$notify({title:"成功",message:"",type:"success",duration:800}),m["a"].$emit("update-sleep-record")}))}}},w=x,y=(a("c5e5"),Object(v["a"])(w,k,I,!1,null,"e22a6406",null)),C=y.exports,T=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("el-card",{staticClass:"box-card",attrs:{shadow:"hover"}},[a("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[a("span",[t._v("睡眠分析")])]),t._v(" "),a("el-row",{attrs:{id:"sleep_analysis_chart"}})],1)},L=[];c()(l.a);var O={data:function(){return{sleep_analysis_chart:null,avgLastWeek:null,avgLastTwoWeek:null,avgLastMonth:null}},created:function(){},mounted:function(){this.featchDataj()},methods:{featchDataj:function(){var t=this;_().then((function(e){t.avgLastWeek=[e.data.duration[0],e.data.start_time[0],e.data.end_time[0]],t.avgLastTwoWeek=[e.data.duration[1],e.data.start_time[1],e.data.end_time[1]],t.avgLastMonth=[e.data.duration[2],e.data.start_time[2],e.data.end_time[2]],console.log(t.avgLastWeek),console.log(t.avgLastTwoWeek),console.log(t.avgLastMonth),t.initChart()}))},timeToString:function(t){var e=parseInt(t/60),a=parseInt(t%60);return parseInt(e/10).toString()+(e%10).toString()+":"+parseInt(a/10).toString()+(a%10).toString()},initChart:function(){this.sleep_analysis_chart=l.a.chart("sleep_analysis_chart",{chart:{type:"bar"},title:{text:"Sleep Analysis"},xAxis:{categories:["平均睡眠时长","平均入睡时间","平均苏醒时间"],title:{text:null}},yAxis:{labels:{formatter:function(){this.value=this.value%1440;var t=parseInt(this.value/60),e=parseInt(this.value%60);return parseInt(t/10).toString()+(t%10).toString()+":"+parseInt(e/10).toString()+(e%10).toString()}},title:{text:"Hours"}},tooltip:{formatter:function(){var t=this.y%1440,e=parseInt(t/60),a=parseInt(t%60),r=parseInt(e/10).toString()+(e%10).toString()+":"+parseInt(a/10).toString()+(a%10).toString();return"Sleep about "+r}},plotOptions:{bar:{dataLabels:{enabled:!0}}},legend:{layout:"vertical",align:"right",verticalAlign:"top",x:-40,y:80,floating:!0,borderWidth:1,backgroundColor:l.a.defaultOptions.legend.backgroundColor||"#FFFFFF",shadow:!0},credits:{enabled:!1},series:[{name:"This Week",data:this.avgLastWeek},{name:"Last Two Week",data:this.avgLastTwoWeek},{name:"Last Month",data:this.avgLastMonth}]})}}},W=O,j=Object(v["a"])(W,T,L,!1,null,null,null),M=j.exports,D={components:{SleepNew:C,SleepChart:S,SleepAnalysis:M}},H=D,$=(a("3a33"),Object(v["a"])(H,r,n,!1,null,"12c88dd8",null));e["default"]=$.exports},c5e5:function(t,e,a){"use strict";var r=a("18e9"),n=a.n(r);n.a},d8ad5:function(t,e,a){"use strict";a.d(e,"a",(function(){return n}));var r=a("2b0e"),n=new r["default"]}}]);