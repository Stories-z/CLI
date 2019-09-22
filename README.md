# 使用方法
`--s [parameter]` 指定起始页数  
`--e [parameter]` 指定结束页数  
`--input [parameter]` 指定输入文件名，若无则设定输入为标准输入流  
`--output [parameter]` 指定输出文件名，若无则设定输出为标准输出流  
`--l [parameter]` 指定每页行数，默认值为66  
`--d [parameter]` 指定打印地址  

# 程序测试
## 输入为标准输入流，输出为标准输出流
`$main`  
`$hello world`  
$hello world

## 输入为文件"1.txt"，输出为标准输出流
`$main 1.txt `  
1   
2  
3  
4  
......  
66  

## 输入为文件"1.txt"，输出为文件"2.txt"
`$main 1.txt --output 2.txt `

## 指定打印地址为1
`$main --d 1`
