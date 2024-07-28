// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"formulago/data/ent/menu"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Menu is the model entity for the Menu schema.
type Menu struct {
	config `json:"-"`
	// ID of the ent.
	// primary key
	ID uint64 `json:"id,omitempty"`
	// created time
	CreatedAt time.Time `json:"created_at,omitempty"`
	// last update time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// parent menu ID | 父菜单ID
	ParentID uint64 `json:"parent_id,omitempty"`
	// menu level | 菜单层级
	MenuLevel uint32 `json:"menu_level,omitempty"`
	// menu type | 菜单类型 0 目录 1 菜单 2 按钮
	MenuType uint32 `json:"menu_type,omitempty"`
	// index path | 菜单路由路径
	Path string `json:"path,omitempty"`
	// index name | 菜单名称
	Name string `json:"name,omitempty"`
	// redirect path | 跳转路径 （外链）
	Redirect string `json:"redirect,omitempty"`
	// the path of vue file | 组件路径
	Component string `json:"component,omitempty"`
	// sorting numbers | 排序编号
	OrderNo uint32 `json:"order_no,omitempty"`
	// disable status | 是否停用
	Disabled bool `json:"disabled,omitempty"`
	// menu name | 菜单显示标题
	Title string `json:"title,omitempty"`
	// menu icon | 菜单图标
	Icon string `json:"icon,omitempty"`
	// hide menu | 是否隐藏菜单
	HideMenu bool `json:"hide_menu,omitempty"`
	// hide the breadcrumb | 隐藏面包屑
	HideBreadcrumb bool `json:"hide_breadcrumb,omitempty"`
	// set the active menu | 激活菜单
	CurrentActiveMenu string `json:"current_active_menu,omitempty"`
	// do not keep alive the tab | 取消页面缓存
	IgnoreKeepAlive bool `json:"ignore_keep_alive,omitempty"`
	// hide the tab header | 隐藏页头
	HideTab bool `json:"hide_tab,omitempty"`
	// show iframe | 内嵌 iframe
	FrameSrc string `json:"frame_src,omitempty"`
	// the route carries parameters or not | 携带参数
	CarryParam bool `json:"carry_param,omitempty"`
	// hide children menu or not | 隐藏所有子菜单
	HideChildrenInMenu bool `json:"hide_children_in_menu,omitempty"`
	// affix tab | Tab 固定
	Affix bool `json:"affix,omitempty"`
	// the maximum number of pages the router can open | 能打开的子TAB数
	DynamicLevel uint32 `json:"dynamic_level,omitempty"`
	// the real path of the route without dynamic part | 菜单路由不包含参数部分
	RealPath string `json:"real_path,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MenuQuery when eager-loading is set.
	Edges MenuEdges `json:"edges"`
}

// MenuEdges holds the relations/edges for other nodes in the graph.
type MenuEdges struct {
	// Roles holds the value of the roles edge.
	Roles []*Role `json:"roles,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *Menu `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*Menu `json:"children,omitempty"`
	// Params holds the value of the params edge.
	Params []*MenuParam `json:"params,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) RolesOrErr() ([]*Role, error) {
	if e.loadedTypes[0] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MenuEdges) ParentOrErr() (*Menu, error) {
	if e.loadedTypes[1] {
		if e.Parent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: menu.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) ChildrenOrErr() ([]*Menu, error) {
	if e.loadedTypes[2] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// ParamsOrErr returns the Params value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) ParamsOrErr() ([]*MenuParam, error) {
	if e.loadedTypes[3] {
		return e.Params, nil
	}
	return nil, &NotLoadedError{edge: "params"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Menu) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case menu.FieldDisabled, menu.FieldHideMenu, menu.FieldHideBreadcrumb, menu.FieldIgnoreKeepAlive, menu.FieldHideTab, menu.FieldCarryParam, menu.FieldHideChildrenInMenu, menu.FieldAffix:
			values[i] = new(sql.NullBool)
		case menu.FieldID, menu.FieldParentID, menu.FieldMenuLevel, menu.FieldMenuType, menu.FieldOrderNo, menu.FieldDynamicLevel:
			values[i] = new(sql.NullInt64)
		case menu.FieldPath, menu.FieldName, menu.FieldRedirect, menu.FieldComponent, menu.FieldTitle, menu.FieldIcon, menu.FieldCurrentActiveMenu, menu.FieldFrameSrc, menu.FieldRealPath:
			values[i] = new(sql.NullString)
		case menu.FieldCreatedAt, menu.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Menu", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Menu fields.
func (m *Menu) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case menu.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = uint64(value.Int64)
		case menu.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case menu.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		case menu.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				m.ParentID = uint64(value.Int64)
			}
		case menu.FieldMenuLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field menu_level", values[i])
			} else if value.Valid {
				m.MenuLevel = uint32(value.Int64)
			}
		case menu.FieldMenuType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field menu_type", values[i])
			} else if value.Valid {
				m.MenuType = uint32(value.Int64)
			}
		case menu.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				m.Path = value.String
			}
		case menu.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		case menu.FieldRedirect:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field redirect", values[i])
			} else if value.Valid {
				m.Redirect = value.String
			}
		case menu.FieldComponent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field component", values[i])
			} else if value.Valid {
				m.Component = value.String
			}
		case menu.FieldOrderNo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_no", values[i])
			} else if value.Valid {
				m.OrderNo = uint32(value.Int64)
			}
		case menu.FieldDisabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field disabled", values[i])
			} else if value.Valid {
				m.Disabled = value.Bool
			}
		case menu.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				m.Title = value.String
			}
		case menu.FieldIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icon", values[i])
			} else if value.Valid {
				m.Icon = value.String
			}
		case menu.FieldHideMenu:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field hide_menu", values[i])
			} else if value.Valid {
				m.HideMenu = value.Bool
			}
		case menu.FieldHideBreadcrumb:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field hide_breadcrumb", values[i])
			} else if value.Valid {
				m.HideBreadcrumb = value.Bool
			}
		case menu.FieldCurrentActiveMenu:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field current_active_menu", values[i])
			} else if value.Valid {
				m.CurrentActiveMenu = value.String
			}
		case menu.FieldIgnoreKeepAlive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field ignore_keep_alive", values[i])
			} else if value.Valid {
				m.IgnoreKeepAlive = value.Bool
			}
		case menu.FieldHideTab:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field hide_tab", values[i])
			} else if value.Valid {
				m.HideTab = value.Bool
			}
		case menu.FieldFrameSrc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field frame_src", values[i])
			} else if value.Valid {
				m.FrameSrc = value.String
			}
		case menu.FieldCarryParam:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field carry_param", values[i])
			} else if value.Valid {
				m.CarryParam = value.Bool
			}
		case menu.FieldHideChildrenInMenu:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field hide_children_in_menu", values[i])
			} else if value.Valid {
				m.HideChildrenInMenu = value.Bool
			}
		case menu.FieldAffix:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field affix", values[i])
			} else if value.Valid {
				m.Affix = value.Bool
			}
		case menu.FieldDynamicLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field dynamic_level", values[i])
			} else if value.Valid {
				m.DynamicLevel = uint32(value.Int64)
			}
		case menu.FieldRealPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field real_path", values[i])
			} else if value.Valid {
				m.RealPath = value.String
			}
		}
	}
	return nil
}

// QueryRoles queries the "roles" edge of the Menu entity.
func (m *Menu) QueryRoles() *RoleQuery {
	return (&MenuClient{config: m.config}).QueryRoles(m)
}

// QueryParent queries the "parent" edge of the Menu entity.
func (m *Menu) QueryParent() *MenuQuery {
	return (&MenuClient{config: m.config}).QueryParent(m)
}

// QueryChildren queries the "children" edge of the Menu entity.
func (m *Menu) QueryChildren() *MenuQuery {
	return (&MenuClient{config: m.config}).QueryChildren(m)
}

// QueryParams queries the "params" edge of the Menu entity.
func (m *Menu) QueryParams() *MenuParamQuery {
	return (&MenuClient{config: m.config}).QueryParams(m)
}

// Update returns a builder for updating this Menu.
// Note that you need to call Menu.Unwrap() before calling this method if this Menu
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Menu) Update() *MenuUpdateOne {
	return (&MenuClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Menu entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Menu) Unwrap() *Menu {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Menu is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Menu) String() string {
	var builder strings.Builder
	builder.WriteString("Menu(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("parent_id=")
	builder.WriteString(fmt.Sprintf("%v", m.ParentID))
	builder.WriteString(", ")
	builder.WriteString("menu_level=")
	builder.WriteString(fmt.Sprintf("%v", m.MenuLevel))
	builder.WriteString(", ")
	builder.WriteString("menu_type=")
	builder.WriteString(fmt.Sprintf("%v", m.MenuType))
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(m.Path)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteString(", ")
	builder.WriteString("redirect=")
	builder.WriteString(m.Redirect)
	builder.WriteString(", ")
	builder.WriteString("component=")
	builder.WriteString(m.Component)
	builder.WriteString(", ")
	builder.WriteString("order_no=")
	builder.WriteString(fmt.Sprintf("%v", m.OrderNo))
	builder.WriteString(", ")
	builder.WriteString("disabled=")
	builder.WriteString(fmt.Sprintf("%v", m.Disabled))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(m.Title)
	builder.WriteString(", ")
	builder.WriteString("icon=")
	builder.WriteString(m.Icon)
	builder.WriteString(", ")
	builder.WriteString("hide_menu=")
	builder.WriteString(fmt.Sprintf("%v", m.HideMenu))
	builder.WriteString(", ")
	builder.WriteString("hide_breadcrumb=")
	builder.WriteString(fmt.Sprintf("%v", m.HideBreadcrumb))
	builder.WriteString(", ")
	builder.WriteString("current_active_menu=")
	builder.WriteString(m.CurrentActiveMenu)
	builder.WriteString(", ")
	builder.WriteString("ignore_keep_alive=")
	builder.WriteString(fmt.Sprintf("%v", m.IgnoreKeepAlive))
	builder.WriteString(", ")
	builder.WriteString("hide_tab=")
	builder.WriteString(fmt.Sprintf("%v", m.HideTab))
	builder.WriteString(", ")
	builder.WriteString("frame_src=")
	builder.WriteString(m.FrameSrc)
	builder.WriteString(", ")
	builder.WriteString("carry_param=")
	builder.WriteString(fmt.Sprintf("%v", m.CarryParam))
	builder.WriteString(", ")
	builder.WriteString("hide_children_in_menu=")
	builder.WriteString(fmt.Sprintf("%v", m.HideChildrenInMenu))
	builder.WriteString(", ")
	builder.WriteString("affix=")
	builder.WriteString(fmt.Sprintf("%v", m.Affix))
	builder.WriteString(", ")
	builder.WriteString("dynamic_level=")
	builder.WriteString(fmt.Sprintf("%v", m.DynamicLevel))
	builder.WriteString(", ")
	builder.WriteString("real_path=")
	builder.WriteString(m.RealPath)
	builder.WriteByte(')')
	return builder.String()
}

// Menus is a parsable slice of Menu.
type Menus []*Menu

func (m Menus) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}