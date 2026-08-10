# Mages Engine Toolkit
English localization of wetor's MagesTools tool.

## Game Compatibility
- Theoretically supports all games powered by the Mages engine
- All MES (msb) and SC3 (scx / scr) scripts can be exported and imported without issues

## Usage guide
* Refer to [How to use](https://github.com/ThePlayer14/MagesTools_en/blob/master/GUIDE.md)


## Usage
```
  -charset string
        [script.optional] Character set containing only text. Must be utf8 encoding. Choose between "charset" and "tbl"
  -compound string
        [script.optional] Compound character table file (e.g. CompoundCharacters.tbl). Format: "[HexCode]=value" or "[HexStart-HexEnd]=value". Enables compound character support. When loaded, compound codes are wrapped as [value] on export and written back to a single code on import.
  -debug int
        [optional] Debug level
            0: Disable debug mode
            1: Show info message
            2: Show warning message (For example, the character table is missing characters)
            3: Not implemented
  -export
        [optional] Export mode. Support folder export
  -format string
        [script.required] Format of script export and import. Case insensitive
            NPCSManager format: "Npcs"
            NPCSManager Plus format: "NpcsP" (default "Npcs")
  -import
        [optional] Import mode
  -input string
        [optional] Usually the import mode requires
  -output string
        [required] Output file or folder
  -skip
        [script.optional] Skip repeated characters in the character table. (default true)
  -source string
        [required] Source files or folder
  -tbl string
        [script.optional] Text in TBL format. Must be utf8 encoding. Choose between "charset" and "tbl"
  -type string
        [required] Source file type.
            Mages Script: "script"
                Supported MES(msb), SC3(scx)
            Diff Binary File: "diff"
                Diff input and output file


```
### Examples

```shell
# Export all files in the folder using the tbl code table in NpcsP format. Do not skip identical characters in the code table. Enable debug mode with value 2.
MagesTools -type=script -export -skip=false -debug=2\
  -format=NpcsP \
  -tbl=./data/CC/MJPN.txt \
  -source=./data/CC/script/mes00 \
  -output=./data/CC/txt


# Export text using the tbl code table in NpcsP format, skipping identical characters in the code table.
MagesTools -type=script -export -skip=true \
  -format=NpcsP \
  -tbl=./data/CC/MJPN.txt \
  -source=./data/temp/1.msb \
  -output=./data/temp/1.msb.txt 

  
# Import text using the tbl code table in NpcsP format, skipping identical characters in the code table.
MagesTools -type=script -import -skip=false \
  -format=NpcsP \
  -tbl=./data/CC/MJPN.txt \
  -source=./data/temp/1.msb \
  -input=./data/temp/1.msb.txt \
  -output=./data/temp/1.msb.txt.msb

# RNE uses the following parameters:
# Export text using the charset character set in Npcs format, without skipping identical characters in the character set.
MagesTools -type=script -export -skip=false \
  -format=Npcs \
  -charset=./data/RNE/Charset_PSV_JP.utf8 \
  -source=./data/temp/1.msb \
  -output=./data/temp/1.msb.txt 

  
# Import text using the charset character set, formatted as Npcs, without skipping identical characters in the character set.
MagesTools -type=script -import -skip=false \
  -format=Npcs \
  -charset=./data/RNE/Charset_PSV_JP.utf8 \
  -source=./data/temp/1.msb \
  -input=./data/temp/1.msb.txt \
  -output=./data/temp/1.msb.txt.msb

# 11eyes CrossOver Xbox 360 parameters
# Exporting script (extract to text), no skipping, debug lv 2, format as Npcs
MagesTools -type=script -export -skip=false -debug=2 -charset ./charset/eleveneyes.utf8 -format=Npcs -source ../magesgame/script -output ../magesgame/scrout
# Importing script (replacing in the file), no skipping, debug lv 2, format as Npcs
MagesTools -type=script -import -skip=false -debug=2 -charset=./charset/eleveneyes.utf8 -format=Npcs -source=../magesgame/script/SC000.scr -input=./script-tl/SC000-tl.txt -output=./script-tl/output/SC000-tex.scr

# File comparison
MagesTools -type=diff \
  -input=./data/temp/1.msb \
  -output=./data/temp/1.msb.txt
```

## Compound Characters

Some games use **compound (multi-character) glyphs**: a single character code that maps to several
characters (for example a combined or decorated glyph). The normal charset/TBL only maps one code to
one display character, so a compound cannot be re-imported correctly — it would be split into
individual characters.

Pass a compound table with `-compound` to handle them. The file format matches the C# SciAdv.Net
tool's `CompoundCharacters.tbl`:

- One entry per line: `[HexCode]=value`
- A range is also supported: `[HexStart-HexEnd]=value`
- `HexCode` is the 16-bit character code (e.g. `0xE000`); `value` is the multi-character string it stands for

Example `CompoundCharacters.tbl`:

```
[E000]=あい
[E001-E003]=うえお
```

Behavior:

- **Export:** a code present in the compound table is written wrapped in brackets, e.g. `[あい]`,
  so it is distinguishable from a literal sequence of those characters.
- **Import:** a bracket group that is not a raw byte sequence (`[0x..]`) and not a name marker
  (`:[..]:`) is looked up in the compound table and written back as the single code. Unknown
  bracket groups fall back to being treated as literal text.

Example:

```shell
# Export with compound character support
MagesTools -type=script -export -format=NpcsP \
  -tbl=./data/CC/MJPN.txt \
  -compound=./data/CC/CompoundCharacters.tbl \
  -source=./data/temp/1.msb \
  -output=./data/temp/1.msb.txt

# Import with compound character support
MagesTools -type=script -import -format=NpcsP \
  -tbl=./data/CC/MJPN.txt \
  -compound=./data/CC/CompoundCharacters.tbl \
  -source=./data/temp/1.msb \
  -input=./data/temp/1.msb.txt \
  -output=./data/temp/1.msb.out.msb
```

When `-compound` is not supplied, behavior is unchanged.

## Script
### Format
The current format (NpcsP) is an optimized version of the NPCSManager format.
- Removed the `[1x01][1x02]` name markers; names are now wrapped with `:[` and `]:` (e.g. `:[name]:`).
- Removed the half-width space that previously followed `]:`.
- All reserved byte data is written with a `0x` prefix, e.g. `[0x04A01414]`.
- Removed the `color` markup `<#` ... `#>`; color is now represented by a plain byte marker, e.g. `[0x04A01414][0x00]`.
- Added basic byte-level parsing of `EvaluateExpression` commands, e.g. `[0x15290AA4B51414008100][0x00]`. Some bugs may still remain.
Script sample:
```
[0x0F][0x1100CC][0x04A01414][0x00]『白い光が見えた』[0x15290AA4B51414008100][0x00][0x03][0xFF]
[0x0F][0x110026][0x04A01414][0x00]『耳鳴りのような音が聞こえた』[0x15290AA4B51414008100][0x00][0x08][0xFF]
[0x0F][0x1100F2]勘違いだと笑ってしまうにはあまりに多くの者たちが[0x1F]体験してしまったこの現象は、原因不明のまま語り継がれ、[0x1F]地震のおかしさを疑う者の手助けをすることとなった。[0x15290AA4B51414008100][0x00][0x08][0xFF]
[0x0F][0x110118]そして、噂にまみれた地震から６年経った２０１５年。[0x15290AA4B51414008100][0x00][0x08][0xFF]
[0x0F][0x1100F2]新しく生まれ変わりつつある渋谷の街で、[0x1F]地震とは別の事件が世間の注目を集めようとしていた。[0x15290AA4B51414008100][0x00][0x08][0xFF]
[0x0F][0x110118]２０１５年９月７日（日）夜[0x15290AA4B51414008100][0x00][0x08][0xFF]
:[男性]:「はい、ではいつも通り３分くらい募集をかけるんで、適当によろです」[0x03][0xFF]
そう言った途端に、コメント欄が一気に流れ出した。[0x03][0xFF]
流れ具合を数秒間確認していると、『ハルちゃんの熱愛報道はいつ？』との依頼を見つけ、[0x09]大谷[0x0A]おおたに[0x0B][0x09]悠馬[0x0A]ゆうま[0x0B]は思わず微笑んだ。[0x03][0xFF]
狙い通りだ。[0x03][0xFF]
依頼の大半はイケメン俳優か女性アイドルに関することだから、視聴者の傾向を読むことは馬鹿みたいに簡単だった。[0x03][0xFF]
問題は、どの人物の名前が挙がるかということで、こればかりは運とその人物の人気による。[0x03][0xFF]
が、[0x04280AA0][0x2D14][0x00]ハルちゃん[0x04800000][0x8113][0x8113]確かなんとかハルコとかいったか[0x8113][0x8113]ならば大丈夫だ。[0x03][0xFF]
先日、行きたくもないイベントに行って、[0x09][0x1E]直接見て来た[0x0A][0x8117][0x8117][0x8117][0x8117][0x8117][0x8117][0x0B]ばかりだ。[0x03][0xFF]
:[大谷]:「……よし」[0x03][0xFF]
```

## Planned features
- Support more formats

## Version history / changelog

### 2026.6.12
- Fixed `SetColor` (color tag `0x04`) reading/swallowing too many bytes: it now reads a single color byte instead of three, so the character byte following a color command is no longer absorbed into the color tag (previously truncated/garbled text such as missing leading characters and bracket glyphs).

### 2024.6.5
- Fixed expression termination detection
- Resolved script parsing issues with empty text
- Fixed folder export write permission issues on Windows
- (Issues identified by discoverer [Fluchw](https://github.com/wetor/MagesTools/issues/5))

### 2022.10.21
- Fixed an encoding error caused by byte data (`:[0xFF]`) following the ‘:’ character. Identified by discoverer [kurikomoe](https://github.com/kurikomoe).

### 2022.3.21
- Supports importing and exporting text files for SC3 (scx / scr) scripts
- Supports folder export (import not currently supported)
- Added basic logging functionality
- Optimized code details

### 2022.3.20 2
- Restructured code architecture to support additional export formats
- Added support for exporting and importing NPCSManager formats
  - Supports importing NPCSManager export files
  - Original NPCSManager cannot import files exported by this program due to minor differences
- Added command-line invocation support
- Enhanced help documentation

### 2022.3.20
- Complete MES (MSB) text import (simple implementation)
- Adjust export format

### 2022.3.19
- Basic framework design
- Completed text export for MES (msb)
### 2022.3.18
- Initial version


## Credits 
This project is made possible due to the efforts of
- [marcussacana](https://github.com/marcussacana)'s [NPCSManager](https://github.com/marcussacana/NPCSManager)
- [liaowm5](https://github.com/SteiensGate)'s msb_tool.py
- [CommitteeOfZero](https://github.com/CommitteeOfZero)'s [sc3ntist](https://github.com/CommitteeOfZero/sc3ntist) and [SciAdv.Net](https://github.com/CommitteeOfZero/SciAdv.Net) projects.
