# Mages Engine Toolkit
English localization of wetor's MagesTools program.

## Game Compatibility
- Theoretically supports all games powered by the Mages engine
- All MES (msb) and SC3 (scx / scr) scripts can be exported and imported without issues

## Usage
```
  -charset string
        [script.optional] Character set containing only text. Must be utf8 encoding. Choose between "charset" and "tbl"
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
  -output=./data/temp/1.msb.txt.msb
```

## Script
### Format
The current format is an optimized version of NPCSManager.
- Delete`name`after`[1x01][1x02]`, using only`:[`value`]:`label name
- Delete`]:`half-width space
- All reserved byte data are implemented using`0x`at the beginning, such as`[0x04A01414]`
- Delete`color`special markings`<#`value`#>`, using only byte markers, such as`[0x04A01414][0x00]`
- Improve support for`EvaluateExpression`simple byte parsing of expressions, such as`[0x15290AA4B51414008100][0x00]`. There may be unknown bugs.

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

### 2024.6.5
- - Fixed expression termination detection
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
