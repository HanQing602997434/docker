
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
*/