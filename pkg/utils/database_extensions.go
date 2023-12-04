package utils

import db_extensions "github.com/je-martinez/2023-go-rest-api/pkg/database/extensions"

func IsSupportedReaction(reaction string) (bool, *db_extensions.ReactionType) {
	reactions := db_extensions.GetSupportedReactions()
	for _, supported_reaction := range reactions {
		if reaction == supported_reaction.ToString() {
			return true, &supported_reaction
		}
	}
	return false, nil
}
