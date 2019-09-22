package main
import (
	"bufio"
	"io"
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
	"os/exec"
)
var (
	start int
	end int 
	len int
	input string
	output string
	printD string
)
func init(){
	flag.IntVar(&start,"s",1,"start position")
	flag.IntVar(&end,"e",1,"end position")
	flag.IntVar(&len,"l",66,"read length")
	flag.StringVar(&input,"input","","input file name,default as stdin")
	flag.StringVar(&output,"output","","output file name,default as stdout")
	flag.StringVar(&printD,"d","","printer destination")
}
func main(){
	flag.Parse()	
	if printD != ""{
		cmd:=exec.Command("lp","-d",printD)
		err:=cmd.Run()
		if err!=nil{
			fmt.Println("printError:",err)
		}	
	}
	if flag.Arg(0) != ""{
		input=flag.Arg(0)
	}
	if output!=""{
		of,err:=os.Create(output)
		defer of.Close()
		if err!= nil{
			fmt.Println("outputFileOpenError:",err)
			return
		}
		if input != ""{
			f,err := os.Open(input)
			if err!=nil{
				fmt.Println("inputFileOpenError:",err)
				return 
			}
			defer f.Close()
			r:=bufio.NewReader(f)
			var line_ctr int=0;
			var page_ctr int=1;
			for {
				buf,err:=r.ReadBytes('\n')
				if err!=nil{
					if err == io.EOF{
						break
					}
					fmt.Println("fileReadError:",err)
				}
				line_ctr++;
				if line_ctr>len{
					page_ctr++;
					line_ctr=1;
				}
				if page_ctr>=start && page_ctr<=end{
					_,err:=of.WriteString(string(buf))
					if err!= nil{
					fmt.Println("fileWriteError:",err)	
					}
				}
			}
		}else{
			var str string;
			var page_ctr int=1;
			for{
				fmt.Scanln(&str);
				if str=="EOF"{
					break;
				}
				if str=="\\f"{
				page_ctr++;
				}
				if page_ctr >= start &&page_ctr <= end{
					_,err:=of.WriteString(string(str+"\n"))
					if err!= nil{
					fmt.Println("fileWriteError:",err)	
					}
				}
			}
		}
	}else{
	if input != ""{
		f,err := os.Open(input)
		if err!=nil{
			fmt.Println("inputFileOpenError:",err)
			return 
		}
		defer f.Close()
		r:=bufio.NewReader(f)
		var line_ctr int=0;
		var page_ctr int=1;
		for {
			buf,err:=r.ReadBytes('\n')
			if err!=nil{
				if err == io.EOF{
					break
				}
				fmt.Println("fileReadError:",err)
			}
			line_ctr++;
			if line_ctr>len{
				page_ctr++;
				line_ctr=1;
			}
			if page_ctr>=start && page_ctr<=end{
				fmt.Printf("%s",string(buf))
			}
		}	
	}else{
		var str string;
		var page_ctr int=1;
		for{
			fmt.Scanln(&str);
			if str=="EOF"{
				break;
			}
			if str=="\\f"{
			page_ctr++;
			}
			if page_ctr >= start &&page_ctr <= end{
				fmt.Println(str);
			}
		}
	}
	}

}
