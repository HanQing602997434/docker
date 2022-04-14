
// Dockerfile
/*
	什么是Dockerfile
		Dockerfile 是一个用来构建镜像的文本文件。

	使用 Dockerfile 定制镜像
		1.下面定制一个nginx镜像（构建好的镜像内会有一个 /usr/share/nginx/html/index.html 文件）
			在一个空目录下，新建一个名为 Dockerfile 文件，并在文件内添加以下内容：
			
			FROM nginx
			RUN echo '这是一个本地构建的nginx镜像' > /usr/share/nginx/html/index.html

		2.FROM 和 RUN 指令的作用
			FROM: 定制的镜像都是基于FROM的镜像，这里的 nginx 就是定制需要的基础镜像。后续的操作都是基于nginx。
			RUN: 用于执行后面跟着的命令行命令。有以下两种格式：
				shell格式：
					RUN <命令行命令>
					# <命令行命令> 等同于，在终端操作的 shell 命令。

				exec格式：
					RUN ["可执行文件", "参数1", "参数2"]
					# RUN ["./test.php", "dev", "offline"] 等价于 RUN ./test.php dev offline

		注意：Dockerfile的指令每执行一次都会在 docker 上新建一层。所以过多无意义的层，会造成镜像膨胀过大。例如：
			FROM centos
			RUN yum -y install wget
			RUN wget -O redis.tar.gz "http://download.redis.io/releases/redis-5.0.3.tar.gz"
			RUN tar -xvf redis.tar.gz

		以上执行会创建 3 层镜像。可简化为以下格式：
			FROM centos
			RUN yum -y install wget \
    			&& wget -O redis.tar.gz "http://download.redis.io/releases/redis-5.0.3.tar.gz" \
				&& tar -xvf redis.tar.gz
				
		如上，以 && 符号连接命令，这样执行后，只会创建 1 层镜像。

	开始构建镜像
		在 Dockerfile 文件的存放目录下，执行构建动作。
		以下示例，通过目录下的 Dockerfile 构建一个 nginx:v3（镜像名称:镜像标签）
		注：最后的 . 代表本次执行的上下文路径。

		出现Successfully tagged nginx:v3，说明已经构建成功。

	上下文路径
		上下文路径，是指 docker 在构建镜像，有时候想要使用到本机的文件（比如复制），docker build 命令的指导这个路径，
		会将路径下的文件所有内容打包。

		解析：由于 docker 的运行模式是 C/S。我们本机是C，docker引擎是S。实际在构建过程是在 docker 引擎下完成的，所以
		这个时候无法用我们本机的文件。这就需要把我们本机的指定目录下的文件一起打包提供给 docker 引擎使用。
		
		如果未说明最后一个参数，那么默认上下文路径就是 Dockerfile 所在的位置。
		注意：上下文路径不要放无用的文件，因为会一起打包发送给 docker 引擎，如果文件过多会造成过程缓慢。

	指令详解
		COPY
			复制指令，从上下文目录中复制文件或者目录到容器里指定路径。
*/