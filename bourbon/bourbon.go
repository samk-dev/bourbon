package bourbon

import (
	"fmt"

	"github.com/joho/godotenv"
)

// const version = "1.0.0"

type Bourbon struct {
	AppName string
	Debug   bool
	Version string
}

func (b *Bourbon) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:     rootPath,
		foldersNames: []string{"handlers", "migrations", "views", "data", "public", "temp", "logs", "middlware"},
	}

	err := b.Init(pathConfig)
	if err != nil {
		return err
	}

	err = b.CheckDotEnv(rootPath)
	if err != nil {
		return err
	}

	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}

	return nil
}

func (b *Bourbon) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.foldersNames {
		err := b.CreateDirIfDoesntExist(root + "/" + path)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Bourbon) CheckDotEnv(path string) error {
	err := b.CreateFileIfDoesntExists(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}

	return nil
}
