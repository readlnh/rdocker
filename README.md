# rdocker
mydocker在ubuntu19.04，内核4.18.0-16-generic上的实现

## 启动流程
main解析命令->runCommand创建namespace隔离容器进程->NewParentProcess返回配置好隔离参数的command参数的对象
启动容器进程->容器自身调用自身，初始化文件内容，挂载proc文件系统

在NewParentProcess中
```golang
  args := []string{"init", command}
  cmd := exec.Command("/proc/self/exe", args...)
```
`/proc/self`指的是当前运行自己的环境，exec是自己调用了自己，后面的args则是参数，即init是传递给本进程的第一个参数

## 关于主机的proc被占用问题
systemd加入linux之后, mount namespace就变成 shared by default, 所以必须显式声明你要这个新的mount namespace独立
