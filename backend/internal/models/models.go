package models

import "time"

type Board struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:100;uniqueIndex" json:"name"`
	Description string    `gorm:"size:255" json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Posts       []Post    `json:"posts,omitempty"`
}

type Topic struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:120" json:"title"`
	Description string    `gorm:"size:255" json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Posts       []Post    `json:"posts,omitempty"`
}

type User struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Username    string    `gorm:"size:80;uniqueIndex" json:"username"`
	DisplayName string    `gorm:"size:120" json:"displayName"`
	AvatarURL   string    `gorm:"size:255" json:"avatarUrl"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Post struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"size:160" json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	BoardID   uint      `json:"boardId"`
	TopicID   uint      `json:"topicId"`
	AuthorID  uint      `json:"authorId"`
	Board     Board     `json:"board"`
	Topic     Topic     `json:"topic"`
	Author    User      `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Comments  []Comment `json:"comments,omitempty"`
}

type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Content   string    `gorm:"type:text" json:"content"`
	PostID    uint      `json:"postId"`
	AuthorID  uint      `json:"authorId"`
	Author    User      `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
