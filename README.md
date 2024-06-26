# Command Agent

Command agent is a go tool who start web api, and link endpoint with shell script.

To start agent you have to create a AgentFile.

It's a JSON file with some mandatory parameters :

```json
{
  "port" : 7890,
  "scripts_folder_path" : "/scripts/",
  "endcommands" : [
    {
      "endpoint_name" : "start",
      "scripts_files_names" : ["start.sh"],
      "method" : "POST"
    }
  ]
}
```

Before running, scripts_folder_path and all scripts_files_names have to exists. 

$ go run main.go

Start process if AgentFile is in same folder.

Basic endpoint : 
