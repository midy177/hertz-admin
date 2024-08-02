// Code generated by ent, DO NOT EDIT.

package user

import (
	"hertz-admin/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v uint8) predicate.User {
	return predicate.User(sql.FieldEQ(FieldStatus, v))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// Nickname applies equality check predicate on the "nickname" field. It's identical to NicknameEQ.
func Nickname(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldNickname, v))
}

// SideMode applies equality check predicate on the "side_mode" field. It's identical to SideModeEQ.
func SideMode(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldSideMode, v))
}

// BaseColor applies equality check predicate on the "base_color" field. It's identical to BaseColorEQ.
func BaseColor(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBaseColor, v))
}

// ActiveColor applies equality check predicate on the "active_color" field. It's identical to ActiveColorEQ.
func ActiveColor(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldActiveColor, v))
}

// RoleID applies equality check predicate on the "role_id" field. It's identical to RoleIDEQ.
func RoleID(v uint64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRoleID, v))
}

// Mobile applies equality check predicate on the "mobile" field. It's identical to MobileEQ.
func Mobile(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldMobile, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// Wecom applies equality check predicate on the "wecom" field. It's identical to WecomEQ.
func Wecom(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldWecom, v))
}

// Avatar applies equality check predicate on the "avatar" field. It's identical to AvatarEQ.
func Avatar(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAvatar, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdatedAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v uint8) predicate.User {
	return predicate.User(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v uint8) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...uint8) predicate.User {
	return predicate.User(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...uint8) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v uint8) predicate.User {
	return predicate.User(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v uint8) predicate.User {
	return predicate.User(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v uint8) predicate.User {
	return predicate.User(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v uint8) predicate.User {
	return predicate.User(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldStatus))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUsername, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPassword, v))
}

// NicknameEQ applies the EQ predicate on the "nickname" field.
func NicknameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldNickname, v))
}

// NicknameNEQ applies the NEQ predicate on the "nickname" field.
func NicknameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldNickname, v))
}

// NicknameIn applies the In predicate on the "nickname" field.
func NicknameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldNickname, vs...))
}

// NicknameNotIn applies the NotIn predicate on the "nickname" field.
func NicknameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldNickname, vs...))
}

// NicknameGT applies the GT predicate on the "nickname" field.
func NicknameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldNickname, v))
}

// NicknameGTE applies the GTE predicate on the "nickname" field.
func NicknameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldNickname, v))
}

// NicknameLT applies the LT predicate on the "nickname" field.
func NicknameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldNickname, v))
}

// NicknameLTE applies the LTE predicate on the "nickname" field.
func NicknameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldNickname, v))
}

// NicknameContains applies the Contains predicate on the "nickname" field.
func NicknameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldNickname, v))
}

// NicknameHasPrefix applies the HasPrefix predicate on the "nickname" field.
func NicknameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldNickname, v))
}

// NicknameHasSuffix applies the HasSuffix predicate on the "nickname" field.
func NicknameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldNickname, v))
}

// NicknameEqualFold applies the EqualFold predicate on the "nickname" field.
func NicknameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldNickname, v))
}

// NicknameContainsFold applies the ContainsFold predicate on the "nickname" field.
func NicknameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldNickname, v))
}

// SideModeEQ applies the EQ predicate on the "side_mode" field.
func SideModeEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldSideMode, v))
}

// SideModeNEQ applies the NEQ predicate on the "side_mode" field.
func SideModeNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldSideMode, v))
}

// SideModeIn applies the In predicate on the "side_mode" field.
func SideModeIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldSideMode, vs...))
}

// SideModeNotIn applies the NotIn predicate on the "side_mode" field.
func SideModeNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldSideMode, vs...))
}

// SideModeGT applies the GT predicate on the "side_mode" field.
func SideModeGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldSideMode, v))
}

// SideModeGTE applies the GTE predicate on the "side_mode" field.
func SideModeGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldSideMode, v))
}

// SideModeLT applies the LT predicate on the "side_mode" field.
func SideModeLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldSideMode, v))
}

// SideModeLTE applies the LTE predicate on the "side_mode" field.
func SideModeLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldSideMode, v))
}

// SideModeContains applies the Contains predicate on the "side_mode" field.
func SideModeContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldSideMode, v))
}

// SideModeHasPrefix applies the HasPrefix predicate on the "side_mode" field.
func SideModeHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldSideMode, v))
}

// SideModeHasSuffix applies the HasSuffix predicate on the "side_mode" field.
func SideModeHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldSideMode, v))
}

// SideModeIsNil applies the IsNil predicate on the "side_mode" field.
func SideModeIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldSideMode))
}

// SideModeNotNil applies the NotNil predicate on the "side_mode" field.
func SideModeNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldSideMode))
}

// SideModeEqualFold applies the EqualFold predicate on the "side_mode" field.
func SideModeEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldSideMode, v))
}

// SideModeContainsFold applies the ContainsFold predicate on the "side_mode" field.
func SideModeContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldSideMode, v))
}

// BaseColorEQ applies the EQ predicate on the "base_color" field.
func BaseColorEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBaseColor, v))
}

// BaseColorNEQ applies the NEQ predicate on the "base_color" field.
func BaseColorNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldBaseColor, v))
}

// BaseColorIn applies the In predicate on the "base_color" field.
func BaseColorIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldBaseColor, vs...))
}

// BaseColorNotIn applies the NotIn predicate on the "base_color" field.
func BaseColorNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldBaseColor, vs...))
}

// BaseColorGT applies the GT predicate on the "base_color" field.
func BaseColorGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldBaseColor, v))
}

// BaseColorGTE applies the GTE predicate on the "base_color" field.
func BaseColorGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldBaseColor, v))
}

// BaseColorLT applies the LT predicate on the "base_color" field.
func BaseColorLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldBaseColor, v))
}

// BaseColorLTE applies the LTE predicate on the "base_color" field.
func BaseColorLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldBaseColor, v))
}

// BaseColorContains applies the Contains predicate on the "base_color" field.
func BaseColorContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldBaseColor, v))
}

// BaseColorHasPrefix applies the HasPrefix predicate on the "base_color" field.
func BaseColorHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldBaseColor, v))
}

// BaseColorHasSuffix applies the HasSuffix predicate on the "base_color" field.
func BaseColorHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldBaseColor, v))
}

// BaseColorIsNil applies the IsNil predicate on the "base_color" field.
func BaseColorIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldBaseColor))
}

// BaseColorNotNil applies the NotNil predicate on the "base_color" field.
func BaseColorNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldBaseColor))
}

// BaseColorEqualFold applies the EqualFold predicate on the "base_color" field.
func BaseColorEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldBaseColor, v))
}

// BaseColorContainsFold applies the ContainsFold predicate on the "base_color" field.
func BaseColorContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldBaseColor, v))
}

// ActiveColorEQ applies the EQ predicate on the "active_color" field.
func ActiveColorEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldActiveColor, v))
}

// ActiveColorNEQ applies the NEQ predicate on the "active_color" field.
func ActiveColorNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldActiveColor, v))
}

// ActiveColorIn applies the In predicate on the "active_color" field.
func ActiveColorIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldActiveColor, vs...))
}

// ActiveColorNotIn applies the NotIn predicate on the "active_color" field.
func ActiveColorNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldActiveColor, vs...))
}

// ActiveColorGT applies the GT predicate on the "active_color" field.
func ActiveColorGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldActiveColor, v))
}

// ActiveColorGTE applies the GTE predicate on the "active_color" field.
func ActiveColorGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldActiveColor, v))
}

// ActiveColorLT applies the LT predicate on the "active_color" field.
func ActiveColorLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldActiveColor, v))
}

// ActiveColorLTE applies the LTE predicate on the "active_color" field.
func ActiveColorLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldActiveColor, v))
}

// ActiveColorContains applies the Contains predicate on the "active_color" field.
func ActiveColorContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldActiveColor, v))
}

// ActiveColorHasPrefix applies the HasPrefix predicate on the "active_color" field.
func ActiveColorHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldActiveColor, v))
}

// ActiveColorHasSuffix applies the HasSuffix predicate on the "active_color" field.
func ActiveColorHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldActiveColor, v))
}

// ActiveColorIsNil applies the IsNil predicate on the "active_color" field.
func ActiveColorIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldActiveColor))
}

// ActiveColorNotNil applies the NotNil predicate on the "active_color" field.
func ActiveColorNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldActiveColor))
}

// ActiveColorEqualFold applies the EqualFold predicate on the "active_color" field.
func ActiveColorEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldActiveColor, v))
}

// ActiveColorContainsFold applies the ContainsFold predicate on the "active_color" field.
func ActiveColorContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldActiveColor, v))
}

// RoleIDEQ applies the EQ predicate on the "role_id" field.
func RoleIDEQ(v uint64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRoleID, v))
}

// RoleIDNEQ applies the NEQ predicate on the "role_id" field.
func RoleIDNEQ(v uint64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldRoleID, v))
}

// RoleIDIn applies the In predicate on the "role_id" field.
func RoleIDIn(vs ...uint64) predicate.User {
	return predicate.User(sql.FieldIn(FieldRoleID, vs...))
}

// RoleIDNotIn applies the NotIn predicate on the "role_id" field.
func RoleIDNotIn(vs ...uint64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldRoleID, vs...))
}

// RoleIDGT applies the GT predicate on the "role_id" field.
func RoleIDGT(v uint64) predicate.User {
	return predicate.User(sql.FieldGT(FieldRoleID, v))
}

// RoleIDGTE applies the GTE predicate on the "role_id" field.
func RoleIDGTE(v uint64) predicate.User {
	return predicate.User(sql.FieldGTE(FieldRoleID, v))
}

// RoleIDLT applies the LT predicate on the "role_id" field.
func RoleIDLT(v uint64) predicate.User {
	return predicate.User(sql.FieldLT(FieldRoleID, v))
}

// RoleIDLTE applies the LTE predicate on the "role_id" field.
func RoleIDLTE(v uint64) predicate.User {
	return predicate.User(sql.FieldLTE(FieldRoleID, v))
}

// RoleIDIsNil applies the IsNil predicate on the "role_id" field.
func RoleIDIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldRoleID))
}

// RoleIDNotNil applies the NotNil predicate on the "role_id" field.
func RoleIDNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldRoleID))
}

// MobileEQ applies the EQ predicate on the "mobile" field.
func MobileEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldMobile, v))
}

// MobileNEQ applies the NEQ predicate on the "mobile" field.
func MobileNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldMobile, v))
}

// MobileIn applies the In predicate on the "mobile" field.
func MobileIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldMobile, vs...))
}

// MobileNotIn applies the NotIn predicate on the "mobile" field.
func MobileNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldMobile, vs...))
}

// MobileGT applies the GT predicate on the "mobile" field.
func MobileGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldMobile, v))
}

// MobileGTE applies the GTE predicate on the "mobile" field.
func MobileGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldMobile, v))
}

// MobileLT applies the LT predicate on the "mobile" field.
func MobileLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldMobile, v))
}

// MobileLTE applies the LTE predicate on the "mobile" field.
func MobileLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldMobile, v))
}

// MobileContains applies the Contains predicate on the "mobile" field.
func MobileContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldMobile, v))
}

// MobileHasPrefix applies the HasPrefix predicate on the "mobile" field.
func MobileHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldMobile, v))
}

// MobileHasSuffix applies the HasSuffix predicate on the "mobile" field.
func MobileHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldMobile, v))
}

// MobileEqualFold applies the EqualFold predicate on the "mobile" field.
func MobileEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldMobile, v))
}

// MobileContainsFold applies the ContainsFold predicate on the "mobile" field.
func MobileContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldMobile, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "email" field.
func EmailIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "email" field.
func EmailNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmail, v))
}

// WecomEQ applies the EQ predicate on the "wecom" field.
func WecomEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldWecom, v))
}

// WecomNEQ applies the NEQ predicate on the "wecom" field.
func WecomNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldWecom, v))
}

// WecomIn applies the In predicate on the "wecom" field.
func WecomIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldWecom, vs...))
}

// WecomNotIn applies the NotIn predicate on the "wecom" field.
func WecomNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldWecom, vs...))
}

// WecomGT applies the GT predicate on the "wecom" field.
func WecomGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldWecom, v))
}

// WecomGTE applies the GTE predicate on the "wecom" field.
func WecomGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldWecom, v))
}

// WecomLT applies the LT predicate on the "wecom" field.
func WecomLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldWecom, v))
}

// WecomLTE applies the LTE predicate on the "wecom" field.
func WecomLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldWecom, v))
}

// WecomContains applies the Contains predicate on the "wecom" field.
func WecomContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldWecom, v))
}

// WecomHasPrefix applies the HasPrefix predicate on the "wecom" field.
func WecomHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldWecom, v))
}

// WecomHasSuffix applies the HasSuffix predicate on the "wecom" field.
func WecomHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldWecom, v))
}

// WecomIsNil applies the IsNil predicate on the "wecom" field.
func WecomIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldWecom))
}

// WecomNotNil applies the NotNil predicate on the "wecom" field.
func WecomNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldWecom))
}

// WecomEqualFold applies the EqualFold predicate on the "wecom" field.
func WecomEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldWecom, v))
}

// WecomContainsFold applies the ContainsFold predicate on the "wecom" field.
func WecomContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldWecom, v))
}

// AvatarEQ applies the EQ predicate on the "avatar" field.
func AvatarEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAvatar, v))
}

// AvatarNEQ applies the NEQ predicate on the "avatar" field.
func AvatarNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldAvatar, v))
}

// AvatarIn applies the In predicate on the "avatar" field.
func AvatarIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldAvatar, vs...))
}

// AvatarNotIn applies the NotIn predicate on the "avatar" field.
func AvatarNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldAvatar, vs...))
}

// AvatarGT applies the GT predicate on the "avatar" field.
func AvatarGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldAvatar, v))
}

// AvatarGTE applies the GTE predicate on the "avatar" field.
func AvatarGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldAvatar, v))
}

// AvatarLT applies the LT predicate on the "avatar" field.
func AvatarLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldAvatar, v))
}

// AvatarLTE applies the LTE predicate on the "avatar" field.
func AvatarLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldAvatar, v))
}

// AvatarContains applies the Contains predicate on the "avatar" field.
func AvatarContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldAvatar, v))
}

// AvatarHasPrefix applies the HasPrefix predicate on the "avatar" field.
func AvatarHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldAvatar, v))
}

// AvatarHasSuffix applies the HasSuffix predicate on the "avatar" field.
func AvatarHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldAvatar, v))
}

// AvatarIsNil applies the IsNil predicate on the "avatar" field.
func AvatarIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldAvatar))
}

// AvatarNotNil applies the NotNil predicate on the "avatar" field.
func AvatarNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldAvatar))
}

// AvatarEqualFold applies the EqualFold predicate on the "avatar" field.
func AvatarEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldAvatar, v))
}

// AvatarContainsFold applies the ContainsFold predicate on the "avatar" field.
func AvatarContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldAvatar, v))
}

// HasToken applies the HasEdge predicate on the "token" edge.
func HasToken() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, TokenTable, TokenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTokenWith applies the HasEdge predicate on the "token" edge with a given conditions (other predicates).
func HasTokenWith(preds ...predicate.Token) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TokenInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, TokenTable, TokenColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
