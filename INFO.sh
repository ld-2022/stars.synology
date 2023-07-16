#!/bin/bash
# Copyright (c) 2000-2020 Synology Inc. All rights reserved.

source /pkgscripts/include/pkg_util.sh

package="starsInstall"
version="1.0.0-0001"
displayname="星空云-安装器"
os_min_ver="7.0-40000"
maintainer="北京现伟科技有限公司"
arch="$(pkg_get_platform)"
description="星空 使用 P2P 技术进行数据传输，避免了中心节点的流量瓶颈，提高了传输效率。"
dsmapp="starsInstall"
dsmuidir="ui"
[ "$(caller)" != "0 NULL" ] && return 0
pkg_dump_info
