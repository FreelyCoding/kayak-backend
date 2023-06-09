package api

import "kayak-backend/global"

func InitRoute() {
	global.Router.GET("/ping", Ping)
	global.Router.GET("/logout", Logout)
	global.Router.POST("/login", Login)
	global.Router.POST("/register", Register)
	global.Router.POST("/change-password", ChangePassword)
	global.Router.POST("/reset-password", ResetPassword)
	global.Router.POST("/send-email", SendEmail)
	global.Router.POST("/weixin-login", WeixinLogin)
	global.Router.POST("/weixin-bind", WeixinBind)
	global.Router.POST("/weixin-complete", WeixinComplete)

	special := global.Router.Group("/special")
	special.Use(global.CheckAuth)
	special.GET("/wrong_problem_set", GetWrongProblemSet)
	special.GET("/favorite_problem_set", GetFavoriteProblemSet)
	special.GET("/featured_problem_set", GetFeaturedProblemSet)
	special.GET("/featured_note", GetFeaturedNote)
	special.GET("/featured_group", GetFeaturedGroup)
	special.POST("/picture_ocr", PictureOCR)
	special.POST("/pdf_ocr", PDFFileOCR)

	user := global.Router.Group("/user")
	user.Use(global.CheckAuth)
	user.GET("/info", GetUserInfo)
	user.GET("/info/:user_id", GetUserInfoById)
	user.PUT("/update", UpdateUserInfo)
	user.GET("/wrong_record", GetUserWrongRecords)

	upload := global.Router.Group("/upload")
	upload.Use(global.CheckAuth)
	upload.POST("/public", UploadPublicFile)
	upload.POST("/avatar", UploadUserAvatar)
	upload.POST("/group_avatar", UploadGroupAvatar)

	note := global.Router.Group("/note")
	note.Use(global.CheckAuth)
	global.Router.GET("/note/all", GetNotes)
	note.POST("/create", CreateNote)
	note.PUT("/update", UpdateNote)
	note.DELETE("/delete/:id", DeleteNote)
	note.POST("/like/:id", LikeNote)
	note.POST("/unlike/:id", UnlikeNote)
	note.POST("/favorite/:id", FavoriteNote)
	note.DELETE("/unfavorite/:id", UnfavoriteNote)
	note.POST("/add_problem/:id", AddProblemToNote)
	note.DELETE("/remove_problem/:id", RemoveProblemFromNote)
	note.GET("/problem_list/:id", GetNoteProblemList)

	wrongRecord := global.Router.Group("/wrong_record")
	wrongRecord.Use(global.CheckAuth)
	wrongRecord.POST("/create/:id", CreateWrongRecord)
	wrongRecord.DELETE("/delete/:id", DeleteWrongRecord)
	wrongRecord.GET("/get/:id", GetWrongRecord)

	problem := global.Router.Group("/problem")
	problem.Use(global.CheckAuth)
	problem.DELETE("/unfavorite/:id", RemoveProblemFromFavorite)
	problem.POST("/favorite/:id", AddProblemToFavorite)
	problem.POST("/batch", AddBatchProblem)

	choiceProblem := problem.Group("/choice")
	global.Router.GET("/problem/choice/all", GetChoiceProblems)
	choiceProblem.POST("/create", CreateChoiceProblem)
	choiceProblem.PUT("/update", UpdateChoiceProblem)
	choiceProblem.DELETE("/delete/:id", DeleteChoiceProblem)
	choiceProblem.GET("/answer/:id", GetChoiceProblemAnswer)

	blankProblem := problem.Group("/blank")
	global.Router.GET("/problem/blank/all", GetBlankProblems)
	blankProblem.POST("/create", CreateBlankProblem)
	blankProblem.PUT("/update", UpdateBlankProblem)
	blankProblem.DELETE("/delete/:id", DeleteBlankProblem)
	blankProblem.GET("/answer/:id", GetBlankProblemAnswer)

	judgeProblem := problem.Group("/judge")
	global.Router.GET("/problem/judge/all", GetJudgeProblems)
	judgeProblem.POST("/create", CreateJudgeProblem)
	judgeProblem.PUT("/update", UpdateJudgeProblem)
	judgeProblem.DELETE("/delete/:id", DeleteJudgeProblem)
	judgeProblem.GET("/answer/:id", GetJudgeProblemAnswer)

	problemSet := global.Router.Group("/problem_set")
	problemSet.Use(global.CheckAuth)
	global.Router.GET("/problem_set/all", GetProblemSets)
	problemSet.POST("/create", CreateProblemSet)
	problemSet.PUT("/update", UpdateProblemSet)
	problemSet.DELETE("/delete/:id", DeleteProblemSet)
	problemSet.GET("/all_problem/:id", GetProblemsInProblemSet)
	problemSet.POST("/add/:id", AddProblemToProblemSet)
	problemSet.POST("/migrate/:id", MigrateProblemToProblemSet)
	problemSet.DELETE("/remove/:id", RemoveProblemFromProblemSet)
	problemSet.POST("/favorite/:id", AddProblemSetToFavorite)
	problemSet.DELETE("/unfavorite/:id", RemoveProblemSetFromFavorite)
	problemSet.GET("/statistic/wrong_count", GetWrongCountOfProblemSet)
	problemSet.GET("/statistic/fav_count", GetFavoriteCountOfProblemSet)

	noteReview := global.Router.Group("/note_review")
	noteReview.Use(global.CheckAuth)
	noteReview.POST("/add", AddNoteReview)
	noteReview.DELETE("/remove/:id", RemoveNoteReview)
	noteReview.GET("/get", GetNoteReviews)
	noteReview.POST("/like/:id", LikeNoteReview)
	noteReview.POST("/unlike/:id", UnlikeNoteReview)

	group := global.Router.Group("/group")
	group.Use(global.CheckAuth)
	group.GET("/all", GetGroups)
	group.POST("/create", CreateGroup)
	group.GET("/invitation/:id", GetGroupInvitation)
	group.DELETE("/delete/:id", DeleteGroup)
	group.GET("/all_user/:id", GetUsersInGroup)
	group.POST("/add/:id", AddUserToGroup)
	group.DELETE("/remove/:id", RemoveUserFromGroup)
	group.DELETE("/quit/:id", QuitGroup)
	group.PUT("/update/:id", UpdateGroupInfo)
	group.POST("/apply", ApplyToJoinGroup)
	group.GET("/application/:id", GetGroupApplication)
	group.PUT("/application", HandleGroupApplication)

	discussion := global.Router.Group("/discussion")
	discussion.Use(global.CheckAuth)
	discussion.GET("/all", GetDiscussions)
	discussion.POST("/create", CreateDiscussion)
	discussion.PUT("/update", UpdateDiscussion)
	discussion.DELETE("/delete/:id", DeleteDiscussion)
	discussion.POST("/like/:id", LikeDiscussion)
	discussion.POST("/unlike/:id", UnlikeDiscussion)
	discussion.POST("/favorite/:id", FavoriteDiscussion)
	discussion.POST("/unfavorite/:id", UnfavoriteDiscussion)

	discussionReview := global.Router.Group("/discussion_review")
	discussionReview.Use(global.CheckAuth)
	discussionReview.POST("/add", AddDiscussionReview)
	discussionReview.DELETE("/remove/:id", RemoveDiscussionReview)
	discussionReview.GET("/get", GetDiscussionReviews)
	discussionReview.POST("/like/:id", LikeDiscussionReview)
	discussionReview.POST("/unlike/:id", UnlikeDiscussionReview)

	search := global.Router.Group("/search")
	search.Use(global.CheckAuth)
	search.POST("/problem_set", SearchProblemSets)
	search.POST("/group", SearchGroups)
	search.POST("/note", SearchNotes)

	check := global.Router.Group("/check")
	check.Use(global.CheckAuth)
	check.GET("/problem_set/:id", CheckProblemSetWriteAuth)

	// Deprecated
	global.Router.GET("/problem/blank/:id", GetBlankProblem)
	global.Router.GET("/problem/choice/:id", GetChoiceProblem)
	problem.GET("/:id/problem_set", GetProblemSetContainsProblem)
	user.GET("/favorite/problem", GetUserFavoriteProblems)
	user.GET("/favorite/problem_set", GetUserFavoriteProblemSets)
	user.GET("/favorite/note", GetUserFavoriteNotes)
	user.GET("/problem/choice", GetUserChoiceProblems)
	user.GET("/problem/blank", GetUserBlankProblems)
	user.GET("/problem_set", GetUserProblemSets)
	user.GET("/note", GetUserNotes)
}
