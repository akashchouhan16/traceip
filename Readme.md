# 🧬 TraceIP

A Command line tool built with **[Go](https://go.dev/ "Go Lang")** to perform geolocation lookups on IP addresses, list MX records, Name Servers (NS) and track the Canonical names for host servers.

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
## 🧬 Preview: go build & run the traceip binary
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
  A command line tool for networking utility and geolocation lookups on IP addresses.
 

  Usage:
    traceip [command]

  Available Commands:
    cname       Perform look up for canonical name (cname) for a particular host.
    completion  Generate the autocompletion script for the specified shell.
    help        Help about any command.
    mx          Perform MX Records look up for a particular host.
    ns          Perform look up for Name Servers for a particular host.
    trace       Geo locate an IPv4 or IPv6 with the trace command.

    Flags:
      -h, --help   help for traceip
      -4, --ipv4   Enforce IPv4 geo tracing only {wip}
      -6, --ipv6   Enforce IPv6 geo tracing only {wip}

    Use "traceip [command] --help" for more information about a command.
  ```

## 🧬 CMD: trace `ip_address`
**To perform geo location lookup on the IP address**
  ```shell
  $ traceip trace 45.32.39.145

    Locating IP...

    =====================================================================
    +---+--------------+-----------------------------------+
    | - | PARAMETERS   | VALUE                             |
    +---+--------------+-----------------------------------+
    | 1 | IP Address   | 45.32.39.145                      |
    | 2 | City         | Ōi                                |
    | 3 | Region       | Saitama                           |
    | 4 | Country      | JP                                |
    | 5 | Coordinates  | 35.6090,139.7302                  |
    | 6 | Timezone     | Asia/Tokyo                        |
    | 7 | Zip/Postal   | 140-8508                          |
    | 8 | Organization | AS20473 The Constant Company, LLC |
    +---+--------------+-----------------------------------+
    =====================================================================
  ```

## 🧬 CMD: mx `host_name`
**To list MX Records for a host**
  ```shell
  $ traceip mx cal.com

    =====================================================================
    MS Records for cal.com
    +---+--------------------------+----------------+
    | - | MX RECORD HOST           | MX RECORD PREF |
    +---+--------------------------+----------------+
    | 1 | aspmx.l.google.com.      |              1 |
    | 2 | alt1.aspmx.l.google.com. |              5 |
    | 3 | alt2.aspmx.l.google.com. |              5 |
    | 4 | alt4.aspmx.l.google.com. |             10 |
    | 5 | alt3.aspmx.l.google.com. |             10 |
    +---+--------------------------+----------------+
    =====================================================================
  ```

## 🧬 CMD: ns `host_name`
**To list the Name Servers for a host**
  ```shell
  $ traceip ns cal.com

    =====================================================================
    Name Server(s) for cal.com
    +---+---------------------------+
    | - | NAME SERVER               |
    +---+---------------------------+
    | 1 | linda.ns.cloudflare.com.  |
    | 2 | carter.ns.cloudflare.com. |
    +---+---------------------------+
    =====================================================================
  ```

## 🧬 CMD: cname `host_name`
**To view the canonical name (CNAME) for a host**
  ```shell
  $ traceip cname telnet.io
    =====================================================================
    canonical name (CNAME) for telnet.io
    +---+----------------+
    | - | CANONICAL NAME |
    +---+----------------+
    | 1 | park.io.       |
    +---+----------------+
    =====================================================================
  ```
---

## 🔖 Contributing
If you wish to contribute, feel free to reach out with ideas or questions for **traceip** or use the issues tab.
- **[akash.c1500@gmail.com](mailto:akash.c1500@gmail.com "Akash's Gmail")**

### License
Copyright (c) **Akash Chouhan**. All rights reserved. Released under the **[MIT License](https://github.com/akashchouhan16/traceip/blob/master/LICENSE "View License")**
