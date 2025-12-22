package main

import (
	"log"
	"strings"

	"github.com/aifuxi/fgo/internal/model"
	"github.com/aifuxi/fgo/pkg/db"
)

func main() {
	db.InitMySQL()

	database := db.GetDB()

	database.AutoMigrate(
		&model.Tag{},
		&model.Blog{},
		&model.Category{},
		&model.User{},
		&model.Role{},
		&model.Permission{},
	)

	permissions := []model.Permission{
		{Name: model.PermissionBlogList, Code: model.PermissionBlogList, Description: "博客列表权限"},
		{Name: model.PermissionBlogView, Code: model.PermissionBlogView, Description: "查看博客权限"},
		{Name: model.PermissionBlogCreate, Code: model.PermissionBlogCreate, Description: "创建博客权限"},
		{Name: model.PermissionBlogUpdate, Code: model.PermissionBlogUpdate, Description: "更新博客权限"},
		{Name: model.PermissionBlogDelete, Code: model.PermissionBlogDelete, Description: "删除博客权限"},

		{Name: model.PermissionCategoryList, Code: model.PermissionCategoryList, Description: "分类列表权限"},
		{Name: model.PermissionCategoryView, Code: model.PermissionCategoryView, Description: "查看分类权限"},
		{Name: model.PermissionCategoryCreate, Code: model.PermissionCategoryCreate, Description: "创建分类权限"},
		{Name: model.PermissionCategoryUpdate, Code: model.PermissionCategoryUpdate, Description: "更新分类权限"},
		{Name: model.PermissionCategoryDelete, Code: model.PermissionCategoryDelete, Description: "删除分类权限"},

		{Name: model.PermissionRoleList, Code: model.PermissionRoleList, Description: "角色列表权限"},
		{Name: model.PermissionRoleView, Code: model.PermissionRoleView, Description: "查看角色权限"},
		{Name: model.PermissionRoleCreate, Code: model.PermissionRoleCreate, Description: "创建角色权限"},
		{Name: model.PermissionRoleUpdate, Code: model.PermissionRoleUpdate, Description: "更新角色权限"},
		{Name: model.PermissionRoleDelete, Code: model.PermissionRoleDelete, Description: "删除角色权限"},

		{Name: model.PermissionUserList, Code: model.PermissionUserList, Description: "用户列表权限"},
		{Name: model.PermissionUserView, Code: model.PermissionUserView, Description: "查看用户权限"},
		{Name: model.PermissionUserCreate, Code: model.PermissionUserCreate, Description: "创建用户权限"},
		{Name: model.PermissionUserUpdate, Code: model.PermissionUserUpdate, Description: "更新用户权限"},
		{Name: model.PermissionUserDelete, Code: model.PermissionUserDelete, Description: "删除用户权限"},

		{Name: model.PermissionTagList, Code: model.PermissionTagList, Description: "标签列表权限"},
		{Name: model.PermissionTagView, Code: model.PermissionTagView, Description: "查看标签权限"},
		{Name: model.PermissionTagCreate, Code: model.PermissionTagCreate, Description: "创建标签权限"},
		{Name: model.PermissionTagUpdate, Code: model.PermissionTagUpdate, Description: "更新标签权限"},
		{Name: model.PermissionTagDelete, Code: model.PermissionTagDelete, Description: "删除标签权限"},
	}

	var allPermissions []model.Permission
	var visitorPermissions []model.Permission

	for _, p := range permissions {
		perm := p
		database.Where("code = ?", perm.Code).FirstOrCreate(&perm)
		allPermissions = append(allPermissions, perm)

		if strings.HasSuffix(perm.Code, ":list") || strings.HasSuffix(perm.Code, ":view") {
			visitorPermissions = append(visitorPermissions, perm)
		}
	}

	var adminRole model.Role
	database.Where("code = ?", "admin").FirstOrCreate(&adminRole, model.Role{
		Name:        "admin",
		Code:        "admin",
		Description: "管理员角色",
	})

	var visitorRole model.Role
	database.Where("code = ?", "visitor").FirstOrCreate(&visitorRole, model.Role{
		Name:        "visitor",
		Code:        "visitor",
		Description: "访客角色",
	})

	if err := database.Model(&adminRole).Association("Permissions").Replace(allPermissions); err != nil {
		log.Fatalf("Failed to replace admin role permissions: %v", err)
	}

	if err := database.Model(&visitorRole).Association("Permissions").Replace(visitorPermissions); err != nil {
		log.Fatalf("Failed to replace visitor role permissions: %v", err)
	}
}
