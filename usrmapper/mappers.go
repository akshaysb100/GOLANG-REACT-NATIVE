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

func ToUserDTO(userDTO entity.User) entity.User {
	return entity.User{
		ID:           userDTO.ID,
		USERNAME:     userDTO.USERNAME,
		MOBILENUMBER: userDTO.MOBILENUMBER,
	}
}

// func ToProductDTOs(products []entity.User) []entity.UserDTO {
// 	users := make([]entity.UserDTO, len(products))

// 	for i, itm := range products {
// 		users[i] = usrmapper.ToUserDTO(itm)
// 	}

// 	return users
// }
