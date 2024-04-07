package adapters

import (
	"fmt"
	"api/src/domain/client"
)


func GetUser(id int) (*client.PropsUser, error) {
	instanceUser := client.PropsUser{}
	entitiesUser := client.UserUseCase(instanceUser)
	
	resEntitiesUser, err := entitiesUser.UserRepository.GetUserById(id)
	fmt.Println(resEntitiesUser)
	if err != nil {
		return nil, nil
	}
	return resEntitiesUser, nil
}