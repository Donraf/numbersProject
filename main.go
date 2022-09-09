package main

import "fmt"

func main() {
	//var expression string = "4+5*(1+2)*(9+3)/3"
	var expression string = "-10+(-20+-30)-4+(5+(6-7)+8)"
	var operators string = "+-*/"
	var newStack Stack
	var outString []string
	var number string = ""
	var prev string = ""
	var flag bool = false
	fmt.Println(expression, operators)
	for _, char := range expression {
		if prev == "-" ||
			prev == "+" ||
			prev == "*" ||
			prev == "/" ||
			prev == "^" ||
			prev == "%" ||
			prev == "(" ||
			prev == "" {
			flag = true
		} else {
			flag = false
		}

		if char >= '0' && char <= '9' {
			number = number + string(char)
		} else {
			if flag && char == '-' {
				number = number + string(char)
			} else if number != "" {
				outString = append(outString, number)
				number = ""
			}
		}

		switch char {
		case '(':
			newStack.Append(string(char))
		case ')':
			{
				length := newStack.Len()
				for i := length; i > 0; i-- {
					outChar := newStack.Pop()
					if outChar == "(" {
						break
					} else {
						outString = append(outString[:], outChar)
					}
				}
			}
		case '-':
			{
				if flag {
					continue
				}
				length := newStack.Len()
				for i := length - 1; i >= 0; i-- {
					if newStack.Get(newStack.Len()-1) != "(" {
						outString = append(outString, newStack.Pop())
					} else {
						newStack.Pop()
						break
					}
				}
				newStack.Append(string(char))
			}
		case '+':
			{
				length := newStack.Len()
				for i := length - 1; i >= 0; i-- {
					if newStack.Get(newStack.Len()-1) != "(" {
						outString = append(outString, newStack.Pop())
					} else {
						break
					}
				}
				newStack.Append(string(char))
			}
		}
		prev = string(char)
	}
	length := newStack.Len()
	for i := length - 1; i >= 0; i-- {
		outString = append(outString, newStack.Pop())
	}
	fmt.Println(outString)
}
