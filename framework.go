
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

	
*/