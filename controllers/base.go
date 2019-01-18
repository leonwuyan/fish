package controllers

import (
	"fish/configs"
	"fish/enums"
	"fish/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	"html/template"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

var cpt *captcha.Captcha

type baseController struct {
	beego.Controller
	sc_Chan chan models.Result
}

func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 4
	cpt.StdHeight = 40
	cpt.StdWidth = 100
	beego.AddFuncMap("year", time.Now().Year)
	beego.AddFuncMap("v", verifyPower)
}
func (c *baseController) Prepare() {
	c.Data["rand"] = rand.Int()
	c.Data["site"] = configs.Site
}
func (c *baseController) jsonData(code enums.ReturnCode, params ...interface{}) models.Result {
	var data interface{}
	var total int
	if len(params) > 0 {
		data = params[0]
		if len(params) > 1 {
			if reflect.TypeOf(params[1]) == reflect.TypeOf(int(1)) {
				total = params[1].(int)
			} else if reflect.TypeOf(params[1]) == reflect.TypeOf(int64(1)) {
				total = int(params[1].(int64))
			}
		}
	}
	return models.Result{
		State: int(code),
		Msg:   code.String(),
		Data:  data,
		Total: total,
	}
}
func verifyPower(admin models.AdminAccount, code int) bool {
	for _, per := range strings.Split(admin.Permissions, ",") {
		iPer, _ := strconv.Atoi(per)
		//999是所有权限，0是不需要权限
		if iPer == 999 || code == 0 {
			return true
		} else {
			if iPer == code {
				return true
			}
		}
	}
	return false
}

func createPostForm(c *baseController, items []map[string]interface{}) {
	htmlFormatter := "<div class=\"form-group input-group\"><span class=\"input-group-addon\">%s</span>%s</div>"
	inputFormatter := "<input class=\"form-control\" type=\"%s\" name=\"%s\" value=\"%s\" placeholder=\"%s\" %s %s/>"
	sliderFormatter := "<input class=\"form-control\" type=\"slider\" id=\"%s\" name=\"%s\" data-slider-min=\"%s\" data-slider-max=\"%s\" data-slider-step=\"%s\" data-slider-value=\"%s\" data-slider-can-min=\"%s\"/><span class=\"input-group-addon\" style=\"width: 65px\"><span id=\"%s_val\">%s</span></span>"
	selectFormatter := "<select class=\"form-control\" id=\"%s\" name=\"%s\">%s</select>"
	optionFormatter := "<option value=\"%d\"%s>%s</option>"
	checkboxFormatter := "<div class=\"col-sm-6 col-md-4 col-lg-4\"><label class=\"checkbox-inline\"><input type=\"checkbox\" name=\"%s\" value=\"%d\"%s>%s</label></div>"
	radioFormatter := "<input class=\"form-control\" type=\"radio\" name=\"%s\" value=\"%s\"%s>%s"
	var html template.HTML
	for _, v := range items {
		if v["id"] == "title" {
			c.Data["title"] = v["value"]
		} else if v["id"] == "url" {
			c.Data["url"] = v["value"]
		} else {
			switch v["type"] {
			case "text", "number", "password":
				var required, readonly string
				if v["required"] == "true" {
					required = "required"
				}
				if v["readonly"] == "true" {
					readonly = "readonly"
				}
				input := template.HTML(fmt.Sprintf(inputFormatter, v["type"], v["id"], v["value"], v["name"], required, readonly))
				html += template.HTML(fmt.Sprintf(htmlFormatter, v["name"], input))
				break
			case "hidden":
				html += template.HTML(fmt.Sprintf("<input type=\"hidden\" name=\"%s\" value=\"%s\">", v["id"], v["value"]))
			case "content":
				html += template.HTML(v["value"].(string))
				break
			case "slider":
				c.Data["hasSlider"] = true
				input := template.HTML(fmt.Sprintf(sliderFormatter, v["id"], v["id"], v["min"], v["max"], v["step"], v["value"], v["can-min"], v["id"], v["value"]))
				html += template.HTML(fmt.Sprintf(htmlFormatter, v["name"], input))
				break
			case "select":
				optionHtml := ""
				options := v["value"].([]map[string]int)
				for _, op := range options {
					for key, value := range op {
						if value == v["selected"] {
							optionHtml += fmt.Sprintf(optionFormatter, value, "selected", key)
						} else {
							optionHtml += fmt.Sprintf(optionFormatter, value, "", key)
						}
					}
				}
				input := template.HTML(fmt.Sprintf(selectFormatter, v["id"], v["id"], optionHtml))
				html += template.HTML(fmt.Sprintf(htmlFormatter, v["name"], input))
				break
			case "checkbox":
				checkboxHtml := ""
				checkbox_items := v["value"].(map[string]map[string]int)
				var ind = 0
				var groups []string
				for group := range checkbox_items {
					groups = append(groups, group)
				}
				sort.Strings(groups)
				for _, group := range groups {
					cls := "form-control-group"
					groupHtml := "<div class=\"" + cls + "\">"
					groupHtml += "<p>" + group + "</p>"
					itemsHtml := ""
					var items []string
					for item := range checkbox_items[group] {
						items = append(items, item)
					}
					sort.Strings(items)
					for _, item := range items {
						checkboxItem := ""
						checkeds := strings.Split(v["checked"].(string), ",")
						//fmt.Printf(fmt.Sprintf("%d,", value))
						item_checked := false
						for _, checked := range checkeds {
							if checked == fmt.Sprintf("%d", checkbox_items[group][item]) {
								item_checked = true
							}
						}
						if item_checked {
							checkboxItem = fmt.Sprintf(checkboxFormatter, v["id"], checkbox_items[group][item], "checked", item)
						} else {
							checkboxItem = fmt.Sprintf(checkboxFormatter, v["id"], checkbox_items[group][item], "", item)
						}
						groupHtml += checkboxItem
					}
					groupHtml += itemsHtml
					groupHtml += "</div>"
					checkboxHtml += groupHtml
					ind++
				}
				tmpHtml := fmt.Sprintf(htmlFormatter, v["name"], checkboxHtml)
				frist := strings.Index(tmpHtml, "form-control-group")
				tmpHtml = tmpHtml[:frist] + "form-control-group form-control-group-top" + tmpHtml[frist+18:]
				last := strings.LastIndex(tmpHtml, "form-control-group")
				tmpHtml = tmpHtml[:last] + "form-control-group form-control-group-bottom" + tmpHtml[last+18:]

				html += template.HTML(tmpHtml)
				break
			case "radio":
				radios := ""
				input := template.HTML(fmt.Sprintf(radioFormatter, v["id"], v["id"], radios))
				html += template.HTML(fmt.Sprintf(htmlFormatter, v["name"], input))
				break
			}
		}
	}
	c.Data["form"] = html
	return
}
