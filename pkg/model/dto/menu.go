package dto

import (
	"boyi/configuration"
	"boyi/internal/claims"
	"database/sql/driver"
	"encoding/csv"
	"encoding/json"
	"os"

	"boyi/pkg/infra/errors"
	"boyi/pkg/model/enums/types"

	"github.com/rs/zerolog/log"
)

// 表單類型
type MenuCategory int

const (
	//大分類
	MenuMainCategory MenuCategory = iota + 1
	//副分類
	MenuSubCategory
	//接口
	MenuApiCategory
)

// Menu 樹經評估不納入資料表內
type Menu struct {
	Name      string         //表單名稱
	Key       ManagerMenuKey //表單鍵值
	SuperKey  ManagerMenuKey //父鍵值
	Router    string         //Api路由
	PublicAPI bool           //是否為開放性接口（略過權限判定）
	Next      []*Menu        //下一層
}

// TableName return database table name
func (s Menu) TableName() string {
	return "menu_tree"
}

type ManagerMenuKey string

func (m ManagerMenuKey) String() string {
	return string(m)
}

const (
	Manager_System ManagerMenuKey = "Manager_System"

	Manager_Management          ManagerMenuKey = "Manager_Management"
	API_Manager_Get             ManagerMenuKey = "API_Manager_Get"
	API_Manager_Update          ManagerMenuKey = "API_Manager_Update"
	API_Manager_Delete          ManagerMenuKey = "API_Manager_Delete"
	API_Manager_Create          ManagerMenuKey = "API_Manager_Create"
	API_Manager_Password_Update ManagerMenuKey = "API_Manager_Password_Update"

	CustomerService_Management          ManagerMenuKey = "CustomerService_Management"
	API_CustomerService_Get             ManagerMenuKey = "API_CustomerService_Get"
	API_CustomerService_Update          ManagerMenuKey = "API_CustomerService_Update"
	API_CustomerService_Delete          ManagerMenuKey = "API_CustomerService_Delete"
	API_CustomerService_Create          ManagerMenuKey = "API_CustomerService_Create"
	API_CustomerService_Password_Update ManagerMenuKey = "API_CustomerService_Password_Update"

	Manager_Role_Management ManagerMenuKey = "Manager_Role_Management"
	API_Manager_Role_Create ManagerMenuKey = "API_Manager_Role_Create"
	API_Manager_Role_Get    ManagerMenuKey = "API_Manager_Role_Get"
	API_Manager_Role_Update ManagerMenuKey = "API_Manager_Role_Update"
	API_Manager_Role_Delete ManagerMenuKey = "API_Manager_Role_Delete"

	Member_System     ManagerMenuKey = "Member_System"
	Member_Management ManagerMenuKey = "Member_Management"
	API_Member_Get    ManagerMenuKey = "API_Member_Get"
	API_Member_Update ManagerMenuKey = "API_Member_Update"
	API_Member_Delete ManagerMenuKey = "API_Member_Delete"

	Member_Role_Management ManagerMenuKey = "Member_Role_Management"
	API_Member_Role_Create ManagerMenuKey = "API_Member_Role_Create"
	API_Member_Role_Get    ManagerMenuKey = "API_Member_Role_Get"
	API_Member_Role_Update ManagerMenuKey = "API_Member_Role_Update"
	API_Member_Role_Delete ManagerMenuKey = "API_Member_Role_Delete"

	Member_Tag_Management ManagerMenuKey = "Member_Tag_Management"
	Member_Tag_Create     ManagerMenuKey = "Member_Tag_Create"
	API_Member_Tag_Create ManagerMenuKey = "API_Member_Tag_Create"
	API_Member_Tag_Get    ManagerMenuKey = "API_Member_Tag_Get"
	API_Member_Tag_Update ManagerMenuKey = "API_Member_Tag_Update"
	API_Member_Tag_Delete ManagerMenuKey = "API_Member_Tag_Delete"

	Member_Room_Management ManagerMenuKey = "Member_Room_Management"
	Member_Room_Get        ManagerMenuKey = "Member_Room_Get"
	API_Member_Room_Get    ManagerMenuKey = "API_Member_Room_Get"
	API_Member_Room_Delete ManagerMenuKey = "API_Member_Room_Delete"

	Chatroom_System ManagerMenuKey = "Chatroom_System"

	MyChatroom_Management ManagerMenuKey = "MyChatroom_Management"
	API_MyChatroom_Get    ManagerMenuKey = "API_MyChatroom_Get"
	API_MyChatroom        ManagerMenuKey = "API_MyChatroom"

	// 客服系統
	CustomerService_System ManagerMenuKey = "CustomerService_System"

	ConsultationChat_Management ManagerMenuKey = "ConsultationChat_Management"
	API_ConsultationChat_Get    ManagerMenuKey = "API_ConsultationChat_Get"
	API_JoinConsultingRoom      ManagerMenuKey = "API_JoinConsultingRoom"
	API_PullInConsultingRoom    ManagerMenuKey = "API_PullInConsultingRoom"
	API_UpdateMessage           ManagerMenuKey = "API_UpdateMessage"
	MothballConsultationChat    ManagerMenuKey = "MothballConsultationChat"

	ContactCustomerService_Management ManagerMenuKey = "ContactCustomerService_Management"
	API_ContactCustomerService_Get    ManagerMenuKey = "API_ContactCustomerService_Get"
	API_ContactCustomerService_Update ManagerMenuKey = "API_ContactCustomerService_Update"
	API_ContactCustomerService_Delete ManagerMenuKey = "API_ContactCustomerService_Delete"
	API_ContactCustomerService_Create ManagerMenuKey = "API_ContactCustomerService_Create"

	// 信息系統
	Message_System ManagerMenuKey = "Message_System"

	Push_Management ManagerMenuKey = "Push_Management"
	API_Push_Get    ManagerMenuKey = "API_Push_Get"
	API_Push_Create ManagerMenuKey = "API_Push_Create"

	// 運營系統
	Operating_System ManagerMenuKey = "Operating_System"

	FAQ_Management ManagerMenuKey = "FAQ_Management"
	API_FAQ_Get    ManagerMenuKey = "API_FAQ_Get"
	API_FAQ_Create ManagerMenuKey = "API_FAQ_Create"

	ConsultingRoomForm_Management ManagerMenuKey = "ConsultingRoomForm_Management"
	API_ConsultingRoomForm_Get    ManagerMenuKey = "API_ConsultingRoomForm_Get"
	API_ConsultingRoomForm_Create ManagerMenuKey = "API_ConsultingRoomForm_Create"
	API_ConsultingRoomForm_Update ManagerMenuKey = "API_ConsultingRoomForm_Update"

	CannedResponse_Management ManagerMenuKey = "CannedResponse_Management"
	API_CannedResponse_Get    ManagerMenuKey = "API_CannedResponse_Get"
	API_CannedResponse_Create ManagerMenuKey = "API_CannedResponse_Create"
	API_CannedResponse_Update ManagerMenuKey = "API_CannedResponse_Update"

	API_CannedResponseCategory_Get    ManagerMenuKey = "API_CannedResponseCategory_Get"
	API_CannedResponseCategory_Create ManagerMenuKey = "API_CannedResponseCategory_Create"
	API_CannedResponseCategory_Update ManagerMenuKey = "API_CannedResponseCategory_Update"

	ConsultingGreeting_Management ManagerMenuKey = "ConsultingGreeting_Management"
	API_ConsultingGreeting_Get    ManagerMenuKey = "API_ConsultingGreeting_Get"
	API_ConsultingGreeting_Update ManagerMenuKey = "API_ConsultingGreeting_Update"

	PlatformMoreInformation_Management ManagerMenuKey = "PlatformMoreInformation_Management"
	API_PlatformMoreInformation_Get    ManagerMenuKey = "API_PlatformMoreInformation_Get"
	API_PlatformMoreInformation_Update ManagerMenuKey = "API_PlatformMoreInformation_Update"

	PlatformSetting_Management               ManagerMenuKey = "PlatformSetting_Management"
	API_ConsultingRoomSetting_Beep           ManagerMenuKey = "API_ConsultingRoomSetting_Beep"
	API_ConsultingRoomSetting_AutoDistribute ManagerMenuKey = "API_ConsultingRoomSetting_AutoDistribute"
	API_ConsultingRoomSetting_RedirectVerify ManagerMenuKey = "API_ConsultingRoomSetting_RedirectVerify"

	ConsultingRoomClient_Management        ManagerMenuKey = "ConsultingRoomClient_Management"
	API_ConsultingRoomClientSetting_Get    ManagerMenuKey = "API_ConsultingRoomClientSetting_Get"
	API_ConsultingRoomClientSetting_Create ManagerMenuKey = "API_ConsultingRoomClientSetting_Create"
	API_ConsultingRoomClientSetting_Update ManagerMenuKey = "API_ConsultingRoomClientSetting_Update"
	API_ConsultingRoomClientSetting_Delete ManagerMenuKey = "API_ConsultingRoomClientSetting_Delete"

	ConsultingRoomOrigin_Management   ManagerMenuKey = "ConsultingRoomOrigin_Management"
	API_SpinachPlatformSetting_ALL    ManagerMenuKey = "API_SpinachPlatformSetting_ALL"
	API_SpinachPlatformSetting_Create ManagerMenuKey = "API_SpinachPlatformSetting_Create"
	API_SpinachPlatformSetting_Update ManagerMenuKey = "API_SpinachPlatformSetting_Update"
	API_SpinachPlatformSetting_Delete ManagerMenuKey = "API_SpinachPlatformSetting_Delete"
	API_ConsultingRoomOrigin_ALL      ManagerMenuKey = "API_ConsultingRoomOrigin_ALL"
	API_ConsultingRoomOrigin_Create   ManagerMenuKey = "API_ConsultingRoomOrigin_Create"
	API_ConsultingRoomOrigin_Update   ManagerMenuKey = "API_ConsultingRoomOrigin_Update"
	API_ConsultingRoomOrigin_Delete   ManagerMenuKey = "API_ConsultingRoomOrigin_Delete"

	ConsultingRoomQuestion_Management ManagerMenuKey = "ConsultingRoomQuestion_Management"
	API_ConsultingRoomQuestion_Get    ManagerMenuKey = "API_ConsultingRoomQuestion_Get"
	API_ConsultingRoomQuestion_Create ManagerMenuKey = "API_ConsultingRoomQuestion_Create"
	API_ConsultingRoomQuestion_Update ManagerMenuKey = "API_ConsultingRoomQuestion_Update"
	API_ConsultingRoomQuestion_Delete ManagerMenuKey = "API_ConsultingRoomQuestion_Delete"

	Merchant_Management ManagerMenuKey = "Merchant_Management"
	API_Merchant_Get    ManagerMenuKey = "API_Merchant_Get"
	API_Merchant_Create ManagerMenuKey = "API_Merchant_Create"
	API_Merchant_Update ManagerMenuKey = "API_Merchant_Update"
	API_Merchant_Delete ManagerMenuKey = "API_Merchant_Delete"

	HostsDeny_Management ManagerMenuKey = "HostsDeny_Management"
	API_HostsDeny_Get    ManagerMenuKey = "API_HostsDeny_Get"
	API_HostsDeny_Create ManagerMenuKey = "API_HostsDeny_Create"
	API_HostsDeny_Update ManagerMenuKey = "API_HostsDeny_Update"
	API_HostsDeny_Delete ManagerMenuKey = "API_HostsDeny_Delete"

	AdTracking_Management ManagerMenuKey = "AdTracking_Management"
	API_AdTracking_Get    ManagerMenuKey = "API_AdTracking_Get"
	API_AdTracking_Create ManagerMenuKey = "API_AdTracking_Create"
	API_AdTracking_Update ManagerMenuKey = "API_AdTracking_Update"
	API_AdTracking_Delete ManagerMenuKey = "API_AdTracking_Delete"

	Report_System                    ManagerMenuKey = "Report_System"
	CommonUserReport_Management      ManagerMenuKey = "CommonUserReport_Management"
	CustomerServiceReport_Management ManagerMenuKey = "CustomerServiceReport_Management"
	ConsultingRoomReport_Management  ManagerMenuKey = "ConsultingRoomReport_Management"

	Surveillance_System     ManagerMenuKey = "Surveillance_System"
	Surveillance_Management ManagerMenuKey = "Surveillance_Management"

	Dashboard ManagerMenuKey = "Dashboard"
)

// 員工系統
// 	|------ 員工管理
// 			角色管理
//          客服管理
// 會員系統
// 	|------ 會員管理
// 			角色管理
// 			標籤管理
//          聊天室管理

//後台Menu 基礎物件，基本上顯示順序與建立順序有關

var (
	_menus   []Menu
	_menuMap map[ManagerMenuKey]Menu
)

func GetMenu() []Menu {

	tmp := make([]Menu, len(_menus))
	_ = copy(tmp, _menus)

	return tmp
}

func GetMenuMap() map[ManagerMenuKey]Menu {
	return _menuMap
}

func SetMenu(app *configuration.App) error {
	if app == nil {
		return errors.NewWithMessagef(errors.ErrInternalError, "config loss")
	}
	if app.MenuFilePath == "" {
		return errors.NewWithMessagef(errors.ErrInternalError, "config menu loss, path: %s", app.MenuFilePath)
	}

	f, err := os.Open(app.MenuFilePath)
	if err != nil {
		return errors.NewWithMessagef(errors.ErrInternalError, "fail to read menu file, err: %+v", err)
	}
	defer func() {
		_ = f.Close()
	}()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return errors.NewWithMessagef(errors.ErrInternalError, "fail to read csv records")
	}

	_menus = make([]Menu, len(records))
	for i := range records {
		if len(records[i]) != 3 {
			log.Err(errors.ErrInternalError).Msgf("fail to read menu csv records: line: %d", i+1)
			continue
		}
		_menus[i].SuperKey = ManagerMenuKey(records[i][0])
		_menus[i].Key = ManagerMenuKey(records[i][1])
		_menus[i].Name = records[i][2]
	}

	_menuMap = make(map[ManagerMenuKey]Menu)
	for i := range _menus {
		_menuMap[_menus[i].Key] = _menus[i]
	}

	return nil
}

func GetAllAuthority() Authority {
	tmp := make(Authority)

	menus := GetMenu()
	for i := range menus {
		tmp[menus[i].Key] = struct{}{}
	}

	return tmp
}

func GetDefaultAuthority(app *configuration.App, accountType types.AccountType) (Authority, error) {
	path := app.MenuDefaultAdminFilePath
	if accountType == types.AccountType__CustomerService {
		path = app.MenuDefaultCSFilePath
	}

	if app == nil {
		return nil, errors.NewWithMessagef(errors.ErrInternalError, "config loss")
	}

	if path == "" {
		return GetAllAuthority(), nil
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, errors.NewWithMessagef(errors.ErrInternalError, "fail to read menu file, err: %+v", err)
	}
	defer func() {
		_ = f.Close()
	}()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, errors.NewWithMessagef(errors.ErrInternalError, "fail to read csv records")
	}

	tmp := make(Authority)
	for i := range records {
		if len(records[i]) != 3 {
			log.Err(errors.ErrInternalError).Msgf("fail to read menu csv records: line: %d", i+1)
			continue
		}
		tmp[ManagerMenuKey(records[i][1])] = struct{}{}
	}

	return tmp, nil
}

func GetParsedMenu(c claims.Claims) []*Menu {
	menus := GetMenu()

	tree := []*Menu{}
	tmp := map[ManagerMenuKey]*Menu{}
	for i := range menus {
		if types.AccountType(c.AccountType) != types.AccountType__System {
			if _, ok := c.Competences[menus[i].Key.String()]; !ok {
				continue
			}
		}
		tmp[menus[i].Key] = &menus[i]
		if menus[i].SuperKey == "" {
			tree = append(tree, &menus[i])
		} else {
			if super, ok := tmp[menus[i].SuperKey]; ok {
				if super.Next == nil {
					super.Next = make([]*Menu, 0)
				}
				super.Next = append(super.Next, &menus[i])
			}
		}
	}

	return tree
}

func GetAuthorityFromScopes(scopes ...string) Authority {
	tmp := make(Authority)
	for i := range scopes {
		tmp[ManagerMenuKey(scopes[i])] = struct{}{}
	}
	return tmp
}

type Authority map[ManagerMenuKey]struct{}

func (m Authority) GetMenus() []*Menu {
	menus := _menuMap
	tmp := make([]*Menu, 0)
	for menuKey := range m {
		if v, ok := menus[menuKey]; ok {
			tmp = append(tmp, &v)
		}
	}

	return MenusConvertToTree(tmp)
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (menuTree *Authority) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	return json.Unmarshal(value.([]byte), menuTree)
}

// Value return json value, implement driver.Valuer interface
func (menuTree Authority) Value() (driver.Value, error) {
	if len(menuTree) == 0 {
		return nil, nil
	}

	b, err := json.Marshal(menuTree)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(b).MarshalJSON()
}

func (menuTree Authority) ToMap() map[string]bool {
	var (
		result = make(map[string]bool)
	)

	for k := range menuTree {
		result[string(k)] = true
	}
	return result
}

func (Authority) GormDataType() string {
	return "json"
}

func MenusConvertToTree(result []*Menu) []*Menu {
	tree := []*Menu{}
	tmp := map[ManagerMenuKey]*Menu{}
	for i := range result {
		tmp[result[i].Key] = result[i]
		if result[i].SuperKey == "" {
			tree = append(tree, result[i])
		}
	}

	for i := range result {
		if result[i].SuperKey != "" {
			if super, ok := tmp[result[i].SuperKey]; ok {
				if super.Next == nil {
					super.Next = make([]*Menu, 0)
				}
				super.Next = append(super.Next, result[i])
			}
		}
	}

	return tree
}
