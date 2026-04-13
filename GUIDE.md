# MagesTools Script editing guide

## 11eyes CrossOver Xbox 360 parameters

# Exporting script (extract to text), no skipping, debug lv 2, format as NpcsP
`MagesTools -type=script -export -skip=false -debug=2 -charset ./charset/eleveneyes.utf8 -format=NpcsP -source ../magesgame/script -output ../magesgame/scrout`
# Importing script (replacing in the file), no skipping, debug lv 2, format as NpcsP
`MagesTools -type=script -import -skip=false -debug=2 -charset=./charset/eleveneyes.utf8 -format=NpcsP -source=../magesgame/script/SC000.scr -input=./script-tl/SC000-tl.txt -output=./script-tl/output/SC000-tex.scr`

* Note that the source `.scr` file is used as a "template" in which the program replaces lines using the input text file provided and saves a new `.scr` file you can use as the new script that the game will load.
* Be sure to keep an eye on the line count consistency when editing the scripts. Using an inconsistent line count script will likely get Xenia throw a nested exception error after the script ends.

## What are the tags
|line_tag    |tag_function      |is_conditional|keep_tag|
|------------|------------------|--------------|--------|
|[0x110022]  |Formatting tag 1  |No            |Yes     |
|[0x110102]  |Formatting tag 2  |No            |Yes     |
|[0x1200F0]  |Formatting tag 3  |No            |Yes     |
|[0x03][0xFF]|Line closing tag 1|Yes           |Yes     |
|[0x08][0xFF]|Line closing tag 2|Yes           |Yes     |
|[0x110088]  |Formatting tag 4  |No            |Yes     |
|[0x0C0301]  |Formatting tag 5  |No            |Yes     |
|[0x0C0542]  |Formatting tag 6  |No            |Yes     |
|[0x0C0783]  |Formatting tag 7  |No            |Yes     |
|[0x09]      |Ruby text base    |No            |No      |
|[0x0A]      |Ruby text         |No            |No      |
|[0x0B]      |Ruby text end     |No            |No      |
|[0x0C03C2]  |Formatting tag 8  |No            |Yes     |
|[0x110044]  |Formatting tag 9  |No            |Yes     |
|[0x12013C]  |Formatting tag 10 |No            |Yes     |
|[0x120174]  |Formatting tag 11 |No            |Yes     |
|[0x00]      |Empty tag         |No            |Yes     |

* While I cannot fully clarify what are the purposes of the tags labeled as "Formatting tag", the ones starting with `0x1100` are likely used for margin adjustment.
* When you encounter a tag starting with `0x04`, it is a color tag, which also needs to be edited with a specific way to not get stray characters in the text. 
For that, swap the last three digits in a color code to three zeros. For example: `[0x04288339]` -> `[0x04288000]`  (the `[0x04]` part is the color set tag, in NpcsP format)
* Name tags are always starting with `:[` followed by the character's name, then closed with `]:`. 
If the NpcsP formatting used, you won't need to keep the `[0x01]`and `[0x02]` tags, because those only appear when using the Npcs format.
* Be sure to keep an eye on the Line closing tags' presence in the script. As to which one needs to be used, refer to the original script file. 
If it's not present will not break the game but will be clumped together with the line after the unclosed line.
* Sometimes you may encounter tags like `[0x15280AA12C142100]`. These are likely the `EvaluateExpression` kind of tags, and these must always be followed by the `[0x00]` Empty tag, because otherwise the game will get stuck. 

# Line tags in JSON
```
[
  {
    "line_tag": "[0x110022]",
    "tag_role": "Formatting tag 1",
    "is_conditional": "No",
    "keep_tag": "Yes"
  },
  {
    "line_tag": "[0x110102]",
    "tag_role": "Formatting tag 2",
    "is_conditional": "No",
    "keep_tag": "Yes"
  },
  {
    "line_tag": "[0x1200F0]",
    "tag_role": "Formatting tag 3",
    "is_conditional": "No",
    "keep_tag": "Yes"
  },
  {
    "line_tag": "[0x03][0xFF]",
    "tag_role": "Line closing tag 1",
    "is_conditional": "Yes",
    "keep_tag": "Yes"
  },
  {
    "line_tag": "[0x08][0xFF]",
    "tag_role": "Line closing tag 2",
    "is_conditional": "Yes",
    "keep_tag": "Yes"
  },
  {
    "line_tag": "[0x110088]",
    "tag_role": "Formatting tag 4",
    "is_conditional": "No",
    "keep_tag": "Yes"
  },
  {
    "line_tag": "[0x0C0301]",
    "tag_role": "Formatting tag 5",
    "is_conditional": "No",
    "keep_tag": "Yes"
  },
  {
    "line_tag": "[0x0C0542]",
    "tag_role": "Formatting tag 6",
    "is_conditional": "No",
    "keep_tag": "Yes"
  },
  {
    "line_tag": "[0x0C0783]",
    "tag_role": "Formatting tag 7",
    "is_conditional": "No",
    "keep_tag": "Yes"
  },
  {
    "line_tag": "[0x09]",
    "tag_role": "Ruby text base",
    "is_conditional": "No",
    "keep_tag": "No"
  },
  {
    "line_tag": "[0x0A]",
    "tag_role": "Ruby text",
    "is_conditional": "No",
    "keep_tag": "No"
  },
  {
    "line_tag": "[0x0B]",
    "tag_role": "Ruby text end",
    "is_conditional": "No",
    "keep_tag": "No"
  },
  {
    "line_tag": "[0x0C03C2]",
    "tag_role": "Formatting tag 8",
    "is_conditional": "No",
    "keep_tag": "Yes"
  },
  {
    "line_tag": "[0x110044]",
    "tag_role": "Formatting tag 9",
    "is_conditional": "No",
    "keep_tag": "Yes"
  },
  {
    "line_tag": "[0x12013C]",
    "tag_role": "Formatting tag 10",
    "is_conditional": "No",
    "keep_tag": "Yes"
  },
  {
    "line_tag": "[0x120174]",
    "tag_role": "Formatting tag 11",
    "is_conditional": "No",
    "keep_tag": "Yes"
  },
  {
    "line_tag": "[0x00]",
    "tag_role": "Empty tag",
    "is_conditional": "No",
    "keep_tag": "Yes"
  }
]
```
