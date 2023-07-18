"use strict";(self["webpackChunkvue_h5_template"]=self["webpackChunkvue_h5_template"]||[]).push([[836],{32894:function(t,e,s){s.d(e,{Z:function(){return c}});var i=function(){var t=this,e=t._self._c;return e("div",{staticClass:"flex align-center justify-center",staticStyle:{height:"45px",width:"45px"},attrs:{"hover-class":"bg-hover-light"},on:{click:function(e){return t.$emit("click")}}},[e("span",{staticClass:"iconfont font-md"},[t._v(t._s(t.icon))])])},n=[],o={props:{icon:{type:String,default:""}}},a=o,r=s(1001),l=(0,r.Z)(a,i,n,!1,null,null,null),c=l.exports},75836:function(t,e,s){s.r(e),s.d(e,{default:function(){return I}});var i=function(){var t=this,e=t._self._c;return e("div",{attrs:{id:"content"}},[e("div",{staticClass:"body",on:{click:function(e){t.show=!1}}},[e("transparent-bar",{attrs:{scrollTop:t.scrollTop,title:t.title},on:{clickRight:t.clickRight}}),e("div",{staticClass:"position-relative",staticStyle:{height:"620rpx"}},[e("van-image",{staticClass:"bg-secondary w-100",attrs:{src:s(48496),fit:"cover",height:"280"}}),e("van-image",{staticClass:"bg-secondary rounded position-absolute",staticStyle:{right:"15px",bottom:"-20px"},attrs:{src:t._f("formatAvatar")(t.userinfo.avatar),width:"60",height:"60",fit:"cover"}}),e("span",{staticClass:"text-white font-sm position-absolute",staticStyle:{bottom:"10px",right:"80px"}},[t._v(t._s(t.userinfo.name))])],1),t.showNotice?e("div",{staticClass:"w-100 text-center mt-1"},[e("van-tag",{attrs:{round:"",type:"primary",size:"medium"},on:{click:t.onRefresh}},[t._v("有新动态哦")])],1):t._e(),e("van-pull-refresh",{on:{refresh:t.onRefresh},model:{value:t.refreshing,callback:function(e){t.refreshing=e},expression:"refreshing"}},[e("van-list",{attrs:{finished:t.finished,"finished-text":"没有更多了"},on:{load:t.onLoad},model:{value:t.loading,callback:function(e){t.loading=e},expression:"loading"}},t._l(t.list,(function(s,i){return e("moment-list",{key:i,attrs:{item:s,index:i},on:{action:t.doAction,reply:t.replyEvent,openVideo:t.openVideo}})})),1)],1)],1),e("van-popup",{attrs:{position:"bottom",overlay:!1,"lock-scroll":!1},model:{value:t.show,callback:function(e){t.show=e},expression:"show"}},[e("div",{staticClass:"bg-light border-top flex align-center px-1",staticStyle:{height:"50px"}},[e("van-field",{staticStyle:{height:"43px",width:"75%"},attrs:{rows:"1",type:"textarea",placeholder:t.placeholder||"文明发言"},model:{value:t.content,callback:function(e){t.content=e},expression:"content"}}),e("icon-button",{attrs:{icon:""},on:{click:t.changeFaceModal}}),e("van-button",{attrs:{type:"primary",size:"small",disabled:0===t.content.length},on:{click:t.send}},[t._v("发送")])],1),t.faceModal?e("div",{staticClass:"flex flex-wrap",staticStyle:{height:"200px",overflow:"auto"}},t._l(t.faceList,(function(s,i){return e("div",{key:i,staticClass:"flex align-center justify-center",staticStyle:{width:"45px",height:"45px"},attrs:{"hover-class":"bg-white"},on:{click:function(e){return t.addFace(s)}}},[e("span",{staticStyle:{"font-size":"24px"}},[t._v(t._s(s))])])})),0):t._e()]),e("van-overlay",{attrs:{show:t.showVideo},on:{click:function(e){t.showVideo=!1}}},[e("div",{staticClass:"wrapper"},[e("video",{staticClass:"w-100",attrs:{src:t.videoUrl,controls:""}})])]),e("van-action-sheet",{attrs:{actions:t.actions,"cancel-text":"取消","close-on-click-action":""},on:{cancel:t.onCancel,select:t.onSelect},model:{value:t.showAction,callback:function(e){t.showAction=e},expression:"showAction"}})],1)},n=[],o=(s(58479),s(39146)),a=(s(57658),s(32894)),r=function(){var t=this,e=t._self._c;return e("div",[e("div",{staticClass:"fixed-top",style:t.navBarStyle},[e("div",{style:"height:"+t.statusBarHeight+"px"}),e("div",{staticClass:"w-100 flex align-center justify-between",staticStyle:{height:"45px"}},[e("div",{staticClass:"flex align-center"},[e("div",{staticClass:"flex align-center justify-center",staticStyle:{height:"40px",width:"40px"},attrs:{"hover-class":"bg-hover-light"},on:{click:t.back}},[e("van-icon",{attrs:{name:"arrow-left",size:"22",color:t.buttonColor}})],1),t.title?e("span",{staticClass:"font"},[t._v(t._s(t.title))]):t._e()]),e("div",{staticClass:"flex align-center"},[e("div",{staticClass:"flex align-center justify-center",staticStyle:{height:"40px",width:"40px"},attrs:{"hover-class":"bg-hover-light"},on:{click:function(e){return t.$emit("clickRight")}}},[e("van-icon",{attrs:{name:"photograph",size:"22",color:t.buttonColor}})],1)])])])])},l=[],c={props:{title:{type:[String,Boolean],default:!1},scrollTop:{type:[Number,String],default:0}},data(){return{statusBarHeight:0,navBarHeight:0}},mounted(){this.navBarHeight=this.statusBarHeight+90},computed:{changeNumber(){let t=200,e=280,s=e-t,i=0;return this.scrollTop>t&&(i=(this.scrollTop-t)/s),i>1?1:i},navBarStyle(){return`background-color: rgba(255,255,255,${this.changeNumber});`},buttonColor(){return this.changeNumber>0?"#000000":"#FFFFFF"}},methods:{back(){this.$router.back()}}},h=c,d=s(1001),p=(0,d.Z)(h,r,l,!1,null,null,null),u=p.exports,m=function(){var t=this,e=t._self._c;return e("div",{staticClass:"px-1 pt-1 flex align-start border-bottom border-light-secondary"},[e("free-avater",{attrs:{src:t.item.user.avatar,uid:t.item.user.id}}),e("div",{staticClass:"pl-1 flex-1 flex flex-column"},[e("span",{staticClass:"text-hover-primary font-sm",staticStyle:{"margin-bottom":"5px"}},[t._v(t._s(t.item.user.name))]),t.item.content?e("span",{staticClass:"text-dark font-sm"},[t._v(t._s(t.item.content))]):t._e(),t.item.image?e("div",{staticClass:"pt-1 flex flex-wrap"},[t._l(t.imgs,(function(s,i){return[1===t.imgs.length?e("van-image",{key:i,staticStyle:{"max-width":"180px","max-height":"240px"},attrs:{src:s,fit:"cover",imageClass:"rounded bg-secondary"},on:{click:function(e){return t.prediv(s)}}}):e("van-image",{key:i,staticClass:"bg-secondary rounded",staticStyle:{margin:"0 5px 5px 0"},attrs:{src:s,fit:"cover",width:"90",height:"90"},on:{click:function(e){return t.prediv(i)}}})]}))],2):t._e(),t.item.video?e("div",{staticClass:"position-relative rounded",on:{click:t.openVideo}},[e("video",{staticStyle:{"max-width":"200px","max-height":"300px"},attrs:{src:t.item.video}}),e("span",{staticClass:"iconfont text-white position-absolute",staticStyle:{"font-size":"35px",width:"35px",height:"35px"},style:t.posterIconStyle},[t._v("")])]):t._e(),e("div",{staticClass:"flex align-center justify-between"},[e("span",{staticClass:"text-light-muted font-sm"},[t._v(t._s(t._f("formatTime")(t.item.created_at)))]),e("div",{staticClass:"px-1"},[e("van-popover",{attrs:{theme:"dark",trigger:"click",placement:"left",actions:t.actions},on:{select:t.onSelect},scopedSlots:t._u([{key:"reference",fn:function(){return[e("van-icon",{attrs:{name:"ellipsis",size:"20"}})]},proxy:!0}]),model:{value:t.showPopover,callback:function(e){t.showPopover=e},expression:"showPopover"}})],1)]),t.item.likes||t.item.comments?e("div",{staticClass:"bg-light mt-1"},[t.item.likes.length?e("div",{staticClass:"border-bottom flex align-start",staticStyle:{padding:"5px"}},[e("van-icon",{attrs:{name:"like-o",size:"16",color:"#0056b3"}}),e("div",{staticClass:"flex flex-1 flex-wrap"},t._l(t.item.likes,(function(s,i){return e("span",{key:i,staticClass:"text-hover-primary ml-1"},[t._v(t._s(s.name))])})),0)],1):t._e(),t.item.comments.length?e("div",{staticClass:"flex align-start",staticStyle:{padding:"5px"}},[e("van-icon",{attrs:{name:"comment-o",size:"16",color:"#0056b3"}}),e("div",{staticClass:"flex flex-column flex-1 ml-1"},t._l(t.item.comments,(function(s,i){return e("div",{key:i,staticClass:"flex"},[s.reply?e("div",{staticClass:"flex align-center"},[e("span",{staticClass:"text-hover-primary font-sm"},[t._v(t._s(s.user.name)+" ")]),e("span",{staticClass:"text-muted font-sm",staticStyle:{margin:"0 2px"}},[t._v("回复")]),e("span",{staticClass:"text-hover-primary font-sm"},[t._v(t._s(s.reply.name)+"：")])]):e("span",{staticClass:"text-hover-primary"},[t._v(t._s(s.user.name)+"：")]),e("span",{staticClass:"text-dark flex-1 font-sm",on:{click:function(e){return e.stopPropagation(),t.$emit("reply",{item:t.item,index:t.index,reply:s.user})}}},[t._v(t._s(s.content))])])})),0)],1):t._e()]):t._e()])],1)},f=[],v=(s(85684),s(10993)),g=function(){var t=this,e=t._self._c;return e("van-image",{class:t.classStyle,attrs:{round:"",width:t.size,height:t.size,src:t._f("formatAvatar")(t.src)},on:{click:t.openUser}})},x=[],y={props:{uid:Number,size:{type:[String,Number],default:40},src:{type:String,default:""},classStyle:{type:String,default:""}},methods:{openUser(){this.$router.push({path:"/user_base",query:{id:this.uid}})}}},_=y,w=(0,d.Z)(_,g,x,!1,null,null,null),k=w.exports,b={components:{freeAvater:k},props:{item:Object,index:Number},data(){return{poster:{w:100,h:100},showPopover:!1,actions:[{text:"赞",event:"like"},{text:"评论",event:"comment"}]}},computed:{imgs(){return this.item.image?this.item.image.split(","):[]},posterIconStyle(){let t=this.poster.w/2-17.5,e=this.poster.h/2-22.5;return`left:${t}px;top:${e}px;`}},mounted(){if(this.item.video){let t=document.querySelector("video");t.addEventListener("canplay",(()=>{this.loadPoster(t.videoWidth,t.videoHeight)}))}},methods:{onSelect({event:t}){this.$emit("action",{event:t,item:this.item,index:this.index})},prediv(t){t<=0&&(t=0),(0,v.Z)({images:this.imgs,startPosition:t,closeable:!0})},openVideo(){this.$emit("openVideo",this.item.video)},loadPoster(t,e){const s=t/e;if(s>1){if(t>200)return this.poster.w=200,void(this.poster.h=parseInt(200/t*e))}else if(e>300)return this.poster.h=300,void(this.poster.h=parseInt(300/e*t));this.poster.w=t,this.poster.h=e}}},C=b,S=(0,d.Z)(C,m,f,!1,null,null,null),Z=S.exports,M=s(20629),N=s(82169),$=s(22837),B=s(89130),T={components:{IconButton:a.Z,TransparentBar:u,MomentList:Z},data(){return{showVideo:!1,videoUrl:"",showNotice:!1,showAction:!1,actions:[{name:"图文",type:"image"},{name:"短视频",type:"video"},{name:"文字",type:"text"}],show:!1,placeholder:"",content:"",scrollTop:0,list:[],faceModal:!1,faceList:["😀","😁","😂","😃","😄","😅","😆","😉","😊","😋","😎","😍","😘","😗","😙","😚","😇","😐","😑","😶","😏","😣","😥","😮","😯","😪","😫","😴","😌","😛","😜","😝","😒","😓","😔","😕","😲","😷","😖","😞","😟","😤","😢","😭","😦","😧","😨","😬","😰","😱","😳","😵","😡","😠"],commentIndex:-1,loading:!1,finished:!1,refreshing:!1,page:1,reply_user:!1,user_id:0,userinfo:{id:0,name:"",avatar:""},title:""}},computed:{...(0,M.rn)({user:t=>t.user.user,chat:t=>t.user.chat}),nickname(){return this.params?this.params.name:this.user.nickname||this.user.username},avatar(){const t=this.params?this.params.avatar:this.user.avatar;return t||s(81051)}},activated(){const t=parseInt(this.$route.query.id)||0;t&&this.user_id!=t&&(this.user_id=t,this.refresh())},mounted(){this.chat.readMoments(),N.Z.$on("momentNotice",this.momentNotice),N.Z.$on("refreshMoment",this.onRefresh),window.addEventListener("scroll",this.scroll,!0)},destroyed(){N.Z.$off("momentNotice",this.momentNotice),N.Z.$off("refreshMoment",this.onRefresh),window.removeEventListener("scroll",this.scroll)},methods:{momentNotice(t){t.user_id&&t.num&&(this.showNotice=!0)},scroll(t){this.scrollTop=document.documentElement.scrollTop,this.scrollTop>240?this.title="朋友圈":this.title=""},onLoad(){this.showNotice=!1,this.getData()},onRefresh(){this.chat.readMoments(),this.refresh()},refresh(){this.page=1,this.finished=!1,this.loading=!0,this.onLoad()},getData(){(0,B.BY)(this.user_id,this.page).then((t=>{this.refreshing&&(this.list=[],this.refreshing=!1),1==this.page&&(this.userinfo=t["user"]),this.list=1===this.page?t["list"]:[...this.list,...t["list"]],this.loading=!1,t["list"].length<$.Z.PAGE_SIZE?this.finished=!0:this.page++}))},doAction({event:t,item:e,index:s}){if("like"===t)return this.doSupport(e);"comment"===t&&(this.show=!0,this.commentIndex=s,this.reply_user=!1)},openVideo(t){this.showVideo=!0,this.videoUrl=t},initComment(){this.content="",this.faceModal=!1,this.reply_user=!1},doSupport(t){(0,B.VU)({id:t.id}).then((()=>{let e=t.likes.findIndex((t=>t.id===this.user.id));-1!==e?t.likes.splice(e,1):t.likes.push({id:this.user.id,name:this.user.nickname||this.user.username}),o.Z.success(-1!==e?"取消点赞成功":"点赞成功")}))},addFace(t){this.content+=t},changeFaceModal(){setTimeout((()=>{this.faceModal=!this.faceModal}),100)},send(){let t=this.list[this.commentIndex];(0,B.B8)({id:t.id,content:this.content,reply_id:this.reply_user?this.reply_user.id:0}).then((()=>{t.comments.push({content:this.content,user:{id:this.user.id,name:this.user.nickname||this.user.username},reply:this.reply?this.reply:null}),o.Z.success("评论成功"),this.initComment()})),this.show=!1},clickRight(){this.showAction=!0},replyEvent(t){this.content="",this.faceModal=!1,this.commentIndex=t.index,this.reply_user=t.reply,this.show=!0,this.placeholder="回复"+t.reply.name+":"},onCancel(){},onSelect({type:t}){this.$router.push({path:"/add_moment",query:{type:t}})}}},z=T,A=(0,d.Z)(z,i,n,!1,null,null,null),I=A.exports},89130:function(t,e,s){s.d(e,{AU:function(){return n},B8:function(){return a},BY:function(){return r},VU:function(){return o}});var i=s(70804);function n(t){return i.Z.post(i.Z.Moment.Create,t)}function o(t){return i.Z.post(i.Z.Moment.Like,t)}function a(t){return i.Z.post(i.Z.Moment.Comment,t)}function r(t,e){return t>0?i.Z.get(i.Z.Moment.List,{user_id:t,p:e}):i.Z.get(i.Z.Moment.Timeline,{p:e})}},48496:function(t,e,s){t.exports=s.p+"static/img/bg.4a5236ea.jpg"}}]);