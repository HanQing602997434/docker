
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
			格式：
				COPY [--chown=<user>:<group>] <源路径1>... <目标路径>
				COPY [--chown=<user>:<group>] ["<源路径1>",...  "<目标路径>"]

			[--chown=<user>:<group>]：可选参数，用户改变复制到容器内文件的拥有者和属组。

			<源路径>：源文件或者源目录，这里可以是通配符表达式，其通配符规则要满足 Go 的filepath.Match规则。例如：
			COPY hom* /mydir/
			COPY hom?.txt /mydir/

			<目标路径>：容器内的指定路径，该路径不用事先建好，路径不存在的话，会自动创建。

		ADD
			ADD 指令和 COPY 的使用格式类似（同样的需求下，官方推荐用COPY）。功能也类似，不同之处如下：
				ADD 的优点：在执行<源文件>为 tar 压缩文件的话，压缩格式为gzip，bzip2 以及 xz 的情况下，会自动复制并
				解压到目标路径。
				ADD 的缺点：在不解压的前提下，无法复制 tar 压缩文件。会令镜像构建缓存失效，从而可能令镜像构建变得缓慢。
				具体是否使用，可以根据是否需要自动解压来决定。

		CMD
			类似于 RUN 指令，用于运行程序，但二者运行的时间点不同：
				CMD 在docker run时运行。
				RUN 是在docker build。

			作用：为启动的容器指定默认要运行的程序，程序运行结束，容器也就结束。CMD指令指定的程序可被docker run命令行参数
			中指定要运行的程序所覆盖。

			注意：如果Dockerfile中存在多个CMD命令，仅最后一个生效。
			格式：
				CMD <shell命令>
				CMD ["<可执行文件或命令>","<param1>","param2",...]
				CMD ["<apram1>","<param2>",...] # 该写法是为 ENTRYPOINT 指令指定的程序提供默认参数

			推荐使用第二种格式，执行过程比较明确。第一种格式实际上在运行的过程中也会自动转换成第二种格式运行，并且默认可执行
			文件是sh。

		ENTRYPOINT
			类似于CMD指令，但其不会被docker run的命令行指定的指令覆盖，而且这些命令行参数会被当作参数送给ENTRYPOINT指令指定
			的程序。
			但是，如果运行docker run时使用了 --entrypoint选项，将覆盖ENTRYPOINT指令指定的程序。
			优点：在执行docker run的时候可以指定ENTRYPOINT运行所需的参数。
			注意：如果Dockerfile中如果存在多个ENTRYPOINT指令，仅最后一个生效。
			格式：
				ENTRYPOINT ["<executeable>","<param1>","<param2>",...]
				可以搭配CMD命令使用：一般是变参才会使用CMD，这里的CMD等于是在给ENTRYPOINT传参，以下示例会提到。
			示例：
				假设已通过Dockerfile构建nginx:test镜像：
					FROM nginx

					ENTRYPOINT ["nginx","-c"] # 定参
					CMD ["/etc/nginx/nginx.conf"] # 变参

			1.不传参运行
				$ docker run nginx:test

				容器内会默认运行以下命令，启动主进程
				nginx -c /etc/nginx/nginx.conf

			2.传参运行
				$ docker run nginx:test -c /etc/nginx/new.conf

				容器会默认运行以下命令，启动主进程（/etc/nginx/new.conf，假设容器内有此文件）
				nginx -c /etc/nginx/new.conf

		ENV
			设置环境变量，定义环境变量，那么在后续的指令中，就可以使用这个环境变量。
			格式：
				ENV <key> <value>
				ENV <key>=<value1> <key2>=<value2>...

			以下示例设置NODE_VERSION = 7.2.0，在后续的指令中可以通过 $NODE_VERSION引用：
				ENV NODE_VERSION 7.2.0

				RUN curl -SLO "https://nodejs.org/dist/v$NODE_VERSION/node-v$NODE_VERSION-linux-x64.tar.xz" \
				  && curl -SLO "https://nodejs.org/dist/v$NODE_VERSION/SHASUMS256.txt.asc"
				  
		ARG
			构建参数，与ENV作用一致，不过作用域不一样。ARG设置的环境变量仅对Dockerfile内有效，也就是说只有docker build
			的过程有效，构建好的镜像内不存在此环境变量。
			构建命令docker build中可以用--build-arg<参数名>=<值>来覆盖。
			格式：
				ARG <参数名>[=<默认值>]

		VOLUME
			定义匿名数据卷。在启动容器时忘记挂在数据卷，会自动挂载到匿名卷。
			作用：
				避免重要的数据，因容器重启而丢失，这是非常致命的。
				避免容器不断变大。
			格式：
				VOLUME ["<路径1>","<路径2>"...]
				VOLUME <路径>
			在启动容器docker run的时候，我们可以通过 -v 参数修改挂载点

		EXPOSE
			仅仅只是声明端口
			作用：
				帮助镜像使用者理解这个镜像服务的守护端口，以方便配置映射。
				在运行时使用随机端口映射时，也就是docker run -P时，会自动随机映射EXPOSE的端口。
			格式：
				EXPOSE <端口1> [<端口2>...]

		WORKDIR
			指定工作目录。用 WORKDIR 指定的工作目录，会在构建镜像的每一层中都存在。（WORKDIR 指定的工作目录，必须是提前创建好的）。
			docker build 构建镜像过程中的，每一个 RUN 命令都是新建的一层。只有通过 WORKDIR 创建的目录才会一直存在。

			格式：
				WORKDIR <工作目录路径>

		USER
			用于指定执行后续命令的用户和用户组，这边只是切换后续命令执行的用户（用户和用户组必须提前已经存在。）
			格式：
				USER <用户名>[:<用户组>]

		HEALTHCEHCK
			用于指定某个程序或者指令来监控docker容器服务的运行状态。
*/