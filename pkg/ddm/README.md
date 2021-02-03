## DDM

动态数据掩码（Dynamic Data Masking，简称为DDM）能够防止把敏感数据暴露给未经授权的用户。

| 类型 | 要求 | 示例 | 说明
| ---- | ---- | ---- | ---- 
| 手机号 | 前 3 后 4 | 132****7986 | 定长 11 位数字
| 邮箱地址 | 前 1 后 1 | l**w@gmail.com | 仅对 @ 之前的邮箱名称进行掩码
| 姓名 | 隐姓 | *鸿章 | 将姓氏隐藏
| 密码 | 不输出 | ****** | 
| 银行卡卡号 | 前 6 后 4 | 622888******5676 | 银行卡卡号最多 19 位数字
| 身份证号 | 前 1 后 1 | 1******7 | 定长 18 位

#### 代码示例

```
// 返回值
type message struct {
	Email     ddm.Email    `json:"email"`
}

msg := new(message)
msg.Email = ddm.Email("xinliangnote@163.com")
...

```
