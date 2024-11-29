# ansible

* ansible安装路径
  1. /etc/ansible
  2.
* ansible configure (优先级)
    1. ANSIBLE_CFG（变量）优先级最高
    2. (当前目录)/"
    3. ~/ansible.cfg
    4. /etc/ansible/ansible.cfg
* ansible 文件组织结构
  1. vars
  2. default
  3. tasks
  4. template
* ansible的变量类型和优先级
  1. inventory 下的变量 host_vars, group_vars {{hostvars.变量名}} {{hostvars.ansible_ssh_user|'root'}}
  2. play variable  在playbook中定义的vars 和vars_files
  3. role variable
  4. tasks variable  register和set_fact
  5. extra variable 执行 playbook  --extra-vars "name=name-extra"

* ansible的内置命令有哪些
   1. FILE
    state: directory/file/absent/hark/link
   2. COPY
   3. Command/Shell
   4. System

* 如何定义一个role
* role的目录结构
* import和include的区别
* ansible的控制机和受控机器的的交互方式

* ansible的模板jinjia2常用的的语法
* 执行策略 linear/free
* inventory（静态和动态）

# shell

* sed
* find
* awk
* top
* iotop
* nginx
* 字符截取
   1. 格式：${string:start:length}
   2. 格式1：${string%substr*}                          #匹配从右往左第一个substr
   3. 格式2：${string%%substr*}                       #匹配从右往左最后一个substr
   4. 格式1：${string#*substr}                          #匹配从左往右第一个substr
   5. 格式2：${string##*substr}                        #匹配从左往右最后一个substr
   6. 用cut命令截取（适合处理管道流或行文本字符)
* 数组定义和循环遍历数组
   1. arr=("a" "b" "c")
   2. ${#arr[@]} 数组的长度
   3. ${arr[@]} 数据的内容
* 文件权限755的含义
* cron 表达式的含义
* systemd type
   1. forking
   2. simple
   3. other
* lsof（查看文件被那个进程引用)

# Git常用命令

# Docker

* 常用命令
    1. FROM
    2. RUN
    3. COPY
    4. volume： volume类型，
    5. CMD
    6. ARG
    7. ENTRYPOINT
    8. 跨架构的构建
    9. 减少镜像大小的手段
    1、基础镜像
* docker 的缓存
* 参考：<https://zhuanlan.zhihu.com/p/670003782>

# 常用的CI/CD工具

* <https://mp.weixin.qq.com/s/bhJ1erCLzW-7WWrrLi1aYA>
* <https://mp.weixin.qq.com/s/LpEmabOXjsRFFksKHMXhcw>
* <https://mp.weixin.qq.com/s/TLoz4pgelJ_Kp5GO1g8CnA>

# kubernetes

* 原理介绍
* 常用组件
* pod、service、deployment、ingress

# 网络

* DNS解析
* HTTP状态码
* HTTPS（加分项）
* 7层、4层网络模型
* Tcp/IP Tcp-Udp区别

# 解答题

* 给定一个路径，如果是文件就输出文件类型， 如果是文件夹就列出这个文件夹下的文件个数并统计不同文件类型的个数。
 实现方式：
  python
  shell

* 服务器负载高，该如何入手查询

# prometheus 部署方式

* federation
* remote_writer
* 参考：<https://prometheus.io/docs/prometheus/latest/federation/>

# mysql数据库

* sql 语句
* 创建数据库的语句
* 创建用户和给用户赋值读写权限
* 主从的设置流程

个人简介
了解过哪些开发语言
教育经历
工具经历
使用的技术
公有云和私有云
排查问题的方式
 介绍思路
  1、日志
