
// Docker 容器使用
/*
	Docker 客户端
		docker客户端非常简单，我们可以直接输入 docker 命令来查看到 Docker 客户端的所有命令选项。
			runoob@runoob:~# docker

		可以通过命令 docker command --help更深入的了解指定的Docker命令使用方法。
		例如我们要查看 docker status 指令的具体使用方法：
			runoob@runoob:~# docker stats --help

	容器使用
		获取镜像
			如果我们本地没有ubuntu镜像，我们可以使用docker pull命令来载入ubuntu镜像：
			$ docker pull ubuntu

		启动容器
			以下命令使用ubuntu镜像启动一个容器，参数以命令行模式进入该容器
			$ docker run -it ubuntu /bin/bash

			参数说明：
				-i：交互式操作。
				-t：终端。
				ubuntu：ubuntu镜像。
				/bin/bash：放在镜像名后的是命令，这里我们希望有个交互式Shell，因此用的是/bin/bash。

			要退出终端，直接输入exit：
				root@ed09e4490c57:/# exit

		启动已停止运行的容器
			查看所有的容器命令如下：
				$ docker ps -a

			使用docker start启动一个已停止的容器：
				$ docker start b750bbbcfd88
*/