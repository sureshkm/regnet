package regnet

import (
	//"bufio"
	"fmt"
	"errors"
	"regexp"
	//"io"
	//"os"
	//"path/filepath"
	"strings"
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

//
func New() (r *Regnet, err error) {
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

//
func (regnet *Regnet) AddPattern(name string, pattern string) (err error) {
	if _,present := regnet.GetPattern(name); present == true {
		return errors.New("regnet: pattern " + name + " already exists.")
	}

	r := regnet.Patterns[blockIdent].Compiled
	slices := r.FindAllString(pattern, -1)

	for indx := range slices {
		key := regnet.Patterns[blockKey].Compiled.FindString(slices[indx])
		value, present := regnet.GetPattern(key)
		if present == false {
			return errors.New("regnet: pattern " + key + " not found. Define it before " + name + " regnet.")
		}else{
			//replace regent this its derefrenced pattern string
			pattern = strings.Replace(pattern, "%{" + key + "}", value.Compiled.String(),-1)
		}
	}

	//  contains only Regnet, so get the value and compile it
	compiled, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	patternCompiled := Pattern{name, compiled}
	regnet.Patterns[name] = patternCompiled
	fmt.Println(regnet.Patterns)

	return nil
}

//regnets.MatchInText("Tue May 15 11:21:42 [conn1047685] moveChunk deleted: 7157")
func (regnet *Regnet) MatchInText(text string) {

}

func (regnet *Regnet) GetPattern(name string) (pattern Pattern, present bool) {
	pattern, present = regnet.Patterns[name]
	return pattern, present
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
