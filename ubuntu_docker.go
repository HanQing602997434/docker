
// Ubuntu Docker安装
/*
	Docker Engine-Community 支持以下的Ubuntu版本
		Xenial 16.04 (LTS)
		Bionic 18.04 (LTS)
		Cosmic 18.10
		Disco 19.04
		其他更新版本...

	Docker Engine-Community 支持上 x86_64 (或amd64) armhf, arm64, s390x (IBM Z),
	和ppc64le (IBM的Power) 架构。

	使用官方安装脚本自动安装
		安装命令：
			curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun

		也可以使用国内daocloud一键安装命令：
			curl -sSL https://get.daocloud.io/docker | sh

	手动安装
		卸载旧版本
			Docker的旧版本被称为docker，docker.io 或 docker-engine。如果已安装，请卸载它们：
				$ sudo apt-get remove docker docker-engine docker.io containerd runc

			当前称为Docker Engine-Community 软件包 docker-ce。
			安装Docker Engine-Community，以下介绍两种方式。
		
		使用Docker仓库进行安装
			在新主机上首次安装Docker Engine-Community 之前，需要设置Docker仓库。之后，可以从
			仓库安装和更新Docker。

		设置仓库
			更新apt包索引
				$ sudo apt-get update

			安装apt依赖包，用于通过HTTPS来获取仓库
				$ sudo apt-get install \
    				apt-transport-https \
    				ca-certificates \
    				curl \
    				gnupg-agent \
					software-properties-common
					
			添加Docker的官方GPG密钥：
				$ curl -fsSL https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu/gpg | sudo apt-key add -

			9DC8 5822 9FC7 DD38 854A E2D8 8D81 803C 0EBF CD88 通过搜索指纹的后8个字符，验证您现在是否拥有带有
			指纹的密钥。
				$ sudo apt-key fingerprint 0EBFCD88
   
				pub   rsa4096 2017-02-22 [SCEA]
      				9DC8 5822 9FC7 DD38 854A  E2D8 8D81 803C 0EBF CD88
				uid           [ unknown] Docker Release (CE deb) <docker@docker.com>
				sub   rsa4096 2017-02-22 [S]

			使用以下指令设置稳定版仓库
				$ sudo add-apt-repository \
   					"deb [arch=amd64] https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu/ \
  					$(lsb_release -cs) \
					  stable"
					  
		安装Docker Engine-Community
			更新apt包索引
				$ sudo apt-get update

			安装最新版本的Docker Engine-Community 和 contained，或者转到下一步安装特定版本：
				$ sudo apt-get install docker-ce docker-ce-cli containerd.io

			要安装特定版本的Docker Engine-Community，请在仓库中列出可用版本，然后选择一种安装。列出您的仓库中可用的版本：
				$ apt-cache madison docker-ce

  				docker-ce | 5:18.09.1~3-0~ubuntu-xenial | https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu  xenial/stable amd64 Packages
  				docker-ce | 5:18.09.0~3-0~ubuntu-xenial | https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu  xenial/stable amd64 Packages
  				docker-ce | 18.06.1~ce~3-0~ubuntu       | https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu  xenial/stable amd64 Packages
  				docker-ce | 18.06.0~ce~3-0~ubuntu       | https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu  xenial/stable amd64 Packages
				...
					  
			使用第二列中的版本字符串安装特定版本，例如5:18.09.1~3-0~ubuntu-xenial
				$ sudo apt-get install docker-ce=<VERSION_STRING> docker-ce-cli=<VERSION_STRING> containerd.io

			测试Docker是否安装成功，输入以下命令，打印出以下信息则安装成功：
				$ sudo docker run hello-world

				Unable to find image 'hello-world:latest' locally
				latest: Pulling from library/hello-world
				1b930d010525: Pull complete                                                                                                                                  
				Digest: sha256:c3b4ada4687bbaa170745b3e4dd8ac3f194ca95b2d0518b417fb47e5879d9b5f
				Status: Downloaded newer image for hello-world:latest


				Hello from Docker!
				This message shows that your installation appears to be working correctly.


				To generate this message, Docker took the following steps:
 					1. The Docker client contacted the Docker daemon.
 					2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    				(amd64)
 					3. The Docker daemon created a new container from that image which runs the
    				executable that produces the output you are currently reading.
 					4. The Docker daemon streamed that output to the Docker client, which sent it
    				to your terminal.

					To try something more ambitious, you can run an Ubuntu container with:
 					$ docker run -it ubuntu bash

					Share images, automate workflows, and more with a free Docker ID:
 					https://hub.docker.com/

					For more examples and ideas, visit:
					https://docs.docker.com/get-started/
						 
		使用Shell脚本进行安装
			Docker在get.docker.com和test.docker.com上提供了方便脚本，用于将快速安装Docker Engine-Community
			的边缘版本和测试版本。脚本的源代码在docker-install仓库中。不建议在生产环境中使用这些脚本，在使用之前
			应该了解潜在的风险：
				1.脚本需要运行root或者具有sudo特权。因此，在运行脚本之前，应仔细检查和审核脚本。

				2.这些脚本尝试检测Linux发行版和版本，并为您配置软件包管理系统。此外，脚本不允许您自定义任何安装
				参数。从Docker的角度或您自己组织的准则和标准的角度来看，这可能导致不支持的配置。

				3.这些脚本将安装软件包管理器的所有依赖和建议，而无需进行确认。这可能会安装大量软件包，具体取决于
				主机的当前配置。

				4.该脚本未提供用于指定要安装哪个版本的Docker的选项。而是安装了在edge通道中发布的最新版本。

				5.如果已使用其他机制将Docker安装在主机上，请不要使用便捷脚本。

			本示例使用get.docker.com上的脚本在Linux上安装最新版本的Docker Engine-Community。要安装最新的测试版本，
			请改用test.docker.com。
			在下面的每个命令，取代每次出现get用test。
				$ curl -fsSL https://get.docker.com -o get-docker.sh
				$ sudo sh get-docker.sh

			如果要使用Docker作为非root用户，则应考虑使用类似以下方式将用户添加到docker组：
				$ sudo usermod -aG docker your-user

		卸载Docker
			删除安装包
				sudo apt-get purge docker-ce

			删除镜像、容器、配置文件等内容：
				sudo rm -rf /var/lib/docker
*/