package util

import (
	"github.com/joho/godotenv"
	//"fmt"
	//"io"
	//"io/ioutil"
	//"bufio"
	//"os"
)

func catch(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func LoadConfig() {
	err := godotenv.Load("./.env")
	//config, err := os.Open("./.env")
	//ioutil.ReadFile
	catch(err)

	//fmt.Println(config)

	// defer config.Close()
	// scanner := bufio.NewScanner(config)
	// scanner.Split(bufio.ScanWords)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }
}
