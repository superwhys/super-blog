"use strict";(self["webpackChunksuperBlog"]=self["webpackChunksuperBlog"]||[]).push([[41],{471:function(t,s,a){a.r(s),a.d(s,{default:function(){return c}});var i=function(){var t=this,s=t._self._c;return s("div",{staticClass:"tagsPage"},[s("base-page",{attrs:{"show-right":!1},scopedSlots:t._u([{key:"blogHeaderImgInnerShow",fn:function(){return[s("div",{staticClass:"flyInFromTop"},[t._v(" "+t._s(t.tag)+" ")])]},proxy:!0},{key:"blogLeftMainShow",fn:function(){return[s("all-tags-module",{staticClass:"flyInFromLeft",staticStyle:{margin:"20px"}}),t._l(t.tagInfoList.tagGroup,(function(a,i){return s("div",{key:i,staticClass:"tagsGroup flyInFromBottom"},[s("div",{staticClass:"tagsGroupKey el-icon-paperclip"},[t._v(" "+t._s(i)+" ")]),s("ul",{staticClass:"tagsGroupPosts"},t._l(a,(function(a){return s("router-link",{key:a.info.title,staticClass:"tagsGroupPost",attrs:{to:"/post"+a.info.toEndPoint,tag:"li"}},[s("div",{staticClass:"tagsGroupPostTitle"},[t._v(" "+t._s(a.info.title)+" ")]),s("div",{staticClass:"tagsGroupPostSubTitle"},[t._v(" "+t._s(a.info.subTitle)+" ")])])})),1)])}))]},proxy:!0},{key:"blogRightMainShow",fn:function(){return[s("div",{staticStyle:{display:"none"}})]},proxy:!0}])})],1)},e=[],o=a(9506),g=a(237),n=a(4203),r=a(2428),l=a(41),u={name:"tagsPage",components:{AllTagsModule:l.Z,TagItem:r.Z,BasePage:o.Z},data(){return{tag:"",tagInfoList:new n.QS({})}},created(){let t=void 0===this.$route.params.tag?"":this.$route.params.tag;this.tag=t;let s=this.$store.getters.getTagGroup(t);void 0===s?((0,g.G)(this.tag).then((s=>{console.debug(s),s.data&&(this.tagInfoList=new n.QS(s.data),this.$store.commit("setTagGroup",{tag:t,tagGrp:this.tagInfoList}))})),this.tag=""===this.tag?"Tags":`Tag: ${this.tag}`):this.tagInfoList=s},watch:{$route(t,s){if("tag"!==t.name)return;let a=void 0===this.$route.params.tag?"":this.$route.params.tag;this.tag=a;let i=this.$store.getters.getTagGroup(a);void 0===i?((0,g.G)(this.tag).then((t=>{console.debug(t),t.data&&(this.tagInfoList=new n.QS(t.data),this.$store.commit("setTagGroup",{tag:a,tagGrp:this.tagInfoList}))})),this.tag=""===this.tag?"Tags":`Tag: ${this.tag}`):this.tagInfoList=i}}},h=u,p=a(1001),f=(0,p.Z)(h,i,e,!1,null,"13f540f0",null),c=f.exports}}]);