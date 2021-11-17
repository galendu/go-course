# 权限判定

常见的2种权限方案:
+ ABAC: Attribute-based Access Control, 基于属性的访问控制
+ RBAC: Role-based Access Control，基于角色的访问控制

keyauth支持2种权限判断方式:
+ 基于用户属性的权限控制, 用于 keyauth本身的权限控制
+ 基于Policy RBAC的权限控制, 基于 对其他子系统的权限控制, 避免子系统权限设计不好, 被提权, 破坏整个权限系统

## keyauth本身权限判定

### 用户分类

首先我们看下用户分类:
```go
// 为了防止越权, 用户可以调整的权限范围只有10已下的权限
type UserType int32

const (
	// 子账号, 无用户中心后台管理权限
	UserType_SUB UserType = 0
	// 组织管理严, 管理该域的组织结构
	UserType_ORG_ADMIN UserType = 5
	// 审计管理员, 可以查看用户中心相关配置, 相当于用户中心只读权限
	UserType_AUDIT_ADMIN UserType = 6
	// 权限管理员, 管理用户的授权策略, 比如空间管理，资源访问策略的配置
	UserType_PERM_ADMIN UserType = 7
	// 域管理员, 有该域的所有管理权限, 协作主账号进行管理
	UserType_DOMAIN_ADMIN UserType = 8
	// 主账号, 具有本域的所有权限
	UserType_PRIMARY UserType = 10
	// 超级管理员, 系统管理员, 万能的人, 不受权限系统约束
	UserType_SUPPER UserType = 15
)
```

### 权限设置

keyauth会和框架集成, 在路由上打上对应的标识, 最终由权限拦截器负责判定

下面是user模块的权限设定
```go
// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	prmary := router.ResourceRouter("primary_account")
	prmary.Allow(types.UserType_SUPPER)
	prmary.BasePath("users")
	prmary.Handle("POST", "/", h.CreatePrimayAccount)
	prmary.Handle("DELETE", "/", h.DestroyPrimaryAccount)

	ram := router.ResourceRouter("ram_account")
	ram.Allow(types.UserType_ORG_ADMIN)
	ram.BasePath("sub_users")
	ram.Handle("POST", "/", h.CreateSubAccount)
	ram.Handle("GET", "/", h.QuerySubAccount)
	ram.Handle("GET", "/:account", h.DescribeSubAccount)
	ram.Handle("PATCH", "/:account", h.PatchSubAccount)
	ram.Handle("DELETE", "/:account", h.DestroySubAccount)
	ram.BasePath("manage")
	ram.Handle("POST", "/block", h.BlockSubAccount)

	portal := router.ResourceRouter("profile")
	portal.BasePath("profile")
	portal.Handle("GET", "/", h.QueryProfile)
	portal.Handle("GET", "/domain", h.QueryDomain)
	portal.Handle("PUT", "/", h.PutProfile)
	portal.Handle("PATCH", "/", h.PatchProfile)

	dom := router.ResourceRouter("domain")
	dom.Allow(types.UserType_DOMAIN_ADMIN)
	dom.BasePath("settings/domain")
	dom.Handle("PUT", "/info", h.UpdateDomainInfo)
	dom.Handle("PUT", "/security", h.UpdateDomainSecurity)

	pass := router.ResourceRouter("password")
	pass.BasePath("password")
	pass.Handle("POST", "/", h.GeneratePassword)
	pass.Handle("PUT", "/", h.UpdatePassword)
}
```

下面是部门管理的权限设定:
```go
	r.BasePath("departments")
	r.Handle("POST", "/", h.Create).SetAllow(types.UserType_ORG_ADMIN)
	r.Handle("GET", "/", h.List)
	r.Handle("GET", "/:id", h.Get)
	r.Handle("PUT", "/:id", h.Put)
	r.Handle("PATCH", "/:id", h.Patch)
	r.Handle("GET", "/:id/subs", h.GetSub)
	r.Handle("DELETE", "/:id", h.Delete).SetAllow(types.UserType_ORG_ADMIN)
```

### 权限判定

用户认证过后, Token信息里面包含当前用户的类型, 只需要拦截, 和权限条目进行匹配就可以判定

```go
func (a *HTTPAuther) ValidatePermission(ctx context.Context, tk *token.Token, e httpb.Entry) error {
	if tk == nil {
		return exception.NewUnauthorized("validate permission need token")
	}

	// 如果是超级管理员不做权限校验, 直接放行
	if tk.UserType.IsIn(types.UserType_SUPPER) {
		a.l.Debugf("[%s] supper admin skip permission check!", tk.Account)
		return nil
	}

	// 检查是否是允许的类型
	if len(e.Allow) > 0 {
		a.l.Debugf("[%s] start check permission to keyauth ...", tk.Account)
		if !e.IsAllow(tk.UserType) {
			return exception.NewPermissionDeny("no permission, allow: %s, but current: %s", e.Allow, tk.UserType)
		}
		a.l.Debugf("[%s] permission check passed", tk.Account)
	}

	return nil
}
```

### 测试

+ 我们先使用超级管理员换取一个token
+ 使用超级管理员创建一个普通用户和一个组织管理员用户
    + 子用户01 ^K%0SFpivmUq
    + 组织管理员01 ^K%0SFpivmUq

然后测试子用户和组织管理员 能否创建用户



## RBAC


## Namespace


## Filter



## 权限对接