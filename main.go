package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"strings"

	"MagesTools/script"
	"MagesTools/script/format"
	"MagesTools/script/utils"
	"github.com/go-restruct/restruct"
)

func main() {

	fmt.Print(`MagesTools 
Version: 0.2.3_2024.06.05
Author: WéΤοr (wetorx@qq.com)
Github: https://github.com/wetor/MagesTools
License: GPL-3.0

`)
	var pType, pSource, pInput, pOutput, pScriptFormat, pCharset, pTbl string
	var pImport, pExport, pSkipChar bool
	var pDebug int
	flag.StringVar(&pType, "type", "", `[required] Source file type.
    Mages Script: "script"
        Supported MES(msb), SC3(scx)
    Diff Binary File: "diff"
        Diff input and output file
`)
	flag.BoolVar(&pExport, "export", false, "[optional] Export mode. Support folder export")
	flag.BoolVar(&pImport, "import", false, "[optional] Import mode")
	flag.IntVar(&pDebug, "debug", 0, `[optional] Debug level
    0: Disable debug mode
    1: Show info message
    2: Show warning message (For example, the character table is missing characters)
    3: Not implemented`)

	flag.StringVar(&pSource, "source", "", `[required] Source files or folder`)

	flag.StringVar(&pInput, "input", "", `[optional] Usually the import mode requires`)
	flag.StringVar(&pOutput, "output", "", `[required] Output file or folder`)

	flag.StringVar(&pScriptFormat, "format", "Npcs", `[script.required] Format of script export and import. Case insensitive
    NPCSManager format: "Npcs"
    NPCSManager Plus format: "NpcsP"`)
	flag.StringVar(&pCharset, "charset", "", `[script.optional] Character set containing only text. Must be utf8 encoding. Choose between "charset" and "tbl"`)
	flag.StringVar(&pTbl, "tbl", "", `[script.optional] Text in TBL format. Must be utf8 encoding. Choose between "charset" and "tbl"`)

	flag.BoolVar(&pSkipChar, "skip", true, "[script.optional] Skip repeated characters in the character table.")

	flag.Parse()
	restruct.EnableExprBeta()

	if pDebug >= 2 {
		utils.ShowWarning = true
	}
	switch pType {
	case "diff":
		if len(pInput) == 0 && len(pOutput) == 0 {
			panic("Diff requires input and output")
		}
		res := utils.FileSameCheck(pInput, pOutput)
		if res {
			fmt.Println("The two files are identical")
		}
	case "script":
		if !pExport && !pImport {
			panic("You must specify export mode or import mode")
		}
		if len(pSource) == 0 {
			panic("The source file or folder must be specified")
		}

		var _format format.Format
		switch strings.ToUpper(pScriptFormat) {
		case "NPCS":
			_format = &format.Npcs{}
		case "NPCSP":
			_format = &format.NpcsP{}
		default:
			panic("Unknown script export format")
		}
		scr := &script.Script{}

		if len(pCharset) > 0 {
			scr.LoadCharset(pCharset, false, pSkipChar)
		} else if len(pTbl) > 0 {
			scr.LoadCharset(pTbl, true, pSkipChar)
		} else {
			panic("A charset file or a tbl file must be specified")
		}

		if pExport {
			if utils.IsDir(pSource) && utils.IsDir(pOutput) {
				files, _ := utils.GetDirFileList(pSource)
				for _, file := range files {
					if pDebug >= 1 {
						fmt.Println(file)
					}
					scr.Open(file, _format)
					scr.Read()
					// Export
					scr.SaveStrings(filepath.Join(pOutput, filepath.Base(file)+".txt"))
				}
			} else if utils.IsFile(pSource) && utils.IsFile(pOutput) {
				scr.Open(pSource, _format)
				scr.Read()
				scr.SaveStrings(pOutput)
			} else {
				panic("Source and output must be files or folders.")
			}
		} else if pImport {
			scr.Open(pSource, _format)
			scr.Read()
			if len(pInput) > 0 {
				scr.LoadStrings(pInput)
			} else {
				panic("An input file must be specified")
			}

			if len(pOutput) > 0 {
				scr.Write(pOutput)
			} else {
				panic("The output file must be specified")
			}
		}
	}

}
