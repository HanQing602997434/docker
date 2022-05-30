
// docker 容器连接
/*
	通过网络端口来访问运行在docker容器内的服务
	容器中运行的网络应用，要让外部也可以访问这些应用，可以通过 -P 或 -p 参数
	来指定端口映射。
	如果实现通过端口连接到一个docker容器：

	网络端口映射
		创建一个Python应用的容器。
			runoob@runoob:~$ docker run -d -P training/webapp python app.py
			fce072cc88cee71b1cdceb57c2821d054a4a59f67da6b416fceb5593f059fc6d

		另外，我们可以指定容器绑定的网络地址，比如绑定127.0.0.1。
		我们使用-P绑定端口号，使用docker ps可以看到容器端口5000绑定主机端口32768.
			runoob@runoob:~$ docker ps
			CONTAINER ID    IMAGE               COMMAND            ...           PORTS                     NAMES
			fce072cc88ce    training/webapp     "python app.py"    ...     0.0.0.0:32768->5000/tcp   grave_hopper

		我们也可以使用-p标识来指定容器端口绑定到主机端口。
		两种方式的区别是：
			-P:是容器内部端口随机映射到主机的端口。
			-p:是容器内部端口绑定到指定的主机端口。
*/