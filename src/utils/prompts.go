package utils

import (
	"fmt"
	"math/rand"
	"net"
	"rconn/src/models"
	"rconn/src/out"
	"regexp"
	"strings"
)

var hostnameRegex = regexp.MustCompile(`^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)*([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?)$`)
var userRegex = regexp.MustCompile(`^[a-zA-Z0-9._\\-]{3,64}$`)

var adjectives = []string{
	"curious", "brave", "quiet", "lucky", "mighty", "gentle",
	"clever", "proud", "nimble", "bold", "shy", "happy", "grumpy",
}

var nouns = []string{
	"salamander", "falcon", "otter", "tiger", "penguin", "eagle",
	"panther", "koala", "badger", "whale", "rhino", "dragon", "wolf",
}

func GenerateConnName() string {
	adj := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]
	return fmt.Sprintf("%s-%s", adj, noun)
}

func validateConnName(input string) error {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if len(trimmed) > 64 {
		return fmt.Errorf("name is too long (max 64 characters)")
	}
	return nil
}

func validateHost(input string) error {
	input = strings.TrimSpace(input)

	if input == "" {
		return fmt.Errorf("hostname cannot be empty")
	}

	if len(input) > 253 {
		return fmt.Errorf("hostname is too long")
	}

	if net.ParseIP(input) != nil {
		return nil
	}

	if !hostnameRegex.MatchString(input) {
		return fmt.Errorf("invalid hostname format")
	}

	return nil
}

func validateUser(input string) error {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return fmt.Errorf("username cannot be empty")
	}

	if !userRegex.MatchString(trimmed) {
		return fmt.Errorf("invalid username format")
	}

	return nil
}

func PromptRDPConnectionParams(isAdhoc bool, params models.RDPConnectionParams) (models.RDPConnectionParams, error) {
	var errs []string

	if !isAdhoc {
		if params.Name == "" {
			params.Name, _ = out.PromptInput(models.PromptInputParams{
				Label:        "Memorable name for the RDP connection",
				DefaultValue: "",
				IsPassword:   false,
			})
		}

		if params.Name == "" {
			params.Name = GenerateConnName()
		}

		if err := validateConnName(params.Name); err != nil {
			errs = append(errs, err.Error())
		}
	}

	if params.Host == "" {
		params.Host, _ = out.PromptInput(models.PromptInputParams{
			Label:        "Hostname or IP Address of the target machine",
			DefaultValue: "",
			IsPassword:   false,
		})

		if err := validateHost(params.Host); err != nil {
			out.Logger.Error("Invalid hostname: " + err.Error())
			errs = append(errs, err.Error())
		}
	}

	if params.User == "" {
		params.User, _ = out.PromptInput(models.PromptInputParams{
			Label:        "User Account to use for the RDP session",
			DefaultValue: "",
			IsPassword:   false,
		})

		if err := validateUser(params.User); err != nil {
			errs = append(errs, err.Error())
		}
	}

	if params.Password == "" {
		params.Password, _ = out.PromptInput(models.PromptInputParams{
			Label:        "Password for the User Account",
			DefaultValue: "",
			IsPassword:   true,
		})
	}

	if len(errs) > 0 {
		return params, fmt.Errorf("failed to parse input: %s", strings.Join(errs, ", "))
	}

	return params, nil
}
