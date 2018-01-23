#!/bin/sh

#---------------------- CEk IP Terdaftar --------------------#
OS=`uname -m`;
MYIP=$(curl -4 icanhazip.com)
if [ $MYIP = "" ]; then
   MYIP=`ifconfig | grep 'inet addr:' | grep -v inet6 | grep -vE '127\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}' | cut -d: -f2 | awk '{ print $1}' | head -1`;
fi
MYIP2="s/xxxxxxxxx/$MYIP/g";

#---------------------- CEk IP Terdaftar --------------------#
vps="zvur";
#vps="aneka";
if [ $vps = "zvur" ]; then
        source="https://raw.githubusercontent.com/refky12/ScriptSSHMaster/master/"
else
        source="https://raw.githubusercontent.com/refky12/ScriptSSHMaster/master/"
exit
fi

#---------------------END CEK IP--------------------------#

#check jika script sudah pernah diinput
scriptname='webpanel';
mkdir -p /var/lib/setup-log
echo " " >> /var/lib/setup-log/setup.txt
scriptchecker=`cat /var/lib/setup-log/setup.txt | grep $scriptname`;
if [ "$scriptchecker" != "" ]; then
		clear
		echo  "-----------------------------------------------------------------------------------------";
		echo  "Error! Anda sudah pernah memasukkan script ini sebelumnya";
		echo  "Script ini hanya boleh dimasukkan 1x saja!";
		echo  "-----------------------------------------------------------------------------------------";
		echo  "";
		echo  "------------------------------------------------------------";
		echo  "Jika Anda sebelumnya gagal dalam instalasi, Mohon untuk reinstall OS VPS Anda lebih dulu!";
		echo  "Anda dapat mereinstall OS VPS Anda melalui VPS Control Panel";
		echo  "Cara Mengakses VPS Control Panel: bit.ly/caraaksesvpspanel";
		echo  "-------------------------------------------------------------------------------------------";
        exit 0;
	else
		echo "";
fi
echo "$scriptname" >> /var/lib/setup-log/setup.txt

# go to root
cd

if [ $(id -u) != "0" ]; then
    echo "Error: You have to login by user root!"
    exit
fi
if [ -f /var/cpanel/cpanel.config ]; then
clear
echo "Your Server installed WHM/Cpanel, if you want to use  VPSSIM"
echo "Lets rebuild VPS, you should use centos 6 or 7 - 64 bit"
echo "Bye !"
exit
fi

if [ -f /etc/psa/.psa.shadow ]; then
clear
echo "Server installed Plesk, if you want to use VPSSIM"
echo "Lets rebuild VPS, you should use centos 6 or 7 - 64 bit"
echo "Bye !"
exit
fi

if [ -f /etc/init.d/directadmin ]; then
clear
echo "Your Server installed DirectAdmin, if you want to use VPSSIM"
echo "Lets rebuild VPS, you should use centos 6 or 7 - 64 bit"
echo "Bye !"
exit
fi

if [ -f /etc/init.d/webmin ]; then
clear
echo "Your Server installed webmin, if you want to use VPSSIM"
echo "Lets rebuild VPS, you should use centos 6 or 7 - 64 bit"
echo "Bye !"
exit
fi

if [ -f /home/vpssim.conf ]; then
clear
echo "========================================================================="
echo "========================================================================="
echo "Your Server is installed VPSSIM"
echo "Use command VPSSIM to access VPSSIM menu"
echo "Bye !"
echo "========================================================================="
echo "========================================================================="
rm -rf install*
exit
fi
yum -y install epel-release
if [ -f /etc/yum.repos.d/epel.repo ]; then
sudo sed -i "s/mirrorlist=https/mirrorlist=http/" /etc/yum.repos.d/epel.repo
fi
#yum -y update
wget -q http://vpssim.com/script/vpssim/calc -O /bin/calc && chmod +x /bin/calc
yum -y install psmisc bc gawk gcc
yum -y -q install virt-what unzip sudo net-tools iproute iproute2 curl deltarpm yum-utils tar
clear 
rm -rf /root/vpssim*

prompt="Type in your choice: "
options=("EngLish" "VietNamese" "Cancel")
echo "========================================================================="
echo "VPSSIM Support Centos 6 (32 - 64 Bit) & Centos 7"
echo "-------------------------------------------------------------------------"
echo "You Should Use Centos 6 - 64 bit To Have The Best Performance."
echo "-------------------------------------------------------------------------"
echo "VPSSIM support: English & VietNamese. "
echo "-------------------------------------------------------------------------"
echo "After finished setup, If you want to change language for VPSSIM, using"
echo "-------------------------------------------------------------------------"
echo "[ Change Language ] function in on VPSSIM main menu."
echo "========================================================================="
echo "                     Choose Language For VPSSIM"
echo "========================================================================="
  
PS3="$prompt"
select opt in "${options[@]}"; do 

    case "$REPLY" in
    1) yourlanguage="english"; break;;
    2) yourlanguage="vietnamese"; break;;
    3) yourlanguage="Cancel"; break;;
    *) echo "You typed wrong, Please type in the ordinal number on the list";continue;;
    esac  
done

if [ "$yourlanguage" = "english" ]; then
echo "----------------------------------------------------------------------------"
echo "Please wait ..."
sleep 2
wget -q https://vpssim.com/script/vpssim/centos$(rpm -q --qf "%{VERSION}" $(rpm -q --whatprovides redhat-release))/vpssim-setup && chmod +x vpssim-setup && clear && ./vpssim-setup
elif [ "$yourlanguage" = "vietnamese" ]; then
echo "----------------------------------------------------------------------------"
echo "Please wait ..."
sleep 2
wget -q https://hostingaz.vn/script/vpssim/centos$(rpm -q --qf "%{VERSION}" $(rpm -q --whatprovides redhat-release))/vpssim-setup && chmod +x vpssim-setup && clear && ./vpssim-setup
else 
rm -rf /root/install* && rm -rf /root/vpssim* && clear
fi




 
