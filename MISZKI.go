package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"golang.org/x/sys/windows/registry"
)

const url = "https://www.google.com/"
const urlErr = "http://www.patriarchia.ru/" //"https://a2-22-61-43.deploy.static.akamaitechnologies.com/"

func main() {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	l := log.New(file, "", log.Ldate|log.Ltime)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1) // выходим из программы
	}
	defer file.Close() // закрываем файл

	//Проверяем наличие интернета
	switch checkWeb(url) {
	case 0:
		l.Println("Данный компьютер подключен к интернету")
		fmt.Println("Данный компьютер подключен к интернету")

		//Проверка наличия фаервола
		switch checkFirewall() {
		case 0:
			l.Println("Межсетевой экран установлен")
			fmt.Println("Межсетевой экран установлен")
			switch checkWeb(urlErr) {
			case 0:
				l.Println("Межсетевой экран функционирует неверно")
				fmt.Println("Межсетевой экран функционирует неверно")
			case 1:
				l.Println("Межсетевой экран функционирует правильно")
				fmt.Println("Межсетевой экран функционирует правильно")
			}
		case 1:
			l.Println("Межсетевой экран не установлен")
			fmt.Println("Межсетевой экран не установлен")
		}

		//Проверка наличие антивируса
		switch checkAnt() {
		case 0:
			l.Println("Установлен Windows Defender")
			fmt.Println("Установлен Windows Defender")
			l.Println("Проверка работоспособности антивирусного ПО")
			fmt.Println("Проверка работоспособности антивирусного ПО")
			checkWorckAnt()

		case 1:
			l.Println("В системе не установлен Антивирус")
			fmt.Println("В системе не установлен Антивирус")
		}

	case 1:
		l.Println("Данный компьютер не подключен к интернету")
		fmt.Println("Данный компьютер не подключен к интернету")
	}

	l.Println("Program terminated")
	//Чтобы не закрывалась консоль при завершении программы
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

}

//Проверка наличия подключения к Интернету и работу межсетевого экрана
func checkWeb(ur string) int {
	resp, err := http.Get(ur) //Отправляем Get запрос на сайт
	if err != nil {
		//		fmt.Println(err)
		return 1
	}
	defer resp.Body.Close()
	//	io.Copy(os.Stdout, resp.Body)
	return 0
}

//Проверкак наличия межсетевого экрана
func checkFirewall() int {
	_, err := ioutil.ReadFile("C:\\Windows\\System32\\WF.msc")
	if err != nil {
		return 1
	} else {
		return 0
	}
}

//Проверка наличия антивируса
func checkAnt() int {

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft`, registry.READ)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	s, err := k.ReadSubKeyNames(240)
	if err != nil {
		log.Fatal(err)
		return 1
	}
	fmt.Printf("Windows system root is %q\n", s[236])
	return 0
}

//Проверка работы антивируса
func checkWorckAnt() {

	const EICAR = "X5O!P%@AP[4\\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*"

	file, err := os.Create("EICAR.txt")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}

	file.WriteString(EICAR)
	file.Close()

	//здеся ннадо как то открыть прям файл на экране

	cmd := exec.Command("notepad", "EICAR.txt")

	err = cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

}
