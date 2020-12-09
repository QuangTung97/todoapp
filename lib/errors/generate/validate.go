package generate

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"sort"
	"strconv"
	"strings"
)

func lowerCaseFirstLetter(s string) string {
	return strings.ToLower(s[:1]) + s[1:]
}

func validateError(name string, info ErrorInfo) error {
	if info.RPCStatus <= 0 || info.RPCStatus > 16 {
		return fmt.Errorf("invalid rpc status '%d'", info.RPCStatus)
	}
	rpcCode := fmt.Sprintf("%02d", info.RPCStatus)
	if !strings.HasPrefix(info.Code, rpcCode) {
		return fmt.Errorf("code must be prefix with '%s'", rpcCode)
	}
	statusCode := codes.Code(info.RPCStatus).String()
	statusCode = lowerCaseFirstLetter(statusCode)
	if !strings.HasPrefix(name, statusCode) {
		return fmt.Errorf("error name must prefix with '%s'", statusCode)
	}
	return nil
}

// Validate validates tags
func Validate(tags map[string]ErrorMap) error {
	codeSet := make(map[string]struct{})
	for _, errorMap := range tags {
		for errorName, info := range errorMap {
			_, existed := codeSet[info.Code]
			if existed {
				return fmt.Errorf("code '%s' duplicated", info.Code)
			}
			codeSet[info.Code] = struct{}{}

			err := validateError(errorName, info)
			if err != nil {
				return err
			}

		}
	}
	return nil
}

func findNextErrorCode(rpcStatus uint32, inputCodes []int) string {
	codeNums := make([]int, 0)
	codeSet := make(map[int]struct{})
	for _, c := range inputCodes {
		_, existed := codeSet[c]
		if !existed {
			codeNums = append(codeNums, c)
			codeSet[c] = struct{}{}
		}
	}
	sort.Ints(codeNums)

	prev := 0
	for _, c := range codeNums {
		if c > prev {
			return fmt.Sprintf("%02d%02d", rpcStatus, prev)
		}
		prev++
	}
	return fmt.Sprintf("%02d%02d", rpcStatus, prev)
}

// NextErrorCodeForRPCStatus finds the next code for rpc status
func NextErrorCodeForRPCStatus(tags map[string]ErrorMap, rpcStatus uint32) (string, error) {
	codeNums := make([]int, 0)
	for _, errorMap := range tags {
		for _, info := range errorMap {
			if info.RPCStatus == rpcStatus {
				n, err := strconv.ParseInt(info.Code[2:], 10, 64)
				if err != nil {
					return "", err
				}
				codeNums = append(codeNums, int(n))
			}
		}
	}

	return findNextErrorCode(rpcStatus, codeNums), nil
}
