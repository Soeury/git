# Git Learn

this is in read.md file <br/>
**learning start** <br/>

---
~~英文能力有限之后就暂时使用中文来写吧~~  <br/>
~~先把git 弄明白再写markdown~~  <br/>
~~竟然有人这个时候还不熟git和markdown是谁就不说了~~  <br/>
~~团队开发的时候不会用就BBQ了~~ <br/>
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

3. **换行**
    * 一句话后面敲两个空格
    * 两句话之间加一个空行
    * 一行文字中间加<br/>表示换行
<br/>
<br/>

4. **引用**
    * 通过符号 **>** 来实现， **>** 后面的空格可有可无
    * 实现如下：

```markdown
> 这是一个引用
>> 这是引用的引用
>>> 这是引用的引用的引用
```
<br/>
<br/>

5. **链接**
    * 实现效果:  [链接名称](https://www.baidu.com)
    * 实现如下：

```markdown
[百度](https://www.baidu.com)
```
<br/>
<br/>

6. **图片**
    * ![rabbit](https://th.bing.com/th/id/OIP.LYCxLaGlvbCayPsLrkNh3gHaEo?w=276&h=180&c=7&r=0&o=5&dpr=1.5&pid=1.7)
    * 实现效果: 
    * 实现如下：

```markdown
![兔斯基](图片地址)
```
<br/>
<br/>

7. **列表**
    * 列表分成有序列表和无序列表，
    * 无序列表：使用 **\*** 、 **\+** 、 **\-** 再加一个空格的方式来实现
    * 有序列表: 使用 **1.** 、 **2.** 、 **3.** 再加一个空格的方式来实现
    * 如果需要控制层级，需要在符号列表之前使用 **Tab**
    * 实现如下:

```markdown
* 无序列表 1
    + 无序列表 1.1
    - 无序列表 1.2
        + 无序列表 1.2.3
+ 无序列表 2
- 无序列表 3

1. 有序列表 1
    1. 有序列表 1.1
2. 有序列表 2
3. 有序列表 3
```
<br/>
<br/>

8. **分割线**
    * 可以使用三连个 **\-** 或者 三连个 **\*** 来实现分割线，注意，分割线的上面需要空一行
    * 实现如下：

```markdown
---
***
```
<br/>
<br/>

9. **删除线**
    * 在需要删除的文字前后分别加两个 **~** 来实现
    * 实现如下:

```markdown
~~这是需要删除的文字~~
```
<br/>
<br/>

10. **代码块**
    * 一行内需要引用代码，则使用一对反引号就可以
    * 如果需要引用代码块，则需要在最前面和最后面的单独一行加上分别加上三个反引号，第一行需要标明代码块使用的语言
    * 实现如下:  ....
    * ~~这要怎么实现啊，反引号包反引号没用。我觉得说得已经很清楚了。(因为懒所以就不给出图片了)~~
<br/>
<br/>

11. **特殊符号**
    * 引用特殊符号在前面加上反斜杠 **\\** 即可
    * 实现如下:

```markdown
\\
\*
\_
\.
\+
```
<br/>
<br/>

12. **待办事项**
    * 支持用 markworn 来完成待办事项， - [ ] 表示未完成 ， - [x] 表示已完成，注意格式
    * 实现如下:

```markdown
- [ ] 吃饭
- [ ] 睡觉
- [x] 打豆豆
```
* 效果如下
- [ ] 吃饭
- [ ] 睡觉
- [x] 打豆豆
<br/>
<br/>

13. **脚注**
    * 脚注[^1]是对文本的注释，在 **markdown** 中编写脚注的方法就是一对中括号+ **^** + 数字 的方式
    * 实现如下:

```markdown
这是一段文本[^1]，介绍的是小熊猫[^2]
[^1]: 文本是用来描述某个事物或者某种规律现象一段文字的组合
[^2]: 小熊猫是国家二级保护动物，不信你去查
```

[^1]:脚注是什么就不多说了
   

