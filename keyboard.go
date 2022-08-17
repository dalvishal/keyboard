package keyboard

import (
	
	"bufio"
	"os"
	"strings"
	"strconv"
)



func Getinput() (float64, error) {

	
	newNuma := bufio.NewReader(os.Stdin)
	newNumFina,err := newNuma.ReadString('\n')
	if err!= nil {
		return 0,err
	}
	newNumFina = strings.TrimSpace(newNumFina)
	numInput,err := strconv.ParseFloat(newNumFina,64)
	if err!= nil {
		return 0,err
	}
	return numInput,nil

}


