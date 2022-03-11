
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
*/