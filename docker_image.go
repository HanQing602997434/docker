
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

	获取一个新的镜像
		当我们在本地主机上使用一个不存在的镜像时 Docker 就会自动下载这个镜像。如果我们预先下载
		这个镜像，我们可以使用docker pull命令下载它。
		下载完成后，我们可以直接使用这个镜像来运行容器。

	查找镜像
		我们可以从Docker Hub网站来搜索镜像，Docker Hub网站为：https://hub.docker.com/
		我们也可以使用docker search命令来搜索镜像。比如我们需要一个httpd的镜像来作为我们
		的web服务，我们可以通过docker search命令搜索httpd来寻找适合我们的镜像。
			docker search httpd

			NAME:镜像仓库源的名称
			DESCRIPTION:镜像的描述
			OFFICIAL:是否docker官方发布
			stars:类似Github里面的star，表示点赞、喜欢的意思
			AUTOMATED:自动构建

	拉取镜像
		我们决定使用上图中的httpd官方版本的镜像，使用命令docker pull来下载镜像
			docker pull httpd

		下载完成后，我们就可以使用这个镜像了
			docker run httpd

	删除镜像
		镜像删除使用docker rmi命令，比如删除hello-world镜像
			docker rmi hello-world

	创建镜像
		当我们从docker镜像仓库下载的镜像不能满足我们的需求时，我们可以通过以下两种方式对镜像进行修改
			1.从已经创建的容器中更新镜像，并且提交这个镜像
			2.使用Dockerfile指令来创建一个新的镜像

	更新镜像
		更新镜像之前，我们需要使用镜像来创建一个容器
			docker run -t -i ubuntu:15.10 /bin/bash

		在运行的容器内使用apt-get update命令进行更新。
		在完成操作之后，输入exit命令来退出这个容器。
		通过命令docker commit来提交容器副本。
			docker commit -m="has update" -a="runoob" e218edb10161 runoob/ubuntu:v2
			sha256:70bf1840fd7c0d2d8ef0a42a817eb29f854c1af8f7c59fc03ac7bdee9545aff8

			各个参数说明：
				-m:提交的描述信息
				-a:指定镜像作者
				e218edb10161:容器ID
				runoob/ubuntu:v2:指定要创建的目标镜像名

		我们可以使用docker images命令来查看我们的新镜像 runoob/ubuntu:v2
			docker images

		使用我们的新镜像 runoob/ubuntu来启动一个容器
			docker run -t -i runoob/ubuntu:v2 /bin/bash
*/