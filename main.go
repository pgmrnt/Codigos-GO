package main

import "fmt"

func main() {

	var input string

	fmt.Println("Qual e Seu nome?")
	fmt.Scanln(&input) //Ele esta pedindo seu nome

	fmt.Println("Voce digitou:", input)

	if input == "Renato" { //aqi esta uma condiçao, se sim, ou se nao.
		fmt.Println("Sim! Este e o Nome certo.") //Se o nome que voce digitou esta certo ira aparecer esta msg.

	} else {
		fmt.Println("O nome nao e Este!") //Se o nome que voce digitou estiver errado ira aparecer esta msg.
	}
	var tentativas int      //este cara vai guardar as tentativas que voce tentou.
	for input != "Renato" { //Este cara aqui e um laco ele vai repetir ate voce certa o nome.
		fmt.Scanln(&input)
		tentativas++

		fmt.Println("Digite novamente!")
	}
	fmt.Printf("Sim! \"%s\" é o nome certo! Você tentou %d vezes.\n", input, tentativas)
	//este cara aqui ira formata para fica facil de entender!
}
