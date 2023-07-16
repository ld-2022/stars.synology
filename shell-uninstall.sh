#!/bin/sh

echo "检查脚本权限"
if [ "$(id -u)" -ne 0 ]; then
   echo "此脚本必须以 sudo 权限执行"
   sudo "$0" "$@"
   exit 1
else
    echo "权限正常"
fi

echo "检查是否有旧的程序"
if [ -f "/opt/stars/stars" ]; then
  echo "卸载旧的安装程序"
  if pgrep "stars" > /dev/null
  then
        echo "stars 进程正在运行，准备杀掉..."
        pkill stars
  else
        echo "没有找到 stars 进程。"
  fi
  rm -rf /opt/stars && rm -f /usr/local/bin/stars || {
    echo "卸载失败"
    exit 1
  }
  echo "星空组网卸载成功！"
fi

echo "执行完成 ^~^"
