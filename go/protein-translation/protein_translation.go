// Package protein decodes codons to amino acids and RNAs to proteins.
package protein

import (
	"errors"
)

// ErrStop indicates that a stop codon was encountered (an unexpected stop in the case of an RNA)
var ErrStop error = errors.New("Stop")

// ErrInvalidBase indicates that an unrecognized codon was encountered.
var ErrInvalidBase error = errors.New("Invalid base")

// FromCodon converts a codon to a amino acid. It returns error ErrStop if the codon encodes a
// stop. It return ErrInvalidBase if the codon is not recognized.
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

// FromRNA converts an RNA to a protein. It stops decoding the RNA when a stop codon is encountered
// or when the end of the RNA is reached (ignoring an incomplete codon at the end if present). It
// returns ErrInvalidBase if an unrecognized codon is encountered. It returns ErrStop if a stop
// codon is encountered before any amino acid encoding codon. When an error is returned, it also
// returns the partial protein decoded thus far.
func FromRNA(rna string) ([]string, error) {
	protein := []string{}
	for codonStart := 0; codonStart < len(rna)-2; codonStart += 3 {
		codon := rna[codonStart : codonStart+3]
		aminoAcid, err := FromCodon(codon)
		switch err {
		case nil:
			protein = append(protein, aminoAcid)
		case ErrStop:
			return protein, nil
		default:
			return protein, err
		}
	}
	return protein, nil
}
