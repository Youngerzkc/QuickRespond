package main
import (
	"fmt"
	"errors"
	"net/http"
	"net/rpc"
)
type Args struct{
	A,B int //大写才可被识别到
}
type Math int
type Quotient struct{
	Quo,Rem int
} 
//有规定,必须是一个类型，方法必须是某个类型
//func ()NAME(ARGS TYPE,REPLAY *TYPE ) error
 func (m *Math) Mutiply(args *Args,reply *int) error  {
	 *reply=args.A*args.B
	 return nil
 }
 func (m *Math)Device(args *Args,reply *Quotient)error  {
	if args.B==0 {
		return errors.New("device by zero")
	}
	reply.Quo=args.A/args.B
	reply.Rem=args.A%args.B
	return nil
 }
 func  main()  {
	 math:=new(Math)
	 rpc.Register(math)
	 rpc.HandleHTTP()
	 err:=http.ListenAndServe(":1234",nil)
	 if err!=nil{
		 fmt.Println(err.Error())
	 }
 }