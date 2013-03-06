package main

import "fmt"
import "reflect"

func main(){
	var x float64 = 3.4;
	
	fmt.Println("type:", reflect.TypeOf(x) )

	v := reflect.ValueOf(x)
	fmt.Println("type=", v.Type() )
	fmt.Println( "v.Float()=",v.Float() )
}
