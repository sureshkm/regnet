package regnet

import (
	//"bufio"
	"fmt"
	//"errors"
	"regexp"
	//"io"
	//"os"
	//"path/filepath"
	//"strings"
)

type Regnet struct {
	Patterns map[string]Pattern
}

type Pattern struct {
	raw      string
	Compiled *regexp.Regexp
}

const (
	blockIdent string = "REGNET_BLOCK"
	blockKey string = "REGNET_KEY"
)

func New() (r *Regnet, err error) {
	//const blockIdent string = "REGNET_BLOCK"
	regentBlock, err := regexp.Compile(`\%{([^}]+)\}`)
	if err != nil {
		return nil, err
	}
	blockPattern := Pattern{blockIdent, regentBlock}
	patterns := make(map[string]Pattern)
	patterns[blockIdent] = blockPattern

	regentKey, err := regexp.Compile(`[\w]+`)
	if err != nil {
		return nil, err
	}
	keyPattern := Pattern{blockKey, regentKey}
	patterns[blockKey] = keyPattern
	return &Regnet{patterns}, nil
}

func (regnet *Regnet) AddPattern(name, pattern string) (err error) {
	r := regnet.Patterns[blockIdent].Compiled
	slices := r.FindAllString(pattern, -1)
	for indx := range slices {
		//fmt.Println(">>" ,regnet.Patterns[blockKey].Compiled.FindString(slices[indx]))

		key := regnet.Patterns[blockKey].Compiled.FindString(slices[indx])
		//fmt.Println("KEY > ", key)
		raw := regnet.Patterns[key].raw
		//fmt.Println("RAW > ", raw)
		if  raw != "" {
			fmt.Println("RAW NOT NULL")
		}else{
			fmt.Println("RAW NULL")
		}
	}

	//  contains only Regnet, so get the value and compile it
	//fmt.Println("slice :", slices, " name " , name ," pattern " , pattern)
	compiled, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	patternCompiled := Pattern{name, compiled}
	regnet.Patterns[name] = patternCompiled
	fmt.Println(regnet.Patterns)

	return nil
	
}


/*func (regnet *Regnet) AddPattern(name, pattern string) (err error) {
	var unResolvedRegnetBlck = regnet.Patterns[blockIdent].Compiled.FindString(pattern)

	for {

		fmt.Println(" unResolvedRegnetBlck >" + unResolvedRegnetBlck)

		var unResolvedRegnetKey = regnet.Patterns[blockKey].Compiled.FindString(unResolvedRegnetBlck)
		fmt.Println(" unResolvedRegnetKey >" + unResolvedRegnetKey)
	}

	///do it here

	compiled, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	regnet.Patterns[name] = Pattern{pattern, compiled}
	return nil
}*/

//regnets.MatchInText("Tue May 15 11:21:42 [conn1047685] moveChunk deleted: 7157")
func (regnet *Regnet) MatchInText(text string) {

}

/*func (regnet *Regnet) GetPattern(name string) (val string, present bool) {
	pattern, present := regnet.patterns[name]
	return pattern, present
}

//Get the pattern from Regnet and compile and replace value (pattern) with compiled result!
// OR
// Create a graph with with compiled regex pattern with relation(name) across
func (regnet *Regnet) Compile(name string) (compiled string, err error) {
	pattern, present := regnet.GetPattern(name)
	if present == false {
		return "", errors.New("regnet: pattern" + name + "not found")
	}

	return pattern, nil


}*/

func main() {
	fmt.Println("here")
}

/*func Patterns(path string) *Regnet {
	_, err := filepath.Glob(path)
	if err == nil {

		}else{
			fmt.Fprintln(os.Stderr, err)
		}
}*/

/*func init() {
	files, err := filepath.Glob("/home/msi/Downloads/logstash-1.4.0/patterns/*")
	if err == nil {
		for file := range files {
			fmt.Println(files[file])

			if patternFile, err := os.Open(files[file]); err == nil {
				defer patternFile.Close()
				reader := bufio.NewReader(patternFile)
				eof := false

				for !eof {
					var line string
					line, err = reader.ReadString('\n')
					if len(line) > 1 && strings.HasPrefix(line,"#") == false {
						fields := strings.Fields(line)

						fmt.Println("line >", fields[0] ,len(line), " :> " , line)
					}

					if err == io.EOF {
            			eof = true
					}

				}
			}
		}
	} else {
		fmt.Fprintln(os.Stderr, err)
	}
}*/
