package text

import "github.com/jybp/render-cli/v2/pkg/tui/views"

func PSQLResultText(result *views.PSQLResult) string {
	return result.Output
}
