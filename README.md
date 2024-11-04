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
    * `git branch name` 创建分支
    * `git branch -v` 查看分支
    * `git checkout name` 切换到指定分支
    * `git merge name` 把指定的name分支合并到当前的分支上

* 总结：
    1. 创建一个分支 ，并切换进入该分支
    2. 在该分支内进行开发，建立一个总文件库

* 感悟：大概了解了,之前个人开发都不怎么用 git，上传文件都是直接 upload ... 
* 总结：~~我是笨笨~~