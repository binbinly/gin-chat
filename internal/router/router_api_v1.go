package router

import (
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api/v1/apply"
	"gin-chat/internal/api/v1/chat"
	"gin-chat/internal/api/v1/collect"
	"gin-chat/internal/api/v1/emoticon"
	"gin-chat/internal/api/v1/friend"
	"gin-chat/internal/api/v1/group"
	"gin-chat/internal/api/v1/moment"
	"gin-chat/internal/api/v1/upload"
	"gin-chat/internal/api/v1/user"
	mw "gin-chat/internal/middleware"
)

func setApiV1(v1 *gin.RouterGroup) {
	// 认证相关路由
	v1.POST("/reg", user.Register)
	v1.POST("/login", user.Login)
	v1.POST("/login_phone", user.PhoneLogin)
	v1.POST("/send_code", user.SendCode)

	up := v1.Group("/upload")
	up.Use(mw.JWT())
	{
		up.POST("/image", upload.Image)
	}
	// 用户模块
	u := v1.Group("/user")
	u.Use(mw.JWT())
	{
		u.POST("/edit", user.Update)
		u.GET("/profile", user.Profile)
		u.GET("/tag", user.Tag)
		u.GET("/logout", user.Logout)
		u.POST("/search", user.Search)
		u.POST("/report", user.Report)
	}

	// 好友申请模块
	a := v1.Group("/apply")
	a.Use(mw.JWT())
	{
		a.POST("/friend", apply.Friend)
		a.POST("/handle", apply.Handle)
		a.GET("/list", apply.List)
		a.GET("/count", apply.Count)
	}

	// 好友模块
	f := v1.Group("/friend")
	f.Use(mw.JWT())
	{
		f.GET("/info", friend.Info)
		f.GET("/list", friend.List)
		f.POST("/black", friend.Black)
		f.POST("/star", friend.Star)
		f.POST("/auth", friend.Auth)
		f.POST("/remark", friend.Remark)
		f.POST("/destroy", friend.Destroy)
		f.GET("/tag_list", friend.TagList)
	}

	// 聊天模块
	c := v1.Group("/chat")
	c.Use(mw.JWT())
	{
		c.POST("/detail", chat.Detail)
		c.POST("/send", chat.Send)
		c.POST("/recall", chat.Recall)
	}

	// 群组模块
	gr := v1.Group("/group")
	gr.Use(mw.JWT())
	{
		gr.POST("/create", group.Create)
		gr.POST("/edit", group.Update)
		gr.POST("/nickname", group.UpdateNickname)
		gr.GET("/list", group.List)
		gr.GET("/info", group.Info)
		gr.GET("/user", group.User)
		gr.GET("/quit", group.Quit)
		gr.GET("/join", group.Join)
		gr.POST("/kickoff", group.KickOff)
		gr.POST("/invite", group.Invite)
	}

	coll := v1.Group("/collect")
	coll.Use(mw.JWT())
	{
		coll.POST("/create", collect.Create)
		coll.GET("/list", collect.List)
		coll.POST("/destroy", collect.Destroy)
	}

	mom := v1.Group("/moment")
	mom.Use(mw.JWT())
	{
		mom.POST("/create", moment.Create)
		mom.GET("/list", moment.List)
		mom.GET("/timeline", moment.Timeline)
		mom.POST("/like", moment.Like)
		mom.POST("/comment", moment.Comment)
	}

	emo := v1.Group("/emoticon")
	emo.Use(mw.JWT())
	{
		emo.GET("/list", emoticon.List)
		emo.GET("/cat", emoticon.Cat)
	}
}
