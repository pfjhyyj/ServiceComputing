# Selpg Homework
Reference [https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html)
## use
### install required package
```
go get github.com/spf13/pflag
```
### usage
```
USAGE: go run selpg.go -sstart_page -eend_page [ -f | -llines_per_page ] [ -ddest ] [ in_filename ]
```
PS: -d is not available now
### test result
```
go run .\selpg.go -s1 -e2 -f .\file1
line 1 of file one
line 2 of file one
line 3 of file one
line 4 of file one
line 5 of file one

line 6 of file one
line 7 of file one
line 8 of file one
line 9 of file one
line 10 of file one

line 11 of file one
line 12 of file one
line 13 of file one
line 14 of file one
line 15 of file one
C:\Users\HEARTU~1\AppData\Local\Temp\go-build978124194\command-line-arguments\_obj\exe\selpg.exe: end_page (2) greater than total pages (1), less output than expected
C:\Users\HEARTU~1\AppData\Local\Temp\go-build978124194\command-line-arguments\_obj\exe\selpg.exe: done
```
