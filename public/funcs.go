package public

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/leekchan/accounting"
)

const (
	ONE_KB = float64(1024)          // 1024B
	ONE_MB = ONE_KB * float64(1024) // 1024KB
	ONE_GB = ONE_MB * float64(1024) // 1024MB
	ONE_TB = ONE_GB * float64(1024) // 1024GB
	ONE_PB = ONE_TB * float64(1024) // 1024TB
	ONE_EB = ONE_PB * float64(1024) // 1024PB
	ONE_ZB = ONE_EB * float64(1024) // 1024EB
	ONE_YB = ONE_ZB * float64(1024) // 1024ZB
	ONE_BB = ONE_YB * float64(1024) // 1024YB
)

// support: B KB MB GB TB PB EB ZB YB BB
func FormatStorageSize(size float64) string {
	if size >= ONE_BB {
		i := Round(size / ONE_BB)
		return fmt.Sprintf("%.1fBB", i)
	}
	if size >= ONE_YB {
		i := Round(size / ONE_YB)
		return fmt.Sprintf("%.1fYB", i)
	}
	if size >= ONE_ZB {
		i := Round(size / ONE_ZB)
		return fmt.Sprintf("%.1fZB", i)
	}
	if size >= ONE_EB {
		i := Round(size / ONE_EB)
		return fmt.Sprintf("%.1fEB", i)
	}
	if size >= ONE_PB {
		i := Round(size / ONE_PB)
		return fmt.Sprintf("%.1fPB", i)
	}
	if size >= ONE_TB {
		i := Round(size / ONE_TB)
		return fmt.Sprintf("%.1fTB", i)
	}
	if size >= ONE_GB {
		i := Round(size / ONE_GB)
		return fmt.Sprintf("%.1fGB", i)
	}
	if size >= ONE_MB {
		i := Round(size / ONE_MB)
		return fmt.Sprintf("%.1fMB", i)
	}
	if size >= ONE_KB {
		i := Round(size / ONE_KB)
		return fmt.Sprintf("%.1fKB", i)
	}
	return fmt.Sprintf("%dB", int64(size))
}

// Round 小数点后n位四舍五入
func Round(a float64, n ...int) float64 {
	_n := 1
	if len(n) > 0 {
		_n = n[0]
		if _n < 1 {
			_n = 1
		}
	}
	pow10 := math.Pow10(_n)
	floor := math.Floor(a*pow10+0.5) / pow10
	return floor
}

func ToJsonBytes(src interface{}, escape ...bool) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	if len(escape) > 0 {
		jsonEncoder.SetEscapeHTML(escape[0])
	} else {
		jsonEncoder.SetEscapeHTML(false)
	}
	err := jsonEncoder.Encode(src)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func ToJsonString(src interface{}, escape ...bool) (string, error) {
	toJsonBytes, err := ToJsonBytes(src, escape...)
	if err != nil {
		return "", err
	}
	return string(toJsonBytes), nil
}

func ToJsonStringWithoutError(src interface{}, escape ...bool) string {
	jsonString, _ := ToJsonString(src, escape...)
	return jsonString
}

// ConvertByJson is 通过json转换对象，参数target为指针
func ConvertByJson(src, target interface{}, escape ...bool) (err error) {
	var marshal []byte
	_type := fmt.Sprintf("%T", src)
	switch _type {
	case "[]byte", "[]uint8":
		marshal = src.([]byte)
	case "string":
		marshal = []byte(src.(string))
	default:
		marshal, err = ToJsonBytes(src, escape...)
		if err != nil {
			return err
		}
	}

	err = json.Unmarshal(marshal, target)
	if err != nil {
		return err
	}
	return nil
}

// FormatTimeToCustom1 格式化时间：今天/昨天/2天前/3个月前/2017-01-01
func FormatTimeToCustom1(createdAt *gtime.Time) string {
	return formatTimeToCustom(createdAt)
}

func formatTimeToCustom(createdAt *gtime.Time) string {
	current := gtime.Now()
	hours := current.Sub(createdAt).Hours()
	if hours < 24 {
		return "今天" + createdAt.Format("H:i")
	}
	if hours < 48 {
		return "昨天"
	}
	if hours < 72 {
		return "2天前"
	}
	if hours < 96 {
		return "3天前"
	}
	return createdAt.Format("Y-m-d")
}

// 格式化金融数字
// 例如：fmt.Println(FormatAccounting(1000.01, "￥", "/年"))
// ￥1,000.01/年
func FormatAccounting(x float64, prefix, suffix string) string {
	x, Precision := RoundAndCount(x)
	ac := accounting.Accounting{Symbol: prefix, Precision: Precision}
	ret := ac.FormatMoney(x)
	return ret + suffix
}

func RoundAndCount(input float64) (round float64, nonZeroDigits int) {
	// 先四舍五入
	round = Round(input, 2)

	// 分离小数部分
	a := strings.Split(gconv.String(round), ".")
	if len(a) > 1 {
		decimalPart := a[1]
		// 计算小数部分非零位数
		nonZeroDigits = len(decimalPart)
	} else {
		nonZeroDigits = 0
	}

	return round, nonZeroDigits
}

// 模拟三目运算
func TrinocularOperation[T any](a bool, b, c T) T {
	return TrinocularOperationT(a, b, c)
}

// 模拟三目运算
func TrinocularOperationStr(a bool, b, c string) string {
	// v := gconv.String(TrinocularOperation(a, b, c))
	v := TrinocularOperationT(a, b, c)
	return v
}

// 模拟三目运算
func TrinocularOperationInt64(a bool, b, c int64) int64 {
	// v := gconv.Int64(TrinocularOperation(a, b, c))
	v := TrinocularOperationT(a, b, c)
	return v
}

// 模拟三目运算
func TrinocularOperationInt(a bool, b, c int) int {
	// v := gconv.Int(TrinocularOperation(a, b, c))
	v := TrinocularOperationT(a, b, c)
	return v
}

// 模拟三目运算
func TrinocularOperationMap(a bool, b, c map[string]interface{}) map[string]interface{} {
	// v := gconv.Map(TrinocularOperation(a, b, c))
	v := TrinocularOperationT(a, b, c)
	return v
}

// 模拟三目运算
func TrinocularOperationT[T any](a bool, b, c T) T {
	if a {
		return b
	}
	return c
}

// 模拟三目运算
func TrinocularOperationTWithFunc[T any](a bool, b, c func() T) T {
	if a {
		return b()
	}
	return c()
}

func ParseJson[T any](jsonStr string, pointer T) (T, error) {
	err := json.Unmarshal([]byte(jsonStr), pointer)
	return pointer, err
}

func ParseJsonWithoutError[T any](jsonStr string, pointer T) T {
	ParseJson(jsonStr, pointer)
	return pointer
}
