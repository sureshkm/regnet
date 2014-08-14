package main

//import "github.com/anaray/regnet"
import (
	//"fmt"
	"github.com/anaray/regnet"
)

func main() {

	/*reg := Reg()
	reg.AddPatternFromFile("/home/msi/Downloads/logstash-1.4.0/patterns/*")
	text := "Tue May 15 11:21:42 [conn1047685] moveChunk deleted: 7157"
	pattern := %{DAY}
	reg.Compile(pattern)
	match := reg.Match(text)

	captures := match.Captures()
	*/

	//regexes := regnet.AddPatternsFromFile("/home/msi/Downloads/logstash-1.4.0/patterns/*")
	/*regnets := regnet.LoadDefinitions("/home/msi/Downloads/logstash-1.4.0/patterns/*")
	pattern := "%{DAY}"

	regnets.Get(pattern)

	Matches matches := regnets.MatchInText("Tue May 15 11:21:42 [conn1047685] moveChunk deleted: 7157")
	for matches.Step() {
		if err := matches.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		fmt.Println(matches.Get())
	}*/

	r, _ := regnet.New()
	r.AddPattern("DAY", `(?:Mon(?:day)?|Tue(?:sday)?|Wed(?:nesday)?|Thu(?:rsday)?|Fri(?:day)?|Sat(?:urday)?|Sun(?:day)?)`)
	//r.AddPattern("YEAR", `(\d\d){1,2}`)
	//r.AddPattern("TEXT", `%{DAY} , %{YEAR}`)

	//fmt.Println(r.Patterns["REGNET_BLOCK"].Compiled)
	//fmt.Println(r.patterns["REGNET_BLOCK"].compiled)
	//}

}
