// You can edit this code!
// Click here and start typing.
package main

type Data struct {
	DataValues []*DataValue
}

type DataValue struct {
	Type  string
	Value []string
}

func main() {
	//d := Data{
	//	DataValues: []*DataValue{
	//		{
	//			Type:  "color",
	//			Value: []string{"red", "blue", "black"},
	//		},
	//		{
	//			Type:  "status",
	//			Value: []string{"new", "used"},
	//		},
	//		{
	//			Type:  "brand",
	//			Value: []string{"apple", "samsung", "sony"},
	//		},
	//	},
	//}

	//dataMap := make(map[string][]string, 0)
	//
	//for i := 0; i < len(d.DataValues[0].Value); i++ {
	//	dataMap[d.DataValues[0].Value[i]] = []string{}
	//}
	//
	//fmt.Println(dataMap)
	//
	//for i := 1; i < len(d.DataValues); i++ {
	//	for _, value := range d.DataValues[i].Value {
	//
	//	}
	//}
}

func recursiveGetData(vs []string) (data [][]string) {
	data = append(data, recursiveGetData(vs)...)
	return
}
