# Git Learn

---
this is in read.md file 
**learning start**
~~英文能力有限之后就暂时使用中文来写吧~~
~~先把git 弄明白再写markdown~~
~~竟然有人这个时候还不熟git和markdown是谁就不说了~~

## Start Basic

---
* 建立远程仓库 复制ssh地址,后面要用,不需要添加 read.md 文件
    * 本地创建和仓库同名的文件夹
    * bash输入 `git init` 初始化本地仓库，这时候会多出俩一个 .git 文件，不要动它
    * bash输入 `git remote add origin copy` 将远程仓库在本地起名为 origin
    * `git add .` 将当前所有修改添加到暂存区
    * `git commit -m "your words"` 将当前暂存区的内容提交到本地仓库，并添加描述
    * `git push -u origin master:main` 第一次提交
    * `git push origin master:main` 之后所有的提交
