package main
import(

 "fmt"
 "time"
 "sync"

)

func main(){

    channalone   := insert_data_to_channel(1,2,3,4,5,6,7,8,9)
    channaltwo   := insert_data_to_channel(10,11,12,13,14,15,16,17,18,19)
    channalthree := insert_data_to_channel(20,21,22,23,24,25,26,27,28,29)

    for v:= range merge(channalone,channaltwo,channalthree){

    	fmt.Println(v)

    }
 
}

// merge the channel
func merge(chans ... <-chan int) <-chan int{

  out := make(chan int) 

	  go func(){
	  	var wg sync.WaitGroup 
	  	wg.Add(len(chans))
	  	for _,c := range chans {
	  	   go func( c <- chan int){
	  	   	 for v:= range c{

	  	   		out <-v 
	  	   	}

	  	   	wg.Done()

	  	}(c)

 }
	wg.Wait()
	close(out)

}()

return out

}

//insert the value in chennal
func insert_data_to_channel(vs ...int) <-chan int{
	c:= make(chan int)
	go func(){
   		for _,v := range vs{
			c <- v
  			time.Sleep(2 * time.Second)

   		}

   		close(c)

	}()

  return c
}
