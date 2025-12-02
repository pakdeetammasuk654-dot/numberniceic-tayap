package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"numberniceic/repositories"
	"regexp"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// --- Helper function for Color Coding ---
func getColorCodeForPairType(pairType string) string {
	// Trim space just in case of CHAR padding issues
	trimmedType := strings.TrimSpace(pairType)
	switch trimmedType {
	// Good numbers (shades of green, dark to light)
	case "D10":
		return "#28a745" // Strong Green
	case "D8":
		return "#5cb85c" // Medium Green
	case "D5":
		return "#90ee90" // Light Green
	// Bad numbers (shades of red, dark to light)
	case "R10":
		return "#dc3545" // Strong Red
	case "R7":
		return "#f0ad4e" // Orange/Warning
	case "R5":
		return "#ffcccb" // Light Pink/Red
	default:
		return "#6c757d" // Neutral Gray for others
	}
}

// --- Input Validation ---
var validInputRegex = regexp.MustCompile(`^[a-zA-Z\p{Thai}]+( [a-zA-Z\p{Thai}]+)*$`)

func validateInputString(input string) bool {
	return validInputRegex.MatchString(input)
}

// --- Generic Handler ---
func NewGenericNumHandler(repo repositories.NumericRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		key, err := url.QueryUnescape(c.Params("key"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid key format"})
		}
		if !regexp.MustCompile(`^[a-zA-Z\p{Thai}]+$`).MatchString(key) {
			return c.Status(400).JSON(fiber.Map{"error": "Input contains invalid characters for a single key"})
		}
		data, err := repo.GetByKey(key)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				log.Printf("Repository error for key '%s': %v", key, err)
			}
			return c.Status(404).JSON(fiber.Map{"error": "Data not found"})
		}
		return c.JSON(data)
	}
}

// --- Decode Handler ---
type DecodeHandler struct {
	satRepo     repositories.NumericRepository
	shaRepo     repositories.NumericRepository
	meaningRepo repositories.NumberMeaningRepository
}

func NewDecodeHandler(satRepo repositories.NumericRepository, shaRepo repositories.NumericRepository, meaningRepo repositories.NumberMeaningRepository) *DecodeHandler {
	return &DecodeHandler{
		satRepo:     satRepo,
		shaRepo:     shaRepo,
		meaningRepo: meaningRepo,
	}
}

// --- Response Structs ---
type CharDetail struct {
	Char    string      `json:"char"`
	SatData interface{} `json:"sat_data"`
	ShaData interface{} `json:"sha_data"`
}
type Interpretation struct {
	PairNumber  string `json:"pair_number"`
	Meaning     string `json:"meaning"`
	Score       int    `json:"score"`
	PairType    string `json:"pair_type"`
	Description string `json:"description"`
	ColorCode   string `json:"color_code"` // Added ColorCode field
}
type InterpretationResult struct {
	TotalScore      int              `json:"total_score"`
	Interpretations []Interpretation `json:"interpretations"`
}
type ComprehensiveResult struct {
	Name             string                `json:"name"`
	RawTotalSatValue int                   `json:"raw_total_sat_value"`
	RawTotalShaValue int                   `json:"raw_total_sha_value"`
	SatResult        *InterpretationResult `json:"sat_result"`
	ShaResult        *InterpretationResult `json:"sha_result"`
	Breakdown        []CharDetail          `json:"breakdown"`
}

// --- Helper Functions ---
func formatSum(sum int) []string {
	if sum < 100 {
		return []string{fmt.Sprintf("%02d", sum)}
	} else if sum < 1000 {
		firstPart := sum / 10
		secondPart := sum % 100
		return []string{strconv.Itoa(firstPart), fmt.Sprintf("%02d", secondPart)}
	} else {
		firstPart := sum / 100
		secondPart := sum % 100
		return []string{strconv.Itoa(firstPart), fmt.Sprintf("%02d", secondPart)}
	}
}

func (h *DecodeHandler) interpretPairs(pairs []string) *InterpretationResult {
	var totalScore int
	interpretations := make([]Interpretation, 0, len(pairs))
	for _, pair := range pairs {
		meaning, err := h.meaningRepo.FindByPairNumber(pair)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				log.Printf("ERROR: Could not query meaning for pair %s: %v", pair, err)
			}
			continue
		}
		totalScore += meaning.PairPoint
		interpretations = append(interpretations, Interpretation{
			PairNumber:  pair,
			Meaning:     meaning.MiracleDetail,
			Score:       meaning.PairPoint,
			PairType:    meaning.PairType,
			Description: meaning.MiracleDesc,
			ColorCode:   getColorCodeForPairType(meaning.PairType), // Assign color code
		})
	}
	return &InterpretationResult{
		TotalScore:      totalScore,
		Interpretations: interpretations,
	}
}

func processChar(repo repositories.NumericRepository, key string) (interface{}, int) {
	if key == " " {
		return nil, 0
	}
	data, err := repo.GetByKey(key)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0
		}
		log.Printf("ERROR: Database error for key '%s': %v", key, err)
		return nil, 0
	}
	return data, data.GetValue()
}

// --- Main Handler Logic ---
func (h *DecodeHandler) DecodeName(c *fiber.Ctx) error {
	name, err := url.QueryUnescape(c.Params("name"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid name format"})
	}
	if !validateInputString(name) {
		return c.Status(400).JSON(fiber.Map{"error": "Input contains invalid characters or spacing"})
	}

	characters := []rune(name)
	breakdown := make([]CharDetail, 0, len(characters))
	var totalSat, totalSha int
	for _, charRune := range characters {
		charStr := string(charRune)
		satDataResult, satValue := processChar(h.satRepo, charStr)
		shaDataResult, shaValue := processChar(h.shaRepo, charStr)
		totalSat += satValue
		totalSha += shaValue
		if charStr != " " {
			breakdown = append(breakdown, CharDetail{Char: charStr, SatData: satDataResult, ShaData: shaDataResult})
		}
	}

	satPairs := formatSum(totalSat)
	shaPairs := formatSum(totalSha)
	satInterpretationResult := h.interpretPairs(satPairs)
	shaInterpretationResult := h.interpretPairs(shaPairs)

	finalResult := ComprehensiveResult{
		Name:             name,
		RawTotalSatValue: totalSat,
		RawTotalShaValue: totalSha,
		SatResult:        satInterpretationResult,
		ShaResult:        shaInterpretationResult,
		Breakdown:        breakdown,
	}
	return c.JSON(finalResult)
}
