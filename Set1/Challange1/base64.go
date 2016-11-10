package main
import (
	"fmt"
	"os"
	"encoding/hex"
	b64 "encoding/base64"
)

func main() {

	for _, arg := range os.Args[1:] {
		ascii, err := hex.DecodeString( arg )
		if err != nil {
			fmt.Println( "Error: %v\n", err )
		}
		fmt.Println( string(b64.StdEncoding.EncodeToString( ascii )) )
		fmt.Println()
	}
}
