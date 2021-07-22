package apply //nolint: typecheck

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"bitbucket.org/git-fsrg/wikifier/internal/app"
	"github.com/fatih/color"
)

func prompt(changes []app.MetaData) bool {
	for _, md := range changes {
		switch s := md.Status; s {
		case app.STATUS_DELETED:
			color.Red("- %s", md.Origin)
		case app.STATUS_PENDING:
			color.Cyan("+ %s", md.Origin)
		}
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Apply (yes):")
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\n")
	return text == "yes" || text == "y"
}
