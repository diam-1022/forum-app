package repository

import (
    "forum-app/backend/internal/models"
    "gorm.io/gorm"
)

type Repository struct {
    db *gorm.DB
}

func New(db *gorm.DB) *Repository {
    return &Repository{db: db}
}

func (r *Repository) ListBoards() ([]models.Board, error) {
    var boards []models.Board
    err := r.db.Find(&boards).Error
    return boards, err
}

func (r *Repository) ListTopics() ([]models.Topic, error) {
    var topics []models.Topic
    err := r.db.Find(&topics).Error
    return topics, err
}

func (r *Repository) ListPosts(boardID, topicID *uint) ([]models.Post, error) {
    var posts []models.Post
    query := r.db.Preload("Author").Preload("Board").Preload("Topic").Order("updated_at desc")
    if boardID != nil {
        query = query.Where("board_id = ?", *boardID)
    }
    if topicID != nil {
        query = query.Where("topic_id = ?", *topicID)
    }
    err := query.Find(&posts).Error
    return posts, err
}

func (r *Repository) GetPostWithComments(id uint) (models.Post, error) {
    var post models.Post
    err := r.db.Preload("Author").Preload("Board").Preload("Topic").Preload("Comments.Author").First(&post, id).Error
    return post, err
}

func (r *Repository) CreatePost(post *models.Post) error {
    return r.db.Create(post).Error
}

func (r *Repository) CreateComment(comment *models.Comment) error {
    return r.db.Create(comment).Error
}

func (r *Repository) EnsureUser(username, displayName string) (models.User, error) {
    user := models.User{}
    err := r.db.Where(models.User{Username: username}).FirstOrCreate(&user, models.User{Username: username, DisplayName: displayName}).Error
    return user, err
}

