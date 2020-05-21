package Driver



func Click(id string) (string, bool) {
	elem, err := utilFind(id)

	if err != nil {
		return err.Error(), false
	}
	if elem != nil {
		err = elem.Click()
		if err != nil {
			return err.Error(), false
		} else {
			return "PASS - " + id + "clicked!", true
		}

	} else {
		return "Element Not Found", false
	}

}

func Type(id string, msg string) (string, bool) {
	elem, err := utilFind(id)

	if err != nil {
		return err.Error(), false
	}
	if elem != nil {
		err = elem.SendKeys(msg)
		if err != nil {
			return err.Error(), false
		} else {
			return "PASS - " + "msg: " + msg + " : Sent to " + id, true
		}

	} else {
		return "Element Not Found", false
	}
}
