## 使用示例

以 `user_demo` 为例：

```go
// 查询：多条 + 分页 
page := 2
num := 2
offset := (page - 1) * num

user, err = user_demo_repo.NewQueryBuilder().
    WhereIdNotIn([]int32{1, 2, 3}).
    WhereUserName(db_repo.EqualPredicate, "tom").
    Limit(num).
    Offset(offset).
    QueryAll(u.db.GetDbR().WithContext(ctx.RequestContext()))

// 查询：总数
count, err := user_demo_repo.NewQueryBuilder().
    WhereIdNotIn([]int32{1, 2, 3}).
    WhereUserName(db_repo.EqualPredicate, "tom").
    Count(u.db.GetDbR().WithContext(ctx.RequestContext()))

// 查询：单条
user, err = user_demo_repo.NewQueryBuilder().
    WhereUserName(db_repo.EqualPredicate, "tom").
    QueryOne(u.db.GetDbR().WithContext(ctx.RequestContext()))

// 创建
model := user_demo_repo.NewModel()
model.UserName = user.UserName
model.NickName = user.NickName
model.Mobile = user.Mobile

id, err = model.Create(u.db.GetDbW().WithContext(ctx.RequestContext()))

```