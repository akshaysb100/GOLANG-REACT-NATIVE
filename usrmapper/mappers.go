package usrmapper

import (
	"GOLANG-REACT-NATIVE/entity"
)

func ToUser(userDTO entity.User) entity.User {
	return entity.User{
		ID:           userDTO.ID,
		USERNAME:     userDTO.USERNAME,
		MOBILENUMBER: userDTO.MOBILENUMBER,
	}
}

func ToIssues(issuesDTO entity.Isuues) entity.Isuues {
	return entity.Isuues{
		ID:    issuesDTO.ID,
		ISSUE: issuesDTO.ISSUE,
	}
}
