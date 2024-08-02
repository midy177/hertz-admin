// Code generated by hertz generator.
// package admin provides the admin service handler.

package admin

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"hertz-admin/api/model/admin"
	"hertz-admin/api/model/base"
	logic "hertz-admin/biz/logic/admin"
	"hertz-admin/data"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// InitDatabase .
// @router /api/initDatabase [GET]
func InitDatabase(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.Empty
	resp := new(base.BaseResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = base.StatusCode_Fail
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	err = logic.NewInitDatabase(data.Default().DBClient, data.CasbinEnforcer()).InitDatabase(ctx)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	resp.StatusCode = base.StatusCode_Success
	resp.StatusMsg = "success"
	c.JSON(consts.StatusOK, resp)
}

// HealthCheck .
// @router /api/health [GET]
func HealthCheck(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.Empty
	resp := new(base.BaseResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = base.StatusCode_Fail
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	resp.StatusCode = base.StatusCode_Success
	resp.StatusMsg = "success"
	c.JSON(consts.StatusOK, resp)
}

// Captcha .
// @router /api/captcha [GET]
func Captcha(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.Empty
	resp := new(admin.CaptchaInfoResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = base.StatusCode_Fail
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	// GetCaptcha
	id, b64s, err := logic.NewCaptcha().GetCaptcha()
	resp.StatusCode = base.StatusCode_Success
	resp.StatusMsg = "success"
	resp.Data = &admin.CaptchaInfo{
		CaptchaID: id,
		ImgPath:   b64s,
	}
	c.JSON(consts.StatusOK, resp)
}

// DeleteStructTag .
// @router /api/deleteStructTag [POST]
func DeleteStructTag(ctx context.Context, c *app.RequestContext) {
	var err error
	var req admin.StructReq
	resp := new(admin.StructResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = base.StatusCode_Fail
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	// delete struct tag
	list := strings.Split(req.StructStr, "\n")
	var sBuilder strings.Builder
	isInner := false
	for _, l := range list {
		l = strings.TrimSpace(l)
		// first struct message line
		if name := getStructName(l); name != "" {
			sBuilder.WriteString(fmt.Sprintf("type %s struct {\n", name))
			isInner = true
			continue
		}
		// no change to comments
		if strings.Contains(l, "//") {
			if isInner {
				sBuilder.WriteString(fmt.Sprintf("  %s\n", l))
				continue
			}
			sBuilder.WriteString(fmt.Sprintf("%s\n", l))
			continue
		}
		// replace multiple spaces to one
		re := regexp.MustCompile(`\s+`)
		l = re.ReplaceAllString(l, " ")
		// delete struct tag
		lList := strings.Split(l, " ")
		if len(lList) >= 2 {
			sBuilder.WriteString(fmt.Sprintf("  %s %s \n", lList[0], lList[1]))
		}
	}
	// end struct message line
	sBuilder.WriteString("}\n")

	resp.StatusCode = base.StatusCode_Success
	resp.StatusMsg = "success"
	resp.StructStr = sBuilder.String()

	c.JSON(consts.StatusOK, resp)
}

// StructToProto .
// @router /api/structToProto [POST]
func StructToProto(ctx context.Context, c *app.RequestContext) {
	var err error
	var req admin.StructReq
	resp := new(admin.ProtoResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusCode = base.StatusCode_Fail
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	// StructToProto
	list := strings.Split(req.StructStr, "\n")
	var sBuilder strings.Builder
	sort := 1
	isInner := false
	for _, l := range list {
		l = strings.TrimSpace(l)
		// first proto message line
		if name := getStructName(l); name != "" {
			sBuilder.WriteString(fmt.Sprintf("message %s {\n", name))
			isInner = true
			continue
		}
		// no change to comments
		if strings.Contains(l, "//") {
			if isInner {
				sBuilder.WriteString(fmt.Sprintf("  %s\n", l))
				continue
			}
			sBuilder.WriteString(fmt.Sprintf("%s\n", l))
			continue
		}
		// transform to proto
		lList := strings.Split(l, " ")
		// delete empty string "" in list
		var nList []string
		for _, ll := range lList {
			if ll != "" {
				nList = append(nList, ll)
			}
		}
		if len(nList) >= 2 {
			sBuilder.WriteString(fmt.Sprintf("  %s %s = %d;\n", getProtoType(nList[1]), smallCamelString(nList[0]), sort))
			sort++
		}
	}
	// end proto message line
	sBuilder.WriteString("}\n")

	resp.StatusCode = base.StatusCode_Success
	resp.StatusMsg = "success"
	resp.ProtoStr = sBuilder.String()

	c.JSON(consts.StatusOK, resp)
}

func smallCamelString(s string) string {
	if len(s) == 0 {
		return s
	}
	if len(s) == 1 {
		return strings.ToLower(s)
	}
	if s == "ID" {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func getProtoType(s string) string {
	switch s {
	case "int":
		return "int32"
	case "time.Time":
		return "string"
	case "float32":
		return "float"
	case "float64":
		return "double"
	case "[]byte":
		return "bytes"
	case "[]string":
		return "repeated string"
	case "[]int":
		return "repeated int32"
	case "[]int64":
		return "repeated int64"
	case "[]float32", "[]float64":
		return "repeated float"
	default:
		s = strings.ReplaceAll(s, "[]", "repeated ")
		return s
	}
}

func getStructName(s string) string {
	re := regexp.MustCompile(`type +([a-zA-Z1-9]+) +struct {`)
	list := re.FindStringSubmatch(s)
	if len(list) > 1 {
		return list[1]
	}
	return ""
}
