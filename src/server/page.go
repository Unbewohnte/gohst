/*
            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
                    Version 2, December 2004

 Copyright (C) 2004 Sam Hocevar <sam@hocevar.net>

 Everyone is permitted to copy and distribute verbatim or modified
 copies of this license document, and changing it is allowed as long
 as the name is changed.

            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

  0. You just DO WHAT THE FUCK YOU WANT TO.
*/

// Kasyanov N.A. (Unbewohnte), 2023

package server

import (
	"Unbewohnte/gohst/logger"
	"html/template"
	"path/filepath"
)

// Constructs a pageName template via inserting basePageName in pagesDir
func getPage(pagesDir string, basePageName string, pageName string) (*template.Template, error) {
	page, err := template.ParseFiles(
		filepath.Join(pagesDir, basePageName),
		filepath.Join(pagesDir, pageName),
	)
	if err != nil {
		logger.Error("Failed to parse page files (pagename is \"%s\"): %s", pageName, err)
		return nil, err
	}

	return page, nil
}
