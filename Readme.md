# 🧬 TraceIP

A Command line tool built with **[Go](https://go.dev/ "Go Lang")** to perform geolocation lookups for IP addresses. 

Retrieve information such as host, country, region, city, latitude, longitude and associated organisation with an IP address.

```bash
$ traceip
```

## 📜 Quick Setup

> **Learn more about go workspace and installation [here](https://go.dev/doc/tutorial/compile-install "Visit go.dev docs").**

- Clone the GitHub repo: 
  - ```bash
    $ git clone https://github.com/akashchouhan16/traceip.git
    ```
  - ```bash
    $ cd traceip
    ```
- Install and compile traceip binary:
  - ```go
    $ go install && go build
    ```
- Run **traceip**:
  - ```
    $ traceip
    ```
* **System Prerequisites**: **Go 1.17** or above installed.


---
## 🐋 Run with Docker 
- Pull the base image for **traceip** from docker hub.
  - ```shell
    $ docker pull akashchouhan16/traceip
    ```
- Run as docker container.
  - ```shell
    $ docker run akashchouhan16/traceip
    ```
- **System Prerequisite:** **Docker runtime** installed locally.
  > **Note**: Incase any issue is faced with the setup, make sure the `GOPATH` is set to pwd and `go version` is updated. Refer the **[Go compilation guide](https://go.dev/doc/tutorial/compile-install "Visit go.dev docs").**

---
## 🧬 Preview: GO build & Run the traceip binary
  ```shell
  $ go build && traceip

     ======================================================
    | ████████╗██████╗  █████╗  ██████╗███████╗ ██╗██████╗  |
    | ╚══██╔══╝██╔══██╗██╔══██╗██╔════╝██╔════╝ ██║██╔══██╗ |
    |    ██║   ██████╔╝███████║██║     █████╗   ██║██████╔╝ |
    |    ██║   ██╔══██╗██╔══██║██║     ██╔══╝   ██║██╔═══╝  |
    |    ██║   ██║  ██║██║  ██║╚██████╗███████╗ ██║██║      |
    |    ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝╚══════╝ ╚═╝╚═╝      |
     ======================================================
    A command line tool for geolocation lookups on IP addresses.


    Usage:
      traceip [command]

    Available Commands:
      completion  Generate the autocompletion script for the specified shell
      help        Help about any command
      trace       Geo locate an IPv4 or IPv6 with the trace command.

    Flags:
      -h, --help   help for traceip
      -4, --ipv4   Enforce IPv4 geo tracing only {wip}
      -6, --ipv6   Enforce IPv6 geo tracing only {wip}

    Use "traceip [command] --help" for more information about a command.
  ```
  ![build-and-run Preview](https://github.com/akashchouhan16/traceip/assets/56465610/b17cf224-f949-44b2-9487-e7cb6e830c7c)

## 🧬 Preview: traceip trace <IPv4 | IPv6>
  ```shell
  $ traceip trace 45.32.39.145

  Locating IP...
  ==============================
        - IP: 45.32.39.145
        - CITY: Ōi
        - REGION: Saitama
        - COUNTRY :JP
        - LOCATION :35.6090,139.7302
        - TIMEZONE:Asia/Tokyo
        - POSTAL: 140-8508
        - Org: AS20473 The Constant Company, LLC
  ==============================
  ```
  <!-- ![trace-cmd Preview](https://github.com/akashchouhan16/traceip/assets/56465610/6fc1bb7c-ac97-4596-8a39-8e8453748265) -->
---

## 🔖 Contributing
### Author
Reach out with questions or ideas for traceip.
- **[akash.c1500@gmail.com](mailto:akash.c1500@gmail.com)**

### License
Copyright (c) **Akash Chouhan**. All rights reserved. Released under the **[MIT License](https://github.com/akashchouhan16/traceip/blob/master/LICENSE "View License")**
