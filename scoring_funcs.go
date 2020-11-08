package main

var DNAFull = ScoringFunc{
	'A': map[uint8]int{
		'A': 5,
		'T': -4,
		'G': -4,
		'C': -4,
	},
	'T': map[uint8]int{
		'A': -4,
		'T': 5,
		'G': -4,
		'C': -4,
	},
	'G': map[uint8]int{
		'A': -4,
		'T': -4,
		'G': 5,
		'C': -4,
	},
	'C': map[uint8]int{
		'A': -4,
		'T': -4,
		'G': -4,
		'C': 5,
	},
}

var SimpleFunc = ScoringFunc{
	'A': map[uint8]int{
		'A':  1 ,
		'R':  -1 ,
		'N':  -1 ,
		'D':  -1 ,
		'C':  -1 ,
		'Q':  -1 ,
		'E':  -1 ,
		'G':  -1 ,
		'H':  -1 ,
		'I':  -1 ,
		'L':  -1 ,
		'K':  -1 ,
		'M':  -1 ,
		'F':  -1 ,
		'P':  -1 ,
		'S':  -1 ,
		'T':  -1 ,
		'W':  -1 ,
		'Y':  -1 ,
		'V':  -1 ,
	},
	'R': map[uint8]int{
		'A':  -1 ,
		'R':  1 ,
		'N':  -1 ,
		'D':  -1 ,
		'C':  -1 ,
		'Q':  -1 ,
		'E':  -1 ,
		'G':  -1 ,
		'H':  -1 ,
		'I':  -1 ,
		'L':  -1 ,
		'K':  -1 ,
		'M':  -1 ,
		'F':  -1 ,
		'P':  -1 ,
		'S':  -1 ,
		'T':  -1 ,
		'W':  -1 ,
		'Y':  -1 ,
		'V':  -1 ,
	},
	'N': map[uint8]int{
		'A':  -1 ,
		'R':  -1 ,
		'N':  1 ,
		'D':  -1 ,
		'C':  -1 ,
		'Q':  -1 ,
		'E':  -1 ,
		'G':  -1 ,
		'H':  -1 ,
		'I':  -1 ,
		'L':  -1 ,
		'K':  -1 ,
		'M':  -1 ,
		'F':  -1 ,
		'P':  -1 ,
		'S':  -1 ,
		'T':  -1 ,
		'W':  -1 ,
		'Y':  -1 ,
		'V':  -1 ,
	},
	'D': map[uint8]int{
		'A':  -1 ,
		'R':  -1 ,
		'N':  -1 ,
		'D':  1 ,
		'C':  -1 ,
		'Q':  -1 ,
		'E':  -1 ,
		'G':  -1 ,
		'H':  -1 ,
		'I':  -1 ,
		'L':  -1 ,
		'K':  -1 ,
		'M':  -1 ,
		'F':  -1 ,
		'P':  -1 ,
		'S':  -1 ,
		'T':  -1 ,
		'W':  -1 ,
		'Y':  -1 ,
		'V':  -1 ,
	},
	'C': map[uint8]int{
		'A':  -1 ,
		'R':  -1 ,
		'N':  -1 ,
		'D':  -1 ,
		'C':  1 ,
		'Q':  -1 ,
		'E':  -1 ,
		'G':  -1 ,
		'H':  -1 ,
		'I':  -1 ,
		'L':  -1 ,
		'K':  -1 ,
		'M':  -1 ,
		'F':  -1 ,
		'P':  -1 ,
		'S':  -1 ,
		'T':  -1 ,
		'W':  -1 ,
		'Y':  -1 ,
		'V':  -1 ,
	},
	'Q': map[uint8]int{
		'A':  -1 ,
		'R':  -1 ,
		'N':  -1 ,
		'D':  -1 ,
		'C':  -1 ,
		'Q':  1 ,
		'E':  -1 ,
		'G':  -1 ,
		'H':  -1 ,
		'I':  -1 ,
		'L':  -1 ,
		'K':  -1 ,
		'M':  -1 ,
		'F':  -1 ,
		'P':  -1 ,
		'S':  -1 ,
		'T':  -1 ,
		'W':  -1 ,
		'Y':  -1 ,
		'V':  -1 ,
	},
	'E': map[uint8]int{
		'A':  -1 ,
		'R':  -1 ,
		'N':  -1 ,
		'D':  -1 ,
		'C':  -1 ,
		'Q':  -1 ,
		'E':  1 ,
		'G':  -1 ,
		'H':  -1 ,
		'I':  -1 ,
		'L':  -1 ,
		'K':  -1 ,
		'M':  -1 ,
		'F':  -1 ,
		'P':  -1 ,
		'S':  -1 ,
		'T':  -1 ,
		'W':  -1 ,
		'Y':  -1 ,
		'V':  -1 ,
	},
	'G': map[uint8]int{
		'A':  -1 ,
		'R':  -1 ,
		'N':  -1 ,
		'D':  -1 ,
		'C':  -1 ,
		'Q':  -1 ,
		'E':  -1 ,
		'G':  1 ,
		'H':  -1 ,
		'I':  -1 ,
		'L':  -1 ,
		'K':  -1 ,
		'M':  -1 ,
		'F':  -1 ,
		'P':  -1 ,
		'S':  -1 ,
		'T':  -1 ,
		'W':  -1 ,
		'Y':  -1 ,
		'V':  -1 ,
	},
	'H': map[uint8]int{
		'A':  -1 ,
		'R':  -1 ,
		'N':  -1 ,
		'D':  -1 ,
		'C':  -1 ,
		'Q':  -1 ,
		'E':  -1 ,
		'G':  -1 ,
		'H':  1 ,
		'I':  -1 ,
		'L':  -1 ,
		'K':  -1 ,
		'M':  -1 ,
		'F':  -1 ,
		'P':  -1 ,
		'S':  -1 ,
		'T':  -1 ,
		'W':  -1 ,
		'Y':  -1 ,
		'V':  -1 ,
	},
	'I': map[uint8]int{
		'A':  -1 ,
		'R':  -1 ,
		'N':  -1 ,
		'D':  -1 ,
		'C':  -1 ,
		'Q':  -1 ,
		'E':  -1 ,
		'G':  -1 ,
		'H':  -1 ,
		'I':  1 ,
		'L':  -1 ,
		'K':  -1 ,
		'M':  -1 ,
		'F':  -1 ,
		'P':  -1 ,
		'S':  -1 ,
		'T':  -1 ,
		'W':  -1 ,
		'Y':  -1 ,
		'V':  -1 ,
	},
	'L': map[uint8]int{
		'A':  -1 ,
		'R':  -1 ,
		'N':  -1 ,
		'D':  -1 ,
		'C':  -1 ,
		'Q':  -1 ,
		'E':  -1 ,
		'G':  -1 ,
		'H':  -1 ,
		'I':  -1 ,
		'L':  1 ,
		'K':  -1 ,
		'M':  -1 ,
		'F':  -1 ,
		'P':  -1 ,
		'S':  -1 ,
		'T':  -1 ,
		'W':  -1 ,
		'Y':  -1 ,
		'V':  -1 ,
	},
	'K': map[uint8]int{
		'A':  -1 ,
		'R':  -1 ,
		'N':  -1 ,
		'D':  -1 ,
		'C':  -1 ,
		'Q':  -1 ,
		'E':  -1 ,
		'G':  -1 ,
		'H':  -1 ,
		'I':  -1 ,
		'L':  -1 ,
		'K':  1 ,
		'M':  -1 ,
		'F':  -1 ,
		'P':  -1 ,
		'S':  -1 ,
		'T':  -1 ,
		'W':  -1 ,
		'Y':  -1 ,
		'V':  -1 ,
	},
	'M': map[uint8]int{
		'A':  -1 ,
		'R':  -1 ,
		'N':  -1 ,
		'D':  -1 ,
		'C':  -1 ,
		'Q':  -1 ,
		'E':  -1 ,
		'G':  -1 ,
		'H':  -1 ,
		'I':  -1 ,
		'L':  -1 ,
		'K':  -1 ,
		'M':  1 ,
		'F':  -1 ,
		'P':  -1 ,
		'S':  -1 ,
		'T':  -1 ,
		'W':  -1 ,
		'Y':  -1 ,
		'V':  -1 ,
	},
	'F': map[uint8]int{
		'A':  -1 ,
		'R':  -1 ,
		'N':  -1 ,
		'D':  -1 ,
		'C':  -1 ,
		'Q':  -1 ,
		'E':  -1 ,
		'G':  -1 ,
		'H':  -1 ,
		'I':  -1 ,
		'L':  -1 ,
		'K':  -1 ,
		'M':  -1 ,
		'F':  1 ,
		'P':  -1 ,
		'S':  -1 ,
		'T':  -1 ,
		'W':  -1 ,
		'Y':  -1 ,
		'V':  -1 ,
	},
	'P': map[uint8]int{
		'A':  -1 ,
		'R':  -1 ,
		'N':  -1 ,
		'D':  -1 ,
		'C':  -1 ,
		'Q':  -1 ,
		'E':  -1 ,
		'G':  -1 ,
		'H':  -1 ,
		'I':  -1 ,
		'L':  -1 ,
		'K':  -1 ,
		'M':  -1 ,
		'F':  -1 ,
		'P':  1 ,
		'S':  -1 ,
		'T':  -1 ,
		'W':  -1 ,
		'Y':  -1 ,
		'V':  -1 ,
	},
	'S': map[uint8]int{
		'A':  -1 ,
		'R':  -1 ,
		'N':  -1 ,
		'D':  -1 ,
		'C':  -1 ,
		'Q':  -1 ,
		'E':  -1 ,
		'G':  -1 ,
		'H':  -1 ,
		'I':  -1 ,
		'L':  -1 ,
		'K':  -1 ,
		'M':  -1 ,
		'F':  -1 ,
		'P':  -1 ,
		'S':  1 ,
		'T':  -1 ,
		'W':  -1 ,
		'Y':  -1 ,
		'V':  -1 ,
	},
	'T': map[uint8]int{
		'A':  -1 ,
		'R':  -1 ,
		'N':  -1 ,
		'D':  -1 ,
		'C':  -1 ,
		'Q':  -1 ,
		'E':  -1 ,
		'G':  -1 ,
		'H':  -1 ,
		'I':  -1 ,
		'L':  -1 ,
		'K':  -1 ,
		'M':  -1 ,
		'F':  -1 ,
		'P':  -1 ,
		'S':  -1 ,
		'T':  1 ,
		'W':  -1 ,
		'Y':  -1 ,
		'V':  -1 ,
	},
	'W': map[uint8]int{
		'A':  -1 ,
		'R':  -1 ,
		'N':  -1 ,
		'D':  -1 ,
		'C':  -1 ,
		'Q':  -1 ,
		'E':  -1 ,
		'G':  -1 ,
		'H':  -1 ,
		'I':  -1 ,
		'L':  -1 ,
		'K':  -1 ,
		'M':  -1 ,
		'F':  -1 ,
		'P':  -1 ,
		'S':  -1 ,
		'T':  -1 ,
		'W':  1 ,
		'Y':  -1 ,
		'V':  -1 ,
	},
	'Y': map[uint8]int{
		'A':  -1 ,
		'R':  -1 ,
		'N':  -1 ,
		'D':  -1 ,
		'C':  -1 ,
		'Q':  -1 ,
		'E':  -1 ,
		'G':  -1 ,
		'H':  -1 ,
		'I':  -1 ,
		'L':  -1 ,
		'K':  -1 ,
		'M':  -1 ,
		'F':  -1 ,
		'P':  -1 ,
		'S':  -1 ,
		'T':  -1 ,
		'W':  -1 ,
		'Y':  1 ,
		'V':  -1 ,
	},
	'V': map[uint8]int{
		'A':  -1 ,
		'R':  -1 ,
		'N':  -1 ,
		'D':  -1 ,
		'C':  -1 ,
		'Q':  -1 ,
		'E':  -1 ,
		'G':  -1 ,
		'H':  -1 ,
		'I':  -1 ,
		'L':  -1 ,
		'K':  -1 ,
		'M':  -1 ,
		'F':  -1 ,
		'P':  -1 ,
		'S':  -1 ,
		'T':  -1 ,
		'W':  -1 ,
		'Y':  -1 ,
		'V':  1 ,
	},
}

var Blosum62 = ScoringFunc{
	'A': map[uint8]int{
		'A': 4,
		'R': -1,
		'N': -2,
		'D': -2,
		'C': 0,
		'Q': -1,
		'E': -1,
		'G': 0,
		'H': -2,
		'I': -1,
		'L': -1,
		'K': -1,
		'M': -1,
		'F': -2,
		'P': -1,
		'S': 1,
		'T': 0,
		'W': -3,
		'Y': -2,
		'V': 0,
	},
	'R': map[uint8]int{
		'A': -1,
		'R': 5,
		'N': 0,
		'D': -2,
		'C': -3,
		'Q': 1,
		'E': 0,
		'G': -2,
		'H': 0,
		'I': -3,
		'L': -2,
		'K': 2,
		'M': -1,
		'F': -3,
		'P': -2,
		'S': -1,
		'T': -1,
		'W': -3,
		'Y': -2,
		'V': -3,
	},
	'N': map[uint8]int{
		'A': -2,
		'R': 0,
		'N': 6,
		'D': 1,
		'C': -3,
		'Q': 0,
		'E': 0,
		'G': 0,
		'H': 1,
		'I': -3,
		'L': -3,
		'K': 0,
		'M': -2,
		'F': -3,
		'P': -2,
		'S': 1,
		'T': 0,
		'W': -4,
		'Y': -2,
		'V': -3,
	},
	'D': map[uint8]int{
		'A': -2,
		'R': -2,
		'N': 1,
		'D': 6,
		'C': -3,
		'Q': 0,
		'E': 2,
		'G': -1,
		'H': -1,
		'I': -3,
		'L': -4,
		'K': -1,
		'M': -3,
		'F': -3,
		'P': -1,
		'S': 0,
		'T': -1,
		'W': -4,
		'Y': -3,
		'V': -3,
	},
	'C': map[uint8]int{
		'A': 0,
		'R': -3,
		'N': -3,
		'D': -3,
		'C': 9,
		'Q': -3,
		'E': -4,
		'G': -3,
		'H': -3,
		'I': -1,
		'L': -1,
		'K': -3,
		'M': -1,
		'F': -2,
		'P': -3,
		'S': -1,
		'T': -1,
		'W': -2,
		'Y': -2,
		'V': -1,
	},
	'Q': map[uint8]int{
		'A': -1,
		'R': 1,
		'N': 0,
		'D': 0,
		'C': -3,
		'Q': 5,
		'E': 2,
		'G': -2,
		'H': 0,
		'I': -3,
		'L': -2,
		'K': 1,
		'M': 0,
		'F': -3,
		'P': -1,
		'S': 0,
		'T': -1,
		'W': -2,
		'Y': -1,
		'V': -2,
	},
	'E': map[uint8]int{
		'A': -1,
		'R': 0,
		'N': 0,
		'D': 2,
		'C': -4,
		'Q': 2,
		'E': 5,
		'G': -2,
		'H': 0,
		'I': -3,
		'L': -3,
		'K': 1,
		'M': -2,
		'F': -3,
		'P': -1,
		'S': 0,
		'T': -1,
		'W': -3,
		'Y': -2,
		'V': -2,
	},
	'G': map[uint8]int{
		'A': 0,
		'R': -2,
		'N': 0,
		'D': -1,
		'C': -3,
		'Q': -2,
		'E': -2,
		'G': 6,
		'H': -2,
		'I': -4,
		'L': -4,
		'K': -2,
		'M': -3,
		'F': -3,
		'P': -2,
		'S': 0,
		'T': -2,
		'W': -2,
		'Y': -3,
		'V': -3,
	},
	'H': map[uint8]int{
		'A': -2,
		'R': 0,
		'N': 1,
		'D': -1,
		'C': -3,
		'Q': 0,
		'E': 0,
		'G': -2,
		'H': 8,
		'I': -3,
		'L': -3,
		'K': -1,
		'M': -2,
		'F': -1,
		'P': -2,
		'S': -1,
		'T': -2,
		'W': -2,
		'Y': 2,
		'V': -3,
	},
	'I': map[uint8]int{
		'A': -1,
		'R': -3,
		'N': -3,
		'D': -3,
		'C': -1,
		'Q': -3,
		'E': -3,
		'G': -4,
		'H': -3,
		'I': 4,
		'L': 2,
		'K': -3,
		'M': 1,
		'F': 0,
		'P': -3,
		'S': -2,
		'T': -1,
		'W': -3,
		'Y': -1,
		'V': 3,
	},
	'L': map[uint8]int{
		'A': -1,
		'R': -2,
		'N': -3,
		'D': -4,
		'C': -1,
		'Q': -2,
		'E': -3,
		'G': -4,
		'H': -3,
		'I': 2,
		'L': 4,
		'K': -2,
		'M': 2,
		'F': 0,
		'P': -3,
		'S': -2,
		'T': -1,
		'W': -2,
		'Y': -1,
		'V': 1,
	},
	'K': map[uint8]int{
		'A': -1,
		'R': 2,
		'N': 0,
		'D': -1,
		'C': -3,
		'Q': 1,
		'E': 1,
		'G': -2,
		'H': -1,
		'I': -3,
		'L': -2,
		'K': 5,
		'M': -1,
		'F': -3,
		'P': -1,
		'S': 0,
		'T': -1,
		'W': -3,
		'Y': -2,
		'V': -2,
	},
	'M': map[uint8]int{
		'A': -1,
		'R': -1,
		'N': -2,
		'D': -3,
		'C': -1,
		'Q': 0,
		'E': -2,
		'G': -3,
		'H': -2,
		'I': 1,
		'L': 2,
		'K': -1,
		'M': 5,
		'F': 0,
		'P': -2,
		'S': -1,
		'T': -1,
		'W': -1,
		'Y': -1,
		'V': 1,
	},
	'F': map[uint8]int{
		'A': -2,
		'R': -3,
		'N': -3,
		'D': -3,
		'C': -2,
		'Q': -3,
		'E': -3,
		'G': -3,
		'H': -1,
		'I': 0,
		'L': 0,
		'K': -3,
		'M': 0,
		'F': 6,
		'P': -4,
		'S': -2,
		'T': -2,
		'W': 1,
		'Y': 3,
		'V': -1,
	},
	'P': map[uint8]int{
		'A': -1,
		'R': -2,
		'N': -2,
		'D': -1,
		'C': -3,
		'Q': -1,
		'E': -1,
		'G': -2,
		'H': -2,
		'I': -3,
		'L': -3,
		'K': -1,
		'M': -2,
		'F': -4,
		'P': 7,
		'S': -1,
		'T': -1,
		'W': -4,
		'Y': -3,
		'V': -2,
	},
	'S': map[uint8]int{
		'A': 1,
		'R': -1,
		'N': 1,
		'D': 0,
		'C': -1,
		'Q': 0,
		'E': 0,
		'G': 0,
		'H': -1,
		'I': -2,
		'L': -2,
		'K': 0,
		'M': -1,
		'F': -2,
		'P': -1,
		'S': 4,
		'T': 1,
		'W': -3,
		'Y': -2,
		'V': -2,
	},
	'T': map[uint8]int{
		'A': 0,
		'R': -1,
		'N': 0,
		'D': -1,
		'C': -1,
		'Q': -1,
		'E': -1,
		'G': -2,
		'H': -2,
		'I': -1,
		'L': -1,
		'K': -1,
		'M': -1,
		'F': -2,
		'P': -1,
		'S': 1,
		'T': 5,
		'W': -2,
		'Y': -2,
		'V': 0,
	},
	'W': map[uint8]int{
		'A': -3,
		'R': -3,
		'N': -4,
		'D': -4,
		'C': -2,
		'Q': -2,
		'E': -3,
		'G': -2,
		'H': -2,
		'I': -3,
		'L': -2,
		'K': -3,
		'M': -1,
		'F': 1,
		'P': -4,
		'S': -3,
		'T': -2,
		'W': 11,
		'Y': 2,
		'V': -3,
	},
	'Y': map[uint8]int{
		'A': -2,
		'R': -2,
		'N': -2,
		'D': -3,
		'C': -2,
		'Q': -1,
		'E': -2,
		'G': -3,
		'H': 2,
		'I': -1,
		'L': -1,
		'K': -2,
		'M': -1,
		'F': 3,
		'P': -3,
		'S': -2,
		'T': -2,
		'W': 2,
		'Y': 7,
		'V': -1,
	},
	'V': map[uint8]int{
		'A': 0,
		'R': -3,
		'N': -3,
		'D': -3,
		'C': -1,
		'Q': -2,
		'E': -2,
		'G': -3,
		'H': -3,
		'I': 3,
		'L': 1,
		'K': -2,
		'M': 1,
		'F': -1,
		'P': -2,
		'S': -2,
		'T': 0,
		'W': -3,
		'Y': -1,
		'V': 4,
	},
}
