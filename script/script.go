package script

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"MagesTools/script/format"
	"MagesTools/script/utils"
)

type Entry struct {
	Index  int `struct:"int32"`
	Offset int `struct:"int32"`
	Length int `struct:"-"`
}

type Strings interface {
	// ReadStrings
	//  Description 读取解析脚本全部字符串
	//  Param readString
	ReadStrings(readString func([]byte) string)
	// GetStrings
	//  Description 取出全部字符串
	//  Return []string
	GetStrings() []string
	// SetStrings
	//  Description 替换全部字符串
	//  Param strings
	SetStrings(strings []string)
	// WriteStrings
	//  Description 写到导入字符串
	//  Param writeString
	WriteStrings(writeString func(string) []byte)
	// GetRaw
	//  Description 获取脚本数据
	//  Return []byte
	GetRaw() []byte
}

type Script struct {
	Name           string
	Strings        Strings
	Format         format.Format
	DecodeCharset  map[uint16]string
	EncodeCharset  map[string]uint16
	DecodeCompound map[uint16]string
	EncodeCompound map[string]uint16
}

// NewScript
//
//	Description 打开脚本文件
//	Param filename string
//	Return *Script
func NewScript(filename string, format format.Format) *Script {
	script := &Script{}
	script.Open(filename, format)
	return script
}

// Open
//
//	Description 打开脚本文件，如果已经使用LoadCharset载入码表，则不需要重新调用LoadCharset
//	Receiver s *Script
//	Param filename string
//	Param format format.Format
func (s *Script) Open(filename string, format format.Format) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s.Name = filepath.Base(filename)
	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	s.Format = format
	switch string(data[0:3]) {
	case "MES":
		s.Strings = LoadMes(data)
	case "SC3":
		s.Strings = LoadSc3(data)
	default:
		panic("Unsupported file type!" + filename)
	}
}

// LoadCharset
//
//	Description 载入码表/字符集
//	Receiver s *Script
//	Param filename string 文件名
//	Param isTBL bool 是否为码表。否则为字符集，字符集从0x8000开始
//	Param skipExist bool 是否检查并跳过重复出现的字符，仅以第一次出现为准
func (s *Script) LoadCharset(filename string, isTBL, skipExist bool) {

	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	decodeCharset := make(map[uint16]string, 65535)
	encodeCharset := make(map[string]uint16, 65535)
	if isTBL {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) > 5 {
				k := utils.BytesToUint16Big(utils.HexToBytes(line[0:4]))
				v := line[5:]
				if skipExist {
					if _, has := encodeCharset[v]; !has {
						decodeCharset[k] = v
						encodeCharset[v] = k
					}
				} else {
					decodeCharset[k] = v
					encodeCharset[v] = k
				}

			}
		}
	} else {
		data, _ := io.ReadAll(f)
		runes := []rune(string(data))
		for i, char := range runes {
			k := uint16(0x8000 + i)
			v := string(char)
			if skipExist {
				if _, has := encodeCharset[v]; !has {
					decodeCharset[k] = v
					encodeCharset[v] = k
				}
			} else {
				decodeCharset[k] = v
				encodeCharset[v] = k
			}
		}
	}
	s.DecodeCharset = decodeCharset
	s.EncodeCharset = encodeCharset
}

// LoadCompound
//
//	Description 载入复合字符表（compound characters）。
//	            文件格式与 C# 工具 CompoundCharacters.tbl 一致：每行 `[HexCode]=value`，
//	            或区间 `[HexStart-HexEnd]=value`。例如 `[E000]=あい`。
//	            载入后，导出会把复合字符包裹为 [value]，导入时 [value] 会编码为单个码值。
//	Receiver s *Script
//	Param filename string 文件名
func (s *Script) LoadCompound(filename string) {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	decodeCompound := make(map[uint16]string)
	encodeCompound := make(map[string]uint16)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || !strings.Contains(line, "=") {
			continue
		}
		eq := strings.Index(line, "=")
		left := line[:eq]
		value := line[eq+1:]
		left = strings.TrimPrefix(left, "[")
		left = strings.TrimSuffix(left, "]")
		if strings.Contains(left, "-") {
			parts := strings.SplitN(left, "-", 2)
			start, err1 := strconv.ParseUint(parts[0], 16, 16)
			end, err2 := strconv.ParseUint(parts[1], 16, 16)
			if err1 != nil || err2 != nil {
				continue
			}
			for i := start; i <= end; i++ {
				code := uint16(i)
				decodeCompound[code] = value
				encodeCompound[value] = code
			}
		} else {
			code, err := strconv.ParseUint(left, 16, 16)
			if err != nil {
				continue
			}
			c := uint16(code)
			decodeCompound[c] = value
			encodeCompound[value] = c
		}
	}
	s.DecodeCompound = decodeCompound
	s.EncodeCompound = encodeCompound
	fmt.Printf("Loaded compound character table: %d entries from %s\n", len(decodeCompound), filename)
}

// Read
//
//	Description 解析文本，需要至少执行一次script.LoadCharset载入码表
//	Receiver s *Script
func (s *Script) Read() {
	if s.DecodeCharset != nil && s.EncodeCharset != nil {
		s.Format.SetCharset(s.DecodeCharset, s.EncodeCharset)
	}
	if s.DecodeCompound != nil && s.EncodeCompound != nil {
		s.Format.SetCompound(s.DecodeCompound, s.EncodeCompound)
	}
	s.Strings.ReadStrings(s.Format.DecodeLine)
}

// SaveStrings
//
//	Description 保存文本，需要先执行script.Read
//	Receiver s *Script
//	Param filename string
func (s *Script) SaveStrings(filename string) {
	strings := s.Strings.GetStrings()
	if len(strings) == 0 {
		fmt.Printf("File %s does not contain any text, skipping\n", s.Name)
		return
	}
	f, _ := os.Create(filename)
	defer f.Close()
	for _, str := range strings {
		f.WriteString(str + "\n")
	}
}

// LoadStrings
//
//	Description 载入文本并导入
//	Receiver s *Script
//	Param filename string
func (s *Script) LoadStrings(filename string) {
	f, _ := os.Open(filename)
	defer f.Close()

	strings := make([]string, 0, len(s.Strings.GetStrings()))
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		strings = append(strings, line)
	}
	s.Strings.SetStrings(strings)
}

// Write
//
//	Description 保存为脚本
//	Receiver s *Script
//	Param filename string
func (s *Script) Write(filename string) {
	s.Strings.WriteStrings(s.Format.EncodeLine)

	f, _ := os.Create(filename)
	defer f.Close()
	f.Write(s.Strings.GetRaw())
}
