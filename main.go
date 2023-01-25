package main

import (
	"fmt"
	"strconv"
	"log"
	"os"
	"bufio"
	"strings"
)

func main(){
	var input string
	fmt.Println("How many bits should each network mask be? (enter a value between 1 and 32)")
	fmt.Scanln(&input)

	n, err := strconv.Atoi(input)
	if err != nil{
		log.Fatal(err)
	} 
	if n < 1 || n > 32{
		log.Fatal("number must be between 1 and 32")
	}
	readFile(n)
}

func readFile(bits int){
	var binary_ips []string
	file, err := os.Open("ips.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		line := scanner.Text()
		binary_ips = append(binary_ips,convertToBinary(line))
	}
	
	subnet_masks := common(binary_ips, bits)
	for i := range subnet_masks{
		fmt.Println(convertToDecimal(subnet_masks[i]))
	}
}

func convertToBinary(ip string) string{
	octets := strings.Split(ip, ".")
	var binary_ip string
	for i := range octets{
		n, err := strconv.Atoi(octets[i])
		if err != nil{
			log.Fatal(err)
		}
		binary_ip += fmt.Sprintf("%08b",n)
	}
	fmt.Println(binary_ip)
	return binary_ip
}

func common(ips []string, bits int) []string{
	var subnet_masks []string
	for i := range ips{
		mask := ips[i][0:bits]
		if !found(mask, subnet_masks){
			subnet_masks = append(subnet_masks, mask)
		}
	}
	return subnet_masks
}

func found(mask string, masks []string) bool{
	for i := range masks{
		if mask == masks[i]{
			return true
		}
	}
	return false
}

func convertToDecimal(m string) string{
	bits := len(m)
	bm := addDotsZeroes(m)
	dm := ""

	octets := strings.Split(bm, ".")
	for i := range octets{
		n, err := strconv.ParseInt(octets[i], 2, 10)
		if err != nil{
			log.Fatal(err)
		}
		if i==0{
			dm += fmt.Sprintf("%d",n)
		} else{
			dm += fmt.Sprintf(".%d",n)
		}
	}
	return fmt.Sprintf("%s /%d",dm,bits)
}

func addDotsZeroes(ip string) string{
	zeroes := "00000000.00000000.00000000.00000000"
	address := ""

	for i := range ip{
		if i%8==0 && i!=0{
			address += "."
		}
		address += string(ip[i])
	}
	return address + zeroes[len(address):]
}

