package app //nolint: typecheck

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// @todo create interface

type MetaData struct {
	Timestamp   string
	Origin      string
	PageTitle   string
	ContentFile string
	Format      string
	Sum         string
	Status      string
	Parent      string
	MetaFile    string
	ID          string
	IsFolder    bool
}

func (c *MetaData) Load(file string) error {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return err
	}
	return nil
}

func (c *MetaData) AutoSave() *MetaData {
	return save(c, c.MetaFile)
}

func (c *MetaData) AutoDelete() {
	if c.ContentFile != "" {
		e := os.Remove(c.ContentFile)
		if e != nil {
			log.Fatal(e)
		}
	}

	e := os.Remove(c.MetaFile)
	if e != nil {
		log.Fatal(e)
	}
}

func (c *MetaData) Save(file string) *MetaData {
	return save(c, file)
}

func save(c *MetaData, file string) *MetaData {
	data, err := yaml.Marshal(c)
	Check(err)
	err = ioutil.WriteFile(file, data, 0755)
	Check(err)
	return c
}

func (c *MetaData) IsRoot() bool {
	return c.PageTitle == "" && c.MetaFile == ""
}
