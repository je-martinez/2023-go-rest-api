package entities

import (
	"time"
)

type Post struct {
	PostID        string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID        string
	Content       string
	Comment_Count int64
	CreatedBy     string
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedBy     string
	UpdatedAt     time.Time `gorm:"default:null"`
	Active        bool      `gorm:"default:true"`
	Comments      []Comment
	Files         []File
	Reactions     []Reaction
}

type PostDTO struct {
	PostID        string
	UserID        string
	Content       string
	Comment_Count int64
	Comments      []CommentDTO
	Files         []FileDTO
	Reactions     []ReactionDTO
}

func (i Post) ToDTO() *PostDTO {
	var comments []CommentDTO
	var files []FileDTO
	var reactions []ReactionDTO

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
