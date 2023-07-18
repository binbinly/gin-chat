"use strict";(self["webpackChunkvue_h5_template"]=self["webpackChunkvue_h5_template"]||[]).push([[67],{31067:function(t,e,n){n.r(e),n.d(e,{default:function(){return k}});var i=function(){var t=this,e=t._self._c;return e("div",[e("van-nav-bar",{attrs:{"left-text":"选择","left-arrow":"",fixed:"",placeholder:""},on:{"click-left":t.onClickLeft},scopedSlots:t._u([{key:"right",fn:function(){return[e("van-button",{attrs:{type:"primary",size:"small"},on:{click:t.submit}},[t._v(t._s(t.buttonText))])]},proxy:!0}])}),"see"===t.type?t._l(t.typeList,(function(n,i){return e("van-cell",{key:i,attrs:{title:n.name,center:""},on:{click:function(e){t.typeIndex=i}},scopedSlots:t._u([{key:"right-icon",fn:function(){return[e("van-checkbox",{attrs:{value:t.typeIndex===i,"checked-color":"#08c060"}})]},proxy:!0}],null,!0)})})):t._e(),"see"!==t.type||"see"===t.type&&(1===t.typeIndex||2===t.typeIndex)?[e("van-index-bar",{attrs:{"sticky-offset-top":45}},[t._l(t.list,(function(n,i){return[e("van-index-anchor",{key:i,attrs:{index:n.title}}),t._l(n.list,(function(n,i){return e("van-cell",{key:i,attrs:{title:n.name,center:""},on:{click:function(e){return t.selectItem(n)}},scopedSlots:t._u([{key:"icon",fn:function(){return[e("van-image",{staticClass:"pr-1",attrs:{round:"",width:"35",height:"35",src:t._f("formatAvatar")(n.avatar)}})]},proxy:!0},{key:"right-icon",fn:function(){return[e("van-checkbox",{ref:"checkboxes",refInFor:!0,attrs:{"checked-color":"#08c060"},model:{value:n.checked,callback:function(e){t.$set(n,"checked",e)},expression:"item2.checked"}})]},proxy:!0}],null,!0)})}))]}))],2)]:t._e()],2)},r=[],s=(n(58479),n(39146)),c=(n(57658),n(20629)),u=n(38810),o=n(82169),a=n(68159),l=n(22837),h={mixins:[u.Z],data(){return{typeIndex:0,typeList:[{name:"公开",key:"all"},{name:"谁可以看",key:"only"},{name:"不给谁看",key:"except"},{name:"私密",key:"none"}],selectList:[],type:"",limit:9,id:0}},mounted(){this.$store.dispatch("contactList")},activated(){this.type=this.$route.query.type,this.limit=parseInt(this.$route.query.limit),this.id=parseInt(this.$route.query.id),this.id&&"inviteGroup"===this.type&&(this.limit=1)},computed:{...(0,c.rn)({list:t=>t.user.friends.map((t=>(t.list=t.list.map((t=>({...t,checked:!1}))),t)))}),buttonText(){let t="发送";return"createGroup"===this.type&&(t="创建群组"),t+" ("+this.selectCount+")"},selectCount(){return this.selectList.length}},methods:{selectItem(t){if(!t.checked&&this.selectCount===this.limit)return(0,s.Z)("最多选中 "+this.limit+" 个");if(t.checked=!t.checked,t.checked)this.selectList.push(t);else{let e=this.selectList.findIndex((e=>e===t));e>-1&&this.selectList.splice(e,1)}},submit(){if("see"!==this.type&&0===this.selectCount)return(0,s.Z)("请先选择");switch(this.type){case"createGroup":(0,a.bq)({ids:this.selectList.map((t=>t.id))}).then((()=>{s.Z.success("创建群聊成功"),this.$router.back()}));break;case"sendCard":let t=this.selectList[0];o.Z.$emit("sendItem",{sendType:"card",content:t.name,type:l.Z.TYPE_CARD,options:{avatar:t.avatar,id:t.id}}),this.$router.back();break;case"remind":o.Z.$emit("sendResult",{type:"remind",data:this.selectList}),this.$router.back();break;case"see":let e=this.typeList[this.typeIndex].key;if("all"!==e&&"none"!==e&&!this.selectCount)return(0,s.Z)("请先选择");o.Z.$emit("sendResult",{type:"see",data:{k:e,v:this.selectList}}),this.$router.back();break;case"inviteGroup":(0,a.Q_)({id:this.id,user_id:this.selectList[0].id}).then((()=>{s.Z.success("邀请成功"),o.Z.$emit("refreshGroupInfo"),this.$router.back()}));break}}}},p=h,d=n(1001),f=(0,d.Z)(p,i,r,!1,null,null,null),k=f.exports},68159:function(t,e,n){n.d(e,{Aj:function(){return p},F8:function(){return o},Ge:function(){return u},Q_:function(){return s},X8:function(){return c},_u:function(){return a},bq:function(){return r},gF:function(){return l},od:function(){return h}});var i=n(70804);function r(t){return i.Z.post(i.Z.Group.Create,t)}function s(t){return i.Z.post(i.Z.Group.Invite,t)}function c(t){return i.Z.get(i.Z.Group.List,{p:t})}function u(t){return i.Z.get(i.Z.Group.Info,{id:t})}function o(t){return i.Z.post(i.Z.Group.Edit,t)}function a(t){return i.Z.post(i.Z.Group.Nickname,t)}function l(t){return i.Z.get(i.Z.Group.Quit,t)}function h(t){return i.Z.post(i.Z.Group.KickOff,t)}function p(t){return i.Z.get(i.Z.Group.User,{id:t})}},38810:function(t,e,n){n(58479);var i=n(39146),r=(n(57658),n(73344));e.Z={created(){let t=(0,r.cF)("token");if(!t)return this.$router.push({path:"/login"})},methods:{onClickLeft(){this.$router.back()},push(t,e){this.$router.push({path:t,params:e})},draw(){(0,i.Z)("待开发")},backToast(t="非法参数"){this.toast(t),setTimeout((()=>{this.$router.back()}),500)},toast(t="非法参数"){(0,i.Z)({title:t,position:"bottom"})}}}}}]);