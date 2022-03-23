
// CentOS Docker安装
/*
	Docker支持以下的64位CentOS版本：
		CentOS 7
		CentOS 8
		更高版本...

	使用官方安装脚本自动安装
		安装命令如下：
			curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun

		也可以使用国内daocloud一键安装命令：
			curl -sSL https://get.daocloud.io/docker | sh

	手动安装
		卸载旧版本
			较旧的Docker版本称为docker或docker-engine。如果已安装这些程序，请卸载它们以及
			相关的依赖项。
				$ sudo yum remove docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
				  docker-engine
				  
	安装Docker Engine-Community
	使用Docker仓库进行安装
		在新主机上首次安装Docker Engine-Community之前，需要设置Docker仓库。之后，可以从仓库
		安装和更新Docker。

		设置仓库
		安装所需的软件包。yum-utils提供了yum-config-manager，并且device mapper存储驱动程序需要
		device-mapper-persistent-data 和 lvm2。
			$ sudo yum install -y yum-utils \
  			device-mapper-persistent-data \
			  lvm2
			  
		使用以下命令来设置稳定的仓库
		使用官方源地址（比较慢）
			$ sudo yum-config-manager \
    		--add-repo \
			https://download.docker.com/linux/centos/docker-ce.repo
			
		可以选择国内的一些源地址
			阿里云
				$ sudo yum-config-manager \
    			--add-repo \
				http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
				
			清华大学源
				$ sudo yum-config-manager \
    			--add-repo \
				   https://mirrors.tuna.tsinghua.edu.cn/docker-ce/linux/centos/docker-ce.repo
				   
	安装 Docker Engine-Community
		安装最新版本的Docker Engine-Community 和 contained，或者转到下一步安装特定版本：
			$ sudo yum install docker-ce docker-ce-cli containerd.io

		如果提示您接受GPG密码，请选是。
			有多个Docker仓库吗？
			如果启用了多个Docker仓库，则在未在yum install 或 yum update 命令中指定版本的情况下，进行
			的安装或更新始终安装最高版本，这可能不适合您的稳定性需求。

		Docker安装完默认未启动。并且已经创建好docker用户组，但该用户组下没有用户。

	要安装特定版本的Docker Engine-Community，请在存储库中列出可用版本，然后选择并安装：
		1.列出并排序您存储库中可用的版本。此示例按版本号（从高到低）对结果进行排序。
			$ yum list docker-ce --showduplicates | sort -r

			docker-ce.x86_64  3:18.09.1-3.el7                     docker-ce-stable
			docker-ce.x86_64  3:18.09.0-3.el7                     docker-ce-stable
			docker-ce.x86_64  18.06.1.ce-3.el7                    docker-ce-stable
			docker-ce.x86_64  18.06.0.ce-3.el7                    docker-ce-stable

		2.通过其完整的软件包名称安装特定版本，该软件包名称（docker-ce）加上版本字符串（第二列），从第一个
		冒号（:）一直到第一个连字符，并用连字符（-）分隔。例如：docker-ce-18.09.1。
			$ sudo yum install docker-ce-<VERSION_STRING> docker-ce-cli-<VERSION_STRING> containerd.io

	启动Docker
		$ sudo systemctl start docker

	通过运行hello-world 映像来验证是否正确安装了Docker Engine-Community。
		$ sudo docker run hello-world
*/