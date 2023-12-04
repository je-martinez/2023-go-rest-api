package entities

import (
	"time"
)

type Post struct {
	PostID        string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID        string
	Content       string
	Comment_Count int64 `gorm:"default:0"`
	CreatedBy     string
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedBy     *string
	UpdatedAt     time.Time  `gorm:"default:null"`
	Active        bool       `gorm:"default:true"`
	Comments      []Comment  `gorm:"constraint:OnDelete:CASCADE;"`
	Files         []File     `gorm:"constraint:OnDelete:CASCADE;"`
	Reactions     []Reaction `gorm:"constraint:OnDelete:CASCADE;"`
}

type PostDTO struct {
	PostID        string        `json:"post_id"`
	UserID        string        `json:"user_id"`
	Content       string        `json:"content"`
	Comment_Count int64         `json:"comment_count"`
	Comments      []CommentDTO  `json:"comments"`
	Files         []FileDTO     `json:"files"`
	Reactions     []ReactionDTO `json:"reactions"`
}

func (i Post) ToDTO() *PostDTO {
	comments := []CommentDTO{}
	files := []FileDTO{}
	reactions := []ReactionDTO{}

	if i.Comments != nil {
		for _, comment := range i.Comments {
			comments = append(comments, *comment.ToDTO())
		}
	}

	if i.Files != nil {
		for _, file := range i.Files {
			files = append(files, *file.ToDTO())
		}
	}

	if i.Reactions != nil {
		for _, reaction := range i.Reactions {
			reactions = append(reactions, *reaction.ToDTO())
		}
	}

	return &PostDTO{
		PostID:        i.PostID,
		UserID:        i.UserID,
		Content:       i.Content,
		Comment_Count: i.Comment_Count,
		Comments:      comments,
		Files:         files,
		Reactions:     reactions,
	}
}
