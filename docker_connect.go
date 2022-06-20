
// docker 容器连接
/*
	通过网络端口来访问运行在docker容器内的服务
	容器中运行的网络应用，要让外部也可以访问这些应用，可以通过 -P 或 -p 参数
	来指定端口映射。
	如果实现通过端口连接到一个docker容器：

	网络端口映射
		创建一个Python应用的容器。
			runoob@runoob:~$ docker run -d -P training/webapp python app.py
			fce072cc88cee71b1cdceb57c2821d054a4a59f67da6b416fceb5593f059fc6d

		另外，我们可以指定容器绑定的网络地址，比如绑定127.0.0.1。
		我们使用-P绑定端口号，使用docker ps可以看到容器端口5000绑定主机端口32768.
			runoob@runoob:~$ docker ps
			CONTAINER ID    IMAGE               COMMAND            ...           PORTS                     NAMES
			fce072cc88ce    training/webapp     "python app.py"    ...     0.0.0.0:32768->5000/tcp   grave_hopper

		我们也可以使用-p标识来指定容器端口绑定到主机端口。
		两种方式的区别是：
			-P:是容器内部端口随机映射到主机的端口。
			-p:是容器内部端口绑定到指定的主机端口。

			runoob@runoob:~$ docker run -d -p 5000:5000 training/webapp python app.py
			33e4523d30aaf0258915c368e66e03b49535de0ef20317d3f639d40222ba6bc0

		另外，我们可以指定容器绑定的网络地址，比如绑定127.0.0.1。
			runoob@runoob:~$ docker run -d -p 127.0.0.1:5001:5000 training/webapp python app.py
			95c6ceef88ca3e71eaf303c2833fd6701d8d1b2572b5613b5a932dfdfe8a857c
			runoob@runoob:~$ docker ps
			CONTAINER ID        IMAGE               COMMAND           ...     PORTS                                NAMES
			95c6ceef88ca        training/webapp     "python app.py"   ...  5000/tcp, 127.0.0.1:5001->5000/tcp   adoring_stonebraker
			33e4523d30aa        training/webapp     "python app.py"   ...  0.0.0.0:5000->5000/tcp               berserk_bartik
			fce072cc88ce        training/webapp     "python app.py"   ...    0.0.0.0:32768->5000/tcp              grave_hopper

		这样我们就可以通过访问127.0.0.1:5001来访问容器的5000端口
		上面的例子中，默认都是绑定TCP端口，如果要绑定UDP端口，可以在端口后面加上/udp
			runoob@runoob:~$ docker run -d -p 127.0.0.1:5000:5000/udp training/webapp python app.py
			6779686f06f6204579c1d655dd8b2b31e8e809b245a97b2d3a8e35abe9dcd22a
			runoob@runoob:~$ docker ps
			CONTAINER ID        IMAGE               COMMAND           ...   PORTS                                NAMES
			6779686f06f6        training/webapp     "python app.py"   ...   5000/tcp, 127.0.0.1:5000->5000/udp   drunk_visvesvaraya
			95c6ceef88ca        training/webapp     "python app.py"   ...    5000/tcp, 127.0.0.1:5001->5000/tcp   adoring_stonebraker
			33e4523d30aa        training/webapp     "python app.py"   ...     0.0.0.0:5000->5000/tcp               berserk_bartik
			fce072cc88ce        training/webapp     "python app.py"   ...    0.0.0.0:32768->5000/tcp              grave_hopper

		docker port命令可以让我们快捷地查看端口的绑定情况。
			runoob@runoob:~$ docker port adoring_stonebrake 5000
			127.0.0.1：5001

		Docker容器互联
			端口映射并不是唯一把docker连接到另一个容器的方法。
			docker有一个连接系统允许将多个容器连接在一起，共享连接信息。
			docker连接会创建一个父子关系，其中父容器可以看到子容器的信息。

		容器命名
			当我们创建一个容器的时候，docker会自动对它进行命名。另外，我们也可以使用--name标识来命名容器，例如：
				runoob@runoob:~$ docker run -d -P --name runoob training/webapp python app.py
				43780a6eabaaf14e590b6e849235c75f3012995403f97749775e38436db9a441

			我们可以使用docker ps命令来查看容器名称。
				runoob@runoob:~$ docker ps -l
				CONTAINER ID     IMAGE            COMMAND           ...    PORTS                     NAMES
				43780a6eabaa     training/webapp   "python app.py"  ...     0.0.0.0:32769->5000/tcp   runoob

	新建网络
		下面先新建一个新的Docker网络
		$ docker network create -d bridge test-net
		
		参数说明：
			-d: 参数指定Docker网络类型，有bridge、overlay。
			其中overlay网络类型用于Swarm mode。

		连接容器
			运行一个容器并连接到新建的test-net网络：
				$ docker run -itd --name test1 test-net ubuntu /bin/bash

			打开新的终端，再运行一个容器并加入到test-net网络：
				$ docker run -itd --name test2 test-net ubuntu /bin/bash

			通过ping可以证明test1和test2容器建立了互联关系。

	配置DNS
		可以在宿主机的/etc/docker/daemon.json文件中增加以下内容来设置全部容器的DNS：
		{
			"dns" : [
				"114.114.114.114",
				"8.8.8.8"
			]
		}
		设置完成后，启动容器的DNS会自动配置为114.114.114.114和8.8.8.8
		配置完成，需要重启docker才能生效。
		查看容器的DNS是否生效可以使用以下命令，它会输出容器的DNS信息：
		$ docker run -it --rm ubuntu cat etc/resolv.conf

		手动指定容器的配置
		如果只想在指定的容器设置DNS，则可以使用以下命令：
			$ docker run -it --rm -h host_ubuntu --dns=114.114.114.114 --dns-search=test.com ubuntu

		参数说明：
			--rm：容器退出时自动清理容器内部的文件系统。
			-h HOSTNAME 或者 --hostname=HOSTNAME: 设定容器的主机名，他会被写到容器内的/etc/hostname 和 /etchosts。
			--dns=IP_ADDRESS：添加DNS服务器到容器的/etc/resolv.conf中，让容器用这个服务器来解析所有不在/etc/hosts中的主机名
*/