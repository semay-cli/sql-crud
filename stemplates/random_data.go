package stemplates

import (
	"math/rand"
	"regexp"
	"strings"
	"text/template"
	"unicode"

	"fmt"
	"strconv"
	"time"
)

// Utility function to generate random strings
func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func formatSliceToString(slice []string) string {
	// Wrap each element with double quotes
	quotedSlice := make([]string, len(slice))
	for i, v := range slice {
		quotedSlice[i] = `"` + v + `"`
	}

	// Join the quoted elements with ", " separator
	return strings.Join(quotedSlice, ", ")
}

// Generates a random email address
func randomEmail() string {
	return randomString(10) + "@example.com"
}

func replaceString(old string) string {
	newStr := strings.ReplaceAll(old, "-", "_")
	return newStr
}

func replaceStringCapitalize(old string) string {
	// Replace all hyphens with underscores
	newStr := strings.ReplaceAll(old, "-", "_")
	// Capitalize the entire string
	newStr = strings.ToUpper(newStr)
	return newStr
}

// Generates a random UUID
func randomUUID() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", rand.Uint32(), rand.Uint32(), rand.Uint32(), rand.Uint32(), rand.Uint32())
}

// Generates a random ID (could be a large number, or a UUID)
func randomID() string {
	return randomUUID() // You can replace this with a custom ID generation method
}

// Generates a random uint (unsigned integer)
func randomUInt() uint {
	return uint(rand.Intn(1000000))
}

// Generates a random float64
func randomFloat64() float64 {
	return rand.Float64() * 1000.0
}

// Generates a random float32
func randomFloat32() float32 {
	return rand.Float32() * 1000.0
}

// Generates a random int32
func randomInt32() int32 {
	return int32(rand.Intn(1000000))
}

// Generates a random int64
func randomInt64() int64 {
	return int64(rand.Intn(1000000))
}

// Generates a random string
func randomGenericString() string {
	return randomString(8) // Adjust length if needed
}

// Generates a random time
func randomTime() time.Time {
	return time.Now().Add(time.Duration(rand.Intn(1000000)) * time.Second)
}

// Parsing functions
func parseTime(dateStr string) time.Time {
	// You can use the time.RFC3339 format for your specific date string
	result, _ := time.Parse(time.RFC3339, dateStr)
	return result
}

func parseInt(intString string) int {
	// You can use the time.RFC3339 format for your specific date string
	result, _ := strconv.Atoi(intString)
	return result
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

// CamelToSnake converts CamelCase or PascalCase to snake_case
func CamelToSnake(s string) string {
	// Insert underscore before all caps that are followed by lowercase letters
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := re.ReplaceAllString(s, "${1}_${2}")

	// Handle acronym-style capitals (e.g., "JSONData" -> "json_data")
	re = regexp.MustCompile("([A-Z]+)([A-Z][a-z])")
	snake = re.ReplaceAllString(snake, "${1}_${2}")

	return strings.ToLower(snake)
}

// CamelToSnake converts CamelCase or PascalCase to snake_case
func ToLowerCaseName(s string) string {
	return strings.ToLower(s)
}

// Capitalize first letter of a word
func capitalize(word string) string {
	if len(word) == 0 {
		return ""
	}
	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// Convert "blue-admin" to "BlueAdmin"
func ToPascalCase(input string) string {
	parts := strings.Split(input, "-")
	for i, part := range parts {
		parts[i] = capitalize(part)
	}
	return strings.Join(parts, "")
}

// PascalToSnake converts a PascalCase string to snake_case
func PascalToSnake(input string) string {
	// Insert an underscore before all capital letters (except the first)
	re := regexp.MustCompile("([A-Z][a-z0-9]*)")
	matches := re.FindAllStringSubmatch(input, -1)

	var words []string
	for _, match := range matches {
		words = append(words, strings.ToLower(match[0]))
	}

	return strings.Join(words, "_")
}

// PascalToWords converts PascalCase or camelCase to "Word Word" format
func PascalToWords(input string) string {
	// Add a space before all capital letters (except the first one)
	re := regexp.MustCompile(`([a-z])([A-Z])`)
	spaced := re.ReplaceAllString(input, `$1 $2`)

	// Optional: capitalize first letter (if input is camelCase)
	return strings.TrimSpace(spaced)
}

// SnakeToWords converts "something_new" → "Something New"
func SnakeToWords(input string) string {
	words := strings.Split(input, "_")
	for i, word := range words {
		if len(word) > 0 {
			// Capitalize first letter
			words[i] = strings.ToUpper(word[:1]) + word[1:]
		}
	}
	return strings.Join(words, " ")
}

// Function to extract the first letter of each word
func getFirstLetters(input string) string {
	// Split the string by hyphen and iterate over the parts
	words := strings.FieldsFunc(input, func(r rune) bool {
		return r == '-' || unicode.IsSpace(r)
	})

	// Collect the first letter of each word
	var result strings.Builder
	for _, word := range words {
		if len(word) > 0 {
			result.WriteRune(unicode.ToLower(rune(word[0]))) // Convert to lowercase if needed
		}
	}

	return result.String()
}

var FuncMap = template.FuncMap{
	"camelToSnake":            CamelToSnake,            // Register CamelToSnake function
	"add":                     add,                     // Register Add function
	"parseTime":               parseTime,               // Register custom function
	"parseInt":                parseInt,                // Register custom function
	"randomEmail":             randomEmail,             // Register random email function
	"randomUUID":              randomUUID,              // Register random UUID function
	"randomID":                randomID,                // Register random ID function
	"randomUInt":              randomUInt,              // Register random uint function
	"randomFloat64":           randomFloat64,           // Register random float64 function
	"randomFloat32":           randomFloat32,           // Register random float32 function
	"randomInt32":             randomInt32,             // Register random int32 function
	"randomInt64":             randomInt64,             // Register random int64 function
	"randomString":            randomGenericString,     // Register random string function
	"randomTime":              randomTime,              // Register random time function
	"randomBool":              randomBool,              // Register random bool function
	"replaceString":           replaceString,           // Register hyphen with underscore
	"replaceStringCapitalize": replaceStringCapitalize, // Register hyphen with underscore
	"formatSliceToString":     formatSliceToString,     // Register format slice to string functionD
	"toLowerCaseName":         ToLowerCaseName,         // Register format slice to string functionD
	"toPascalCase":            ToPascalCase,            // Register toPascalCase function
	"toSnakeCase":             PascalToSnake,           // Register PascalToSnake function
	"pascalToWords":           PascalToWords,           // Register PascalToWords function
	"snakeToWord":             SnakeToWords,            // Register SnakeToWords function
	"capitalize":              capitalize,              // Register capitalize function
	"getFirstLetters":         getFirstLetters,         // Register getFirstLetters function
}

func add(a int, b int) int {
	return a + b
}
