

## 测试
1. 编译，复制到服务器
```
env GOOS=linux GOARCH=amd64 go build -v main.go
scp main root@10.23.2.11:/root/zzg/      
```

2. 在服务器上运行

```
./main
```

3. 在服务器上开个窗口监控内存变化
```
# 找到pid
ps -ef|grep main
# 跟踪内存变化
watch 'pmap -xp 779504'
```

3. 访问服务，观察内存增长
```
curl http://10.23.2.11:6060/
```
