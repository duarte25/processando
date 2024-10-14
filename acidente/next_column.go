package accident

import "strings"

// nextColumn retorna a próxima coluna a partir do índice atual.
// Ela também atualiza o índice atual para a próxima posição.
func nextColumn(line string, idx *int, sep string) string {
	startPos := *idx
	endPos := strings.Index(line[startPos:], sep)

	if endPos == -1 {
		*idx = len(line) // Atualiza o índice para o final da linha
		return line[startPos:]
	} else {
		*idx = startPos + endPos + 1 // Atualiza o índice para a próxima posição após o separador
		return line[startPos : startPos+endPos]
	}
}
