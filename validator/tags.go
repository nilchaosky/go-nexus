package validator

import (
	"fmt"
)

var (
	tagMap map[string]string
)

// init 包加载时初始化标签映射表
// 记录所有validator库支持的验证标签及其格式化字符串模板
func init() {
	// 预分配容量，包含所有常用标签
	tagMap = make(map[string]string, 100)

	// 手动添加所有常用标签的格式化模板（中文）
	setCustomTemplates()
}

// getDefaultTemplate 获取标签的默认格式化模板
func getDefaultTemplate(tag string) string {
	return fmt.Sprintf("字段未通过 '%s' 验证", tag)
}

// setCustomTemplates 设置自定义的格式化模板（中文）
func setCustomTemplates() {
	customTemplates := map[string]string{
		"required":                      "{0}为必填项",
		"required_if":                   "当{0}存在时，{1}为必填项",
		"required_unless":               "当{0}不存在时，{1}为必填项",
		"required_with":                 "当{0}存在时，{1}为必填项",
		"required_with_all":             "当所有指定字段存在时，{0}为必填项",
		"required_without":              "当{0}不存在时，{1}为必填项",
		"required_without_all":          "当所有指定字段不存在时，{0}为必填项",
		"excluded_if":                   "当{0}存在时，{1}不能存在",
		"excluded_unless":               "当{0}不存在时，{1}不能存在",
		"excluded_with":                 "当{0}存在时，{1}不能存在",
		"excluded_without":              "当{0}不存在时，{1}不能存在",
		"isdefault":                     "{0}必须为默认值",
		"len":                           "{0}长度必须为{1}",
		"min":                           "{0}最小值为{1}",
		"max":                           "{0}最大值为{1}",
		"eq":                            "{0}必须等于{1}",
		"ne":                            "{0}不能等于{1}",
		"lt":                            "{0}必须小于{1}",
		"lte":                           "{0}必须小于等于{1}",
		"gt":                            "{0}必须大于{1}",
		"gte":                           "{0}必须大于等于{1}",
		"eqfield":                       "{0}必须等于{1}字段",
		"eqcsfield":                     "{0}必须等于{1}字段（忽略大小写）",
		"necsfield":                     "{0}不能等于{1}字段（忽略大小写）",
		"gtcsfield":                     "{0}必须大于{1}字段（忽略大小写）",
		"gtecsfield":                    "{0}必须大于等于{1}字段（忽略大小写）",
		"ltcsfield":                     "{0}必须小于{1}字段（忽略大小写）",
		"ltecsfield":                    "{0}必须小于等于{1}字段（忽略大小写）",
		"nefield":                       "{0}不能等于{1}字段",
		"gtefield":                      "{0}必须大于等于{1}字段",
		"gtfield":                       "{0}必须大于{1}字段",
		"ltefield":                      "{0}必须小于等于{1}字段",
		"ltfield":                       "{0}必须小于{1}字段",
		"alpha":                         "{0}只能包含字母",
		"alphanum":                      "{0}只能包含字母和数字",
		"alphaunicode":                  "{0}只能包含Unicode字母",
		"alphanumunicode":               "{0}只能包含Unicode字母和数字",
		"numeric":                       "{0}必须为数字",
		"number":                        "{0}必须为数字",
		"hexadecimal":                   "{0}必须为十六进制",
		"hexcolor":                      "{0}必须为十六进制颜色值",
		"rgb":                           "{0}必须为RGB颜色值",
		"rgba":                          "{0}必须为RGBA颜色值",
		"hsl":                           "{0}必须为HSL颜色值",
		"hsla":                          "{0}必须为HSLA颜色值",
		"email":                         "{0}必须为有效的邮箱地址",
		"url":                           "{0}必须为有效的URL",
		"uri":                           "{0}必须为有效的URI",
		"base64":                        "{0}必须为Base64编码",
		"base64url":                     "{0}必须为Base64URL编码",
		"contains":                      "{0}必须包含{1}",
		"containsany":                   "{0}必须包含{1}中的任意字符",
		"containsrune":                  "{0}必须包含字符{1}",
		"excludes":                      "{0}不能包含{1}",
		"excludesall":                   "{0}不能包含{1}中的任意字符",
		"excludesrune":                  "{0}不能包含字符{1}",
		"startswith":                    "{0}必须以{1}开头",
		"endswith":                      "{0}必须以{1}结尾",
		"startsnotwith":                 "{0}不能以{1}开头",
		"endsnotwith":                   "{0}不能以{1}结尾",
		"isbn":                          "{0}必须为有效的ISBN",
		"isbn10":                        "{0}必须为有效的ISBN-10",
		"isbn13":                        "{0}必须为有效的ISBN-13",
		"uuid":                          "{0}必须为有效的UUID",
		"uuid3":                         "{0}必须为有效的UUID v3",
		"uuid4":                         "{0}必须为有效的UUID v4",
		"uuid5":                         "{0}必须为有效的UUID v5",
		"ulid":                          "{0}必须为有效的ULID",
		"ascii":                         "{0}只能包含ASCII字符",
		"printascii":                    "{0}只能包含可打印的ASCII字符",
		"multibyte":                     "{0}必须包含多字节字符",
		"datauri":                       "{0}必须为有效的Data URI",
		"latitude":                      "{0}必须为有效的纬度",
		"longitude":                     "{0}必须为有效的经度",
		"ssn":                           "{0}必须为有效的SSN",
		"ip":                            "{0}必须为有效的IP地址",
		"ipv4":                          "{0}必须为有效的IPv4地址",
		"ipv6":                          "{0}必须为有效的IPv6地址",
		"cidr":                          "{0}必须为有效的CIDR",
		"cidrv4":                        "{0}必须为有效的CIDRv4",
		"cidrv6":                        "{0}必须为有效的CIDRv6",
		"tcp_addr":                      "{0}必须为有效的TCP地址",
		"tcp4_addr":                     "{0}必须为有效的TCP4地址",
		"tcp6_addr":                     "{0}必须为有效的TCP6地址",
		"udp_addr":                      "{0}必须为有效的UDP地址",
		"udp4_addr":                     "{0}必须为有效的UDP4地址",
		"udp6_addr":                     "{0}必须为有效的UDP6地址",
		"ip_addr":                       "{0}必须为有效的IP地址",
		"ip4_addr":                      "{0}必须为有效的IP4地址",
		"ip6_addr":                      "{0}必须为有效的IP6地址",
		"unix_addr":                     "{0}必须为有效的Unix地址",
		"mac":                           "{0}必须为有效的MAC地址",
		"hostname":                      "{0}必须为有效的主机名",
		"hostname_rfc1123":              "{0}必须为符合RFC1123的主机名",
		"fqdn":                          "{0}必须为有效的FQDN",
		"unique":                        "{0}必须为唯一值",
		"oneof":                         "{0}必须为以下值之一：{1}",
		"json":                          "{0}必须为有效的JSON",
		"jwt":                           "{0}必须为有效的JWT",
		"lowercase":                     "{0}必须为小写",
		"uppercase":                     "{0}必须为大写",
		"datetime":                      "{0}必须为有效的日期时间",
		"timezone":                      "{0}必须为有效的时区",
		"iso3166_1_alpha2":              "{0}必须为有效的ISO3166-1 alpha-2国家代码",
		"iso3166_1_alpha3":              "{0}必须为有效的ISO3166-1 alpha-3国家代码",
		"iso3166_1_alpha_numeric":       "{0}必须为有效的ISO3166-1数字国家代码",
		"iso3166_2":                     "{0}必须为有效的ISO3166-2代码",
		"iso4217":                       "{0}必须为有效的ISO4217货币代码",
		"iso8601":                       "{0}必须为有效的ISO8601日期",
		"postcode_iso3166_alpha2":       "{0}必须为有效的邮政编码（基于ISO3166-1 alpha-2）",
		"postcode_iso3166_alpha2_field": "{0}必须为有效的邮政编码（基于字段{1}）",
		"boolean":                       "{0}必须为布尔值",
		"image":                         "{0}必须为有效的图片",
	}

	// 将所有自定义模板添加到tagMap
	for tag, template := range customTemplates {
		tagMap[tag] = template
	}
}

// getTagTemplate 获取指定标签的格式化模板
// tag 为验证标签名称
func getTagTemplate(tag string) string {
	if template, exists := tagMap[tag]; exists {
		return template
	}
	return getDefaultTemplate(tag)
}
