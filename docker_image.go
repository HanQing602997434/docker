
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
*/