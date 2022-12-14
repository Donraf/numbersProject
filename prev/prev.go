package prev

import "fmt"

type Operator struct {
	Type     string
	Number   int
	Priority int
}

func GetOperators(operators *string, arr *string) []Operator {
	operatorsSlice := make([]Operator, 0)
	var number int = 1
	var bracePriority int = 0
	var prev string = ""
	for _, elemInArr := range *arr {
		var flag bool = false
		if elemInArr == '(' {
			bracePriority += 2
		} else if elemInArr == ')' {
			bracePriority -= 2
		}
		for _, operator := range *operators {
			if elemInArr == operator {
				var basePriority int = 0
				switch elemInArr {
				case '+':
					basePriority = 0
				case '-':
					for _, operator2 := range *operators {
						if prev == string(operator2) {
							flag = true
						}
					}
					if prev == "(" {
						flag = true
					}
					basePriority = 0
				case '*':
					basePriority = 1
				case '/':
					basePriority = 1
				case '%':
					basePriority = 1
				case '^':
					basePriority = 1
				}
				if !flag {
					var newOperator Operator = Operator{
						Type:     string(elemInArr),
						Number:   number,
						Priority: basePriority + bracePriority,
					}
					number++
					operatorsSlice = append(operatorsSlice, newOperator)
				}
			}
		}
		prev = string(elemInArr)
	}
	SortOperators(&operatorsSlice)
	return operatorsSlice
}

func SortOperators(operators *[]Operator) {
	var operatorsNum int = len(*operators)
	for i := 1; i < operatorsNum; i++ {
		var operatorPriotiry int = (*operators)[i].Priority
		for j := i - 1; j >= 0; j-- {
			if operatorPriotiry > (*operators)[j].Priority {
				(*operators)[j], (*operators)[j+1] = (*operators)[j+1], (*operators)[j]
			}
		}
	}
}

func GetOperands(operators *string, arr *string) []string {
	var prev string = ""
	var operands []string
	for _, char := range *arr {
		var flag bool = false
		if char >= '0' && char <= '9' {
			fmt.Print(char-'0', ' ')
		}
		if char == '-' {
			for _, operator2 := range *operators {
				if prev == string(operator2) {
					flag = true
					break
				}
			}
			if prev == "(" {
				flag = true
			}
			if flag {
				fmt.Print('-', ' ')
			}
		}
		prev = string(char)
	}
	return operands
}

func main() {
	var operatorsSet string = "-+*/%^"
	//var expression string = "1+2*9+2*(8+3)"
	var expression string = "1*2(3*(4+(-5))*(5+6))*7"
	var operatorsSlice []Operator
	operatorsSlice = GetOperators(&operatorsSet, &expression)
	operandsSlice := make([]int, len(operatorsSlice)+1)
	fmt.Println(operatorsSlice)
	fmt.Println(operandsSlice)
	GetOperands(&operatorsSet, &expression)
	//for i := 0; i < len(expression); i++ {
	//	char := expression[i]
	//	if char >= '0' && char <= '9' {
	//		fmt.Println(char - '0')
	//	}
	//}
}
