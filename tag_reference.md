## Specific tag format differences (sc3tools vs MagesTools)

| Element | sc3tools (C#) Format | MagesTools Format |
|---------|--------------|------------|
| Line break (`0x00`) | `[linebreak]` | `[0x00]` |
| Character name start (`0x01`) | `[name]` | `:[` |
| Dialogue line start (`0x02`) | `[line]` | `]:` |
| Present (`0x03`) | `[%p]` | `[0x03]` |
| Present - reset alignment (`0x08`) | `[%e]` | `[0x08]` |
| Present - unknown (`0x18`) | `[%18]` | `[0x18]` |
| Color (`0x04` + index) | `[color index="02"]` | `[0x0402]` |
| Ruby base (`0x09`) | `[ruby-base]` | `[0x09]` |
| Ruby text start (`0x0A`) | `[ruby-text-start]` | `[0x0A]` |
| Ruby text end (`0x0B`) | `[ruby-text-end]` | `[0x0B]` |
| Font size (`0x0C` + 2 bytes) | `[font size="1000"]` | `[0x0C03E8]` |
| Print in parallel (`0x0E`) | `[parallel]` | `[0x0E]` |
| Center text (`0x0F`) | `[center]` | `[0x0F]` |
| Set margin top (`0x11` + 2 bytes) | `[margin top="228"]` | `[0x1100E4]` |
| Set margin left (`0x12` + 2 bytes) | `[margin left="142"]` | `[0x12008E]` |
| Get hardcoded value (`0x13` + 2 bytes) | `[hardcoded-value index="0"]` | `[0x130000]` |
| Evaluate expression (`0x15` + bytes) | `[evaluate expr="..."]` | `[0x15...]` |
| Auto-forward | `[auto-forward]` | `[0x19]` |
| Compound character | `[①]` (resolved name) | raw private-use code point (e.g. ``) |
| String terminator | (none) | `[0xFF]` appended to every line |
| Full-width chars | normalized to half-width | preserved |
