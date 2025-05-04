# hackerhotel-ical

A lightweight Go application that generates an iCalendar (.ics) file for the HackerHotel event. This tool allows attendees to easily import the event schedule into their preferred calendar applications.

## Features

* Fetches the latest HackerHotel event schedule and filters the location (pretalx room)
* Facilitates easy import into calendar applications like Google Calendar, Apple Calendar, Outlook, and more.
* Used in Dakboard for per room schedules

## Installation

### Prerequisites

* Go 1.20 or higher

### Clone the Repository

```bash
git clone https://github.com/JeitoBV/hackerhotel-ical.git
cd hackerhotel-ical
```

### Build the Application

```bash
go build
```

## Usage

After building the application, run it using:

```bash
./hackerhotel-ical --help
Usage: hackerhotel-ical [--schedule <url>] --token <token> [--listen <bind address>]

Options:
  --schedule <url>, -s <url>
                         URL to the ical schedule [default: https://pretalx.hackerhotel.nl/2025/schedule/export/schedule.ics, env: SCHEDULE_URL]
  --token <token>, -t <token>
                         Authentication Token [env: TOKEN]
  --listen <bind address>, -l <bind address>
                         Port to listen on [default: 0.0.0.0:5000, env: LISTEN_ADDRESS]
  --help, -h             display this help and exit
```

## Example systemd service

Extract and copy the binary to `/usr/local/bin/`
```
sudo tar xvfz release.tar.gz -C /usr/local/bin hackerhotel-ical
sudo chmod 755 /usr/local/bin/hackerhotel-ical
```

Create the system service file:
```
sudo vi /etc/systemd/system/conference-ical.service
```

```
[Unit]
Description=Conference ICAL configurator
ConditionPathExists=/usr/local/bin/hackerhotel-ical
After=network.target

[Service]
Type=simple
User=nobody
Group=users

WorkingDirectory=/usr/local/bin
Environment="TOKEN=APITOKENFROMPRETALX"
Environment="SCHEDULE_URL=https://conference.wireshark.org/sharkfest-25-us-2024/schedule/export/schedule.ics"
ExecStart=/usr/local/bin/hackerhotel-ical
Restart=on-failure
RestartSec=10

StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=hackerhotel-ical

[Install]
WantedBy=multi-user.target
```

Reload systemd, enable and start conference-ical
```
systemctl daemon-reload
systemctl enable conference-ical
systemctl start conference-ical
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

*Note: This project is currently in its early stages. Features and functionalities are subject to change.*

---

