"use strict";(self["webpackChunkvue_h5_template"]=self["webpackChunkvue_h5_template"]||[]).push([[211],{80211:function(t,o,e){e.r(o),e.d(o,{default:function(){return c}});var i=function(){var t=this,o=t._self._c;return o("div",[o("van-nav-bar",{attrs:{"left-text":"设置朋友圈动态权限","left-arrow":""},on:{"click-left":t.onClickLeft}}),o("van-cell",{attrs:{center:"",title:"不让他看我"},scopedSlots:t._u([{key:"right-icon",fn:function(){return[o("van-switch",{attrs:{size:"24","active-color":"#08c060"},on:{change:t.onChangeMe},model:{value:t.form.lookme,callback:function(o){t.$set(t.form,"lookme",o)},expression:"form.lookme"}})]},proxy:!0}])}),o("van-cell",{attrs:{center:"",title:"不看他"},scopedSlots:t._u([{key:"right-icon",fn:function(){return[o("van-switch",{attrs:{size:"24","active-color":"#08c060"},on:{change:t.onChangeHim},model:{value:t.form.lookhim,callback:function(o){t.$set(t.form,"lookhim",o)},expression:"form.lookhim"}})]},proxy:!0}])})],1)},r=[],s=(e(58479),e(39146)),n=e(38810),l=e(95014),h={mixins:[n.Z],data(){return{id:0,form:{lookme:!0,lookhim:!0}}},mounted(){this.id=parseInt(this.$route.query.id),this.form.lookhim=1==this.$route.query.lookhim,this.form.lookme=1==this.$route.query.lookme},methods:{onChangeMe(t){this.form.lookme=t,this.submit()},onChangeHim(t){this.form.lookhim=t,this.submit()},submit(){(0,l.Hj)({user_id:this.id,lookme:this.form.lookme?1:0,lookhim:this.form.lookhim?1:0}).then((t=>{s.Z.success("修改成功")}))}}},a=h,u=e(1001),m=(0,u.Z)(a,i,r,!1,null,null,null),c=m.exports},38810:function(t,o,e){e(58479);var i=e(39146),r=(e(57658),e(73344));o.Z={created(){let t=(0,r.cF)("token");if(!t)return this.$router.push({path:"/login"})},methods:{onClickLeft(){this.$router.back()},push(t,o){this.$router.push({path:t,params:o})},draw(){(0,i.Z)("待开发")},backToast(t="非法参数"){this.toast(t),setTimeout((()=>{this.$router.back()}),500)},toast(t="非法参数"){(0,i.Z)({title:t,position:"bottom"})}}}}}]);