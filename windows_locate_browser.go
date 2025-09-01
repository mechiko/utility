package utility

import (
	"os"
	"os/exec"
)

// ChromeExecutable returns a string which points to the preferred Chrome
// executable file.
var ChromeExecutable = LocateChrome

// LocateChrome returns a path to the Chrome binary, or an empty string if
// Chrome installation is not found.
func LocateChrome() string {
	pathsChrome := []string{
		os.Getenv("LocalAppData") + "/Google/Chrome/Application/chrome.exe",
		os.Getenv("ProgramFiles") + "/Google/Chrome/Application/chrome.exe",
		os.Getenv("ProgramFiles(x86)") + "/Google/Chrome/Application/chrome.exe",
		os.Getenv("LocalAppData") + "/Chromium/Application/chrome.exe",
		os.Getenv("ProgramFiles") + "/Chromium/Application/chrome.exe",
		os.Getenv("ProgramFiles(x86)") + "/Chromium/Application/chrome.exe",
		os.Getenv("ProgramFiles(x86)") + "/Microsoft/Edge/Application/msedge.exe",
		os.Getenv("ProgramFiles") + "/Microsoft/Edge/Application/msedge.exe",
	}
	for _, path := range pathsChrome {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}
		return path
	}
	return ""
}

func LocateYandex() string {
	pathsChrome := []string{
		os.Getenv("LocalAppData") + "/Yandex/YandexBrowser/Application/browser.exe",
	}
	for _, path := range pathsChrome {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}
		return path
	}
	return ""
}

func LocateFox() string {
	pathsChrome := []string{
		os.Getenv("ProgramFiles") + "/Firefox Developer Edition/firefox.exe",
		os.Getenv("ProgramFiles(x86)") + "/Firefox Developer Edition/firefox.exe",
		os.Getenv("ProgramFiles") + "/Mozilla Firefox/firefox.exe",
		os.Getenv("ProgramFiles(x86)") + "/Mozilla Firefox/firefox.exe",
	}
	for _, path := range pathsChrome {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}
		return path
	}
	return ""
}

func LocateEdge() string {
	pathsChrome := []string{
		os.Getenv("ProgramFiles(x86)") + "/Microsoft/Edge/Application/msedge.exe",
	}
	for _, path := range pathsChrome {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}
		return path
	}
	return ""
}

func StartBrowser(browser Browser, url string) (err error) {
	switch browser {
	case Chrome:
		args := []string{}
		appUrl := "--app=" + url
		args = append(args, appUrl)
		cmd := exec.Command(LocateChrome(), args...)
		err = cmd.Start()
	case Firefox:
		argsFox := []string{}
		appUrlFox := url
		argsFox = append(argsFox, appUrlFox)
		cmd := exec.Command(LocateFox(), argsFox...)
		err = cmd.Start()
	case Edge:
		args := []string{}
		appUrl := "--app=" + url
		args = append(args, appUrl)
		cmd := exec.Command(LocateEdge(), args...)
		err = cmd.Start()
	case Yandex:
		args := []string{}
		appUrl := "--app=" + url
		args = append(args, appUrl)
		cmd := exec.Command(LocateYandex(), args...)
		err = cmd.Start()
	default:
		err = OpenHttpLinkInShell(url)
	}
	return err
}
