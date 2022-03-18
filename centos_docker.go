
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
*/