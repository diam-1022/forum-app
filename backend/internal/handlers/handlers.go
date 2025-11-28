package handlers

import (
    "net/http"
    "strconv"

    "forum-app/backend/internal/models"
    "forum-app/backend/internal/repository"
    "github.com/gin-gonic/gin"
)

type Handler struct {
    repo *repository.Repository
}

func New(repo *repository.Repository) *Handler {
    return &Handler{repo: repo}
}

func (h *Handler) Health(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) ListBoards(c *gin.Context) {
    boards, err := h.repo.ListBoards()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, boards)
}

func (h *Handler) ListTopics(c *gin.Context) {
    topics, err := h.repo.ListTopics()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, topics)
}

func (h *Handler) ListPosts(c *gin.Context) {
    var boardIDPtr, topicIDPtr *uint
    if boardID := c.Query("boardId"); boardID != "" {
        if v, err := strconv.Atoi(boardID); err == nil {
            val := uint(v)
            boardIDPtr = &val
        }
    }
    if topicID := c.Query("topicId"); topicID != "" {
        if v, err := strconv.Atoi(topicID); err == nil {
            val := uint(v)
            topicIDPtr = &val
        }
    }
    posts, err := h.repo.ListPosts(boardIDPtr, topicIDPtr)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, posts)
}

func (h *Handler) GetPost(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
        return
    }
    post, err := h.repo.GetPostWithComments(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, post)
}

type createPostRequest struct {
    Title       string `json:"title" binding:"required"`
    Content     string `json:"content" binding:"required"`
    BoardID     uint   `json:"boardId" binding:"required"`
    TopicID     uint   `json:"topicId" binding:"required"`
    Username    string `json:"username" binding:"required"`
    DisplayName string `json:"displayName" binding:"required"`
}

func (h *Handler) CreatePost(c *gin.Context) {
    var req createPostRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    user, err := h.repo.EnsureUser(req.Username, req.DisplayName)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    post := models.Post{
        Title:    req.Title,
        Content:  req.Content,
        BoardID:  req.BoardID,
        TopicID:  req.TopicID,
        AuthorID: user.ID,
    }
    if err := h.repo.CreatePost(&post); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, post)
}

type createCommentRequest struct {
    Content     string `json:"content" binding:"required"`
    Username    string `json:"username" binding:"required"`
    DisplayName string `json:"displayName" binding:"required"`
}

func (h *Handler) CreateComment(c *gin.Context) {
    postID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
        return
    }
    var req createCommentRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    user, err := h.repo.EnsureUser(req.Username, req.DisplayName)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    comment := models.Comment{
        PostID:   uint(postID),
        Content:  req.Content,
        AuthorID: user.ID,
    }
    if err := h.repo.CreateComment(&comment); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, comment)
}


