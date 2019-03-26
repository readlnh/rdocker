# rdocker
mydocker在ubuntu19.04，内核4.18.0-16-generic上的实现

## 关于主机的proc被占用问题
systemd加入linux之后, mount namespace就变成 shared by default, 所以必须显式声明你要这个新的mount namespace独立
