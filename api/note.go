package api

import (
	"github.com/gin-gonic/gin"
	"kayak-backend/global"
	"net/http"
	"time"
)

type NoteResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
type NoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// GetNotes godoc
// @Schemes http
// @Description 获取当前登录用户的所有笔记
// @Success 200 {object} []NoteResponse "笔记列表"
// @Failure default {string} string "服务器错误"
// @Router /note [get]
// @Security ApiKeyAuth
func GetNotes(c *gin.Context) {
	var notes []NoteResponse
	sqlString := `SELECT id, title, content FROM note WHERE user_id = $1`
	if err := global.Database.Select(&notes, sqlString, c.GetInt("UserId")); err != nil {
		c.String(http.StatusInternalServerError, "服务器错误")
		return
	}
	c.JSON(http.StatusOK, notes)
}

// CreateNote godoc
// @Schemes http
// @Description 创建笔记
// @Param note body NoteRequest true "笔记信息"
// @Param is_public query bool false "是否公开"
// @Success 200 {string} string "创建成功"
// @Failure 400 {string} string "请求解析失败"
// @Failure default {string} string "服务器错误"
// @Router /note/create [post]
// @Security ApiKeyAuth
func CreateNote(c *gin.Context) {
	var note NoteRequest
	if err := c.ShouldBindJSON(&note); err != nil {
		c.String(http.StatusBadRequest, "请求解析失败")
		return
	}
	sqlString := `INSERT INTO note (title, content, created_at, updated_at, user_id, is_public) VALUES ($1, $2, $3, $4, $5, $6)`
	if _, err := global.Database.Exec(sqlString, note.Title, note.Content, time.Now(),
		time.Now(), c.GetInt("UserId"), c.Query("is_public")); err != nil {
		c.String(http.StatusInternalServerError, "服务器错误")
		return
	}
	c.String(http.StatusOK, "创建成功")
}

// UpdateNote godoc
// @Schemes http
// @Description 更新笔当前登录用户的笔记
// @Param note body NoteResponse true "笔记信息"
// @Param is_public query bool false "是否公开"
// @Success 200 {string} string "更新成功"
// @Failure 400 {string} string "请求解析失败"
// @Failure 403 {string} string "没有权限"
// @Failure default {string} string "服务器错误"
// @Router /note/update [put]
// @Security ApiKeyAuth
func UpdateNote(c *gin.Context) {
	var note NoteResponse
	if err := c.ShouldBindJSON(&note); err != nil {
		c.String(http.StatusBadRequest, "请求解析失败")
		return
	}
	sqlString := `SELECT user_id FROM note WHERE id = $1`
	var userId int
	if err := global.Database.Get(&userId, sqlString, note.ID); err != nil {
		c.String(http.StatusInternalServerError, "服务器错误")
		return
	}
	if userId != c.GetInt("UserId") {
		c.String(http.StatusForbidden, "没有权限")
		return
	}
	sqlString = `UPDATE note SET title = $1, content = $2, updated_at = $3, is_public = $4 WHERE id = $5`
	if _, err := global.Database.Exec(sqlString, note.Title, note.Content, time.Now(),
		c.Query("is_public"), note.ID); err != nil {
		c.String(http.StatusInternalServerError, "服务器错误")
		return
	}
	c.String(http.StatusOK, "更新成功")
}

// DeleteNote godoc
// @Schemes http
// @Description 删除当前登录用户的笔记
// @Param id path int true "笔记ID"
// @Success 200 {string} string "删除成功"
// @Failure 400 {string} string "请求解析失败"
// @Failure 403 {string} string "没有权限"
// @Failure default {string} string "服务器错误"
// @Router /note/delete/{id} [delete]
// @Security ApiKeyAuth
func DeleteNote(c *gin.Context) {
	userId := c.GetInt("UserId")
	noteId := c.Param("id")
	sqlString := `SELECT user_id FROM problem_type WHERE id = $1`
	var noteUserId int
	if err := global.Database.Get(&noteUserId, sqlString, noteId); err != nil {
		c.String(http.StatusInternalServerError, "服务器错误")
		return
	}
	if userId != noteUserId {
		c.String(http.StatusForbidden, "没有权限")
		return
	}
	sqlString = `DELETE FROM note WHERE id = $1`
	if _, err := global.Database.Exec(sqlString, noteId); err != nil {
		c.String(http.StatusInternalServerError, "服务器错误")
		return
	}
	c.String(http.StatusOK, "删除成功")
}
