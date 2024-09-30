package utils

import (
	"fmt"
	"go-bt-project-structure-generator/constants"
	"os"
	"os/exec"

	"github.com/manifoldco/promptui"
)

var SupportedORMs = map[string]string{
	"GORM": `"gorm.io/gorm"`,
}

func ConfirmInstallationORM() (bool, error) {

	Items := []string{"Yes", "No"}

	prompt := promptui.Select{
		Label: "Do you want to add Gorm to project? ",
		Items: Items,
	}

	_, result, err := prompt.Run()

	if err != nil {
		return false, fmt.Errorf("selection error: %v", err)
	}

	return result == "Yes", nil
}

func InstallORM() error {

	orm := []string{"Orm"}

	cmd := exec.Command("go", "get", "-u", constants.GormPackage)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("something went wrong while intalling %s. Error: %v", orm[0], err)
	}

	return nil
}
