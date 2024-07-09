#!/usr/bin/env bash

# This could be done via 'wget -i' but I wanted to ensure things were commented and could be easily understood.

# Raw Binary
wget https://ftp.gnu.org/gnu/binutils/binutils-2.42.tar.gz -O ./fixtures/binutils-2.42.tar.gz
# APK for Alpine 3.20
wget http://dl-cdn.alpinelinux.org/alpine/v3.20/main/x86_64/binutils-2.42-r0.apk -O ./fixtures/binutils-2.42-r0.apk
# rpm for openSUSE Leap 15.6
wget https://ftp.lysator.liu.se/pub/opensuse/distribution/leap/15.6/repo/oss/x86_64/binutils-2.41-150100.7.46.1.x86_64.rpm -O ./fixtures/binutils-2.41-150100.7.46.1.x86_64.rpm
# deb from Debian 12
wget http://ftp.de.debian.org/debian/pool/main/b/binutils/binutils_2.40-2_amd64.deb -O ./fixtures/binutils_2.40-2_amd64.deb
# deb from Ubuntu 24.04 LTS
wget http://archive.ubuntu.com/ubuntu/pool/main/b/binutils/binutils_2.42-4ubuntu2_amd64.deb -O ./fixtures/binutils_2.42-4ubuntu2_amd64.deb
# amazon linux 2
wget https://cdn.amazonlinux.com/2/core/2.0/x86_64/6b0225ccc542f3834c95733dcf321ab9f1e77e6ca6817469771a8af7c49efe6c/../../../../../blobstore/9e065039031da4efbce0e0b51d9ef41fbdc1f51da852b1ef386c6805cded946e/binutils-2.29.1-31.amzn2.x86_64.rpm -O ./fixtures/binutils-2.29.1-31.amzn2.x86_64.rpm
# Arch Linux
wget https://ftp5.gwdg.de/pub/linux/archlinux/core/os/x86_64/binutils-2.42+r91+g6224493e457-1-x86_64.pkg.tar.zst -O ./fixtures/binutils-2.42+r91+g6224493e457-1-x86_64.pkg.tar.zst
# Oracle Linux
wget https://yum.oracle.com/repo/OracleLinux/OL9/baseos/latest/x86_64/getPackage/binutils-2.35.2-43.0.1.el9.x86_64.rpm -O ./fixtures/binutils-2.35.2-43.0.1.el9.x86_64.rpm

wget https://dlcdn.apache.org/spark/spark-3.5.1/spark-3.5.1-bin-without-hadoop.tgz -O ./fixtures/spark-3.5.1-bin-without-hadoop.tgz
#wget https://ftp.gnu.org/gnu/binutils/binutils-2.42.tar.gz -O ./fixtures/binutils-2.42.tar.gz
#wget https://ftp.gnu.org/gnu/binutils/binutils-2.42.tar.gz -O ./fixtures/binutils-2.42.tar.gz
#wget https://ftp.gnu.org/gnu/binutils/binutils-2.42.tar.gz -O ./fixtures/binutils-2.42.tar.gz

