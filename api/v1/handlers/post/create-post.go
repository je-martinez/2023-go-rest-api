package post_handlers

import (
	"fmt"
	"main/pkg/DTOs"
	"main/pkg/constants"
	"main/pkg/database"
	"main/pkg/utils"

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

	fmt.Printf("%+v\n", post)

	newPost := post.ToEntity(post.Content, currentUser.UserID)

	database.PostRepository.Create(newPost)

	utils.GinApiResponse(c, 200, "", newPost.ToDTO(), nil)
}
