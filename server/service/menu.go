package service

import (
	"errors"
	"server/global"
	"server/model/entity"
	"server/model/request"
	"strconv"

	"gorm.io/gorm"
)

//@description: 获取路由总树map

type MenuService struct{}

var MenuServiceApp = new(MenuService)

func (menuService *MenuService) getMenuTreeMap(authorityId uint) (treeMap map[string][]entity.SysMenu, err error) {
	var allMenus []entity.SysMenu
	var baseMenu []entity.SysBaseMenu
	var btns []entity.SysAuthorityBtn
	treeMap = make(map[string][]entity.SysMenu)

	var SysAuthorityMenus []entity.SysAuthorityMenu
	err = global.DB.Where("sys_authority_authority_id = ?", authorityId).Find(&SysAuthorityMenus).Error
	if err != nil {
		return
	}

	var MenuIds []string

	for i := range SysAuthorityMenus {
		MenuIds = append(MenuIds, SysAuthorityMenus[i].MenuId)
	}

	err = global.DB.Where("id in (?)", MenuIds).Order("sort").Preload("Parameters").Find(&baseMenu).Error
	if err != nil {
		return
	}

	for i := range baseMenu {
		allMenus = append(allMenus, entity.SysMenu{
			SysBaseMenu: baseMenu[i],
			AuthorityId: authorityId,
			MenuId:      strconv.Itoa(int(baseMenu[i].ID)),
			Parameters:  baseMenu[i].Parameters,
		})
	}

	err = global.DB.Where("authority_id = ?", authorityId).Preload("SysBaseMenuBtn").Find(&btns).Error
	if err != nil {
		return
	}
	var btnMap = make(map[uint]map[string]uint)
	for _, v := range btns {
		if btnMap[v.SysMenuID] == nil {
			btnMap[v.SysMenuID] = make(map[string]uint)
		}
		btnMap[v.SysMenuID][v.SysBaseMenuBtn.Name] = authorityId
	}
	for _, v := range allMenus {
		v.Btns = btnMap[v.ID]
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}

//@description: 获取动态菜单树

func (menuService *MenuService) GetMenuTree(authorityId uint) (menus []entity.SysMenu, err error) {
	menuTree, err := menuService.getMenuTreeMap(authorityId)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = menuService.getChildrenList(&menus[i], menuTree)
	}
	return menus, err
}

//@description: 获取子菜单

func (menuService *MenuService) getChildrenList(menu *entity.SysMenu, treeMap map[string][]entity.SysMenu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//@description: 获取路由分页

func (menuService *MenuService) GetInfoList() (list interface{}, total int64, err error) {
	var menuList []entity.SysBaseMenu
	treeMap, err := menuService.getBaseMenuTreeMap()
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = menuService.getBaseChildrenList(&menuList[i], treeMap)
	}
	return menuList, total, err
}

//@description: 获取菜单的子菜单

func (menuService *MenuService) getBaseChildrenList(menu *entity.SysBaseMenu, treeMap map[string][]entity.SysBaseMenu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//@description: 添加基础路由

func (menuService *MenuService) AddBaseMenu(menu entity.SysBaseMenu) error {
	if !errors.Is(global.DB.Where("name = ?", menu.Name).First(&entity.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	return global.DB.Create(&menu).Error
}

//@description: 获取路由总树map

func (menuService *MenuService) getBaseMenuTreeMap() (treeMap map[string][]entity.SysBaseMenu, err error) {
	var allMenus []entity.SysBaseMenu
	treeMap = make(map[string][]entity.SysBaseMenu)
	err = global.DB.Order("sort").Preload("MenuBtn").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}

//@description: 获取基础路由树

func (menuService *MenuService) GetBaseMenuTree() (menus []entity.SysBaseMenu, err error) {
	treeMap, err := menuService.getBaseMenuTreeMap()
	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = menuService.getBaseChildrenList(&menus[i], treeMap)
	}
	return menus, err
}

//@description: 为角色增加menu树

func (menuService *MenuService) AddMenuAuthority(menus []entity.SysBaseMenu, authorityId uint) (err error) {
	var auth entity.SysAuthority
	auth.AuthorityId = authorityId
	auth.SysBaseMenus = menus
	err = AuthorityServiceApp.SetMenuAuthority(&auth)
	return err
}

//@description: 查看当前角色树

func (menuService *MenuService) GetMenuAuthority(info *request.GetAuthorityId) (menus []entity.SysMenu, err error) {
	var baseMenu []entity.SysBaseMenu
	var SysAuthorityMenus []entity.SysAuthorityMenu
	err = global.DB.Where("sys_authority_authority_id = ?", info.AuthorityId).Find(&SysAuthorityMenus).Error
	if err != nil {
		return
	}

	var MenuIds []string

	for i := range SysAuthorityMenus {
		MenuIds = append(MenuIds, SysAuthorityMenus[i].MenuId)
	}

	err = global.DB.Where("id in (?) ", MenuIds).Order("sort").Find(&baseMenu).Error

	for i := range baseMenu {
		menus = append(menus, entity.SysMenu{
			SysBaseMenu: baseMenu[i],
			AuthorityId: info.AuthorityId,
			MenuId:      strconv.Itoa(int(baseMenu[i].ID)),
			Parameters:  baseMenu[i].Parameters,
		})
	}
	// sql := "SELECT authority_menu.keep_alive,authority_menu.default_menu,authority_menu.created_at,authority_menu.updated_at,authority_menu.deleted_at,authority_menu.menu_level,authority_menu.parent_id,authority_menu.path,authority_menu.`name`,authority_menu.hidden,authority_menu.component,authority_menu.title,authority_menu.icon,authority_menu.sort,authority_menu.menu_id,authority_menu.authority_id FROM authority_menu WHERE authority_menu.authority_id = ? ORDER BY authority_menu.sort ASC"
	// err = global.GVA_DB.Raw(sql, authorityId).Scan(&menus).Error
	return menus, err
}

// UserAuthorityDefaultRouter 用户角色默认路由检查

func (menuService *MenuService) UserAuthorityDefaultRouter(user *entity.SysUser) {
	var menuIds []string
	err := global.DB.Model(&entity.SysAuthorityMenu{}).Where("sys_authority_authority_id = ?", user.AuthorityId).Pluck("sys_base_menu_id", &menuIds).Error
	if err != nil {
		return
	}
	var am entity.SysBaseMenu
	err = global.DB.First(&am, "name = ? and id in (?)", user.Authority.DefaultRouter, menuIds).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user.Authority.DefaultRouter = "404"
	}
}
