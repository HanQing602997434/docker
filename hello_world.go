
// hello world
/*
	Docker允许你在容器内运行应用程序，使用 docker run 命令来在容器内运行一个应用程序。
	输出hello world
		runoob@runoob:~$ docker run ubuntu:15.10 /bin/echo "Hello world"
		Hello world

	各个参数解析：
		docker: Docker的二进制执行文件。

		run: 与前面的 docker 组合来运行一个容器。

		ubuntu:15.10 指定要运行的镜像，Docker首先从本地主机上查找镜像是否存在，如果不存在，
		Docker就会从镜像仓库 Docker Hub 下载公共镜像。

		/bin/echo "Hello world": 在启动的容器里执行的命令。
*/