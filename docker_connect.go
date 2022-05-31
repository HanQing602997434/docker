
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
*/