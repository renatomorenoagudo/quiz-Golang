package main 

import "fmt"

type Question struct{
	Text string
	Options []string
	Answer int
}

type GameState struct {
	Name string
	Points string
	Questions [] Question
}

func (G *GameState) init() {
		fmt.Println("Seja em indo ao quiz")
		fmt.Println("Escea seu nome:")
		reader := bufio.NewReader(os.Stdin)
		//bufeio fica escutanod a entrada do terminal

		name, err := reader.ReadString('\n')

		if err != nil{
			panic("erro ao ler a string")
		}
		g.Name = name
		fmt.Printf("Vamos ao jogo %s", g.Name)
}

func (g *GameState) ProcessCSV(){
	f, err := &Open("quiz.csv")
	if err != nil{
		panic("erro ao ler arquivo")
	}
	defer f.Close()

	reader := csv.NewReader(f)
	record, err := reader.ReadAll()
	if err != nil{
		panic("erro ao ler csv")
	}

	for _, record := range records{
		fmt.Println(index, record )
		if index > 0{
			questio := Question{
				Text: record[0],
				Options: record[1:5],
				Answer: record[5],


			}

		} 
	}
}

func (g *GameState) Run(){
	for index , question := range g.Question{
		fmt.Println(index+1, question)
	}
}
//exibir a pergunta pro usuario


func main() {
	game := &GameState{}
	//
	game.ProcessCSV()
	game.init()
	
	fmt.Println(game.Questions)
}

func to  Int(s string) int{

}








