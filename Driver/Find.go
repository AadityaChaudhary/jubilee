package Driver

import (
	"errors"
	"github.com/tebeka/selenium"
	"strings"
)

func Find(id string) (string, bool) {
	_, err := utilFind(id)

	if err == nil {
		return "PASS - Element Found", true
	}

	if err.Error() == "invalid syntax" {
		return "FAILED - INVALID SYNTAX, REFERENCE TO WEB ELEMENT SHOULD BEGIN WITH link=, xpath=, or id=", false
	} else {
		return "FAILED - Element Not Found", false
	}


}

func utilFind(id string) (selenium.WebElement, error) {
	if strings.HasPrefix(id, "id=") {
		return wd.FindElement(selenium.ByID,strings.TrimPrefix(id,"id="))
	} else if strings.HasPrefix(id, "xpath=") {
		return wd.FindElement(selenium.ByXPATH,strings.TrimPrefix(id,"xpath="))
	} else if strings.HasPrefix(id, "link=") {
		return wd.FindElement(selenium.ByLinkText,strings.TrimPrefix(id,"link="))
	} else {
		return nil, errors.New("invalid syntax")
	}
}

func WaitForElem(id string, timeout int) (string, bool) {
	for x := 0; x < timeout; x++ {
		_, err := utilFind(id)
		if err == nil {
			return "PASS - " + id + " Found after " + string(x) + " ms", true
		}
	}
	return "FAILED - " + id + " never found", false
}
