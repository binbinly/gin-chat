"use strict";(self["webpackChunkvue_h5_template"]=self["webpackChunkvue_h5_template"]||[]).push([[856],{70856:function(n,t,e){e.r(t),e.d(t,{default:function(){return l}});var r=function(){var n=this,t=n._self._c;return t("div",[t("form",[t("van-search",{attrs:{"show-action":"",fixed:"",placeholder:"搜索用户"},on:{search:n.onSearch,cancel:n.onCancel},model:{value:n.keyword,callback:function(t){n.keyword=t},expression:"keyword"}})],1),n._l(n.list,(function(e,r){return t("van-cell",{attrs:{center:"",value:e.nickname,"is-link":""},on:{click:function(t){return n.openUserBase(e.id)}},scopedSlots:n._u([{key:"icon",fn:function(){return[t("van-image",{staticClass:"pr-1",attrs:{round:"",width:"35",height:"35",src:n._f("formatAvatar")(e.avatar)}})]},proxy:!0},{key:"title",fn:function(){return[t("span",{staticClass:"custom-title pr-1"},[n._v(n._s(e.username))]),t("van-tag",{attrs:{type:"danger"}},[n._v(n._s(e.phone))])]},proxy:!0}],null,!0)})}))],2)},o=[],u=e(2718),a={data(){return{keyword:"",list:[]}},methods:{onSearch(n){(0,u.qz)({keyword:this.keyword}).then((n=>{this.list=n}))},onCancel(){this.$router.back()},openUserBase(n){this.$router.push({path:"/user_base",query:{id:n}})}}},c=a,i=e(1001),s=(0,i.Z)(c,r,o,!1,null,null,null),l=s.exports},2718:function(n,t,e){e.d(t,{S0:function(){return i},YD:function(){return u},Zr:function(){return c},cT:function(){return s},qz:function(){return a},x4:function(){return o}});var r=e(7705);function o(n){return r.Z.post(r.Z.Login,n,!1)}function u(n){return r.Z.post(r.Z.Reg,n,!1)}function a(n){return r.Z.post(r.Z.SearchUser,n)}function c(){return r.Z.get(r.Z.EmoticonCat)}function i(n){return r.Z.get(r.Z.Emoticon,{cat:n},!0)}function s(n){return new Promise(((t,e)=>{r.Z.post(r.Z.Upload,{file:n},!0,!0).then((n=>n)).catch((n=>{e(n)}))}))}}}]);