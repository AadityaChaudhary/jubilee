package Driver

import (
	"fmt"
	"github.com/tebeka/selenium"
)

var service *selenium.Service
var wd selenium.WebDriver

func Start(seleniumPath string, driverPath string, port int, browser string) {
	opts := []selenium.ServiceOption{
		//selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(driverPath), // Specify the path to GeckoDriver in order to use Firefox.
		//selenium.Output(os.Stdout),            // Output debug information to STDERR.
	}
	selenium.SetDebug(false)
	var err error
	service, err = selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}



	caps := selenium.Capabilities{"browserName": browser}
	wd, err = selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}


	if err := wd.Get("http://play.golang.org/?simple=1"); err != nil {
		panic(err)
	}
}

func Stop() {
	wd.Quit()
	service.Stop()
}

func OpenSite(link string) (string, bool) {
	err := wd.Get(link)
	if err != nil {
		return "FAIL - " + link + " Not Opened", false
	}
	return "PASS - " + link + " Opened", true
}
