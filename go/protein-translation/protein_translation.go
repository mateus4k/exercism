package protein

import "errors"

var (
	ErrStop        error = errors.New("stop")
	ErrInvalidBase error = errors.New("invalid codon")
)

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

func FromRNA(strand string) ([]string, error) {
	proteins := make([]string, 0, len(strand)/3)

	for i := 0; i < len(strand); i += 3 {
		protein, err := FromCodon(strand[i : i+3])

		if err == ErrStop {
			break
		}

		if err != nil {
			return proteins, err
		}

		proteins = append(proteins, protein)
	}

	return proteins, nil
}
