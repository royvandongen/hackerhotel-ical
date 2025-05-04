# hackerhotel-ical

**hackerhotel-ical** is a lightweight Go application that generates an iCalendar (.ics) file for the HackerHotel event. It allows attendees to easily import the event schedule into their preferred calendar applications.

## ✨ Features

- Fetches the latest HackerHotel schedule from Pretalx
- Filters schedule by room/location
- Supports import into Google Calendar, Apple Calendar, Outlook, and more
- Compatible with Dakboard and other digital signage solutions (per-room schedules)

## 🛠️ Installation

### Prerequisites

- Go 1.20 or higher

### Clone the Repository

```bash
git clone https://github.com/JeitoBV/hackerhotel-ical.git
cd hackerhotel-ical
```

### Build the Application

```bash
go build
```

## 🚀 Usage

After building the application, run it with:

```bash
./hackerhotel-ical --help
```

### Example

```bash
./hackerhotel-ical --token mysecrettoken
```

### Options

```
Usage: hackerhotel-ical [--schedule <url>] --token <token> [--listen <bind address>]

Options:
  --schedule <url>, -s <url>
      URL of the Pretalx schedule
      [default: https://pretalx.hackerhotel.nl/2025/schedule/export/schedule.ics]
      [env: SCHEDULE_URL]

  --token <token>, -t <token>
      Required authentication token
      [env: TOKEN]

  --listen <bind address>, -l <bind address>
      Address and port to bind to
      [default: 0.0.0.0:5000]
      [env: LISTEN_ADDRESS]

  --help, -h
      Show this help message and exit
```

## 🤝 Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

> **Note:** This project is currently in its early stages. Features and functionality may change over time.
