.PHONY: all servers clean

# 默认目标
all: servers

# 调用批处理脚本来处理所有 .thrift 文件
servers:
	processidl.bat

# 清理目标
clean:
	@echo Cleaning up...
