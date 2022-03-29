
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

	以上命令的完整意思可以解释为：Docker以ubuntu15.10镜像创建一个新容器，然后在容器里执行
	bin/echo "Hello world"，然后输出结果。

	进行交互式的容器
		我们通过docker的两个参数-i -t，让docker运行的容器实现"对话"的能力：
			runoob@runoob:~$ docker run -i -t ubuntu15.10 /bin/bash
			root@0123ce188bd8:/#

		各个参数解析：
			-t：在新容器内指定一个伪终端或终端。
			-i：允许你对容器内的标准输入（STDIN）进行交互。
*/