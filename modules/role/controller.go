package role

import (
	"net/http"
	role "online-shop-home-test/modules/role/dto"

	"github.com/gin-gonic/gin"
)

func routes(rg *gin.RouterGroup) {
	roleRoute := rg.Group("/role")

	roleRoute.GET("", getAllRole)
	roleRoute.POST("", createRole)
	roleRoute.PUT("/:id", updateRole)
	roleRoute.DELETE("/:id", deleteRole)

}

// Role      		godoc
// @Summary     	Get all role
// @Description  	Get all role that exist
// @Param        	page    	query     string  false  "Number of page to load"
// @Param        	pageSize    query     string  false  "Number page size or limit to display role data"
// @Tags         	role
// @Produce      	json
// @Success      	200  {array} models.Role
// @Router       	/role [get]
func getAllRole(ctx *gin.Context) {
	count, data, err := GetAllRole(ctx.Query("page"), ctx.Query("pageSize"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    data,
		"count":   count,
		"message": "success",
	})
}

// Role      		godoc
// @Summary     	Create role
// @Description  	Create new role
// @Param			role body role.CreateRoleDTO true  "Create role"
// @Tags         	role
// @Produce      	json
// @Success      	201  {object} models.Role
// @Router       	/role [post]
func createRole(ctx *gin.Context) {
	var input role.CreateRoleDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data, err := CreateRole(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data":    data,
		"message": "success",
	})
}

// Role      		godoc
// @Summary     	Update role by ID
// @Description  	Update role by ID
// @Param			id path string true "Role id to update"
// @Param			role body role.UpdateRoleDTO true  "Update role"
// @Tags         	role
// @Produce      	json
// @Success      	200  {object} models.Role
// @Router       	/role/{id} [put]
func updateRole(ctx *gin.Context) {
	var input role.UpdateRoleDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data, err := UpdateRole(ctx.Param("id"), input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": "success",
	})
}

// Role      		godoc
// @Summary     	Delete role by ID
// @Description  	Delete role by ID
// @Param			id path string true "Role id to delete"
// @Tags         	role
// @Produce      	json
// @Success      	200  {object} models.Role
// @Router       	/role/{id} [delete]
func deleteRole(ctx *gin.Context) {
	err := DeleteRole(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    struct{}{},
		"message": "success",
	})
}
