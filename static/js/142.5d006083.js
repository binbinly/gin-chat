"use strict";(self["webpackChunkvue_h5_template"]=self["webpackChunkvue_h5_template"]||[]).push([[142],{57142:function(t,n,o){o.r(n),o.d(n,{default:function(){return m}});var e=function(){var t=this,n=t._self._c;return n("div",[n("van-nav-bar",{attrs:{title:"表情包",fixed:"",placeholder:"","left-arrow":""},on:{"click-left":t.onClickLeft}}),t._l(t.list,(function(o,e){return n("van-cell",{key:e,attrs:{title:o.category,center:"","is-link":""},on:{click:function(n){return t.handle(o)}},scopedSlots:t._u([{key:"icon",fn:function(){return[n("van-image",{staticClass:"pr-1",attrs:{"lazy-load":"",width:"60",height:"60",fit:"cover",src:o.url}})]},proxy:!0},{key:"right-icon",fn:function(){return[o.isAdd?n("span",[t._v("已添加")]):n("van-button",{attrs:{type:"primary",size:"small"},on:{click:function(n){return n.stopPropagation(),t.submit(o)}}},[t._v("添加")])]},proxy:!0}],null,!0)})}))],2)},r=[],i=o(71568),u=o(20629),c=o(2718),a=o(18128),s={mixins:[i.Z],data(){return{list:[]}},computed:{...(0,u.rn)({emoCat:t=>t.user.emoCat})},mounted(){this.initData(),a.Z.$on("onEmoticon",this.onEmoticon)},destroy(){a.Z.$off("onEmoticon",this.onEmoticon)},methods:{onEmoticon(t){this.list.forEach((n=>{n.category==t&&(n.isAdd=!0)}))},initData(){(0,c.Zr)().then((t=>{this.list=t.map((t=>{let n=-1;return this.emoCat.length>0&&(n=this.emoCat.indexOf(t.category)),{category:t.category,isAdd:n>-1,url:t.url}}))}))},handle(t){this.$router.push({path:"/emoticon",query:{cat:t.category}})},submit(t){this.$store.commit("addEmo",t.category),t.isAdd=!0,a.Z.$emit("onEmoticon",t.category)}}},l=s,h=o(1001),f=(0,h.Z)(l,e,r,!1,null,null,null),m=f.exports},2718:function(t,n,o){o.d(n,{S0:function(){return a},YD:function(){return i},Zr:function(){return c},cT:function(){return s},qz:function(){return u},x4:function(){return r}});var e=o(7705);function r(t){return e.Z.post(e.Z.Login,t,!1)}function i(t){return e.Z.post(e.Z.Reg,t,!1)}function u(t){return e.Z.post(e.Z.SearchUser,t)}function c(){return e.Z.get(e.Z.EmoticonCat)}function a(t){return e.Z.get(e.Z.Emoticon,{cat:t},!0)}function s(t){return e.Z.upload(e.Z.Upload,{file:t}).then((t=>t)).catch((t=>{reject(t)}))}},71568:function(t,n,o){o(58479);var e=o(39146),r=o(76927);n["Z"]={created(){let t=(0,r.cF)("token");if(!t)return this.$router.push({path:"/login"})},methods:{onClickLeft(){this.$router.back()},push(t,n){this.$router.push({path:t,params:n})},draw(){(0,e.Z)("待开发")},backToast(t="非法参数"){this.toast(t),setTimeout((()=>{this.$router.back()}),500)},toast(t="非法参数"){(0,e.Z)({title:t,position:"bottom"})}}}}}]);