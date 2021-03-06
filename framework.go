
// docker架构
/*
	docker包括三个概念：
		镜像（Image）：docker镜像（Image），就相当于是一个root文件系统。比如官方镜像Ubuntu:16.04
		就包含了完整的一套Ubuntu 16.04最小系统的root文件系统。

		容器（Container）：镜像（Image）和容器（Container）的关系，就像是面向对象程序设计中类和实
		例一样，镜像是静态的定义，容器是镜像运行时的实体。容器可以被创建、启动、停止、删除、暂停等等。

		仓库（Repository）：仓库可看成一个代码控制中心，用来保存镜像

	docker使用客户端-服务器（C/S）架构模式，使用远程API来管理和创建docker容器。

	docker容器通过docker镜像来创建。

	容器与镜像的关系类似于面向对象编程中的对象与类
		docker							面向对象
		容器							  对象
		镜像							  类

	概念							   说明
	docker镜像(Images)			   docker镜像是用于创建docker容器的模板，比如Ubuntu系统。
	docker容器(Container)			   容器是独立运行的一个或一组应用，是镜像运行时的实体。
	docker客户端(Client)			  docker客户端通过命令行或其他工具使用docker sdk(https://docs.docker.com/develop/sdk/) 与docker的守护进程通信。
	docker主机(Host)				   一个物理或虚拟的机器用于执行docker守护进程和容器。
	docker registry					docker仓库用来保存镜像，可以理解为代码控制中的代码仓库。
									docker hub(https://hub.docker.com) 提供了庞大的镜像集合供使用。
									一个docker registry中包括多个仓库(reponsitory)；每个仓库包括多个标签(tag)；
									每个标签对应一个镜像。
									通常，一个仓库会包含一个软件不同版本的镜像，而标签就常用于对应该软件的各个版本。
									我们可以通过标签来获取对应版本的镜像，如果没有标签，以lastest作为默认标签。
	docker machine					docker machine是一个简化docker安装的命令行工具，通过一个简单的命令即可在相应的
									平台安装docker，比如VirtualBox、Digital Ocean、Microsoft Azure。
*/