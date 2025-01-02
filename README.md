# TouchBarWeather

TouchBarWeather is a macOS application that integrates weather updates into the MacBook Touch Bar. It displays real-time weather information using the OpenWeather API, providing a convenient and aesthetically pleasing way to check weather details.

## Features

- **Live Weather Updates:** Displays temperature, weather conditions, and location directly on the Touch Bar.
- **Menu Bar Integration:** Adds a weather icon in the menu bar with an option to refresh weather data.
- **Customizable Refresh Interval:** Updates every 5 minutes by default.

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/touchbarweather.git
    cd touchbarweather
    ```

2. Install dependencies:
    - [menuet](https://github.com/caseymrm/menuet)
    - [macdriver](https://github.com/progrium/macdriver)

   Use `go get` to fetch the required modules:
    ```bash
    go get github.com/caseymrm/menuet
    go get github.com/progrium/macdriver
    ```

3. Replace the placeholder `YOUR_API_KEY` in `touchbarweather.go` with your OpenWeather API key.

4. Build and run the application:
    ```bash
    go run touchbarweather.go
    ```

## Usage

- **Touch Bar:** Displays the current weather as a button on the Touch Bar.
- **Menu Bar:** Click the weather icon in the menu bar to refresh weather data manually.

## Configuration

- **Default Location:** The default coordinates are set to latitude `37.7749` and longitude `-122.4194` (San Francisco). Update these in the `touchbarweather.go` file to your desired location.

- **API Key:** Obtain a free API key from [OpenWeather](https://openweathermap.org/) and update it in the code.

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.

## Contributions

Contributions, issues, and feature requests are welcome! Feel free to open a pull request or submit an issue.
