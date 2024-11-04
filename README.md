# Git Learn

this is in read.md file <br/>
**learning start** <br/>

---
~~英文能力有限之后就暂时使用中文来写吧~~ <br/>
~~先把git 弄明白再写markdown~~ <br/>
~~竟然有人这个时候还不熟git和markdown是谁就不说了~~ <br/>

## Start Basic

* 建立远程仓库 复制ssh地址,后面要用,不需要添加 read.md 文件
    * 本地创建和仓库同名的文件夹
    * `git init` 初始化本地仓库，这时候会多出俩一个 .git 文件，不要动它
    * `git remote add origin copy` 将远程仓库在本地起名为 origin
    * `git add .` 将当前所有修改添加到暂存区
    * `git commit -m "your words"` 将当前暂存区的内容提交到本地仓库，并添加描述
    * `git push -u origin master:main` 第一次提交
    * `git push origin master:main` 之后所有的提交
    * `git status` 查看本地状态

* 分支操作
    * `git branch my_branch_name` 创建分支
    * `git branch -v` 查看分支
    * `git checkout name` 切换到指定分支
    * `git merge name` 把指定的name分支合并到当前的分支上

* 团队开发
    * 首先将自己本地代码提交到本地仓库中
    * `git pull origin my_branch_name` 合并远程仓库的修改到本地的指定分支中
    * 如果发生了合并冲突，git 会使用特殊的标记 **<<<<<<**   **======**    **>>>>>>** 来标记冲突的部分
    * 需要手动解决冲突的部分，之后使用 git add . 标记冲突已解决
    * 最后本地仓库提交到远程仓库

* 总结：
    1. 创建一个分支 ，并切换进入该分支
    2. 在该分支内进行开发，建立一个文件库
    3. 。。。。。。
    4. 应该可以直接在本地的主分支开发，远程再开一个分支 (? . ?) 

* 感悟：大概了解了,之前个人开发都不怎么用 git，上传文件都是直接 upload ... 
* 总结：~~我是笨笨~~

---
## Start MarkDown


1. **标题**
    * 用 **#** 表示，可以表示1-6级标题，随 **#** 个数递增，一级标题最大，六级最小
    * 效果如下：

```markdown
# 一级标题
## 二级标题
### 三级标题
#### 四级标题
##### 五级标题
###### 六级标题
```


2. **字体**
    * 星号和下划线都可以，单是斜体，双是粗体，三是斜粗体
    * 效果如下：

```markdown
*这是斜体*        <br/>
_这是斜体_        <br/>
**这是粗体**      <br/>
__这是粗体__      <br/>
***这是斜粗体***  <br/>
___这是斜粗体___  <br/>
```


3. 换行
    * 一句话后面敲两个空格
    * 两句话之间加一个空行
    * 一行文字中间加<br/>表示换行


4. 引用
    * 通过符号 **>** 来实现， **>** 后面的空格可有可无
    * 效果如下：

```markdown
> 这是一个引用
>> 这是引用的引用
>>> 这是引用的引用的引用
```


5. 链接
    * 实现方式:  [链接名称](链接地址)  
    * 效果如下：

```markdown
[百度](https://www.baidu.com)
```


6. 图片
    * 实现方式:  ![图片描述](图片地址)
    * 效果如下：

```markdown
![兔斯基](https://www.bing.com/images/search?view=detailV2&ccid=LYCxLaGl&id=EF628405A669E5DF89F44CC0F75A59509C65BB4F&thid=OIP.LYCxLaGlvbCayPsLrkNh3gHaEo&mediaurl=https%3a%2f%2fth.bing.com%2fth%2fid%2fR.2d80b12da1a5bdb09ac8fb0bae4361de%3frik%3dT7tlnFBZWvfATA%26riu%3dhttp%253a%252f%252fb.zol-img.com.cn%252fdesk%252fbizhi%252fimage%252f1%252f960x600%252f1348730411542.jpg%26ehk%3dRE8fis3mdsbNlYL%252fGJfjlhfNZPCqoWVVIP1laYV02Gw%253d%26risl%3d%26pid%3dImgRaw%26r%3d0&exph=600&expw=960&q=%e5%85%94%e6%96%af%e5%9f%ba&simid=607996125399494177&FORM=IRPRST&ck=1B94F09A725611A0D224A1B608C32E75&selectedIndex=2&itb=0)
```
