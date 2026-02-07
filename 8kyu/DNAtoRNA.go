package main
import "fmt"

func DNAtoRNA(dna string) string{
	rna := ""
	for i := 0; i < len(dna); i++{
		if dna[i] != 'T'{
			rna += string(dna[i])
		}else{
			rna += "U"
		}
	}

	return rna
}
