
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

	构建镜像
		我们使用命令docker build，从零开始来创建一个新的镜像。为此我们需要创建一个Dockerfile文件，
		其中包含一组指令来告诉Docker如何构建我们的镜像。
			runoob@runoob:~$ cat Dockerfile 
			FROM    centos:6.7
			MAINTAINER      Fisher "fisher@sudops.com"

			RUN     /bin/echo 'root:123456' |chpasswd
			RUN     useradd runoob
			RUN     /bin/echo 'runoob:123456' |chpasswd
			RUN     /bin/echo -e "LANG=\"en_US.UTF-8\"" >/etc/default/local
			EXPOSE  22
			EXPOSE  80
			CMD     /usr/sbin/sshd -D

		每一个指令都会在镜像上创建一个新的层，每个指令的前缀都必须是大写的。
		每一条FROM，指定使用哪个镜像源。
		RUN指令告诉docker在镜像内执行命令，安装了什么。。。
		然后，我们使用Dockerfile文件，通过docker build命令来构建一个镜像。
			runoob@runoob:~$ docker build -t runoob/centos:6.7 .
			Sending build context to Docker daemon 17.92 KB
			Step 1 : FROM centos:6.7
			 ---&gt; d95b5ca17cc3
			Step 2 : MAINTAINER Fisher "fisher@sudops.com"
			 ---&gt; Using cache
			 ---&gt; 0c92299c6f03
			Step 3 : RUN /bin/echo 'root:123456' |chpasswd
			 ---&gt; Using cache
			 ---&gt; 0397ce2fbd0a
			Step 4 : RUN useradd runoob
			......

		参数说明：
			-t : 指定要创建的目标镜像名
			. : Dockerfile文件所在目录，可以指定Dockerfile的绝对路径
		使用docker images查看已经创建的镜像已经在列表中存在，镜像ID为860c279d2fec

		可以使用新的镜像来创建容器
			docker run -t -i runoob/centos:6.7 /bin/bash
			[root@41c28d18b5fb /]# id runoob
			uid=500(runoob) gid=500(runoob) groups=500(runoob)
		从上面看到新镜像已经包含我们创建的用户runoob

	设置镜像标签
		我们可以使用docker tag命令，为镜像添加一个新的标签。
			docker tag 860c279d2fec runoob/centos:dev

		docker tag 镜像ID，这里是860c279d2fec，用户名称、镜像源名(responsitory name)
		和新的签名(tag)。
		使用docker images命令可以看到，ID为860c279d2fec的镜像多了一个标签
			runoob@runoob:~$ docker images
			REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
			runoob/centos       6.7                 860c279d2fec        5 hours ago         190.6 MB
			runoob/centos       dev                 860c279d2fec        5 hours ago         190.6 MB
			runoob/ubuntu       v2                  70bf1840fd7c        22 hours ago        158.5 MB
			ubuntu              14.04               90d5884b1ee0        6 days ago          188 MB
			php                 5.6                 f40e9e0f10c8        10 days ago         444.8 MB
			nginx               latest              6f8d099c3adc        13 days ago         182.7 MB
			mysql               5.6                 f2e8d6c772c0        3 weeks ago         324.6 MB
			httpd               latest              02ef73cf1bc0        3 weeks ago         194.4 MB
			ubuntu              15.10               4e3b13c8a266        5 weeks ago         136.3 MB
			hello-world         latest              690ed74de00f        6 months ago        960 B
			centos              6.7                 d95b5ca17cc3        6 months ago        190.6 MB
			training/webapp     latest              6fae60ef3446        12 months ago       348.8 MB
*/