// calculate
package calculate

func addOne(data int) int {
	data = data + 1
	return data
}

func getMaxData(data1 int,data2 int) int {
	if(data1 >= data2){
		return data1
	}else{
		return data2
	}
}

func HanNuo(data int) int {
	if(data == 1){
		return 1
	}else{
		return (data + HanNuo(data -1))
	}
}
