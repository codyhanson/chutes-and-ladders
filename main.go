package main

import "fmt"
import "math/rand"
import "time"

func main() {

	seed := time.Now().UnixNano()
	fmt.Printf("Random seed: %d\n", seed)
	rand.Seed(seed)

	var board = []int{
	0, //starting spot
	37, // 1, ladder to 38
	0, //2
	0, //3
	10, //4, ladder to 14
	0, //5
	0, //6
	0, //7
	0, //8
	22, //9, ladder to 31
	0, //10
	0, //11
	0, //12
	0, //13
	0, //14
	0, //15
	-10, //16, chute to 6
	0, //17
	0, //18
	0, //19
	0, //20
	21,//21, ladder to 42
	0, //22
	0, //23
	0, //24
	0, //25
	0, //26
	0, //27
	56, //28, ladder to 84,
	0, //29
	0, //30
	0, //31
	0, //32
	0, //33
	0, //34
	0, //35
	8, //36, ladder to 44,
	0, //37
	0, //38
	0, //39
	0, //40
	0, //41
	0, //42
	0, //43
	0, //44
	0, //45
	0, //46
	-21, //47, chute to 26
	0, //48
	-38, //49, chute to 11
	0, //50
	16, //51, ladder to 67
	0, //52
	0, //53
	0, //54
	0, //55
	-3, //56, chute to 53
	0, //57
	0, //58
	0, //59
	0, //60
	0, //61
	-43, //62, chute to 19
	0, //63
	-4, //64, chute to 60
	0, //65
	0, //66
	0, //67
	0, //68
	0, //69
	0, //70
	20, //71, ladder to 91
	0, //72
	0, //73
	0, //74
	0, //75
	0, //76
	0, //77
	0, //78
	0, //79
	20, //80
	0, //81
	0, //82
	0, //83
	0, //84
	0, //85
	0, //86
	-63, //87, chute to 24
	0, //88
	0, //89
	0, //90
	0, //91
	0, //92
	-20, //93, chute to 73
	0, //94
	-20, //95, chute to 75
	0, //96
	0, //97
	-20, //98, chute to 78
	0, //99
	0, //100!
	}

	moveCounter := 0
	i := 0
	for ;; {
		moveCounter += 1
		fmt.Printf("------\nMove %d, currently at %d\n", moveCounter, i)
		spin := spinner()
		fmt.Printf("spun %d\n", spin)
		if (i + spin) >= len(board) {
			break
		}
		i += spin
		cl := board[i]
		if cl > 0 {
			fmt.Printf("Ladder! %d\n", cl)
		}
		if cl < 0 {
			fmt.Printf("chute! %d\n", cl)
		}
		i += cl
		if (i + spin) >= len(board) {
			break
		}
	}

	fmt.Printf("You won in %d moves!\n", moveCounter)

}

func spinner() int {
	return rand.Intn(6) + 1
}
