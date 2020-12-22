# dedao-dl

> 用 go 写的一个 《得到》 APP 课程下载工具，使用 cookie 登录后，可在终端查看已购买的课程，听书书架，电子书架，锦囊等



## 特性

* 可查看购买的课程，课程文章内容
* 可查看听书书架，电子书架列表
* 可查看已购买的锦囊
* 课程可生成PDF，也可生成 mp3 文件

## 安装

### 安装依赖

* chromedp
  > 生成 PDF 需要借助 [chromedp](https://github.com/chromedp/chromedp), 该功能必须安装 [Google-Chrome](https://www.google.cn/intl/zh-CN/chrome/) 浏览器
* ffmpeg
  > 音频需要借助 [ffmpeg](https://ffmpeg.org/) 合成
  >
  >说明: 代码里集成了 u3m8 下载器，会导致编译后程序体积增大 7MB 左右，如果不想安装 ffmpeg， 请打开相关代码的注释即可。

### 使用 `go get` 安装

使用如下命令安装：

`go get github.com/yann0917/dedao-dl`

## 使用方法

* 电脑端登录 [得到](https://www.dedao.cn)，以便生成 cookie ☆☆☆☆☆
* 电脑端登录 [得到](https://www.dedao.cn)，以便生成 cookie ☆☆☆☆☆
* 电脑端登录 [得到](https://www.dedao.cn)，以便生成 cookie ☆☆☆☆☆

`dedao-dl -h` 可查看帮助说明，每个命令都有 `-h` 参数可查看该命令的用法

```bash
Usage:
  dedao-dl [command]

Available Commands:
  ace         获取我的锦囊
  article     获取文章详情
  cat         获取课程分类
  course      获取我购买过课程
  dl          下载已购买课程，并转换成 PDF & 音频
  ebook       获取我的电子书架
  help        Help about any command
  login       登录得到 pc 端 https://www.dedao.cn
  odob        获取我的听书书架
  who         查看当前登录的用户
```

`dedao-dl login` 登录，不带任何参数的情况下，默认使用 [`go-rod/rod`](https://github.com/go-rod/rod) 从浏览器获取cookie

`dedao-dl cat` 获取课程分类

```text
+---+----------+------+----------+
| # |   名称   | 统计 | 分类标签 |
+---+----------+------+----------+
| 0 | 全部     | 1696 | all      |
| 1 | 课程     |   64 | bauhinia |
| 2 | 听书书架 | 1407 | odob     |
| 3 | 电子书架 |  210 | ebook    |
| 4 | 锦囊     |   15 | compass  |
+---+----------+------+----------+
```

`dedao-dl course` 获取我购买过课程

```text
+----+-----+----------------------------+---------------------------------------+---------------------+--------+----------+
| #  | ID  |          课程名称          |                 作者                  |      购买日期       |  价格  | 学习进度 |
+----+-----+----------------------------+---------------------------------------+---------------------+--------+----------+
|  0 |  51 | 张潇雨·个人投资课          | 张潇雨·对冲基金管理人                 | 2020-08-17 09:59:56 | 199.00 |      100 |
|  1 | 560 | 吴军·硅谷来信³（年度日更） | 吴军·计算机科学家                     | 2020-11-05 09:28:23 | 299.00 |       14 |
|  2 | 571 | 年度得到·何帆中国经济报告  | 何帆·著名经济学者                     | 2020-12-10 07:22:28 |  99.00 |       36 |
|  3 |  21 | 古典·超级个体              | 古典·资深生涯规划师                   | 2019-09-23 17:11:02 | 199.00 |       84 |
|  4 |  79 | 老喻的人生算法课           | 老喻·思考者                           | 2019-09-23 17:13:51 |  99.00 |      100 |
|  5 |  84 | 给忙碌者的女性健康课       | 常青·协和医学院博士                   | 2020-10-27 21:30:55 |  19.90 |      100 |
|  6 |  16 | 陈海贤·自我发展心理学      | 陈海贤·心理学博士                     | 2019-09-23 17:13:46 |  99.00 |      100 |
|  7 | 544 | 听书番外篇                 | 阿狮·每天听本书负责人                 | 2020-08-18 15:43:44 |   0.00 |      100 |
|  8 | 530 | 跟李松蔚学心理咨询         | 李松蔚·心理咨询师                     | 2020-07-01 11:01:44 |  99.00 |      100 |
|  9 | 484 | 王立铭·抑郁症医学课        | 王立铭·浙江大学生命科学研究院教授     | 2020-10-27 13:15:36 |  39.90 |      100 |
| 10 |  82 | 陈海贤·亲密关系30讲        | 陈海贤·心理学博士                     | 2019-11-05 00:02:21 |  99.00 |      100 |
+----+-----+----------------------------+---------------------------------------+---------------------+--------+----------+
```

`dedao-dl course -i xxx` 查看课程信息

```text
专栏名称：年度得到·何帆中国经济报告
专栏作者：何帆·著名经济学者
更新进度：10/26
课程亮点：1.这份报告和你想象中的不一样。何帆老师会告诉你，比起“新冠疫情”“中美博弈”，2020-2021，你更应该关注“本土时代”的到来。在这个时代，中国经济会更强调“防守”，更强调“向内寻找解决方案”；

2.“本土时代的生存策略”是这份报告要交付的方法论，包括企业如何转型、如何获得竞争优势，也包括个人如何打造新型的关系网；

3.这份报告还会带你看中国经济出现的“新物种”，让你了解宏观趋势；

4.这份报告的初心值得你知晓。何帆老师要用30年时间，每年写一份中国经济报告，记录中国从2019-2049年的变化。这是他的第三份报告。加入课程，见证这个注定伟大的传统；

5.你可以在课程中向何帆老师建言，哪些城市、产业、公司值得他研究。你的建议很有可能会出现在明年的《何帆·中国经济报告》中。

+---+------+-----------------------------------+---------------------+--------------+
| # |  ID  |               章节                |      更新时间       | 是否更新完成 |
+---+------+-----------------------------------+---------------------+--------------+
| 0 | 1661 | 发刊词(1讲)                       | 2020-12-17 13:55:24 | ✔            |
| 1 | 1654 | 模块一：本土时代(4讲)             | 2020-12-17 13:55:28 | ✔            |
| 2 | 1655 | 模块二：网中网(4讲)               | 2020-12-20 22:06:06 | ✔            |
| 3 | 1656 | 模块三：“变形金刚”策略(已更新1讲) | 2020-12-21 00:00:04 | ❌           |
| 4 | 1657 | 模块四：“领跑员”策略              | 2020-12-08 18:15:34 | ❌           |
| 5 | 1658 | 模块五：“自己人”策略              | 2020-12-08 18:16:13 | ❌           |
| 6 | 1659 | 模块六：成长的烦恼                | 2020-12-08 18:17:21 | ❌           |
| 7 | 1660 | 模块七：总结                      | 2020-12-08 18:17:45 | ❌           |
+---+------+-----------------------------------+---------------------+--------------+
```

## 参考

* [geektime-dl](https://github.com/mmzou/geektime-dl)
* [annie](https://github.com/iawia002/annie)
* [m3u8](https://github.com/oopsguy/m3u8)

仅供个人学习使用，请尊重版权，内容版权均为得到所有，请勿传播内容！！！

仅供个人学习使用，请尊重版权，内容版权均为得到所有，请勿传播内容！！！

仅供个人学习使用，请尊重版权，内容版权均为得到所有，请勿传播内容！！！

---
