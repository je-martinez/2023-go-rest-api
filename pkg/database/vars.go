package database

import (
	e "github.com/je-martinez/2023-go-rest-api/pkg/database/entities"
	r "github.com/je-martinez/2023-go-rest-api/pkg/database/repository"
)

var UserRepository *r.GormRepository[e.User]
var PostRepository *r.GormRepository[e.Post]
var FileRepository *r.GormRepository[e.File]
