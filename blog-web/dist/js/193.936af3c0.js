"use strict";(self["webpackChunksuperBlog"]=self["webpackChunksuperBlog"]||[]).push([[193],{811:function(t,e,o){o.d(e,{Z:function(){return C}});var r=function(){var t=this,e=t._self._c;return e("div",{staticClass:"basePage"},[e("el-container",{staticStyle:{"min-height":"100vh"}},[e("div",{staticClass:"header-menu",style:{backgroundColor:t.backgroundColor}},[e("div",{staticClass:"header-menu-left",style:{color:t.navTextColor}},[t._v(" SuperYong ")]),e("div",{staticClass:"header-menu-right"},[e("router-link",{style:{color:t.navTextColor},attrs:{to:"/"}},[t._v("HOME")]),e("router-link",{style:{color:t.navTextColor},attrs:{to:"/about"}},[t._v("ABOUT")]),e("router-link",{style:{color:t.navTextColor},attrs:{to:"/tag"}},[t._v("TAGS")])],1)]),e("el-header",{style:t.backgroundImageStyle},[e("blog-header",{staticStyle:{width:"100%",height:"100%"},scopedSlots:t._u([{key:"headerImgInnerShow",fn:function(){return[t._t("blogHeaderImgInnerShow")]},proxy:!0}],null,!0)})],1),e("el-main",{staticStyle:{overflow:"visible",height:"100%",padding:"50px 0","min-height":"calc(100vh - 420px - 5vh)"}},[e("div",{staticClass:"mainContainer"},[e("div",{staticClass:"leftContainer"},[t._t("blogLeftMainShow")],2),e("div",{staticClass:"rightContainer"},[t._t("blogRightMainShow")],2)])]),e("el-footer",{staticStyle:{height:"5vh"}},[t._t("blogFooterShow",(function(){return[e("div",{staticClass:"blogFooter"},[t._v(" Copyright © SuperYong Blog 2023 ")])]}))],2)],1)],1)},n=[],a=function(){var t=this,e=t._self._c;return e("div",{staticClass:"header-container"},[e("div",{staticClass:"header-content"},[e("div",{staticClass:"header-text"},[t._t("headerImgInnerShow",(function(){return[e("h1",[t._v("Welcome to My Blog")])]}))],2)])])},l=[],s={name:"blogHeader",data(){return{}}},i=s,c=o(1001),u=(0,c.Z)(i,a,l,!1,null,"16998cf2",null),d=u.exports,h={name:"basePage",components:{BlogHeader:d},props:{headerImgUrl:{type:String,default:o(4748)}},data(){return{navTextColor:"white",backgroundColor:"transparent"}},computed:{backgroundImageStyle(){return{height:"420px",padding:"0",backgroundImage:`url(${this.headerImgUrl})`}}},mounted(){window.addEventListener("scroll",this.changeColor)},beforeDestroy(){window.removeEventListener("scroll",this.changeColor)},methods:{changeColor(){window.scrollY>400?(this.backgroundColor="rgb(255,255,255, 0.8)",this.navTextColor="black"):(this.backgroundColor="transparent",this.navTextColor="white")}}},g=h,v=(0,c.Z)(g,r,n,!1,null,"670947e8",null),C=v.exports},1193:function(t,e,o){o.r(e),o.d(e,{default:function(){return u}});var r=function(){var t=this,e=t._self._c;return e("div",{staticClass:"aboutPage"},[e("base-page",{scopedSlots:t._u([{key:"blogHeaderImgInnerShow",fn:function(){return[e("div",{staticClass:"title"},[t._v("About")]),e("div",{staticClass:"subTitle"},[t._v("Hey, this is SuperYong")])]},proxy:!0}])})],1)},n=[],a=o(811),l={name:"aboutPage",components:{BasePage:a.Z}},s=l,i=o(1001),c=(0,i.Z)(s,r,n,!1,null,"5a61e807",null),u=c.exports},4748:function(t,e,o){t.exports=o.p+"img/post-bg-desk.6c52d023.jpg"}}]);