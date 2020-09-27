# sql注入脚本常用的算法

(持续更新)

二分法

用法

```
sqli.FindByDichotomy(f)


//f是自定义的函数，
//给定一个当前值 burp，能够当target > burp时返回真，否则返回假
//target一般就是ascii(substr((select databse()),1,1))
func f(burp int)bool{
	//里面就是http请求以及判断的函数（时间盲注或者布尔盲注）
}
```

