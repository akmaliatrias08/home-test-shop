package role

import (
	"online-shop-home-test/modules/helpers"
	role "online-shop-home-test/modules/role/dto"
	"online-shop-home-test/modules/role/models"
	"online-shop-home-test/modules/utils"

	"gorm.io/gorm/clause"
)

func GetAllRole(page, pageSize string) (int64, []models.Role, error) {
	var count int64
	var roles []models.Role

	//handle pagination
	limit, offset, err := helpers.Pagination(page, pageSize)
	if err != nil {
		return 0, roles, err
	}

	//count all roles
	if err := utils.DB.Model(&roles).Count(&count).Error; err != nil {
		return count, roles, err
	}

	//get all roles
	if err := utils.DB.Model(&roles).Limit(limit).Offset(offset).Find(&roles).Error; err != nil {
		return count, roles, err
	}

	return count, roles, nil
}

func GetRoleById(id string) (models.Role, error) {
	var role models.Role

	if err := utils.DB.Model(&role).Where("id", id).First(&role).Error; err != nil {
		return role, err
	}

	return role, nil
}

func CreateRole(createRoleDTO role.CreateRoleDTO) (models.Role, error) {
	var role models.Role

	role.Name = createRoleDTO.Name

	if err := utils.DB.Create(&role).Error; err != nil {
		return role, err
	}

	return role, nil
}

func UpdateRole(id string, UpdateRoleDTO role.UpdateRoleDTO) (models.Role, error) {
	var role models.Role

	//validate if id exist
	_, err := GetRoleById(id)
	if err != nil {
		return role, err
	}

	if err := utils.DB.Model(&role).Where("id", id).Clauses(clause.Returning{}).Update("name", UpdateRoleDTO.Name).Error; err != nil {
		return role, err
	}

	return role, nil
}

func DeleteRole(id string) error {
	var role models.Role

	//validate if id exist
	_, err := GetRoleById(id)
	if err != nil {
		return err
	}

	if err := utils.DB.Model(&role).Where("id", id).Delete(&role).Error; err != nil {
		return err
	}

	return nil
}
