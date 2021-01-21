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

// func ToUserDTO(user entity.User) entity.UserDTO {
// 	return entity.UserDTO{
// 		ID:           user.ID,
// 		USERNAME:     user.USERNAME,
// 		MOBILENUMBER: user.MOBILENUMBER,
// 	}
// }

// func ToProductDTOs(products []entity.User) []entity.UserDTO {
// 	users := make([]entity.UserDTO, len(products))

// 	for i, itm := range products {
// 		users[i] = ToUserDTO(itm)
// 	}

// 	return users
// }
