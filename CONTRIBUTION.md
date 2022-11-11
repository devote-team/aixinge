# 贡献手册

本文为想要参与 AiXinGe 开源项目的指导手册，欢迎大家积极参与贡献代码。

## 贡献方式

- 提交 PullRequest：包括但不限于针对项目提供新特性、修复缺陷、完善注释、修正拼写问题等方式。
- 创建 Issue：为项目进行测验，发现问题并提出问题
- 参与讨论：欢迎加群或者联系作者进行讨论

## 代码规范

- 本项目集成了 editorconfig，相关的缩进配置都在里面，有不清楚的可以查询。
- 代码编写的时候尽可能完善注释，如果可以，最好是英文注释。
- 命名要清晰，尽可能做到见名知意。

## Git 提交规范

为了方便管理，我们的 Git
提交遵循 [AngularJS Git 提交规范](https://docs.google.com/document/d/1QrDFcIiPjSLDn3EL15IJygNPiHORgU1_OOAqWjiDU5Y/edit)
，这是一个相对标准并受到大部分人认可的 Commit 模板，主要说明如下：

一次 Git 提交信息格式类似于：

```
type(scope): short description

long description
```

其中 type 主要有以下几种类型：

- feat：新功能相关改动
- fix：缺陷相关的修复(如果有 Issue 编号请带上)
- docs：文档相关变化
- style：代码格式调整(不影响代码运行的变动)
- refactor：代码的重构或优化(既不增加新功能，也不是修复bug)
- perf：性能优化相关变动
- test：测试文件相关变动
- ci：持续集成相关文件的变动
- chore：其他文件的变动（不涉及源码和测试源码）
- revert：回退至某一次提交

而 scope 则表示了当前改动的范围，short description 为本次提交的短描述，long description 为本次提交的长描述（长描述非必选，在有必要的情况下进行填写）

故此，一个正确的提交规范类似于：

```
docs(md): add contribution

add project contribution guide
```

如果有不清楚的地方，也可以搜索相关文件或教程，查看更多的说明。

### 插件

Jetbrains 系列 IDE 中，有插件可以帮助我们快捷添加满足规范的 Commit Message。

**插件地址：**[git-commit-template](https://plugins.jetbrains.com/plugin/9861-git-commit-template)

**使用方法：**
![](https://gitee.com/aixinge/aixinge/raw/master/wiki/img/Git-Commit-Template-Open.jpg)
![](https://gitee.com/aixinge/aixinge/raw/master/wiki/img/Git-Commit-Template-Use.jpg)

除了插件以外，大家还可以搜索到更多的工具，在此就不赘述了，大家感兴趣可以自行了解。
