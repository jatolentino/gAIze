## Backend Setup
Follow these 4 steps to setup the Go API

### 1. SDK Download 
Download the Nvidia SDK
>For [windows](https://catalog.ngc.nvidia.com/orgs/nvidia/teams/maxine/resources/maxine_windows_ar_sdk_ga/files)

>For [Linux](https://catalog.ngc.nvidia.com/orgs/nvidia/teams/maxine/resources/maxine_linux_ar_sdk/files)

Uncompress the rar and rename the extracted folder to:
> NVIDIA_AR_SDK_0.8.2.0

### 2. Folder Structure
You should get a folder structure like:
```bash
root
|-- GazeRedirect.exe
|-- api
|   `-- handler
|       |-- compress_video.go
|       |-- constants.go
|       |-- process_video.go
|       |-- receive_video.go
|       `-- send_video.go
|-- cmd
|   `-- gaizeapi
|       `-- main.go
|-- docs
|   `-- docs.txt
|-- gAIze.bat
|-- go.mod
|-- go.sum
|-- internal
|   `-- config
|       `-- paths.go
|-- main.go
|-- MAXINE-AR-SDK
|-- NVIDIA_AR_SDK_0.8.2.0
|-- pkg
|   `-- files.go
|-- readme.md
|-- results
|   `-- output
|-- scripts
|   |-- GazeRedirect.exe
|   `-- gAIze.bat
|-- tests
|   `-- test.go
`-- uploads
    `-- input

12 directories, 20 files
```

### 3. Install packages
```bash
go get github.com/gin-contrib/cors
go get github.com/gin-gonic/gin
```

### 4. Run the API
```bash
go run main.go
```