package main
import (
	"os"
	"fmt"
	"log"
	"net/rpc"
)
//两边的类型必须一致
type Args struct{
	A,B int //大写才可被识别到
}
type Quotient struct{
	Quo,Rem int
} 
func  main()  {
	if len(os.Args)!=2{
		fmt.Println("usage:",os.Args[0],"server")
		os.Exit(1)
	}
	serverAddr:=os.Args[1]
	client,err:=rpc.DialHTTP("tcp",serverAddr+":1234")
	if err!=nil{
		log.Fatal("dailing",err)
	}
	args:=Args{43,8}
	var reply int
	err=client.Call("Math.Mutiply", args,&reply)
	if err!=nil{
		log.Fatal("Math error:",err)
	}
	fmt.Printf("%d * %d =%d\n",args.A,args.B,reply)
	var Quo Quotient
	err=client.Call("Math.Device", args,&Quo)
	if err!=nil{
		log.Fatal("Math error:",err)
	}
	fmt.Printf("%d / %d = %d :re %d\n",args.A,args.B,Quo.Quo,Quo.Rem)
}

