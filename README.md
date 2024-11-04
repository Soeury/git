# Git Learn

this is in read.md file <br/>
**learning start** <br/>

---
~~英文能力有限之后就暂时使用中文来写吧~~  <br/>
~~先把git 弄明白再写markdown~~  <br/>
~~竟然有人这个时候还不熟git和markdown是谁就不说了~~  <br/>
~~团队开发的时候不会用就BBQ了~~
~~呜呜~~  <br/>

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
    * 实现如下：

```markdown
# 一级标题
## 二级标题
### 三级标题
#### 四级标题
##### 五级标题
###### 六级标题
```
<br/>
<br/>

2. **字体**
    * 星号和下划线都可以，单是斜体，双是粗体，三是斜粗体
    * 实现如下：

```markdown
*这是斜体*        
_这是斜体_        
**这是粗体**      
__这是粗体__      
***这是斜粗体***  
___这是斜粗体___  
```
<br/>
<br/>

3. 换行
    * 一句话后面敲两个空格
    * 两句话之间加一个空行
    * 一行文字中间加<br/>表示换行
<br/>
<br/>

4. 引用
    * 通过符号 **>** 来实现， **>** 后面的空格可有可无
    * 实现如下：

```markdown
> 这是一个引用
>> 这是引用的引用
>>> 这是引用的引用的引用
```
<br/>
<br/>

5. 链接
    * 实现效果:  [链接名称](https://www.baidu.com)
    * 实现如下：

```markdown
[百度](https://www.baidu.com)
```
<br/>
<br/>

6. 图片
    * 实现效果: 
    ![rabbit](https://th.bing.com/th/id/OIP.LYCxLaGlvbCayPsLrkNh3gHaEo?w=276&h=180&c=7&r=0&o=5&dpr=1.5&pid=1.7)
    * 实现如下：

```markdown
![兔斯基](图片地址)
```
