<h1 align="center"> identicon-avatar-go </h1>

[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu)
[![LICENSE](https://img.shields.io/badge/license-Anti%20996-blue.svg)](https://github.com/996icu/996.ICU/blob/master/LICENSE)


<p align="center">
identicon 根据一个字符串(可以是用户的ID或者IP)的哈希值生成大量不重复的头像
很多大型IT网站上可以见到，比如 Github、Sourceforge、Stackoveflow。
</p>

<p align="center">
基于go基础库
</p>

## 安装

```shell
$ go get github.com/valiner/identicon-avatar-go
```

## 使用
```go
//生成本地图片
identicon.SaveAvatar("sdp", 125, "test.jpg")

//图片base64 uri
identicon.GetAvatarDataUri("sdp", 125)
```

## License

MIT