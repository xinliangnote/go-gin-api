## 与 UrlSign 对应的 PHP 加密算法

```php
// 对 params key 进行排序
ksort($params);

// 对 sortParams 进行 Encode
$sortParamsEncode = http_build_query($params);

// 加密字符串规则 path + method + sortParamsEncode + secret
$encryptStr = $path . $method . $sortParamsEncode . $secret

// 对加密字符串进行 md5
$md5Str = md5($encryptStr);

// 对 md5Str 进行 base64 encode
$tokenString = base64_encode($md5Str);

echo $tokenString;
```