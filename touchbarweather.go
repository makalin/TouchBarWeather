package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"

    "github.com/progrium/darwinkit/macos/appkit"
    "github.com/progrium/darwinkit/macos/foundation"
    "github.com/progrium/darwinkit/objc"
)

const (
    apiKey    = "YOUR_API_KEY" // Replace with your OpenWeather API key
    baseURL   = "http://api.openweathermap.org/data/2.5/weather"
    latitude  = 37.7749        // Default latitude
    longitude = -122.4194      // Default longitude
)

type WeatherData struct {
    Main struct {
        Temp     float64 `json:"temp"`
        Humidity int     `json:"humidity"`
    } `json:"main"`
    Weather []struct {
        Main        string `json:"main"`
        Description string `json:"description"`
    } `json:"weather"`
    Name string `json:"name"`
}

type TouchBarWeather struct {
    app        *appkit.NSApplication
    touchBar   objc.Object
    weatherBtn objc.Object
    statusItem objc.Object
    timer      *time.Ticker
}

func NewTouchBarWeather() *TouchBarWeather {
    w := &TouchBarWeather{}
    w.setupTouchBar()
    w.setupStatusBar()
    w.timer = time.NewTicker(5 * time.Minute)
    return w
}

func (w *TouchBarWeather) setupTouchBar() {
    // Create a button for the Touch Bar
    w.weatherBtn = appkit.NSButton_New()
    w.weatherBtn.Call("setBezelStyle:", appkit.NSBezelStyleRounded)
    w.weatherBtn.Call("setTitle:", foundation.String("Loading weather..."))

    // Create and configure the Touch Bar
    w.touchBar = objc.Get("NSTouchBar").Alloc().Init()
    w.touchBar.Call("setIdentifier:", foundation.String("com.weather.touchbar"))

    itemID := foundation.String("weatherItem")
    touchBarItem := objc.Get("NSCustomTouchBarItem").Alloc().Init()
    touchBarItem.Call("setIdentifier:", itemID)
    touchBarItem.Call("setView:", w.weatherBtn)

    w.touchBar.Call("setDefaultItemIdentifiers:", foundation.Array(itemID))
    w.touchBar.Call("setTemplateItems:", foundation.Set(touchBarItem))
}

func (w *TouchBarWeather) setupStatusBar() {
    statusBar := appkit.NSStatusBar_System()
    w.statusItem = statusBar.Call("statusItemWithLength:", appkit.NSVariableStatusItemLength)

    // Create a menu
    menu := appkit.NSMenu_New()

    refreshItem := appkit.NSMenuItem_New()
    refreshItem.Call("setTitle:", foundation.String("Refresh"))
    refreshItem.Call("setAction:", objc.Sel("refreshWeather:"))
    menu.Call("addItem:", refreshItem)

    quitItem := appkit.NSMenuItem_New()
    quitItem.Call("setTitle:", foundation.String("Quit"))
    quitItem.Call("setAction:", objc.Sel("terminate:"))
    menu.Call("addItem:", quitItem)

    w.statusItem.Call("setMenu:", menu)
    w.statusItem.Call("button").Call("setTitle:", foundation.String("☁️"))
}

func (w *TouchBarWeather) updateWeather() {
    url := fmt.Sprintf("%s?lat=%f&lon=%f&appid=%s&units=metric", baseURL, latitude, longitude, apiKey)

    resp, err := http.Get(url)
    if err != nil {
        w.updateDisplay("Weather update failed")
        return
    }
    defer resp.Body.Close()

    var weather WeatherData
    if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
        w.updateDisplay("Failed to parse weather data")
        return
    }

    display := fmt.Sprintf("\ud83c\udf21\ufe0f %.1f°C | %s | %s", 
        weather.Main.Temp,
        weather.Weather[0].Main,
        weather.Name,
    )
    w.updateDisplay(display)
}

func (w *TouchBarWeather) updateDisplay(text string) {
    foundation.Dispatch(func() {
        w.weatherBtn.Call("setTitle:", foundation.String(text))
    })
}

func (w *TouchBarWeather) Run() {
    w.updateWeather()

    go func() {
        for range w.timer.C {
            w.updateWeather()
        }
    }()

    delegate := appkit.DefaultDelegateClass.Instantiate("AppDelegate")
    delegate.AddMethod("refreshWeather:", func(_ objc.Object) {
        w.updateWeather()
    })

    w.app = appkit.NSApp()
    w.app.Call("setDelegate:", delegate)
    w.app.Call("setActivationPolicy:", appkit.NSApplicationActivationPolicyAccessory)
    w.app.Run()
}

func main() {
    app := NewTouchBarWeather()
    app.Run()
}
