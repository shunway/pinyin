# pinyin


golang实现中文汉字转拼音

demo
```go
package main

import(
	"fmt"
	 "github.com/shunway/pinyin"
)

func main()  {
    str, err := pinyin.New("光荣与梦想").Split("").Mode(InitialsInCapitals).Convert()
	if err != nil {
		// 错误处理
	}else{
		fmt.Println(str)
	}

	str, err = pinyin.New("光荣与梦想").Split(" ").Mode(pinyin.WithoutTone).Convert()
	if err != nil {
		// 错误处理
    }else{
    	fmt.Println(str)
    }

	str, err = pinyin.New("光荣与梦想").Split("-").Mode(pinyin.Tone).Convert()
	if err != nil {
		// 错误处理
    }else{
    	fmt.Println(str)
    }

	str, err = pinyin.New("光荣与梦想").Convert()
	if err != nil {
		// 错误处理
    }else{
    	fmt.Println(str)
    }	
}
```

输出
```bash
GuangRongYuMengXiang
guang rong yu meng xiang
guāng-róng-yǔ-mèng-xiǎng
guang rong yu meng xiang
```

Mode 介绍

* `InitialsInCapitals`: 首字母大写, 不带音调
* `WithoutTone`: 全小写,不带音调
* `Tone`: 全小写带音调

Split 介绍

split 方法是两个汉字之间的分隔符.