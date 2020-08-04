// SPDX-License-Identifier: MIT

package locale

func init() {
	setMessages("cmn-Hant", cmnHant)
	setMessages("cmn-Hans", cmnHans)
}

// 各个语言需要翻译的所有字符串
//
// 有关命令行的相关翻译项，其第一行数据会被提取出来同时作为官网的翻译数据，
// 需要注释其第一行必须得是一个完整的句子。
// 所有以 CmdXxUsage 的都是子命令的说明语句。
//
const (
	// 与 flag 包相关的处理
	CmdUsage       = "%s 是一个 RESTful API 文档生成工具\n"
	CmdUsageFooter = `详细文档可访问官网 %s
源码以 MIT 许可发布于 %s
`
	CmdUsageOptions  = "选项："
	CmdUsageCommands = "子命令："
	CmdHelpUsage     = "显示帮助信息\n"
	CmdVersionUsage  = "显示版本信息\n"
	CmdLangUsage     = "显示所有支持的语言\n"
	CmdLocaleUsage   = "显示所有支持的本地化内容\n"
	CmdDetectUsage   = "根据目录下的内容生成配置文件\n"
	CmdSyntaxUsage   = "测试语法的正确性\n"
	CmdMockUsage     = `启用 mock 服务

mock 服务会根据接口定义检测用户提交的数据是否合法，并生成随机的数据返回给用户。
对于数据只作检测是否合规，但是无法理解其内容，比如提交地址中添加了 size=20，
只会检测 20 的类型是否符合 size 的要求，但是不会只返回给用户 20 条数据。
`
	CmdBuildUsage  = "生成文档内容\n"
	CmdStaticUsage = "启用静态文件服务\n"
	CmdLSPUsage    = "启动 language server protocol 服务\n"
	Version        = "版本：%s\n文档：%s\nLSP：%s\nopenapi：%s\nGo：%s"
	CmdNotFound    = "子命令 %s 未找到\n"

	FlagSyntaxDirUsage         = "以 `URI` 形式表示测试项目地址"
	FlagBuildDirUsage          = "以 `URI` 形式表示的项目地址"
	FlagMockPortUsage          = "指定 mock 服务的端口号"
	FlagMockServersUsage       = "指定 mock 服务时，文档中 server 变量对应的路由前缀"
	FlagMockIndentUsage        = "指定缩进内容"
	FlagMockSliceSizeUsage     = "生成数组大小的范围，格式为 [min,max]。"
	FlagMockNumSliceUsage      = "生成数值类型的数据时的数值范围，格式为 [min,max]。"
	FlagMockNumFloatUsage      = "生成的数值是否允许有浮点数存在"
	FlagMockPathUsage          = "指定文档的 `URI` 格式路径，根据此文档的内容生成 mock 数据。"
	FlagMockStringSizeUsage    = "生成字符串类型数据时字符串的长度范围，格式为 [min,max]。"
	FlagMockStringAlphaUsage   = "生成的字符串中允许出现的字符"
	FlagMockUsernameSizeUsage  = "生成邮箱地址时，用户名的长度范围，格式为 [min,max]。"
	FlagMockEmailDomainsUsage  = "生成邮箱地址时所可用的域名列表，多个用半角逗号分隔。"
	FlagMockURLDomainsUsage    = "生成 URL 地址时所可用的域名列表，多个用半角逗号分隔。"
	FlagMockImagePrefixUsage   = "生成图片类型数据的基地址"
	FlagMockDateRangeUsage     = "生成可用的日期范围，格式为 [start,end]，start 和 end 均为 RFC3339 格式。"
	FlagDetectRecursiveUsage   = "detect 子命令是否检测子目录的值"
	FlagDetectDirUsage         = "以 `URI` 形式表示检测项目地址"
	FlagDetectWrite            = "是否将配置内容写入文件，如果为 true，会将配置内容写入检测目录下的 .apidoc.yaml 文件。"
	FlagStaticPortUsage        = "指定 static 服务的端口号"
	FlagStaticDocsUsage        = "指定 static 服务静态文件所在的 `URI`"
	FlagStaticStylesheetUsage  = "指定 static 是否只启用样式文件内容"
	FlagStaticContentTypeUsage = "指定 static 的 content-type 值，不指定，则根据扩展名自动获取"
	FlagStaticURLUsage         = "指定 static 服务中文档的输出地址"
	FlagStaticPathUsage        = "指定 static 服务 `URI` 格式的文档路径，如果未指定，则不生成相关的文档内容。"
	FlagLSPPortUsage           = "指定 language server protocol 服务的端口号"
	FlagLSPModeUsage           = "指定 language server protocol 的运行方式，可以是 stdio、tcp、unix、ipc 和 udp"
	FlagLSPHeaderUsage         = "指定 language server protocol 传递内容是否带报头信息"
	FlagVersionKindUsage       = "只显示该类型的版本号，可以是 apidoc、doc、lsp、openapi 和 all"

	VersionInCompatible = "当前程序与配置文件中指定的版本号不兼容"
	Complete            = "完成！文档保存在：%s，总用时：%v"
	ConfigWriteSuccess  = "配置内容成功写入 %s"
	TestSuccess         = "语法没有问题！"
	LangID              = "ID"
	LangName            = "名称"
	LangExts            = "扩展名"
	LoadAPI             = "加载 API：%s %s"
	RequestAPI          = "访问 API：%s %s"
	DeprecatedWarn      = "%s %s 将于 %s 被废弃"
	GeneratorBy         = "当前文档由 %s 生成"
	ServerStart         = "服务启动，可通过 %s 访问"
	RequestRPC          = "访问 RPC：%s"
	UnimplementedRPC    = "未实现该 RPC 服务 %s"
	PackFileHeader      = "文档由 %s 自动生成，请勿手动修改！"
	CloseLSPFolder      = "关闭项目 %s"

	// 文档树中各个字段的介绍
	UsageAPIDoc              = "usage-apidoc"
	UsageAPIDocAPIDoc        = "usage-apidoc-apidoc"
	UsageAPIDocLang          = "usage-apidoc-lang"
	UsageAPIDocLogo          = "usage-apidoc-logo"
	UsageAPIDocCreated       = "usage-apidoc-created"
	UsageAPIDocVersion       = "usage-apidoc-version"
	UsageAPIDocTitle         = "usage-apidoc-title"
	UsageAPIDocDescription   = "usage-apidoc-description"
	UsageAPIDocContact       = "usage-apidoc-contact"
	UsageAPIDocLicense       = "usage-apidoc-license"
	UsageAPIDocTags          = "usage-apidoc-tags"
	UsageAPIDocServers       = "usage-apidoc-servers"
	UsageAPIDocAPIs          = "usage-apidoc-apis"
	UsageAPIDocHeaders       = "usage-apidoc-headers"
	UsageAPIDocResponses     = "usage-apidoc-responses"
	UsageAPIDocMimetypes     = "usage-apidoc-mimetypes"
	UsageAPIDocXMLNamespaces = "usage-apidoc-xml-namespaces"

	UsageXMLNamespace       = "usage-xml-namespace"
	UsageXMLNamespacePrefix = "usage-xml-namespace-prefix"
	UsageXMLNamespaceURN    = "usage-xml-namespace-urn"

	UsageAPI            = "usage-api"
	UsageAPIVersion     = "usage-api-version"
	UsageAPIMethod      = "usage-api-method"
	UsageAPIID          = "usage-api-id"
	UsageAPIPath        = "usage-api-path"
	UsageAPISummary     = "usage-api-summary"
	UsageAPIDescription = "usage-api-description"
	UsageAPIRequests    = "usage-api-requests"
	UsageAPIResponses   = "usage-api-responses"
	UsageAPICallback    = "usage-api-callback"
	UsageAPIDeprecated  = "usage-api-deprecated"
	UsageAPIHeaders     = "usage-api-headers"
	UsageAPITags        = "usage-api-tags"
	UsageAPIServers     = "usage-api-servers"

	UsageLink     = "usage-link"
	UsageLinkText = "usage-link-text"
	UsageLinkURL  = "usage-link-url"

	UsageContact      = "usage-contact"
	UsageContactName  = "usage-contact-name"
	UsageContactURL   = "usage-contact-url"
	UsageContactEmail = "usage-contact-email"

	UsageCallback            = "usage-callback"
	UsageCallbackMethod      = "usage-callback-method"
	UsageCallbackPath        = "usage-callback-path"
	UsageCallbackSummary     = "usage-callback-summary"
	UsageCallbackDeprecated  = "usage-callback-deprecated"
	UsageCallbackDescription = "usage-callback-description"
	UsageCallbackResponses   = "usage-callback-responses"
	UsageCallbackRequests    = "usage-callback-requests"
	UsageCallbackHeaders     = "usage-callback-headers"

	UsageEnum            = "usage-enum"
	UsageEnumDeprecated  = "usage-enum-deprecated"
	UsageEnumValue       = "usage-enum-value"
	UsageEnumSummary     = "usage-enum-summary"
	UsageEnumDescription = "usage-enum-description"

	UsageExample         = "usage-example"
	UsageExampleMimetype = "usage-example-mimetype"
	UsageExampleSummary  = "usage-example-summary"
	UsageExampleContent  = "usage-example-content"

	UsageParam            = "usage-param"
	UsageParamName        = "usage-param-name"
	UsageParamType        = "usage-param-type"
	UsageParamDeprecated  = "usage-param-deprecated"
	UsageParamDefault     = "usage-param-default"
	UsageParamOptional    = "usage-param-optional"
	UsageParamArray       = "usage-param-array"
	UsageParamItems       = "usage-param-items"
	UsageParamSummary     = "usage-param-summary"
	UsageParamEnums       = "usage-param-enums"
	UsageParamDescription = "usage-param-description"
	UsageParamArrayStyle  = "usage-param-array-style"

	UsagePath        = "usage-path"
	UsagePathPath    = "usage-path-path"
	UsagePathParams  = "usage-path-params"
	UsagePathQueries = "usage-path-queries"

	UsageRequest            = "usage-request"
	UsageRequestName        = "usage-request-name"
	UsageRequestType        = "usage-request-type"
	UsageRequestDeprecated  = "usage-request-deprecated"
	UsageRequestArray       = "usage-request-array"
	UsageRequestItems       = "usage-request-items"
	UsageRequestSummary     = "usage-request-summary"
	UsageRequestStatus      = "usage-request-status"
	UsageRequestEnums       = "usage-request-enums"
	UsageRequestDescription = "usage-request-description"
	UsageRequestMimetype    = "usage-request-mimetype"
	UsageRequestExamples    = "usage-request-examples"
	UsageRequestHeaders     = "usage-request-headers"

	UsageRichtext     = "usage-richtext"
	UsageRichtextType = "usage-richtext-type"
	UsageRichtextText = "usage-richtext-text"

	UsageTag           = "usage-tag"
	UsageTagName       = "usage-tag-name"
	UsageTagTitle      = "usage-tag-title"
	UsageTagDeprecated = "usage-tag-deprecated"

	UsageServer            = "usage-server"
	UsageServerName        = "usage-server-name"
	UsageServerTitle       = "usage-server-title"
	UsageServerURL         = "usage-server-url"
	UsageServerDeprecated  = "usage-server-deprecated"
	UsageServerSummary     = "usage-server-summary"
	UsageServerDescription = "usage-server-description"

	UsageXMLAttr    = "usage-xml-attr"
	UsageXMLExtract = "usage-xml-extract"
	UsageXMLCData   = "usage-xml-cdata"
	UsageXMLPrefix  = "usage-xml-prefix"
	UsageXMLWrapped = "usage-xml-wrapped"

	// 基本类型
	UsageString  = "usage-string"
	UsageNumber  = "usage-number"
	UsageBool    = "usage-bool"
	UsageVersion = "usage-version"
	UsageDate    = "usage-date"
	UsageType    = "usage-type"

	// 以下是有关 build.Config 的字段说明
	UsageConfigVersion               = "usage-config-version"
	UsageConfigInputs                = "usage-config-inputs"
	UsageConfigInputsLang            = "usage-config-inputs.lang"
	UsageConfigInputsDir             = "usage-config-inputs.dir"
	UsageConfigInputsExts            = "usage-config-inputs.exts"
	UsageConfigInputsRecursive       = "usage-config-inputs.recursive"
	UsageConfigInputsEncoding        = "usage-config-inputs.encoding"
	UsageConfigOutput                = "usage-config-output"
	UsageConfigOutputType            = "usage-config-output.type"
	UsageConfigOutputPath            = "usage-config-output.path"
	UsageConfigOutputTags            = "usage-config-output.tags"
	UsageConfigOutputStyle           = "usage-config-output.style"
	UsageConfigOutputNamespace       = "usage-config-output.namespace"
	UsageConfigOutputNamespacePrefix = "usage-config-output.namespace-prefix"

	// 错误信息，可能在地方用到
	ErrInvalidUTF8Character      = "无效的 UTF8 字符"
	ErrInvalidXML                = "无效的 XML 文档"
	ErrMultipleRootTag           = "存在多个根标签"
	ErrIsNotAPIDoc               = "并非有效的 apidoc 的文档格式"
	ErrInvalidContentTypeCharset = "报头 ContentType 中指定的字符集无效 "
	ErrInvalidContentLength      = "报头 ContentLength 无效"
	ErrBodyIsEmpty               = "请求的报文为空"
	ErrInvalidHeaderFormat       = "无效的报头格式"
	ErrRequired                  = "不能为空"
	ErrInvalidFormat             = "格式不正确"
	ErrDirNotExists              = "目录不存在"
	ErrNotFoundEndFlag           = "找不到结束符号"
	ErrNotFoundEndTag            = "找不到结束标签"
	ErrNotFoundSupportedLang     = "该目录下没有支持的语言文件"
	ErrDirIsEmpty                = "目录下没有需要解析的文件"
	ErrInvalidValue              = "无效的值"
	ErrInvalidTag                = "无效的标签"
	ErrPathNotMatchParams        = "地址参数不匹配"
	ErrDuplicateValue            = "重复的值"
	ErrMessage                   = "%s 位于 %s"
	ErrNotFound                  = "未找到该值"
	ErrReadRemoteFile            = "读取远程文件 %s 时返回状态码 %d"
	ErrServerNotInitialized      = "服务未初始化"
	ErrInvalidLSPState           = "无效的 LSP 状态"
	ErrInvalidURIScheme          = "无效的 URI 协议"
	ErrFileNotFound              = "未找到文件 %s"

	// logs
	InfoPrefix    = "[INFO] "
	WarnPrefix    = "[WARN] "
	ErrorPrefix   = "[ERRO] "
	SuccessPrefix = "[SUCC] "
)
