#!/bin/bash
passwd='811115Abc&'
/usr/bin/expect <<-EOF
set time 30
spawn ssh -p18330 yaogang@10.149.19.73
expect {
"*yes/no" { send "yes\r"; exp_continue }
"*password:" { send "$passwd\r" }
}
expect "*#"
send "cd /home/trunk\r"
expect "*#"
send "svn up\r"
expect "*#"
send "exit\r"
interact
expect eof
EOF
