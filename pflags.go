package pflags

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/PGo-Projects/output"
	"github.com/PGo-Projects/pflags/internal/config"
)

var (
	feature = "feature"
	array   = "array"
	hashmap = "hashmap"

	namePattern = regexp.MustCompile(`[a-zA-Z][-_a-zA-Z\d]*`)

	headerPatterns = map[string]*regexp.Regexp{
		feature: regexp.MustCompile(`\[\[([-_a-zA-Z]+)\]\](\.default)?`),
		array:   regexp.MustCompile(`\[([-_a-zA-Z]+)\]`),
		hashmap: regexp.MustCompile(`{([-_a-zA-Z]+)}`),
	}

	valuePatterns = []*regexp.Regexp{
		regexp.MustCompile(`(-?\d+)`),                    // int64
		regexp.MustCompile(`(-?(0|([1-9]\d*))?\.\d+)`),   // float
		regexp.MustCompile(`"([^"\\]*(?:\\.[^"\\]*)*)"`), // string
		regexp.MustCompile(`'([^']?)'`),                  // char
		regexp.MustCompile(`((false)|(true))`),           // bool
	}
)

type header struct {
	kind    string
	name    string
	info    []string
	lineNum int
}

func SetDebug(debugMode bool) {
	output.DEBUG = debugMode
}

func Parse(filename string, features ...string) (*config.Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cfg := config.New(len(features) > 1)
	scanner := NewScanner(file)
	err = parseFeatures(features, scanner, cfg)
	if err != nil {
		output.DebugErrorln(err)
	}

	return cfg, err
}

func parseFeatures(features []string, scanner *scanner, cfg *config.Config) error {
	var reachedEOF bool
	var err error

	parsingFeature := false
	for scanner.Scan() {
		line, lineNum := scanner.Text()
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "[") || strings.HasPrefix(line, "{") {
			heading, isHeader := parseForHeader(line, lineNum)
			if isHeader {
				if heading.kind == feature && contains(features, heading.name) {
					parseFeatureHeader(cfg, heading)
					parsingFeature = true
					continue
				}
			for_loop:
				for parsingFeature {
					switch heading.kind {
					case array:
						heading, reachedEOF, err = parseArray(scanner, cfg, heading.name)
					case hashmap:
						heading, reachedEOF, err = parseHashmap(scanner, cfg, heading.name)
					case feature:
						if contains(features, heading.name) {
							parseFeatureHeader(cfg, heading)
							break for_loop
						} else {
							parsingFeature = false
						}
					default:
						return fmt.Errorf("Header of type %s on line %d can not be handled", heading.kind, lineNum)
					}

					if err != nil || reachedEOF {
						return err
					}
				}

			}
		}
	}
	return nil
}

func parseArray(scanner *scanner, cfg *config.Config, name string) (*header, bool, error) {
	for scanner.Scan() {
		line, lineNum := scanner.Text()
		line = strings.TrimSpace(line)
		if line != "" {
			if h, isHeader := parseForHeader(line, lineNum); isHeader {
				return h, false, nil
			}

			elems := strings.Split(line, " ")
			value, ok := getValue(elems[0])
			if len(elems) > 1 || !ok {
				return nil, false, fmt.Errorf("%s on line %d is not a valid value", line, lineNum)
			}
			cfg.Array.Add(name, value)
		}
	}
	return nil, true, nil
}

func parseHashmap(scanner *scanner, cfg *config.Config, name string) (*header, bool, error) {
	for scanner.Scan() {
		line, lineNum := scanner.Text()
		line = strings.TrimSpace(line)
		if line != "" {
			if h, isHeader := parseForHeader(line, lineNum); isHeader {
				return h, false, nil
			}

			elems := strings.Split(line, " ")
			match := namePattern.FindString(elems[0])
			if elems[0] != match {
				return nil, false, fmt.Errorf("%s on line %d is not a valid name", elems[0], lineNum)
			}
			if elems[1] != "=" {
				return nil, false, fmt.Errorf("%s on line %d is not a valid delimiter", elems[1], lineNum)
			}

			value, ok := getValue(elems[2])
			if !ok {
				return nil, false, fmt.Errorf("%s on line %d is not a valid value", elems[2], lineNum)
			}
			cfg.HashMap.Add(name, elems[0], value)
		}
	}
	return nil, true, nil
}

func parseForHeader(line string, lineNum int) (*header, bool) {
	for headerType, headerPattern := range headerPatterns {
		matches := headerPattern.FindStringSubmatch(line)
		if len(matches) > 0 && matches[0] == line {
			h := &header{
				kind:    headerType,
				name:    matches[1],
				info:    matches[2:],
				lineNum: lineNum,
			}
			return h, true
		}
	}
	return nil, false
}

func parseFeatureHeader(cfg *config.Config, heading *header) {
	if len(heading.info) > 1 && heading.info[0] == ".default" {
		cfg.DefaultFeature = heading.name
	}
	output.DebugStringln(
		fmt.Sprintf("Parsing feature %s on line %d\n", heading.name, heading.lineNum),
		output.BLUE,
	)
}

func getValue(s string) (string, bool) {
	for _, pattern := range valuePatterns {
		matches := pattern.FindStringSubmatch(s)
		if len(matches) > 0 && matches[0] == s {
			return matches[1], true
		}
	}
	return "", false
}

func contains(array []string, s string) bool {
	for _, elem := range array {
		if elem == s {
			return true
		}
	}
	return false
}
