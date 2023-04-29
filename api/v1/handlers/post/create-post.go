package post_handlers

import (
	"github.com/je-martinez/2023-go-rest-api/pkg/DTOs"
	"github.com/je-martinez/2023-go-rest-api/pkg/constants"
	"github.com/je-martinez/2023-go-rest-api/pkg/database"
	"github.com/je-martinez/2023-go-rest-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {

	currentUser, errCurrentUser := utils.ExtractUserFromToken(c)

	if errCurrentUser != nil || currentUser == nil {
		utils.GinApiResponse(c, 500, constants.ERR_CURRENT_USER, nil, utils.ValidateStructErrors(errCurrentUser))
		return
	}

	var post DTOs.CreatePostDTO

	err := c.ShouldBind(&post)
	if err != nil {
		utils.GinApiResponse(c, 400, constants.ERR_BIND_MULTIPART, nil, []string{err.Error()})
		return
	}

	newPost := post.ToEntity(post.Content, currentUser.UserID)

	database.PostRepository.Create(newPost)

	utils.GinApiResponse(c, 200, "", newPost.ToDTO(), nil)
}
