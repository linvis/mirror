(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-2d0cf120"],{"61c2":function(e,t,a){"use strict";a.r(t);var o=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("v-card",{staticClass:"mx-4",attrs:{"min-width":"30%"}},[a("v-row",[a("v-col",{attrs:{cols:"6"}},[a("v-card-title",[e._v("Settings")])],1),a("v-col",{attrs:{col:"6"}},[a("v-card-actions",{staticClass:"pr-4",attrs:{"align-end":""}},[a("v-spacer"),a("v-btn",{attrs:{color:"green accent-4"},on:{click:e.submitSettingProgram}},[e._v(" Update ")])],1)],1)],1),a("v-container",[a("form",[a("v-text-field",{attrs:{label:"Leetcode User Name",outlined:"",clearable:""},model:{value:e.leetcodeName,callback:function(t){e.leetcodeName=t},expression:"leetcodeName"}}),a("v-text-field",{attrs:{label:"Leetcode Cookies",outlined:"",clearable:""},model:{value:e.leetcodeCookies,callback:function(t){e.leetcodeCookies=t},expression:"leetcodeCookies"}}),a("v-text-field",{attrs:{label:"Github User Name",outlined:"",clearable:""},model:{value:e.githubName,callback:function(t){e.githubName=t},expression:"githubName"}})],1)])],1)},c=[],n=a("b775");function l(e){return Object(n["a"])({url:"/setting/program",method:"post",data:e})}function i(e){return Object(n["a"])({url:"/setting/program",method:"get",params:{token:e}})}var r={data:function(){return{leetcodeName:"",leetcodeCookies:"",githubName:""}},mounted:function(){this.featchData()},methods:{featchData:function(){var e=this;i().then((function(t){e.leetcodeName=t.data.leetcode_user_name,e.leetcodeCookies=t.data.cookie,e.githubName=t.data.github_user_name}))},handleSubmit:function(e){var t={leetcode_user_name:this.leetcodeName,cookie:this.leetcodeCookies,github_user_name:this.githubName};l(t).then((function(e){}))}}},s=r,d=a("2877"),u=a("6544"),m=a.n(u),b=a("8336"),h=a("b0af"),f=a("99d9"),v=a("62ad"),g=a("a523"),k=a("0fd9"),p=a("2fa4"),C=a("8654"),N=Object(d["a"])(s,o,c,!1,null,null,null);t["default"]=N.exports;m()(N,{VBtn:b["a"],VCard:h["a"],VCardActions:f["a"],VCardTitle:f["c"],VCol:v["a"],VContainer:g["a"],VRow:k["a"],VSpacer:p["a"],VTextField:C["a"]})}}]);
//# sourceMappingURL=chunk-2d0cf120.5b665014.js.map