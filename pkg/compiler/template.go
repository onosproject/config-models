// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package compiler

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func (c *ModelCompiler) applyTemplate(name, tplPath, outPath string) error {
	var funcs template.FuncMap = map[string]interface{}{
		"quote": func(value interface{}) string {
			return fmt.Sprintf("\"%s\"", value)
		},
		"replace": func(search, replace string, value interface{}) string {
			return strings.ReplaceAll(fmt.Sprint(value), search, replace)
		},
		"capitalize": func(s string) string {
			caser := cases.Title(language.English)
			return caser.String(s)
		},
	}

	tpl, err := template.New(name).
		Funcs(funcs).
		ParseFiles(tplPath)
	if err != nil {
		return err
	}

	file, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return tpl.Execute(file, c.dictionary)
}

func (c *ModelCompiler) getTemplatePath(name string) string {
	return filepath.Join("templates", name)
}

func (c *ModelCompiler) createDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Debugf("Creating '%s'", dir)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			log.Errorf("Creating '%s' failed: %s", dir, err)
		}
	}
}
