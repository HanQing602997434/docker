
// Docker镜像使用
/*
	当运行容器时，使用的镜像如果在本地不存在，docker就会自动从docker镜像仓库下载，默认是
	从Docker Hub公共镜像源下载。
	
	列出镜像列表
		我们可以使用docker images来列出本地主机上的镜像。
		各个选项说明：
			REPOSITORY：表示镜像的仓库源
			TAG：镜像的标签
			IMAGE ID：镜像ID
			CREATED：镜像创建的时间
			SIZE：镜像大小
		
		同一仓库可以有多个TAG，代表这个仓库源的不同个版本，如ubuntu仓库源里，有15.10、14.04
		等多个版本，我们使用REPOSITORY:TAG来定义不同的镜像。

		所以，我们如果要使用版本为15.10的ubuntu系统镜像来运行容器时，命令如下：
		docker run -t -i ubuntu:15.10 /bin/bash

		参数说明：
			-i：交互式操作
			-t：终端
			ubuntu15.10：这个是指用ubuntu 15.10版本镜像为基础来启动容器
			/bin/bash：放在镜像名后的是命令，这里我们希望有个交互式Shell，因此用的是/bin/bash
*/