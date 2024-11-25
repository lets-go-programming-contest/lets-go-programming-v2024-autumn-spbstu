//go:generate go vet ./...
package output

import (
	_ "embed"
	"fmt"

	"github.com/fatih/color"

	"github.com/mrqiz/task-8/internal/feats"
	"github.com/mrqiz/task-8/internal/version"
)

//go:embed JustBuyThisGoddamnThing.txt
var proLicenseAdvert string

func CliHeader() {
	fmt.Printf("%s (%s)\n", version.AppName, version.Version)
}

func CliTooFewArgs() {
	fmt.Println("err: not enough args")
	fmt.Println()
	cliUsage()
}

func cliUsage() {
	fmt.Printf("usage: EnterpriseCalc [opearation] [...operands]\n\n")
}

func CliOpNotFound() {
	color.Set(color.FgRed)
	fmt.Println("err: unknown feature")
	color.Unset()
	fmt.Println("available feats: ")

	for _, op := range feats.Features {
		fmt.Println(op.Name)
	}

	if version.Edition == "base" {
		cliProAdvert()
	}
}

func cliProAdvert() {
	color.Set(color.FgGreen, color.Bold)
	fmt.Print("\n", proLicenseAdvert)
	color.Unset()
}
