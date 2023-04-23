package database

import (
	e "main/pkg/database/entities"
	r "main/pkg/database/repository"
)

var UserRepository *r.GormRepository[e.User]
var PostRepository *r.GormRepository[e.Post]
